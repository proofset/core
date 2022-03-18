package proof_test

import (
	"testing"

	keepertest "github.com/proofset/core/testutil/keeper"
	"github.com/proofset/core/testutil/nullify"
	"github.com/proofset/core/x/proof"
	"github.com/proofset/core/x/proof/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ProofKeeper(t)
	proof.InitGenesis(ctx, *k, genesisState)
	got := proof.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
