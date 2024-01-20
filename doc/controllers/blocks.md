# Blocks

Read-only endpoints to obtain Stacks block details

```go
blocksController := client.BlocksController()
```

## Class Name

`BlocksController`

## Methods

- [Get Blocks by Burn Block](../../doc/controllers/blocks.md#get-blocks-by-burn-block)
- [Get Blocks](../../doc/controllers/blocks.md#get-blocks)
- [Get Block](../../doc/controllers/blocks.md#get-block)
- [Get Block List](../../doc/controllers/blocks.md#get-block-list)
- [Get Block by Hash](../../doc/controllers/blocks.md#get-block-by-hash)
- [Get Block by Height](../../doc/controllers/blocks.md#get-block-by-height)
- [Get Block by Burn Block Hash](../../doc/controllers/blocks.md#get-block-by-burn-block-hash)
- [Get Block by Burn Block Height](../../doc/controllers/blocks.md#get-block-by-burn-block-height)

# Get Blocks by Burn Block

Retrieves a list of blocks confirmed by a specific burn block

```go
GetBlocksByBurnBlock(
    ctx context.Context,
    heightOrHash interface{},
    limit *int,
    offset *int) (
    models.ApiResponse[models.NakamotoBlockListResponse],
    error)
```

## Parameters

| Parameter      | Type          | Tags               | Description                                                                                          |
| -------------- | ------------- | ------------------ | ---------------------------------------------------------------------------------------------------- |
| `heightOrHash` | `interface{}` | Template, Required | filter by burn block height, hash, or the constant `latest` to filter for the most recent burn block |
| `limit`        | `*int`        | Query, Optional    | max number of blocks to fetch                                                                        |
| `offset`       | `*int`        | Query, Optional    | index of first burn block to fetch                                                                   |

## Response Type

[`models.NakamotoBlockListResponse`](../../doc/models/nakamoto-block-list-response.md)

## Example Usage

```go
ctx := context.Background()
heightOrHash := interface{}("[key1, val1][key2, val2]")
limit := 20
offset := 0

apiResponse, err := blocksController.GetBlocksByBurnBlock(ctx, heightOrHash, &limit, &offset)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Blocks

Retrieves a list of recently mined blocks

```go
GetBlocks(
    ctx context.Context,
    limit *int,
    offset *int) (
    models.ApiResponse[models.NakamotoBlockListResponse],
    error)
```

## Parameters

| Parameter | Type   | Tags            | Description                        |
| --------- | ------ | --------------- | ---------------------------------- |
| `limit`   | `*int` | Query, Optional | max number of blocks to fetch      |
| `offset`  | `*int` | Query, Optional | index of first burn block to fetch |

## Response Type

[`models.NakamotoBlockListResponse`](../../doc/models/nakamoto-block-list-response.md)

## Example Usage

```go
ctx := context.Background()
limit := 20
offset := 0

apiResponse, err := blocksController.GetBlocks(ctx, &limit, &offset)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Block

Retrieves a single block

```go
GetBlock(
    ctx context.Context,
    heightOrHash interface{}) (
    models.ApiResponse[models.NakamotoBlock],
    error)
```

## Parameters

| Parameter      | Type          | Tags               | Description                                                                                                 |
| -------------- | ------------- | ------------------ | ----------------------------------------------------------------------------------------------------------- |
| `heightOrHash` | `interface{}` | Template, Required | filter by block height, hash, index block hash or the constant `latest` to filter for the most recent block |

## Response Type

[`models.NakamotoBlock`](../../doc/models/nakamoto-block.md)

## Example Usage

```go
ctx := context.Background()
heightOrHash := interface{}("[key1, val1][key2, val2]")

apiResponse, err := blocksController.GetBlock(ctx, heightOrHash)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Block List

**This endpoint is deprecated.**

**NOTE:** This endpoint is deprecated in favor of [Get blocks](#operation/get_blocks).

Retrieves a list of recently mined blocks

If you need to actively monitor new blocks, we highly recommend subscribing to [WebSockets or Socket.io](https://github.com/hirosystems/stacks-blockchain-api/tree/master/client) for real-time updates.

```go
GetBlockList(
    ctx context.Context,
    limit *int,
    offset *int) (
    models.ApiResponse[models.BlockListResponse],
    error)
```

## Parameters

| Parameter | Type   | Tags            | Description                   |
| --------- | ------ | --------------- | ----------------------------- |
| `limit`   | `*int` | Query, Optional | max number of blocks to fetch |
| `offset`  | `*int` | Query, Optional | index of first block to fetch |

## Response Type

[`models.BlockListResponse`](../../doc/models/block-list-response.md)

## Example Usage

```go
ctx := context.Background()
limit := 20
offset := 42000

apiResponse, err := blocksController.GetBlockList(ctx, &limit, &offset)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Block by Hash

**This endpoint is deprecated.**

**NOTE:** This endpoint is deprecated in favor of [Get block](#operation/get_block).

Retrieves block details of a specific block for a given chain height. You can use the hash from your latest block ('get_block_list' API) to get your block details.

```go
GetBlockByHash(
    ctx context.Context,
    hash string) (
    models.ApiResponse[models.Block],
    error)
```

## Parameters

| Parameter | Type     | Tags               | Description       |
| --------- | -------- | ------------------ | ----------------- |
| `hash`    | `string` | Template, Required | Hash of the block |

## Response Type

[`models.Block`](../../doc/models/block.md)

## Example Usage

```go
ctx := context.Background()
hash := "0x4839a8b01cfb39ffcc0d07d3db31e848d5adf5279d529ed5062300b9f353ff79"

apiResponse, err := blocksController.GetBlockByHash(ctx, hash)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description               | Exception Class |
| ---------------- | ------------------------------- | --------------- |
| 404              | Cannot find block with given ID | `ApiError`      |

# Get Block by Height

**This endpoint is deprecated.**

**NOTE:** This endpoint is deprecated in favor of [Get block](#operation/get_block).

Retrieves block details of a specific block at a given block height

```go
GetBlockByHeight(
    ctx context.Context,
    height float64) (
    models.ApiResponse[models.Block],
    error)
```

## Parameters

| Parameter | Type      | Tags               | Description         |
| --------- | --------- | ------------------ | ------------------- |
| `height`  | `float64` | Template, Required | Height of the block |

## Response Type

[`models.Block`](../../doc/models/block.md)

## Example Usage

```go
ctx := context.Background()
height := float64(10000)

apiResponse, err := blocksController.GetBlockByHeight(ctx, height)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description                   | Exception Class |
| ---------------- | ----------------------------------- | --------------- |
| 404              | Cannot find block with given height | `ApiError`      |

# Get Block by Burn Block Hash

**This endpoint is deprecated.**

**NOTE:** This endpoint is deprecated in favor of [Get blocks](#operation/get_blocks).

Retrieves block details of a specific block for a given burnchain block hash

```go
GetBlockByBurnBlockHash(
    ctx context.Context,
    burnBlockHash string) (
    models.ApiResponse[models.Block],
    error)
```

## Parameters

| Parameter       | Type     | Tags               | Description                 |
| --------------- | -------- | ------------------ | --------------------------- |
| `burnBlockHash` | `string` | Template, Required | Hash of the burnchain block |

## Response Type

[`models.Block`](../../doc/models/block.md)

## Example Usage

```go
ctx := context.Background()
burnBlockHash := "0x00000000000000000002bba732926cf68b6eda3e2cdbc2a85af79f10efeeeb10"

apiResponse, err := blocksController.GetBlockByBurnBlockHash(ctx, burnBlockHash)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description                   | Exception Class |
| ---------------- | ----------------------------------- | --------------- |
| 404              | Cannot find block with given height | `ApiError`      |

# Get Block by Burn Block Height

**This endpoint is deprecated.**

**NOTE:** This endpoint is deprecated in favor of [Get blocks](#operation/get_blocks).

Retrieves block details of a specific block for a given burn chain height

```go
GetBlockByBurnBlockHeight(
    ctx context.Context,
    burnBlockHeight float64) (
    models.ApiResponse[models.Block],
    error)
```

## Parameters

| Parameter         | Type      | Tags               | Description                    |
| ----------------- | --------- | ------------------ | ------------------------------ |
| `burnBlockHeight` | `float64` | Template, Required | Height of the burn chain block |

## Response Type

[`models.Block`](../../doc/models/block.md)

## Example Usage

```go
ctx := context.Background()
burnBlockHeight := float64(744603)

apiResponse, err := blocksController.GetBlockByBurnBlockHeight(ctx, burnBlockHeight)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description                   | Exception Class |
| ---------------- | ----------------------------------- | --------------- |
| 404              | Cannot find block with given height | `ApiError`      |
