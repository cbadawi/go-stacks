# Faucets

Endpoints to request STX or BTC tokens (not possible on Mainnet)

```go
faucetsController := client.FaucetsController()
```

## Class Name

`FaucetsController`

## Methods

- [Run Faucet Stx](../../doc/controllers/faucets.md#run-faucet-stx)
- [Run Faucet Btc](../../doc/controllers/faucets.md#run-faucet-btc)

# Run Faucet Stx

Add 500 STX tokens to the specified testnet address. Testnet STX addresses begin with `ST`. If the `stacking`
parameter is set to `true`, the faucet will add the required number of tokens for individual stacking to the
specified testnet address.

The endpoint returns the transaction ID, which you can use to view the transaction in the
[Stacks Explorer](https://explorer.hiro.so/?chain=testnet). The tokens are delivered once the transaction has
been included in an anchor block.

A common reason for failed faucet transactions is that the faucet has run out of tokens. If you are experiencing
failed faucet transactions to a testnet address, you can get help in [Discord](https://stacks.chat).

**Note:** This is a testnet only endpoint. This endpoint will not work on the mainnet.

```go
RunFaucetStx(
    ctx context.Context,
    address string,
    stacking *bool) (
    models.ApiResponse[models.RunFaucetResponse],
    error)
```

## Parameters

| Parameter  | Type     | Tags            | Description                                                             |
| ---------- | -------- | --------------- | ----------------------------------------------------------------------- |
| `address`  | `string` | Query, Required | A valid testnet STX address                                             |
| `stacking` | `*bool`  | Query, Optional | Request the amount of STX tokens needed for individual address stacking |

## Response Type

[`models.RunFaucetResponse`](../../doc/models/run-faucet-response.md)

## Example Usage

```go
ctx := context.Background()
address := "ST3M7N9Q9HDRM7RVP1Q26P0EE69358PZZAZD7KMXQ"
stacking := false

apiResponse, err := faucetsController.RunFaucetStx(ctx, address, &stacking)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description  | Exception Class |
| ---------------- | ------------------ | --------------- |
| 500              | Faucet call failed | `ApiError`      |

# Run Faucet Btc

Add 1 BTC token to the specified testnet BTC address.

The endpoint returns the transaction ID, which you can use to view the transaction in a testnet Bitcoin block
explorer. The tokens are delivered once the transaction has been included in a block.

**Note:** This is a testnet only endpoint. This endpoint will not work on the mainnet.

```go
RunFaucetBtc(
    ctx context.Context,
    address string,
    body *models.ExtendedV1FaucetsBtcRequest) (
    models.ApiResponse[models.RunFaucetResponse],
    error)
```

## Parameters

| Parameter | Type                                                                                         | Tags            | Description                 |
| --------- | -------------------------------------------------------------------------------------------- | --------------- | --------------------------- |
| `address` | `string`                                                                                     | Query, Required | A valid testnet BTC address |
| `body`    | [`*models.ExtendedV1FaucetsBtcRequest`](../../doc/models/extended-v1-faucets-btc-request.md) | Body, Optional  | -                           |

## Response Type

[`models.RunFaucetResponse`](../../doc/models/run-faucet-response.md)

## Example Usage

```go
ctx := context.Background()
address := "2N4M94S1ZPt8HfxydXzL2P7qyzgVq7MHWts"

body := models.ExtendedV1FaucetsBtcRequest{
    Address: models.ToPointer("2N4M94S1ZPt8HfxydXzL2P7qyzgVq7MHWts"),
}

apiResponse, err := faucetsController.RunFaucetBtc(ctx, address, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description               | Exception Class                                                                                            |
| ---------------- | ------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| 403              | BTC Faucet not fully configured | [`ExtendedV1FaucetsBtc403ErrorException`](../../doc/models/extended-v1-faucets-btc-403-error-exception.md) |
| 500              | Faucet call failed              | `ApiError`                                                                                                 |
