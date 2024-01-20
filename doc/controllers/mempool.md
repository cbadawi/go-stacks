# Mempool

Endpoints to obtain Mempool information

```go
mempoolController := client.MempoolController()
```

## Class Name

`MempoolController`

# Get Mempool Fee Priorities

Returns estimated fee priorities (in micro-STX) for all transactions that are currently in the mempool. Also returns priorities separated by transaction type.

```go
GetMempoolFeePriorities(
    ctx context.Context) (
    models.ApiResponse[models.MempoolFeePriorities],
    error)
```

## Response Type

[`models.MempoolFeePriorities`](../../doc/models/mempool-fee-priorities.md)

## Example Usage

```go
ctx := context.Background()
apiResponse, err := mempoolController.GetMempoolFeePriorities(ctx)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```
