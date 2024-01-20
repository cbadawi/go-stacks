# Mempool Transaction List Response

GET request that returns transactions

## Structure

`MempoolTransactionListResponse`

## Fields

| Name      | Type                                                                     | Tags     | Description |
| --------- | ------------------------------------------------------------------------ | -------- | ----------- |
| `Limit`   | `int`                                                                    | Required | -           |
| `Offset`  | `int`                                                                    | Required | -           |
| `Total`   | `int`                                                                    | Required | -           |
| `Results` | [`[]models.MempoolTransaction`](../../doc/models/mempool-transaction.md) | Required | -           |

## Example (as JSON)

```json
{
  "limit": 234,
  "offset": 118,
  "total": 140,
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
