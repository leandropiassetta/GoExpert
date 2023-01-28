package main

import (
	"encoding/json"
	"os"
)

type bankAccount struct {
	Number  int `json:"-"`                          // When i dont want show this field, i put this dash
	Balance int `json:"otherThing" validate:"gt=0"` // this are tags of anotation
}

func main() {
	account := bankAccount{Number: 1, Balance: 100}

	// serialization json
	// MARSHAL -> transform in a json and save the resul in a variable
	res, err := json.Marshal(account)
	if err != nil {
		println(err)
	}

	println(string(res))
	// ENCODER -> Sometimes i don't save the result inside a variable but already return the result of json. I want receive a value and make it a record somewhere else, this place can be a file or a screen of terminal.

	// os.Stdout -> will exit to the default screen, i.e in my output terminal
	err = json.NewEncoder(os.Stdout).Encode(account)

	if err != nil {
		println(err)
	}

	// deserialization json
	// UNMARSHAL -> transform a json in a other struct
	pureJson := []byte(`{"Number":2, "otherThing":200}`)
	var accountAA bankAccount

	err = json.Unmarshal(pureJson, &accountAA)
	if err != nil {
		println(err)
	}
	println(accountAA.Balance)
}
