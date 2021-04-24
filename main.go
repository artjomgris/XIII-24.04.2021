package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type Token struct {
	Header struct {
		Alg string `json:"alg"`
		Typ string	`json:"typ"`
	}
	Payload struct {
		Sub string `json:"sub"`
		Name string `json:"name"`
		Iat int `json:"iat"`
	}
	Signature struct{
		Key string `json:"key"`
		Secret string `json:"secret"`
	}

}

func main() {

	var m Token
	m.Header.Alg = "HS256"
	m.Header.Typ = "JWT"
	
	m.Payload.Sub = "1234567890"
	m.Payload.Name = "John Doe"
	m.Payload.Iat = 1516239022

	m.Signature.Secret = "KUKUSHKA"

	data, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}


	
	str := base64.StdEncoding.EncodeToString(m.Header."."."")
	fmt.Println(str)
}