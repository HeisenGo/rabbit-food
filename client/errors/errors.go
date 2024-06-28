package errors

type Error struct {
	Field   string
	Message string
}

func newError(field string, msg string) *Error {
	return &Error{
		Field:   field,
		Message: msg,
	}
}

func (e *Error) Error() string {
	return e.Message
}

var (
	ErrConnectionFailed            = newError("connection", "connection failed to be established")
	ErrEncodingRequest             = newError("request", "encoding request failed")
	ErrWritingToServer             = newError("connection", "writing request to server failed")
	ErrReadingResponse             = newError("connection", "failed to read response from server")
	ErrDecodingResponse            = newError("response", "decoding response failed")
	ErrDecodingServerErrorResponse = newError("response", "error in decoding error response")
	ErrDecodingSuccessfulResponse  = newError("response", "error in decoding a successful response")
	ErrDecodingToken               = newError("token", "error decoding token")
	ErrDataType                    = newError("data", "data type error")
)
