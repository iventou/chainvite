package ticket

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/iventou/chainvite/api/chainvite/ticket"
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
					RpcMethod: "TicketAll",
					Use:       "list-ticket",
					Short:     "List all ticket",
				},
				{
					RpcMethod:      "Ticket",
					Use:            "show-ticket [id]",
					Short:          "Shows a ticket",
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
					RpcMethod:      "CreateTicket",
					Use:            "create-ticket [index] [ticketId] [eventId] [ticketType] [valid] [owner] [metadata]",
					Short:          "Create a new ticket",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "ticketId"}, {ProtoField: "eventId"}, {ProtoField: "ticketType"}, {ProtoField: "valid"}, {ProtoField: "owner"}, {ProtoField: "metadata"}},
				},
				{
					RpcMethod:      "UpdateTicket",
					Use:            "update-ticket [index] [ticketId] [eventId] [ticketType] [valid] [owner] [metadata]",
					Short:          "Update ticket",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "ticketId"}, {ProtoField: "eventId"}, {ProtoField: "ticketType"}, {ProtoField: "valid"}, {ProtoField: "owner"}, {ProtoField: "metadata"}},
				},
				{
					RpcMethod:      "DeleteTicket",
					Use:            "delete-ticket [index]",
					Short:          "Delete ticket",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "MintTicket",
					Use:            "mint-ticket [event-id] [ticket-type] [owner] [metadata]",
					Short:          "Send a mint-ticket tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "eventId"}, {ProtoField: "ticketType"}, {ProtoField: "owner"}, {ProtoField: "metadata"}},
				},
				{
					RpcMethod:      "ValidateTicket",
					Use:            "validate-ticket [ticket-id] [device-id]",
					Short:          "Send a validate-ticket tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "ticketId"}, {ProtoField: "deviceId"}},
				},
				{
					RpcMethod:      "TransferTicket",
					Use:            "transfer-ticket [ticket-id] [new-owner]",
					Short:          "Send a transfer-ticket tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "ticketId"}, {ProtoField: "newOwner"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
