package keeper_test

import (
	"testing"

	testkeeper "chancoin/testutil/keeper"
	"chancoin/x/chancoin/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ChancoinKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
