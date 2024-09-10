package stocks

import (
	"encoding/json"
	"net/http"

	"github.com/StreamFinance-Protocol/stream-chain/protocol/daemons/pricefeed/client/price_function"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/daemons/pricefeed/types"
)

// StocksTicker is our representation of ticker information returned in Stocks api response.
// It implements interface `Ticker` in util.go.
type StocksTicker struct {
	Ticker    string `json:"symbol" validate:"required"`
	AskPrice  string `json:"askPrice" validate:"required,positive-float-string"`
	BidPrice  string `json:"bidPrice" validate:"required,positive-float-string"`
	LastPrice string `json:"lastPrice" validate:"required,positive-float-string"`
}

// Ensure that StocksTicker implements the Ticker interface at compile time.
var _ price_function.Ticker = (*StocksTicker)(nil)

func (t StocksTicker) GetPair() string {
	// needs to be wrapped in quotes to be consistent with the API request format.
	return t.Ticker
}

func (t StocksTicker) GetAskPrice() string {
	return t.AskPrice
}

func (t StocksTicker) GetBidPrice() string {
	return t.BidPrice
}

func (t StocksTicker) GetLastPrice() string {
	return t.LastPrice
}

// StockPriceFunction transforms an API response from stock api into a map of tickers to prices that have been
// shifted by a market specific exponent.
func StocksPriceFunction(
	response *http.Response,
	tickerToExponent map[string]int32,
	resolver types.Resolver,
) (tickerToPrice map[string]uint64, unavailableTickers map[string]error, err error) {
	// Unmarshal response body into a list of tickers.
	var stockTickers []StocksTicker
	err = json.NewDecoder(response.Body).Decode(&stockTickers)
	if err != nil {
		return nil, nil, err
	}

	return price_function.GetMedianPricesFromTickers(
		stockTickers,
		tickerToExponent,
		resolver,
	)
}
