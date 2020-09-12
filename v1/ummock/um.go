package ummock

import (
	"log"
	"regexp"

	"github.com/genghisjahn/go-usermanage/v1/engine"
	"github.com/genghisjahn/go-usermanage/v1/primitives"
	"golang.org/x/crypto/bcrypt"
)

//UMMock struct that defines properties and functions for a mock/memory implemetation of the usermanager interface
type UMMock struct {
	Email  string `json:"email"`
	config engine.EmailPWConfig
}

//NewUserManager returns an instance of a struct that fufills the usermanager interface
func NewUserManager(c engine.EmailPWConfig) UMMock {
	um := UMMock{config: c}
	return um
}

//CreateUser accepts email and password arguments and attemps to create a user record.  ServiceError is the custom error type returned
func (u UMMock) CreateUser(email string, password []byte) (string, *primitives.ServiceError) {
	matched, err := regexp.Match(u.config.EmailRegEx, []byte(email))
	if err != nil {
		return "", primitives.NewServiceError(5, err.Error())
	}
	if !matched {
		return "", primitives.NewServiceError(4, "invalid email")
	}
	matched, err = regexp.Match(u.config.PasswordRegEx, password)
	if err != nil {
		return "", primitives.NewServiceError(5, err.Error())
	}
	if !matched {
		return "", primitives.NewServiceError(4, "invalid password")
	}
	pwhash, pwErr := hashPassword(password, u.config.BcryptCost)
	if pwErr != nil {
		return "", primitives.NewServiceError(5, pwErr.Error())
	}
	log.Println(pwhash)
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

func hashPassword(password []byte, c int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, c)
	return string(bytes), err
}
