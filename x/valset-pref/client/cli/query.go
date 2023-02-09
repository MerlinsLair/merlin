package valsetprefcli

import (
	"github.com/spf13/cobra"

	"github.com/merlinslair/merlin/osmoutils/osmocli"
	"github.com/merlinslair/merlin/x/valset-pref/client/queryproto"
	"github.com/merlinslair/merlin/x/valset-pref/types"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd() *cobra.Command {
	cmd := osmocli.QueryIndexCmd(types.ModuleName)
	cmd.AddCommand(GetCmdValSetPref())
	return cmd
}

// GetCmdValSetPref takes the  address and returns the existing validator set for that address.
func GetCmdValSetPref() *cobra.Command {
	return osmocli.SimpleQueryCmd[*queryproto.UserValidatorPreferencesRequest](
		"val-set [address]",
		"Query the validator set for a specific user address", "",
		types.ModuleName, queryproto.NewQueryClient,
	)
}