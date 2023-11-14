package vutest

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"vutest/testutil/sample"
	vutestsimulation "vutest/x/vutest/simulation"
	"vutest/x/vutest/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = vutestsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateBaseDenom = "op_weight_msg_base_denom"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateBaseDenom int = 100

	opWeightMsgUpdateBaseDenom = "op_weight_msg_base_denom"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateBaseDenom int = 100

	opWeightMsgDeleteBaseDenom = "op_weight_msg_base_denom"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteBaseDenom int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	vutestGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&vutestGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateBaseDenom int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateBaseDenom, &weightMsgCreateBaseDenom, nil,
		func(_ *rand.Rand) {
			weightMsgCreateBaseDenom = defaultWeightMsgCreateBaseDenom
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateBaseDenom,
		vutestsimulation.SimulateMsgCreateBaseDenom(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateBaseDenom int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateBaseDenom, &weightMsgUpdateBaseDenom, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateBaseDenom = defaultWeightMsgUpdateBaseDenom
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateBaseDenom,
		vutestsimulation.SimulateMsgUpdateBaseDenom(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteBaseDenom int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteBaseDenom, &weightMsgDeleteBaseDenom, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteBaseDenom = defaultWeightMsgDeleteBaseDenom
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteBaseDenom,
		vutestsimulation.SimulateMsgDeleteBaseDenom(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateBaseDenom,
			defaultWeightMsgCreateBaseDenom,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				vutestsimulation.SimulateMsgCreateBaseDenom(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateBaseDenom,
			defaultWeightMsgUpdateBaseDenom,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				vutestsimulation.SimulateMsgUpdateBaseDenom(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteBaseDenom,
			defaultWeightMsgDeleteBaseDenom,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				vutestsimulation.SimulateMsgDeleteBaseDenom(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
