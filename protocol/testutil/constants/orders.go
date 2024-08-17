package constants

import (
	"github.com/StreamFinance-Protocol/stream-chain/protocol/dtypes"
	clobtypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/clob/types"
)

var (
	// Short-term orders.
	Order_Alice_Num0_Id0_Clob0_Buy5_Price10_GTB15 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 0, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 15},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num0_Id0_Clob0_Buy10_Price10_GTB16 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 0, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        TenQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 16},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num0_Id0_Clob1_Buy5_Price10_GTB15 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 0, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 15},
	}
	Order_Alice_Num0_Id0_Clob2_Buy5_Price10_GTB15 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 0, ClobPairId: 2},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 15},
	}
	Order_Alice_Num0_Id0_Clob0_Buy5_Price10_GTB20 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 0, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num0_Id0_Clob0_Sell5_Price10_GTB20 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 0, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num0_Id0_Clob0_Buy5_Price5_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_005,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Alice_Num0_Id0_Clob0_Buy6_Price10_GTB20 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 0, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        SixQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num0_Id0_Clob0_Buy7_Price10_GTB20 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 0, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        SevenQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num0_Id0_Clob0_Buy35_Price10_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     ThirtyFiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Alice_Num0_Id1_Clob0_Sell5_Price15_GTB15 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 1, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_015,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 15},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num0_Id1_Clob0_Sell10_Price15_GTB15 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 1, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        TenQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_015,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 15},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num0_Id2_Clob1_Sell5_Price10_GTB15 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 2, ClobPairId: 1},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 15},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num0_Id3_Clob1_Sell5_Price10_GTB15 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 3, ClobPairId: 1},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 15},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num0_Id4_Clob1_Buy25_Price5_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 4, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TwentyFiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_005,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Alice_Num0_Id4_Clob2_Buy25_Price5_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 4, ClobPairId: 2},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TwentyFiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_005,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Alice_Num0_Id5_Clob1_Sell25_Price15_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 5, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyFiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Alice_Num0_Id6_Clob0_Buy25_Price5_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 6, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TwentyFiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_005,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Alice_Num0_Id7_Clob0_Sell25_Price15_GTB20 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 7, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        TwentyFiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_015,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num0_Id8_Clob1_Sell25_PriceMax_GTB20 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 8, ClobPairId: 1},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        TwentyFiveQuantumsSerializableInt,
		Subticks:                        dtypes.MaxUint256SerializableInt(),
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num0_Id9_Clob1_Buy15_Price45_GTB19 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 9, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FifteenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_045,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 19},
	}
	Order_Alice_Num0_Id10_Clob0_Sell25_Price15_GTB20 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 10, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        TwentyFiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_015,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num0_Id10_Clob0_Sell35_Price15_GTB25 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 10, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     ThirtyFiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 25},
	}
	Order_Alice_Num0_Id0_Clob0_Sell200BTC_Price101_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyBillionQuantumsSerializableInt, // 200 BTC
		Subticks:     Dollars_Uusdc_101,                    // $101
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Alice_Num0_Id0_Clob0_Sell100BTC_Price102_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenBillionQuantumsSerializableInt, // 100 BTC
		Subticks:     Dollars_Uusdc_102,                 // $102
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Alice_Num0_Id0_Clob0_Sell100BTC_Price106_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenBillionQuantumsSerializableInt, // 100 BTC
		Subticks:     Dollars_Uusdc_106,                 // $106
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Alice_Num1_Id0_Clob0_Sell10_Price10_GTB20 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 0, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        TenQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num1_Id0_Clob0_Sell10_Price10_GTB30 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 30},
	}
	Order_Alice_Num1_Id0_Clob0_Sell10_Price15_GTB20 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 0, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        TenQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_015,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num1_Id2_Clob1_Buy10_Price10_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 2, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Alice_Num1_Id2_Clob1_Buy10_Price10_GTB26 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 2, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 26},
	}
	Order_Alice_Num1_Id1_Clob1_Sell10_Price15_GTB20 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 1, ClobPairId: 1},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        TenQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_015,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num1_Id2_Clob1_Buy67_Price5_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 2, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     SixtySevenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_005,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Alice_Num1_Id3_Clob1_Buy7_Price5 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 3, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     SevenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_005,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Alice_Num1_Clob0_Id4_Buy10_Price45_GTB20 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 4, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        TenQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_045,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num1_Id5_Clob1_Sell50_Price40_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 5, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     FiftyQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_040,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Alice_Num1_Id6_Clob1_Sell15_Price22_GTB30 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 6, ClobPairId: 1},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        FifteenQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_022,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 30},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num1_Id7_Clob1_Buy35_PriceMax_GTB30 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 7, ClobPairId: 1},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        ThirtyFiveQuantumsSerializableInt,
		Subticks:                        dtypes.MaxUint256SerializableInt(),
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 30},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num1_Id8_Clob0_Buy15_Price25_GTB31 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 8, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FifteenQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_025,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 31},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num1_Id9_Clob0_Sell10_Price10_GTB31 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 9, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 31},
	}
	Order_Alice_Num1_Id10_Clob0_Buy5_Price30_GTB31 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 10, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_030,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 31},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num1_Id10_Clob0_Buy5_Price30_GTB32 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 10, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_030,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 32},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num1_Id10_Clob0_Buy6_Price30_GTB32 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 10, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        SixQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_030,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 32},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num1_Id10_Clob0_Buy7_Price30_GTB33 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 10, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     SevenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_030,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 33},
	}
	Order_Alice_Num1_Id10_Clob0_Buy10_Price30_GTB33 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 10, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        TenQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_030,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 33},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num1_Id10_Clob0_Buy15_Price30_GTB33 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 10, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FifteenQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_030,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 33},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num1_Id10_Clob0_Buy10_Price30_GTB34 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 10, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_030,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 34},
	}
	Order_Alice_Num1_Id10_Clob0_Buy5_Price30_GTB34 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 10, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_030,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 34},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num1_Id11_Clob1_Buy10_Price45_GTB20 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 11, ClobPairId: 1},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        TenQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_045,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num1_Id12_Clob0_Sell20_Price5_GTB25 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 12, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_005,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 25},
	}
	Order_Alice_Num1_Id13_Clob0_Buy30_Price50_GTB25 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 13, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        ThirtyQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_050,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 25},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num1_Id13_Clob0_Buy50_Price50_GTB30 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 13, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiftyQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_050,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 30},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num1_Id0_Clob0_Sell100_Price500000_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_500_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Alice_Num1_Id0_Clob0_Sell100_Price51000_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_51_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Alice_Num1_Id0_Clob0_Sell100_Price100000_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_1_000_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Alice_Num1_Id3_Clob0_Sell100_Price100000_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 3, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_1_000_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Alice_Num1_Id5_Clob1_Buy10_Price15_GTB23 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 5, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 23},
	}
	Order_Bob_Num0_Id0_Clob1_Sell10_Price15_GTB20 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 0, ClobPairId: 1},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        TenQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_015,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Bob_Num0_Id0_Clob2_Sell10_Price15_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 0, ClobPairId: 2},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Bob_Num0_Id1_Clob1_Sell11_Price16_GTB18 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 1, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     ElevenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_016,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 18},
	}
	Order_Bob_Num0_Id1_Clob1_Sell11_Price16_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 1, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     ElevenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_016,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Bob_Num0_Id2_Clob1_Sell12_Price13_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 2, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwelveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_013,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Bob_Num0_Id0_Clob0_Sell1BTC_Price50000_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredMillionQuantumsSerializableInt, // 1 BTC
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Bob_Num0_Id3_Clob1_Buy10_Price10_GTB20 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 3, ClobPairId: 1},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        TenQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Bob_Num0_Id4_Clob1_Buy20_Price35_GTB22 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 4, ClobPairId: 1},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        TwentyQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_035,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 22},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Bob_Num0_Id5_Clob0_Buy20_Price10_GTB22 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 5, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TwentyQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 22},
	}
	Order_Bob_Num0_Id6_Clob0_Buy20_Price1000_GTB22 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 6, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TwentyQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_001,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 22},
	}
	Order_Bob_Num0_Id7_Clob0_Buy20_Price10000_GTB22 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 7, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TwentyQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 22},
	}
	Order_Bob_Num0_Id8_Clob1_Sell5_Price10_GTB22 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 8, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 22},
	}
	Order_Bob_Num0_Id8_Clob1_Sell20_Price10_GTB22 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 8, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 22},
	}
	Order_Bob_Num0_Id8_Clob0_Sell20_Price10_GTB22 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 8, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        TwentyQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 22},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Bob_Num0_Id9_Clob0_Sell20_Price1000 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 9, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_001,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 22},
	}
	Order_Bob_Num0_Id10_Clob0_Sell20_Price10000 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 10, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 22},
	}
	Order_Bob_Num0_Id11_Clob1_Sell5_Price15_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 11, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Bob_Num0_Id11_Clob1_Buy5_Price40_GTB20 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 11, ClobPairId: 1},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_040,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Bob_Num0_Id12_Clob0_Buy5_Price5_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 12, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_005,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Bob_Num0_Id12_Clob0_Buy5_Price40_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 12, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_040,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Bob_Num0_Id12_Clob1_Buy5_Price40_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 12, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_040,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Bob_Num0_Id11_Clob1_Buy5_Price40_GTB32 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 11, ClobPairId: 1},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_040,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 32},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Bob_Num0_Id12_Clob0_Sell20_Price5_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 12, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_005,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Bob_Num0_Id12_Clob0_Sell20_Price15_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 12, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Bob_Num0_Id12_Clob0_Sell20_Price35_GTB32 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 12, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_035,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 32},
	}
	Order_Bob_Num0_Id13_Clob0_Sell35_Price35_GTB30 = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 13, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        ThirtyFiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_035,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 30},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Bob_Num0_Id14_Clob0_Sell10_Price10_GTB25 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 14, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 25},
	}
	Order_Bob_Num0_Id1_Clob0_Buy35_Price55_GTB32 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     ThirtyFiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_055,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 32},
	}
	Order_Bob_Num0_Id2_Clob0_Sell25_Price95_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 2, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyFiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_095,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Bob_Num0_Id1_Clob0_Buy100BTC_Price98_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenBillionQuantumsSerializableInt, // 100 BTC
		Subticks:     Dollars_Uusdc_98,                  // $98
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Bob_Num0_Id1_Clob0_Buy100BTC_Price99_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenBillionQuantumsSerializableInt, // 100 BTC
		Subticks:     Dollars_Uusdc_99,                  // $99
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Bob_Num0_Id0_Clob0_Sell100BTC_Price101_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenBillionQuantumsSerializableInt, // 100 BTC
		Subticks:     Dollars_Uusdc_101,                 // $101
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Bob_Num0_Id0_Clob0_Sell200BTC_Price101_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyBillionQuantumsSerializableInt, // 200 BTC
		Subticks:     Dollars_Uusdc_101,                    // $101
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Bob_Num1_Id1_Clob1_Sell25_Price85_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num1, ClientId: 1, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyFiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_085,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id0_Clob0_Buy1BTC_Price5subticks_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_005,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id0_Clob0_Sell1BTC_Price5000_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_5_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id0_Clob0_Buy1BTC_Price500000_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_500_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id0_Clob0_Buy025BTC_Price500000_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TwentyFiveMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_500_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id0_Clob0_Sell1BTC_Price500000_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_500_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id0_Clob0_Buy70_Price500000_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     SeventyQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_500_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id0_Clob0_Buy110_Price500000_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredAndTenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_500_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id0_Clob0_Buy110_Price50000_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredAndTenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id0_Clob0_Buy10_Price500000_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_500_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Carl_Num0_Id0_Clob0_Buy80_Price500000_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     EightyQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_500_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Carl_Num0_Id0_Clob0_Buy10_Price50000_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_500_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Carl_Num0_Id0_Clob0_Buy110_Price50000_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredAndTenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_500_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Carl_Num0_Id2_Clob0_Sell5_Price10_GTB15 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 2, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 15},
	}
	Order_Carl_Num0_Id1_Clob0_Buy01BTC_Price49500_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_49_500,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id0_Clob0_Buy1BTC_Price49500_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_49_500,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id0_Clob0_Buy1BTC_Price49800_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_49_800,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id0_Clob0_Buy1BTC_Price50000_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id2_Clob0_Buy1BTC_Price50500_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 2, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_500,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id1_Clob0_Buy1BTC_Price49999 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_49_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id2_Clob0_Buy05BTC_Price50000 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 2, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FiftyMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id3_Clob0_Buy025BTC_Price49500 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 3, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TwentyFiveMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_49_500,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id3_Clob0_Buy025BTC_Price49800 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 3, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TwentyFiveMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_49_800,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id3_Clob0_Buy025BTC_Price50000 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 3, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TwentyFiveMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id4_Clob0_Buy05BTC_Price40000 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 4, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FiftyMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_40_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id5_Clob0_Buy2BTC_Price50000 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 5, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TwoHundredyMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id2_Clob1_Buy10ETH_Price3000 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 2, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenBillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_3_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id3_Clob1_Buy1ETH_Price3000 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 3, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneBillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_3_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id4_Clob1_Buy01ETH_Price3000 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 4, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_3_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id0_Clob0_Buy10QtBTC_Price100000QuoteQt = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_100_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Carl_Num0_Id0_Clob0_Buy10QtBTC_Price100001QuoteQt = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_100_001,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Carl_Num0_Id0_Clob0_Sell1kQtBTC_Price50000 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id1_Clob0_Sell1kQtBTC_Price50000 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num0_Id0_Clob0_Buy100BTC_Price99_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenBillionQuantumsSerializableInt, // 100 BTC
		Subticks:     Dollars_Uusdc_99,                  // $99
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Carl_Num0_Id1_Clob0_Buy100BTC_Price100_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenBillionQuantumsSerializableInt, // 100 BTC
		Subticks:     Dollars_Uusdc_100,                 // $100
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Carl_Num0_Id1_Clob0_Buy100BTC_Price101_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenBillionQuantumsSerializableInt, // 100 BTC
		Subticks:     Dollars_Uusdc_101,                 // $101
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Carl_Num1_Id0_Clob0_Buy1kQtBTC_Price50000 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num1, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num1_Id0_Clob0_Buy1kQtBTC_Price60000 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num1, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_60_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num1_Id1_Clob0_Buy1kQtBTC_Price50000 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num1, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num1_Id0_Clob0_Buy1BTC_Price50000 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num1, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_49_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num1_Id0_Clob0_Buy1BTC_Price50000_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num1, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num1_Id0_Clob0_Buy1BTC_Price50003_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num1, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_003,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num1_Id0_Clob0_Buy1BTC_Price50500_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num1, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_500,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Carl_Num1_Id0_Clob0_Buy1BTC_Price51500_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num1, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_51_500,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Dave_Num0_Id2_Clob0_Sell1BTC_Price49500_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 2, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_49_500,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Dave_Num0_Id0_Clob0_Sell1BTC_Price49999_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_49_999,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Dave_Num0_Id0_Clob0_Sell1BTC_Price50000_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Dave_Num0_Id0_Clob0_Sell1BTC_Price50498_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_498,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Dave_Num0_Id1_Clob0_Sell01BTC_Price50500_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_500,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Dave_Num0_Id0_Clob0_Sell1BTC_Price50500_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_500,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Dave_Num0_Id0_Clob0_Sell1BTC_Price60000_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_60_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Dave_Num0_Id1_Clob0_Sell025BTC_Price50000_GTB11 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyFiveMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 11},
	}
	Order_Dave_Num0_Id1_Clob0_Sell025BTC_Price50498_GTB11 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyFiveMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_498,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 11},
	}
	Order_Dave_Num0_Id1_Clob0_Sell025BTC_Price50500_GTB11 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyFiveMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_500,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 11},
	}
	// Replacement for the above order
	Order_Dave_Num0_Id0_Clob0_Sell1BTC_Price50000 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Dave_Num0_Id1_Clob0_Sell025BTC_Price50000_GTB12 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyFiveMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 12},
	}
	Order_Dave_Num0_Id2_Clob0_Sell025BTC_Price50000_GTB12 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 2, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyFiveMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 12},
	}
	Order_Dave_Num0_Id2_Clob0_Sell025BTC_Price50500_GTB12 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 2, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyFiveMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_500,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 12},
	}
	Order_Dave_Num0_Id0_Clob0_Buy100BTC_Price101_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenBillionQuantumsSerializableInt, // 100 BTC
		Subticks:     Dollars_Uusdc_101,                 // $101
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Dave_Num0_Id0_Clob0_Buy100BTC_Price102_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenBillionQuantumsSerializableInt, // 100 BTC
		Subticks:     Dollars_Uusdc_102,                 // $102
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Dave_Num0_Id1_Clob0_Buy100BTC_Price104_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenBillionQuantumsSerializableInt, // 100 BTC
		Subticks:     Dollars_Uusdc_104,                 // $104
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Dave_Num0_Id3_Clob1_Sell1ETH_Price3000 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 3, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneBillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_3_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Dave_Num0_Id4_Clob1_Sell1ETH_Price3000 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 4, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneBillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_3_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Dave_Num0_Id4_Clob1_Sell1ETH_Price3020 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 4, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneBillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_3_020,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Dave_Num0_Id4_Clob1_Sell1ETH_Price3030 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 4, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneBillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_3_030,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Dave_Num1_Id0_Clob0_Sell1BTC_Price48500_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num1, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_48_500,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Dave_Num1_Id0_Clob0_Sell1BTC_Price49500_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num1, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_49_500,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Dave_Num1_Id0_Clob0_Sell1BTC_Price49997_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num1, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_49_997,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Dave_Num1_Id0_Clob0_Sell025BTC_Price49999_GTB10 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num1, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyFiveMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_49_999,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
	}
	Order_Dave_Num1_Id0_Clob0_Buy100BTC_Price101_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num1, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenBillionQuantumsSerializableInt, // 100 BTC
		Subticks:     Dollars_Uusdc_101,                 // $101
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Dave_Num1_Id0_Clob0_Sell100BTC_Price101_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num1, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenBillionQuantumsSerializableInt, // 100 BTC
		Subticks:     Dollars_Uusdc_101,                 // $101
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}
	Order_Dave_Num1_Id0_Clob0_Sell100BTC_Price102_GTB20 = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num1, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenBillionQuantumsSerializableInt, // 100 BTC
		Subticks:     Dollars_Uusdc_102,                 // $102
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
	}

	// IOC orders.
	Order_Alice_Num0_Id1_Clob0_Buy5_Price15_GTB20_IOC = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_IOC,
	}
	Order_Alice_Num0_Id1_Clob1_Buy5_Price15_GTB20_IOC = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 1, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_IOC,
	}
	Order_Alice_Num0_Id1_Clob1_Sell5_Price15_GTB20_IOC = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 1, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_IOC,
	}
	Order_Alice_Num0_Id1_Clob1_Buy10_Price15_GTB20_IOC = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 1, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_IOC,
	}
	Order_Alice_Num0_Id1_Clob1_Sell10_Price15_GTB20_IOC = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 1, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_IOC,
	}
	Order_Alice_Num1_Id1_Clob1_Sell10_Price15_GTB20_IOC = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 1, ClobPairId: 1},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        TenQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_015,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:                     clobtypes.Order_TIME_IN_FORCE_IOC,
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
	}
	Order_Alice_Num1_Id1_Clob1_Sell10_Price15_GTB21_IOC = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 1, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 21},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_IOC,
	}

	// Fill-or-kill orders.
	Order_Alice_Num0_Id0_Clob1_Sell10_Price15_GTB20_FOK = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 0, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
	}
	Order_Alice_Num0_Id0_Clob1_Buy10_Price15_GTB20_FOK = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 0, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
	}
	Order_Alice_Num0_Id0_Clob1_Buy20_Price15_GTB20_FOK = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 0, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TwentyQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
	}
	Order_Alice_Num1_Id1_Clob1_Sell10_Price15_GTB20_FOK = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 1, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
	}
	Order_Alice_Num1_Id1_Clob1_Sell10_Price15_GTB21_FOK = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 1, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 21},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
	}
	Order_Alice_Num0_Id0_Clob0_Buy1BTC_Price50000_GTB10_FOK = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt, // 1 BTC
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
	}
	Order_Bob_Num0_Id1_Clob1_Buy20_Price35_GTB22_FOK = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 1, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TwentyQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_035,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 22},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
	}
	Order_Carl_Num0_Id0_Clob0_Buy05BTC_Price50000_GTB10_FOK = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FiftyMillionQuantumsSerializableInt, // 0.5 BTC
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
	}
	Order_Carl_Num0_Id0_Clob0_Buy1BTC_Price50000_GTB10_FOK = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt, // 1 BTC
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
	}
	Order_Carl_Num0_Id0_Clob0_Buy1BTC_Price50000_GTB20_FOK = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt, // 1 BTC
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
	}
	Order_Carl_Num0_Id0_Clob0_Buy075BTC_Price50000_GTB11_FOK = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     SeventyFiveMillionQuantumsSerializableInt, // 0.75 BTC
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 11},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
	}
	Order_Dave_Num0_Id0_Clob0_Sell1BTC_Price50000_GTB10_FOK = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredMillionQuantumsSerializableInt, // 1 BTC
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
	}
	// FOK + RO orders.
	Order_Alice_Num1_Id1_Clob1_Sell10_Price15_GTB20_FOK_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 1, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
		ReduceOnly:   true,
	}
	Order_Alice_Num1_Id1_Clob1_Buy10_Price15_GTB20_FOK_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 1, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
		ReduceOnly:   true,
	}
	Order_Alice_Num1_Id1_Clob0_Sell10_Price15_GTB20_FOK_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
		ReduceOnly:   true,
	}
	Order_Alice_Num1_Id1_Clob0_Buy10_Price15_GTB20_FOK_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
		ReduceOnly:   true,
	}
	Order_Alice_Num1_Id0_Clob0_Buy110_Price50000_GTB21_FOK_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredAndTenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 21},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
		ReduceOnly:   true,
	}
	Order_Alice_Num1_Id0_Clob0_Sell110_Price50000_GTB21_FOK_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredAndTenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 21},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
		ReduceOnly:   true,
	}
	Order_Alice_Num1_Id1_Clob0_Sell15_Price500000_GTB20_FOK_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     FifteenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_500_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
		ReduceOnly:   true,
	}
	// IOC + RO orders.
	Order_Alice_Num1_Id1_Clob1_Sell10_Price15_GTB20_IOC_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 1, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_IOC,
		ReduceOnly:   true,
	}
	Order_Alice_Num1_Id1_Clob1_Buy10_Price15_GTB20_IOC_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 1, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_IOC,
		ReduceOnly:   true,
	}
	Order_Alice_Num1_Id1_Clob0_Sell10_Price15_GTB20_IOC_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_IOC,
		ReduceOnly:   true,
	}
	Order_Alice_Num1_Id1_Clob0_Buy10_Price15_GTB20_IOC_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_IOC,
		ReduceOnly:   true,
	}
	Order_Alice_Num1_Id1_Clob0_Sell15_Price500000_GTB20_IOC_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     FifteenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_500_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_IOC,
		ReduceOnly:   true,
	}
	Order_Alice_Num1_Id0_Clob0_Sell110_Price50000_GTB21_IOC_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredAndTenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 21},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_IOC,
		ReduceOnly:   true,
	}
	Order_Alice_Num1_Id0_Clob0_Buy110_Price50000_GTB21_IOC_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredAndTenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 21},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_IOC,
		ReduceOnly:   true,
	}

	// Reduce-only orders.
	Order_Alice_Num1_Id1_Clob0_Sell10_Price15_GTB20_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ReduceOnly:   true,
	}
	Order_Alice_Num1_Id2_Clob0_Buy20_Price30_GTB20_RO = clobtypes.Order{
		OrderId:                         clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 2, ClobPairId: 0},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        TwentyQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_030,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ConditionalOrderTriggerSubticks: dtypes.NewInt(0),
		ReduceOnly:                      true,
	}
	Order_Alice_Num1_Id3_Clob1_Buy30_Price35_GTB25_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 3, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     ThirtyQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_035,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 25},
		ReduceOnly:   true,
	}
	Order_Alice_Num1_Id4_Clob0_Sell15_Price20_GTB20_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 4, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     FifteenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_020,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ReduceOnly:   true,
	}
	Order_Alice_Num1_Id5_Clob1_Sell10_Price15_GTB20_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 5, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ReduceOnly:   true,
	}
	Order_Alice_Num1_Id6_Clob0_Buy10_Price5_GTB20_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 6, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_005,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ReduceOnly:   true,
	}
	Order_Bob_Num0_Id1_Clob0_Sell15_Price50_GTB20_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     FifteenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_050,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ReduceOnly:   true,
	}
	Order_Bob_Num0_Id2_Clob0_Sell10_Price35_GTB20_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 2, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_035,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ReduceOnly:   true,
	}
	Order_Bob_Num0_Id3_Clob0_Sell20_Price10_GTB20_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 3, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ReduceOnly:   true,
	}
	Order_Carl_Num0_Id0_Clob0_Buy1BTC_Price50000_GTB10_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Carl_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
		ReduceOnly:   true,
	}
	Order_Dave_Num0_Id0_Clob0_Sell1BTC_Price50000_GTB10_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 0, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 10},
		ReduceOnly:   true,
	}
	Order_Dave_Num0_Id2_Clob0_Sell25BTC_Price50000_GTB12_RO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Dave_Num0, ClientId: 2, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyFiveMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 12},
		ReduceOnly:   true,
	}

	// Post-only orders.
	Order_Alice_Num0_Id1_Clob0_Sell15_Price10_GTB18_PO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     FifteenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 18},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_POST_ONLY,
	}
	Order_Alice_Num0_Id1_Clob0_Buy15_Price10_GTB18_PO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num0, ClientId: 1, ClobPairId: 0},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FifteenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 18},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_POST_ONLY,
	}
	Order_Alice_Num1_Id1_Clob1_Sell10_Price15_GTB20_PO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 1, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_POST_ONLY,
	}
	Order_Alice_Num1_Id4_Clob1_Sell10_Price15_GTB20_PO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Alice_Num1, ClientId: 4, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_POST_ONLY,
	}
	Order_Bob_Num0_Id4_Clob1_Buy20_Price35_GTB22_PO = clobtypes.Order{
		OrderId:      clobtypes.OrderId{SubaccountId: Bob_Num0, ClientId: 4, ClobPairId: 1},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TwentyQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_035,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 22},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_POST_ONLY,
	}
)
