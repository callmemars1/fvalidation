package utils

type Optional[T any] struct {
	value T
	set   bool
}

func (o *Optional[T]) Value() T {
	return o.value
}

func (o *Optional[T]) HasValue() bool {
	return o.set
}

func Some[T any](value T) Optional[T] {
	return Optional[T]{
		value: value,
		set:   true,
	}
}

func None[T any]() Optional[T] {
	return Optional[T]{
		set: false,
	}
}
