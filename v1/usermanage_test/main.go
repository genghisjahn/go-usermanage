package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/genghisjahn/go-usermanage/v1/engine"
	"github.com/genghisjahn/go-usermanage/v1/ummock"
)

func main() {
	path := "config/config.json"
	defaultConfig := engine.EmailPWConfig{BcryptCost: 12, EmailRegEx: "[!-?a-~]+[@][!-?a-~]+[.][!-?a-~]+", PasswordRegEx: "[ -~]{8,}$"}
	dat, err := ioutil.ReadFile(path)
	epwc := engine.EmailPWConfig{}
	if err != nil {
		log.Printf("Config file not found %v\nDefaults Used: %#v", path, epwc)
	}
	err = json.Unmarshal(dat, &epwc)
	if err != nil {
		epwc = defaultConfig
	}

	um := ummock.NewUserManager(epwc)
	eng, engErr := engine.NewEngine(um)
	if engErr != nil {
		log.Fatal(engErr)
	}

	guid, errU := eng.CreateUser("j@s.c", []byte("abcd1234"))
	if errU != nil {
		log.Fatal(errU)
	}
	fmt.Println(guid)
}
