# Download a genesis.json for testing. The node that you this on will be your "validator"
# It should be on version v4.x

merlind init --chain-id=testing testing --home=$HOME/.merlind
merlind keys add validator --keyring-backend=test --home=$HOME/.merlind
merlind add-genesis-account $(merlind keys show validator -a --keyring-backend=test --home=$HOME/.merlind) 1000000000umer,1000000000valtoken --home=$HOME/.merlind
sed -i -e "s/stake/umer/g" $HOME/.merlind/config/genesis.json
merlind gentx validator 500000000umer --commission-rate="0.0" --keyring-backend=test --home=$HOME/.merlind --chain-id=testing
merlind collect-gentxs --home=$HOME/.merlind

cat $HOME/.merlind/config/genesis.json | jq '.initial_height="711800"' > $HOME/.merlind/config/tmp_genesis.json && mv $HOME/.merlind/config/tmp_genesis.json $HOME/.merlind/config/genesis.json
cat $HOME/.merlind/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"]["denom"]="valtoken"' > $HOME/.merlind/config/tmp_genesis.json && mv $HOME/.merlind/config/tmp_genesis.json $HOME/.merlind/config/genesis.json
cat $HOME/.merlind/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"]["amount"]="100"' > $HOME/.merlind/config/tmp_genesis.json && mv $HOME/.merlind/config/tmp_genesis.json $HOME/.merlind/config/genesis.json
cat $HOME/.merlind/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="120s"' > $HOME/.merlind/config/tmp_genesis.json && mv $HOME/.merlind/config/tmp_genesis.json $HOME/.merlind/config/genesis.json
cat $HOME/.merlind/config/genesis.json | jq '.app_state["staking"]["params"]["min_commission_rate"]="0.050000000000000000"' > $HOME/.merlind/config/tmp_genesis.json && mv $HOME/.merlind/config/tmp_genesis.json $HOME/.merlind/config/genesis.json

# Now setup a second full node, and peer it with this v3.0.0-rc0 node.

# start the chain on both machines
merlind start
# Create proposals

merlind tx gov submit-proposal --title="existing passing prop" --description="passing prop"  --from=validator --deposit=1000valtoken --chain-id=testing --keyring-backend=test --broadcast-mode=block  --type="Text"
merlind tx gov vote 1 yes --from=validator --keyring-backend=test --chain-id=testing --yes
merlind tx gov submit-proposal --title="prop with enough mer deposit" --description="prop w/ enough deposit"  --from=validator --deposit=500000000umer --chain-id=testing --keyring-backend=test --broadcast-mode=block  --type="Text"
# Check that we have proposal 1 passed, and proposal 2 in deposit period
merlind q gov proposals
# CHeck that validator commission is under min_commission_rate
merlind q staking validators
# Wait for upgrade block.
# Upgrade happened
# your full node should have crashed with consensus failure

# Now we test post-upgrade behavior is as intended

# Everything in deposit stayed in deposit
merlind q gov proposals
# Check that commissions was bumped to min_commission_rate
merlind q staking validators
# pushes 2 into voting period
merlind tx gov deposit 2 1valtoken --from=validator --keyring-backend=test --chain-id=testing --yes