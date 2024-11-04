package yahoo

import (
	"github.com/StreamFinance-Protocol/stream-chain/protocol/daemons/pricefeed/client/constants/exchange_common"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/daemons/pricefeed/client/types"
)

var (
	YahooDetails = types.ExchangeQueryDetails{
		Exchange:      exchange_common.EXCHANGE_ID_YAHOO,
		Url:           "https://apidojo-yahoo-finance-v1.p.rapidapi.com/market/v2/get-quotes?region=US&symbols=$",
		PriceFunction: YahooPriceFunction,
		IsMultiMarket: true,
	}
)