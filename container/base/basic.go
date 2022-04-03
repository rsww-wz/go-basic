package base

type UnSign interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Sign interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Float interface {
	~float32 | ~float64
}

type Integer interface {
	UnSign | Sign
}

type Number interface {
	Integer | Float
}

type Complex interface {
	~complex64 | ~complex128
}

type Order interface {
	Number | Complex | string
}

type Map[T Order] map[T]any

type Chan[T any] chan T
