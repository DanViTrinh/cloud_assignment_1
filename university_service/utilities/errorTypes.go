package utilities

// Client error
//
// 	OrigErr - the original error
// 	StatusCode - status code to be sent to user
// 	Message - message for client
type ClientError struct {
	OrigErr    error
	StatusCode int
	Message    string
}

// Server error
//
// 	OrigErr - original error
// 	StatusCode - status code to be sent to user
// 	UsrMessage - message for client
// 	DevMessage - extra message for dev
type ServerError struct {
	OrigErr    error
	StatusCode int
	UsrMessage string
	DevMessage string
}

// Not found error
// return this error if not found is received from foreign api

// returns original error.Error() for ClientError
func (e ClientError) Error() string {

	return e.OrigErr.Error()
}

// returns original error.Error() for ServerError
func (e ServerError) Error() string {
	return e.OrigErr.Error()
}

// Creates a new client error wrapper for original error
//
// Parameters:
//
// 	origErr - original error
// 	statusCode - status code to be sent to client
// 	msg - message for client
//
// Returns:
//
// 	ClientError - new ClientError with params
func NewClientError(origErr error, statusCode int, msg string) error {
	return ClientError{
		OrigErr:    origErr,
		StatusCode: statusCode,
		Message:    msg,
	}
}

// Creates a new client error wrapper for original error
//
// Parameters:
//
// 	origErr - original error
// 	statusCode - status code to be sent to client
// 	usrMsg - message for client
// 	devMsg - extra message for dev
//
// Returns:
//
// 	ServerError - new ServerError with params
func NewServerError(origErr error, statusCode int, usrMsg, devMsg string) error {
	return ServerError{
		OrigErr:    origErr,
		StatusCode: statusCode,
		UsrMessage: usrMsg,
		DevMessage: devMsg,
	}
}

// Unwraps to give the original error for ClientError
func (e ClientError) Unwrap() error {
	return e.OrigErr
}

// Unwraps to give the original error for ServerError
func (e ServerError) Unwrap() error {
	return e.OrigErr
}
