package get_price

import (
	"log"

	"github.com/spf13/cobra"
)

// getPriceCmd represents the getPrice command
var GetPriceCmd = &cobra.Command{
	Use:   "get-price",
	Short: "Retrieves the current market price of a specified stock ticker",
	Long:  `This command allows users to quickly access real-time pricing information for a particular stock`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("flags: %v", cmd.Flag("ticker").Value)

	},
}

func init() {

	// Flags
	GetPriceCmd.PersistentFlags().String("ticker", "", "Ticker symbol of the stock")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getPriceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
