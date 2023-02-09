package v15

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	packetforwardtypes "github.com/strangelove-ventures/packet-forward-middleware/v4/router/types"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/merlinslair/merlin/app/keepers"
	appParams "github.com/merlinslair/merlin/app/params"
	"github.com/merlinslair/merlin/app/upgrades"
	gammkeeper "github.com/merlinslair/merlin/x/gamm/keeper"
	"github.com/merlinslair/merlin/x/poolmanager"
	poolmanagertypes "github.com/merlinslair/merlin/x/poolmanager/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	bpm upgrades.BaseAppParamManager,
	keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		poolmanagerParams := poolmanagertypes.NewParams(keepers.GAMMKeeper.GetParams(ctx).PoolCreationFee)

		keepers.PoolManagerKeeper.SetParams(ctx, poolmanagerParams)
		keepers.PacketForwardKeeper.SetParams(ctx, packetforwardtypes.DefaultParams())

		// N.B: pool id in gamm is to be deprecated in the future
		// Instead,it is moved to poolmanager.
		migrateNextPoolId(ctx, keepers.GAMMKeeper, keepers.PoolManagerKeeper)

		//  N.B.: this is done to avoid initializing genesis for poolmanager module.
		// Otherwise, it would overwrite migrations with InitGenesis().
		// See RunMigrations() for details.
		fromVM[poolmanagertypes.ModuleName] = 0

		// Metadata for umer and uion were missing prior to this upgrade.
		// They are added in this upgrade.
		registerMerIonMetadata(ctx, keepers.BankKeeper)

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}

func migrateNextPoolId(ctx sdk.Context, gammKeeper *gammkeeper.Keeper, poolmanagerKeeper *poolmanager.Keeper) {
	// N.B: pool id in gamm is to be deprecated in the future
	// Instead,it is moved to poolmanager.
	// nolint: staticcheck
	nextPoolId := gammKeeper.GetNextPoolId(ctx)
	poolmanagerKeeper.SetNextPoolId(ctx, nextPoolId)

	for poolId := uint64(1); poolId < nextPoolId; poolId++ {
		poolType, err := gammKeeper.GetPoolType(ctx, poolId)
		if err != nil {
			panic(err)
		}

		poolmanagerKeeper.SetPoolRoute(ctx, poolId, poolType)
	}
}

func registerMerIonMetadata(ctx sdk.Context, bankKeeper bankkeeper.Keeper) {
	umerMetadata := banktypes.Metadata{
		Description: "The native token of Merlin",
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    appParams.BaseCoinUnit,
				Exponent: 0,
				Aliases:  nil,
			},
			{
				Denom:    appParams.HumanCoinUnit,
				Exponent: appParams.MerExponent,
				Aliases:  nil,
			},
		},
		Base:    appParams.BaseCoinUnit,
		Display: appParams.HumanCoinUnit,
	}

	uionMetadata := banktypes.Metadata{
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    "uion",
				Exponent: 0,
				Aliases:  nil,
			},
			{
				Denom:    "ion",
				Exponent: 6,
				Aliases:  nil,
			},
		},
		Base:    "uion",
		Display: "ion",
	}

	bankKeeper.SetDenomMetaData(ctx, umerMetadata)
	bankKeeper.SetDenomMetaData(ctx, uionMetadata)
}
