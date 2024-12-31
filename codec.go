package sequin

import "reflect"

type Codec interface {
	Encode(reflect.Value) ([]byte, error)
	Decode([]byte, reflect.Type) (reflect.Value, error)
}
