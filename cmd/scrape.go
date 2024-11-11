package cmd

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
	"github.com/spf13/cobra"
)

func scrape(url string) {
	c := colly.NewCollector(
		colly.Async(true),
	)

	// Set custom User-Agent
	c.UserAgent = "GoBagWebScraper/1.0"

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Found link: %s\n", link)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Error:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Scraping complete.")
	})

	// Start scraping
	if err := c.Visit(url); err != nil {
		log.Fatalf("Failed to visit URL: %v", err)
	}

	c.Wait()
}

var scrapeCmd = &cobra.Command{
	Use:   "scraper",
	Short: "A simple web scraper",
	Long:  "This tool scrapes a given URL and prints all links found on the page.",
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		fmt.Printf("Scraping URL: %s\n", url)
		scrape(url)
	},
}

func init() {
	rootCmd.AddCommand(scrapeCmd)
}
