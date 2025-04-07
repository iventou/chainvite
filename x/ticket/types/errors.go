package types

import "cosmossdk.io/errors"

// Ticket module error codes
var (
	ErrTicketNotFound         = errors.Register(ModuleName, 1, "ticket not found")
	ErrInvalidTicketValidity  = errors.Register(ModuleName, 2, "invalid ticket validity")
	ErrUnauthorizedAccess     = errors.Register(ModuleName, 3, "unauthorized access")
	ErrInvalidTicketOperation = errors.Register(ModuleName, 4, "invalid ticket operation")
)
