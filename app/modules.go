package app

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/client"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	icq "github.com/strangelove-ventures/async-icq"

	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v4/modules/core"
	ibchost "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v4/modules/core/keeper"

	ibchookstypes "github.com/osmosis-labs/osmosis/x/ibc-hooks/types"

	ica "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"

	icqtypes "github.com/strangelove-ventures/async-icq/types"

	downtimemodule "github.com/merlinslair/merlin/x/downtime-detector/module"
	downtimetypes "github.com/merlinslair/merlin/x/downtime-detector/types"

	ibc_hooks "github.com/osmosis-labs/osmosis/x/ibc-hooks"

	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	appparams "github.com/merlinslair/merlin/app/params"
	_ "github.com/merlinslair/merlin/client/docs/statik"
	"github.com/merlinslair/merlin/simulation/simtypes"
	concentratedliquidity "github.com/merlinslair/merlin/x/concentrated-liquidity/clmodule"
	concentratedliquiditytypes "github.com/merlinslair/merlin/x/concentrated-liquidity/types"
	"github.com/merlinslair/merlin/x/epochs"
	epochstypes "github.com/merlinslair/merlin/x/epochs/types"
	"github.com/merlinslair/merlin/x/gamm"
	gammtypes "github.com/merlinslair/merlin/x/gamm/types"
	"github.com/merlinslair/merlin/x/incentives"
	incentivestypes "github.com/merlinslair/merlin/x/incentives/types"
	"github.com/merlinslair/merlin/x/lockup"
	lockuptypes "github.com/merlinslair/merlin/x/lockup/types"
	"github.com/merlinslair/merlin/x/mint"
	minttypes "github.com/merlinslair/merlin/x/mint/types"
	poolincentives "github.com/merlinslair/merlin/x/pool-incentives"
	poolincentivestypes "github.com/merlinslair/merlin/x/pool-incentives/types"
	poolmanager "github.com/merlinslair/merlin/x/poolmanager/module"
	poolmanagertypes "github.com/merlinslair/merlin/x/poolmanager/types"
	"github.com/merlinslair/merlin/x/protorev"
	protorevtypes "github.com/merlinslair/merlin/x/protorev/types"
	superfluid "github.com/merlinslair/merlin/x/superfluid"
	superfluidtypes "github.com/merlinslair/merlin/x/superfluid/types"
	"github.com/merlinslair/merlin/x/tokenfactory"
	tokenfactorytypes "github.com/merlinslair/merlin/x/tokenfactory/types"
	"github.com/merlinslair/merlin/x/twap/twapmodule"
	twaptypes "github.com/merlinslair/merlin/x/twap/types"
	"github.com/merlinslair/merlin/x/txfees"
	txfeestypes "github.com/merlinslair/merlin/x/txfees/types"
	valsetpreftypes "github.com/merlinslair/merlin/x/valset-pref/types"
	valsetprefmodule "github.com/merlinslair/merlin/x/valset-pref/valpref-module"
	"github.com/osmosis-labs/osmosis/osmoutils/partialord"
)

// moduleAccountPermissions defines module account permissions
// TODO: Having to input nil's here is unacceptable, we need a way to automatically derive this.
var moduleAccountPermissions = map[string][]string{
	authtypes.FeeCollectorName:               nil,
	distrtypes.ModuleName:                    nil,
	ibchookstypes.ModuleName:                 nil,
	icatypes.ModuleName:                      nil,
	icqtypes.ModuleName:                      nil,
	minttypes.ModuleName:                     {authtypes.Minter, authtypes.Burner},
	minttypes.DeveloperVestingModuleAcctName: nil,
	stakingtypes.BondedPoolName:              {authtypes.Burner, authtypes.Staking},
	stakingtypes.NotBondedPoolName:           {authtypes.Burner, authtypes.Staking},
	govtypes.ModuleName:                      {authtypes.Burner},
	ibctransfertypes.ModuleName:              {authtypes.Minter, authtypes.Burner},
	gammtypes.ModuleName:                     {authtypes.Minter, authtypes.Burner},
	incentivestypes.ModuleName:               {authtypes.Minter, authtypes.Burner},
	protorevtypes.ModuleName:                 {authtypes.Minter, authtypes.Burner},
	lockuptypes.ModuleName:                   {authtypes.Minter, authtypes.Burner},
	poolincentivestypes.ModuleName:           nil,
	superfluidtypes.ModuleName:               {authtypes.Minter, authtypes.Burner},
	txfeestypes.ModuleName:                   nil,
	txfeestypes.NonNativeFeeCollectorName:    nil,
	wasm.ModuleName:                          {authtypes.Burner},
	tokenfactorytypes.ModuleName:             {authtypes.Minter, authtypes.Burner},
	valsetpreftypes.ModuleName:               {authtypes.Staking},
}

