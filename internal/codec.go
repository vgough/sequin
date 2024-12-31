package internal

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"reflect"

	"google.golang.org/protobuf/proto"
)

type vtProto interface {
	MarshalVT() ([]byte, error)
	UnmarshalVT([]byte) error
}

var protoType = reflect.TypeFor[proto.Message]()
var vtProtoType = reflect.TypeFor[vtProto]()
var contextType = reflect.TypeFor[context.Context]()

type DefaultCodec struct {
}

func (c *DefaultCodec) Encode(v reflect.Value) ([]byte, error) {
	at := v.Type()
	switch {
	case at.Implements(vtProtoType):
		return v.Interface().(vtProto).MarshalVT()
	case at.Implements(protoType):
		return proto.Marshal(v.Interface().(proto.Message))
	case v.IsZero():
		return []byte{}, nil
	default:
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		if err := enc.Encode(v.Interface()); err != nil {
			return nil, fmt.Errorf("gob encode failed on type %q: %w", at.Name(), err)
		}
		return buf.Bytes(), nil
	}
}

func (c *DefaultCodec) Decode(data []byte, vt reflect.Type) (reflect.Value, error) {
	typ := vt
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	val := reflect.New(typ)

	switch {
	case len(data) == 0:
	case vt.Implements(vtProtoType):
		if err := val.Interface().(vtProto).UnmarshalVT(data); err != nil {
			return val, fmt.Errorf("proto unmarshal failed on type %q: %w", vt.Name(), err)
		}
	case vt.Implements(protoType):
		if err := proto.Unmarshal(data, val.Interface().(proto.Message)); err != nil {
			return val, fmt.Errorf("proto unmarshal failed on type %q: %w", vt.Name(), err)
		}
	default:
		dec := gob.NewDecoder(bytes.NewReader(data))
		if err := dec.DecodeValue(val); err != nil {
			return val, fmt.Errorf("gob decode failed on type %q: %w", vt.Name(), err)
		}
	}

	if vt.Kind() == reflect.Ptr {
		return val, nil
	}
	return val.Elem(), nil
}
