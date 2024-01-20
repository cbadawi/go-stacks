# Transaction Results

GET request that returns transactions

## Structure

`TransactionResults`

## Fields

| Name      | Type                                                      | Tags     | Description                                                       |
| --------- | --------------------------------------------------------- | -------- | ----------------------------------------------------------------- |
| `Limit`   | `int`                                                     | Required | The number of transactions to return<br>**Constraints**: `<= 200` |
| `Offset`  | `int`                                                     | Required | The number to transactions to skip (starting at `0`)              |
| `Total`   | `int`                                                     | Required | The number of transactions available                              |
| `Results` | [`[]models.Transaction`](../../doc/models/transaction.md) | Required | -                                                                 |

## Example (as JSON)

```json
{
  "limit": 120,
  "offset": 216,
  "total": 42,
  "results": [
    {
      "tx_id": "tx_id0",
      "nonce": 240,
      "fee_rate": "fee_rate8",
      "sender_address": "sender_address8",
      "sponsor_nonce": 228
    }
  ]
}
```
