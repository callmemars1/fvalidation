package testing

import (
	"fmt"
	fv "github.com/callmemars1/fvalidation"
	"github.com/callmemars1/fvalidation/rules"
	"regexp"
)

type Person struct {
	email    string
	name     string
	age      uint8
	passport string
	emails   []string
}

func HaveMaxLength(n int) fv.Rule[string] {
	return fv.NewRule[string](
		func(s string) (bool, error) {
			return len(s) <= n, nil
		},
		fv.FailurePayload{},
	)
}

var EmailRegex = regexp.MustCompile("")

func validate() {
	fv.BuildBaseValidator(func(p Person, vCtx *fv.ValidationContext, _ any) {
		fv.
			RuleFor(vCtx, "name", p.name).
			Must(rules.BeNotEmptyOrWhitespace()).
			WithMessage("Имя должно быть заполнено!").
			And().
			Must(HaveMaxLength(100)).
			WithMessage("Имя должно быть короче 100 символов!")

		fv.
			When(vCtx, func() (bool, error) {
				return p.age > 14, nil
			}, func() {
				fv.
					RuleFor(vCtx, "passport", p.passport).
					Must(HaveMaxLength(10)).
					WithMessage("Номер паспорта не более 10 символов!")
			})

		fv.
			RuleForEach(vCtx, p.emails, func(i int) string {
				return fmt.Sprintf("emails[%d]", i)
			}).
			Must(rules.BeNotEmptyOrWhitespace()).
			WithMessage("Email должен быть заполнен").
			And().
			Must(rules.MatchRegex(EmailRegex)).
			WithMessage("Некорректный формат электронной почты")
	}, nil)
}
