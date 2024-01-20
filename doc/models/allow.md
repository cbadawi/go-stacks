# Allow

Allow specifies supported Operation status, Operation types, and all possible error statuses. This Allow object is used by clients to validate the correctness of a Rosetta Server implementation. It is expected that these clients will error if they receive some response that contains any of the above information that is not specified here.

## Structure

`Allow`

## Fields

| Name                      | Type                                                                              | Tags     | Description                                                                                                                                                  |
| ------------------------- | --------------------------------------------------------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `OperationStatuses`       | [`[]models.RosettaOperationStatus`](../../doc/models/rosetta-operation-status.md) | Required | All Operation.Status this implementation supports. Any status that is returned during parsing that is not listed here will cause client validation to error. |
| `OperationTypes`          | `[]string`                                                                        | Required | All Operation.Type this implementation supports. Any type that is returned during parsing that is not listed here will cause client validation to error.     |
| `Errors`                  | [`[]models.RosettaErrorNoDetails`](../../doc/models/rosetta-error-no-details.md)  | Required | All Errors that this implementation could return. Any error that is returned during parsing that is not listed here will cause client validation to error.   |
| `HistoricalBalanceLookup` | `bool`                                                                            | Required | Any Rosetta implementation that supports querying the balance of an account at any height in the past should set this to true.                               |

## Example (as JSON)

```json
{
  "operation_statuses": [
    {
      "status": "status0",
      "successful": false
    }
  ],
  "operation_types": ["operation_types7"],
  "errors": [
    {
      "code": 144,
      "message": "message0",
      "retriable": false
    }
  ],
  "historical_balance_lookup": false
}
```
