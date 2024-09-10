package stocks

import (
	"github.com/StreamFinance-Protocol/stream-chain/protocol/daemons/pricefeed/client/constants/exchange_common"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/daemons/pricefeed/client/types"
)

var (
	BinanceDetails = types.ExchangeQueryDetails{
		Exchange:      exchange_common.EXCHANGE_ID_STOCKS,
		Url:           "",
		PriceFunction: StocksPriceFunction,
		IsMultiMarket: true,
	}

	BinanceUSDetails = types.ExchangeQueryDetails{
		Exchange:      exchange_common.EXCHANGE_ID_STOCKS,
		Url:           "",
		PriceFunction: StocksPriceFunction,
		IsMultiMarket: true,
	}
)
