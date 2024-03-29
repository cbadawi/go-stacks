# Client Class Documentation

The following parameters are configurable for the API Client:

| Parameter           | Type                                         | Description                                                     |
| ------------------- | -------------------------------------------- | --------------------------------------------------------------- |
| `environment`       | Environment                                  | The API environment. <br> **Default: `Environment.PRODUCTION`** |
| `httpConfiguration` | [`HttpConfiguration`](http-configuration.md) | Configurable http client options like timeout and retries.      |

The API client can be initialized as follows:

```go
config := stacks.CreateConfiguration(
    stacks.WithHttpConfiguration(
        stacks.CreateHttpConfiguration(
            stacks.WithTimeout(0),
            stacks.WithTransport(http.DefaultTransport),
            stacks.WithRetryConfiguration(
                stacks.CreateRetryConfiguration(
                    stacks.WithMaxRetryAttempts(0),
                    stacks.WithRetryOnTimeout(true),
                    stacks.WithRetryInterval(1),
                    stacks.WithMaximumRetryWaitTime(0),
                    stacks.WithBackoffFactor(2),
                    stacks.WithHttpStatusCodesToRetry([]int64{408, 413, 429, 500, 502, 503, 504, 521, 522, 524}),
                    stacks.WithHttpMethodsToRetry([]string{"GET", "PUT"}),
                ),
            ),
        ),
    ),
    stacks.WithBaseUri(stacks.MAINNET_URI),
    stacks.WithVerbose(true),

)
client := stacks.NewClient(config)
```

## Stacks Blockchain API Client

The gateway for the SDK. This class acts as a factory for the Controllers and also holds the configuration of the SDK.

## Controllers

| Name              | Description                      |
| ----------------- | -------------------------------- |
| accounts          | Gets AccountsController          |
| blocks            | Gets BlocksController            |
| burnBlocks        | Gets BurnBlocksController        |
| faucets           | Gets FaucetsController           |
| fees              | Gets FeesController              |
| info              | Gets InfoController              |
| microblocks       | Gets MicroblocksController       |
| names             | Gets NamesController             |
| nonFungibleTokens | Gets NonFungibleTokensController |
| rosetta           | Gets RosettaController           |
| search            | Gets SearchController            |
| smartContracts    | Gets SmartContractsController    |
| stackingRewards   | Gets StackingRewardsController   |
| transactions      | Gets TransactionsController      |
| mempool           | Gets MempoolController           |
| stacking          | Gets StackingController          |
