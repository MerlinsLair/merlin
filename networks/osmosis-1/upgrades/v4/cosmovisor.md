# Install and setup Cmervisor

We highly recommend validators use cmervisor to run their nodes. This
will make low-downtime upgrades smoother, as validators don't have to
manually upgrade binaries during the upgrade, and instead can preinstall
new binaries, and cmervisor will automatically update them based on
on-chain SoftwareUpgrade proposals.

You should review the docs for cmervisor located here:
<https://docs.cosmos.network/master/run-node/cmervisor.html>

If you choose to use cmervisor, please continue with these
instructions:

To install Cmervisor:

    git clone https://github.com/cosmos/cosmos-sdk
    cd cosmos-sdk
    git checkout v0.42.9
    make cmervisor
    cp cmervisor/cmervisor $GOPATH/bin/cmervisor
    cd $HOME

After this, you must make the necessary folders for cosmosvisor in your
daemon home directory (\~/.merlin).

``` {.sh}
mkdir -p ~/.merlin
mkdir -p ~/.merlin/cmervisor
mkdir -p ~/.merlin/cmervisor/genesis
mkdir -p ~/.merlin/cmervisor/genesis/bin
mkdir -p ~/.merlin/cmervisor/upgrades
```

Cmervisor requires some ENVIRONMENT VARIABLES be set in order to
function properly. We recommend setting these in your `.profile` so it
is automatically set in every session.

For validators we recommmend setting

- `DAEMON_ALLOW_DOWNLOAD_BINARIES=false` for security reasons
- `DAEMON_LOG_BUFFER_SIZE=512` to avoid a bug with extra long log
    lines crashing the server.
- `DAEMON_RESTART_AFTER_UPGRADE=true` for unattended upgrades

```{=html}
<!-- -->
```

    echo "# Setup Cmervisor" >> ~/.profile
    echo "export DAEMON_NAME=merlin" >> ~/.profile
    echo "export DAEMON_HOME=$HOME/.merlin" >> ~/.profile
    echo "export DAEMON_ALLOW_DOWNLOAD_BINARIES=false" >> ~/.profile
    echo "export DAEMON_LOG_BUFFER_SIZE=512" >> ~/.profile
    echo "export DAEMON_RESTART_AFTER_UPGRADE=true" >> ~/.profile
    source ~/.profile

Finally, you should copy the current merlin binary into the
cmervisor/genesis folder.

    cp $GOPATH/bin/merlin ~/.merlin/cmervisor/genesis/bin

Prepare for upgrade (v4)
------------------------

To prepare for the upgrade, you need to create some folders, and build
and install the new binary.

    mkdir -p ~/.merlin/cmervisor/upgrades/v4/bin
    git clone https://github.com/merlinslair/merlin
    cd merlin
    git checkout v4.0.0
    make build
    cp build/merlin ~/.merlin/cmervisor/upgrades/v4/bin

Now cmervisor will run with the current binary, and will automatically
upgrade to this new binary at the appropriate height if run with:

    cmervisor start

Please note, this does not automatically update your
`$GOPATH/bin/merlin` binary, to do that after the upgrade, please run
`make install` in the merlin source folder.
