package ummock

import (
	"regexp"

	"github.com/genghisjahn/go-usermanage/v1/engine"
	"github.com/genghisjahn/go-usermanage/v1/primitives"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var mockusers map[string]mockuser

func init() {
	mockusers = make(map[string]mockuser)
	mockusers["TEST_UUID"] = mockuser{"j@j.com", []byte("hashpassword")}
}

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
func (u UMMock) CreateUser(email string, password []byte) *primitives.ServiceError {
	matched, err := regexp.Match(u.config.EmailRegEx, []byte(email))
	if err != nil {
		return primitives.NewServiceError(5, err.Error())
	}
	if !matched {
		return primitives.NewServiceError(4, u.config.InvalidEmailMsg)
	}
	matched, err = regexp.Match(u.config.PasswordRegEx, password)
	if err != nil {
		return primitives.NewServiceError(5, err.Error())
	}
	if !matched {
		return primitives.NewServiceError(4, u.config.InvalidPasswordMsg)
	}
	pwhash, pwErr := hashPassword(password, u.config.BcryptCost)
	if pwErr != nil {
		return primitives.NewServiceError(5, pwErr.Error())
	}
	id := uuid.New().String()
	if mockusers == nil {
		mockusers = make(map[string]mockuser)
	}
	mockusers[id] = mockuser{Email: email, PWHash: pwhash}
	return nil
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

func hashPassword(password []byte, c int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, c)

}

type mockuser struct {
	Email  string
	PWHash []byte
}
