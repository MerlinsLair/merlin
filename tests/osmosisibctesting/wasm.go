package merlinibctesting

import (
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/stretchr/testify/require"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	transfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	"github.com/stretchr/testify/suite"

	"github.com/merlinslair/merlin/x/ibc-rate-limit/types"
)

func (chain *TestChain) StoreContractCode(suite *suite.Suite, path string) {
	merlinApp := chain.GetMerlinApp()

	govKeeper := merlinApp.GovKeeper
	wasmCode, err := os.ReadFile(path)
	suite.Require().NoError(err)

	addr := merlinApp.AccountKeeper.GetModuleAddress(govtypes.ModuleName)
	src := wasmtypes.StoreCodeProposalFixture(func(p *wasmtypes.StoreCodeProposal) {
		p.RunAs = addr.String()
		p.WASMByteCode = wasmCode
		checksum := sha256.Sum256(wasmCode)
		p.CodeHash = checksum[:]
	})

	// when stored
	storedProposal, err := govKeeper.SubmitProposal(chain.GetContext(), src, false)
	suite.Require().NoError(err)

	// and proposal execute
	handler := govKeeper.Router().GetRoute(storedProposal.ProposalRoute())
	err = handler(chain.GetContext(), storedProposal.GetContent())
	suite.Require().NoError(err)
}

func (chain *TestChain) InstantiateRLContract(suite *suite.Suite, quotas string) sdk.AccAddress {
	merlinApp := chain.GetMerlinApp()
	transferModule := merlinApp.AccountKeeper.GetModuleAddress(transfertypes.ModuleName)
	govModule := merlinApp.AccountKeeper.GetModuleAddress(govtypes.ModuleName)

	initMsgBz := []byte(fmt.Sprintf(`{
           "gov_module":  "%s",
           "ibc_module":"%s",
           "paths": [%s]
        }`,
		govModule, transferModule, quotas))

	contractKeeper := wasmkeeper.NewDefaultPermissionKeeper(merlinApp.WasmKeeper)
	codeID := uint64(1)
	creator := merlinApp.AccountKeeper.GetModuleAddress(govtypes.ModuleName)
	addr, _, err := contractKeeper.Instantiate(chain.GetContext(), codeID, creator, creator, initMsgBz, "rate limiting contract", nil)
	suite.Require().NoError(err)
	return addr
}

func (chain *TestChain) InstantiateContract(suite *suite.Suite, msg string, codeID uint64) sdk.AccAddress {
	merlinApp := chain.GetMerlinApp()
	contractKeeper := wasmkeeper.NewDefaultPermissionKeeper(merlinApp.WasmKeeper)
	creator := merlinApp.AccountKeeper.GetModuleAddress(govtypes.ModuleName)
	addr, _, err := contractKeeper.Instantiate(chain.GetContext(), codeID, creator, creator, []byte(msg), "contract", nil)
	suite.Require().NoError(err)
	return addr
}

func (chain *TestChain) QueryContract(suite *suite.Suite, contract sdk.AccAddress, key []byte) string {
	merlinApp := chain.GetMerlinApp()
	state, err := merlinApp.WasmKeeper.QuerySmart(chain.GetContext(), contract, key)
	suite.Require().NoError(err)
	return string(state)
}

func (chain *TestChain) RegisterRateLimitingContract(addr []byte) {
	addrStr, err := sdk.Bech32ifyAddressBytes("mer", addr)
	require.NoError(chain.T, err)
	params, err := types.NewParams(addrStr)
	require.NoError(chain.T, err)
	merlinApp := chain.GetMerlinApp()
	paramSpace, ok := merlinApp.AppKeepers.ParamsKeeper.GetSubspace(types.ModuleName)
	require.True(chain.T, ok)
	paramSpace.SetParamSet(chain.GetContext(), &params)
}
