# RetryConfiguration

The following parameters are configurable for the RetryConfiguration:

## Properties

| Name                     | Type            | Description                                                                                                         |
| ------------------------ | --------------- | ------------------------------------------------------------------------------------------------------------------- |
| `maxRetryAttempts`       | `int64`         | Maximum number of retries.<br>_Default_: `0`                                                                        |
| `retryOnTimeout`         | `bool`          | Whether to retry on request timeout.<br>_Default_: `true`                                                           |
| `retryInterval`          | `time.Duration` | Interval before next retry. Used in calculation of wait time for next request in case of failure.<br>_Default_: `1` |
| `maximumRetryWaitTime`   | `time.Duration` | Overall wait time for the requests getting retried.<br>_Default_: `0`                                               |
| `backoffFactor`          | `int64`         | Used in calculation of wait time for next request in case of failure.<br>_Default_: `2`                             |
| `httpStatusCodesToRetry` | `[]int64`       | Http status codes to retry against.<br>_Default_: `[]int64{408, 413, 429, 500, 502, 503, 504, 521, 522, 524}`       |
| `httpMethodsToRetry`     | `[]string`      | Http methods to retry against.<br>_Default_: `[]string{"GET", "PUT"}`                                               |

The retryConfiguration can be initialized as follows:

```go
retryConfiguration := CreateRetryConfiguration(
    stacks.WithMaxRetryAttempts(0),
    stacks.WithRetryOnTimeout(true),
    stacks.WithRetryInterval(1),
    stacks.WithMaximumRetryWaitTime(0),
    stacks.WithBackoffFactor(2),
    stacks.WithHttpStatusCodesToRetry([]int64{408, 413, 429, 500, 502, 503, 504, 521, 522, 524}),
    stacks.WithHttpMethodsToRetry([]string{"GET", "PUT"}),
)
```
