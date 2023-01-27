package keeper

import (
	"github.com/anpan2val/anpan-chain/x/anpanchain/types"
)

var _ types.QueryServer = Keeper{}
