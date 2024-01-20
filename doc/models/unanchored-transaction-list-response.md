# Unanchored Transaction List Response

GET request that returns unanchored transactions

## Structure

`UnanchoredTransactionListResponse`

## Fields

| Name      | Type                                                      | Tags     | Description                                     |
| --------- | --------------------------------------------------------- | -------- | ----------------------------------------------- |
| `Total`   | `int`                                                     | Required | The number of unanchored transactions available |
| `Results` | [`[]models.Transaction`](../../doc/models/transaction.md) | Required | -                                               |

## Example (as JSON)

```json
{
  "total": 132,
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
