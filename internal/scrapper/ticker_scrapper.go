package ticker_scrapper

import (
	"github.com/gocolly/colly"
	"github.com/zeeshanahmad0201/ticker_trawler/pkg/validation"
)

func ScrapePrice(ticker string, sources []string) (float64, error) {
	price := 0.0

	var domains []string
	for _, source := range sources {
		// check if domains don't have source
		if !contains(domains, source) {
			domains = append(domains, validation.ParseDomain(source))
		}
	}

	if len(domains) == 0 {
		domains = append(domains, "finance.yahoo.com")
	}

	c := colly.NewCollector(colly.AllowedDomains(domains...))

	// Scrape the price from the specified source

	return price, nil
}
