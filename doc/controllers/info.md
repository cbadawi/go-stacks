# Info

Read-only endpoints to obtain network, Proof-of-Transfer, Stacking, STX token, and node information

```go
infoController := client.InfoController()
```

## Class Name

`InfoController`

## Methods

- [Get Core Api Info](../../doc/controllers/info.md#get-core-api-info)
- [Get Status](../../doc/controllers/info.md#get-status)
- [Get Network Block Times](../../doc/controllers/info.md#get-network-block-times)
- [Get Network Block Time by Network](../../doc/controllers/info.md#get-network-block-time-by-network)
- [Get Stx Supply](../../doc/controllers/info.md#get-stx-supply)
- [Get Stx Supply Total Supply Plain](../../doc/controllers/info.md#get-stx-supply-total-supply-plain)
- [Get Stx Supply Circulating Plain](../../doc/controllers/info.md#get-stx-supply-circulating-plain)
- [Get Total Stx Supply Legacy Format](../../doc/controllers/info.md#get-total-stx-supply-legacy-format)
- [Get Pox Info](../../doc/controllers/info.md#get-pox-info)

# Get Core Api Info

Retrieves information about the Core API including the server version

```go
GetCoreApiInfo(
    ctx context.Context) (
    models.ApiResponse[models.CoreNodeInfoResponse],
    error)
```

## Response Type

[`models.CoreNodeInfoResponse`](../../doc/models/core-node-info-response.md)

## Example Usage

```go
ctx := context.Background()
apiResponse, err := infoController.GetCoreApiInfo(ctx)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Status

Retrieves the running status of the Stacks Blockchain API, including the server version and current chain tip information.

```go
GetStatus(
    ctx context.Context) (
    models.ApiResponse[models.ServerStatusResponse],
    error)
```

## Response Type

[`models.ServerStatusResponse`](../../doc/models/server-status-response.md)

## Example Usage

```go
ctx := context.Background()
apiResponse, err := infoController.GetStatus(ctx)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Network Block Times

Retrieves the target block times for mainnet and testnet. The block time is hardcoded and will change throughout the implementation phases of the testnet.

```go
GetNetworkBlockTimes(
    ctx context.Context) (
    models.ApiResponse[models.NetworkBlockTimesResponse],
    error)
```

## Response Type

[`models.NetworkBlockTimesResponse`](../../doc/models/network-block-times-response.md)

## Example Usage

```go
ctx := context.Background()
apiResponse, err := infoController.GetNetworkBlockTimes(ctx)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Network Block Time by Network

Retrieves the target block time for a given network. The network can be mainnet or testnet. The block time is hardcoded and will change throughout the implementation phases of the testnet.

```go
GetNetworkBlockTimeByNetwork(
    ctx context.Context,
    network models.NetworkEnum) (
    models.ApiResponse[models.TargetBlockTime],
    error)
```

## Parameters

| Parameter | Type                                                     | Tags               | Description                                                   |
| --------- | -------------------------------------------------------- | ------------------ | ------------------------------------------------------------- |
| `network` | [`models.NetworkEnum`](../../doc/models/network-enum.md) | Template, Required | the target block time for a given network (testnet, mainnet). |

## Response Type

[`models.TargetBlockTime`](../../doc/models/target-block-time.md)

## Example Usage

```go
ctx := context.Background()
network := models.NetworkEnum("mainnet")

apiResponse, err := infoController.GetNetworkBlockTimeByNetwork(ctx, network)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Stx Supply

Retrieves the total and unlocked STX supply. More information on Stacking can be found [here] (https://docs.stacks.co/understand-stacks/stacking).
**Note:** This uses the estimated future total supply for the year 2050.

```go
GetStxSupply(
    ctx context.Context,
    height *float64) (
    models.ApiResponse[models.GetStxSupplyResponse],
    error)
```

## Parameters

| Parameter | Type       | Tags            | Description                                                                                                                                                                                                         |
| --------- | ---------- | --------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `height`  | `*float64` | Query, Optional | Supply details are queried from specified block height. If the block height is not specified, the latest block height is taken as default value. Note that the `block height` is referred to the stacks blockchain. |

## Response Type

[`models.GetStxSupplyResponse`](../../doc/models/get-stx-supply-response.md)

## Example Usage

```go
ctx := context.Background()
height := float64(200)

apiResponse, err := infoController.GetStxSupply(ctx, &height)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Stx Supply Total Supply Plain

Retrieves the total supply for STX tokens as plain text.
**Note:** this uses the estimated future total supply for the year 2050.

```go
GetStxSupplyTotalSupplyPlain(
    ctx context.Context) (
    models.ApiResponse[string],
    error)
```

## Response Type

`string`

## Example Usage

```go
ctx := context.Background()
apiResponse, err := infoController.GetStxSupplyTotalSupplyPlain(ctx)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response

```
"123.456789"
```

# Get Stx Supply Circulating Plain

Retrieves the STX tokens currently in circulation that have been unlocked as plain text.

```go
GetStxSupplyCirculatingPlain(
    ctx context.Context) (
    models.ApiResponse[string],
    error)
```

## Response Type

`string`

## Example Usage

```go
ctx := context.Background()
apiResponse, err := infoController.GetStxSupplyCirculatingPlain(ctx)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response

```
"123.456789"
```

# Get Total Stx Supply Legacy Format

Retrieves total supply of STX tokens including those currently in circulation that have been unlocked.
**Note:** this uses the estimated future total supply for the year 2050.

```go
GetTotalStxSupplyLegacyFormat(
    ctx context.Context,
    height *float64) (
    models.ApiResponse[models.GetStxSupplyLegacyFormatResponse],
    error)
```

## Parameters

| Parameter | Type       | Tags            | Description                                                                                                                                      |
| --------- | ---------- | --------------- | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `height`  | `*float64` | Query, Optional | Supply details are queried from specified block height. If the block height is not specified, the latest block height is taken as default value. |

## Response Type

[`models.GetStxSupplyLegacyFormatResponse`](../../doc/models/get-stx-supply-legacy-format-response.md)

## Example Usage

```go
ctx := context.Background()
height := float64(200)

apiResponse, err := infoController.GetTotalStxSupplyLegacyFormat(ctx, &height)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Pox Info

Retrieves Proof-of-Transfer (PoX) information. Can be used for Stacking.

```go
GetPoxInfo(
    ctx context.Context) (
    models.ApiResponse[models.CoreNodePoxResponse],
    error)
```

## Response Type

[`models.CoreNodePoxResponse`](../../doc/models/core-node-pox-response.md)

## Example Usage

```go
ctx := context.Background()
apiResponse, err := infoController.GetPoxInfo(ctx)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```
