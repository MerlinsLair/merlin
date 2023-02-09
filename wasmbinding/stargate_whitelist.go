package wasmbinding

import (
	"fmt"
	"sync"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	downtimequerytypes "github.com/merlinslair/merlin/x/downtime-detector/client/queryproto"
	epochtypes "github.com/merlinslair/merlin/x/epochs/types"
	gammtypes "github.com/merlinslair/merlin/x/gamm/types"
	gammv2types "github.com/merlinslair/merlin/x/gamm/v2types"
	incentivestypes "github.com/merlinslair/merlin/x/incentives/types"
	lockuptypes "github.com/merlinslair/merlin/x/lockup/types"
	minttypes "github.com/merlinslair/merlin/x/mint/types"
	poolincentivestypes "github.com/merlinslair/merlin/x/pool-incentives/types"
	poolmanagerqueryproto "github.com/merlinslair/merlin/x/poolmanager/client/queryproto"
	superfluidtypes "github.com/merlinslair/merlin/x/superfluid/types"
	tokenfactorytypes "github.com/merlinslair/merlin/x/tokenfactory/types"
	twapquerytypes "github.com/merlinslair/merlin/x/twap/client/queryproto"
	txfeestypes "github.com/merlinslair/merlin/x/txfees/types"
)

// stargateWhitelist keeps whitelist and its deterministic
// response binding for stargate queries.
//
// The query can be multi-thread, so we have to use
// thread safe sync.Map.
var stargateWhitelist sync.Map

