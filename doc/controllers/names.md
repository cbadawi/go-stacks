# Names

Read-only endpoints realted to the Blockchain Naming System on Stacks

Stacks Documentation - Blockchain Naming System: [https://docs.stacks.co/build-apps/references/bns](https://docs.stacks.co/build-apps/references/bns)

```go
namesController := client.NamesController()
```

## Class Name

`NamesController`

## Methods

- [Get Namespace Price](../../doc/controllers/names.md#get-namespace-price)
- [Get Name Price](../../doc/controllers/names.md#get-name-price)
- [Get All Namespaces](../../doc/controllers/names.md#get-all-namespaces)
- [Get Namespace Names](../../doc/controllers/names.md#get-namespace-names)
- [Get All Names](../../doc/controllers/names.md#get-all-names)
- [Get Name Info](../../doc/controllers/names.md#get-name-info)
- [Fetch Subdomains List for Name](../../doc/controllers/names.md#fetch-subdomains-list-for-name)
- [Fetch Zone File](../../doc/controllers/names.md#fetch-zone-file)
- [Get Historical Zone File](../../doc/controllers/names.md#get-historical-zone-file)
- [Get Names Owned by Address](../../doc/controllers/names.md#get-names-owned-by-address)

# Get Namespace Price

Retrieves the price of a namespace. The `amount` given will be in the smallest possible units of the currency.

```go
GetNamespacePrice(
    ctx context.Context,
    tld string) (
    models.ApiResponse[models.BnsGetNamespacePriceResponse],
    error)
```

## Parameters

| Parameter | Type     | Tags               | Description                      |
| --------- | -------- | ------------------ | -------------------------------- |
| `tld`     | `string` | Template, Required | the namespace to fetch price for |

## Response Type

[`models.BnsGetNamespacePriceResponse`](../../doc/models/bns-get-namespace-price-response.md)

## Example Usage

```go
ctx := context.Background()
tld := "id"

apiResponse, err := namesController.GetNamespacePrice(ctx, tld)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Name Price

Retrieves the price of a name. The `amount` given will be in the smallest possible units of the currency.

```go
GetNamePrice(
    ctx context.Context,
    name string) (
    models.ApiResponse[models.BnsGetNamePriceResponse],
    error)
```

## Parameters

| Parameter | Type     | Tags               | Description                             |
| --------- | -------- | ------------------ | --------------------------------------- |
| `name`    | `string` | Template, Required | the name to query price information for |

## Response Type

[`models.BnsGetNamePriceResponse`](../../doc/models/bns-get-name-price-response.md)

## Example Usage

```go
ctx := context.Background()
name := "muneeb.id"

apiResponse, err := namesController.GetNamePrice(ctx, name)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get All Namespaces

Retrieves a list of all namespaces known to the node.

```go
GetAllNamespaces(
    ctx context.Context) (
    models.ApiResponse[models.BnsGetAllNamespacesResponse],
    error)
```

## Response Type

[`models.BnsGetAllNamespacesResponse`](../../doc/models/bns-get-all-namespaces-response.md)

## Example Usage

```go
ctx := context.Background()
apiResponse, err := namesController.GetAllNamespaces(ctx)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Get Namespace Names

Retrieves a list of names within a given namespace.

```go
GetNamespaceNames(
    ctx context.Context,
    tld string,
    page *int) (
    models.ApiResponse[[]string],
    error)
```

## Parameters

| Parameter | Type     | Tags               | Description                                                                                                                         |
| --------- | -------- | ------------------ | ----------------------------------------------------------------------------------------------------------------------------------- |
| `tld`     | `string` | Template, Required | the namespace to fetch names from.                                                                                                  |
| `page`    | `*int`   | Query, Optional    | namespace values are defaulted to page 1 with 100 results. You can query specific page results by using the 'page' query parameter. |

## Response Type

`[]string`

## Example Usage

```go
ctx := context.Background()
tld := "id"
page := 22

apiResponse, err := namesController.GetNamespaceNames(ctx, tld, &page)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                |
| ---------------- | ----------------- | -------------------------------------------------------------- |
| 400              | Error             | [`BnsErrorException`](../../doc/models/bns-error-exception.md) |
| 404              | Error             | [`BnsErrorException`](../../doc/models/bns-error-exception.md) |

# Get All Names

Retrieves a list of all names known to the node.

```go
GetAllNames(
    ctx context.Context,
    page *int) (
    models.ApiResponse[[]string],
    error)
```

## Parameters

| Parameter | Type   | Tags            | Description                                                                                                              |
| --------- | ------ | --------------- | ------------------------------------------------------------------------------------------------------------------------ |
| `page`    | `*int` | Query, Optional | names are defaulted to page 1 with 100 results. You can query specific page results by using the 'page' query parameter. |

## Response Type

`[]string`

## Example Usage

```go
ctx := context.Background()
page := 22

apiResponse, err := namesController.GetAllNames(ctx, &page)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                |
| ---------------- | ----------------- | -------------------------------------------------------------- |
| 400              | Error             | [`BnsErrorException`](../../doc/models/bns-error-exception.md) |

# Get Name Info

Retrieves details of a given name including the `address`, `status` and last transaction id - `last_txid`.

```go
GetNameInfo(
    ctx context.Context,
    name string) (
    models.ApiResponse[models.BnsGetNameInfoResponse],
    error)
```

## Parameters

| Parameter | Type     | Tags               | Description          |
| --------- | -------- | ------------------ | -------------------- |
| `name`    | `string` | Template, Required | fully-qualified name |

## Response Type

[`models.BnsGetNameInfoResponse`](../../doc/models/bns-get-name-info-response.md)

## Example Usage

```go
ctx := context.Background()
name := "muneeb.id"

apiResponse, err := namesController.GetNameInfo(ctx, name)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                |
| ---------------- | ----------------- | -------------------------------------------------------------- |
| 400              | Error             | [`BnsErrorException`](../../doc/models/bns-error-exception.md) |
| 404              | Error             | [`BnsErrorException`](../../doc/models/bns-error-exception.md) |

# Fetch Subdomains List for Name

Retrieves the list of subdomains for a specific name

```go
FetchSubdomainsListForName(
    ctx context.Context,
    name string) (
    models.ApiResponse[[]string],
    error)
```

## Parameters

| Parameter | Type     | Tags               | Description          |
| --------- | -------- | ------------------ | -------------------- |
| `name`    | `string` | Template, Required | fully-qualified name |

## Response Type

`[]string`

## Example Usage

```go
ctx := context.Background()
name := "id.blockstack"

apiResponse, err := namesController.FetchSubdomainsListForName(ctx, name)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

# Fetch Zone File

Retrieves a user’s raw zone file. This only works for RFC-compliant zone files. This method returns an error for names that have non-standard zone files.

```go
FetchZoneFile(
    ctx context.Context,
    name string) (
    models.ApiResponse[models.BnsFetchFileZoneResponse2],
    error)
```

## Parameters

| Parameter | Type     | Tags               | Description          |
| --------- | -------- | ------------------ | -------------------- |
| `name`    | `string` | Template, Required | fully-qualified name |

## Response Type

[`models.BnsFetchFileZoneResponse2`](../../doc/models/bns-fetch-file-zone-response-2.md)

## Example Usage

```go
ctx := context.Background()
name := "bar.test"

apiResponse, err := namesController.FetchZoneFile(ctx, name)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                |
| ---------------- | ----------------- | -------------------------------------------------------------- |
| 400              | Error             | [`BnsErrorException`](../../doc/models/bns-error-exception.md) |
| 404              | Error             | [`BnsErrorException`](../../doc/models/bns-error-exception.md) |

# Get Historical Zone File

Retrieves the historical zonefile specified by the username and zone hash.

```go
GetHistoricalZoneFile(
    ctx context.Context,
    name string,
    zoneFileHash string) (
    models.ApiResponse[models.BnsFetchHistoricalZoneFileResponse2],
    error)
```

## Parameters

| Parameter      | Type     | Tags               | Description          |
| -------------- | -------- | ------------------ | -------------------- |
| `name`         | `string` | Template, Required | fully-qualified name |
| `zoneFileHash` | `string` | Template, Required | zone file hash       |

## Response Type

[`models.BnsFetchHistoricalZoneFileResponse2`](../../doc/models/bns-fetch-historical-zone-file-response-2.md)

## Example Usage

```go
ctx := context.Background()
name := "muneeb.id"
zoneFileHash := "b100a68235244b012854a95f9114695679002af9"

apiResponse, err := namesController.GetHistoricalZoneFile(ctx, name, zoneFileHash)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                |
| ---------------- | ----------------- | -------------------------------------------------------------- |
| 400              | Error             | [`BnsErrorException`](../../doc/models/bns-error-exception.md) |
| 404              | Error             | [`BnsErrorException`](../../doc/models/bns-error-exception.md) |

# Get Names Owned by Address

Retrieves a list of names owned by the address provided.

```go
GetNamesOwnedByAddress(
    ctx context.Context,
    blockchain string,
    address string) (
    models.ApiResponse[models.BnsNamesOwnByAddressResponse],
    error)
```

## Parameters

| Parameter    | Type     | Tags               | Description                            |
| ------------ | -------- | ------------------ | -------------------------------------- |
| `blockchain` | `string` | Template, Required | the layer-1 blockchain for the address |
| `address`    | `string` | Template, Required | the address to lookup                  |

## Response Type

[`models.BnsNamesOwnByAddressResponse`](../../doc/models/bns-names-own-by-address-response.md)

## Example Usage

```go
ctx := context.Background()
blockchain := "bitcoin"
address := "1QJQxDas5JhdiXhEbNS14iNjr8auFT96GP"

apiResponse, err := namesController.GetNamesOwnedByAddress(ctx, blockchain, address)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class                                                |
| ---------------- | ----------------- | -------------------------------------------------------------- |
| 404              | Error             | [`BnsErrorException`](../../doc/models/bns-error-exception.md) |
