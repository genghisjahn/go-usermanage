package usermanage

const errorTypeClient = 4
const errorTypeServer = 5

//ServiceError is a customer error type to help differentiate between client and server side errors
type ServiceError struct {
	Message   string
	ErrorType int
}

//NewServiceError accepts a 4 or a 5 and a message to create a ServiceError
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

//Engine is the struct that holds all of the parts of the functions, properties, configuration and interface implemntations that interact with the data store.
type Engine struct {
}

type Config struct {
	EmailRegEx    string
	PasswordRegEx string
	BycryptCost   int
}