//nolint:staticcheck
func init() {
	// cosmos-sdk queries

	// auth
	setWhitelistedQuery("/cosmos.auth.v1beta1.Query/Account", &authtypes.QueryAccountResponse{})
	setWhitelistedQuery("/cosmos.auth.v1beta1.Query/Params", &authtypes.QueryParamsResponse{})

	// bank
	setWhitelistedQuery("/cosmos.bank.v1beta1.Query/Balance", &banktypes.QueryBalanceResponse{})
	setWhitelistedQuery("/cosmos.bank.v1beta1.Query/DenomMetadata", &banktypes.QueryDenomsMetadataResponse{})
	setWhitelistedQuery("/cosmos.bank.v1beta1.Query/Params", &banktypes.QueryParamsResponse{})
	setWhitelistedQuery("/cosmos.bank.v1beta1.Query/SupplyOf", &banktypes.QuerySupplyOfResponse{})

	// distribution
	setWhitelistedQuery("/cosmos.distribution.v1beta1.Query/Params", &distributiontypes.QueryParamsResponse{})
	setWhitelistedQuery("/cosmos.distribution.v1beta1.Query/DelegatorWithdrawAddress", &distributiontypes.QueryDelegatorWithdrawAddressResponse{})
	setWhitelistedQuery("/cosmos.distribution.v1beta1.Query/ValidatorCommission", &distributiontypes.QueryValidatorCommissionResponse{})

	// gov
	setWhitelistedQuery("/cosmos.gov.v1beta1.Query/Deposit", &govtypes.QueryDepositResponse{})
	setWhitelistedQuery("/cosmos.gov.v1beta1.Query/Params", &govtypes.QueryParamsResponse{})
	setWhitelistedQuery("/cosmos.gov.v1beta1.Query/Vote", &govtypes.QueryVoteResponse{})

	// slashing
	setWhitelistedQuery("/cosmos.slashing.v1beta1.Query/Params", &slashingtypes.QueryParamsResponse{})
	setWhitelistedQuery("/cosmos.slashing.v1beta1.Query/SigningInfo", &slashingtypes.QuerySigningInfoResponse{})

	// staking
	setWhitelistedQuery("/cosmos.staking.v1beta1.Query/Delegation", &stakingtypes.QueryDelegationResponse{})
	setWhitelistedQuery("/cosmos.staking.v1beta1.Query/Params", &stakingtypes.QueryParamsResponse{})
	setWhitelistedQuery("/cosmos.staking.v1beta1.Query/Validator", &stakingtypes.QueryValidatorResponse{})

	// merlin queries

	// epochs
	setWhitelistedQuery("/merlin.epochs.v1beta1.Query/EpochInfos", &epochtypes.QueryEpochsInfoResponse{})
	setWhitelistedQuery("/merlin.epochs.v1beta1.Query/CurrentEpoch", &epochtypes.QueryCurrentEpochResponse{})

	// gamm
	setWhitelistedQuery("/merlin.gamm.v1beta1.Query/NumPools", &gammtypes.QueryNumPoolsResponse{})
	setWhitelistedQuery("/merlin.gamm.v1beta1.Query/TotalLiquidity", &gammtypes.QueryTotalLiquidityResponse{})
	setWhitelistedQuery("/merlin.gamm.v1beta1.Query/Pool", &gammtypes.QueryPoolResponse{})
	setWhitelistedQuery("/merlin.gamm.v1beta1.Query/PoolParams", &gammtypes.QueryPoolParamsResponse{})
	setWhitelistedQuery("/merlin.gamm.v1beta1.Query/TotalPoolLiquidity", &gammtypes.QueryTotalPoolLiquidityResponse{})
	setWhitelistedQuery("/merlin.gamm.v1beta1.Query/TotalShares", &gammtypes.QueryTotalSharesResponse{})
	setWhitelistedQuery("/merlin.gamm.v1beta1.Query/CalcJoinPoolShares", &gammtypes.QueryCalcJoinPoolSharesResponse{})
	setWhitelistedQuery("/merlin.gamm.v1beta1.Query/CalcExitPoolCoinsFromShares", &gammtypes.QueryCalcExitPoolCoinsFromSharesResponse{})
	setWhitelistedQuery("/merlin.gamm.v1beta1.Query/CalcJoinPoolNoSwapShares", &gammtypes.QueryCalcJoinPoolNoSwapSharesResponse{})
	setWhitelistedQuery("/merlin.gamm.v1beta1.Query/PoolType", &gammtypes.QueryPoolTypeResponse{})
	setWhitelistedQuery("/merlin.gamm.v2.Query/SpotPrice", &gammv2types.QuerySpotPriceResponse{})
	setWhitelistedQuery("/merlin.gamm.v1beta1.Query/EstimateSwapExactAmountIn", &gammtypes.QuerySwapExactAmountInResponse{})
	setWhitelistedQuery("/merlin.gamm.v1beta1.Query/EstimateSwapExactAmountOut", &gammtypes.QuerySwapExactAmountOutResponse{})

	// incentives
	setWhitelistedQuery("/merlin.incentives.Query/ModuleToDistributeCoins", &incentivestypes.ModuleToDistributeCoinsResponse{})
	setWhitelistedQuery("/merlin.incentives.Query/LockableDurations", &incentivestypes.QueryLockableDurationsResponse{})

	// lockup
	setWhitelistedQuery("/merlin.lockup.Query/ModuleBalance", &lockuptypes.ModuleBalanceResponse{})
	setWhitelistedQuery("/merlin.lockup.Query/ModuleLockedAmount", &lockuptypes.ModuleLockedAmountResponse{})
	setWhitelistedQuery("/merlin.lockup.Query/AccountUnlockableCoins", &lockuptypes.AccountUnlockableCoinsResponse{})
	setWhitelistedQuery("/merlin.lockup.Query/AccountUnlockingCoins", &lockuptypes.AccountUnlockingCoinsResponse{})
	setWhitelistedQuery("/merlin.lockup.Query/LockedDenom", &lockuptypes.LockedDenomResponse{})
	setWhitelistedQuery("/merlin.lockup.Query/LockedByID", &lockuptypes.LockedResponse{})

	// mint
	setWhitelistedQuery("/merlin.mint.v1beta1.Query/EpochProvisions", &minttypes.QueryEpochProvisionsResponse{})
	setWhitelistedQuery("/merlin.mint.v1beta1.Query/Params", &minttypes.QueryParamsResponse{})

	// pool-incentives
	setWhitelistedQuery("/merlin.poolincentives.v1beta1.Query/GaugeIds", &poolincentivestypes.QueryGaugeIdsResponse{})

	// superfluid
	setWhitelistedQuery("/merlin.superfluid.Query/Params", &superfluidtypes.QueryParamsResponse{})
	setWhitelistedQuery("/merlin.superfluid.Query/AssetType", &superfluidtypes.AssetTypeResponse{})
	setWhitelistedQuery("/merlin.superfluid.Query/AllAssets", &superfluidtypes.AllAssetsResponse{})
	setWhitelistedQuery("/merlin.superfluid.Query/AssetMultiplier", &superfluidtypes.AssetMultiplierResponse{})

	// poolmanager
	setWhitelistedQuery("/merlin.poolmanager.v1beta1.Query/NumPools", &poolmanagerqueryproto.NumPoolsResponse{})
	setWhitelistedQuery("/merlin.poolmanager.v1beta1.Query/EstimateSwapExactAmountIn", &poolmanagerqueryproto.EstimateSwapExactAmountInResponse{})
	setWhitelistedQuery("/merlin.poolmanager.v1beta1.Query/EstimateSwapExactAmountOut", &poolmanagerqueryproto.EstimateSwapExactAmountOutRequest{})

	// txfees
	setWhitelistedQuery("/merlin.txfees.v1beta1.Query/FeeTokens", &txfeestypes.QueryFeeTokensResponse{})
	setWhitelistedQuery("/merlin.txfees.v1beta1.Query/DenomSpotPrice", &txfeestypes.QueryDenomSpotPriceResponse{})
	setWhitelistedQuery("/merlin.txfees.v1beta1.Query/DenomPoolId", &txfeestypes.QueryDenomPoolIdResponse{})
	setWhitelistedQuery("/merlin.txfees.v1beta1.Query/BaseDenom", &txfeestypes.QueryBaseDenomResponse{})

	// tokenfactory
	setWhitelistedQuery("/merlin.tokenfactory.v1beta1.Query/params", &tokenfactorytypes.QueryParamsResponse{})
	setWhitelistedQuery("/merlin.tokenfactory.v1beta1.Query/DenomAuthorityMetadata", &tokenfactorytypes.QueryDenomAuthorityMetadataResponse{})
	// Does not include denoms_from_creator, TBD if this is the index we want contracts to use instead of admin

	// twap
	setWhitelistedQuery("/merlin.twap.v1beta1.Query/ArithmeticTwap", &twapquerytypes.ArithmeticTwapResponse{})
	setWhitelistedQuery("/merlin.twap.v1beta1.Query/ArithmeticTwapToNow", &twapquerytypes.ArithmeticTwapToNowResponse{})
	setWhitelistedQuery("/merlin.twap.v1beta1.Query/GeometricTwap", &twapquerytypes.GeometricTwapResponse{})
	setWhitelistedQuery("/merlin.twap.v1beta1.Query/GeometricTwapToNow", &twapquerytypes.GeometricTwapToNowResponse{})
	setWhitelistedQuery("/merlin.twap.v1beta1.Query/Params", &twapquerytypes.ParamsResponse{})

	// downtime-detector
	setWhitelistedQuery("/merlin.downtimedetector.v1beta1.Query/RecoveredSinceDowntimeOfLength", &downtimequerytypes.RecoveredSinceDowntimeOfLengthResponse{})
}

// GetWhitelistedQuery returns the whitelisted query at the provided path.
// If the query does not exist, or it was setup wrong by the chain, this returns an error.
func GetWhitelistedQuery(queryPath string) (codec.ProtoMarshaler, error) {
	protoResponseAny, isWhitelisted := stargateWhitelist.Load(queryPath)
	if !isWhitelisted {
		return nil, wasmvmtypes.UnsupportedRequest{Kind: fmt.Sprintf("'%s' path is not allowed from the contract", queryPath)}
	}
	protoResponseType, ok := protoResponseAny.(codec.ProtoMarshaler)
	if !ok {
		return nil, wasmvmtypes.Unknown{}
	}
	return protoResponseType, nil
}

func setWhitelistedQuery(queryPath string, protoType codec.ProtoMarshaler) {
	stargateWhitelist.Store(queryPath, protoType)
}
