package main

import (
	"testing"
)

func TestAskAI(t *testing.T) {
	res, err := AskAI("Hello they", []byte{})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}

func TestAskAIWithFormat(t *testing.T) {
	var mess = `
		About Wingram
		my email is wingram@gmail.com
		my phone is 123123123
		my address is VietNam
	`

	res, err := AskAI(mess, []byte(`
		{
			"type": "object",
			"properties": {
				"email": {
					"type": "string"
				},
				"phone": {
					"type": "string"
				}
			},
			"required": [
				"email",
				"phone"
			]
		}
	`))

	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}