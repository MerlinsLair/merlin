package cli_test

import (
	gocontext "context"
	"testing"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/merlinslair/merlin/app/apptesting"
	"github.com/merlinslair/merlin/x/txfees/types"
)

type QueryTestSuite struct {
	apptesting.KeeperTestHelper
	queryClient types.QueryClient
}

func (s *QueryTestSuite) SetupSuite() {
	s.Setup()
	s.queryClient = types.NewQueryClient(s.QueryHelper)

	// set up pool
	poolAssets := []sdk.Coin{
		sdk.NewInt64Coin("umer", 1000000),
		sdk.NewInt64Coin("stake", 120000000),
	}
	s.PrepareBalancerPoolWithCoins(poolAssets...)

	// set up fee token
	upgradeProp := types.NewUpdateFeeTokenProposal(
		"Test Proposal",
		"test",
		types.FeeToken{
			Denom:  "umer",
			PoolID: 1,
		},
	)
	err := s.App.TxFeesKeeper.HandleUpdateFeeTokenProposal(s.Ctx, &upgradeProp)
	s.Require().NoError(err)

	s.Commit()
}

func (s *QueryTestSuite) TestQueriesNeverAlterState() {
	testCases := []struct {
		name   string
		query  string
		input  interface{}
		output interface{}
	}{
		{
			"Query base denom",
			"/merlin.txfees.v1beta1.Query/BaseDenom",
			&types.QueryBaseDenomRequest{},
			&types.QueryBaseDenomResponse{},
		},
		{
			"Query poolID by denom",
			"/merlin.txfees.v1beta1.Query/DenomPoolId",
			&types.QueryDenomPoolIdRequest{Denom: "umer"},
			&types.QueryDenomPoolIdResponse{},
		},
		{
			"Query spot price by denom",
			"/merlin.txfees.v1beta1.Query/DenomSpotPrice",
			&types.QueryDenomSpotPriceRequest{Denom: "umer"},
			&types.QueryDenomSpotPriceResponse{},
		},
		{
			"Query fee tokens",
			"/merlin.txfees.v1beta1.Query/FeeTokens",
			&types.QueryFeeTokensRequest{},
			&types.QueryFeeTokensResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			s.SetupSuite()
			err := s.QueryHelper.Invoke(gocontext.Background(), tc.query, tc.input, tc.output)
			s.Require().NoError(err)
			s.StateNotAltered()
		})
	}
}

func TestQueryTestSuite(t *testing.T) {
	suite.Run(t, new(QueryTestSuite))
}
