# Microblocks

Read-only endpoints to obtain microblocks details

Stacks Documentation - Microblocks: [https://docs.stacks.co/understand-stacks/microblocks](https://docs.stacks.co/understand-stacks/microblocks)

```go
microblocksController := client.MicroblocksController()
```

## Class Name

`MicroblocksController`

## Methods

- [Get Microblock List](../../doc/controllers/microblocks.md#get-microblock-list)
- [Get Microblock by Hash](../../doc/controllers/microblocks.md#get-microblock-by-hash)
- [Get Unanchored Txs](../../doc/controllers/microblocks.md#get-unanchored-txs)

# Get Microblock List

Retrieves a list of microblocks.

If you need to actively monitor new microblocks, we highly recommend subscribing to [WebSockets or Socket.io](https://github.com/hirosystems/stacks-blockchain-api/tree/master/client) for real-time updates.

```go
GetMicroblockList(
    ctx context.Context,
    limit *int,
    offset *int) (
    models.ApiResponse[models.MicroblockListResponse],
    error)
```

## Parameters

| Parameter | Type   | Tags            | Description                            |
| --------- | ------ | --------------- | -------------------------------------- |
| `limit`   | `*int` | Query, Optional | Max number of microblocks to fetch     |
| `offset`  | `*int` | Query, Optional | Index of the first microblock to fetch |

## Response Type

[`models.MicroblockListResponse`](../../doc/models/microblock-list-response.md)

## Example Usage

```go
ctx := context.Background()
limit := 100
offset := 42000

apiResponse, err := microblocksController.GetMicroblockList(ctx, &limit, &offset)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Microblock by Hash

Retrieves a specific microblock by `hash`

```go
GetMicroblockByHash(
    ctx context.Context,
    hash string) (
    models.ApiResponse[models.Microblock],
    error)
```

## Parameters

| Parameter | Type     | Tags               | Description            |
| --------- | -------- | ------------------ | ---------------------- |
| `hash`    | `string` | Template, Required | Hash of the microblock |

## Response Type

[`models.Microblock`](../../doc/models/microblock.md)

## Example Usage

```go
ctx := context.Background()
hash := "0x3bfcdf84b3012adb544cf0f6df4835f93418c2269a3881885e27b3d58eb82d47"

apiResponse, err := microblocksController.GetMicroblockByHash(ctx, hash)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description                      | Exception Class |
| ---------------- | -------------------------------------- | --------------- |
| 404              | Cannot find microblock with given hash | `ApiError`      |

# Get Unanchored Txs

Retrieves transactions that have been streamed in microblocks but not yet accepted or rejected in an anchor block

```go
GetUnanchoredTxs(
    ctx context.Context) (
    models.ApiResponse[models.UnanchoredTransactionListResponse],
    error)
```

## Response Type

[`models.UnanchoredTransactionListResponse`](../../doc/models/unanchored-transaction-list-response.md)

## Example Usage

```go
ctx := context.Background()
apiResponse, err := microblocksController.GetUnanchoredTxs(ctx)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```
