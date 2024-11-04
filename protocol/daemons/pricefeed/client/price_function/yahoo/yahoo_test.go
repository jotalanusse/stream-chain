
package yahoo_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"
	"os"

	"github.com/StreamFinance-Protocol/stream-chain/protocol/daemons/pricefeed/client/price_function/testutil"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/daemons/pricefeed/client/price_function/yahoo"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/lib"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/daemons/pricefeed"
	"github.com/stretchr/testify/require"
)

// Test tickers for Binance.
const (
	AAPL_TICKER = "AAPL"
)

// Test exponent maps.
var (
	YahooExponentMap = map[string]int32{
		AAPL_TICKER: -5,
	}
)

func TestYahooPriceFunction_Mixed(t *testing.T) {
	// Test response strings.

	tests := map[string]struct {
		// parameters
		responseJsonString  string
		exponentMap         map[string]int32
		medianFunctionFails bool

		// expectations
		expectedPriceMap       map[string]uint64
		expectedUnavailableMap map[string]error
		expectedError          error
	}{
		"Success - integers": {
			responseJsonString: `{"quoteResponse": {"result": [{"symbol":"AAPL","regularMarketPrice":"1780.00","bid":"1780.00","ask":"1780.25000000"}]}}`,
			exponentMap:        YahooExponentMap,
			expectedPriceMap: map[string]uint64{
				AAPL_TICKER: uint64(1_780_000_00),
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			response := getYahooResponse()

			var prices map[string]uint64
			var unavailable map[string]error
			var err error
			if tc.medianFunctionFails {
				prices, unavailable, err = yahoo.YahooPriceFunction(response, tc.exponentMap, testutil.MedianErr)
			} else {
				prices, unavailable, err = yahoo.YahooPriceFunction(response, tc.exponentMap, lib.Median[uint64])
			}
			fmt.Println(response)

			if tc.expectedError != nil {
				require.EqualError(t, err, tc.expectedError.Error())
				require.Nil(t, prices)
				require.Nil(t, unavailable)
			} else {
				require.Equal(t, tc.expectedPriceMap, prices)
				pricefeed.ErrorMapsEqual(t, tc.expectedUnavailableMap, unavailable)
				require.NoError(t, err)
			}
			response.Body.Close()
		})

	}
}

func getYahooResponse() *http.Response {
	url := "https://apidojo-yahoo-finance-v1.p.rapidapi.com/market/v2/get-quotes?region=US&symbols=AAPL"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("x-rapidapi-host", "apidojo-yahoo-finance-v1.p.rapidapi.com")
	apiKey := os.Getenv("YAHOO_FINANCE_API_KEY")
	if apiKey == "" {
		panic("YAHOO_FINANCE_API_KEY environment variable is not set")
	}
	req.Header.Add("x-rapidapi-key", apiKey) // Replace with your actual API key
	client := &http.Client{}
	fmt.Println(req)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Error: received status code %d", resp.StatusCode))
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(body)
	// Create a new response with the body read
	newResp := &http.Response{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Proto:      resp.Proto,
		ProtoMajor: resp.ProtoMajor,
		ProtoMinor: resp.ProtoMinor,
		Header:     resp.Header,
		Body:       io.NopCloser(bytes.NewBuffer(body)),
	}

	return newResp
}