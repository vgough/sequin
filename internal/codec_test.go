package internal

import (
	"log/slog"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCodec(t *testing.T) {
	var c DefaultCodec
	arg1 := int(42)
	data, err := c.Encode(reflect.ValueOf(arg1))
	require.NoError(t, err)

	arg2, err := c.Decode(data, reflect.TypeOf(arg1))
	require.NoError(t, err)
	require.EqualValues(t, arg1, arg2.Int())
}

func TestCodec_Proto(t *testing.T) {
	var c DefaultCodec
	arg1 := &fakeProto{data: "hello"}
	data, err := c.Encode(reflect.ValueOf(arg1))
	require.NoError(t, err)
	require.Len(t, data, len(arg1.data))

	res, err := c.Decode(data, reflect.TypeOf(arg1))
	require.NoError(t, err)
	require.EqualValues(t, arg1, res.Interface())
}

type fakeProto struct {
	data string
}

func (p *fakeProto) MarshalVT() ([]byte, error) {
	slog.Info("marshaling", "data", p.data)
	return []byte(p.data), nil
}

func (p *fakeProto) UnmarshalVT(data []byte) error {
	slog.Info("unmarshaling", "data", string(data))
	p.data = string(data)
	return nil
}
