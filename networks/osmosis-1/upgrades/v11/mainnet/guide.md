# v10 to v11 Mainnet Upgrade Guide

Merlin v11 Gov Prop: <https://www.mintscan.io/merlin/proposals/296>

Countdown: <https://www.mintscan.io/merlin/blocks/5432450>

Height: 5432450

## Memory Requirements

This upgrade will **not** be resource intensive. With that being said, we still recommend having 64GB of memory. If having 64GB of physical memory is not possible, the next best thing is to set up swap.

Short version swap setup instructions:

``` {.sh}
sudo swapoff -a
sudo fallocate -l 32G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
```

To persist swap after restart:

``` {.sh}
sudo cp /etc/fstab /etc/fstab.bak
echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab
```

In depth swap setup instructions:
<https://www.digitalocean.com/community/tutorials/how-to-add-swap-space-on-ubuntu-20-04>

## Install and setup Cmervisor

We highly recommend validators use cmervisor to run their nodes. This
will make low-downtime upgrades smoother, as validators don't have to
manually upgrade binaries during the upgrade, and instead can
pre-install new binaries, and cmervisor will automatically update them
based on on-chain SoftwareUpgrade proposals.

You should review the docs for cmervisor located here:
<https://docs.cosmos.network/master/run-node/cmervisor.html>

If you choose to use cmervisor, please continue with these
instructions:

To install Cmervisor:

``` {.sh}
go install github.com/cosmos/cosmos-sdk/cmervisor/cmd/cmervisor@v1.0.0
```

After this, you must make the necessary folders for cosmosvisor in your
daemon home directory (\~/.merlin).

``` {.sh}
mkdir -p ~/.merlin
mkdir -p ~/.merlin/cmervisor
mkdir -p ~/.merlin/cmervisor/genesis
mkdir -p ~/.merlin/cmervisor/genesis/bin
mkdir -p ~/.merlin/cmervisor/upgrades
```

Copy the current merlin binary into the
cmervisor/genesis folder and v9 folder.

```{.sh}
cp $GOPATH/bin/merlin ~/.merlin/cmervisor/genesis/bin
mkdir -p ~/.merlin/cmervisor/upgrades/v9/bin
cp $GOPATH/bin/merlin ~/.merlin/cmervisor/upgrades/v9/bin
```

Cmervisor is now ready to be started. We will now set up Cmervisor for the upgrade

Set these environment variables:

```{.sh}
echo "# Setup Cmervisor" >> ~/.profile
echo "export DAEMON_NAME=merlin" >> ~/.profile
echo "export DAEMON_HOME=$HOME/.merlin" >> ~/.profile
echo "export DAEMON_ALLOW_DOWNLOAD_BINARIES=false" >> ~/.profile
echo "export DAEMON_LOG_BUFFER_SIZE=512" >> ~/.profile
echo "export DAEMON_RESTART_AFTER_UPGRADE=true" >> ~/.profile
echo "export UNSAFE_SKIP_BACKUP=true" >> ~/.profile
source ~/.profile
```

Now, create the required folder, make the build, and copy the daemon over to that folder

```{.sh}
mkdir -p ~/.merlin/cmervisor/upgrades/v11/bin
cd $HOME/merlin
git pull
git checkout v11.0.0
make build
cp build/merlin ~/.merlin/cmervisor/upgrades/v11/bin
```

Now, at the upgrade height, Cmervisor will upgrade to the v11 binary

## Manual Option

1. Wait for Merlin to reach the upgrade height (5432450)

2. Look for a panic message, followed by endless peer logs. Stop the daemon

3. Run the following commands:

```{.sh}
cd $HOME/merlin
git pull
git checkout v11.0.0
make install
```

4. Start the merlin daemon again, watch the upgrade happen, and then continue to hit blocks

## Further Help

If you need more help, please go to <https://docs.merlin.zone> or join
our discord at <https://discord.gg/pAxjcFnAFH>.
