package anpanchain

import (
	"math/rand"

	"github.com/anpan2val/anpan-chain/testutil/sample"
	anpanchainsimulation "github.com/anpan2val/anpan-chain/x/anpanchain/simulation"
	"github.com/anpan2val/anpan-chain/x/anpanchain/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = anpanchainsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgGreeting = "op_weight_msg_greeting"
	// TODO: Determine the simulation weight value
	defaultWeightMsgGreeting int = 100

	opWeightMsgCreatePeople = "op_weight_msg_people"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePeople int = 100

	opWeightMsgUpdatePeople = "op_weight_msg_people"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePeople int = 100

	opWeightMsgDeletePeople = "op_weight_msg_people"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeletePeople int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	anpanchainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PeopleList: []types.People{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		PeopleCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&anpanchainGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgGreeting int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgGreeting, &weightMsgGreeting, nil,
		func(_ *rand.Rand) {
			weightMsgGreeting = defaultWeightMsgGreeting
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgGreeting,
		anpanchainsimulation.SimulateMsgGreeting(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreatePeople int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreatePeople, &weightMsgCreatePeople, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePeople = defaultWeightMsgCreatePeople
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePeople,
		anpanchainsimulation.SimulateMsgCreatePeople(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatePeople int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdatePeople, &weightMsgUpdatePeople, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePeople = defaultWeightMsgUpdatePeople
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePeople,
		anpanchainsimulation.SimulateMsgUpdatePeople(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeletePeople int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeletePeople, &weightMsgDeletePeople, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePeople = defaultWeightMsgDeletePeople
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePeople,
		anpanchainsimulation.SimulateMsgDeletePeople(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
