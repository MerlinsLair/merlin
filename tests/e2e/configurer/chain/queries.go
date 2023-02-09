package chain

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
	tmabcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/merlinslair/merlin/tests/e2e/util"
	cltypes "github.com/merlinslair/merlin/x/concentrated-liquidity/types"
	epochstypes "github.com/merlinslair/merlin/x/epochs/types"
	gammtypes "github.com/merlinslair/merlin/x/gamm/types"
	poolmanagertypes "github.com/merlinslair/merlin/x/poolmanager/types"
	superfluidtypes "github.com/merlinslair/merlin/x/superfluid/types"
	twapqueryproto "github.com/merlinslair/merlin/x/twap/client/queryproto"
)

func (n *NodeConfig) QueryGRPCGateway(path string, parameters ...string) ([]byte, error) {
	if len(parameters)%2 != 0 {
		return nil, fmt.Errorf("invalid number of parameters, must follow the format of key + value")
	}

	// add the URL for the given validator ID, and pre-pend to to path.
	hostPort, err := n.containerManager.GetHostPort(n.Name, "1317/tcp")
	require.NoError(n.t, err)
	endpoint := fmt.Sprintf("http://%s", hostPort)
	fullQueryPath := fmt.Sprintf("%s/%s", endpoint, path)

	var resp *http.Response
	require.Eventually(n.t, func() bool {
		req, err := http.NewRequest("GET", fullQueryPath, nil)
		if err != nil {
			return false
		}

		if len(parameters) > 0 {
			q := req.URL.Query()
			for i := 0; i < len(parameters); i += 2 {
				q.Add(parameters[i], parameters[i+1])
			}
			req.URL.RawQuery = q.Encode()
		}

		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			n.t.Logf("error while executing HTTP request: %s", err.Error())
			return false
		}

		return resp.StatusCode != http.StatusServiceUnavailable
	}, time.Minute, time.Millisecond*10, "failed to execute HTTP request")

	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(bz))
	}
	return bz, nil
}

func (n *NodeConfig) QueryNumPools() uint64 {
	path := "merlin/gamm/v1beta1/num_pools"

	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	//nolint:staticcheck
	var numPools gammtypes.QueryNumPoolsResponse
	err = util.Cdc.UnmarshalJSON(bz, &numPools)
	require.NoError(n.t, err)
	return numPools.NumPools
}

func (n *NodeConfig) QueryConcentratedPositions(address string) []cltypes.FullPositionByOwnerResult {
	path := fmt.Sprintf("/merlin/concentratedliquidity/v1beta1/positions/%s", address)

	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var positionsResponse cltypes.QueryUserPositionsResponse
	err = util.Cdc.UnmarshalJSON(bz, &positionsResponse)
	require.NoError(n.t, err)
	return positionsResponse.Positions
}
func (n *NodeConfig) QueryConcentratedPool(poolId uint64) (cltypes.ConcentratedPoolExtension, error) {
	path := fmt.Sprintf("/merlin/concentratedliquidity/v1beta1/pools/%d", poolId)
	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var poolResponse cltypes.QueryPoolResponse
	err = util.Cdc.UnmarshalJSON(bz, &poolResponse)
	require.NoError(n.t, err)

	var pool poolmanagertypes.PoolI
	err = util.Cdc.UnpackAny(poolResponse.Pool, &pool)
	require.NoError(n.t, err)

	poolCLextension, ok := pool.(cltypes.ConcentratedPoolExtension)

	if !ok {
		return nil, fmt.Errorf("invalid pool type: %T", pool)
	}

	return poolCLextension, nil
}

// QueryBalancer returns balances at the address.
func (n *NodeConfig) QueryBalances(address string) (sdk.Coins, error) {
	path := fmt.Sprintf("cosmos/bank/v1beta1/balances/%s", address)
	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var balancesResp banktypes.QueryAllBalancesResponse
	if err := util.Cdc.UnmarshalJSON(bz, &balancesResp); err != nil {
		return sdk.Coins{}, err
	}
	return balancesResp.GetBalances(), nil
}

func (n *NodeConfig) QuerySupplyOf(denom string) (sdk.Int, error) {
	path := fmt.Sprintf("cosmos/bank/v1beta1/supply/%s", denom)
	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var supplyResp banktypes.QuerySupplyOfResponse
	if err := util.Cdc.UnmarshalJSON(bz, &supplyResp); err != nil {
		return sdk.NewInt(0), err
	}
	return supplyResp.Amount.Amount, nil
}

