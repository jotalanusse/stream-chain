package constants

import (
	clobtypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/clob/types"
)

var (
	// Long-term orders.
	LongTermOrder_Alice_Num0_Id0_Clob0_Buy5_Price10_GTB5 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 5},
	}
	LongTermOrder_Alice_Num0_Id0_Clob0_Buy5_Price5_GTBT5 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_005,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 5},
	}
	LongTermOrder_Alice_Num0_Id0_Clob0_Buy5_Price10_GTBT5 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 5},
	}
	LongTermOrder_Alice_Num0_Id0_Clob1_Buy5_Price10_GTBT5 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   1,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 5},
	}
	LongTermOrder_Alice_Num0_Id0_Clob0_Buy5_Price10_GTBT15 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
	}
	LongTermOrder_Alice_Num1_Id0_Clob0_Buy5_Price10_GTBT5 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num1,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 5},
	}
	LongTermOrder_Alice_Num1_Id1_Clob0_Buy02BTC_Price10_GTB15 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num1,
			ClientId:     1,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TwentyMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
	}
	LongTermOrder_Alice_Num1_Id2_Clob0_Sell02BTC_Price10_GTB15 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num1,
			ClientId:     2,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
	}
	LongTermOrder_Alice_Num0_Id0_Clob0_Buy100_Price10_GTBT15 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
	}
	LongTermOrder_Alice_Num0_Id0_Clob0_Buy5_Price10_GTBT20 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 20},
	}
	LongTermOrder_Alice_Num0_Id1_Clob1_Sell65_Price15_GTBT25 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     1,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   1,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     SixtyFiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_015,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 25},
	}
	LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     2,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     SixtyFiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 25},
	}
	LongTermOrder_Alice_Num1_Id4_Clob0_Buy10_Price45_GTBT20 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num1,
			ClientId:     4,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_045,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 20},
	}
	LongTermOrder_Alice_Num1_Id0_Clob0_Sell15_Price5_GTBT10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num1,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     FifteenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_005,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
	}
	LongTermOrder_Alice_Num1_Id1_Clob0_Sell25_Price30_GTBT10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num1,
			ClientId:     1,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyFiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_030,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
	}
	LongTermOrder_Alice_Num1_Id2_Clob0_Buy10_Price40_GTBT10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num1,
			ClientId:     2,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_040,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
	}
	LongTermOrder_Alice_Num0_Id1_Clob0_Sell20_Price10_GTBT10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     1,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
	}
	LongTermOrder_Alice_Num0_Id0_Clob0_Buy1BTC_Price50000_GTBT15 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
	}
	LongTermOrder_Alice_Num0_Id1_Clob0_Buy1BTC_Price50000_GTBT15 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     1,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
	}
	LongTermOrder_Bob_Num0_Id0_Clob0_Sell2_Price5_GTBT10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Bob_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwoQuantumSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_005,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
	}
	LongTermOrder_Bob_Num0_Id0_Clob0_Sell5_Price5_GTBT10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Bob_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_005,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
	}
	LongTermOrder_Bob_Num0_Id1_Clob0_Sell5_Price10_GTBT10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Bob_Num0,
			ClientId:     1,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
	}
	LongTermOrder_Bob_Num0_Id1_Clob0_Sell50_Price10_GTBT15 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Bob_Num0,
			ClientId:     1,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     FiftyQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
	}
	LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Bob_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TwentyFiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_030,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
	}
	LongTermOrder_Bob_Num0_Id0_Clob0_Buy35_Price30_GTBT11 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Bob_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     ThirtyFiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_030,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 11},
	}
	LongTermOrder_Bob_Num0_Id1_Clob0_Buy45_Price10_GTBT10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Bob_Num0,
			ClientId:     1,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FourtyFiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
	}
	LongTermOrder_Bob_Num0_Id2_Clob0_Buy15_Price5_GTBT10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Bob_Num0,
			ClientId:     2,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FifteenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_005,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
	}
	LongTermOrder_Bob_Num1_Id3_Clob0_Buy10_Price40_GTBT10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Bob_Num1,
			ClientId:     3,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_040,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
	}
	LongTermOrder_Carl_Num0_Id0_Clob0_Sell1BTC_Price50000_GTBT10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Carl_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
	}
	LongTermOrder_Carl_Num0_Id0_Clob0_Buy1BTC_Price49500_GTBT10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Carl_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_49_500,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
	}
	LongTermOrder_Carl_Num0_Id0_Clob0_Buy1BTC_Price50000_GTBT10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Carl_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
	}
	LongTermOrder_Dave_Num0_Id0_Clob0_Sell025BTC_Price50000_GTBT10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Dave_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyFiveMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
	}
	LongTermOrder_Dave_Num0_Id1_Clob0_Sell025BTC_Price50001_GTBT10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Dave_Num0,
			ClientId:     1,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TwentyFiveMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_001,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
	}
	LongTermOrder_Dave_Num0_Id0_Clob0_Sell1BTC_Price50000_GTBT10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Dave_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
	}
	LongTermOrder_Dave_Num0_Id0_Clob0_Buy1BTC_Price50000_GTBT10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Dave_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
	}

	// Conditional orders.
	ConditionalOrder_Alice_Num0_Id0_Clob0_Buy5_Price10_GTBT15_StopLoss20 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_020,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob0_Buy5_Price20_GTBT15_StopLoss20 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_020,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_020,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob0_Buy5_Price10_GTBT15_TakeProfit20 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_020,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob0_Buy5_Price10_GTBT15_TakeProfit10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_010,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob0_Buy5_Price10_GTBT15_TakeProfit25 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_025,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob0_Sell5_Price10_GTBT15_StopLoss20 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_020,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob0_Sell5_Price10_GTBT15_TakeProfit20 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_020,
	}
	ConditionalOrder_Alice_Num0_Id1_Clob0_Buy15_Price10_GTBT15_StopLoss20 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     1,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FifteenQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_020,
	}
	ConditionalOrder_Alice_Num0_Id1_Clob0_Buy15_Price25_GTBT15_StopLoss25 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     1,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FifteenQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_025,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_025,
	}
	ConditionalOrder_Alice_Num0_Id1_Clob0_Buy15_Price10_GTBT15_TakeProfit20 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     1,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FifteenQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_020,
	}
	ConditionalOrder_Alice_Num0_Id1_Clob0_Buy15_Price10_GTBT15_TakeProfit5 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     1,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FifteenQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_005,
	}
	ConditionalOrder_Alice_Num0_Id2_Clob0_Buy20_Price10_GTBT15_StopLoss20 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     2,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        TwentyQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_020,
	}
	ConditionalOrder_Alice_Num0_Id2_Clob0_Sell20_Price10_GTBT15_StopLoss20 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     2,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        TwentyQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_020,
	}
	ConditionalOrder_Alice_Num0_Id2_Clob0_Sell20_Price20_GTBT15_TakeProfit20 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     2,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        TwentyQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_020,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_020,
	}
	ConditionalOrder_Alice_Num0_Id2_Clob0_Buy20_Price10_GTBT15_TakeProfit10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     2,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        TwentyQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_010,
	}
	ConditionalOrder_Alice_Num0_Id3_Clob0_Buy25_Price10_GTBT15_StopLoss20 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     3,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        TwentyFiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_020,
	}
	ConditionalOrder_Alice_Num0_Id3_Clob0_Buy25_Price25_GTBT15_StopLoss25 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     3,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        TwentyFiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_025,
	}
	ConditionalOrder_Alice_Num0_Id3_Clob0_Buy25_Price10_GTBT15_TakeProfit20 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     3,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        TwentyFiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_020,
	}
	ConditionalOrder_Alice_Num0_Id3_Clob0_Sell25_Price10_GTBT15_StopLoss10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     3,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        TwentyFiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_010,
	}
	ConditionalOrder_Alice_Num0_Id3_Clob1_Buy25_Price10_GTBT15_StopLoss20 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     3,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   1,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        TwentyFiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_020,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob1_Buy5_Price10_GTBT15_StopLoss20 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   1,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_020,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob1_Buy5_Price10_GTBT15_TakeProfit20 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   1,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_020,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob1_Buy5_Price10_GTBT15_TakeProfit30 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   1,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_030,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob1_Sell5_Price10_GTBT15_StopLoss20 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   1,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_020,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob1_Sell5_Price10_GTBT15_TakeProfit20 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   1,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_020,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob0_Buy5_Price50_GTBT10_StopLoss51_FOK = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_050,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		TimeInForce:                     clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_051,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob0_Buy5_Price50_GTBT10_StopLoss51_IOC = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_050,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		TimeInForce:                     clobtypes.Order_TIME_IN_FORCE_IOC,
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_051,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob0_Buy5_Price10_GTBT20 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 20},
	}
	ConditionalOrder_Alice_Num1_Id0_Clob0_Sell5_Price10_GTBT15_StopLoss15 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num1,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_015,
	}
	ConditionalOrder_Alice_Num1_Id1_Clob0_Sell50_Price5_GTBT30_TakeProfit10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num1,
			ClientId:     1,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        FiftyQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_005,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 30},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_010,
	}
	ConditionalOrder_Alice_Num1_Id0_Clob0_Sell5_Price10_GTB15 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num1,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     FiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 15},
	}
	ConditionalOrder_Alice_Num1_Id1_Clob0_Sell50_Price5_GTB30 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num1,
			ClientId:     1,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     FiftyQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_005,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 30},
	}
	ConditionalOrder_Alice_Num1_Id1_Clob0_Sell50_Price5_GTB30_TakeProfit20 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num1,
			ClientId:     1,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        FiftyQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_005,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlock{GoodTilBlock: 30},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_0_000_020,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob0_Buy1BTC_Price50000_GTBT10_TP_48700 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        OneHundredMillionQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_48_700,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob0_Buy1BTC_Price50000_GTBT10_TP_49700 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        OneHundredMillionQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_49_700,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob0_Buy1BTC_Price50000_GTBT10_TP_49995 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        OneHundredMillionQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_49_995,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob0_Buy1BTC_Price50000_GTBT10_TP_49999 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        OneHundredMillionQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_49_999,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob0_Buy1BTC_Price50000_GTBT10_TP_49999_IOC = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        OneHundredMillionQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_49_999,
		TimeInForce:                     clobtypes.Order_TIME_IN_FORCE_IOC,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob0_Buy1BTC_Price50000_GTBT10_TP_49999_FOK = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        OneHundredMillionQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_49_999,
		TimeInForce:                     clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob0_Buy1BTC_Price50000_GTBT10_SL_50001 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        OneHundredMillionQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_50_001,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob0_Buy1BTC_Price50000_GTBT10_SL_50005 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        OneHundredMillionQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_50_005,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob0_Buy1BTC_Price50000_GTBT10_SL_50300 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        OneHundredMillionQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_50_300,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob0_Buy1BTC_Price50000_GTBT10_SL_51300 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        OneHundredMillionQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_51_300,
	}
	ConditionalOrder_Bob_Num0_Id0_Clob0_Sell1BTC_Price50000_GTBT10_TP_50001 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Bob_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        OneHundredMillionQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_50_001,
	}
	ConditionalOrder_Bob_Num0_Id0_Clob0_Sell1BTC_Price50000_GTBT10_TP_50005 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Bob_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        OneHundredMillionQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_50_005,
	}
	ConditionalOrder_Bob_Num0_Id0_Clob0_Sell1BTC_Price50000_GTBT10_TP_50300 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Bob_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        OneHundredMillionQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_50_300,
	}
	ConditionalOrder_Bob_Num0_Id0_Clob0_Sell1BTC_Price50000_GTBT10_TP_51300 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Bob_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        OneHundredMillionQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_51_300,
	}
	ConditionalOrder_Bob_Num0_Id0_Clob0_Sell1BTC_Price50000_GTBT10_SL_48700 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Bob_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        OneHundredMillionQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_48_700,
	}
	ConditionalOrder_Bob_Num0_Id0_Clob0_Sell1BTC_Price50000_GTBT10_SL_49700 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Bob_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        OneHundredMillionQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_49_700,
	}
	ConditionalOrder_Bob_Num0_Id0_Clob0_Sell1BTC_Price50000_GTBT10_SL_49995 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Bob_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        OneHundredMillionQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_49_995,
	}
	ConditionalOrder_Bob_Num0_Id0_Clob0_Sell1BTC_Price50000_GTBT10_SL_49999 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Bob_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        OneHundredMillionQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_49_999,
	}
	ConditionalOrder_Carl_Num0_Id0_Clob0_Sell1BTC_Price50000_GTBT10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Carl_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
	}
	ConditionalOrder_Carl_Num0_Id0_Clob0_Buy1BTC_Price50000_GTBT10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Carl_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
	}
	ConditionalOrder_Dave_Num0_Id0_Clob0_Sell1BTC_Price50000_GTBT10 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Dave_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
	}

	ConditionalOrder_Bob_Num0_Id0_Clob0_Sell10_Price10_GTBT10_PO_SL_15 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Bob_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        TenQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		TimeInForce:                     clobtypes.Order_TIME_IN_FORCE_POST_ONLY,
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_15,
	}
	ConditionalOrder_Alice_Num0_Id1_Clob0_Sell20_Price10_GTBT10_SL_15 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     1,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        TwentyQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_15,
	}
	ConditionalOrder_Bob_Num0_Id1_Clob0_Sell50_Price10_GTBT15_SL_15 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Bob_Num0,
			ClientId:     1,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        FiftyQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_15,
	}
	ConditionalOrder_Alice_Num0_Id0_Clob0_Buy5_Price10_GTBT5_SL_15 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiveQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_0_000_010,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 5},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_15,
	}
	ConditionalOrder_Dave_Num0_Id0_Clob0_Sell1BTC_Price50000_GTBT10_SL_50003 = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Dave_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        OneHundredMillionQuantumsSerializableInt,
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_50_003,
	}
	ConditionalOrder_Carl_Num0_Id0_Clob0_Buy05BTC_Price50000_GTBT10_SL_50003_FOK = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Carl_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiftyMillionQuantumsSerializableInt, // 0.5 BTC
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		TimeInForce:                     clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_50_003,
	}
	ConditionalOrder_Carl_Num0_Id0_Clob0_Sell05BTC_Price50000_GTBT10_TP_50003_FOK = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Carl_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        FiftyMillionQuantumsSerializableInt, // 0.5 BTC
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		TimeInForce:                     clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_50_003,
	}
	ConditionalOrder_Carl_Num0_Id0_Clob0_Buy05BTC_Price50000_GTBT10_SL_50003_IOC = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Carl_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiftyMillionQuantumsSerializableInt, // 0.5 BTC
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		TimeInForce:                     clobtypes.Order_TIME_IN_FORCE_IOC,
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_STOP_LOSS,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_50_003,
	}
	ConditionalOrder_Carl_Num0_Id0_Clob0_Buy05BTC_Price50000_GTBT10_TP_49999_PO = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Carl_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
			ClobPairId:   0,
		},
		Side:                            clobtypes.Order_SIDE_BUY,
		Quantums:                        FiftyMillionQuantumsSerializableInt, // 0.5 BTC
		Subticks:                        Dollars_Uusdc_50_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		TimeInForce:                     clobtypes.Order_TIME_IN_FORCE_POST_ONLY,
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_49_999,
	}

	// Conditional FOK/IOC RO orders.
	ConditionalOrder_Alice_Num1_Id1_Clob0_Sell05BTC_Price500000_GTBT20_TP_50001_IOC_RO = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num1,
			ClientId:     1,
			ClobPairId:   0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        FiftyMillionQuantumsSerializableInt, // 0.5 BTC
		Subticks:                        Dollars_Uusdc_500_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 20},
		TimeInForce:                     clobtypes.Order_TIME_IN_FORCE_IOC,
		ReduceOnly:                      true,
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_50_001,
	}
	ConditionalOrder_Alice_Num1_Id1_Clob0_Sell05BTC_Price500000_GTBT20_TP_50001_FOK_RO = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num1,
			ClientId:     1,
			ClobPairId:   0,
			OrderFlags:   clobtypes.OrderIdFlags_Conditional,
		},
		Side:                            clobtypes.Order_SIDE_SELL,
		Quantums:                        FiftyMillionQuantumsSerializableInt, // 0.5 BTC
		Subticks:                        Dollars_Uusdc_500_000,
		GoodTilOneof:                    &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 20},
		TimeInForce:                     clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
		ReduceOnly:                      true,
		ConditionType:                   clobtypes.Order_CONDITION_TYPE_TAKE_PROFIT,
		ConditionalOrderTriggerSubticks: Dollars_Uusdc_50_001,
	}

	// Long-Term post-only orders.
	LongTermOrder_Alice_Num0_Id0_Clob0_Buy100_Price10_GTBT15_PO = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 15},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_POST_ONLY,
	}
	LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25_PO = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Alice_Num0,
			ClientId:     2,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     SixtyFiveQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 25},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_POST_ONLY,
	}
	LongTermOrder_Bob_Num0_Id0_Clob0_Sell10_Price10_GTBT10_PO = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Bob_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_010,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_POST_ONLY,
	}
	LongTermOrder_Dave_Num0_Id0_Clob0_Buy1BTC_Price50000_GTBT10_PO = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Dave_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_BUY,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_POST_ONLY,
	}

	// Long-Term reduce-only orders.
	LongTermOrder_Bob_Num0_Id2_Clob0_Sell10_Price35_GTB20_RO = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Bob_Num0,
			ClientId:     2,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     TenQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_0_000_035,
		GoodTilOneof: &clobtypes.Order_GoodTilBlock{GoodTilBlock: 20},
		ReduceOnly:   true,
	}

	// Long-Term Fill Or Kill Orders.
	LongTermOrder_Carl_Num0_Id0_Clob0_Sell1BTC_Price50000_GTBT10_FOK = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Carl_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_FILL_OR_KILL,
	}

	// Long-Term Immediate Or Cancel Orders.
	LongTermOrder_Carl_Num0_Id0_Clob0_Sell1BTC_Price50000_GTBT10_IOC = clobtypes.Order{
		OrderId: clobtypes.OrderId{
			SubaccountId: Carl_Num0,
			ClientId:     0,
			OrderFlags:   clobtypes.OrderIdFlags_LongTerm,
			ClobPairId:   0,
		},
		Side:         clobtypes.Order_SIDE_SELL,
		Quantums:     OneHundredMillionQuantumsSerializableInt,
		Subticks:     Dollars_Uusdc_50_000,
		GoodTilOneof: &clobtypes.Order_GoodTilBlockTime{GoodTilBlockTime: 10},
		TimeInForce:  clobtypes.Order_TIME_IN_FORCE_IOC,
	}
)
