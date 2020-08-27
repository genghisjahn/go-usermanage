## Technical specification for version 1

This document is subordinate to the prose_spec document in this directory.

Fucntions Implemented:

1. `func createUser(email, string, password []byte)(string,error)`
    1. Parameters
        1. `email` is the email address supplised by the user
        1. `password` is the password to verify the user, value is also supplied by the user
    1. Return values
        1. `string` - the confirmation GUID that will be used in a later to call to verify the user.  It is inteneded to be embedded as a querstry argument in a REST call later.
        1. `error` - An error value for the function.  If everything goes write, it'll be `nil`.  If something goes wrong it will return the error _and_ an error type [ServiceError](#ServiceError).  The error type will be a string value of `client` or `server`.  At this level we aren't talking about error codes in a REST API response, but we have an eye towards that.  
            1. An error will have an errortype of `4` (client) if the email address is invalid or already used.  Also if the password fails password requirements.  For version 1 the requiements will be that it must be 8 bytes long.
            2. An error type of `5`(server) will be returned if there is an error when writing to the datastore
1. `func verifyUser(confirmationGUID string)error`
    1. Parameters
        1. `confirmationGUID` string is the GUID that was returned by the `createUser` function above
    1. Return values
        1. A `ServiceError`. `4`(client) errors will be returned if the GUID does not exist or has already been verified.  `5`(server) errors will be returned if something went wrong with communicating with the data layers. `nil` means that the account was verified successfully.
1. `func loginUser(email, string, password []byte) error`
    1. Parameters
        1. `email` is the email address supplised by the user
        1. `password` is the password to verify the user, value is also supplied by the user
    2. Return Values
        1. `error` is a ServiceError. `4`(client) for various validation errors (around email address or a password that doesn't meet requirements) and `5`(server) for problems talking to the database.


### ServiceError
```go

const ErrorTypeClient = 4
const ErrorTypeServer = 5

type ServiceError struct{
    Message string
    ErrorType int
}

func NewServiceError(errortype int, message string) ServiceError{
    if errortype != 4 || errortype!=5 {
        panic("Invalid Error Type " & m.ErrorType)
    }
    return ServiceError{Message: message, ErrorType: errortype}
}

func (m ServiceError) Error() string {
    if m.ErrorType == 4{
        return "Client Error: " & m.Message
    }
    if m.ErrorType == 5{
        return "Server Error: " & m.Message
    }
    panic("Invalid Error Type " & m.ErrorType)
}

```