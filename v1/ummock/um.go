package ummock

import (
	"github.com/genghisjahn/go-usermanage/v1/engine"
	"github.com/genghisjahn/go-usermanage/v1/primitives"
)

//UMMock struct that defines properties and functions for a mock/memory implemetation of the usermanager interface
type UMMock struct {
	Email string `json:"email"`
}

//NewUserManager returns an instance of a struct that fufills the usermanager interface
func NewUserManager(c engine.EmailPWConfig) UMMock {
	return UMMock{}
}

//CreateUser accepts email and password arguments and attemps to create a user record.  ServiceError is the custom error type returned
func (u UMMock) CreateUser(email, string, password []byte) (string, *primitives.ServiceError) {
	//Where do we get the bcrypt cost from?
	//How do we know where to store things?
	//How do we know what the password requirements are?
	//How do we know what the email validate is?
	return "", nil
}

//VerifyUser accepts a GUID and attempts to update the user record so that it is marked as verified.  ServiceError is the customer error type returned
func (u UMMock) VerifyUser(confirmationGUID string) *primitives.ServiceError {
	//How do we connect to the data store?
	return nil
}

//LoginUser accepts email and password, if the email address has been validated and the email & password are correct, then return nil
func (u UMMock) LoginUser(email string, password []byte) *primitives.ServiceError {
	//Validate the email first, make sure it's not gargage
	return nil
}
