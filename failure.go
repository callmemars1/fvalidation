package fvalidation

import u "github.com/callmemars1/fvalidation/utils"

type FailurePayload struct {
	Code    u.Optional[string]
	Message u.Optional[string]
}

type FailurePayloadMap map[string]string

func (p FailurePayload) ToMap() FailurePayloadMap {
	m := make(FailurePayloadMap)

	if p.Code.HasValue() {
		m["code"] = p.Code.Value()
	}

	if p.Message.HasValue() {
		m["message"] = p.Message.Value()
	}

	return m
}

type Failure struct {
	field   string
	payload *FailurePayload
}
