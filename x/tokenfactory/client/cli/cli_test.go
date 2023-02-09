package cli_test

import (
	"testing"

	"github.com/merlinslair/merlin/osmoutils/osmocli"
	"github.com/merlinslair/merlin/x/tokenfactory/client/cli"
	"github.com/merlinslair/merlin/x/tokenfactory/types"
)

func TestGetCmdDenomAuthorityMetadata(t *testing.T) {
	desc, _ := cli.GetCmdDenomAuthorityMetadata()
	tcs := map[string]osmocli.QueryCliTestCase[*types.QueryDenomAuthorityMetadataRequest]{
		"basic test": {
			Cmd: "uatom",
			ExpectedQuery: &types.QueryDenomAuthorityMetadataRequest{
				Denom: "uatom",
			},
		},
	}
	osmocli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdDenomsFromCreator(t *testing.T) {
	desc, _ := cli.GetCmdDenomsFromCreator()
	tcs := map[string]osmocli.QueryCliTestCase[*types.QueryDenomsFromCreatorRequest]{
		"basic test": {
			Cmd: "mer1test",
			ExpectedQuery: &types.QueryDenomsFromCreatorRequest{
				Creator: "mer1test",
			},
		},
	}
	osmocli.RunQueryTestCases(t, desc, tcs)
}