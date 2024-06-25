package fvalidation

import u "github.com/callmemars1/fvalidation/utils"

type RuleValidationFailure[T any] struct {
	field   *Field[T]
	payload *FailurePayload
}

func (r *RuleValidationFailure[T]) And() *Field[T] {
	return r.field
}

func (r *RuleValidationFailure[T]) WithCode(code string) *RuleValidationFailure[T] {
	r.payload.Code = u.Some(code)
	return r
}

func (r *RuleValidationFailure[T]) WithMessage(message string) *RuleValidationFailure[T] {
	r.payload.Message = u.Some(message)
	return r
}
