# Smart Contracts

```go
smartContractsController := client.SmartContractsController()
```

## Class Name

`SmartContractsController`

## Methods

- [Get Contract by Id](../../doc/controllers/smart-contracts.md#get-contract-by-id)
- [Get Contracts by Trait](../../doc/controllers/smart-contracts.md#get-contracts-by-trait)
- [Get Contract Events by Id](../../doc/controllers/smart-contracts.md#get-contract-events-by-id)
- [Get Contract Interface](../../doc/controllers/smart-contracts.md#get-contract-interface)
- [Get Contract Data Map Entry](../../doc/controllers/smart-contracts.md#get-contract-data-map-entry)
- [Get Contract Source](../../doc/controllers/smart-contracts.md#get-contract-source)
- [Call Read Only Function](../../doc/controllers/smart-contracts.md#call-read-only-function)

# Get Contract by Id

Retrieves details of a contract with a given `contract_id`

```go
GetContractById(
    ctx context.Context,
    contractId string,
    unanchored *bool) (
    models.ApiResponse[models.SmartContract7],
    error)
```

## Parameters

| Parameter    | Type     | Tags               | Description                                                             |
| ------------ | -------- | ------------------ | ----------------------------------------------------------------------- |
| `contractId` | `string` | Template, Required | Contract identifier formatted as `<contract_address>.<contract_name>`   |
| `unanchored` | `*bool`  | Query, Optional    | Include transaction data from unanchored (i.e. unconfirmed) microblocks |

## Response Type

[`models.SmartContract7`](../../doc/models/smart-contract-7.md)

## Example Usage

```go
ctx := context.Background()
contractId := "SP6P4EJF0VG8V0RB3TQQKJBHDQKEF6NVRD1KZE3C.satoshibles"
unanchored := true

apiResponse, err := smartContractsController.GetContractById(ctx, contractId, &unanchored)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description                | Exception Class |
| ---------------- | -------------------------------- | --------------- |
| 404              | Cannot find contract of given ID | `ApiError`      |

# Get Contracts by Trait

Retrieves a list of contracts based on the following traits listed in JSON format - functions, variables, maps, fungible tokens and non-fungible tokens

```go
GetContractsByTrait(
    ctx context.Context,
    traitAbi string,
    limit *int,
    offset *int) (
    models.ApiResponse[models.ContractListResponse],
    error)
```

## Parameters

| Parameter  | Type     | Tags            | Description                            |
| ---------- | -------- | --------------- | -------------------------------------- |
| `traitAbi` | `string` | Query, Required | JSON abi of the trait.                 |
| `limit`    | `*int`   | Query, Optional | max number of contracts fetch          |
| `offset`   | `*int`   | Query, Optional | index of first contract event to fetch |

## Response Type

[`models.ContractListResponse`](../../doc/models/contract-list-response.md)

## Example Usage

```go
ctx := context.Background()
traitAbi := "trait_abi8"
offset := 42000

apiResponse, err := smartContractsController.GetContractsByTrait(ctx, traitAbi, nil, &offset)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Contract Events by Id

Retrieves a list of events that have been triggered by a given `contract_id`

```go
GetContractEventsById(
    ctx context.Context,
    contractId string,
    limit *int,
    offset *int,
    unanchored *bool) (
    models.ApiResponse[models.TransactionEvent],
    error)
```

## Parameters

| Parameter    | Type     | Tags               | Description                                                             |
| ------------ | -------- | ------------------ | ----------------------------------------------------------------------- |
| `contractId` | `string` | Template, Required | Contract identifier formatted as `<contract_address>.<contract_name>`   |
| `limit`      | `*int`   | Query, Optional    | max number of contract events to fetch                                  |
| `offset`     | `*int`   | Query, Optional    | index of first contract event to fetch                                  |
| `unanchored` | `*bool`  | Query, Optional    | Include transaction data from unanchored (i.e. unconfirmed) microblocks |

## Response Type

[`models.TransactionEvent`](../../doc/models/transaction-event.md)

## Example Usage

```go
ctx := context.Background()
contractId := "SP6P4EJF0VG8V0RB3TQQKJBHDQKEF6NVRD1KZE3C.satoshibles"
offset := 42000
unanchored := true

apiResponse, err := smartContractsController.GetContractEventsById(ctx, contractId, nil, &offset, &unanchored)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Contract Interface

Retrieves a contract interface with a given `contract_address` and `contract name`

```go
GetContractInterface(
    ctx context.Context,
    contractAddress string,
    contractName string,
    tip *string) (
    models.ApiResponse[models.ContractInterfaceResponse],
    error)
```

## Parameters

| Parameter         | Type      | Tags               | Description                        |
| ----------------- | --------- | ------------------ | ---------------------------------- |
| `contractAddress` | `string`  | Template, Required | Stacks address                     |
| `contractName`    | `string`  | Template, Required | Contract name                      |
| `tip`             | `*string` | Query, Optional    | The Stacks chain tip to query from |

## Response Type

[`models.ContractInterfaceResponse`](../../doc/models/contract-interface-response.md)

## Example Usage

```go
ctx := context.Background()
contractAddress := "SP6P4EJF0VG8V0RB3TQQKJBHDQKEF6NVRD1KZE3C"
contractName := "satoshibles"

apiResponse, err := smartContractsController.GetContractInterface(ctx, contractAddress, contractName, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Contract Data Map Entry

Attempt to fetch data from a contract data map. The contract is identified with Stacks Address `contract_address` and Contract Name `contract_address` in the URL path. The map is identified with [Map Name].

The key to lookup in the map is supplied via the POST body. This should be supplied as the hex string serialization of the key (which should be a Clarity value). Note, this is a JSON string atom.

In the response, `data` is the hex serialization of the map response. Note that map responses are Clarity option types, for non-existent values, this is a serialized none, and for all other responses, it is a serialized (some ...) object.

```go
GetContractDataMapEntry(
    ctx context.Context,
    contractAddress string,
    contractName string,
    mapName string,
    body string,
    proof *int,
    tip *string) (
    models.ApiResponse[models.MapEntryResponse],
    error)
```

## Parameters

| Parameter         | Type      | Tags               | Description                                                                  |
| ----------------- | --------- | ------------------ | ---------------------------------------------------------------------------- |
| `contractAddress` | `string`  | Template, Required | Stacks address                                                               |
| `contractName`    | `string`  | Template, Required | Contract name                                                                |
| `mapName`         | `string`  | Template, Required | Map name                                                                     |
| `body`            | `string`  | Body, Required     | Hex string serialization of the lookup key (which should be a Clarity value) |
| `proof`           | `*int`    | Query, Optional    | Returns object without the proof field when set to 0                         |
| `tip`             | `*string` | Query, Optional    | The Stacks chain tip to query from                                           |

## Response Type

[`models.MapEntryResponse`](../../doc/models/map-entry-response.md)

## Example Usage

```go
ctx := context.Background()
contractAddress := "SPSCWDV3RKV5ZRN1FQD84YE1NQFEDJ9R1F4DYQ11"
contractName := "newyorkcitycoin-core-v2"
mapName := "approved-contracts"
body := "{\"$ref\":\"./entities/contracts/get-specific-data-map-inside-contract.example.json\"}"

apiResponse, err := smartContractsController.GetContractDataMapEntry(ctx, contractAddress, contractName, mapName, body, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description       | Exception Class |
| ---------------- | ----------------------- | --------------- |
| 400              | Failed loading data map | `ApiError`      |

# Get Contract Source

Retrieves the Clarity source code of a given contract, along with the block height it was published in, and the MARF proof for the data

```go
GetContractSource(
    ctx context.Context,
    contractAddress string,
    contractName string,
    proof *int,
    tip *string) (
    models.ApiResponse[models.ContractSourceResponse],
    error)
```

## Parameters

| Parameter         | Type      | Tags               | Description                                        |
| ----------------- | --------- | ------------------ | -------------------------------------------------- |
| `contractAddress` | `string`  | Template, Required | Stacks address                                     |
| `contractName`    | `string`  | Template, Required | Contract name                                      |
| `proof`           | `*int`    | Query, Optional    | Returns object without the proof field if set to 0 |
| `tip`             | `*string` | Query, Optional    | The Stacks chain tip to query from                 |

## Response Type

[`models.ContractSourceResponse`](../../doc/models/contract-source-response.md)

## Example Usage

```go
ctx := context.Background()
contractAddress := "SP6P4EJF0VG8V0RB3TQQKJBHDQKEF6NVRD1KZE3C"
contractName := "satoshibles"

apiResponse, err := smartContractsController.GetContractSource(ctx, contractAddress, contractName, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Call Read Only Function

Call a read-only public function on a given smart contract.

The smart contract and function are specified using the URL path. The arguments and the simulated tx-sender are supplied via the POST body in the following JSON format:

```go
CallReadOnlyFunction(
    ctx context.Context,
    contractAddress string,
    contractName string,
    functionName string,
    body models.ReadOnlyFunctionArgs,
    tip *string) (
    models.ApiResponse[models.ReadOnlyFunctionSuccessResponse],
    error)
```

## Parameters

| Parameter         | Type                                                                         | Tags               | Description                                                                                                                                                                       |
| ----------------- | ---------------------------------------------------------------------------- | ------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `contractAddress` | `string`                                                                     | Template, Required | Stacks address                                                                                                                                                                    |
| `contractName`    | `string`                                                                     | Template, Required | Contract name                                                                                                                                                                     |
| `functionName`    | `string`                                                                     | Template, Required | Function name                                                                                                                                                                     |
| `body`            | [`models.ReadOnlyFunctionArgs`](../../doc/models/read-only-function-args.md) | Body, Required     | map of arguments and the simulated tx-sender where sender is either a Contract identifier or a normal Stacks address, and arguments is an array of hex serialized Clarity values. |
| `tip`             | `*string`                                                                    | Query, Optional    | The Stacks chain tip to query from                                                                                                                                                |

## Response Type

[`models.ReadOnlyFunctionSuccessResponse`](../../doc/models/read-only-function-success-response.md)

## Example Usage

```go
ctx := context.Background()
contractAddress := "SP187Y7NRSG3T9Z9WTSWNEN3XRV1YSJWS81C7JKV7"
contractName := "imaginary-friends-zebras"
functionName := "get-token-uri"

body := models.ReadOnlyFunctionArgs{
    Sender:    "sender2",
    Arguments: []string{"arguments3", "arguments4", "arguments5"},
}

apiResponse, err := smartContractsController.CallReadOnlyFunction(ctx, contractAddress, contractName, functionName, &body, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```
