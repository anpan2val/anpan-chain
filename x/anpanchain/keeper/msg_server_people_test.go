package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/anpan2val/anpan-chain/x/anpanchain/types"
)

func TestPeopleMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreatePeople(ctx, &types.MsgCreatePeople{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestPeopleMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdatePeople
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdatePeople{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdatePeople{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdatePeople{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreatePeople(ctx, &types.MsgCreatePeople{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdatePeople(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestPeopleMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeletePeople
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeletePeople{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeletePeople{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeletePeople{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreatePeople(ctx, &types.MsgCreatePeople{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeletePeople(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
