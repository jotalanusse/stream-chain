package client_test

import (
	"context"
	"testing"

	"cosmossdk.io/log"

	"github.com/StreamFinance-Protocol/stream-chain/protocol/daemons/deleveraging/api"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/daemons/deleveraging/client"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/daemons/flags"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/dtypes"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/mocks"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/constants"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/grpc"
	blocktimetypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/blocktime/types"
	clobtypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/clob/types"
	satypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/subaccounts/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRunDeleveragingDaemonTaskLoop(t *testing.T) {
	tests := map[string]struct {
		// mocks
		setupMocks func(ctx context.Context, mck *mocks.QueryClient)

		// expectations
		expectedDeleveragingSubaccountIds []satypes.SubaccountId
		expectedError                     error
	}{
		"Can get liquidatable subaccount with short position": {
			setupMocks: func(ctx context.Context, mck *mocks.QueryClient) {
				// Block height.
				res := &blocktimetypes.QueryPreviousBlockInfoResponse{
					Info: &blocktimetypes.BlockInfo{
						Height:    uint32(50),
						Timestamp: constants.TimeTen,
					},
				}
				mck.On("PreviousBlockInfo", mock.Anything, mock.Anything).Return(res, nil)

				// Subaccount.
				res2 := &satypes.QuerySubaccountAllResponse{
					Subaccount: []satypes.Subaccount{
						constants.Carl_Num0_1BTC_Short_54999USD,
					},
				}
				mck.On("SubaccountAll", mock.Anything, mock.Anything).Return(res2, nil)

				// Sends liquidatable subaccount ids to the server.
				req := &api.DeleveragingSubaccountsRequest{
					SubaccountOpenPositionInfo: []clobtypes.SubaccountOpenPositionInfo{
						{
							PerpetualId:                 0,
							SubaccountsWithLongPosition: []satypes.SubaccountId{},
							SubaccountsWithShortPosition: []satypes.SubaccountId{
								constants.Carl_Num0,
							},
						},
					},
				}
				response3 := &api.DeleveragingSubaccountsResponse{}
				mck.On("DeleverageSubaccounts", ctx, req).Return(response3, nil)
			},
		},
		"Can get liquidatable subaccount with long position": {
			setupMocks: func(ctx context.Context, mck *mocks.QueryClient) {
				// Block height.
				res := &blocktimetypes.QueryPreviousBlockInfoResponse{
					Info: &blocktimetypes.BlockInfo{
						Height:    uint32(50),
						Timestamp: constants.TimeTen,
					},
				}
				mck.On("PreviousBlockInfo", mock.Anything, mock.Anything).Return(res, nil)

				// Subaccount.
				res2 := &satypes.QuerySubaccountAllResponse{
					Subaccount: []satypes.Subaccount{
						constants.Dave_Num0_1BTC_Long_45001USD_Short,
					},
				}
				mck.On("SubaccountAll", mock.Anything, mock.Anything).Return(res2, nil)

				// Sends liquidatable subaccount ids to the server.
				req := &api.DeleveragingSubaccountsRequest{
					SubaccountOpenPositionInfo: []clobtypes.SubaccountOpenPositionInfo{
						{
							PerpetualId: 0,
							SubaccountsWithLongPosition: []satypes.SubaccountId{
								constants.Dave_Num0,
							},
							SubaccountsWithShortPosition: []satypes.SubaccountId{},
						},
					},
				}
				response3 := &api.DeleveragingSubaccountsResponse{}
				mck.On("DeleverageSubaccounts", ctx, req).Return(response3, nil)
			},
		},
		"Skip well collateralized subaccounts": {
			setupMocks: func(ctx context.Context, mck *mocks.QueryClient) {
				// Block height.
				res := &blocktimetypes.QueryPreviousBlockInfoResponse{
					Info: &blocktimetypes.BlockInfo{
						Height:    uint32(50),
						Timestamp: constants.TimeTen,
					},
				}
				mck.On("PreviousBlockInfo", mock.Anything, mock.Anything).Return(res, nil)

				// Subaccount.
				res2 := &satypes.QuerySubaccountAllResponse{
					Subaccount: []satypes.Subaccount{
						constants.Carl_Num0_1BTC_Short_55000USD,
						constants.Dave_Num0_1BTC_Long_45000USD_Short,
					},
				}
				mck.On("SubaccountAll", mock.Anything, mock.Anything).Return(res2, nil)

				// Sends liquidatable subaccount ids to the server.
				req := &api.DeleveragingSubaccountsRequest{
					SubaccountOpenPositionInfo: []clobtypes.SubaccountOpenPositionInfo{
						{
							PerpetualId: 0,
							SubaccountsWithLongPosition: []satypes.SubaccountId{
								constants.Dave_Num0,
							},
							SubaccountsWithShortPosition: []satypes.SubaccountId{
								constants.Carl_Num0,
							},
						},
					},
				}
				response3 := &api.DeleveragingSubaccountsResponse{}
				mck.On("DeleverageSubaccounts", ctx, req).Return(response3, nil)
			},
		},
		"Skip subaccounts with no open positions": {
			setupMocks: func(ctx context.Context, mck *mocks.QueryClient) {
				// Block height.
				res := &blocktimetypes.QueryPreviousBlockInfoResponse{
					Info: &blocktimetypes.BlockInfo{
						Height:    uint32(50),
						Timestamp: constants.TimeTen,
					},
				}
				mck.On("PreviousBlockInfo", mock.Anything, mock.Anything).Return(res, nil)

				// Subaccount.
				res2 := &satypes.QuerySubaccountAllResponse{
					Subaccount: []satypes.Subaccount{
						constants.Alice_Num0_100_000USD,
					},
				}
				mck.On("SubaccountAll", mock.Anything, mock.Anything).Return(res2, nil)

				// Sends liquidatable subaccount ids to the server.
				req := &api.DeleveragingSubaccountsRequest{
					SubaccountOpenPositionInfo: []clobtypes.SubaccountOpenPositionInfo{},
				}
				response3 := &api.DeleveragingSubaccountsResponse{}
				mck.On("DeleverageSubaccounts", ctx, req).Return(response3, nil)
			},
		},
		"Can get subaccount that become undercollateralized with funding payments (short)": {
			setupMocks: func(ctx context.Context, mck *mocks.QueryClient) {
				// Block height.
				res := &blocktimetypes.QueryPreviousBlockInfoResponse{
					Info: &blocktimetypes.BlockInfo{
						Height:    uint32(50),
						Timestamp: constants.TimeTen,
					},
				}
				mck.On("PreviousBlockInfo", mock.Anything, mock.Anything).Return(res, nil)

				// Subaccount.
				res2 := &satypes.QuerySubaccountAllResponse{
					Subaccount: []satypes.Subaccount{
						// Without funding, Carl has a TNC of $5,000, MMR of $5,000, and is
						// well-collateralized.
						// However, funding index for Carl's position is 10,000 and perpetual's funding index
						// is 0. Index delta is -10,000, so Carl has to make a funding payment of $1 and
						// become under-collateralized.
						{
							Id: &constants.Carl_Num0,
							AssetPositions: []*satypes.AssetPosition{
								{
									AssetId:  0,
									Quantums: dtypes.NewInt(55_000_000_000), // $55,000
								},
							},
							PerpetualPositions: []*satypes.PerpetualPosition{
								{
									PerpetualId:  0,
									Quantums:     dtypes.NewInt(-100_000_000), // -1 BTC
									FundingIndex: dtypes.NewInt(10_000),
								},
							},
						},
					},
				}
				mck.On("SubaccountAll", mock.Anything, mock.Anything).Return(res2, nil)

				// Sends liquidatable subaccount ids to the server.
				req := &api.DeleveragingSubaccountsRequest{
					SubaccountOpenPositionInfo: []clobtypes.SubaccountOpenPositionInfo{
						{
							PerpetualId:                 0,
							SubaccountsWithLongPosition: []satypes.SubaccountId{},
							SubaccountsWithShortPosition: []satypes.SubaccountId{
								constants.Carl_Num0,
							},
						},
					},
				}
				response3 := &api.DeleveragingSubaccountsResponse{}
				mck.On("DeleverageSubaccounts", ctx, req).Return(response3, nil)
			},
		},
		"Can get subaccount that become liquidatable with funding payments (long)": {
			setupMocks: func(ctx context.Context, mck *mocks.QueryClient) {
				// Block height.
				res := &blocktimetypes.QueryPreviousBlockInfoResponse{
					Info: &blocktimetypes.BlockInfo{
						Height:    uint32(50),
						Timestamp: constants.TimeTen,
					},
				}
				mck.On("PreviousBlockInfo", mock.Anything, mock.Anything).Return(res, nil)

				// Subaccount.
				res2 := &satypes.QuerySubaccountAllResponse{
					Subaccount: []satypes.Subaccount{
						// Without funding, Dave has a TNC of $5,000, MMR of $5,000, and is
						// well-collateralized.
						// However, funding index for Dave's position is -10,000 and perpetual's funding index
						// is 0. Index delta is 10,000, so Dave has to make a funding payment of $1 and
						// become under-collateralized.
						{
							Id: &constants.Dave_Num0,
							AssetPositions: []*satypes.AssetPosition{
								{
									AssetId:  0,
									Quantums: dtypes.NewInt(-45_000_000_000), // -$45,000
								},
							},
							PerpetualPositions: []*satypes.PerpetualPosition{
								{
									PerpetualId:  0,
									Quantums:     dtypes.NewInt(100_000_000), // 1 BTC
									FundingIndex: dtypes.NewInt(-10_000),
								},
							},
						},
					},
				}
				mck.On("SubaccountAll", mock.Anything, mock.Anything).Return(res2, nil)

				// Sends liquidatable subaccount ids to the server.
				req := &api.DeleveragingSubaccountsRequest{
					SubaccountOpenPositionInfo: []clobtypes.SubaccountOpenPositionInfo{
						{
							PerpetualId: 0,
							SubaccountsWithLongPosition: []satypes.SubaccountId{
								constants.Dave_Num0,
							},
							SubaccountsWithShortPosition: []satypes.SubaccountId{},
						},
					},
				}
				response3 := &api.DeleveragingSubaccountsResponse{}
				mck.On("DeleverageSubaccounts", ctx, req).Return(response3, nil)
			},
		},
		"Skips subaccount that become well-collateralized with funding payments (short)": {
			setupMocks: func(ctx context.Context, mck *mocks.QueryClient) {
				// Block height.
				res := &blocktimetypes.QueryPreviousBlockInfoResponse{
					Info: &blocktimetypes.BlockInfo{
						Height:    uint32(50),
						Timestamp: constants.TimeTen,
					},
				}
				mck.On("PreviousBlockInfo", mock.Anything, mock.Anything).Return(res, nil)

				// Subaccount.
				res2 := &satypes.QuerySubaccountAllResponse{
					Subaccount: []satypes.Subaccount{
						// Without funding, Carl has a TNC of $4,999, MMR of $5,000, and is
						// under-collateralized.
						// However, funding index for Carl's position is -10,000 and perpetual's funding index
						// is 0. Index delta is 10,000, so Carl would receive a funding payment of $1 and
						// become well-collateralized.
						{
							Id: &constants.Carl_Num0,
							AssetPositions: []*satypes.AssetPosition{
								{
									AssetId:  0,
									Quantums: dtypes.NewInt(54_999_000_000), // $54,999
								},
							},
							PerpetualPositions: []*satypes.PerpetualPosition{
								{
									PerpetualId:  0,
									Quantums:     dtypes.NewInt(-100_000_000), // -1 BTC
									FundingIndex: dtypes.NewInt(-10_000),
								},
							},
						},
					},
				}
				mck.On("SubaccountAll", mock.Anything, mock.Anything).Return(res2, nil)

				// Sends liquidatable subaccount ids to the server.
				req := &api.DeleveragingSubaccountsRequest{
					SubaccountOpenPositionInfo: []clobtypes.SubaccountOpenPositionInfo{
						{
							PerpetualId:                 0,
							SubaccountsWithLongPosition: []satypes.SubaccountId{},
							SubaccountsWithShortPosition: []satypes.SubaccountId{
								constants.Carl_Num0,
							},
						},
					},
				}
				response3 := &api.DeleveragingSubaccountsResponse{}
				mck.On("DeleverageSubaccounts", ctx, req).Return(response3, nil)
			},
		},
		"Skips subaccount that become well-collateralized with funding payments (long)": {
			setupMocks: func(ctx context.Context, mck *mocks.QueryClient) {
				// Block height.
				res := &blocktimetypes.QueryPreviousBlockInfoResponse{
					Info: &blocktimetypes.BlockInfo{
						Height:    uint32(50),
						Timestamp: constants.TimeTen,
					},
				}
				mck.On("PreviousBlockInfo", mock.Anything, mock.Anything).Return(res, nil)

				// Subaccount.
				res2 := &satypes.QuerySubaccountAllResponse{
					Subaccount: []satypes.Subaccount{
						// Without funding, Dave has a TNC of $4,999, MMR of $5,000, and is
						// under-collateralized.
						// However, funding index for Dave's position is 10,000 and perpetual's funding index
						// is 0. Index delta is -10,000, so Dave would receive a funding payment of $1 and
						// become well-collateralized.
						{
							Id: &constants.Dave_Num0,
							AssetPositions: []*satypes.AssetPosition{
								{
									AssetId:  0,
									Quantums: dtypes.NewInt(-44_999_000_000), // -$44,999
								},
							},
							PerpetualPositions: []*satypes.PerpetualPosition{
								{
									PerpetualId:  0,
									Quantums:     dtypes.NewInt(100_000_000), // 1 BTC
									FundingIndex: dtypes.NewInt(10_000),
								},
							},
						},
					},
				}
				mck.On("SubaccountAll", mock.Anything, mock.Anything).Return(res2, nil)

				// Sends liquidatable subaccount ids to the server.
				req := &api.DeleveragingSubaccountsRequest{
					SubaccountOpenPositionInfo: []clobtypes.SubaccountOpenPositionInfo{
						{
							PerpetualId: 0,
							SubaccountsWithLongPosition: []satypes.SubaccountId{
								constants.Dave_Num0,
							},
							SubaccountsWithShortPosition: []satypes.SubaccountId{},
						},
					},
				}
				response3 := &api.DeleveragingSubaccountsResponse{}
				mck.On("DeleverageSubaccounts", ctx, req).Return(response3, nil)
			},
		},
		"Can get negative tnc subaccount with short position": {
			setupMocks: func(ctx context.Context, mck *mocks.QueryClient) {
				// Block height.
				res := &blocktimetypes.QueryPreviousBlockInfoResponse{
					Info: &blocktimetypes.BlockInfo{
						Height:    uint32(50),
						Timestamp: constants.TimeTen,
					},
				}
				mck.On("PreviousBlockInfo", mock.Anything, mock.Anything).Return(res, nil)

				// Subaccount.
				res2 := &satypes.QuerySubaccountAllResponse{
					Subaccount: []satypes.Subaccount{
						// Carl has TNC of -$1.
						constants.Carl_Num0_1BTC_Short_49999USD,
					},
				}
				mck.On("SubaccountAll", mock.Anything, mock.Anything).Return(res2, nil)

				// Sends liquidatable subaccount ids to the server.
				req := &api.DeleveragingSubaccountsRequest{
					SubaccountOpenPositionInfo: []clobtypes.SubaccountOpenPositionInfo{
						{
							PerpetualId:                 0,
							SubaccountsWithLongPosition: []satypes.SubaccountId{},
							SubaccountsWithShortPosition: []satypes.SubaccountId{
								constants.Carl_Num0,
							},
						},
					},
				}
				response3 := &api.DeleveragingSubaccountsResponse{}
				mck.On("DeleverageSubaccounts", ctx, req).Return(response3, nil)
			},
		},
		"Can get negative tnc subaccount with long position": {
			setupMocks: func(ctx context.Context, mck *mocks.QueryClient) {
				// Block height.
				res := &blocktimetypes.QueryPreviousBlockInfoResponse{
					Info: &blocktimetypes.BlockInfo{
						Height:    uint32(50),
						Timestamp: constants.TimeTen,
					},
				}
				mck.On("PreviousBlockInfo", mock.Anything, mock.Anything).Return(res, nil)

				// Subaccount.
				res2 := &satypes.QuerySubaccountAllResponse{
					Subaccount: []satypes.Subaccount{
						// Dave has TNC of -$1.
						constants.Dave_Num0_1BTC_Long_50001USD_Short,
					},
				}
				mck.On("SubaccountAll", mock.Anything, mock.Anything).Return(res2, nil)

				// Sends liquidatable subaccount ids to the server.
				req := &api.DeleveragingSubaccountsRequest{
					SubaccountOpenPositionInfo: []clobtypes.SubaccountOpenPositionInfo{
						{
							PerpetualId: 0,
							SubaccountsWithLongPosition: []satypes.SubaccountId{
								constants.Dave_Num0,
							},
							SubaccountsWithShortPosition: []satypes.SubaccountId{},
						},
					},
				}
				response3 := &api.DeleveragingSubaccountsResponse{}
				mck.On("DeleverageSubaccounts", ctx, req).Return(response3, nil)
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			queryClientMock := &mocks.QueryClient{}
			tc.setupMocks(grpc.Ctx, queryClientMock)
			s := client.SubTaskRunnerImpl{}

			c := client.NewClient(log.NewNopLogger())
			c.SubaccountQueryClient = queryClientMock
			c.DeleveragingServiceClient = queryClientMock
			c.BlocktimeQueryClient = queryClientMock

			err := s.RunDeleveragingDaemonTaskLoop(
				grpc.Ctx,
				c,
				flags.GetDefaultDaemonFlags().Deleveraging,
			)
			if tc.expectedError != nil {
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
				queryClientMock.AssertExpectations(t)
			}
		})
	}
}