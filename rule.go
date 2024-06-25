package fvalidation

type Predicate[T any] func(T) (bool, error)

type Rule[T any] struct {
	validate       Predicate[T]
	defaultPayload FailurePayload
}

func (r *Rule[T]) Validate(value T) (bool, error) {
	return r.validate(value)
}

func (r *Rule[T]) DefaultPayload() FailurePayload {
	return r.defaultPayload
}

func NewRule[T any](predicate Predicate[T], payload FailurePayload) Rule[T] {
	return Rule[T]{
		validate:       predicate,
		defaultPayload: payload,
	}
}
