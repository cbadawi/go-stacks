# Post Core Node Transactions Error Exception

GET request that returns transactions

## Structure

`PostCoreNodeTransactionsErrorException`

## Fields

| Name         | Type          | Tags     | Description                   |
| ------------ | ------------- | -------- | ----------------------------- |
| `Error`      | `string`      | Required | The error                     |
| `Reason`     | `string`      | Required | The reason for the error      |
| `ReasonData` | `interface{}` | Required | More details about the reason |
| `Txid`       | `string`      | Required | The relevant transaction id   |

## Example (as JSON)

```json
{
  "error": "error8",
  "reason": "reason0",
  "reason_data": {
    "key1": "val1",
    "key2": "val2"
  },
  "txid": "txid8"
}
```
