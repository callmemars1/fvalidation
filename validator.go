package fvalidation

type Validator[T any] interface {
	Validate(T) (ValidationResult, error)
}

type BaseValidationPipeline[T any, TDep any] func(T, *ValidationContext, TDep)

type BaseValidator[T any, TDep any] struct {
	pipeline     BaseValidationPipeline[T, TDep]
	dependencies TDep
}

func BuildBaseValidator[T any, TDep any](pipeline BaseValidationPipeline[T, TDep], dep TDep) BaseValidator[T, TDep] {
	return BaseValidator[T, TDep]{pipeline: pipeline, dependencies: dep}
}

func (v *BaseValidator[T, TDep]) Validate(value T) (*ValidationResult, error) {
	vCtx := &ValidationContext{
		failures: make([]Failure, 0),
	}

	v.pipeline(value, vCtx, v.dependencies)

	if vCtx.err != nil {
		return nil, vCtx.err
	}

	return &ValidationResult{
		failures: vCtx.failures,
	}, nil
}
