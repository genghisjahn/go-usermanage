package um

const ErrorTypeClient = 4
const ErrorTypeServer = 5

type ServiceError struct {
	Message   string
	ErrorType int
}

func NewServiceError(errortype int, message string) ServiceError {
	if errortype != 4 && errortype != 5 {
		panic("Invalid Error Type " + string(errortype))
	}
	return ServiceError{Message: message, ErrorType: errortype}
}

func (m ServiceError) Error() string {
	if m.ErrorType == 4 {
		return "Client Error: " + m.Message
	}
	if m.ErrorType == 5 {
		return "Server Error: " + m.Message
	}
	panic("Invalid Error Type " + string(m.ErrorType))
}
