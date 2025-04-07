package device

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/iventou/chainvite/api/chainvite/device"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "DeviceAll",
					Use:       "list-device",
					Short:     "List all device",
				},
				{
					RpcMethod:      "Device",
					Use:            "show-device [id]",
					Short:          "Shows a device",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateDevice",
					Use:            "create-device [index] [deviceId] [manufacturer] [location] [eventId] [status]",
					Short:          "Create a new device",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "deviceId"}, {ProtoField: "manufacturer"}, {ProtoField: "location"}, {ProtoField: "eventId"}, {ProtoField: "status"}},
				},
				{
					RpcMethod:      "UpdateDevice",
					Use:            "update-device [index] [deviceId] [manufacturer] [location] [eventId] [status]",
					Short:          "Update device",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "deviceId"}, {ProtoField: "manufacturer"}, {ProtoField: "location"}, {ProtoField: "eventId"}, {ProtoField: "status"}},
				},
				{
					RpcMethod:      "DeleteDevice",
					Use:            "delete-device [index]",
					Short:          "Delete device",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "RegisterDevice",
					Use:            "register-device [manufacturer] [location] [event-id]",
					Short:          "Send a register-device tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "manufacturer"}, {ProtoField: "location"}, {ProtoField: "eventId"}},
				},
				{
					RpcMethod:      "UpdateDeviceStatus",
					Use:            "update-device-status [device-id] [status]",
					Short:          "Send a update-device-status tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "deviceId"}, {ProtoField: "status"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
