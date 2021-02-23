package main

import (
	"io/ioutil"
)

import "github.com/pelletier/go-toml"

type Docusign struct {
	IntegrationKey string
	ApiUsername    string
	PrivateKey     string
}
type Secrets struct {
	Docusign Docusign
}

func loadSecrets() Secrets {
	data, err := ioutil.ReadFile("secrets.toml")
	check(err)
	secrets := Secrets{}
	err = toml.Unmarshal(data, &secrets)
	check(err)
	return secrets
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
