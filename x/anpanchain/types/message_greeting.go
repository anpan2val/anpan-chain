package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGreeting = "greeting"

var _ sdk.Msg = &MsgGreeting{}

func NewMsgGreeting(creator string, fromAddress string, toAddress string) *MsgGreeting {
	return &MsgGreeting{
		Creator:     creator,
		FromAddress: fromAddress,
		ToAddress:   toAddress,
	}
}

func (msg *MsgGreeting) Route() string {
	return RouterKey
}

func (msg *MsgGreeting) Type() string {
	return TypeMsgGreeting
}

func (msg *MsgGreeting) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgGreeting) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgGreeting) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
