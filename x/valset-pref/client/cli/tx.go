package valsetprefcli

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/merlinslair/merlin/x/valset-pref/types"
	"github.com/osmosis-labs/osmosis/osmoutils"
	"github.com/osmosis-labs/osmosis/osmoutils/osmocli"
)

func GetTxCmd() *cobra.Command {
	txCmd := osmocli.TxIndexCmd(types.ModuleName)
	osmocli.AddTxCmd(txCmd, NewSetValSetCmd)
	osmocli.AddTxCmd(txCmd, NewDelValSetCmd)
	osmocli.AddTxCmd(txCmd, NewUnDelValSetCmd)
	osmocli.AddTxCmd(txCmd, NewReDelValSetCmd)
	osmocli.AddTxCmd(txCmd, NewWithRewValSetCmd)
	return txCmd
}

func NewSetValSetCmd() (*osmocli.TxCliDesc, *types.MsgSetValidatorSetPreference) {
	return &osmocli.TxCliDesc{
		Use:              "set-valset [delegator_addr] [validators] [weights]",
		Short:            "Creates a new validator set for the delegator with valOperAddress and weight",
		Example:          "merlin tx valset-pref set-valset mer1... mervaloper1abc...,mervaloper1def...  0.56,0.44",
		NumArgs:          3,
		ParseAndBuildMsg: NewMsgSetValidatorSetPreference,
	}, &types.MsgSetValidatorSetPreference{}
}

func NewDelValSetCmd() (*osmocli.TxCliDesc, *types.MsgDelegateToValidatorSet) {
	return &osmocli.TxCliDesc{
		Use:     "delegate-valset [delegator_addr] [amount]",
		Short:   "Delegate tokens to existing valset using delegatorAddress and tokenAmount.",
		Example: "merlin tx valset-pref delegate-valset mer1... 100stake",
		NumArgs: 2,
	}, &types.MsgDelegateToValidatorSet{}
}

func NewUnDelValSetCmd() (*osmocli.TxCliDesc, *types.MsgUndelegateFromValidatorSet) {
	return &osmocli.TxCliDesc{
		Use:     "undelegate-valset [delegator_addr] [amount]",
		Short:   "UnDelegate tokens from existing valset using delegatorAddress and tokenAmount.",
		Example: "merlin tx valset-pref undelegate-valset mer1... 100stake",
		NumArgs: 2,
	}, &types.MsgUndelegateFromValidatorSet{}
}

func NewReDelValSetCmd() (*osmocli.TxCliDesc, *types.MsgRedelegateValidatorSet) {
	return &osmocli.TxCliDesc{
		Use:              "redelegate-valset [delegator_addr] [validators] [weights]",
		Short:            "Redelegate tokens from existing validators to new sets of validators",
		Example:          "merlin tx valset-pref redelegate-valset  mer1... mervaloper1efg...,mervaloper1hij...  0.56,0.44",
		NumArgs:          3,
		ParseAndBuildMsg: NewMsgReDelValidatorSetPreference,
	}, &types.MsgRedelegateValidatorSet{}
}

func NewWithRewValSetCmd() (*osmocli.TxCliDesc, *types.MsgWithdrawDelegationRewards) {
	return &osmocli.TxCliDesc{
		Use:     "withdraw-reward-valset [delegator_addr]",
		Short:   "Withdraw delegation reward form the new validator set.",
		Example: "merlin tx valset-pref withdraw-valset mer1...",
		NumArgs: 1,
	}, &types.MsgWithdrawDelegationRewards{}
}

func NewMsgSetValidatorSetPreference(clientCtx client.Context, args []string, fs *pflag.FlagSet) (sdk.Msg, error) {
	delAddr, err := sdk.AccAddressFromBech32(args[0])
	if err != nil {
		return nil, err
	}

	valset, err := ValidateValAddrAndWeight(args)
	if err != nil {
		return nil, err
	}

	return types.NewMsgSetValidatorSetPreference(
		delAddr,
		valset,
	), nil
}

func NewMsgReDelValidatorSetPreference(clientCtx client.Context, args []string, fs *pflag.FlagSet) (sdk.Msg, error) {
	delAddr, err := sdk.AccAddressFromBech32(args[0])
	if err != nil {
		return nil, err
	}

	valset, err := ValidateValAddrAndWeight(args)
	if err != nil {
		return nil, err
	}

	return types.NewMsgRedelegateValidatorSet(
		delAddr,
		valset,
	), nil
}

func ValidateValAddrAndWeight(args []string) ([]types.ValidatorPreference, error) {
	var valAddrs []string
	valAddrs = append(valAddrs, strings.Split(args[1], ",")...)

	weights, err := osmoutils.ParseSdkDecFromString(args[2], ",")
	if err != nil {
		return nil, err
	}

	if len(valAddrs) != len(weights) {
		return nil, fmt.Errorf("the length of validator addresses and weights not matched")
	}

	if len(valAddrs) == 0 {
		return nil, fmt.Errorf("records is empty")
	}

	var valset []types.ValidatorPreference
	for i, val := range valAddrs {
		valset = append(valset, types.ValidatorPreference{
			ValOperAddress: val,
			Weight:         weights[i],
		})
	}

	return valset, nil
}
