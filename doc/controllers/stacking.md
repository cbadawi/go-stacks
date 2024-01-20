# Stacking

```go
stackingController := client.StackingController()
```

## Class Name

`StackingController`

# Get Pool Delegations

Retrieves the list of stacking pool members for a given delegator principal.

```go
GetPoolDelegations(
    ctx context.Context,
    poolPrincipal string,
    afterBlock *int,
    unanchored *bool,
    limit *int,
    offset *int) (
    models.ApiResponse[models.PoolDelegationsResponse],
    error)
```

## Parameters

| Parameter       | Type     | Tags               | Description                                                                 |
| --------------- | -------- | ------------------ | --------------------------------------------------------------------------- |
| `poolPrincipal` | `string` | Template, Required | Address principal of the stacking pool delegator                            |
| `afterBlock`    | `*int`   | Query, Optional    | If specified, only delegation events after the given block will be included |
| `unanchored`    | `*bool`  | Query, Optional    | whether or not to include Stackers from unconfirmed transactions            |
| `limit`         | `*int`   | Query, Optional    | number of items to return                                                   |
| `offset`        | `*int`   | Query, Optional    | number of items to skip                                                     |

## Response Type

[`models.PoolDelegationsResponse`](../../doc/models/pool-delegations-response.md)

## Example Usage

```go
ctx := context.Background()
poolPrincipal := "SPSCWDV3RKV5ZRN1FQD84YE1NQFEDJ9R1F4DYQ11"
unanchored := true
limit := 100
offset := 300

apiResponse, err := stackingController.GetPoolDelegations(ctx, poolPrincipal, nil, &unanchored, &limit, &offset)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```
