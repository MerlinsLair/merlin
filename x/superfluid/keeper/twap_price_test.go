package keeper_test

import (
	"github.com/merlinslair/merlin/x/superfluid/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestMerEquivalentMultiplierSetGetDeleteFlow() {
	suite.SetupTest()

	// initial check
	multipliers := suite.App.SuperfluidKeeper.GetAllMerEquivalentMultipliers(suite.Ctx)
	suite.Require().Len(multipliers, 0)

	// set multiplier
	suite.App.SuperfluidKeeper.SetMerEquivalentMultiplier(suite.Ctx, 1, "gamm/pool/1", sdk.NewDec(2))

	// get multiplier
	multiplier := suite.App.SuperfluidKeeper.GetMerEquivalentMultiplier(suite.Ctx, "gamm/pool/1")
	suite.Require().Equal(multiplier, sdk.NewDec(2))

	// check multipliers
	expectedMultipliers := []types.MerEquivalentMultiplierRecord{
		{
			EpochNumber: 1,
			Denom:       "gamm/pool/1",
			Multiplier:  sdk.NewDec(2),
		},
	}
	multipliers = suite.App.SuperfluidKeeper.GetAllMerEquivalentMultipliers(suite.Ctx)
	suite.Require().Equal(multipliers, expectedMultipliers)

	// test last epoch price
	multiplier = suite.App.SuperfluidKeeper.GetMerEquivalentMultiplier(suite.Ctx, "gamm/pool/1")
	suite.Require().Equal(multiplier, sdk.NewDec(2))

	// delete multiplier
	suite.App.SuperfluidKeeper.DeleteMerEquivalentMultiplier(suite.Ctx, "gamm/pool/1")

	// get multiplier
	multiplier = suite.App.SuperfluidKeeper.GetMerEquivalentMultiplier(suite.Ctx, "gamm/pool/1")
	suite.Require().Equal(multiplier, sdk.NewDec(0))

	// check multipliers
	multipliers = suite.App.SuperfluidKeeper.GetAllMerEquivalentMultipliers(suite.Ctx)
	suite.Require().Len(multipliers, 0)

	// test last epoch price
	multiplier = suite.App.SuperfluidKeeper.GetMerEquivalentMultiplier(suite.Ctx, "gamm/pool/1")
	suite.Require().Equal(multiplier, sdk.NewDec(0))
}
