# Search

Read-only endpoints to search for accounts, blocks, smart contracts, and transactions

```go
searchController := client.SearchController()
```

## Class Name

`SearchController`

# Search by Id

Search blocks, transactions, contracts, or accounts by hash/ID

```go
SearchById(
    ctx context.Context,
    id string,
    includeMetadata *bool) (
    models.ApiResponse[models.SearchResult],
    error)
```

## Parameters

| Parameter         | Type     | Tags               | Description                                                                          |
| ----------------- | -------- | ------------------ | ------------------------------------------------------------------------------------ |
| `id`              | `string` | Template, Required | The hex hash string for a block or transaction, account address, or contract address |
| `includeMetadata` | `*bool`  | Query, Optional    | This includes the detailed data for purticular hash in the response                  |

## Response Type

[`models.SearchResult`](../../doc/models/search-result.md)

## Example Usage

```go
ctx := context.Background()
id := "0xcf8b233f19f6c07d2dc1963302d2436efd36e9afac127bf6582824a13961c06d"

apiResponse, err := searchController.SearchById(ctx, id, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
| ---------------- | ----------------- | --------------- |
| 404              | Not found         | `ApiError`      |
