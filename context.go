package fvalidation

type ValidationContext struct {
	err      error
	failures []Failure
}
