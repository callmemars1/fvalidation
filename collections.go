package fvalidation

type NameProvider func(int) string

type Collection[T any] struct {
	nameProvider NameProvider
	values       []T
	vCtx         *ValidationContext
}

func RuleForEach[T any](vCtx *ValidationContext, values []T, nameProvider NameProvider) *Collection[T] {
	return &Collection[T]{
		nameProvider: nameProvider,
		values:       values,
		vCtx:         vCtx,
	}
}

func (col *Collection[T]) Must(rule Rule[T]) *CollectionRuleValidationFailure[T] {
	if col.vCtx.err != nil {
		return &CollectionRuleValidationFailure[T]{
			collection:      col,
			failurePayloads: nil,
		}
	}

	payloads := make([]*FailurePayload, 0)
	for i, value := range col.values {
		valid, err := rule.validate(value)
		if err != nil {
			col.vCtx.err = err
			return &CollectionRuleValidationFailure[T]{
				collection:      col,
				failurePayloads: nil,
			}
		}

		if !valid {
			payload := rule.DefaultPayload()
			payloads = append(payloads, &payload)

			failure := Failure{
				field:   col.nameProvider(i),
				payload: &payload,
			}
			col.vCtx.failures = append(col.vCtx.failures, failure)
		}
	}

	return &CollectionRuleValidationFailure[T]{
		collection:      col,
		failurePayloads: payloads,
	}
}
