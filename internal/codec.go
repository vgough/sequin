package internal

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"

	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

var errorType = reflect.TypeFor[error]()

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
	case at == errorType:
		s := EncodeError(v.Interface().(error))
		return proto.Marshal(s.Proto())
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
	case vt == errorType:
		var s spb.Status
		if err := proto.Unmarshal(data, &s); err != nil {
			return val, fmt.Errorf("proto unmarshal failed: %w", err)
		}
		return reflect.ValueOf(DecodeError(&s)), nil
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

func EncodeError(err error) *status.Status {
	if err == nil {
		return status.New(codes.OK, "")
	}

	// If it's already a status, return it directly
	if s, ok := status.FromError(err); ok {
		return s
	}

	// Convert regular error to status with UNKNOWN code
	return status.New(codes.Unknown, err.Error())
}

func DecodeError(s *spb.Status) error {
	if s == nil || s.Code == int32(codes.OK) {
		return nil
	}
	return status.FromProto(s).Err()
}
