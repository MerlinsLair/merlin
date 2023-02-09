package v6

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/merlinslair/merlin/app/keepers"
)

// RunForkLogic executes height-gated on-chain fork logic for the Merlin v6
// upgrade.
//
// NOTE: All the height gated fork logic is actually in the Merlin ibc-go fork.
// See: https://github.com/osmosis-labs/ibc-go/releases/tag/v2.0.2-mer
func RunForkLogic(ctx sdk.Context, _ *keepers.AppKeepers) {
	ctx.Logger().Info("Applying emergency hard fork for v6, allows IBC to create new channels.")
}