func (n *NodeConfig) QueryContractsFromId(codeId int) ([]string, error) {
	path := fmt.Sprintf("/cosmwasm/wasm/v1/code/%d/contracts", codeId)
	bz, err := n.QueryGRPCGateway(path)

	require.NoError(n.t, err)

	var contractsResponse wasmtypes.QueryContractsByCodeResponse
	if err := util.Cdc.UnmarshalJSON(bz, &contractsResponse); err != nil {
		return nil, err
	}

	return contractsResponse.Contracts, nil
}

func (n *NodeConfig) QueryLatestWasmCodeID() uint64 {
	path := "/cosmwasm/wasm/v1/code"

	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var response wasmtypes.QueryCodesResponse
	err = util.Cdc.UnmarshalJSON(bz, &response)
	require.NoError(n.t, err)
	if len(response.CodeInfos) == 0 {
		return 0
	}
	return response.CodeInfos[len(response.CodeInfos)-1].CodeID
}

func (n *NodeConfig) QueryWasmSmart(contract string, msg string) (map[string]interface{}, error) {
	// base64-encode the msg
	encodedMsg := base64.StdEncoding.EncodeToString([]byte(msg))
	path := fmt.Sprintf("/cosmwasm/wasm/v1/contract/%s/smart/%s", contract, encodedMsg)

	bz, err := n.QueryGRPCGateway(path)
	if err != nil {
		return nil, err
	}

	var response wasmtypes.QuerySmartContractStateResponse
	err = util.Cdc.UnmarshalJSON(bz, &response)
	if err != nil {
		return nil, err
	}

	var responseJSON map[string]interface{}
	err = json.Unmarshal(response.Data, &responseJSON)
	if err != nil {
		return nil, err
	}
	return responseJSON, nil
}

func (n *NodeConfig) QueryPropTally(proposalNumber int) (sdk.Int, sdk.Int, sdk.Int, sdk.Int, error) {
	path := fmt.Sprintf("cosmos/gov/v1beta1/proposals/%d/tally", proposalNumber)
	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var balancesResp govtypes.QueryTallyResultResponse
	if err := util.Cdc.UnmarshalJSON(bz, &balancesResp); err != nil {
		return sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt(), err
	}
	noTotal := balancesResp.Tally.No
	yesTotal := balancesResp.Tally.Yes
	noWithVetoTotal := balancesResp.Tally.NoWithVeto
	abstainTotal := balancesResp.Tally.Abstain

	return noTotal, yesTotal, noWithVetoTotal, abstainTotal, nil
}

func (n *NodeConfig) QueryPropStatus(proposalNumber int) (string, error) {
	path := fmt.Sprintf("cosmos/gov/v1beta1/proposals/%d", proposalNumber)
	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var propResp govtypes.QueryProposalResponse
	if err := util.Cdc.UnmarshalJSON(bz, &propResp); err != nil {
		return "", err
	}
	proposalStatus := propResp.Proposal.Status

	return proposalStatus.String(), nil
}

func (n *NodeConfig) QueryIntermediaryAccount(denom string, valAddr string) (int, error) {
	intAccount := superfluidtypes.GetSuperfluidIntermediaryAccountAddr(denom, valAddr)
	path := fmt.Sprintf(
		"cosmos/staking/v1beta1/validators/%s/delegations/%s",
		valAddr, intAccount,
	)

	bz, err := n.QueryGRPCGateway(path)
	require.NoError(n.t, err)

	var stakingResp stakingtypes.QueryDelegationResponse
	err = util.Cdc.UnmarshalJSON(bz, &stakingResp)
	require.NoError(n.t, err)

	intAccBalance := stakingResp.DelegationResponse.Balance.Amount.String()
	intAccountBalance, err := strconv.Atoi(intAccBalance)
	require.NoError(n.t, err)
	return intAccountBalance, err
}

func (n *NodeConfig) QueryCurrentEpoch(identifier string) int64 {
	path := "merlin/epochs/v1beta1/current_epoch"

	bz, err := n.QueryGRPCGateway(path, "identifier", identifier)
	require.NoError(n.t, err)

	var response epochstypes.QueryCurrentEpochResponse
	err = util.Cdc.UnmarshalJSON(bz, &response)
	require.NoError(n.t, err)
	return response.CurrentEpoch
}

