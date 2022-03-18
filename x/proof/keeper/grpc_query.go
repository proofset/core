package keeper

import (
	"github.com/proofset/core/x/proof/types"
)

var _ types.QueryServer = Keeper{}
