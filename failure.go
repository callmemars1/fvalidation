package fvalidation

import u "github.com/callmemars1/fvalidation/utils"

type FailurePayload struct {
	Code    u.Optional[string]
	Message u.Optional[string]
}

type Failure struct {
	field   string
	payload *FailurePayload
}
