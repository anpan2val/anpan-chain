package keeper

import (
	"context"

	"github.com/anpan2val/anpan-chain/x/anpanchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Greeting(goCtx context.Context, msg *types.MsgGreeting) (*types.MsgGreetingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgGreetingResponse{}, nil
}
