package v5

import (
	"github.com/merlinslair/merlin/app/upgrades"

	store "github.com/cosmos/cosmos-sdk/store/types"
)

// UpgradeName defines the on-chain upgrade name for the Merlin v5 upgrade.
const UpgradeName = "v5"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades:        store.StoreUpgrades{},
}
