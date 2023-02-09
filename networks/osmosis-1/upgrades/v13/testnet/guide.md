# v12 to v13 Testnet Upgrade Guide

Countdown: <https://testnet.mintscan.io/merlin-testnet/blocks/7759340>

Height: 7759340

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

## First Time Cmervisor Setup

If you have never setup Cmervisor before, follow the following instructions.

If you have already setup Cmervisor, skip to the next section.

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
daemon home directory (\~/.merlind).

``` {.sh}
mkdir -p ~/.merlind
mkdir -p ~/.merlind/cmervisor
mkdir -p ~/.merlind/cmervisor/genesis
mkdir -p ~/.merlind/cmervisor/genesis/bin
mkdir -p ~/.merlind/cmervisor/upgrades
```

Copy the current v12 merlind binary into the
cmervisor/genesis folder and v12 folder.

```{.sh}
cp $GOPATH/bin/merlind ~/.merlind/cmervisor/genesis/bin
mkdir -p ~/.merlind/cmervisor/upgrades/v12/bin
cp $GOPATH/bin/merlind ~/.merlind/cmervisor/upgrades/v12/bin
```

Cmervisor is now ready to be set up for v13.

Set these environment variables:

```{.sh}
echo "# Setup Cmervisor" >> ~/.profile
echo "export DAEMON_NAME=merlind" >> ~/.profile
echo "export DAEMON_HOME=$HOME/.merlind" >> ~/.profile
echo "export DAEMON_ALLOW_DOWNLOAD_BINARIES=false" >> ~/.profile
echo "export DAEMON_LOG_BUFFER_SIZE=512" >> ~/.profile
echo "export DAEMON_RESTART_AFTER_UPGRADE=true" >> ~/.profile
echo "export UNSAFE_SKIP_BACKUP=true" >> ~/.profile
source ~/.profile
```

## Cmervisor Upgrade

Create the v13 folder, make the build, and copy the daemon over to that folder

```{.sh}
mkdir -p ~/.merlind/cmervisor/upgrades/v13/bin
cd $HOME/merlin
git pull
git checkout v13.0.0-rc3
make build
cp build/merlind ~/.merlind/cmervisor/upgrades/v13/bin
```

Now, at the upgrade height, Cmervisor will upgrade to the v13 binary

## Manual Option

1. Wait for Merlin to reach the upgrade height (7759340)

2. Look for a panic message, followed by endless peer logs. Stop the daemon

3. Run the following commands:

```{.sh}
cd $HOME/merlin
git pull
git checkout v13.0.0-rc3
make install
```

4. Start the merlin daemon again, watch the upgrade happen, and then continue to hit blocks

## Further Help

If you need more help, please go to <https://docs.merlin.zone> or join
our discord at <https://discord.gg/pAxjcFnAFH>.
