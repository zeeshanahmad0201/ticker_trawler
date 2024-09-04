package get_price

import (
	"github.com/spf13/cobra"
	"github.com/zeeshanahmad0201/ticker_trawler/pkg/validation"
)

const (
	FlagTicker = "ticker"
	FlagSource = "source"
)

// getPriceCmd represents the getPrice command
var GetPriceCmd = &cobra.Command{
	Use:   "get-price",
	Short: "Retrieves the current market price of a specified stock ticker",
	Long:  `This command allows users to quickly access real-time pricing information for a particular stock`,
	Run: func(cmd *cobra.Command, args []string) {
		// validate ticker
		err := validation.ValidateTicker(cmd.Flag(FlagTicker).Value.String())
		if err != nil {
			cmd.PrintErr(err)
		}
		// fetch price

	},
}

func init() {

	// Flags
	GetPriceCmd.PersistentFlags().String(FlagTicker, "", "Ticker symbol of the stock (e.g., AAPL for Apple Inc.)")
	GetPriceCmd.PersistentFlags().String(FlagSource, "", "Source of the stock price data (e.g., finance.yahoo.com, investing.com)")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getPriceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
