package scrapper

import (
	"log"
	"sync"

	"github.com/gocolly/colly"
)

func ScrapePriceFromYahoo(ticker string) (float64, error) {

	// Log the start of the function
	log.Println("Starting to scrape for ticker:", ticker)

	// The specific Yahoo Finance page for the stock ticker
	scrapeURL := "https://sarmaaya.pk/psx/company/ENGRO"

	// Create a new collector
	c := colly.NewCollector()
	priceChan := make(chan float64)
	errChan := make(chan error)
	var wg sync.WaitGroup

	// Set headers to mimic a browser request
	c.OnRequest(func(r *colly.Request) {
		log.Println("Sending request to:", r.URL.String())
	})

	// Error handling
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err.Error())
		errChan <- err
	})

	// Find the price inside the #quote-header-info div
	c.OnHTML("div.row", func(h *colly.HTMLElement) {

		log.Printf("Scraped data: %s", h.ChildText("div.col-lg-2"))

		// // Find the span containing the current stock price
		// priceText := h.ChildText("fin-streamer[data-field='regularMarketPrice']")
		// log.Println("Scraped price text:", priceText)

		// // Clean up the price and convert it to float64
		// priceText = strings.ReplaceAll(priceText, ",", "") // Remove commas from the price
		// price, err := strconv.ParseFloat(priceText, 64)
		// if err != nil {
		// 	log.Println("Failed to parse price:", err)
		// 	errChan <- fmt.Errorf("failed to parse price: %v", err)
		// 	return
		// }

		// log.Println("Parsed price:", price)

		// Send the price to the priceChan
		priceChan <- 64
	})

	// Log before visiting the URL
	log.Println("Visiting URL:", scrapeURL)

	// Visit the URL
	wg.Add(1)
	go func() {
		defer wg.Done()
		c.Visit(scrapeURL)
	}()

	go func() {
		wg.Wait()
		close(priceChan)
		close(errChan)
	}()

	// Log before the select block
	log.Println("Waiting for result...")

	// Wait for either the price or an error to be returned
	select {
	case price := <-priceChan:
		log.Println("Received price:", price)
		return price, nil
	case err := <-errChan:
		log.Println("Error received:", err)
		return 0.0, err
	}
}
