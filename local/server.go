package local

import (
	"context"
	"errors"
	"reflect"

	"github.com/vgough/sequin"
	"github.com/vgough/sequin/registry"
)

type Server struct {
	codec sequin.Codec
}

var _ sequin.Runtime = &Server{}

func NewServer() *Server {
	return &Server{
		codec: &DefaultCodec{},
	}
}

func (s *Server) Exec(ep *registry.Endpoint, args []reflect.Value) []reflect.Value {
	// Marshal the arguments.
	data, err := s.encodeValues(args)
	if err != nil {
		return ep.MakeError(err)
	}

	// call backend
	results, err := s.exec(ep.Name, data)
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

func (s *Server) exec(name string, args [][]byte) ([][]byte, error) {
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

	out := ep.Exec(in)
	return s.encodeValues(out)
}

func (s *Server) encodeValues(values []reflect.Value) ([][]byte, error) {
	data := make([][]byte, len(values))
	for i, v := range values {
		if v.Type() == registry.ContextType {
			continue
		}
		d, err := s.codec.Encode(v)
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
		val, err := s.codec.Decode(d, types[i])
		if err != nil {
			return nil, err
		}
		values[i] = val
	}
	return values, nil
}
