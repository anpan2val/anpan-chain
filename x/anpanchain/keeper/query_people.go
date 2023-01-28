package keeper

import (
	"context"

	"github.com/anpan2val/anpan-chain/x/anpanchain/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PeopleAll(goCtx context.Context, req *types.QueryAllPeopleRequest) (*types.QueryAllPeopleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var peoples []types.People
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	peopleStore := prefix.NewStore(store, types.KeyPrefix(types.PeopleKey))

	pageRes, err := query.Paginate(peopleStore, req.Pagination, func(key []byte, value []byte) error {
		var people types.People
		if err := k.cdc.Unmarshal(value, &people); err != nil {
			return err
		}

		peoples = append(peoples, people)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPeopleResponse{People: peoples, Pagination: pageRes}, nil
}

func (k Keeper) People(goCtx context.Context, req *types.QueryGetPeopleRequest) (*types.QueryGetPeopleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	people, found := k.GetPeople(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetPeopleResponse{People: people}, nil
}
