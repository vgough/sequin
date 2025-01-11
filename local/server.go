package local

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/sync/singleflight"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/vgough/sequin"
	sequinv1 "github.com/vgough/sequin/gen/sequin/v1"
	"github.com/vgough/sequin/internal"
	"github.com/vgough/sequin/registry"
)

var encodingKey []byte = []byte("sequin")
var requestIDMD = internal.MDKey[string]{}

type Server struct {
	sf singleflight.Group

	storage Store
}

var _ sequin.Runtime = &Server{}

type ServerOptions func(*Server)

func WithStorage(storage Store) ServerOptions {
	return func(s *Server) {
		s.storage = storage
	}
}

func NewServer(opts ...ServerOptions) *Server {
	s := &Server{}
	for _, opt := range opts {
		opt(s)
	}
	if s.storage == nil {
		s.storage = NewMemStore()
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

func (s *Server) runInternal(ctx context.Context, requestID string,
	ep *registry.Endpoint, data [][]byte) ([][]byte, error) {

	funcoOp := sequinv1.FuncOperation{
		Name: ep.Name,
		Args: data,
	}
	any, err := anypb.New(&funcoOp)
	if err != nil {
		return nil, err
	}

	op := &sequinv1.Operation{
		RequestId: requestID,
		Detail:    any,
	}

	created, err := s.storage.AddOperation(ctx, op)
	if err != nil {
		return nil, err
	}

	var state *OpState
	if !created {
		// get existing state.
		state, err = s.storage.GetState(ctx, requestID)
		if err != nil {
			return nil, err
		}
		if len(state.Results) > 0 {
			return nil, ErrOperationAlreadyFinished
		}
	}
	if state == nil {
		state = &OpState{}
	}

	success, results, err := s.exec(ep.Name, requestID, data, state)
	if err != nil {
		return nil, err
	}

	if success {
		state.Results = results
		err = s.storage.SetState(ctx, requestID, state)
		if err != nil {
			return results, fmt.Errorf("failed to set terminal state: %w", err)
		}
	}

	return results, nil
}

func (s *Server) run(ctx context.Context, requestID string,
	ep *registry.Endpoint, data [][]byte) ([][]byte, error) {

	// use singleflight to avoid duplicate requests.
	res := s.sf.DoChan(requestID, func() (interface{}, error) {
		return s.runInternal(ctx, requestID, ep, data)
	})

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case res := <-res:
		if res.Err != nil {
			return nil, res.Err
		}
		results := res.Val.([][]byte)
		return results, nil
	}
}

func (s *Server) exec(name string, requestID string, args [][]byte,
	state *OpState) (bool, [][]byte, error) {

	ep := registry.GetEndpoint(name)
	if ep == nil {
		return false, nil, errors.New("unknown function: " + name)
	}

	in, err := s.decodeValues(args, ep.InputTypes)
	if err != nil {
		return false, nil, err
	}

	// TODO: chain to incoming context.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rt := NewRuntime(s.storage, requestID, state)
	ctx = internal.WithOpRuntime(ctx, rt)

	ctx = sequin.WithRuntime(ctx, s)
	ctx = requestIDMD.Set(ctx, requestID)
	ep.SetContext(ctx, in)

	out := ep.Exec(in)
	data, err := s.encodeValues(out)
	if err != nil {
		return false, nil, err
	}
	success := ep.GetError(out) == nil
	return success, data, nil
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
