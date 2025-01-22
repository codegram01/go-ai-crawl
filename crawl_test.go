package main

import "testing"

func TestCrawlHtml(t *testing.T) {
	html, err := CrawlHtml("https://wingram.org", "footer")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(html)
}


func TestCrawlText(t *testing.T) {
	text, err := CrawlText("https://wingram.org", "footer")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(text)
}

func TestCrawlAttributes(t *testing.T) {
	attrs, err := CrawlAttributes("https://wingram.org/about/sites/", "a", "href")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(attrs)
}