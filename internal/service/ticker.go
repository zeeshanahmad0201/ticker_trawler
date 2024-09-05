package ticker

import (
	"strings"

	"github.com/zeeshanahmad0201/ticker_trawler/internal/scrapper"
	"github.com/zeeshanahmad0201/ticker_trawler/pkg/helpers"
)

func FetchPrice(ticker string, source string) (float64, error) {
	price := 0.0

	// validate the ticker symbol
	if err := helpers.ValidateTicker(ticker); err != nil {
		return price, err
	}

	if source == "" {
		source = "finance.yahoo.com"
	}

	source = strings.ToLower(source)

	switch source {
	case "finance.yahoo.com":
		return scrapper.ScrapePriceFromYahoo(ticker)
	}

	return price, nil
}
