package keeper

import (
	"context"
	"fmt"

	"github.com/anpan2val/anpan-chain/x/anpanchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreatePeople(goCtx context.Context, msg *types.MsgCreatePeople) (*types.MsgCreatePeopleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var people = types.People{
		Creator: msg.Creator,
		Address: msg.Address,
		Name:    msg.Name,
	}

	id := k.AppendPeople(
		ctx,
		people,
	)

	return &types.MsgCreatePeopleResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdatePeople(goCtx context.Context, msg *types.MsgUpdatePeople) (*types.MsgUpdatePeopleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var people = types.People{
		Creator: msg.Creator,
		Id:      msg.Id,
		Address: msg.Address,
		Name:    msg.Name,
	}

	// Checks that the element exists
	val, found := k.GetPeople(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetPeople(ctx, people)

	return &types.MsgUpdatePeopleResponse{}, nil
}

func (k msgServer) DeletePeople(goCtx context.Context, msg *types.MsgDeletePeople) (*types.MsgDeletePeopleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetPeople(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePeople(ctx, msg.Id)

	return &types.MsgDeletePeopleResponse{}, nil
}
