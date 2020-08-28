package um

//CreateUser accepts email and password arguments and attemps to create a user record.  ServiceError is the custom error type returned
func CreateUser(email, string, password []byte) (string, error) {
	//Where do we get the bcrypt cost from?
	//How do we know where to store things?
	//How do we know what the password requirements are?
	//How do we know what the email validate is?
	return "", nil
}

//VerifyUser accepts a GUID and attempts to update the user record so that it is marked as verified.  ServiceError is the customer error type returned
func VerifyUser(confirmationGUID string) error {
	//How do we connect to the data store?
	return nil
}

//LoginUser accepts email and password, if the email address has been validated and the email & password are correct, then return nil
func LoginUser(email, string, password []byte) error {
	//Validate the email first, make sure it's not gargage
	return nil
}
