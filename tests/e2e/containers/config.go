package containers

// ImageConfig contains all images and their respective tags
// needed for running e2e tests.
type ImageConfig struct {
	InitRepository string
	InitTag        string

	MerlinRepository string
	MerlinTag        string

	RelayerRepository string
	RelayerTag        string
}

//nolint:deadcode
const (
	// Current Git branch merlin repo/version. It is meant to be built locally.
	// It is used when skipping upgrade by setting MERLIN_E2E_SKIP_UPGRADE to true).
	// This image should be pre-built with `make docker-build-debug` either in CI or locally.
	CurrentBranchMerRepository = "merlin"
	CurrentBranchMerTag        = "debug"
	// Pre-upgrade merlin repo/tag to pull.
	// It should be uploaded to Docker Hub. MERLIN_E2E_SKIP_UPGRADE should be unset
	// for this functionality to be used.
	previousVersionMerRepository = "merlabs/merlin-dev"
	previousVersionMerTag        = "v14.x-6b742c84-1675172347"
	// Pre-upgrade repo/tag for merlin initialization (this should be one version below upgradeVersion)
	previousVersionInitRepository = "merlabs/merlin-e2e-init-chain"
	previousVersionInitTag        = "v14.x-f6813bcb-1674140667-manual"
	// Hermes repo/version for relayer
	relayerRepository = "merlabs/hermes"
	relayerTag        = "0.13.0"
)

// Returns ImageConfig needed for running e2e test.
// If isUpgrade is true, returns images for running the upgrade
// If isFork is true, utilizes provided fork height to initiate fork logic
func NewImageConfig(isUpgrade, isFork bool) ImageConfig {
	config := ImageConfig{
		RelayerRepository: relayerRepository,
		RelayerTag:        relayerTag,
	}

	if !isUpgrade {
		// If upgrade is not tested, we do not need InitRepository and InitTag
		// because we directly call the initialization logic without
		// the need for Docker.
		config.MerlinRepository = CurrentBranchMerRepository
		config.MerlinTag = CurrentBranchMerTag
		return config
	}

	// If upgrade is tested, we need to utilize InitRepository and InitTag
	// to initialize older state with Docker
	config.InitRepository = previousVersionInitRepository
	config.InitTag = previousVersionInitTag

	if isFork {
		// Forks are state compatible with earlier versions before fork height.
		// Normally, validators switch the binaries pre-fork height
		// Then, once the fork height is reached, the state breaking-logic
		// is run.
		config.MerlinRepository = CurrentBranchMerRepository
		config.MerlinTag = CurrentBranchMerTag
	} else {
		// Upgrades are run at the time when upgrade height is reached
		// and are submitted via a governance proposal. Thefore, we
		// must start running the previous Merlin version. Then, the node
		// should auto-upgrade, at which point we can restart the updated
		// Merlin validator container.
		config.MerlinRepository = previousVersionMerRepository
		config.MerlinTag = previousVersionMerTag
	}

	return config
}
