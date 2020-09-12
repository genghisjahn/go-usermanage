package engine

import (
	"fmt"

	"github.com/genghisjahn/go-usermanage/v1/primitives"
)

//Engine is the struct that holds all of the parts of the functions, properties, configuration and interface implemntations that interact with the data store.
type Engine struct {
	primitives.UserManager
	epwconfig EmailPWConfig
}

//EmailPWConfig holds the regex for email, password and bcrypt cost.  It's expected to be loaded by a config file, sane defaults should be provided if the config file is missing.
type EmailPWConfig struct {
	EmailRegEx         string `json:"email_regex"`
	PasswordRegEx      string `json:"password_regex"`
	BcryptCost         int    `json:"bcrypt_cost"`
	InvalidEmailMsg    string `json:"invalid_email_message"`
	InvalidPasswordMsg string `json:"invalid_password_message"`
}

//NewEngine takes a variadic param to get the interfaces to build the usermanager engine
func NewEngine(parts ...interface{}) (Engine, error) {
	eng := Engine{}
	for _, v := range parts {
		var s = ""
		c, ok := v.(EmailPWConfig)
		if ok {
			s = "Config"
			eng.epwconfig = c
		}
		um, ok := v.(primitives.UserManager)
		if ok {
			s = "UserManager"
			eng.UserManager = um
		}
		if s == "" {
			return Engine{}, fmt.Errorf("Value does not fit a required interface %v %T", v, v)
		}
	}
	return eng, nil
}

//TODO

//I think we need the create,verify, login stuff here too
//And we do the verification here
//and then pass it on to the engine part that handles it
