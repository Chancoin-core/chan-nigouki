package keeper_test

import (
	"context"
	"testing"

	keepertest "chancoin/testutil/keeper"
	"chancoin/x/chancoin/keeper"
	"chancoin/x/chancoin/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ChancoinKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
