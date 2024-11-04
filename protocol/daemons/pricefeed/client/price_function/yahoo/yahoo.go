package yahoo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/StreamFinance-Protocol/stream-chain/protocol/daemons/pricefeed/client/price_function"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/daemons/pricefeed/types"
)

// YahooTicker is our representation of ticker information returned in Yahoo response.
// It implements interface `Ticker` in util.go.
type YahooTicker struct {
	Pair      string `json:"symbol" validate:"required"`
	AskPrice  string `json:"ask" validate:"required"`
	BidPrice  string `json:"bid" validate:"required"`
	LastPrice string `json:"regularMarketPrice" validate:"required"`
}

type YahooResponseBody struct {
	QuoteResponse struct {
		Result []struct {
			Symbol             string  `json:"symbol"`
			Bid                float64 `json:"bid"`
			Ask                float64 `json:"ask"`
			RegularMarketPrice float64 `json:"regularMarketPrice"`
		} `json:"result"`
	} `json:"quoteResponse"`
}

// Ensure that YahooTicker implements the Ticker interface at compile time.
var _ price_function.Ticker = (*YahooTicker)(nil)

func (t YahooTicker) WithPair(pair string) YahooTicker {
	t.Pair = pair
	return t
}

func (t YahooTicker) GetPair() string {
	// needs to be wrapped in quotes to be consistent with the API request format.
	return t.Pair
}

func (t YahooTicker) GetAskPrice() string {
	return t.AskPrice
}

func (t YahooTicker) GetBidPrice() string {
	return t.BidPrice
}

func (t YahooTicker) GetLastPrice() string {
	return t.LastPrice
}

// YahooPriceFunction transforms an API response from Yahoo into a map of tickers to prices that have been
// shifted by a market specific exponent.
func YahooPriceFunction(
	response *http.Response,
	tickerToExponent map[string]int32,
	resolver types.Resolver,
) (tickerToPrice map[string]uint64, unavailableTickers map[string]error, err error) {
	// Unmarshal response body into a list of tickers.
	var responseBody YahooResponseBody
	err = json.NewDecoder(response.Body).Decode(&responseBody)

	if err != nil {
		panic(fmt.Errorf("Error decoding yahoo response body: %w", err))
	}

	YahooTickers := make([]YahooTicker, 0, len(responseBody.QuoteResponse.Result))
	for _, result := range responseBody.QuoteResponse.Result {
		YahooTickers = append(YahooTickers, YahooTicker{
			Pair:      result.Symbol,
			AskPrice:  strconv.FormatFloat(result.Ask, 'f', -1, 64),
			BidPrice:  strconv.FormatFloat(result.Bid, 'f', -1, 64),
			LastPrice: strconv.FormatFloat(result.RegularMarketPrice, 'f', -1, 64),
		})
	}

	return price_function.GetMedianPricesFromTickers(
		YahooTickers,
		tickerToExponent,
		resolver,
	)
}
