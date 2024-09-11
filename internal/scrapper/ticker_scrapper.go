package scrapper

import (
	"log"
	"sync"

	"github.com/gocolly/colly"
)

func ScrapePriceFromYahoo(ticker string) (float64, error) {
	log.Println("Starting to scrape for ticker:", ticker)
	scrapeURL := "https://finance.yahoo.com/quote/" + ticker

	c := colly.NewCollector()

	priceChan := make(chan float64)
	errChan := make(chan error)
	var wg sync.WaitGroup

	// c.OnError(func(r *colly.Response, err error) {
	// 	log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err.Error())
	// 	errChan <- err
	// })

	c.OnHTML(".company_snapshot_content .table-sm-responsive", func(h *colly.HTMLElement) {
		defer wg.Done()
		priceText := h.ChildText("th")
		log.Println("Scraped price text:", priceText)

		// priceText = strings.ReplaceAll(priceText, ",", "")
		// price, err := strconv.ParseFloat(priceText, 64)
		// if err != nil {
		// 	log.Println("Failed to parse price:", err)
		// 	errChan <- err
		// 	return
		// }

		// log.Println("Parsed price:", price)
		priceChan <- 10
	})

	log.Println("Visiting URL:", scrapeURL)

	wg.Add(1)
	go func() {
		c.Visit(scrapeURL)
	}()

	go func() {
		wg.Wait()
		close(priceChan)
		close(errChan)
	}()

	select {
	case price := <-priceChan:
		return price, nil
	case err := <-errChan:
		return -1, err
	}
}
