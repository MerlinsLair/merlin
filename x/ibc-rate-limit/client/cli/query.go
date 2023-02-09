package cli

import (
	"github.com/spf13/cobra"

	"github.com/merlinslair/merlin/x/ibc-rate-limit/types"
	"github.com/osmosis-labs/osmosis/osmoutils/osmocli"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd() *cobra.Command {
	cmd := osmocli.QueryIndexCmd(types.ModuleName)

	cmd.AddCommand(
		osmocli.GetParams[*types.QueryParamsRequest](
			types.ModuleName, types.NewQueryClient),
	)

	return cmd
}
