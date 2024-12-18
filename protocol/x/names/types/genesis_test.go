package types_test

import (
	"testing"

	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/names/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := map[string]struct {
		genState    *types.GenesisState
		expectedErr error
	}{
		"default is valid": {
			genState: types.DefaultGenesis(),
		},
		"valid genesis state": {
			genState: &types.GenesisState{
				Names: []types.Name{
					{
						Id:   0,
						Name: types.NameJota.Name,
					},
				},
			},
		},
		"empty genesis state": {
			genState: &types.GenesisState{
				Names: []types.Name{},
			},
			expectedErr: types.ErrNoNameInGenesis,
		},
		"name[0] not jota": {
			genState: &types.GenesisState{
				Names: []types.Name{
					{
						Id:   0,
						Name: types.NameSolal.Name,
					},
				},
			},
			expectedErr: types.ErrJotaMustBeNameZero,
		},
		"duplicated name id": {
			genState: &types.GenesisState{
				Names: []types.Name{
					{
						Id:   0,
						Name: types.NameJota.Name,
					},
					{
						Id:   0,
						Name: types.NameSolal.Name,
					},
				},
			},
			expectedErr: types.ErrNameIdAlreadyExists,
		},
		"duplicated name": {
			genState: &types.GenesisState{
				Names: []types.Name{
					{
						Id:   0,
						Name: types.NameJota.Name,
					},
					{
						Id:   1,
						Name: types.NameJota.Name,
					},
				},
			},
			expectedErr: types.ErrNameNameAlreadyExists,
		},
		"gaps in name id": {
			genState: &types.GenesisState{
				Names: []types.Name{
					{
						Id:   0,
						Name: types.NameJota.Name,
					},
					{
						Id:   2,
						Name: types.NameSolal.Name,
					},
				},
			},
			expectedErr: types.ErrGapFoundInNameId,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.expectedErr == nil {
				require.NoError(t, err)
			} else {
				require.ErrorIs(t, err, tc.expectedErr)
			}
		})
	}
}
