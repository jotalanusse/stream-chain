package memclob

import (
	"testing"

	clobtest "github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/clob"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/constants"
	testutil_memclob "github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/memclob"
	sdktest "github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/sdk"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/clob/types"
	satypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/subaccounts/types"
)

func TestPlaceOrder_LongTerm(t *testing.T) {
	ctx, _, _ := sdktest.NewSdkContextWithMultistore()
	ctx = ctx.WithIsCheckTx(true)
	tests := map[string]struct {
		// State.
		placedMatchableOrders  []types.MatchableOrder
		collateralizationCheck map[int]testutil_memclob.CollateralizationCheck

		// Parameters.
		order types.Order

		// Expectations.
		expectedFilledSize         satypes.BaseQuantums
		expectedOrderStatus        types.OrderStatus
		expectedRemainingBids      []OrderWithRemainingSize
		expectedRemainingAsks      []OrderWithRemainingSize
		expectedOperations         []types.Operation
		expectedInternalOperations []types.InternalOperation
		expectedErr                error
	}{
		"Can place a valid Long-Term buy order on an empty orderbook": {
			placedMatchableOrders: []types.MatchableOrder{},
			collateralizationCheck: map[int]testutil_memclob.CollateralizationCheck{
				0: {
					CollatCheck: map[satypes.SubaccountId][]types.PendingOpenOrder{
						constants.Alice_Num0: {
							{
								RemainingQuantums: constants.LongTermOrder_Alice_Num0_Id0_Clob0_Buy5_Price10_GTBT15.GetBaseQuantums(),
								IsBuy:             constants.LongTermOrder_Alice_Num0_Id0_Clob0_Buy5_Price10_GTBT15.IsBuy(),
								Subticks:          constants.LongTermOrder_Alice_Num0_Id0_Clob0_Buy5_Price10_GTBT15.GetOrderSubticks(),
								ClobPairId:        constants.LongTermOrder_Alice_Num0_Id0_Clob0_Buy5_Price10_GTBT15.GetClobPairId(),
							},
						},
					},
					Result: map[satypes.SubaccountId]satypes.UpdateResult{
						constants.Alice_Num0: satypes.Success,
					},
				},
			},

			order: constants.LongTermOrder_Alice_Num0_Id0_Clob0_Buy5_Price10_GTBT15,

			expectedFilledSize:  constants.BaseQuantums_0,
			expectedOrderStatus: types.Success,
			expectedRemainingBids: []OrderWithRemainingSize{
				{
					Order:         constants.LongTermOrder_Alice_Num0_Id0_Clob0_Buy5_Price10_GTBT15,
					RemainingSize: constants.BaseQuantums_5,
				},
			},
			expectedRemainingAsks: []OrderWithRemainingSize{},
			expectedOperations: []types.Operation{
				clobtest.NewOrderPlacementOperation(
					constants.LongTermOrder_Alice_Num0_Id0_Clob0_Buy5_Price10_GTBT15,
				),
			},
			expectedInternalOperations: []types.InternalOperation{},
		},
		`Matches a buy order when it overlaps the orderbook, and with no orders on the other side it places the remaining
		size on the orderbook`: {
			placedMatchableOrders: []types.MatchableOrder{
				&constants.Order_Alice_Num0_Id1_Clob0_Sell10_Price15_GTB15,
			},
			collateralizationCheck: map[int]testutil_memclob.CollateralizationCheck{
				0: {
					CollatCheck: map[satypes.SubaccountId][]types.PendingOpenOrder{
						constants.Alice_Num0: {
							{
								RemainingQuantums: constants.BaseQuantums_10,
								IsBuy:             false,
								Subticks:          15,
								ClobPairId:        0,
							},
						},
						constants.Bob_Num0: {
							{
								RemainingQuantums: constants.BaseQuantums_10,
								IsBuy:             true,
								Subticks:          15,
								ClobPairId:        0,
							},
						},
					},
					Result: map[satypes.SubaccountId]satypes.UpdateResult{
						constants.Alice_Num0: satypes.Success,
						constants.Bob_Num0:   satypes.Success,
					},
				},
				1: {
					CollatCheck: map[satypes.SubaccountId][]types.PendingOpenOrder{
						constants.Bob_Num0: {
							{
								RemainingQuantums: constants.BaseQuantums_15,
								IsBuy:             true,
								Subticks:          30,
								ClobPairId:        0,
							},
						},
					},
					Result: map[satypes.SubaccountId]satypes.UpdateResult{
						constants.Bob_Num0: satypes.Success,
					},
				},
			},

			order: constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,

			expectedFilledSize:  constants.BaseQuantums_10,
			expectedOrderStatus: types.Success,
			expectedOperations: []types.Operation{
				clobtest.NewOrderPlacementOperation(
					constants.Order_Alice_Num0_Id1_Clob0_Sell10_Price15_GTB15,
				),
				clobtest.NewOrderPlacementOperation(
					constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,
				),
				clobtest.NewMatchOperation(
					&constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,
					[]types.MakerFill{
						{
							MakerOrderId: constants.Order_Alice_Num0_Id1_Clob0_Sell10_Price15_GTB15.OrderId,
							FillAmount:   constants.TenQuantumsSerializableInt,
						},
					},
				),
			},
			expectedInternalOperations: []types.InternalOperation{
				types.NewShortTermOrderPlacementInternalOperation(
					constants.Order_Alice_Num0_Id1_Clob0_Sell10_Price15_GTB15,
				),
				types.NewPreexistingStatefulOrderPlacementInternalOperation(
					constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,
				),
				types.NewMatchOrdersInternalOperation(
					constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,
					[]types.MakerFill{
						{
							MakerOrderId: constants.Order_Alice_Num0_Id1_Clob0_Sell10_Price15_GTB15.OrderId,
							FillAmount:   constants.TenQuantumsSerializableInt,
						},
					},
				),
			},
			expectedRemainingBids: []OrderWithRemainingSize{
				{
					Order:         constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,
					RemainingSize: constants.BaseQuantums_15,
				},
			},
			expectedRemainingAsks: []OrderWithRemainingSize{},
		},
		`Fully matches a Long-Term sell order with other Long-Term buy orders when it overlaps the
		orderbook`: {
			placedMatchableOrders: []types.MatchableOrder{
				&constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,
				&constants.LongTermOrder_Bob_Num0_Id1_Clob0_Buy45_Price10_GTBT10,
			},
			collateralizationCheck: map[int]testutil_memclob.CollateralizationCheck{
				0: {
					CollatCheck: map[satypes.SubaccountId][]types.PendingOpenOrder{
						constants.Alice_Num0: {
							{
								RemainingQuantums: constants.BaseQuantums_25,
								IsBuy:             false,
								Subticks:          30,
								ClobPairId:        0,
							},
						},
						constants.Bob_Num0: {
							{
								RemainingQuantums: constants.BaseQuantums_25,
								IsBuy:             true,
								Subticks:          30,
								ClobPairId:        0,
							},
						},
					},
					Result: map[satypes.SubaccountId]satypes.UpdateResult{
						constants.Alice_Num0: satypes.Success,
						constants.Bob_Num0:   satypes.Success,
					},
				},
				1: {
					CollatCheck: map[satypes.SubaccountId][]types.PendingOpenOrder{
						constants.Alice_Num0: {
							{
								RemainingQuantums: constants.BaseQuantums_40,
								IsBuy:             false,
								Subticks:          10,
								ClobPairId:        0,
							},
						},
						constants.Bob_Num0: {
							{
								RemainingQuantums: constants.BaseQuantums_40,
								IsBuy:             true,
								Subticks:          10,
								ClobPairId:        0,
							},
						},
					},
					Result: map[satypes.SubaccountId]satypes.UpdateResult{
						constants.Alice_Num0: satypes.Success,
						constants.Bob_Num0:   satypes.Success,
					},
				},
			},

			order: constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25,

			expectedFilledSize:  constants.BaseQuantums_65,
			expectedOrderStatus: types.Success,
			expectedOperations: []types.Operation{
				clobtest.NewOrderPlacementOperation(
					constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,
				),
				clobtest.NewOrderPlacementOperation(
					constants.LongTermOrder_Bob_Num0_Id1_Clob0_Buy45_Price10_GTBT10,
				),
				clobtest.NewOrderPlacementOperation(
					constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25,
				),
				clobtest.NewMatchOperation(
					&constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25,
					[]types.MakerFill{
						{
							MakerOrderId: constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10.OrderId,
							FillAmount:   constants.TwentyFiveQuantumsSerializableInt,
						},
						{
							MakerOrderId: constants.LongTermOrder_Bob_Num0_Id1_Clob0_Buy45_Price10_GTBT10.OrderId,
							FillAmount:   constants.FourtyQuantumsSerializableInt,
						},
					},
				),
			},
			expectedInternalOperations: []types.InternalOperation{
				types.NewPreexistingStatefulOrderPlacementInternalOperation(
					constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,
				),
				types.NewPreexistingStatefulOrderPlacementInternalOperation(
					constants.LongTermOrder_Bob_Num0_Id1_Clob0_Buy45_Price10_GTBT10,
				),
				types.NewPreexistingStatefulOrderPlacementInternalOperation(
					constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25,
				),
				types.NewMatchOrdersInternalOperation(
					constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25,
					[]types.MakerFill{
						{
							MakerOrderId: constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10.OrderId,
							FillAmount:   constants.TwentyFiveQuantumsSerializableInt,
						},
						{
							MakerOrderId: constants.LongTermOrder_Bob_Num0_Id1_Clob0_Buy45_Price10_GTBT10.OrderId,
							FillAmount:   constants.FourtyQuantumsSerializableInt,
						},
					},
				),
			},
			expectedRemainingBids: []OrderWithRemainingSize{
				{
					Order:         constants.LongTermOrder_Bob_Num0_Id1_Clob0_Buy45_Price10_GTBT10,
					RemainingSize: constants.BaseQuantums_5,
				},
			},
			expectedRemainingAsks: []OrderWithRemainingSize{},
		},
		`Short-Term taker order can fully match with Long-Term maker order`: {
			placedMatchableOrders: []types.MatchableOrder{
				&constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,
				&constants.LongTermOrder_Bob_Num0_Id1_Clob0_Buy45_Price10_GTBT10,
			},
			collateralizationCheck: map[int]testutil_memclob.CollateralizationCheck{
				0: {
					CollatCheck: map[satypes.SubaccountId][]types.PendingOpenOrder{
						constants.Alice_Num0: {
							{
								RemainingQuantums: constants.BaseQuantums_5,
								IsBuy:             false,
								Subticks:          30,
								ClobPairId:        0,
							},
						},
						constants.Bob_Num0: {
							{
								RemainingQuantums: constants.BaseQuantums_5,
								IsBuy:             true,
								Subticks:          30,
								ClobPairId:        0,
							},
						},
					},
					Result: map[satypes.SubaccountId]satypes.UpdateResult{
						constants.Alice_Num0: satypes.Success,
						constants.Bob_Num0:   satypes.Success,
					},
				},
			},

			order: constants.Order_Alice_Num0_Id0_Clob0_Sell5_Price10_GTB20,

			expectedFilledSize:  constants.BaseQuantums_5,
			expectedOrderStatus: types.Success,
			expectedOperations: []types.Operation{
				clobtest.NewOrderPlacementOperation(
					constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,
				),
				clobtest.NewOrderPlacementOperation(
					constants.LongTermOrder_Bob_Num0_Id1_Clob0_Buy45_Price10_GTBT10,
				),
				clobtest.NewOrderPlacementOperation(
					constants.Order_Alice_Num0_Id0_Clob0_Sell5_Price10_GTB20,
				),
				clobtest.NewMatchOperation(
					&constants.Order_Alice_Num0_Id0_Clob0_Sell5_Price10_GTB20,
					[]types.MakerFill{
						{
							MakerOrderId: constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10.OrderId,
							FillAmount:   constants.FiveQuantumsSerializableInt,
						},
					},
				),
			},
			expectedInternalOperations: []types.InternalOperation{
				types.NewPreexistingStatefulOrderPlacementInternalOperation(
					constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,
				),
				types.NewShortTermOrderPlacementInternalOperation(
					constants.Order_Alice_Num0_Id0_Clob0_Sell5_Price10_GTB20,
				),
				types.NewMatchOrdersInternalOperation(
					constants.Order_Alice_Num0_Id0_Clob0_Sell5_Price10_GTB20,
					[]types.MakerFill{
						{
							MakerOrderId: constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10.OrderId,
							FillAmount:   constants.FiveQuantumsSerializableInt,
						},
					},
				),
			},
			expectedRemainingBids: []OrderWithRemainingSize{
				{
					Order:         constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,
					RemainingSize: constants.BaseQuantums_20,
				},
				{
					Order:         constants.LongTermOrder_Bob_Num0_Id1_Clob0_Buy45_Price10_GTBT10,
					RemainingSize: constants.BaseQuantums_45,
				},
			},
			expectedRemainingAsks: []OrderWithRemainingSize{},
		},
		`A Long-Term sell order can partially match with a Long-Term buy order, fail collateralization
			checks while matching, and all existing matches are considered valid`: {
			placedMatchableOrders: []types.MatchableOrder{
				&constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,
				&constants.LongTermOrder_Bob_Num0_Id1_Clob0_Buy45_Price10_GTBT10,
			},
			collateralizationCheck: map[int]testutil_memclob.CollateralizationCheck{
				0: {
					CollatCheck: map[satypes.SubaccountId][]types.PendingOpenOrder{
						constants.Alice_Num0: {
							{
								RemainingQuantums: constants.BaseQuantums_25,
								IsBuy:             false,
								Subticks:          30,
								ClobPairId:        0,
							},
						},
						constants.Bob_Num0: {
							{
								RemainingQuantums: constants.BaseQuantums_25,
								IsBuy:             true,
								Subticks:          30,
								ClobPairId:        0,
							},
						},
					},
					Result: map[satypes.SubaccountId]satypes.UpdateResult{
						constants.Alice_Num0: satypes.Success,
						constants.Bob_Num0:   satypes.Success,
					},
				},
				1: {
					CollatCheck: map[satypes.SubaccountId][]types.PendingOpenOrder{
						constants.Alice_Num0: {
							{
								RemainingQuantums: constants.BaseQuantums_40,
								IsBuy:             false,
								Subticks:          10,
								ClobPairId:        0,
							},
						},
						constants.Bob_Num0: {
							{
								RemainingQuantums: constants.BaseQuantums_40,
								IsBuy:             true,
								Subticks:          10,
								ClobPairId:        0,
							},
						},
					},
					Result: map[satypes.SubaccountId]satypes.UpdateResult{
						constants.Alice_Num0: satypes.StillUndercollateralized,
						constants.Bob_Num0:   satypes.NewlyUndercollateralized,
					},
				},
			},

			order: constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25,

			expectedFilledSize:  constants.BaseQuantums_25,
			expectedOrderStatus: types.Undercollateralized,
			expectedOperations: []types.Operation{
				clobtest.NewOrderPlacementOperation(
					constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,
				),
				clobtest.NewOrderPlacementOperation(
					constants.LongTermOrder_Bob_Num0_Id1_Clob0_Buy45_Price10_GTBT10,
				),
				clobtest.NewOrderPlacementOperation(
					constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25,
				),
				clobtest.NewMatchOperation(
					&constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25,
					[]types.MakerFill{
						{
							MakerOrderId: constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10.OrderId,
							FillAmount:   constants.TwentyFiveQuantumsSerializableInt,
						},
					},
				),
			},
			expectedInternalOperations: []types.InternalOperation{
				types.NewOrderRemovalInternalOperation(
					constants.LongTermOrder_Bob_Num0_Id1_Clob0_Buy45_Price10_GTBT10.OrderId,
					types.OrderRemoval_REMOVAL_REASON_UNDERCOLLATERALIZED,
				),
				types.NewPreexistingStatefulOrderPlacementInternalOperation(
					constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,
				),
				types.NewPreexistingStatefulOrderPlacementInternalOperation(
					constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25,
				),
				types.NewMatchOrdersInternalOperation(
					constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25,
					[]types.MakerFill{
						{
							MakerOrderId: constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10.OrderId,
							FillAmount:   constants.TwentyFiveQuantumsSerializableInt,
						},
					},
				),
				types.NewOrderRemovalInternalOperation(
					constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25.OrderId,
					types.OrderRemoval_REMOVAL_REASON_UNDERCOLLATERALIZED,
				),
			},
			expectedRemainingBids: []OrderWithRemainingSize{},
			expectedRemainingAsks: []OrderWithRemainingSize{},
		},
		`A Long-Term sell order can partially match with a Long-Term buy order, fail collateralization
			checks when adding to orderbook, and all existing matches are considered valid`: {
			placedMatchableOrders: []types.MatchableOrder{
				&constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,
			},
			collateralizationCheck: map[int]testutil_memclob.CollateralizationCheck{
				0: {
					CollatCheck: map[satypes.SubaccountId][]types.PendingOpenOrder{
						constants.Alice_Num0: {
							{
								RemainingQuantums: constants.BaseQuantums_25,
								IsBuy:             false,
								Subticks:          30,
								ClobPairId:        0,
							},
						},
						constants.Bob_Num0: {
							{
								RemainingQuantums: constants.BaseQuantums_25,
								IsBuy:             true,
								Subticks:          30,
								ClobPairId:        0,
							},
						},
					},
					Result: map[satypes.SubaccountId]satypes.UpdateResult{
						constants.Alice_Num0: satypes.Success,
						constants.Bob_Num0:   satypes.Success,
					},
				},
				1: {
					CollatCheck: map[satypes.SubaccountId][]types.PendingOpenOrder{
						constants.Alice_Num0: {
							{
								RemainingQuantums: constants.BaseQuantums_40,
								IsBuy:             false,
								Subticks:          10,
								ClobPairId:        0,
							},
						},
					},
					Result: map[satypes.SubaccountId]satypes.UpdateResult{
						constants.Alice_Num0: satypes.NewlyUndercollateralized,
					},
				},
			},

			order: constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25,

			expectedFilledSize:  constants.BaseQuantums_25,
			expectedOrderStatus: types.Undercollateralized,
			expectedOperations: []types.Operation{
				clobtest.NewOrderPlacementOperation(
					constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,
				),
				clobtest.NewOrderPlacementOperation(
					constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25,
				),
				clobtest.NewMatchOperation(
					&constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25,
					[]types.MakerFill{
						{
							MakerOrderId: constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10.OrderId,
							FillAmount:   constants.TwentyFiveQuantumsSerializableInt,
						},
					},
				),
			},
			expectedInternalOperations: []types.InternalOperation{
				types.NewPreexistingStatefulOrderPlacementInternalOperation(
					constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,
				),
				types.NewPreexistingStatefulOrderPlacementInternalOperation(
					constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25,
				),
				types.NewMatchOrdersInternalOperation(
					constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25,
					[]types.MakerFill{
						{
							MakerOrderId: constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10.OrderId,
							FillAmount:   constants.TwentyFiveQuantumsSerializableInt,
						},
					},
				),
				types.NewOrderRemovalInternalOperation(
					constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25.OrderId,
					types.OrderRemoval_REMOVAL_REASON_UNDERCOLLATERALIZED,
				),
			},
			expectedRemainingBids: []OrderWithRemainingSize{},
			expectedRemainingAsks: []OrderWithRemainingSize{},
		},
		`A Long-Term post-only sell order can partially match with a Long-Term buy order,
				all existing matches are reverted and it's not added to pendingStatefulOrders`: {
			placedMatchableOrders: []types.MatchableOrder{
				&constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,
			},
			collateralizationCheck: map[int]testutil_memclob.CollateralizationCheck{
				0: {
					CollatCheck: map[satypes.SubaccountId][]types.PendingOpenOrder{
						constants.Alice_Num0: {
							{
								RemainingQuantums: constants.BaseQuantums_25,
								IsBuy:             false,
								Subticks:          30,
								ClobPairId:        0,
							},
						},
						constants.Bob_Num0: {
							{
								RemainingQuantums: constants.BaseQuantums_25,
								IsBuy:             true,
								Subticks:          30,
								ClobPairId:        0,
							},
						},
					},
					Result: map[satypes.SubaccountId]satypes.UpdateResult{
						constants.Alice_Num0: satypes.Success,
						constants.Bob_Num0:   satypes.Success,
					},
				},
			},

			order: constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25_PO,

			expectedFilledSize:  constants.BaseQuantums_0,
			expectedOrderStatus: types.Success,
			expectedOperations: []types.Operation{
				clobtest.NewOrderPlacementOperation(
					constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,
				),
			},
			expectedInternalOperations: []types.InternalOperation{
				types.NewOrderRemovalInternalOperation(
					constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25_PO.OrderId,
					types.OrderRemoval_REMOVAL_REASON_POST_ONLY_WOULD_CROSS_MAKER_ORDER,
				),
			},
			expectedRemainingBids: []OrderWithRemainingSize{
				{
					Order:         constants.LongTermOrder_Bob_Num0_Id0_Clob0_Buy25_Price30_GTBT10,
					RemainingSize: constants.BaseQuantums_25,
				},
			},
			expectedRemainingAsks: []OrderWithRemainingSize{},
			expectedErr:           types.ErrPostOnlyWouldCrossMakerOrder,
		},
		`A Long-term buy order can self-match against a Long-term sell order from the same subaccount,
			causing the maker order to be removed`: {
			placedMatchableOrders: []types.MatchableOrder{
				&constants.LongTermOrder_Alice_Num0_Id0_Clob0_Buy5_Price10_GTBT15,
			},
			collateralizationCheck: map[int]testutil_memclob.CollateralizationCheck{
				0: {
					CollatCheck: map[satypes.SubaccountId][]types.PendingOpenOrder{
						constants.Alice_Num0: {
							{
								RemainingQuantums: constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25.GetBaseQuantums(),
								IsBuy:             constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25.IsBuy(),
								Subticks:          constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25.GetOrderSubticks(),
								ClobPairId:        constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25.GetClobPairId(),
							},
						},
					},
					Result: map[satypes.SubaccountId]satypes.UpdateResult{
						constants.Alice_Num0: satypes.Success,
					},
				},
			},

			order: constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25,

			expectedFilledSize:  constants.BaseQuantums_0,
			expectedOrderStatus: types.Success,
			expectedOperations:  []types.Operation{},
			expectedInternalOperations: []types.InternalOperation{
				types.NewOrderRemovalInternalOperation(
					constants.LongTermOrder_Alice_Num0_Id0_Clob0_Buy5_Price10_GTBT15.OrderId,
					types.OrderRemoval_REMOVAL_REASON_INVALID_SELF_TRADE,
				),
			},
			expectedRemainingBids: []OrderWithRemainingSize{},
			expectedRemainingAsks: []OrderWithRemainingSize{
				{
					Order:         constants.LongTermOrder_Alice_Num0_Id2_Clob0_Sell65_Price10_GTBT25,
					RemainingSize: constants.BaseQuantums_65,
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// Setup memclob state and test expectations.
			memclob, fakeMemClobKeeper, expectedNumCollateralizationChecks, numCollateralChecks := simplePlaceOrderTestSetup(
				t,
				ctx,
				tc.placedMatchableOrders,
				tc.collateralizationCheck,
				constants.GetStatePosition_ZeroPositionSize,
				&tc.order,
			)

			// Run the test case and verify expectations.
			placeOrderAndVerifyExpectationsOperations(
				t,
				ctx,
				memclob,
				tc.order,
				numCollateralChecks,
				tc.expectedFilledSize,
				tc.expectedFilledSize,
				tc.expectedOrderStatus,
				tc.expectedErr,
				expectedNumCollateralizationChecks,
				tc.expectedRemainingBids,
				tc.expectedRemainingAsks,
				tc.expectedOperations,
				tc.expectedInternalOperations,
				fakeMemClobKeeper,
			)

			// TODO(DEC-1296): Verify the correct offchain update messages were returned for Long-Term orders.
		})
	}
}

func TestPlaceOrder_PreexistingStatefulOrder(t *testing.T) {
	// Setup memclob state and test expectations.
	ctx, _, _ := sdktest.NewSdkContextWithMultistore()
	ctx = ctx.WithIsCheckTx(true)
	longTermOrder := constants.LongTermOrder_Alice_Num0_Id0_Clob0_Buy5_Price10_GTBT15
	collateralizationCheck := map[int]testutil_memclob.CollateralizationCheck{
		0: {
			CollatCheck: map[satypes.SubaccountId][]types.PendingOpenOrder{
				constants.Alice_Num0: {
					{
						RemainingQuantums: constants.BaseQuantums_5,
						IsBuy:             true,
						Subticks:          10,
						ClobPairId:        0,
					},
				},
			},
			Result: map[satypes.SubaccountId]satypes.UpdateResult{
				constants.Alice_Num0: satypes.Success,
			},
		},
	}
	memclob, fakeMemClobKeeper, expectedNumCollateralizationChecks, numCollateralChecks := simplePlaceOrderTestSetup(
		t,
		ctx,
		[]types.MatchableOrder{},
		collateralizationCheck,
		constants.GetStatePosition_ZeroPositionSize,
		&longTermOrder,
	)

	fakeMemClobKeeper.SetLongTermOrderPlacement(ctx, longTermOrder, uint32(5))

	// Run the test case and verify expectations.
	placeOrderAndVerifyExpectations(
		t,
		ctx,
		memclob,
		longTermOrder,
		numCollateralChecks,
		constants.BaseQuantums_0,
		constants.BaseQuantums_0,
		types.Success,
		nil,
		expectedNumCollateralizationChecks,
		[]OrderWithRemainingSize{
			{
				Order:         longTermOrder,
				RemainingSize: constants.BaseQuantums_5,
			},
		},
		[]OrderWithRemainingSize{},
		[]expectedMatch{},
		fakeMemClobKeeper,
	)
}
