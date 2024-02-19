# go-stacks

## Introduction

`go-stacks` is a golang sdk for interacting with the stacks blockchain.
This SDK was generated using the Hiro's OpenAPI specification & [API Matic](https://www.apimatic.io/) with some modifications. It only has one dependancy, and that dependancy has none.

### Requirements

The SDK requires **Go version 1.22 or above**.

## Building

### Install Dependencies

Resolve all the SDK dependencies, using the `go get` command.

## TODOs

- [ ] fix and improve testing
- [ ] implement (de)serializing transactions

## Quick Start

**_Note:_** Documentation for the client can be found [here.](doc/client.md)

```go
package main

import (
	"context"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks"
)

func main() {
	config := stacks.CreateConfiguration(
		stacks.WithBaseUri(stacks.MAINNET_URI),
		stacks.WithVerbose(true),
	)

	principal := "SP3KANBW2C4E5BRPWNTWZCCDGF2F87CW9D9KV0FFK"
	unanchored := false
	untilBlock := "600000"

	client := stacks.NewClient(config)
	accountsController := client.AccountsController()
	ctx := context.Background()

	apiResponse, err := accountsController.GetAccountBalance(ctx, principal, &unanchored, &untilBlock)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(apiResponse.Data.Stx.Balance)
}

```

## List of APIs

- [Burn Blocks](doc/controllers/burn-blocks.md)
- [Non-Fungible Tokens](doc/controllers/non-fungible-tokens.md)
- [Smart Contracts](doc/controllers/smart-contracts.md)
- [Stacking Rewards](doc/controllers/stacking-rewards.md)
- [Accounts](doc/controllers/accounts.md)
- [Blocks](doc/controllers/blocks.md)
- [Faucets](doc/controllers/faucets.md)
- [Fees](doc/controllers/fees.md)
- [Info](doc/controllers/info.md)
- [Microblocks](doc/controllers/microblocks.md)
- [Names](doc/controllers/names.md)
- [Rosetta](doc/controllers/rosetta.md)
- [Search](doc/controllers/search.md)
- [Transactions](doc/controllers/transactions.md)
- [Mempool](doc/controllers/mempool.md)
- [Stacking](doc/controllers/stacking.md)

## Classes Documentation

- [HttpConfiguration](doc/http-configuration.md)
- [RetryConfiguration](doc/retry-configuration.md)
