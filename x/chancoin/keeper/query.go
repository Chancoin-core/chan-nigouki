package keeper

import (
	"chancoin/x/chancoin/types"
)

var _ types.QueryServer = Keeper{}
