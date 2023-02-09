#!/usr/bin/env bash

export REPO=https://github.com/merlinslair/merlin.git
export BRANCH=main
export NODE=merlin
export CLI=merlin
export MONIKER=testnet_node_moniker
# export GENESIS=https://raw.githubusercontent.com/cosmos/launch/master/genesis.json
# Thanatos - first ransomware to accept payment in bitcoin cash!
# (Also some greek god for death)
export CHAINID="mer-testnet-thanatos"

sudo apt update -y
sudo apt upgrade -y
sudo apt install -y make
sudo apt install -y build-essential
sudo apt install -y gcc

ulimit -n 65536
ulimit -u 65536

curl https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh | bash
source /root/.bashrc

export GOROOT=/root/.go
export PATH=$GOROOT/bin:$PATH
export GOPATH=/root/go
export PATH=$GOPATH/bin:$PATH

git clone $REPO
echo $(basename $REPO .git)

make LEDGER_ENABLED=false build
cp ./build/merlin /root/go/bin/

# install docker (https://docs.docker.com/engine/install/ubuntu/)

sudo apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo \
  "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

 sudo apt-get update -y
 sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-compose