package client

import (
	"github.com/merlinslair/merlin/x/pool-incentives/client/cli"
	"github.com/merlinslair/merlin/x/pool-incentives/client/rest"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var (
	UpdatePoolIncentivesHandler  = govclient.NewProposalHandler(cli.NewCmdSubmitUpdatePoolIncentivesProposal, rest.ProposalUpdatePoolIncentivesRESTHandler)
	ReplacePoolIncentivesHandler = govclient.NewProposalHandler(cli.NewCmdSubmitReplacePoolIncentivesProposal, rest.ProposalReplacePoolIncentivesRESTHandler)
)
