package fvalidation

type Field[T any] struct {
	title string
	value T
	vCtx  *ValidationContext
}

func RuleFor[T any](vCtx *ValidationContext, title string, value T) *Field[T] {
	return &Field[T]{
		vCtx:  vCtx,
		title: title,
		value: value,
	}
}

func (f *Field[T]) Must(rule Rule[T]) *RuleValidationFailure[T] {
	if f.vCtx.err != nil {
		return &RuleValidationFailure[T]{
			field:   f,
			payload: nil,
		}
	}

	valid, err := rule.validate(f.value)
	if err != nil {
		f.vCtx.err = err
		return &RuleValidationFailure[T]{
			field:   f,
			payload: nil,
		}
	}

	payload := rule.DefaultPayload()

	if valid {
		return &RuleValidationFailure[T]{
			field:   f,
			payload: &payload,
		}
	}

	failure := Failure{
		field:   f.title,
		payload: &payload,
	}
	f.vCtx.failures = append(f.vCtx.failures, failure)

	return &RuleValidationFailure[T]{
		field:   f,
		payload: &payload,
	}
}
