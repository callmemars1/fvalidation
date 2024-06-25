package fvalidation

type ConditionalResult struct {
	ctx              *ValidationContext
	conditionSucceed bool
}

func When(
	ctx *ValidationContext,
	condition func() (bool, error),
	flow func()) ConditionalResult {
	if ctx.err != nil {
		return ConditionalResult{
			ctx: ctx,
		}
	}

	ok, err := condition()
	if err != nil {
		ctx.err = err
		return ConditionalResult{
			ctx: ctx,
		}
	}

	if ok {
		flow()
	}

	return ConditionalResult{
		conditionSucceed: ok,
		ctx:              ctx,
	}
}

func (cr *ConditionalResult) OtherwiseWhen(
	condition func() (bool, error),
	flow func()) ConditionalResult {

	if cr.ctx.err != nil {
		return ConditionalResult{
			ctx: cr.ctx,
		}
	}

	if cr.conditionSucceed {
		return *cr
	}

	ok, err := condition()
	if err != nil {
		cr.ctx.err = err
		return ConditionalResult{
			ctx: cr.ctx,
		}
	}

	if ok {
		flow()
	}

	return ConditionalResult{
		conditionSucceed: ok,
		ctx:              cr.ctx,
	}
}

func (cr *ConditionalResult) Otherwise(flow func()) {
	if !cr.conditionSucceed {
		flow()
	}
}
