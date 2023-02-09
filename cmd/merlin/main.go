package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	merlin "github.com/merlinslair/merlin/app"
	"github.com/merlinslair/merlin/app/params"
	"github.com/merlinslair/merlin/cmd/merlin/cmd"
)

func main() {
	params.SetAddressPrefixes()
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, merlin.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
