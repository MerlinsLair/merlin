# CosmWasm support

This package contains CosmWasm integration points.

This package provides first class support for:

- Queries
  - Denoms
  - Pools
  - Prices
- Messages / Execution
  - Minting / controlling of new native tokens
  - Swap

## Command line interface (CLI)

- Commands

```sh
  merlind tx wasm -h
```

- Query

```sh
  merlind query wasm -h
```

## Tests

This contains a few high level tests that `x/wasm` is properly
integrated.

Since the code tested is not in this repo, and we are just testing the
application integration (app.go), I figured this is the most suitable
location for it.
