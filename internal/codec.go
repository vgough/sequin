package internal

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
)

func EncodeVarint(value int, buf [10]byte) []byte {
	n := 0
	for value >= 0x80 {
		buf[n] = byte(value) | 0x80
		value >>= 7
		n++
	}
	buf[n] = byte(value)
	return buf[:n+1]
}

func Encode(v reflect.Value) ([]byte, error) {
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

func Decode(data []byte, vt reflect.Type) (reflect.Value, error) {
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
