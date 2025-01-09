package local

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"reflect"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/sync/singleflight"

	"github.com/vgough/sequin"
	"github.com/vgough/sequin/internal"
	"github.com/vgough/sequin/registry"
	"github.com/vgough/sequin/storage"
)

var encodingKey []byte = []byte("sequin")
var requestIDMD = internal.MDKey[string]{}

type Server struct {
	sf      singleflight.Group
	storage storage.Store

	mu sync.Mutex

	// maps from request id to current or recent requests.
	cache map[string]*requestState
}

type requestState struct {
	requestID string
	results   [][]byte
}

var _ sequin.Runtime = &Server{}

type ServerOptions func(*Server)

func WithStorage(storage storage.Store) ServerOptions {
	return func(s *Server) {
		s.storage = storage
	}
}

func NewServer(opts ...ServerOptions) *Server {
	s := &Server{
		cache: make(map[string]*requestState),
	}
	for _, opt := range opts {
		opt(s)
	}
	if s.storage == nil {
		s.storage = storage.NewTestObjectStore()
	}
	return s
}

func (s *Server) Exec(ep *registry.Endpoint, args []reflect.Value) []reflect.Value {
	// Marshal the arguments.
	data, err := s.encodeValues(args)
	if err != nil {
		return ep.MakeError(err)
	}

	// create unique id from data.
	ctx := ep.GetContext(args)
	parentID := requestIDMD.Get(ctx)
	if opt, ok := ep.Metadata[internal.GlobalIDGen]; ok {
		if boolVal, ok := opt.(bool); ok && boolVal {
			parentID = ""
		}
	}

	requestID := computeUniqueID(parentID, data)

	results, err := s.run(ctx, requestID, ep, data)
	if err != nil {
		return ep.MakeError(err)
	}

	// decode results
	out, err := s.decodeValues(results, ep.OutputTypes)
	if err != nil {
		return ep.MakeError(err)
	}

	return out
}

func (s *Server) run(ctx context.Context, requestID string,
	ep *registry.Endpoint, data [][]byte) ([][]byte, error) {

	// use singleflight to avoid duplicate requests.
	res := s.sf.DoChan(requestID, func() (interface{}, error) {
		s.mu.Lock()
		state, ok := s.cache[requestID]
		s.mu.Unlock()
		if ok {
			return state, nil
		}

		results, err := s.exec(ep.Name, requestID, data)
		if err != nil {
			return nil, err
		}
		state = &requestState{requestID: requestID, results: results}
		s.mu.Lock()
		defer s.mu.Unlock()
		// persist results.
		s.cache[requestID] = state

		return state, nil
	})

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case res := <-res:
		if res.Err != nil {
			return nil, res.Err
		}
		state := res.Val.(*requestState)
		return state.results, nil
	}
}

func (s *Server) exec(name string, requestID string, args [][]byte) ([][]byte, error) {
	ep := registry.GetEndpoint(name)
	if ep == nil {
		return nil, errors.New("unknown function: " + name)
	}

	in, err := s.decodeValues(args, ep.InputTypes)
	if err != nil {
		return nil, err
	}
	if ep.ContextIndex >= 0 {
		in[ep.ContextIndex] = reflect.ValueOf(context.Background())
	}

	// TODO: chain to incoming context.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx = sequin.WithRuntime(ctx, s)
	ctx = requestIDMD.Set(ctx, requestID)
	ep.SetContext(ctx, in)

	out := ep.Exec(in)
	return s.encodeValues(out)
}

func computeUniqueID(parentID string, data [][]byte) string {
	hash := hmac.New(sha256.New, encodingKey)

	var tmp [10]byte
	hash.Write(internal.EncodeVarint(len(parentID), tmp))
	hash.Write([]byte(parentID))

	for _, d := range data {
		hash.Write(internal.EncodeVarint(len(d), tmp))
		hash.Write(d)
	}
	digest := hash.Sum(nil)
	return base64.RawStdEncoding.EncodeToString(digest[:])
}

func (s *Server) encodeValues(values []reflect.Value) ([][]byte, error) {
	data := make([][]byte, len(values))
	for i, v := range values {
		if v.Type() == registry.ContextType {
			continue
		}
		d, err := internal.Encode(v)
		if err != nil {
			return nil, err
		}
		data[i] = d
	}
	return data, nil
}

func (s *Server) decodeValues(data [][]byte, types []reflect.Type) ([]reflect.Value, error) {
	values := make([]reflect.Value, len(data))
	for i, d := range data {
		if types[i] == registry.ContextType {
			continue
		}
		val, err := internal.Decode(d, types[i])
		if err != nil {
			return nil, err
		}
		values[i] = val
	}
	return values, nil
}
