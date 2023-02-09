#!/bin/sh
MERLIN_HOME="/tmp/merlin$(date +%s)"
RANDOM_KEY="randommerlinvalidatorkey"
CHAIN_ID=merlin-1
DENOM=umer
MAXBOND=50000000000000 # 500 Million MER

GENTX_FILE=$(find ./$CHAIN_ID/gentxs -iname "*.json")
LEN_GENTX=$(echo ${#GENTX_FILE})

# Gentx Start date
start="2021-06-03 15:00:00Z"
# Compute the seconds since epoch for start date
stTime=$(date --date="$start" +%s)

# Gentx End date
end="2021-07-12 23:59:59Z"
# Compute the seconds since epoch for end date
endTime=$(date --date="$end" +%s)

# Current date
current=$(date +%Y-%m-%d\ %H:%M:%S)
# Compute the seconds since epoch for current date
curTime=$(date --date="$current" +%s)

if [[ $curTime < $stTime ]]; then
    echo "start=$stTime:curent=$curTime:endTime=$endTime"
    echo "Gentx submission is not open yet. Please close the PR and raise a new PR after 04-June-2021 23:59:59"
    exit 0
else
    if [[ $curTime > $endTime ]]; then
        echo "start=$stTime:curent=$curTime:endTime=$endTime"
        echo "Gentx submission is closed"
        exit 0
    else
        echo "Gentx is now open"
        echo "start=$stTime:curent=$curTime:endTime=$endTime"
    fi
fi

if [ $LEN_GENTX -eq 0 ]; then
    echo "No new gentx file found."
else
    set -e

    echo "GentxFile::::"
    echo $GENTX_FILE

    echo "...........Init Merlin.............."

    git clone https://github.com/merlinslair/merlin
    cd merlin
    git checkout gentx-launch
    make build
    chmod +x ./build/merlin

    ./build/merlin keys add $RANDOM_KEY --keyring-backend test --home $MERLIN_HOME

    ./build/merlin init --chain-id $CHAIN_ID validator --home $MERLIN_HOME

    echo "..........Fetching genesis......."
    rm -rf $MERLIN_HOME/config/genesis.json
    curl -s https://raw.githubusercontent.com/osmosis-labs/networks/main/$CHAIN_ID/pregenesis.json >$MERLIN_HOME/config/genesis.json

    # this genesis time is different from original genesis time, just for validating gentx.
    sed -i '/genesis_time/c\   \"genesis_time\" : \"2021-03-29T00:00:00Z\",' $MERLIN_HOME/config/genesis.json

    GENACC=$(cat ../$GENTX_FILE | sed -n 's|.*"delegator_address":"\([^"]*\)".*|\1|p')
    denomquery=$(jq -r '.body.messages[0].value.denom' ../$GENTX_FILE)
    amountquery=$(jq -r '.body.messages[0].value.amount' ../$GENTX_FILE)

    echo $GENACC
    echo $amountquery
    echo $denomquery

    # only allow $DENOM tokens to be bonded
    if [ $denomquery != $DENOM ]; then
        echo "invalid denomination"
        exit 1
    fi

    # limit the amount that can be bonded

    if [ $amountquery -gt $MAXBOND ]; then
        echo "bonded too much: $amountquery > $MAXBOND"
        exit 1
    fi

    ./build/merlin add-genesis-account $RANDOM_KEY 100000000000000$DENOM --home $MERLIN_HOME \
        --keyring-backend test

    ./build/merlin gentx $RANDOM_KEY 90000000000000$DENOM --home $MERLIN_HOME \
        --keyring-backend test --chain-id $CHAIN_ID

    cp ../$GENTX_FILE $MERLIN_HOME/config/gentx/

    echo "..........Collecting gentxs......."
    ./build/merlin collect-gentxs --home $MERLIN_HOME
    sed -i '/persistent_peers =/c\persistent_peers = ""' $MERLIN_HOME/config/config.toml

    ./build/merlin validate-genesis --home $MERLIN_HOME

    echo "..........Starting node......."
    ./build/merlin start --home $MERLIN_HOME &

    sleep 1800s

    echo "...checking network status.."

    ./build/merlin status --node http://localhost:26657

    echo "...Cleaning the stuff..."
    killall merlin >/dev/null 2>&1
    rm -rf $MERLIN_HOME >/dev/null 2>&1
fi
