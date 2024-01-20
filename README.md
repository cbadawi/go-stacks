# Getting Started with Stacks Blockchain API

## Introduction

Welcome to the API reference overview for the [Stacks Blockchain API](https://docs.hiro.so/get-started/stacks-blockchain-api).

[Download Postman collection](https://hirosystems.github.io/stacks-blockchain-api/collection.json)

### Requirements

The SDK requires **Go version 1.18 or above**.

## Building

### Install Dependencies

Resolve all the SDK dependencies, using the `go get` command.

## Installation

The following section explains how to use the stacksblockchainapi library in a new project.

### 1. Add SDK as a Dependency to the Application

- Add the following lines to your application's `go.mod` file:

```go
replace stacksblockchainapi => ".\\Stacks Blockchain API" // local path to the SDK

require stacksblockchainapi v0.0.0
```

- Resolve the dependencies in the updated `go.mod` file, using the `go get` command.

## Test the SDK

`Go Testing` is used as the testing framework. To run the unit tests of the SDK, navigate to the `tests` directory of the SDK and run the following command in the terminal:

```bash
$ go test
```

## Initialize the API Client

**_Note:_** Documentation for the client can be found [here.](doc/client.md)

The following parameters are configurable for the API Client:

| Parameter           | Type                                             | Description                                                     |
| ------------------- | ------------------------------------------------ | --------------------------------------------------------------- |
| `environment`       | Environment                                      | The API environment. <br> **Default: `Environment.PRODUCTION`** |
| `httpConfiguration` | [`HttpConfiguration`](doc/http-configuration.md) | Configurable http client options like timeout and retries.      |

The API client can be initialized as follows:

```go
config := stacksblockchainapi.CreateConfiguration(
    stacksblockchainapi.WithHttpConfiguration(
        stacksblockchainapi.CreateHttpConfiguration(
            stacksblockchainapi.WithTimeout(0),
            stacksblockchainapi.WithTransport(http.DefaultTransport),
            stacksblockchainapi.WithRetryConfiguration(
                stacksblockchainapi.CreateRetryConfiguration(
                    stacksblockchainapi.WithMaxRetryAttempts(0),
                    stacksblockchainapi.WithRetryOnTimeout(true),
                    stacksblockchainapi.WithRetryInterval(1),
                    stacksblockchainapi.WithMaximumRetryWaitTime(0),
                    stacksblockchainapi.WithBackoffFactor(2),
                    stacksblockchainapi.WithHttpStatusCodesToRetry([]int64{408, 413, 429, 500, 502, 503, 504, 521, 522, 524}),
                    stacksblockchainapi.WithHttpMethodsToRetry([]string{"GET", "PUT"}),
                ),
            ),
        ),
    ),
    stacksblockchainapi.WithEnvironment(stacksblockchainapi.PRODUCTION),
)
client := stacksblockchainapi.NewClient(config)
```

## Environments

The SDK can be configured to use a different environment for making API calls. Available environments are:

### Fields

| Name         | Description         |
| ------------ | ------------------- |
| production   | **Default** Mainnet |
| environment2 | Testnet             |
| environment3 | Local               |

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
