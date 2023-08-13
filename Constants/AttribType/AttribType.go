package AttribType

import "reflect"

type AttribType int

const (
	Int     AttribType = 0
	Uint    AttribType = 1
	Float   AttribType = 2
	String  AttribType = 3
	Bool    AttribType = 4
	Invalid AttribType = -1
)

func AttribTypeFromKind(kind reflect.Kind) AttribType {
	switch kind {
	case reflect.Int:
		return Int
	case reflect.Uint:
		return Uint
	case reflect.Float64 | reflect.Float32:
		return Float
	case reflect.String:
		return String
	case reflect.Bool:
		return Bool
	default:
		return Invalid
	}
}
