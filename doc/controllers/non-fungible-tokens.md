# Non-Fungible Tokens

```go
nonFungibleTokensController := client.NonFungibleTokensController()
```

## Class Name

`NonFungibleTokensController`

## Methods

- [Get Nft Holdings](../../doc/controllers/non-fungible-tokens.md#get-nft-holdings)
- [Get Nft History](../../doc/controllers/non-fungible-tokens.md#get-nft-history)
- [Get Nft Mints](../../doc/controllers/non-fungible-tokens.md#get-nft-mints)

# Get Nft Holdings

Retrieves the list of Non-Fungible Tokens owned by the given principal (STX address or Smart Contract ID).
Results can be filtered by one or more asset identifiers and can include metadata about the transaction that made the principal own each token.

More information on Non-Fungible Tokens on the Stacks blockchain can be found [here](https://docs.stacks.co/write-smart-contracts/tokens#non-fungible-tokens-nfts).

```go
GetNftHoldings(
    ctx context.Context,
    principal string,
    assetIdentifiers []string,
    limit *int,
    offset *int,
    unanchored *bool,
    txMetadata *bool) (
    models.ApiResponse[models.NonFungibleTokenHoldingsList],
    error)
```

## Parameters

| Parameter          | Type       | Tags            | Description                                                                                                                                          |
| ------------------ | ---------- | --------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `principal`        | `string`   | Query, Required | token owner's STX address or Smart Contract ID                                                                                                       |
| `assetIdentifiers` | `[]string` | Query, Optional | identifiers of the token asset classes to filter for                                                                                                 |
| `limit`            | `*int`     | Query, Optional | max number of tokens to fetch                                                                                                                        |
| `offset`           | `*int`     | Query, Optional | index of first tokens to fetch                                                                                                                       |
| `unanchored`       | `*bool`    | Query, Optional | whether or not to include tokens from unconfirmed transactions                                                                                       |
| `txMetadata`       | `*bool`    | Query, Optional | whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times. |

## Response Type

[`models.NonFungibleTokenHoldingsList`](../../doc/models/non-fungible-token-holdings-list.md)

## Example Usage

```go
ctx := context.Background()
principal := "SPNWZ5V2TPWGQGVDR6T7B6RQ4XMGZ4PXTEE0VQ0S.marketplace-v3"
assetIdentifiers := []string{"SPQZF23W7SEYBFG5JQ496NMY0G7379SRYEDREMSV.Candy::candy"}
limit := 50
offset := 42000
unanchored := true
txMetadata := false

apiResponse, err := nonFungibleTokensController.GetNftHoldings(ctx, principal, assetIdentifiers, &limit, &offset, &unanchored, &txMetadata)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Nft History

Retrieves all events relevant to a Non-Fungible Token. Useful to determine the ownership history of a particular asset.

More information on Non-Fungible Tokens on the Stacks blockchain can be found [here](https://docs.stacks.co/write-smart-contracts/tokens#non-fungible-tokens-nfts).

```go
GetNftHistory(
    ctx context.Context,
    assetIdentifier string,
    value string,
    limit *int,
    offset *int,
    unanchored *bool,
    txMetadata *bool) (
    models.ApiResponse[models.NonFungibleTokenHistoryEventList],
    error)
```

## Parameters

| Parameter         | Type     | Tags            | Description                                                                                                                                          |
| ----------------- | -------- | --------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `assetIdentifier` | `string` | Query, Required | token asset class identifier                                                                                                                         |
| `value`           | `string` | Query, Required | hex representation of the token's unique value                                                                                                       |
| `limit`           | `*int`   | Query, Optional | max number of events to fetch                                                                                                                        |
| `offset`          | `*int`   | Query, Optional | index of first event to fetch                                                                                                                        |
| `unanchored`      | `*bool`  | Query, Optional | whether or not to include events from unconfirmed transactions                                                                                       |
| `txMetadata`      | `*bool`  | Query, Optional | whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times. |

## Response Type

[`models.NonFungibleTokenHistoryEventList`](../../doc/models/non-fungible-token-history-event-list.md)

## Example Usage

```go
ctx := context.Background()
assetIdentifier := "SP2X0TZ59D5SZ8ACQ6YMCHHNR2ZN51Z32E2CJ173.the-explorer-guild::The-Explorer-Guild"
value := "0x0100000000000000000000000000000803"
limit := 50
offset := 42000
unanchored := true
txMetadata := false

apiResponse, err := nonFungibleTokensController.GetNftHistory(ctx, assetIdentifier, value, &limit, &offset, &unanchored, &txMetadata)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Nft Mints

Retrieves all mint events for a Non-Fungible Token asset class. Useful to determine which NFTs of a particular collection have been claimed.

More information on Non-Fungible Tokens on the Stacks blockchain can be found [here](https://docs.stacks.co/write-smart-contracts/tokens#non-fungible-tokens-nfts).

```go
GetNftMints(
    ctx context.Context,
    assetIdentifier string,
    limit *int,
    offset *int,
    unanchored *bool,
    txMetadata *bool) (
    models.ApiResponse[models.NonFungibleTokenMintList],
    error)
```

## Parameters

| Parameter         | Type     | Tags            | Description                                                                                                                                          |
| ----------------- | -------- | --------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `assetIdentifier` | `string` | Query, Required | token asset class identifier                                                                                                                         |
| `limit`           | `*int`   | Query, Optional | max number of events to fetch                                                                                                                        |
| `offset`          | `*int`   | Query, Optional | index of first event to fetch                                                                                                                        |
| `unanchored`      | `*bool`  | Query, Optional | whether or not to include events from unconfirmed transactions                                                                                       |
| `txMetadata`      | `*bool`  | Query, Optional | whether or not to include the complete transaction metadata instead of just `tx_id`. Enabling this option can affect performance and response times. |

## Response Type

[`models.NonFungibleTokenMintList`](../../doc/models/non-fungible-token-mint-list.md)

## Example Usage

```go
ctx := context.Background()
assetIdentifier := "SP2X0TZ59D5SZ8ACQ6YMCHHNR2ZN51Z32E2CJ173.the-explorer-guild::The-Explorer-Guild"
limit := 50
offset := 42000
unanchored := true
txMetadata := false

apiResponse, err := nonFungibleTokensController.GetNftMints(ctx, assetIdentifier, &limit, &offset, &unanchored, &txMetadata)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```
