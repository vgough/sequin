package local

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
)

type DefaultCodec struct {
}

func (c *DefaultCodec) Encode(v reflect.Value) ([]byte, error) {
	at := v.Type()
	switch {
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
