package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/merlinslair/merlin/x/protorev/types"
)

func TestGenesisStateValidate(t *testing.T) {
	cases := []struct {
		description string
		genState    *types.GenesisState
		valid       bool
	}{
		{
			description: "Default parameters with no routes",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
			},
			valid: true,
		},
		{
			description: "Default parameters with valid routes",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				TokenPairs: []types.TokenPairArbRoutes{
					{
						ArbRoutes: []*types.Route{{
							Trades: []*types.Trade{
								{
									Pool:     1,
									TokenIn:  types.AtomDenomination,
									TokenOut: "Juno",
								},
								{
									Pool:     0,
									TokenIn:  "Juno",
									TokenOut: types.MerlinDenomination,
								},
								{
									Pool:     3,
									TokenIn:  types.MerlinDenomination,
									TokenOut: types.AtomDenomination,
								},
							},
						}},
						TokenIn:  types.MerlinDenomination,
						TokenOut: "Juno",
					},
				},
			},
			valid: true,
		},
		{
			description: "Default parameters with invalid routes (duplicate token pairs)",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				TokenPairs: []types.TokenPairArbRoutes{
					{
						ArbRoutes: []*types.Route{
							{
								Trades: []*types.Trade{
									{
										Pool:     1,
										TokenIn:  types.AtomDenomination,
										TokenOut: "Juno",
									},
									{
										Pool:     0,
										TokenIn:  "Juno",
										TokenOut: types.MerlinDenomination,
									},
									{
										Pool:     3,
										TokenIn:  types.MerlinDenomination,
										TokenOut: types.AtomDenomination,
									},
								},
							},
						},
						TokenIn:  types.MerlinDenomination,
						TokenOut: "Juno",
					},
					{
						ArbRoutes: []*types.Route{
							{
								Trades: []*types.Trade{
									{
										Pool:     1,
										TokenIn:  types.AtomDenomination,
										TokenOut: "Juno",
									},
									{
										Pool:     0,
										TokenIn:  "Juno",
										TokenOut: types.MerlinDenomination,
									},
									{
										Pool:     3,
										TokenIn:  types.MerlinDenomination,
										TokenOut: types.AtomDenomination,
									},
								},
							},
						},
						TokenIn:  types.MerlinDenomination,
						TokenOut: "Juno",
					},
				},
			},
			valid: false,
		},
		{
			description: "Default parameters with nil routes",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				TokenPairs: []types.TokenPairArbRoutes{
					{
						ArbRoutes: nil,
						TokenIn:   types.MerlinDenomination,
						TokenOut:  "Juno",
					},
				},
			},
			valid: false,
		},
		{
			description: "Default parameters with invalid routes (too few trades in a route)",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				TokenPairs: []types.TokenPairArbRoutes{
					{
						ArbRoutes: []*types.Route{
							{
								Trades: []*types.Trade{
									{
										Pool:     3,
										TokenIn:  types.MerlinDenomination,
										TokenOut: types.AtomDenomination,
									},
								},
							},
						},
						TokenIn:  types.MerlinDenomination,
						TokenOut: "Juno",
					},
				},
			},
			valid: false,
		},
		{
			description: "Default parameters with invalid routes (mismatch in and out denoms)",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				TokenPairs: []types.TokenPairArbRoutes{
					{
						ArbRoutes: []*types.Route{{
							Trades: []*types.Trade{
								{
									Pool:     1,
									TokenIn:  types.AtomDenomination,
									TokenOut: "Juno",
								},
								{
									Pool:     0,
									TokenIn:  "Juno",
									TokenOut: types.MerlinDenomination,
								},
								{
									Pool:     3,
									TokenIn:  types.MerlinDenomination,
									TokenOut: "eth",
								},
							},
						}},
						TokenIn:  types.MerlinDenomination,
						TokenOut: "Juno",
					},
				},
			},
			valid: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			err := tc.genState.Validate()

			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
