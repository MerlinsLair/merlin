package wasmbinding

import (
	"crypto/sha256"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/CosmWasm/wasmd/x/wasm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/merlinslair/merlin/app"
)

func TestNoStorageWithoutProposal(t *testing.T) {
	// we use default config
	merlin, ctx := CreateTestInput()

	wasmKeeper := merlin.WasmKeeper
	// this wraps wasmKeeper, providing interfaces exposed to external messages
	contractKeeper := keeper.NewDefaultPermissionKeeper(wasmKeeper)

	_, _, creator := keyPubAddr()

	// upload reflect code
	wasmCode, err := os.ReadFile("../testdata/hackatom.wasm")
	require.NoError(t, err)
	_, _, err = contractKeeper.Create(ctx, creator, wasmCode, nil)
	require.Error(t, err)
}

func storeCodeViaProposal(t *testing.T, ctx sdk.Context, merlin *app.MerlinApp, addr sdk.AccAddress) {
	govKeeper := merlin.GovKeeper
	wasmCode, err := os.ReadFile("../testdata/hackatom.wasm")
	require.NoError(t, err)

	src := types.StoreCodeProposalFixture(func(p *types.StoreCodeProposal) {
		p.RunAs = addr.String()
		p.WASMByteCode = wasmCode
		checksum := sha256.Sum256(wasmCode)
		p.CodeHash = checksum[:]
	})

	// when stored
	storedProposal, err := govKeeper.SubmitProposal(ctx, src, false)
	require.NoError(t, err)

	// and proposal execute
	handler := govKeeper.Router().GetRoute(storedProposal.ProposalRoute())
	err = handler(ctx, storedProposal.GetContent())
	require.NoError(t, err)
}

func TestStoreCodeProposal(t *testing.T) {
	merlin, ctx := CreateTestInput()
	myActorAddress := RandomAccountAddress()
	wasmKeeper := merlin.WasmKeeper

	storeCodeViaProposal(t, ctx, merlin, myActorAddress)

	// then
	cInfo := wasmKeeper.GetCodeInfo(ctx, 1)
	require.NotNil(t, cInfo)
	assert.Equal(t, myActorAddress.String(), cInfo.Creator)
	assert.True(t, wasmKeeper.IsPinnedCode(ctx, 1))

	storedCode, err := wasmKeeper.GetByteCode(ctx, 1)
	require.NoError(t, err)
	wasmCode, err := os.ReadFile("../testdata/hackatom.wasm")
	require.NoError(t, err)
	assert.Equal(t, wasmCode, storedCode)
}

type HackatomExampleInitMsg struct {
	Verifier    sdk.AccAddress `json:"verifier"`
	Beneficiary sdk.AccAddress `json:"beneficiary"`
}

func TestInstantiateContract(t *testing.T) {
	merlin, ctx := CreateTestInput()
	funder := RandomAccountAddress()
	benefit, arb := RandomAccountAddress(), RandomAccountAddress()
	FundAccount(t, ctx, merlin, funder)

	storeCodeViaProposal(t, ctx, merlin, funder)
	contractKeeper := keeper.NewDefaultPermissionKeeper(merlin.WasmKeeper)
	codeID := uint64(1)

	initMsg := HackatomExampleInitMsg{
		Verifier:    arb,
		Beneficiary: benefit,
	}
	initMsgBz, err := json.Marshal(initMsg)
	require.NoError(t, err)

	funds := sdk.NewInt64Coin("umer", 123456)
	_, _, err = contractKeeper.Instantiate(ctx, codeID, funder, funder, initMsgBz, "demo contract", sdk.Coins{funds})
	require.NoError(t, err)
}
