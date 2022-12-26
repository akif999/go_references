package references

var builtins map[string]struct{} = map[string]struct{}{

	"bool":        struct{}{},
	"true":        struct{}{},
	"false":       struct{}{},
	"uint8":       struct{}{},
	"uint16":      struct{}{},
	"uint32":      struct{}{},
	"uint64":      struct{}{},
	"int8":        struct{}{},
	"int16":       struct{}{},
	"int32":       struct{}{},
	"int64":       struct{}{},
	"float32":     struct{}{},
	"float64":     struct{}{},
	"complex64":   struct{}{},
	"complex128":  struct{}{},
	"string":      struct{}{},
	"int":         struct{}{},
	"uint":        struct{}{},
	"uintptr":     struct{}{},
	"byte":        struct{}{},
	"rune":        struct{}{},
	"Type":        struct{}{},
	"Type1":       struct{}{},
	"IntegerType": struct{}{},
	"FloatType":   struct{}{},
	"ComplexType": struct{}{},
	"append":      struct{}{},
	"copy":        struct{}{},
	"delete":      struct{}{},
	"len":         struct{}{},
	"cap":         struct{}{},
	"make":        struct{}{},
	"new":         struct{}{},
	"complex":     struct{}{},
	"real":        struct{}{},
	"imag":        struct{}{},
	"close":       struct{}{},
	"panic":       struct{}{},
	"recover":     struct{}{},
	"print":       struct{}{},
	"println":     struct{}{},
	"error":       struct{}{},
}

func isBuiltin(ident string) bool {
	_, ok := builtins[ident]
	return ok
}
