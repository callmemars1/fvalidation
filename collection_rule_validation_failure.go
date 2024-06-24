package fvalidation

import u "github.com/callmemars1/fvalidation/utils"

type CollectionRuleValidationFailure[T any] struct {
	collection      *Collection[T]
	failurePayloads []*FailurePayload
}

func (r *CollectionRuleValidationFailure[T]) And() *Collection[T] {
	return r.collection
}

func (r *CollectionRuleValidationFailure[T]) WithCode(code string) *CollectionRuleValidationFailure[T] {
	for _, payload := range r.failurePayloads {
		payload.Code = u.Some(code)
	}
	return r
}

func (r *CollectionRuleValidationFailure[T]) WithMessage(message string) *CollectionRuleValidationFailure[T] {
	for _, payload := range r.failurePayloads {
		payload.Message = u.Some(message)
	}
	return r
}
