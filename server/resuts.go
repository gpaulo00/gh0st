package server

// DoneResult is a basic response of the server
type DoneResult struct {
	Result bool `json:"result"`
}

// ErrorResult is an error response of the server
type ErrorResult struct {
	Error string `json:"error"`
}

// Error creates a new ErrorResult from error
func Error(e error) ErrorResult {
	return ErrorResult{Error: e.Error()}
}

// Done is the default DoneResult instance
var Done = DoneResult{Result: true}
