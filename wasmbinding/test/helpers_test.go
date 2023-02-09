package wasmbinding

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/merlinslair/merlin/app"
)

func CreateTestInput() (*app.MerlinApp, sdk.Context) {
	merlin := app.Setup(false)
	ctx := merlin.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "merlin-1", Time: time.Now().UTC()})
	return merlin, ctx
}

func FundAccount(t *testing.T, ctx sdk.Context, merlin *app.MerlinApp, acct sdk.AccAddress) {
	err := simapp.FundAccount(merlin.BankKeeper, ctx, acct, sdk.NewCoins(
		sdk.NewCoin("umer", sdk.NewInt(10000000000)),
	))
	require.NoError(t, err)
}

// we need to make this deterministic (same every test run), as content might affect gas costs
func keyPubAddr() (crypto.PrivKey, crypto.PubKey, sdk.AccAddress) {
	key := ed25519.GenPrivKey()
	pub := key.PubKey()
	addr := sdk.AccAddress(pub.Address())
	return key, pub, addr
}

func RandomAccountAddress() sdk.AccAddress {
	_, _, addr := keyPubAddr()
	return addr
}

func RandomBech32AccountAddress() string {
	return RandomAccountAddress().String()
}
