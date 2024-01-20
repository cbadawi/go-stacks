# Transactions

Endpoints to obtain transaction details and to broadcast transactions to the network

Hiro Documentation - Transactions: [https://docs.hiro.so/get-started/transactions](https://docs.hiro.so/get-started/transactions)

```go
transactionsController := client.TransactionsController()
```

## Class Name

`TransactionsController`

## Methods

- [Get Transaction List](../../doc/controllers/transactions.md#get-transaction-list)
- [Get Mempool Transaction List](../../doc/controllers/transactions.md#get-mempool-transaction-list)
- [Get Dropped Mempool Transaction List](../../doc/controllers/transactions.md#get-dropped-mempool-transaction-list)
- [Get Mempool Transaction Stats](../../doc/controllers/transactions.md#get-mempool-transaction-stats)
- [Get Tx List Details](../../doc/controllers/transactions.md#get-tx-list-details)
- [Get Transaction by Id](../../doc/controllers/transactions.md#get-transaction-by-id)
- [Get Raw Transaction by Id](../../doc/controllers/transactions.md#get-raw-transaction-by-id)
- [Post Core Node Transactions](../../doc/controllers/transactions.md#post-core-node-transactions)
- [Get Transactions by Block](../../doc/controllers/transactions.md#get-transactions-by-block)
- [Get Transactions by Block Hash](../../doc/controllers/transactions.md#get-transactions-by-block-hash)
- [Get Transactions by Block Height](../../doc/controllers/transactions.md#get-transactions-by-block-height)
- [Get Address Mempool Transactions](../../doc/controllers/transactions.md#get-address-mempool-transactions)
- [Get Filtered Events](../../doc/controllers/transactions.md#get-filtered-events)

# Get Transaction List

Retrieves all recently mined transactions

If using TypeScript, import typings for this response from our types package:

`import type { TransactionResults } from '@stacks/stacks-blockchain-api-types';`

```go
GetTransactionList(
    ctx context.Context,
    limit *int,
    offset *int,
    mType []models.TypeEnum,
    unanchored *bool) (
    models.ApiResponse[models.TransactionResults],
    error)
```

## Parameters

| Parameter    | Type                                                 | Tags            | Description                                                             |
| ------------ | ---------------------------------------------------- | --------------- | ----------------------------------------------------------------------- |
| `limit`      | `*int`                                               | Query, Optional | max number of transactions to fetch                                     |
| `offset`     | `*int`                                               | Query, Optional | index of first transaction to fetch                                     |
| `mType`      | [`[]models.TypeEnum`](../../doc/models/type-enum.md) | Query, Optional | Filter by transaction type                                              |
| `unanchored` | `*bool`                                              | Query, Optional | Include transaction data from unanchored (i.e. unconfirmed) microblocks |

## Response Type

[`models.TransactionResults`](../../doc/models/transaction-results.md)

## Example Usage

```go
ctx := context.Background()
limit := 100
offset := 42000
unanchored := true

apiResponse, err := transactionsController.GetTransactionList(ctx, &limit, &offset, nil, &unanchored)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Mempool Transaction List

Retrieves all transactions that have been recently broadcast to the mempool. These are pending transactions awaiting confirmation.

If you need to monitor new transactions, we highly recommend subscribing to [WebSockets or Socket.io](https://github.com/hirosystems/stacks-blockchain-api/tree/master/client) for real-time updates.

```go
GetMempoolTransactionList(
    ctx context.Context,
    senderAddress *string,
    recipientAddress *string,
    address *string,
    orderBy *models.OrderByEnum,
    order *models.OrderEnum,
    limit *int,
    offset *int,
    unanchored *bool) (
    models.ApiResponse[models.MempoolTransactionListResponse],
    error)
```

## Parameters

| Parameter          | Type                                                       | Tags            | Description                                                                                                                            |
| ------------------ | ---------------------------------------------------------- | --------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `senderAddress`    | `*string`                                                  | Query, Optional | Filter to only return transactions with this sender address.                                                                           |
| `recipientAddress` | `*string`                                                  | Query, Optional | Filter to only return transactions with this recipient address (only applicable for STX transfer tx types).                            |
| `address`          | `*string`                                                  | Query, Optional | Filter to only return transactions with this address as the sender or recipient (recipient only applicable for STX transfer tx types). |
| `orderBy`          | [`*models.OrderByEnum`](../../doc/models/order-by-enum.md) | Query, Optional | Option to sort results by transaction age, size, or fee rate.                                                                          |
| `order`            | [`*models.OrderEnum`](../../doc/models/order-enum.md)      | Query, Optional | Option to sort results in ascending or descending order.                                                                               |
| `limit`            | `*int`                                                     | Query, Optional | max number of mempool transactions to fetch                                                                                            |
| `offset`           | `*int`                                                     | Query, Optional | index of first mempool transaction to fetch                                                                                            |
| `unanchored`       | `*bool`                                                    | Query, Optional | Include transaction data from unanchored (i.e. unconfirmed) microblocks                                                                |

## Response Type

[`models.MempoolTransactionListResponse`](../../doc/models/mempool-transaction-list-response.md)

## Example Usage

```go
ctx := context.Background()
senderAddress := "SP1GPBP8NBRXDRJBFQBV7KMAZX1Z7W2RFWJEH0V10"
orderBy := models.OrderByEnum("fee")
order := models.OrderEnum("asc")
limit := 20
offset := 42000
unanchored := true

apiResponse, err := transactionsController.GetMempoolTransactionList(ctx, &senderAddress, nil, nil, &orderBy, &order, &limit, &offset, &unanchored)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Dropped Mempool Transaction List

Retrieves all recently-broadcast transactions that have been dropped from the mempool.

Transactions are dropped from the mempool if:

- they were stale and awaiting garbage collection or,
- were expensive, or
- were replaced with a new fee

```go
GetDroppedMempoolTransactionList(
    ctx context.Context,
    limit *int,
    offset *int) (
    models.ApiResponse[models.MempoolTransactionListResponse],
    error)
```

## Parameters

| Parameter | Type   | Tags            | Description                                 |
| --------- | ------ | --------------- | ------------------------------------------- |
| `limit`   | `*int` | Query, Optional | max number of mempool transactions to fetch |
| `offset`  | `*int` | Query, Optional | index of first mempool transaction to fetch |

## Response Type

[`models.MempoolTransactionListResponse`](../../doc/models/mempool-transaction-list-response.md)

## Example Usage

```go
ctx := context.Background()
limit := 96
offset := 42000

apiResponse, err := transactionsController.GetDroppedMempoolTransactionList(ctx, &limit, &offset)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Mempool Transaction Stats

Queries for transactions counts, age (by block height), fees (simple average), and size.
All results broken down by transaction type and percentiles (p25, p50, p75, p95).

```go
GetMempoolTransactionStats(
    ctx context.Context) (
    models.ApiResponse[models.MempoolTransactionStatsResponse],
    error)
```

## Response Type

[`models.MempoolTransactionStatsResponse`](../../doc/models/mempool-transaction-stats-response.md)

## Example Usage

```go
ctx := context.Background()
apiResponse, err := transactionsController.GetMempoolTransactionStats(ctx)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Tx List Details

Retrieves a list of transactions for a given list of transaction IDs

If using TypeScript, import typings for this response from our types package:

`import type { Transaction } from '@stacks/stacks-blockchain-api-types';`

```go
GetTxListDetails(
    ctx context.Context,
    txId []string,
    eventOffset *int,
    eventLimit *int,
    unanchored *bool) (
    models.ApiResponse[map[string]models.TransactionList],
    error)
```

## Parameters

| Parameter     | Type       | Tags            | Description                                                             |
| ------------- | ---------- | --------------- | ----------------------------------------------------------------------- |
| `txId`        | `[]string` | Query, Required | Array of transaction ids                                                |
| `eventOffset` | `*int`     | Query, Optional | The number of events to skip                                            |
| `eventLimit`  | `*int`     | Query, Optional | The numbers of events to return                                         |
| `unanchored`  | `*bool`    | Query, Optional | Include transaction data from unanchored (i.e. unconfirmed) microblocks |

## Response Type

[`map[string]models.TransactionList`](../../doc/models/transaction-list.md)

## Example Usage

```go
ctx := context.Background()
txId := []string{"0x0a411719e3bfde95f9e227a2d7f8fac3d6c646b1e6cc186db0e2838a2c6cd9c0", "0xfbe6412b46e9db87df991a0d043ff47eb58019f770208cf48e2380337cc31785", "0x178b77678a758d9f29a147d6399315c83d0f1baf114431fbe4d3587aa5fbba6f"}
eventOffset := 0
eventLimit := 96
unanchored := true

apiResponse, err := transactionsController.GetTxListDetails(ctx, txId, &eventOffset, &eventLimit, &unanchored)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description                    | Exception Class |
| ---------------- | ------------------------------------ | --------------- |
| 404              | Could not find any transaction by ID | `ApiError`      |

# Get Transaction by Id

Retrieves transaction details for a given transaction ID

`import type { Transaction } from '@stacks/stacks-blockchain-api-types';`

```go
GetTransactionById(
    ctx context.Context,
    txId string,
    eventOffset *int,
    eventLimit *int,
    unanchored *bool) (
    models.ApiResponse[models.Transaction],
    error)
```

## Parameters

| Parameter     | Type     | Tags               | Description                                                             |
| ------------- | -------- | ------------------ | ----------------------------------------------------------------------- |
| `txId`        | `string` | Template, Required | Hash of transaction                                                     |
| `eventOffset` | `*int`   | Query, Optional    | The number of events to skip                                            |
| `eventLimit`  | `*int`   | Query, Optional    | The numbers of events to return                                         |
| `unanchored`  | `*bool`  | Query, Optional    | Include transaction data from unanchored (i.e. unconfirmed) microblocks |

## Response Type

[`models.Transaction`](../../doc/models/transaction.md)

## Example Usage

```go
ctx := context.Background()
txId := "0x0a411719e3bfde95f9e227a2d7f8fac3d6c646b1e6cc186db0e2838a2c6cd9c0"
eventOffset := 0
eventLimit := 96
unanchored := true

apiResponse, err := transactionsController.GetTransactionById(ctx, txId, &eventOffset, &eventLimit, &unanchored)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description                    | Exception Class |
| ---------------- | ------------------------------------ | --------------- |
| 404              | Cannot find transaction for given ID | `ApiError`      |

# Get Raw Transaction by Id

Retrieves a hex encoded serialized transaction for a given ID

```go
GetRawTransactionById(
    ctx context.Context,
    txId string) (
    models.ApiResponse[models.GetRawTransactionResult],
    error)
```

## Parameters

| Parameter | Type     | Tags               | Description         |
| --------- | -------- | ------------------ | ------------------- |
| `txId`    | `string` | Template, Required | Hash of transaction |

## Response Type

[`models.GetRawTransactionResult`](../../doc/models/get-raw-transaction-result.md)

## Example Usage

```go
ctx := context.Background()
txId := "0x0a411719e3bfde95f9e227a2d7f8fac3d6c646b1e6cc186db0e2838a2c6cd9c0"

apiResponse, err := transactionsController.GetRawTransactionById(ctx, txId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description                    | Exception Class |
| ---------------- | ------------------------------------ | --------------- |
| 404              | Cannot find transaction for given ID | `ApiError`      |

# Post Core Node Transactions

Broadcasts raw transactions on the network. You can use the [@stacks/transactions](https://github.com/blockstack/stacks.js) project to generate a raw transaction payload.

```go
PostCoreNodeTransactions(
    ctx context.Context,
    body *models.FileWrapper) (
    models.ApiResponse[string],
    error)
```

## Parameters

| Parameter | Type                  | Tags           | Description |
| --------- | --------------------- | -------------- | ----------- |
| `body`    | `*models.FileWrapper` | Form, Optional | -           |

## Response Type

`string`

## Example Usage

```go
ctx := context.Background()

apiResponse, err := transactionsController.PostCoreNodeTransactions(ctx, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description                | Exception Class                                                                                             |
| ---------------- | -------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| 400              | Rejections result in a 400 error | [`PostCoreNodeTransactionsErrorException`](../../doc/models/post-core-node-transactions-error-exception.md) |

# Get Transactions by Block

Retrieves transactions confirmed in a single block

```go
GetTransactionsByBlock(
    ctx context.Context,
    heightOrHash interface{}) (
    models.ApiResponse[models.TransactionResults],
    error)
```

## Parameters

| Parameter      | Type          | Tags               | Description                                                                                                 |
| -------------- | ------------- | ------------------ | ----------------------------------------------------------------------------------------------------------- |
| `heightOrHash` | `interface{}` | Template, Required | filter by block height, hash, index block hash or the constant `latest` to filter for the most recent block |

## Response Type

[`models.TransactionResults`](../../doc/models/transaction-results.md)

## Example Usage

```go
ctx := context.Background()
heightOrHash := interface{}("[key1, val1][key2, val2]")

apiResponse, err := transactionsController.GetTransactionsByBlock(ctx, heightOrHash)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Transactions by Block Hash

**This endpoint is deprecated.**

**NOTE:** This endpoint is deprecated in favor of [Get transactions by block](#operation/get_transactions_by_block).

Retrieves a list of all transactions within a block for a given block hash.

```go
GetTransactionsByBlockHash(
    ctx context.Context,
    blockHash string,
    limit *int,
    offset *int) (
    models.ApiResponse[models.TransactionResults],
    error)
```

## Parameters

| Parameter   | Type     | Tags               | Description                         |
| ----------- | -------- | ------------------ | ----------------------------------- |
| `blockHash` | `string` | Template, Required | Hash of block                       |
| `limit`     | `*int`   | Query, Optional    | max number of transactions to fetch |
| `offset`    | `*int`   | Query, Optional    | index of first transaction to fetch |

## Response Type

[`models.TransactionResults`](../../doc/models/transaction-results.md)

## Example Usage

```go
ctx := context.Background()
blockHash := "0x0a83d82a65460a9e711f85a44616350280040b75317dbe486a923c1131b5ff99"
limit := 10
offset := 42000

apiResponse, err := transactionsController.GetTransactionsByBlockHash(ctx, blockHash, &limit, &offset)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Transactions by Block Height

**This endpoint is deprecated.**

**NOTE:** This endpoint is deprecated in favor of [Get transactions by block](#operation/get_transactions_by_block).

Retrieves all transactions within a block at a given height

```go
GetTransactionsByBlockHeight(
    ctx context.Context,
    height int,
    limit *int,
    offset *int,
    unanchored *bool) (
    models.ApiResponse[models.TransactionResults],
    error)
```

## Parameters

| Parameter    | Type    | Tags               | Description                                                             |
| ------------ | ------- | ------------------ | ----------------------------------------------------------------------- |
| `height`     | `int`   | Template, Required | Height of block                                                         |
| `limit`      | `*int`  | Query, Optional    | max number of transactions to fetch                                     |
| `offset`     | `*int`  | Query, Optional    | index of first transaction to fetch                                     |
| `unanchored` | `*bool` | Query, Optional    | Include transaction data from unanchored (i.e. unconfirmed) microblocks |

## Response Type

[`models.TransactionResults`](../../doc/models/transaction-results.md)

## Example Usage

```go
ctx := context.Background()
height := 66119
limit := 10
offset := 42000
unanchored := true

apiResponse, err := transactionsController.GetTransactionsByBlockHeight(ctx, height, &limit, &offset, &unanchored)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Address Mempool Transactions

Retrieves all transactions for a given address that are currently in mempool

```go
GetAddressMempoolTransactions(
    ctx context.Context,
    address string,
    limit *int,
    offset *int,
    unanchored *bool) (
    models.ApiResponse[models.MempoolTransactionListResponse],
    error)
```

## Parameters

| Parameter    | Type     | Tags               | Description                                                             |
| ------------ | -------- | ------------------ | ----------------------------------------------------------------------- |
| `address`    | `string` | Template, Required | Transactions for the address                                            |
| `limit`      | `*int`   | Query, Optional    | max number of transactions to fetch                                     |
| `offset`     | `*int`   | Query, Optional    | index of first transaction to fetch                                     |
| `unanchored` | `*bool`  | Query, Optional    | Include transaction data from unanchored (i.e. unconfirmed) microblocks |

## Response Type

[`models.MempoolTransactionListResponse`](../../doc/models/mempool-transaction-list-response.md)

## Example Usage

```go
ctx := context.Background()
address := "SP197DVH8KTJGX4STM61QN0WJV8Y9QJWXV83ZGNR9"
limit := 90
offset := 42000
unanchored := true

apiResponse, err := transactionsController.GetAddressMempoolTransactions(ctx, address, &limit, &offset, &unanchored)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Filtered Events

Retrieves the list of events filtered by principal (STX address or Smart Contract ID), transaction id or event types. The list of event types is ('smart_contract_log', 'stx_lock', 'stx_asset', 'fungible_token_asset', 'non_fungible_token_asset').

```go
GetFilteredEvents(
    ctx context.Context,
    txId *string,
    address *string,
    limit *int,
    offset *int,
    mType []models.Type2Enum) (
    models.ApiResponse[models.TransactionEventsResponse],
    error)
```

## Parameters

| Parameter | Type                                                    | Tags            | Description                             |
| --------- | ------------------------------------------------------- | --------------- | --------------------------------------- |
| `txId`    | `*string`                                               | Query, Optional | Hash of transaction                     |
| `address` | `*string`                                               | Query, Optional | Stacks address or a Contract identifier |
| `limit`   | `*int`                                                  | Query, Optional | number of items to return               |
| `offset`  | `*int`                                                  | Query, Optional | number of items to skip                 |
| `mType`   | [`[]models.Type2Enum`](../../doc/models/type-2-enum.md) | Query, Optional | Filter the events on event type         |

## Response Type

[`models.TransactionEventsResponse`](../../doc/models/transaction-events-response.md)

## Example Usage

```go
ctx := context.Background()
txId := "0x29e25515652dad41ef675bd0670964e3d537b80ec19cf6ca6f1dd65d5bc642c5"
address := "ST1HB64MAJ1MBV4CQ80GF01DZS4T1DSMX20ADCRA4"
limit := 100
offset := 42000

apiResponse, err := transactionsController.GetFilteredEvents(ctx, &txId, &address, &limit, &offset, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```
