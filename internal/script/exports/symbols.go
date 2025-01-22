package exports

import "reflect"

var Symbols = map[string]map[string]reflect.Value{}

// Provide access to exported APIs.
//go:generate go run github.com/traefik/yaegi/cmd/yaegi extract github.com/vgough/sequin/internal/script/runtime
