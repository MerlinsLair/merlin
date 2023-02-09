# this script runs under the assumption that a three-validator environment is running on your local machine(multinode-local-testnet.sh)
# this script would do basic setup that has to be achieved to actual superfluid staking
# prior to running this script, have the following json file in the directory running this script
#
# stake-umer.json
# {
# 	"weights": "5stake,5umer",
# 	"initial-deposit": "1000000stake,1000000umer",
# 	"swap-fee": "0.01",
# 	"exit-fee": "0.01",
# 	"future-governor": "168h"
# }

# create pool
merlind tx gamm create-pool --pool-file=./stake-umer.json --from=validator1 --keyring-backend=test --chain-id=testing --yes --home=$HOME/.merlind/validator1
sleep 7

# test swap in pool created
merlind tx gamm swap-exact-amount-in 100000umer 50000 --swap-route-pool-ids=1 --swap-route-denoms=stake --from=validator1 --keyring-backend=test --chain-id=testing --yes --home=$HOME/.merlind/validator1
sleep 7

# create a lock up with lockable duration 360h
merlind tx lockup lock-tokens 10000000000000000000gamm/pool/1 --duration=360h --from=validator1 --keyring-backend=test --chain-id=testing --broadcast-mode=block --yes --home=$HOME/.merlind/validator1
sleep 7

# submit and pass proposal for superfluid
merlind tx gov submit-proposal set-superfluid-assets-proposal --title="set superfluid assets" --description="set superfluid assets description" --superfluid-assets="gamm/pool/1" --deposit=10000000umer --from=validator1 --chain-id=testing --keyring-backend=test --broadcast-mode=block --yes --home=$HOME/.merlind/validator1
sleep 7

merlind tx gov deposit 1 10000000stake --from=validator1 --keyring-backend=test --chain-id=testing --broadcast-mode=block --yes --home=$HOME/.merlind/validator1
sleep 7

merlind tx gov vote 1 yes --from=validator1 --keyring-backend=test --chain-id=testing --yes --home=$HOME/.merlind/validator1
sleep 7
merlind tx gov vote 1 yes --from=validator2 --keyring-backend=test --chain-id=testing --yes --home=$HOME/.merlind/validator2
sleep 7
