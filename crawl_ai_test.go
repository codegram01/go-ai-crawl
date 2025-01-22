package main

import "testing"

func TestCrawlWithAI(t *testing.T) {
	text, err := CrawlText("https://wingram.org/about/sites/", "body")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(text)

	res, err := AskAI(text, []byte(`
		{
			"type": "object",
			"properties": {
				"email": {
					"type": "string"
				}
			},
			"required": [
				"email"
			]
		}
	`))

	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}