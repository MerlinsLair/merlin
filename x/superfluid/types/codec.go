package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSuperfluidDelegate{}, "merlin/superfluid-delegate", nil)
	cdc.RegisterConcrete(&MsgSuperfluidUndelegate{}, "merlin/superfluid-undelegate", nil)
	cdc.RegisterConcrete(&MsgLockAndSuperfluidDelegate{}, "merlin/lock-and-superfluid-delegate", nil)
	cdc.RegisterConcrete(&MsgSuperfluidUnbondLock{}, "merlin/superfluid-unbond-lock", nil)
	cdc.RegisterConcrete(&MsgSuperfluidUndelegateAndUnbondLock{}, "merlin/sf-undelegate-and-unbond-lock", nil)
	cdc.RegisterConcrete(&SetSuperfluidAssetsProposal{}, "merlin/set-superfluid-assets-proposal", nil)
	cdc.RegisterConcrete(&UpdateUnpoolWhiteListProposal{}, "merlin/update-unpool-whitelist", nil)
	cdc.RegisterConcrete(&RemoveSuperfluidAssetsProposal{}, "merlin/del-superfluid-assets-proposal", nil)
	cdc.RegisterConcrete(&MsgUnPoolWhitelistedPool{}, "merlin/unpool-whitelisted-pool", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgSuperfluidDelegate{},
		&MsgSuperfluidUndelegate{},
		&MsgLockAndSuperfluidDelegate{},
		&MsgSuperfluidUnbondLock{},
		&MsgSuperfluidUndelegateAndUnbondLock{},
		&MsgUnPoolWhitelistedPool{},
	)

	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&SetSuperfluidAssetsProposal{},
		&RemoveSuperfluidAssetsProposal{},
		&UpdateUnpoolWhiteListProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	RegisterCodec(authzcodec.Amino)
	amino.Seal()
}
