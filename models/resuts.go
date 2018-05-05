package models

import "errors"

// DoneResult is a basic response of the server
type DoneResult struct {
	Result bool `json:"result"`
}

// ErrorResult is an error response of the server
type ErrorResult struct {
	Error string `json:"error"`
}

// Err converts the ErrorResult to an Error
func (e ErrorResult) Err() error {
	if e.Error == "" {
		return nil
	}
	return errors.New(e.Error)
}

// Error creates a new ErrorResult from error
func Error(e error) ErrorResult {
	return ErrorResult{Error: e.Error()}
}

// Done is the default DoneResult instance
var Done = DoneResult{Result: true}
