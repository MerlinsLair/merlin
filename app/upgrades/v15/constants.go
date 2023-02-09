package v15

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	icqtypes "github.com/strangelove-ventures/async-icq/types"
	packetforwardtypes "github.com/strangelove-ventures/packet-forward-middleware/v4/router/types"

	"github.com/merlinslair/merlin/app/upgrades"
	cltypes "github.com/merlinslair/merlin/x/concentrated-liquidity/types"
	poolmanagertypes "github.com/merlinslair/merlin/x/poolmanager/types"
	protorevtypes "github.com/merlinslair/merlin/x/protorev/types"
	valsetpreftypes "github.com/merlinslair/merlin/x/valset-pref/types"
)

// UpgradeName defines the on-chain upgrade name for the Merlin v15 upgrade.
const UpgradeName = "v15"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{poolmanagertypes.StoreKey, cltypes.StoreKey, valsetpreftypes.StoreKey, protorevtypes.StoreKey, icqtypes.StoreKey, packetforwardtypes.StoreKey},
		Deleted: []string{},
	},
}
