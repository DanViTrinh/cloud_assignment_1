package utilities

//kind of errors like enum
type ErrorType int

const (
	ClientError ErrorType = iota + 1
	ServerError
	UnsensitiveServerError
)

type RestErrorWraper struct {
	OriginalError error
	StatusCode    int
	Message       string
	ErrorKind     ErrorType
}

func (e RestErrorWraper) Error() string {
	return e.OriginalError.Error()
}

func NewRestErrorWrapper(originalError error, statusCode int,
	message string, errorType ErrorType) error {
	return RestErrorWraper{
		OriginalError: originalError,
		StatusCode:    statusCode,
		Message:       message,
		ErrorKind:     errorType,
	}
}

func (e RestErrorWraper) Unwrap() error {
	return e.OriginalError
}
