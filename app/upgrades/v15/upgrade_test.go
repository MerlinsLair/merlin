package v15_test

import (
	"reflect"
	"testing"

	gamm "github.com/merlinslair/merlin/x/gamm/keeper"

	"github.com/stretchr/testify/suite"

	"github.com/merlinslair/merlin/app/apptesting"
	v15 "github.com/merlinslair/merlin/app/upgrades/v15"
)

type UpgradeTestSuite struct {
	apptesting.KeeperTestHelper
}

func (suite *UpgradeTestSuite) SetupTest() {
	suite.Setup()
}

func TestUpgradeTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeTestSuite))
}

func (suite *UpgradeTestSuite) TestMigrateNextPoolIdAndCreatePool() {
	suite.SetupTest() // reset

	const (
		expectedNextPoolId uint64 = 1
	)

	var (
		gammKeeperType = reflect.TypeOf(&gamm.Keeper{})
	)

	ctx := suite.Ctx
	gammKeeper := suite.App.GAMMKeeper
	poolmanagerKeeper := suite.App.PoolManagerKeeper

	nextPoolId := gammKeeper.GetNextPoolId(ctx)
	suite.Require().Equal(expectedNextPoolId, nextPoolId)

	// system under test.
	v15.MigrateNextPoolId(ctx, gammKeeper, poolmanagerKeeper)

	// validate poolmanager's next pool id.
	actualNextPoolId := poolmanagerKeeper.GetNextPoolId(ctx)
	suite.Require().Equal(expectedNextPoolId, actualNextPoolId)

	// create a pool after migration.
	actualCreatedPoolId := suite.PrepareBalancerPool()
	suite.Require().Equal(expectedNextPoolId, actualCreatedPoolId)

	// validate that module route mapping has been created for each pool id.
	for poolId := uint64(1); poolId < expectedNextPoolId; poolId++ {
		swapModule, err := poolmanagerKeeper.GetPoolModule(ctx, poolId)
		suite.Require().NoError(err)

		suite.Require().Equal(gammKeeperType, reflect.TypeOf(swapModule))
	}

	// validate params
	gammPoolCreationFee := gammKeeper.GetParams(ctx).PoolCreationFee
	poolmanagerPoolCreationFee := poolmanagerKeeper.GetParams(ctx).PoolCreationFee
	suite.Require().Equal(gammPoolCreationFee, poolmanagerPoolCreationFee)
}

func (suite *UpgradeTestSuite) TestRegisterMerIonMetadata() {
	suite.SetupTest() // reset

	expectedUmerdenom := "umer"
	expectedUiondenom := "uion"

	ctx := suite.Ctx
	bankKeeper := suite.App.BankKeeper

	// meta data should not be found pre-registration of meta data
	umerMetadata, found := suite.App.BankKeeper.GetDenomMetaData(ctx, "umer")
	suite.Require().False(found)

	uionMetadata, found := suite.App.BankKeeper.GetDenomMetaData(ctx, "uion")
	suite.Require().False(found)

	// system under test.
	v15.RegisterMerIonMetadata(ctx, *bankKeeper)

	umerMetadata, found = suite.App.BankKeeper.GetDenomMetaData(ctx, "umer")
	suite.Require().True(found)

	uionMetadata, found = suite.App.BankKeeper.GetDenomMetaData(ctx, "uion")
	suite.Require().True(found)

	suite.Require().Equal(expectedUmerdenom, umerMetadata.Base)
	suite.Require().Equal(expectedUiondenom, uionMetadata.Base)
}
