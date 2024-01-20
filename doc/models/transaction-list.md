# Transaction List

## Structure

`TransactionList`

## Fields

| Name     | Type                                              | Tags     | Description |
| -------- | ------------------------------------------------- | -------- | ----------- |
| `Found`  | `*bool`                                           | Optional | -           |
| `Result` | [`*models.Result2`](../../doc/models/result-2.md) | Optional | -           |

## Example (as JSON)

```json
{
  "found": false,
  "result": {
    "tx_id": "tx_id0",
    "nonce": 46,
    "fee_rate": "fee_rate8",
    "sender_address": "sender_address8",
    "sponsor_nonce": 90,
    "sponsored": false
  }
}
```
