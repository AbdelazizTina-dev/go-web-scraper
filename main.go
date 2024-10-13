package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("To use the command, please provide an URL, e.g: scrape <url>")
		return
	}

	base_url := os.Args[1]

	c := colly.NewCollector()

	file, err := os.Create("data.txt") // This will create or overwrite the file

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		parsedURL, err := url.Parse(link)
		if err != nil {
			fmt.Println("Error parsing URL:", err)
			return
		}

		fullUrl := link

		if !parsedURL.IsAbs() {
			fullUrl = base_url + link
		}

		line := fmt.Sprintf("Link found: %s\n", fullUrl)
		file.WriteString(line)

		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(base_url)
}