func (n *NodeConfig) QueryArithmeticTwapToNow(poolId uint64, baseAsset, quoteAsset string, startTime time.Time) (sdk.Dec, error) {
	path := "merlin/twap/v1beta1/ArithmeticTwapToNow"

	bz, err := n.QueryGRPCGateway(
		path,
		"pool_id", strconv.FormatInt(int64(poolId), 10),
		"base_asset", baseAsset,
		"quote_asset", quoteAsset,
		"start_time", startTime.Format(time.RFC3339Nano),
	)
	if err != nil {
		return sdk.Dec{}, err
	}

	var response twapqueryproto.ArithmeticTwapToNowResponse
	err = util.Cdc.UnmarshalJSON(bz, &response)
	require.NoError(n.t, err) // this error should not happen
	return response.ArithmeticTwap, nil
}

func (n *NodeConfig) QueryArithmeticTwap(poolId uint64, baseAsset, quoteAsset string, startTime time.Time, endTime time.Time) (sdk.Dec, error) {
	path := "merlin/twap/v1beta1/ArithmeticTwap"

	bz, err := n.QueryGRPCGateway(
		path,
		"pool_id", strconv.FormatInt(int64(poolId), 10),
		"base_asset", baseAsset,
		"quote_asset", quoteAsset,
		"start_time", startTime.Format(time.RFC3339Nano),
		"end_time", endTime.Format(time.RFC3339Nano),
	)
	if err != nil {
		return sdk.Dec{}, err
	}

	var response twapqueryproto.ArithmeticTwapResponse
	err = util.Cdc.UnmarshalJSON(bz, &response)
	require.NoError(n.t, err) // this error should not happen
	return response.ArithmeticTwap, nil
}

func (n *NodeConfig) QueryGeometricTwapToNow(poolId uint64, baseAsset, quoteAsset string, startTime time.Time) (sdk.Dec, error) {
	path := "merlin/twap/v1beta1/GeometricTwapToNow"

	bz, err := n.QueryGRPCGateway(
		path,
		"pool_id", strconv.FormatInt(int64(poolId), 10),
		"base_asset", baseAsset,
		"quote_asset", quoteAsset,
		"start_time", startTime.Format(time.RFC3339Nano),
	)
	if err != nil {
		return sdk.Dec{}, err
	}

	var response twapqueryproto.GeometricTwapToNowResponse
	err = util.Cdc.UnmarshalJSON(bz, &response)
	require.NoError(n.t, err)
	return response.GeometricTwap, nil
}

func (n *NodeConfig) QueryGeometricTwap(poolId uint64, baseAsset, quoteAsset string, startTime time.Time, endTime time.Time) (sdk.Dec, error) {
	path := "merlin/twap/v1beta1/GeometricTwap"

	bz, err := n.QueryGRPCGateway(
		path,
		"pool_id", strconv.FormatInt(int64(poolId), 10),
		"base_asset", baseAsset,
		"quote_asset", quoteAsset,
		"start_time", startTime.Format(time.RFC3339Nano),
		"end_time", endTime.Format(time.RFC3339Nano),
	)
	if err != nil {
		return sdk.Dec{}, err
	}

	var response twapqueryproto.GeometricTwapResponse
	err = util.Cdc.UnmarshalJSON(bz, &response)
	require.NoError(n.t, err)
	return response.GeometricTwap, nil
}

// QueryHashFromBlock gets block hash at a specific height. Otherwise, error.
func (n *NodeConfig) QueryHashFromBlock(height int64) (string, error) {
	block, err := n.rpcClient.Block(context.Background(), &height)
	if err != nil {
		return "", err
	}
	return block.BlockID.Hash.String(), nil
}

// QueryCurrentHeight returns the current block height of the node or error.
func (n *NodeConfig) QueryCurrentHeight() (int64, error) {
	status, err := n.rpcClient.Status(context.Background())
	if err != nil {
		return 0, err
	}
	return status.SyncInfo.LatestBlockHeight, nil
}

// QueryLatestBlockTime returns the latest block time.
func (n *NodeConfig) QueryLatestBlockTime() time.Time {
	status, err := n.rpcClient.Status(context.Background())
	require.NoError(n.t, err)
	return status.SyncInfo.LatestBlockTime
}

// QueryListSnapshots gets all snapshots currently created for a node.
func (n *NodeConfig) QueryListSnapshots() ([]*tmabcitypes.Snapshot, error) {
	abciResponse, err := n.rpcClient.ABCIQuery(context.Background(), "/app/snapshots", nil)
	if err != nil {
		return nil, err
	}

	var listSnapshots tmabcitypes.ResponseListSnapshots
	if err := json.Unmarshal(abciResponse.Response.Value, &listSnapshots); err != nil {
		return nil, err
	}

	return listSnapshots.Snapshots, nil
}