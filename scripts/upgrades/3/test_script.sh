# Create a genesis.json for testing. The node that you this on will be your "validator"
# It should be on version v3.0.0-rc0
merlin init --chain-id=testing testing --home=$HOME/.merlin
merlin keys add validator --keyring-backend=test --home=$HOME/.merlin
merlin add-genesis-account $(merlin keys show validator -a --keyring-backend=test --home=$HOME/.merlin) 1000000000umer,1000000000valtoken --home=$HOME/.merlin
sed -i -e "s/stake/umer/g" $HOME/.merlin/config/genesis.json
merlin gentx validator 500000000umer --commission-rate="0.0" --keyring-backend=test --home=$HOME/.merlin --chain-id=testing
merlin collect-gentxs --home=$HOME/.merlin

cat $HOME/.merlin/config/genesis.json | jq '.initial_height="711800"' > $HOME/.merlin/config/tmp_genesis.json && mv $HOME/.merlin/config/tmp_genesis.json $HOME/.merlin/config/genesis.json
cat $HOME/.merlin/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"]["denom"]="valtoken"' > $HOME/.merlin/config/tmp_genesis.json && mv $HOME/.merlin/config/tmp_genesis.json $HOME/.merlin/config/genesis.json
cat $HOME/.merlin/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"]["amount"]="100"' > $HOME/.merlin/config/tmp_genesis.json && mv $HOME/.merlin/config/tmp_genesis.json $HOME/.merlin/config/genesis.json
cat $HOME/.merlin/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="120s"' > $HOME/.merlin/config/tmp_genesis.json && mv $HOME/.merlin/config/tmp_genesis.json $HOME/.merlin/config/genesis.json
cat $HOME/.merlin/config/genesis.json | jq '.app_state["staking"]["params"]["min_commission_rate"]="0.050000000000000000"' > $HOME/.merlin/config/tmp_genesis.json && mv $HOME/.merlin/config/tmp_genesis.json $HOME/.merlin/config/genesis.json

# Now setup a second full node, and peer it with this v3.0.0-rc0 node.

# start the chain on both machines
merlin start
# Create proposals

merlin tx gov submit-proposal --title="existing passing prop" --description="passing prop"  --from=validator --deposit=1000valtoken --chain-id=testing --keyring-backend=test --broadcast-mode=block  --type="Text"
merlin tx gov vote 1 yes --from=validator --keyring-backend=test --chain-id=testing --yes
merlin tx gov submit-proposal --title="prop with enough mer deposit" --description="prop w/ enough deposit"  --from=validator --deposit=500000000umer --chain-id=testing --keyring-backend=test --broadcast-mode=block  --type="Text"
# Check that we have proposal 1 passed, and proposal 2 in deposit period
merlin q gov proposals
# CHeck that validator commission is under min_commission_rate
merlin q staking validators
# Wait for upgrade block.
# Upgrade happened
# your full node should have crashed with consensus failure

# Now we test post-upgrade behavior is as intended

# Everything in deposit stayed in deposit
merlin q gov proposals
# Check that commissions was bumped to min_commission_rate
merlin q staking validators
# pushes 2 into voting period
merlin tx gov deposit 2 1valtoken --from=validator --keyring-backend=test --chain-id=testing --yes