package main

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func Crawl(url string, actions ...chromedp.Action) error {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.106 Safari/537.36"),
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(context.Background())
	defer cancel()

	actions = append([]chromedp.Action{
		chromedp.Navigate(url)},
		actions...,
	)

	err := chromedp.Run(ctx, actions...)

	return err
}

func CrawlAttributes(link string, selector string, attr string) ([]string, error) {
	var nodes []*cdp.Node
	var contents []string

	err := Crawl(
		link,
		chromedp.Nodes(selector, &nodes, chromedp.ByQueryAll),
	)

	if err != nil {
		return contents, err
	}

	for _, node := range nodes {
		content := node.AttributeValue(attr)
		contents = append(contents, content)
	}

	return contents, nil
}

func CrawlText(link string, selector string) (string, error) {
	var text string
	err := Crawl(link,
		chromedp.Text(selector, &text),
	)

	return text, err
}

func CrawlHtml(link string, selector string) (string, error) {
	var text string
	err := Crawl(link,
		chromedp.InnerHTML(selector, &text),
	)

	return text, err
}
