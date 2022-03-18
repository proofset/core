package keeper_test

import (
	"testing"

	testkeeper "github.com/proofset/core/testutil/keeper"
	"github.com/proofset/core/x/proof/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ProofKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
