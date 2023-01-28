package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgGreeting{}, "anpanchain/Greeting", nil)
	cdc.RegisterConcrete(&MsgCreatePeople{}, "anpanchain/CreatePeople", nil)
	cdc.RegisterConcrete(&MsgUpdatePeople{}, "anpanchain/UpdatePeople", nil)
	cdc.RegisterConcrete(&MsgDeletePeople{}, "anpanchain/DeletePeople", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgGreeting{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePeople{},
		&MsgUpdatePeople{},
		&MsgDeletePeople{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