// appModules return modules to initialize module manager.
func appModules(
	app *MerlinApp,
	encodingConfig appparams.EncodingConfig,
	skipGenesisInvariants bool,
) []module.AppModule {
	appCodec := encodingConfig.Marshaler

	return []module.AppModule{
		genutil.NewAppModule(
			app.AccountKeeper,
			app.StakingKeeper,
			app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, *app.AccountKeeper, nil),
		vesting.NewAppModule(*app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(appCodec, *app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		crisis.NewAppModule(app.CrisisKeeper, skipGenesisInvariants),
		gov.NewAppModule(appCodec, *app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, *app.MintKeeper, app.AccountKeeper, app.BankKeeper),
		slashing.NewAppModule(appCodec, *app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, *app.StakingKeeper),
		distr.NewAppModule(appCodec, *app.DistrKeeper, app.AccountKeeper, app.BankKeeper, *app.StakingKeeper),
		downtimemodule.NewAppModule(*app.DowntimeKeeper),
		staking.NewAppModule(appCodec, *app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		upgrade.NewAppModule(*app.UpgradeKeeper),
		wasm.NewAppModule(appCodec, app.WasmKeeper, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		evidence.NewAppModule(*app.EvidenceKeeper),
		authzmodule.NewAppModule(appCodec, *app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		ibc.NewAppModule(app.IBCKeeper),
		ica.NewAppModule(nil, app.ICAHostKeeper),
		params.NewAppModule(*app.ParamsKeeper),
		app.RawIcs20TransferAppModule,
		gamm.NewAppModule(appCodec, *app.GAMMKeeper, app.AccountKeeper, app.BankKeeper),
		poolmanager.NewAppModule(*app.PoolManagerKeeper, app.GAMMKeeper),
		twapmodule.NewAppModule(*app.TwapKeeper),
		concentratedliquidity.NewAppModule(appCodec, *app.ConcentratedLiquidityKeeper),
		protorev.NewAppModule(appCodec, *app.ProtoRevKeeper, app.AccountKeeper, app.BankKeeper, app.EpochsKeeper, app.GAMMKeeper),
		txfees.NewAppModule(*app.TxFeesKeeper),
		incentives.NewAppModule(*app.IncentivesKeeper, app.AccountKeeper, app.BankKeeper, app.EpochsKeeper),
		lockup.NewAppModule(*app.LockupKeeper, app.AccountKeeper, app.BankKeeper),
		poolincentives.NewAppModule(*app.PoolIncentivesKeeper),
		epochs.NewAppModule(*app.EpochsKeeper),
		superfluid.NewAppModule(
			*app.SuperfluidKeeper,
			app.AccountKeeper,
			app.BankKeeper,
			app.StakingKeeper,
			app.LockupKeeper,
			app.GAMMKeeper,
			app.EpochsKeeper,
		),
		tokenfactory.NewAppModule(*app.TokenFactoryKeeper, app.AccountKeeper, app.BankKeeper),
		valsetprefmodule.NewAppModule(appCodec, *app.ValidatorSetPreferenceKeeper),
		ibc_hooks.NewAppModule(app.AccountKeeper),
		icq.NewAppModule(*app.AppKeepers.ICQKeeper),
	}
}

// orderBeginBlockers returns the order of BeginBlockers, by module name.
func orderBeginBlockers(allModuleNames []string) []string {
	ord := partialord.NewPartialOrdering(allModuleNames)
	// Upgrades should be run VERY first
	// Epochs is set to be next right now, this in principle could change to come later / be at the end.
	// But would have to be a holistic change with other pipelines taken into account.
	ord.FirstElements(upgradetypes.ModuleName, epochstypes.ModuleName, capabilitytypes.ModuleName)

	// Staking ordering
	// TODO: Perhaps this can be relaxed, left to future work to analyze.
	ord.Sequence(distrtypes.ModuleName, slashingtypes.ModuleName, evidencetypes.ModuleName, stakingtypes.ModuleName)
	// superfluid must come after distribution & epochs.
	// TODO: we actually set it to come after staking, since thats what happened before, and want to minimize chance of break.
	ord.After(superfluidtypes.ModuleName, stakingtypes.ModuleName)
	// TODO: This can almost certainly be un-constrained, but we keep the constraint to match prior functionality.
	// IBChost came after staking, before superfluid.
	// TODO: Come back and delete this line after testing the base change.
	ord.Sequence(stakingtypes.ModuleName, ibchost.ModuleName, superfluidtypes.ModuleName)
	// We leave downtime-detector un-constrained.
	// every remaining module's begin block is a no-op.
	return ord.TotalOrdering()
}

// OrderEndBlockers returns EndBlockers (crisis, govtypes, staking) with no relative order.
func OrderEndBlockers(allModuleNames []string) []string {
	ord := partialord.NewPartialOrdering(allModuleNames)

	// Staking must be after gov.
	ord.FirstElements(govtypes.ModuleName)
	ord.LastElements(stakingtypes.ModuleName)

	// only Merlin modules with endblock code are: twap, crisis, govtypes, staking
	// we don't care about the relative ordering between them.
	return ord.TotalOrdering()
}

// OrderInitGenesis returns module names in order for init genesis calls.
func OrderInitGenesis(allModuleNames []string) []string {
	// NOTE: The genutils moodule must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	return []string{
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		downtimetypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		ibchost.ModuleName,
		icatypes.ModuleName,
		gammtypes.ModuleName,
		poolmanagertypes.ModuleName,
		protorevtypes.ModuleName,
		twaptypes.ModuleName,
		txfeestypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		ibctransfertypes.ModuleName,
		poolincentivestypes.ModuleName,
		superfluidtypes.ModuleName,
		tokenfactorytypes.ModuleName,
		valsetpreftypes.ModuleName,
		incentivestypes.ModuleName,
		epochstypes.ModuleName,
		lockuptypes.ModuleName,
		authz.ModuleName,
		concentratedliquiditytypes.ModuleName,
		// wasm after ibc transfer
		wasm.ModuleName,
		// ibc_hooks after auth keeper
		ibchookstypes.ModuleName,
		icqtypes.ModuleName,
	}
}

// ModuleAccountAddrs returns all the app's module account addresses.
func ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

func (app *MerlinApp) GetAccountKeeper() simtypes.AccountKeeper {
	return app.AppKeepers.AccountKeeper
}

func (app *MerlinApp) GetBankKeeper() simtypes.BankKeeper {
	return app.AppKeepers.BankKeeper
}

// Required for ibctesting
func (app *MerlinApp) GetStakingKeeper() stakingkeeper.Keeper {
	return *app.AppKeepers.StakingKeeper // Dereferencing the pointer
}

func (app *MerlinApp) GetIBCKeeper() *ibckeeper.Keeper {
	return app.AppKeepers.IBCKeeper // This is a *ibckeeper.Keeper
}

func (app *MerlinApp) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	return app.AppKeepers.ScopedIBCKeeper
}

func (app *MerlinApp) GetTxConfig() client.TxConfig {
	return MakeEncodingConfig().TxConfig
}
