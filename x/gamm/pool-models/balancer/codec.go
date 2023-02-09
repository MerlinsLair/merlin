package balancer

import (
	types "github.com/merlinslair/merlin/x/gamm/types"
	poolmanagertypes "github.com/merlinslair/merlin/x/poolmanager/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	proto "github.com/gogo/protobuf/proto"
)

// RegisterLegacyAminoCodec registers the necessary x/gamm interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&Pool{}, "merlin/gamm/BalancerPool", nil)
	cdc.RegisterConcrete(&MsgCreateBalancerPool{}, "merlin/gamm/create-balancer-pool", nil)
	cdc.RegisterConcrete(&PoolParams{}, "merlin/gamm/BalancerPoolParams", nil)
	cdc.RegisterConcrete(&MsgMigrateSharesToFullRangeConcentratedPosition{}, "merlin/gamm/MigratePosition", nil)
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"merlin.poolmanager.v1beta1.PoolI",
		(*poolmanagertypes.PoolI)(nil),
		&Pool{},
	)
	registry.RegisterInterface(
		"merlin.gamm.v1beta1.PoolI", // N.B.: the old proto-path is preserved for backwards-compatibility.
		(*types.CFMMPoolI)(nil),
		&Pool{},
	)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateBalancerPool{},
		&MsgMigrateSharesToFullRangeConcentratedPosition{},
	)
	registry.RegisterImplementations(
		(*proto.Message)(nil),
		&PoolParams{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/bank module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/staking and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	amino.Seal()
}
