## Technical specification for version 1

This document is subordinate to the prose_spec document in this directory.

Fucntions Implemented:

1. createUser(email, string, password []byte)(string,error)
    1. Parameters
        1. `email` is the email address supplised by the user
        1. `password` is the password to verify the user, value is also supplied by the user
    1. Return values
        1. `string` - the confirmation GUID that will be used in a later to call to verify the user.  It is inteneded to be embedded as a querstry argument in a REST call later.
        1. `error` - An error value for the function.  If everything goes write, it'll be `nil`.  If something goes wrong it will return the error _and_ an error type.  The error type will be a string value of `client` or `server`.  At this level we aren't talking about error codes in a REST API response, but we have an eye towards that.  
            1. An error will have an errortype of `client` if the email address is invalid or already used.  Also if the password fails password requirements.  For version 1 the requiements will be that it must be 8 bytes long.
            2. An error type of `server` will be returned if there is an error when writing to the datastore