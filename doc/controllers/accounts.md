# Accounts

Read-only endpoints to obtain Stacks account details

Stacks Documentation - Accounts: [https://docs.stacks.co/understand-stacks/accounts](https://docs.stacks.co/understand-stacks/accounts)

```go
accountsController := client.AccountsController()
```

## Class Name

`AccountsController`

## Methods

- [Get Account Balance](../../doc/controllers/accounts.md#get-account-balance)
- [Get Account Stx Balance](../../doc/controllers/accounts.md#get-account-stx-balance)
- [Get Account Transactions](../../doc/controllers/accounts.md#get-account-transactions)
- [Get Single Transaction With Transfers](../../doc/controllers/accounts.md#get-single-transaction-with-transfers)
- [Get Account Transactions With Transfers](../../doc/controllers/accounts.md#get-account-transactions-with-transfers)
- [Get Account Nonces](../../doc/controllers/accounts.md#get-account-nonces)
- [Get Account Assets](../../doc/controllers/accounts.md#get-account-assets)
- [Get Account Inbound](../../doc/controllers/accounts.md#get-account-inbound)
- [Get Account Info](../../doc/controllers/accounts.md#get-account-info)

# Get Account Balance

Retrieves total account balance information for a given Address or Contract Identifier. This includes the balances of STX Tokens, Fungible Tokens and Non-Fungible Tokens for the account.

```go
GetAccountBalance(
    ctx context.Context,
    principal string,
    unanchored *bool,
    untilBlock *string) (
    models.ApiResponse[models.AddressBalanceResponse],
    error)
```

## Parameters

| Parameter    | Type      | Tags               | Description                                                                                      |
| ------------ | --------- | ------------------ | ------------------------------------------------------------------------------------------------ |
| `principal`  | `string`  | Template, Required | Stacks address or a Contract identifier                                                          |
| `unanchored` | `*bool`   | Query, Optional    | Include transaction data from unanchored (i.e. unconfirmed) microblocks                          |
| `untilBlock` | `*string` | Query, Optional    | returned data representing the state up until that point in time, rather than the current block. |

## Response Type

[`models.AddressBalanceResponse`](../../doc/models/address-balance-response.md)

## Example Usage

```go
ctx := context.Background()
principal := "SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0"
unanchored := true
untilBlock := "60000"

apiResponse, err := accountsController.GetAccountBalance(ctx, principal, &unanchored, &untilBlock)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Account Stx Balance

Retrieves STX token balance for a given Address or Contract Identifier.

```go
GetAccountStxBalance(
    ctx context.Context,
    principal string,
    unanchored *bool,
    untilBlock *string) (
    models.ApiResponse[models.AddressStxBalanceResponse],
    error)
```

## Parameters

| Parameter    | Type      | Tags               | Description                                                                                                                                                        |
| ------------ | --------- | ------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `principal`  | `string`  | Template, Required | Stacks address or a Contract identifier.                                                                                                                           |
| `unanchored` | `*bool`   | Query, Optional    | Include transaction data from unanchored (i.e. unconfirmed) microblocks.                                                                                           |
| `untilBlock` | `*string` | Query, Optional    | returned data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. |

## Response Type

[`models.AddressStxBalanceResponse`](../../doc/models/address-stx-balance-response.md)

## Example Usage

```go
ctx := context.Background()
principal := "SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0"
unanchored := true
untilBlock := "60000"

apiResponse, err := accountsController.GetAccountStxBalance(ctx, principal, &unanchored, &untilBlock)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Account Transactions

Retrieves a list of all Transactions for a given Address or Contract Identifier. More information on Transaction types can be found [here](https://docs.stacks.co/understand-stacks/transactions#types).

If you need to actively monitor new transactions for an address or contract id, we highly recommend subscribing to [WebSockets or Socket.io](https://github.com/hirosystems/stacks-blockchain-api/tree/master/client) for real-time updates.

```go
GetAccountTransactions(
    ctx context.Context,
    principal string,
    limit *int,
    offset *int,
    height *float64,
    unanchored *bool,
    untilBlock *string) (
    models.ApiResponse[models.AddressTransactionsListResponse],
    error)
```

## Parameters

| Parameter    | Type       | Tags               | Description                                                                                                                                                        |
| ------------ | ---------- | ------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `principal`  | `string`   | Template, Required | Stacks address or a Contract identifier                                                                                                                            |
| `limit`      | `*int`     | Query, Optional    | max number of account transactions to fetch                                                                                                                        |
| `offset`     | `*int`     | Query, Optional    | index of first account transaction to fetch                                                                                                                        |
| `height`     | `*float64` | Query, Optional    | Filter for transactions only at this given block height                                                                                                            |
| `unanchored` | `*bool`    | Query, Optional    | Include transaction data from unanchored (i.e. unconfirmed) microblocks                                                                                            |
| `untilBlock` | `*string`  | Query, Optional    | returned data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. |

## Response Type

[`models.AddressTransactionsListResponse`](../../doc/models/address-transactions-list-response.md)

## Example Usage

```go
ctx := context.Background()
principal := "SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0"
limit := 42000
offset := 42000
height := float64(42000)
unanchored := true
untilBlock := "60000"

apiResponse, err := accountsController.GetAccountTransactions(ctx, principal, &limit, &offset, &height, &unanchored, &untilBlock)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Single Transaction With Transfers

Retrieves transaction details for a given Transaction Id `tx_id`, for a given account or contract Identifier.

```go
GetSingleTransactionWithTransfers(
    ctx context.Context,
    principal string,
    txId string) (
    models.ApiResponse[models.AddressTransactionWithTransfers],
    error)
```

## Parameters

| Parameter   | Type     | Tags               | Description                             |
| ----------- | -------- | ------------------ | --------------------------------------- |
| `principal` | `string` | Template, Required | Stacks address or a contract identifier |
| `txId`      | `string` | Template, Required | Transaction id                          |

## Response Type

[`models.AddressTransactionWithTransfers`](../../doc/models/address-transaction-with-transfers.md)

## Example Usage

```go
ctx := context.Background()
principal := "SP3FBR2AGK5H9QBDH3EEN6DF8EK8JY7RX8QJ5SVTE"
txId := "0x34d79c7cfc2fe525438736733e501a4bf0308a5556e3e080d1e2c0858aad7448"

apiResponse, err := accountsController.GetSingleTransactionWithTransfers(ctx, principal, txId)
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

# Get Account Transactions With Transfers

Retrieve all transactions for an account or contract identifier including STX transfers for each transaction.

```go
GetAccountTransactionsWithTransfers(
    ctx context.Context,
    principal string,
    limit *int,
    offset *int,
    height *float64,
    unanchored *bool,
    untilBlock *string) (
    models.ApiResponse[models.AddressTransactionsWithTransfersListResponse],
    error)
```

## Parameters

| Parameter    | Type       | Tags               | Description                                                                                      |
| ------------ | ---------- | ------------------ | ------------------------------------------------------------------------------------------------ |
| `principal`  | `string`   | Template, Required | Stacks address or a Contract identifier                                                          |
| `limit`      | `*int`     | Query, Optional    | max number of account transactions to fetch                                                      |
| `offset`     | `*int`     | Query, Optional    | index of first account transaction to fetch                                                      |
| `height`     | `*float64` | Query, Optional    | Filter for transactions only at this given block height                                          |
| `unanchored` | `*bool`    | Query, Optional    | Include transaction data from unanchored (i.e. unconfirmed) microblocks                          |
| `untilBlock` | `*string`  | Query, Optional    | returned data representing the state up until that point in time, rather than the current block. |

## Response Type

[`models.AddressTransactionsWithTransfersListResponse`](../../doc/models/address-transactions-with-transfers-list-response.md)

## Example Usage

```go
ctx := context.Background()
principal := "SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0"
limit := 20
offset := 10
height := float64(66119)
unanchored := true
untilBlock := "60000"

apiResponse, err := accountsController.GetAccountTransactionsWithTransfers(ctx, principal, &limit, &offset, &height, &unanchored, &untilBlock)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Account Nonces

Retrieves the latest nonce values used by an account by inspecting the mempool, microblock transactions, and anchored transactions.

```go
GetAccountNonces(
    ctx context.Context,
    principal string,
    blockHeight *float64,
    blockHash *string) (
    models.ApiResponse[models.AddressNonces],
    error)
```

## Parameters

| Parameter     | Type       | Tags               | Description                                                                                                       |
| ------------- | ---------- | ------------------ | ----------------------------------------------------------------------------------------------------------------- |
| `principal`   | `string`   | Template, Required | Stacks address                                                                                                    |
| `blockHeight` | `*float64` | Query, Optional    | Optionally get the nonce at a given block height.                                                                 |
| `blockHash`   | `*string`  | Query, Optional    | Optionally get the nonce at a given block hash. Note - Use either of the query parameters but not both at a time. |

## Response Type

[`models.AddressNonces`](../../doc/models/address-nonces.md)

## Example Usage

```go
ctx := context.Background()
principal := "SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0"
blockHeight := float64(66119)
blockHash := "0x72d53f3cba39e149dcd42708e535bdae03d73e60d2fe853aaf61c0b392f521e9"

apiResponse, err := accountsController.GetAccountNonces(ctx, principal, &blockHeight, &blockHash)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Account Assets

Retrieves a list of all assets events associated with an account or a Contract Identifier. This includes Transfers, Mints.

```go
GetAccountAssets(
    ctx context.Context,
    principal string,
    limit *int,
    offset *int,
    unanchored *bool,
    untilBlock *string) (
    models.ApiResponse[models.AddressAssetsListResponse],
    error)
```

## Parameters

| Parameter    | Type      | Tags               | Description                                                                                                                                                  |
| ------------ | --------- | ------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `principal`  | `string`  | Template, Required | Stacks address or a Contract identifier                                                                                                                      |
| `limit`      | `*int`    | Query, Optional    | max number of account assets to fetch                                                                                                                        |
| `offset`     | `*int`    | Query, Optional    | index of first account assets to fetch                                                                                                                       |
| `unanchored` | `*bool`   | Query, Optional    | Include transaction data from unanchored (i.e. unconfirmed) microblocks                                                                                      |
| `untilBlock` | `*string` | Query, Optional    | returned data representing the state at that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. |

## Response Type

[`models.AddressAssetsListResponse`](../../doc/models/address-assets-list-response.md)

## Example Usage

```go
ctx := context.Background()
principal := "SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0"
limit := 20
offset := 42000
unanchored := true
untilBlock := "60000"

apiResponse, err := accountsController.GetAccountAssets(ctx, principal, &limit, &offset, &unanchored, &untilBlock)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Account Inbound

Retrieves a list of STX transfers with memos to the given principal. This includes regular transfers from a stx-transfer transaction type,
and transfers from contract-call transactions a the `send-many-memo` bulk sending contract.

```go
GetAccountInbound(
    ctx context.Context,
    principal string,
    limit *int,
    offset *int,
    height *float64,
    unanchored *bool,
    untilBlock *string) (
    models.ApiResponse[models.AddressStxInboundListResponse],
    error)
```

## Parameters

| Parameter    | Type       | Tags               | Description                                                                                                                                                        |
| ------------ | ---------- | ------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `principal`  | `string`   | Template, Required | Stacks address or a Contract identifier                                                                                                                            |
| `limit`      | `*int`     | Query, Optional    | number of items to return                                                                                                                                          |
| `offset`     | `*int`     | Query, Optional    | number of items to skip                                                                                                                                            |
| `height`     | `*float64` | Query, Optional    | Filter for transfers only at this given block height                                                                                                               |
| `unanchored` | `*bool`    | Query, Optional    | Include transaction data from unanchored (i.e. unconfirmed) microblocks                                                                                            |
| `untilBlock` | `*string`  | Query, Optional    | returned data representing the state up until that point in time, rather than the current block. Note - Use either of the query parameters but not both at a time. |

## Response Type

[`models.AddressStxInboundListResponse`](../../doc/models/address-stx-inbound-list-response.md)

## Example Usage

```go
ctx := context.Background()
principal := "SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0"
offset := 42000
unanchored := true
untilBlock := "60000"

apiResponse, err := accountsController.GetAccountInbound(ctx, principal, nil, &offset, nil, &unanchored, &untilBlock)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Account Info

Retrieves the account data for a given Account or a Contract Identifier

Where balance is the hex encoding of a unsigned 128-bit integer (big-endian), nonce is an unsigned 64-bit integer, and the proofs are provided as hex strings.

For non-existent accounts, this does not return a 404 error, rather it returns an object with balance and nonce of 0.

```go
GetAccountInfo(
    ctx context.Context,
    principal string,
    proof *int,
    tip *string) (
    models.ApiResponse[models.AccountDataResponse],
    error)
```

## Parameters

| Parameter   | Type      | Tags               | Description                                        |
| ----------- | --------- | ------------------ | -------------------------------------------------- |
| `principal` | `string`  | Template, Required | Stacks address or a Contract identifier            |
| `proof`     | `*int`    | Query, Optional    | Returns object without the proof field if set to 0 |
| `tip`       | `*string` | Query, Optional    | The Stacks chain tip to query from                 |

## Response Type

[`models.AccountDataResponse`](../../doc/models/account-data-response.md)

## Example Usage

```go
ctx := context.Background()
principal := "SP31DA6FTSJX2WGTZ69SFY11BH51NZMB0ZW97B5P0"

apiResponse, err := accountsController.GetAccountInfo(ctx, principal, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```
