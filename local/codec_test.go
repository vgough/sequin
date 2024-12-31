package local

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCodec(t *testing.T) {
	arg1 := int(42)
	data, err := Encode(reflect.ValueOf(arg1))
	require.NoError(t, err)

	arg2, err := Decode(data, reflect.TypeOf(arg1))
	require.NoError(t, err)
	require.EqualValues(t, arg1, arg2.Int())
}

func TestCodec_Proto(t *testing.T) {
	arg1 := &fakeProto{Data: "hello"}
	data, err := Encode(reflect.ValueOf(arg1))
	require.NoError(t, err)

	res, err := Decode(data, reflect.TypeOf(arg1))
	require.NoError(t, err)
	require.EqualValues(t, arg1, res.Interface())
}

type fakeProto struct {
	Data string
}
