package rules

import (
	fv "github.com/callmemars1/fvalidation"
	"github.com/callmemars1/fvalidation/utils"
	"regexp"
	"strings"
)

func MatchRegex(r *regexp.Regexp) fv.Rule[string] {
	return fv.NewRule(
		func(s string) (bool, error) {
			return r.MatchString(s), nil
		},
		fv.FailurePayload{
			Code: utils.Some("REGEX_PATTERN_MISMATCH"),
		},
	)
}

func BeNotEmptyOrWhitespace() fv.Rule[string] {
	return fv.NewRule(
		func(s string) (bool, error) {
			trimmed := strings.TrimSpace(s)
			return len(trimmed) > 0, nil
		},
		fv.FailurePayload{
			Code: utils.Some("STRING_EMPTY_OR_WHITESPACE"),
		})
}
