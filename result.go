package fvalidation

type ValidationResult struct {
	failures []Failure
}

func (vr ValidationResult) IsValid() bool {
	return len(vr.failures) == 0
}

type ValidationResultMap map[string][]FailurePayloadMap

func (vr ValidationResult) ToMap() ValidationResultMap {
	result := make(ValidationResultMap)
	for _, failure := range vr.failures {
		var failures []FailurePayloadMap
		if result[failure.field] == nil {
			failures = make([]FailurePayloadMap, 0)
			result[failure.field] = failures
		}

		failures = result[failure.field]
		failures = append(failures, failure.payload.ToMap())
		result[failure.field] = failures
	}

	return result
}
