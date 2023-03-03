package utilities

// Client error
// OrigErr is the original error
// StatusCode is the status code to be sent to user
// Message is message for client
type ClientError struct {
	OrigErr    error
	StatusCode int
	Message    string
}

// returns original error
func (e ClientError) Error() string {

	return e.OrigErr.Error()
}

// Creates a new client error rapper for original error
// origErr is the original error
// statusCode is the status code to be sent to client
// msg is message for client
func NewClientError(origErr error, statusCode int, msg string) error {
	return ClientError{
		OrigErr:    origErr,
		StatusCode: statusCode,
		Message:    msg,
	}
}

// Unwraps to give the original error
func (e ClientError) Unwrap() error {
	return e.OrigErr
}

// Server error
// OrigErr is the original error
// StatusCode is the status code to be sent to user
// UsrMessage is message for client
// DevMessage is extra message for dev
type ServerError struct {
	OrigErr    error
	StatusCode int
	UsrMessage string
	DevMessage string
}

// returns original error
func (e ServerError) Error() string {
	return e.OrigErr.Error()
}

// Creates a new client error rapper for original error
// origErr is the original error
// statusCode is the status code to be sent to client
// usrMsg is message for user
// devMsg is extra message for dev
func NewServerError(origErr error, statusCode int, usrMsg, devMsg string) error {
	return ServerError{
		OrigErr:    origErr,
		StatusCode: statusCode,
		UsrMessage: usrMsg,
		DevMessage: devMsg,
	}
}

// Unwraps to give the original error
func (e ServerError) Unwrap() error {
	return e.OrigErr
}
