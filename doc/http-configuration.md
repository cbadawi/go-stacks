# HttpConfiguration

The following parameters are configurable for the HttpConfiguration:

## Properties

| Name                 | Type                                                               | Description                                                                                     |
| -------------------- | ------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------- |
| `timeout`            | `float64`                                                          | Timeout in milliseconds.<br>_Default_: `0`                                                      |
| `transport`          | `http.RoundTripper`                                                | Establishes network connection and caches them for reuse.<br>_Default_: `http.DefaultTransport` |
| `retryConfiguration` | [`stacksblockchainapi.RetryConfiguration`](retry-configuration.md) | Configurations to retry requests.<br>_Default_: `DefaultRetryConfiguration()`                   |

The httpConfiguration can be initialized as follows:

```go
httpConfiguration := CreateHttpConfiguration(
    stacksblockchainapi.WithTimeout(0),
    stacksblockchainapi.WithTransport(http.DefaultTransport),
    stacksblockchainapi.WithRetryConfiguration(DefaultRetryConfiguration()),
)
```
