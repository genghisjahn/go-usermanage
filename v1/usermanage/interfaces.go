package usermanage

//UserManager is the interface that shows what functions are required to create, verify and login a user record
type UserManager interface {
	CreateUser(string, []byte) (string, error)
	VerifyUser(string) error
	LoginUser(string, []byte) error
	NewUserManager() UserManager
}
