# Rosetta

Endpoints to support the Rosetta API open blockchain standard

Hiro Documentation - Rosetta Support: [https://docs.hiro.so/get-started/stacks-blockchain-api#rosetta-support](https://docs.hiro.so/get-started/stacks-blockchain-api#rosetta-support)

```go
rosettaController := client.RosettaController()
```

## Class Name

`RosettaController`

## Methods

- [Rosetta Network List](../../doc/controllers/rosetta.md#rosetta-network-list)
- [Rosetta Network Options](../../doc/controllers/rosetta.md#rosetta-network-options)
- [Rosetta Network Status](../../doc/controllers/rosetta.md#rosetta-network-status)
- [Rosetta Account Balance](../../doc/controllers/rosetta.md#rosetta-account-balance)
- [Rosetta Block](../../doc/controllers/rosetta.md#rosetta-block)
- [Rosetta Block Transaction](../../doc/controllers/rosetta.md#rosetta-block-transaction)
- [Rosetta Mempool](../../doc/controllers/rosetta.md#rosetta-mempool)
- [Rosetta Mempool Transaction](../../doc/controllers/rosetta.md#rosetta-mempool-transaction)
- [Rosetta Construction Derive](../../doc/controllers/rosetta.md#rosetta-construction-derive)
- [Rosetta Construction Hash](../../doc/controllers/rosetta.md#rosetta-construction-hash)
- [Rosetta Construction Metadata](../../doc/controllers/rosetta.md#rosetta-construction-metadata)
- [Rosetta Construction Parse](../../doc/controllers/rosetta.md#rosetta-construction-parse)
- [Rosetta Construction Preprocess](../../doc/controllers/rosetta.md#rosetta-construction-preprocess)
- [Rosetta Construction Submit](../../doc/controllers/rosetta.md#rosetta-construction-submit)
- [Rosetta Construction Payloads](../../doc/controllers/rosetta.md#rosetta-construction-payloads)
- [Rosetta Construction Combine](../../doc/controllers/rosetta.md#rosetta-construction-combine)

# Rosetta Network List

Retrieves a list of NetworkIdentifiers that the Rosetta server supports.

```go
RosettaNetworkList(
    ctx context.Context) (
    models.ApiResponse[models.RosettaNetworkListResponse],
    error)
```

## Response Type

[`models.RosettaNetworkListResponse`](../../doc/models/rosetta-network-list-response.md)

## Example Usage

```go
ctx := context.Background()
apiResponse, err := rosettaController.RosettaNetworkList(ctx)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                        |
| ---------------- | ----------------- | ---------------------------------------------------------------------- |
| 400              | Error             | [`RosettaErrorException`](../../doc/models/rosetta-error-exception.md) |

# Rosetta Network Options

Retrieves the version information and allowed network-specific types for a NetworkIdentifier.
Any NetworkIdentifier returned by /network/list should be accessible here.
Because options are retrievable in the context of a NetworkIdentifier, it is possible to define unique options for each network.

```go
RosettaNetworkOptions(
    ctx context.Context,
    body models.RosettaOptionsRequest) (
    models.ApiResponse[models.RosettaNetworkOptionsResponse],
    error)
```

## Parameters

| Parameter | Type                                                                          | Tags           | Description |
| --------- | ----------------------------------------------------------------------------- | -------------- | ----------- |
| `body`    | [`models.RosettaOptionsRequest`](../../doc/models/rosetta-options-request.md) | Body, Required | -           |

## Response Type

[`models.RosettaNetworkOptionsResponse`](../../doc/models/rosetta-network-options-response.md)

## Example Usage

```go
ctx := context.Background()

bodyNetworkIdentifier := models.NetworkIdentifier{
    Blockchain:           "blockchain2",
    Network:              "network6",
}

body := models.RosettaOptionsRequest{
    NetworkIdentifier: bodyNetworkIdentifier,
}

apiResponse, err := rosettaController.RosettaNetworkOptions(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                        |
| ---------------- | ----------------- | ---------------------------------------------------------------------- |
| 400              | Error             | [`RosettaErrorException`](../../doc/models/rosetta-error-exception.md) |

# Rosetta Network Status

Retrieves the current status of the network requested.
Any NetworkIdentifier returned by /network/list should be accessible here.

```go
RosettaNetworkStatus(
    ctx context.Context,
    body models.RosettaStatusRequest) (
    models.ApiResponse[models.RosettaNetworkStatusResponse],
    error)
```

## Parameters

| Parameter | Type                                                                        | Tags           | Description |
| --------- | --------------------------------------------------------------------------- | -------------- | ----------- |
| `body`    | [`models.RosettaStatusRequest`](../../doc/models/rosetta-status-request.md) | Body, Required | -           |

## Response Type

[`models.RosettaNetworkStatusResponse`](../../doc/models/rosetta-network-status-response.md)

## Example Usage

```go
ctx := context.Background()

bodyNetworkIdentifier := models.NetworkIdentifier{
    Blockchain:           "blockchain2",
    Network:              "network6",
}

body := models.RosettaStatusRequest{
    NetworkIdentifier: bodyNetworkIdentifier,
}

apiResponse, err := rosettaController.RosettaNetworkStatus(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                        |
| ---------------- | ----------------- | ---------------------------------------------------------------------- |
| 400              | Error             | [`RosettaErrorException`](../../doc/models/rosetta-error-exception.md) |

# Rosetta Account Balance

An AccountBalanceRequest is utilized to make a balance request on the /account/balance endpoint.
If the block_identifier is populated, a historical balance query should be performed.

```go
RosettaAccountBalance(
    ctx context.Context,
    body models.RosettaAccountBalanceRequest) (
    models.ApiResponse[models.RosettaAccountBalanceResponse],
    error)
```

## Parameters

| Parameter | Type                                                                                         | Tags           | Description |
| --------- | -------------------------------------------------------------------------------------------- | -------------- | ----------- |
| `body`    | [`models.RosettaAccountBalanceRequest`](../../doc/models/rosetta-account-balance-request.md) | Body, Required | -           |

## Response Type

[`models.RosettaAccountBalanceResponse`](../../doc/models/rosetta-account-balance-response.md)

## Example Usage

```go
ctx := context.Background()

bodyNetworkIdentifier := models.NetworkIdentifier{
    Blockchain:           "blockchain2",
    Network:              "network6",
}

bodyAccountIdentifier := models.RosettaAccount{
    Address:    "address6",
}

body := models.RosettaAccountBalanceRequest{
    NetworkIdentifier: bodyNetworkIdentifier,
    AccountIdentifier: bodyAccountIdentifier,
}

apiResponse, err := rosettaController.RosettaAccountBalance(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                        |
| ---------------- | ----------------- | ---------------------------------------------------------------------- |
| 400              | Error             | [`RosettaErrorException`](../../doc/models/rosetta-error-exception.md) |

# Rosetta Block

Retrieves the Block information for a given block identifier including a list of all transactions in the block.

```go
RosettaBlock(
    ctx context.Context,
    body models.RosettaBlockRequest) (
    models.ApiResponse[models.RosettaBlockResponse],
    error)
```

## Parameters

| Parameter | Type                                                                      | Tags           | Description |
| --------- | ------------------------------------------------------------------------- | -------------- | ----------- |
| `body`    | [`models.RosettaBlockRequest`](../../doc/models/rosetta-block-request.md) | Body, Required | -           |

## Response Type

[`models.RosettaBlockResponse`](../../doc/models/rosetta-block-response.md)

## Example Usage

```go
ctx := context.Background()

bodyNetworkIdentifier := models.NetworkIdentifier{
    Blockchain:           "blockchain2",
    Network:              "network6",
}

bodyBlockIdentifier := models.RosettaPartialBlockIdentifier{
}

body := models.RosettaBlockRequest{
    NetworkIdentifier: bodyNetworkIdentifier,
    BlockIdentifier:   bodyBlockIdentifier,
}

apiResponse, err := rosettaController.RosettaBlock(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                        |
| ---------------- | ----------------- | ---------------------------------------------------------------------- |
| 400              | Error             | [`RosettaErrorException`](../../doc/models/rosetta-error-exception.md) |

# Rosetta Block Transaction

Retrieves a Transaction included in a block that is not returned in a BlockResponse.

```go
RosettaBlockTransaction(
    ctx context.Context,
    body models.RosettaBlockTransactionRequest) (
    models.ApiResponse[models.RosettaBlockTransactionResponse],
    error)
```

## Parameters

| Parameter | Type                                                                                             | Tags           | Description |
| --------- | ------------------------------------------------------------------------------------------------ | -------------- | ----------- |
| `body`    | [`models.RosettaBlockTransactionRequest`](../../doc/models/rosetta-block-transaction-request.md) | Body, Required | -           |

## Response Type

[`models.RosettaBlockTransactionResponse`](../../doc/models/rosetta-block-transaction-response.md)

## Example Usage

```go
ctx := context.Background()

bodyNetworkIdentifier := models.NetworkIdentifier{
    Blockchain:           "blockchain2",
    Network:              "network6",
}

bodyBlockIdentifier := models.RosettaPartialBlockIdentifier{
}

bodyTransactionIdentifier := models.TransactionIdentifier{
    Hash: "hash6",
}

body := models.RosettaBlockTransactionRequest{
    NetworkIdentifier:     bodyNetworkIdentifier,
    BlockIdentifier:       bodyBlockIdentifier,
    TransactionIdentifier: bodyTransactionIdentifier,
}

apiResponse, err := rosettaController.RosettaBlockTransaction(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                        |
| ---------------- | ----------------- | ---------------------------------------------------------------------- |
| 400              | Error             | [`RosettaErrorException`](../../doc/models/rosetta-error-exception.md) |

# Rosetta Mempool

Retrieves a list of transactions currently in the mempool for a given network.

```go
RosettaMempool(
    ctx context.Context,
    body models.RosettaMempoolRequest) (
    models.ApiResponse[models.RosettaMempoolResponse],
    error)
```

## Parameters

| Parameter | Type                                                                          | Tags           | Description |
| --------- | ----------------------------------------------------------------------------- | -------------- | ----------- |
| `body`    | [`models.RosettaMempoolRequest`](../../doc/models/rosetta-mempool-request.md) | Body, Required | -           |

## Response Type

[`models.RosettaMempoolResponse`](../../doc/models/rosetta-mempool-response.md)

## Example Usage

```go
ctx := context.Background()

bodyNetworkIdentifier := models.NetworkIdentifier{
    Blockchain:           "blockchain2",
    Network:              "network6",
}

body := models.RosettaMempoolRequest{
    NetworkIdentifier: bodyNetworkIdentifier,
}

apiResponse, err := rosettaController.RosettaMempool(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                        |
| ---------------- | ----------------- | ---------------------------------------------------------------------- |
| 400              | Error             | [`RosettaErrorException`](../../doc/models/rosetta-error-exception.md) |

# Rosetta Mempool Transaction

Retrieves transaction details from the mempool for a given transaction id from a given network.

```go
RosettaMempoolTransaction(
    ctx context.Context,
    body models.RosettaMempoolTransactionRequest) (
    models.ApiResponse[models.RosettaMempoolTransactionResponse],
    error)
```

## Parameters

| Parameter | Type                                                                                                 | Tags           | Description |
| --------- | ---------------------------------------------------------------------------------------------------- | -------------- | ----------- |
| `body`    | [`models.RosettaMempoolTransactionRequest`](../../doc/models/rosetta-mempool-transaction-request.md) | Body, Required | -           |

## Response Type

[`models.RosettaMempoolTransactionResponse`](../../doc/models/rosetta-mempool-transaction-response.md)

## Example Usage

```go
ctx := context.Background()

bodyNetworkIdentifier := models.NetworkIdentifier{
    Blockchain:           "blockchain2",
    Network:              "network6",
}

bodyTransactionIdentifier := models.TransactionIdentifier{
    Hash: "hash6",
}

body := models.RosettaMempoolTransactionRequest{
    NetworkIdentifier:     bodyNetworkIdentifier,
    TransactionIdentifier: bodyTransactionIdentifier,
}

apiResponse, err := rosettaController.RosettaMempoolTransaction(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                        |
| ---------------- | ----------------- | ---------------------------------------------------------------------- |
| 400              | Error             | [`RosettaErrorException`](../../doc/models/rosetta-error-exception.md) |

# Rosetta Construction Derive

Retrieves the Account Identifier information based on a Public Key for a given network

```go
RosettaConstructionDerive(
    ctx context.Context,
    body models.RosettaConstructionDeriveRequest) (
    models.ApiResponse[models.RosettaConstructionDeriveResponse],
    error)
```

## Parameters

| Parameter | Type                                                                                                 | Tags           | Description |
| --------- | ---------------------------------------------------------------------------------------------------- | -------------- | ----------- |
| `body`    | [`models.RosettaConstructionDeriveRequest`](../../doc/models/rosetta-construction-derive-request.md) | Body, Required | -           |

## Response Type

[`models.RosettaConstructionDeriveResponse`](../../doc/models/rosetta-construction-derive-response.md)

## Example Usage

```go
ctx := context.Background()

bodyNetworkIdentifier := models.NetworkIdentifier{
    Blockchain:           "blockchain2",
    Network:              "network6",
}

bodyPublicKey := models.RosettaPublicKey{
    HexBytes:  "hex_bytes4",
    CurveType: models.CurveTypeEnum("secp256k1"),
}

body := models.RosettaConstructionDeriveRequest{
    NetworkIdentifier: bodyNetworkIdentifier,
    PublicKey:         bodyPublicKey,
}

apiResponse, err := rosettaController.RosettaConstructionDerive(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                        |
| ---------------- | ----------------- | ---------------------------------------------------------------------- |
| 400              | Error             | [`RosettaErrorException`](../../doc/models/rosetta-error-exception.md) |

# Rosetta Construction Hash

Retrieves the network-specific transaction hash for a signed transaction.

```go
RosettaConstructionHash(
    ctx context.Context,
    body models.RosettaConstructionHashRequest) (
    models.ApiResponse[models.RosettaConstructionHashResponse],
    error)
```

## Parameters

| Parameter | Type                                                                                             | Tags           | Description |
| --------- | ------------------------------------------------------------------------------------------------ | -------------- | ----------- |
| `body`    | [`models.RosettaConstructionHashRequest`](../../doc/models/rosetta-construction-hash-request.md) | Body, Required | -           |

## Response Type

[`models.RosettaConstructionHashResponse`](../../doc/models/rosetta-construction-hash-response.md)

## Example Usage

```go
ctx := context.Background()

bodyNetworkIdentifier := models.NetworkIdentifier{
    Blockchain:           "blockchain2",
    Network:              "network6",
}

body := models.RosettaConstructionHashRequest{
    SignedTransaction: "signed_transaction8",
    NetworkIdentifier: bodyNetworkIdentifier,
}

apiResponse, err := rosettaController.RosettaConstructionHash(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                        |
| ---------------- | ----------------- | ---------------------------------------------------------------------- |
| 400              | Error             | [`RosettaErrorException`](../../doc/models/rosetta-error-exception.md) |

# Rosetta Construction Metadata

To Do

```go
RosettaConstructionMetadata(
    ctx context.Context,
    body models.RosettaConstructionMetadataRequest) (
    models.ApiResponse[models.RosettaConstructionMetadataResponse],
    error)
```

## Parameters

| Parameter | Type                                                                                                     | Tags           | Description |
| --------- | -------------------------------------------------------------------------------------------------------- | -------------- | ----------- |
| `body`    | [`models.RosettaConstructionMetadataRequest`](../../doc/models/rosetta-construction-metadata-request.md) | Body, Required | -           |

## Response Type

[`models.RosettaConstructionMetadataResponse`](../../doc/models/rosetta-construction-metadata-response.md)

## Example Usage

```go
ctx := context.Background()

bodyNetworkIdentifier := models.NetworkIdentifier{
    Blockchain:           "blockchain2",
    Network:              "network6",
}

bodyOptions := models.RosettaOptions{
}

body := models.RosettaConstructionMetadataRequest{
    NetworkIdentifier: bodyNetworkIdentifier,
    Options:           bodyOptions,
}

apiResponse, err := rosettaController.RosettaConstructionMetadata(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                        |
| ---------------- | ----------------- | ---------------------------------------------------------------------- |
| 400              | Error             | [`RosettaErrorException`](../../doc/models/rosetta-error-exception.md) |

# Rosetta Construction Parse

TODO

```go
RosettaConstructionParse(
    ctx context.Context,
    body models.RosettaConstructionParseRequest) (
    models.ApiResponse[models.RosettaConstructionParseResponse],
    error)
```

## Parameters

| Parameter | Type                                                                                               | Tags           | Description |
| --------- | -------------------------------------------------------------------------------------------------- | -------------- | ----------- |
| `body`    | [`models.RosettaConstructionParseRequest`](../../doc/models/rosetta-construction-parse-request.md) | Body, Required | -           |

## Response Type

[`models.RosettaConstructionParseResponse`](../../doc/models/rosetta-construction-parse-response.md)

## Example Usage

```go
ctx := context.Background()

bodyNetworkIdentifier := models.NetworkIdentifier{
    Blockchain:           "blockchain2",
    Network:              "network6",
}

body := models.RosettaConstructionParseRequest{
    Signed:            false,
    Transaction:       "transaction2",
    NetworkIdentifier: bodyNetworkIdentifier,
}

apiResponse, err := rosettaController.RosettaConstructionParse(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                        |
| ---------------- | ----------------- | ---------------------------------------------------------------------- |
| 400              | Error             | [`RosettaErrorException`](../../doc/models/rosetta-error-exception.md) |

# Rosetta Construction Preprocess

TODO

```go
RosettaConstructionPreprocess(
    ctx context.Context,
    body models.RosettaConstructionPreprocessRequest) (
    models.ApiResponse[models.RosettaConstructionPreprocessResponse],
    error)
```

## Parameters

| Parameter | Type                                                                                                         | Tags           | Description |
| --------- | ------------------------------------------------------------------------------------------------------------ | -------------- | ----------- |
| `body`    | [`models.RosettaConstructionPreprocessRequest`](../../doc/models/rosetta-construction-preprocess-request.md) | Body, Required | -           |

## Response Type

[`models.RosettaConstructionPreprocessResponse`](../../doc/models/rosetta-construction-preprocess-response.md)

## Example Usage

```go
ctx := context.Background()

bodyNetworkIdentifier := models.NetworkIdentifier{
    Blockchain:           "blockchain2",
    Network:              "network6",
}

bodyOperations0OperationIdentifier := models.RosettaOperationIdentifier{
    Index:        254,
}

bodyOperations0 := models.RosettaOperation{
    Type:                "type8",
    OperationIdentifier: bodyOperations0OperationIdentifier,
}

bodyOperations := []models.RosettaOperation{bodyOperations0}
body := models.RosettaConstructionPreprocessRequest{
    NetworkIdentifier:      bodyNetworkIdentifier,
    Operations:             bodyOperations,
}

apiResponse, err := rosettaController.RosettaConstructionPreprocess(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                        |
| ---------------- | ----------------- | ---------------------------------------------------------------------- |
| 400              | Error             | [`RosettaErrorException`](../../doc/models/rosetta-error-exception.md) |

# Rosetta Construction Submit

Submit a pre-signed transaction to the node. The examples below are illustrative only. You'll need to use your wallet to generate actual values to use them in the request payload.

```go
RosettaConstructionSubmit(
    ctx context.Context,
    body models.RosettaConstructionSubmitRequest) (
    models.ApiResponse[models.RosettaConstructionSubmitResponse],
    error)
```

## Parameters

| Parameter | Type                                                                                                 | Tags           | Description |
| --------- | ---------------------------------------------------------------------------------------------------- | -------------- | ----------- |
| `body`    | [`models.RosettaConstructionSubmitRequest`](../../doc/models/rosetta-construction-submit-request.md) | Body, Required | -           |

## Response Type

[`models.RosettaConstructionSubmitResponse`](../../doc/models/rosetta-construction-submit-response.md)

## Example Usage

```go
ctx := context.Background()

bodyNetworkIdentifier := models.NetworkIdentifier{
    Blockchain:           "blockchain2",
    Network:              "network6",
}

body := models.RosettaConstructionSubmitRequest{
    SignedTransaction: "signed_transaction8",
    NetworkIdentifier: bodyNetworkIdentifier,
}

apiResponse, err := rosettaController.RosettaConstructionSubmit(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                        |
| ---------------- | ----------------- | ---------------------------------------------------------------------- |
| 400              | Error             | [`RosettaErrorException`](../../doc/models/rosetta-error-exception.md) |

# Rosetta Construction Payloads

Generate an unsigned transaction from operations and metadata. The examples below are illustrative only. You'll need to use your wallet to generate actual values to use them in the request payload.

```go
RosettaConstructionPayloads(
    ctx context.Context,
    body models.RosettaConstructionPayloadsRequest) (
    models.ApiResponse[models.RosettaConstructionPayloadResponse],
    error)
```

## Parameters

| Parameter | Type                                                                                                     | Tags           | Description |
| --------- | -------------------------------------------------------------------------------------------------------- | -------------- | ----------- |
| `body`    | [`models.RosettaConstructionPayloadsRequest`](../../doc/models/rosetta-construction-payloads-request.md) | Body, Required | -           |

## Response Type

[`models.RosettaConstructionPayloadResponse`](../../doc/models/rosetta-construction-payload-response.md)

## Example Usage

```go
ctx := context.Background()

bodyNetworkIdentifier := models.NetworkIdentifier{
    Blockchain:           "blockchain2",
    Network:              "network6",
}

bodyOperations0OperationIdentifier := models.RosettaOperationIdentifier{
    Index:        254,
}

bodyOperations0 := models.RosettaOperation{
    Type:                "type8",
    OperationIdentifier: bodyOperations0OperationIdentifier,
}

bodyOperations := []models.RosettaOperation{bodyOperations0}
body := models.RosettaConstructionPayloadsRequest{
    NetworkIdentifier: bodyNetworkIdentifier,
    Operations:        bodyOperations,
}

apiResponse, err := rosettaController.RosettaConstructionPayloads(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                        |
| ---------------- | ----------------- | ---------------------------------------------------------------------- |
| 400              | Error             | [`RosettaErrorException`](../../doc/models/rosetta-error-exception.md) |

# Rosetta Construction Combine

Take unsigned transaction and signature, combine both and return signed transaction. The examples below are illustrative only. You'll need to use your wallet to generate actual values to use them in the request payload.

```go
RosettaConstructionCombine(
    ctx context.Context,
    body models.RosettaConstructionCombineRequest) (
    models.ApiResponse[models.RosettaConstructionCombineResponse],
    error)
```

## Parameters

| Parameter | Type                                                                                                   | Tags           | Description |
| --------- | ------------------------------------------------------------------------------------------------------ | -------------- | ----------- |
| `body`    | [`models.RosettaConstructionCombineRequest`](../../doc/models/rosetta-construction-combine-request.md) | Body, Required | -           |

## Response Type

[`models.RosettaConstructionCombineResponse`](../../doc/models/rosetta-construction-combine-response.md)

## Example Usage

```go
ctx := context.Background()

bodyNetworkIdentifier := models.NetworkIdentifier{
    Blockchain:           "blockchain2",
    Network:              "network6",
}

bodySignatures0SigningPayload := models.SigningPayload{
    HexBytes:          "hex_bytes4",
}

bodySignatures0PublicKey := models.RosettaPublicKey{
    HexBytes:  "hex_bytes4",
    CurveType: models.CurveTypeEnum("secp256k1"),
}

bodySignatures0 := models.RosettaSignature{
    SignatureType:  models.SignatureTypeEnum(""),
    HexBytes:       "hex_bytes0",
    SigningPayload: bodySignatures0SigningPayload,
    PublicKey:      bodySignatures0PublicKey,
}

bodySignatures := []models.RosettaSignature{bodySignatures0}
body := models.RosettaConstructionCombineRequest{
    UnsignedTransaction: "unsigned_transaction8",
    NetworkIdentifier:   bodyNetworkIdentifier,
    Signatures:          bodySignatures,
}

apiResponse, err := rosettaController.RosettaConstructionCombine(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                        |
| ---------------- | ----------------- | ---------------------------------------------------------------------- |
| 400              | Error             | [`RosettaErrorException`](../../doc/models/rosetta-error-exception.md) |
