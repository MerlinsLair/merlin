package v15

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	gammkeeper "github.com/merlinslair/merlin/x/gamm/keeper"
	poolmanagerkeeper "github.com/merlinslair/merlin/x/poolmanager"
)

func MigrateNextPoolId(ctx sdk.Context, gammKeeper *gammkeeper.Keeper, poolmanagerKeeper *poolmanagerkeeper.Keeper) {
	migrateNextPoolId(ctx, gammKeeper, poolmanagerKeeper)
}

func RegisterMerIonMetadata(ctx sdk.Context, bankKeeper bankkeeper.Keeper) {
	registerMerIonMetadata(ctx, bankKeeper)
}
