package keeper

import (
	"github.com/gogo/protobuf/proto"

	gammtypes "github.com/merlinslair/merlin/x/gamm/types"
	"github.com/merlinslair/merlin/x/superfluid/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// This function calculates the mer equivalent worth of an LP share.
// It is intended to eventually use the TWAP of the worth of an LP share
// once that is exposed from the gamm module.
func (k Keeper) calculateMerBackingPerShare(pool gammtypes.CFMMPoolI, merInPool sdk.Int) sdk.Dec {
	twap := merInPool.ToDec().Quo(pool.GetTotalShares().ToDec())
	return twap
}

func (k Keeper) SetMerEquivalentMultiplier(ctx sdk.Context, epoch int64, denom string, multiplier sdk.Dec) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixTokenMultiplier)
	priceRecord := types.MerEquivalentMultiplierRecord{
		EpochNumber: epoch,
		Denom:       denom,
		Multiplier:  multiplier,
	}
	bz, err := proto.Marshal(&priceRecord)
	if err != nil {
		panic(err)
	}
	prefixStore.Set([]byte(denom), bz)
}

func (k Keeper) GetSuperfluidOSMOTokens(ctx sdk.Context, denom string, amount sdk.Int) sdk.Int {
	multiplier := k.GetMerEquivalentMultiplier(ctx, denom)
	if multiplier.IsZero() {
		return sdk.ZeroInt()
	}

	decAmt := multiplier.Mul(amount.ToDec())
	asset := k.GetSuperfluidAsset(ctx, denom)
	return k.GetRiskAdjustedMerValue(ctx, asset, decAmt.RoundInt())
}

func (k Keeper) DeleteMerEquivalentMultiplier(ctx sdk.Context, denom string) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixTokenMultiplier)
	prefixStore.Delete([]byte(denom))
}

func (k Keeper) GetMerEquivalentMultiplier(ctx sdk.Context, denom string) sdk.Dec {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixTokenMultiplier)
	bz := prefixStore.Get([]byte(denom))
	if bz == nil {
		return sdk.ZeroDec()
	}
	priceRecord := types.MerEquivalentMultiplierRecord{}
	err := proto.Unmarshal(bz, &priceRecord)
	if err != nil {
		panic(err)
	}
	return priceRecord.Multiplier
}

func (k Keeper) GetAllMerEquivalentMultipliers(ctx sdk.Context) []types.MerEquivalentMultiplierRecord {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.KeyPrefixTokenMultiplier)
	iterator := prefixStore.Iterator(nil, nil)
	defer iterator.Close()

	priceRecords := []types.MerEquivalentMultiplierRecord{}
	for ; iterator.Valid(); iterator.Next() {
		priceRecord := types.MerEquivalentMultiplierRecord{}

		err := proto.Unmarshal(iterator.Value(), &priceRecord)
		if err != nil {
			panic(err)
		}

		priceRecords = append(priceRecords, priceRecord)
	}
	return priceRecords
}
