package device

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/iventou/chainvite/testutil/sample"
	devicesimulation "github.com/iventou/chainvite/x/device/simulation"
	"github.com/iventou/chainvite/x/device/types"
)

// avoid unused import issue
var (
	_ = devicesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateDevice = "op_weight_msg_device"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDevice int = 100

	opWeightMsgUpdateDevice = "op_weight_msg_device"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDevice int = 100

	opWeightMsgDeleteDevice = "op_weight_msg_device"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteDevice int = 100

	opWeightMsgRegisterDevice = "op_weight_msg_register_device"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterDevice int = 100

	opWeightMsgUpdateDeviceStatus = "op_weight_msg_update_device_status"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDeviceStatus int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	deviceGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		DeviceList: []types.Device{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&deviceGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateDevice int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateDevice, &weightMsgCreateDevice, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDevice = defaultWeightMsgCreateDevice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDevice,
		devicesimulation.SimulateMsgCreateDevice(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateDevice int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateDevice, &weightMsgUpdateDevice, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDevice = defaultWeightMsgUpdateDevice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDevice,
		devicesimulation.SimulateMsgUpdateDevice(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteDevice int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteDevice, &weightMsgDeleteDevice, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteDevice = defaultWeightMsgDeleteDevice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteDevice,
		devicesimulation.SimulateMsgDeleteDevice(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRegisterDevice int
	simState.AppParams.GetOrGenerate(opWeightMsgRegisterDevice, &weightMsgRegisterDevice, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterDevice = defaultWeightMsgRegisterDevice
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterDevice,
		devicesimulation.SimulateMsgRegisterDevice(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateDeviceStatus int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateDeviceStatus, &weightMsgUpdateDeviceStatus, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDeviceStatus = defaultWeightMsgUpdateDeviceStatus
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDeviceStatus,
		devicesimulation.SimulateMsgUpdateDeviceStatus(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateDevice,
			defaultWeightMsgCreateDevice,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				devicesimulation.SimulateMsgCreateDevice(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateDevice,
			defaultWeightMsgUpdateDevice,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				devicesimulation.SimulateMsgUpdateDevice(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteDevice,
			defaultWeightMsgDeleteDevice,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				devicesimulation.SimulateMsgDeleteDevice(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRegisterDevice,
			defaultWeightMsgRegisterDevice,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				devicesimulation.SimulateMsgRegisterDevice(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateDeviceStatus,
			defaultWeightMsgUpdateDeviceStatus,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				devicesimulation.SimulateMsgUpdateDeviceStatus(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
