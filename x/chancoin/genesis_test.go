package chancoin_test

import (
	"testing"

	keepertest "chancoin/testutil/keeper"
	"chancoin/testutil/nullify"
	"chancoin/x/chancoin"
	"chancoin/x/chancoin/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ChancoinKeeper(t)
	chancoin.InitGenesis(ctx, *k, genesisState)
	got := chancoin.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
