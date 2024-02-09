# go-stacks

## Introduction

`go-stacks` is a golang sdk for interacting with the stacks blockchain.
This SDK was generated using the Hiro's OpenAPI specification & [API Matic](https://www.apimatic.io/) with some modifications. It only has one dependancy, and that dependancy has none.

### Requirements

The SDK requires **Go version 1.22 or above**.

## Building

### Install Dependencies

Resolve all the SDK dependencies, using the `go get` command.

## Test the SDK

`Go Testing` is used as the testing framework. To run the unit tests of the SDK, navigate to the `tests` directory of the SDK and run the following command in the terminal:

```bash
$ go test
```

## Quick Start

**_Note:_** Documentation for the client can be found [here.](doc/client.md)

```go
import stacks "github.com/cbadawi/go-stacks"

func main() {
	config := stacks.CreateConfiguration(
		stacks.WithBaseUri(stacks.MAINNET_URI),
	)

	principal := "SP18W9NPMBEBKWRJA5RNRJKGKW1FHQQ8BF9RXVEGN"
	unanchored := false
	untilBlock := "600000"

	client := stacks.NewClient(config)
	accountsController := client.AccountsController()
	ctx := context.Background()

	apiResponse, err := accountsController.GetAccountBalance(ctx, principal, &unanchored, &untilBlock)
	fmt.Println(apiResponse)
	fmt.Println(err)
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
