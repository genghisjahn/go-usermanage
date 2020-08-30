package usermanage

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/genghisjahn/go-usermanage/v1/engine"
	"github.com/genghisjahn/go-usermanage/v1/ummock"
)

func main() {
	path := "config/config.json"
	defaultConfig := engine.Config{BcryptCost: 12, EmailRegEx: "[!-?a-~]+[@][!-?a-~]+[.][!-?a-~]+", PasswordRegEx: "[ -~]{8,}$"}
	dat, err := ioutil.ReadFile(path)
	c := engine.Config{}
	if err != nil {
		log.Printf("Config file not found %v\nDefaults Used: %#v", path, c)
	}
	err = json.Unmarshal(dat, &c)
	if err != nil {
		c = defaultConfig
	}

	um := ummock.NewUserManager()
	eng, engErr := engine.NewEngine(c, um)
	if engErr != nil {
		log.Fatal(engErr)
	}
	_ = eng
	eng.UserManager.CreateUser("hello", []byte("temp"))
}
