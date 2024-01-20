# Burn Blocks

```go
burnBlocksController := client.BurnBlocksController()
```

## Class Name

`BurnBlocksController`

## Methods

- [Get Burn Blocks](../../doc/controllers/burn-blocks.md#get-burn-blocks)
- [Get Burn Block](../../doc/controllers/burn-blocks.md#get-burn-block)

# Get Burn Blocks

Retrieves a list of recent burn blocks

```go
GetBurnBlocks(
    ctx context.Context,
    limit *int,
    offset *int) (
    models.ApiResponse[models.BurnBlockListResponse],
    error)
```

## Parameters

| Parameter | Type   | Tags            | Description                        |
| --------- | ------ | --------------- | ---------------------------------- |
| `limit`   | `*int` | Query, Optional | max number of burn blocks to fetch |
| `offset`  | `*int` | Query, Optional | index of first burn block to fetch |

## Response Type

[`models.BurnBlockListResponse`](../../doc/models/burn-block-list-response.md)

## Example Usage

```go
ctx := context.Background()
limit := 20
offset := 42000

apiResponse, err := burnBlocksController.GetBurnBlocks(ctx, &limit, &offset)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Burn Block

Retrieves a single burn block

```go
GetBurnBlock(
    ctx context.Context,
    heightOrHash interface{}) (
    models.ApiResponse[models.BurnBlock],
    error)
```

## Parameters

| Parameter      | Type          | Tags               | Description                                                                                          |
| -------------- | ------------- | ------------------ | ---------------------------------------------------------------------------------------------------- |
| `heightOrHash` | `interface{}` | Template, Required | filter by burn block height, hash, or the constant `latest` to filter for the most recent burn block |

## Response Type

[`models.BurnBlock`](../../doc/models/burn-block.md)

## Example Usage

```go
ctx := context.Background()
heightOrHash := interface{}("[key1, val1][key2, val2]")

apiResponse, err := burnBlocksController.GetBurnBlock(ctx, heightOrHash)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```
