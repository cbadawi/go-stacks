# Stacking Rewards

```go
stackingRewardsController := client.StackingRewardsController()
```

## Class Name

`StackingRewardsController`

## Methods

- [Get Burnchain Reward Slot Holders](../../doc/controllers/stacking-rewards.md#get-burnchain-reward-slot-holders)
- [Get Burnchain Reward Slot Holders by Address](../../doc/controllers/stacking-rewards.md#get-burnchain-reward-slot-holders-by-address)
- [Get Burnchain Reward List](../../doc/controllers/stacking-rewards.md#get-burnchain-reward-list)
- [Get Burnchain Reward List by Address](../../doc/controllers/stacking-rewards.md#get-burnchain-reward-list-by-address)
- [Get Burnchain Rewards Total by Address](../../doc/controllers/stacking-rewards.md#get-burnchain-rewards-total-by-address)

# Get Burnchain Reward Slot Holders

Retrieves a list of the Bitcoin addresses that would validly receive Proof-of-Transfer commitments.

```go
GetBurnchainRewardSlotHolders(
    ctx context.Context,
    limit *int,
    offset *int) (
    models.ApiResponse[models.BurnchainRewardSlotHolderListResponse],
    error)
```

## Parameters

| Parameter | Type   | Tags            | Description                       |
| --------- | ------ | --------------- | --------------------------------- |
| `limit`   | `*int` | Query, Optional | max number of items to fetch      |
| `offset`  | `*int` | Query, Optional | index of the first items to fetch |

## Response Type

[`models.BurnchainRewardSlotHolderListResponse`](../../doc/models/burnchain-reward-slot-holder-list-response.md)

## Example Usage

```go
ctx := context.Background()
limit := 96
offset := 42000

apiResponse, err := stackingRewardsController.GetBurnchainRewardSlotHolders(ctx, &limit, &offset)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Burnchain Reward Slot Holders by Address

Retrieves a list of the Bitcoin addresses that would validly receive Proof-of-Transfer commitments for a given reward slot holder recipient address.

```go
GetBurnchainRewardSlotHoldersByAddress(
    ctx context.Context,
    address string,
    limit *int,
    offset *int) (
    models.ApiResponse[models.BurnchainRewardSlotHolderListResponse],
    error)
```

## Parameters

| Parameter | Type     | Tags               | Description                                                                                                                                                                                                          |
| --------- | -------- | ------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `address` | `string` | Template, Required | Reward slot holder recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format |
| `limit`   | `*int`   | Query, Optional    | max number of items to fetch                                                                                                                                                                                         |
| `offset`  | `*int`   | Query, Optional    | index of the first items to fetch                                                                                                                                                                                    |

## Response Type

[`models.BurnchainRewardSlotHolderListResponse`](../../doc/models/burnchain-reward-slot-holder-list-response.md)

## Example Usage

```go
ctx := context.Background()
address := "36hQtSEXBMevo5chpxhfAGiCTSC34QKgda"
offset := 42000

apiResponse, err := stackingRewardsController.GetBurnchainRewardSlotHoldersByAddress(ctx, address, nil, &offset)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Burnchain Reward List

Retrieves a list of recent burnchain (e.g. Bitcoin) reward recipients with the associated amounts and block info

```go
GetBurnchainRewardList(
    ctx context.Context,
    limit *int,
    offset *int) (
    models.ApiResponse[models.BurnchainRewardListResponse],
    error)
```

## Parameters

| Parameter | Type   | Tags            | Description                     |
| --------- | ------ | --------------- | ------------------------------- |
| `limit`   | `*int` | Query, Optional | max number of rewards to fetch  |
| `offset`  | `*int` | Query, Optional | index of first rewards to fetch |

## Response Type

[`models.BurnchainRewardListResponse`](../../doc/models/burnchain-reward-list-response.md)

## Example Usage

```go
ctx := context.Background()
limit := 96
offset := 42000

apiResponse, err := stackingRewardsController.GetBurnchainRewardList(ctx, &limit, &offset)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Burnchain Reward List by Address

Retrieves a list of recent burnchain (e.g. Bitcoin) rewards for the given recipient with the associated amounts and block info

```go
GetBurnchainRewardListByAddress(
    ctx context.Context,
    address string,
    limit *int,
    offset *int) (
    models.ApiResponse[models.BurnchainRewardListResponse],
    error)
```

## Parameters

| Parameter | Type     | Tags               | Description                                                                                                                                                                                              |
| --------- | -------- | ------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `address` | `string` | Template, Required | Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format |
| `limit`   | `*int`   | Query, Optional    | max number of rewards to fetch                                                                                                                                                                           |
| `offset`  | `*int`   | Query, Optional    | index of first rewards to fetch                                                                                                                                                                          |

## Response Type

[`models.BurnchainRewardListResponse`](../../doc/models/burnchain-reward-list-response.md)

## Example Usage

```go
ctx := context.Background()
address := "36hQtSEXBMevo5chpxhfAGiCTSC34QKgda"
offset := 42000

apiResponse, err := stackingRewardsController.GetBurnchainRewardListByAddress(ctx, address, nil, &offset)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Burnchain Rewards Total by Address

Retrieves the total burnchain (e.g. Bitcoin) rewards for a given recipient `address`

```go
GetBurnchainRewardsTotalByAddress(
    ctx context.Context,
    address string) (
    models.ApiResponse[models.BurnchainRewardsTotal],
    error)
```

## Parameters

| Parameter | Type     | Tags               | Description                                                                                                                                                                                              |
| --------- | -------- | ------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `address` | `string` | Template, Required | Reward recipient address. Should either be in the native burnchain's format (e.g. B58 for Bitcoin), or if a STX principal address is provided it will be encoded as into the equivalent burnchain format |

## Response Type

[`models.BurnchainRewardsTotal`](../../doc/models/burnchain-rewards-total.md)

## Example Usage

```go
ctx := context.Background()
address := "36hQtSEXBMevo5chpxhfAGiCTSC34QKgda"

apiResponse, err := stackingRewardsController.GetBurnchainRewardsTotalByAddress(ctx, address)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```
