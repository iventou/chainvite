package keeper

import (
	"github.com/iventou/chainvite/x/ticket/types"
)

var _ types.QueryServer = Keeper{}
