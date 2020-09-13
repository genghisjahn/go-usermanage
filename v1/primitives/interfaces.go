package primitives

//UserManager is the interface that shows what functions are required to create, verify and login a user record
type UserManager interface {
	CreateUser(string, []byte) *ServiceError
	VerifyUser(string) *ServiceError
	LoginUser(string, []byte) *ServiceError
	//NewUserManager() UserManager
}
