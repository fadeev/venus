package keeper

import (
	"github.com/fadeev/venus/x/venus/types"
)

var _ types.QueryServer = Keeper{}
