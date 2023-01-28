package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreatePeople = "create_people"
	TypeMsgUpdatePeople = "update_people"
	TypeMsgDeletePeople = "delete_people"
)

var _ sdk.Msg = &MsgCreatePeople{}

func NewMsgCreatePeople(creator string, address string, name string) *MsgCreatePeople {
	return &MsgCreatePeople{
		Creator: creator,
		Address: address,
		Name:    name,
	}
}

func (msg *MsgCreatePeople) Route() string {
	return RouterKey
}

func (msg *MsgCreatePeople) Type() string {
	return TypeMsgCreatePeople
}

func (msg *MsgCreatePeople) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePeople) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePeople) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePeople{}

func NewMsgUpdatePeople(creator string, id uint64, address string, name string) *MsgUpdatePeople {
	return &MsgUpdatePeople{
		Id:      id,
		Creator: creator,
		Address: address,
		Name:    name,
	}
}

func (msg *MsgUpdatePeople) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePeople) Type() string {
	return TypeMsgUpdatePeople
}

func (msg *MsgUpdatePeople) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePeople) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePeople) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePeople{}

func NewMsgDeletePeople(creator string, id uint64) *MsgDeletePeople {
	return &MsgDeletePeople{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeletePeople) Route() string {
	return RouterKey
}

func (msg *MsgDeletePeople) Type() string {
	return TypeMsgDeletePeople
}

func (msg *MsgDeletePeople) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeletePeople) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePeople) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
