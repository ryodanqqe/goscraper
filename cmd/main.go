package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type PageData struct {
	Rank       string
	Nickname   string
	Name       string
	Followers  string
	Category   string
	Country    string
	Authentic  string
	Engagement string
}

func main() {
	var pageData []PageData

	c := colly.NewCollector()

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36"

	c.OnHTML("div.row", func(e *colly.HTMLElement) {
		rank := e.ChildText(".row .row-cell.rank span[data-v-2e6a30b8]")
		nickname := e.ChildText("div.contributor__name-content")
		name := e.ChildText("div.contributor__title")
		followers := e.ChildText("div.row-cell.subscribers")
		category := e.ChildText("div.tag[data-v-595cc10b]")
		country := e.ChildText("div.row-cell.audience")
		authentic := e.ChildText("div.row-cell.authentic")
		engagement := e.ChildText("div.row-cell.engagement")

		data := PageData{
			Rank:       rank,
			Nickname:   nickname,
			Name:       name,
			Followers:  followers,
			Category:   category,
			Country:    country,
			Authentic:  authentic,
			Engagement: engagement,
		}
		pageData = append(pageData, data)

	})

	c.Visit("https://hypeauditor.com/top-instagram-all-russia/#")

	csvFile, csvErr := os.Create("out.csv")
	if csvErr != nil {
		log.Fatalf("%v", csvErr)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	headers := []string{
		"Rank",
		"Nickname",
		"Name",
		"Followers",
		"Category",
		"Country",
		"Authentic",
		"Engagement",
	}
	writer.Write(headers)

	for _, data := range pageData {
		record := []string{
			data.Rank,
			data.Nickname,
			data.Name,
			data.Followers,
			data.Category,
			data.Country,
			data.Authentic,
			data.Engagement,
		}

		writer.Write(record)
	}

}
