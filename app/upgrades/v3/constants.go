package v3

import "github.com/merlinslair/merlin/app/upgrades"

const (
	// UpgradeName defines the on-chain upgrade name for the Merlin v3 upgrade.
	UpgradeName = "v3"

	// UpgradeHeight defines the block height at which the Merlin v3 upgrade is
	// triggered.
	UpgradeHeight = 712_000
)

var Fork = upgrades.Fork{
	UpgradeName:    UpgradeName,
	UpgradeHeight:  UpgradeHeight,
	BeginForkLogic: RunForkLogic,
}
