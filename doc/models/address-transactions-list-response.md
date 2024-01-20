# Address Transactions List Response

GET request that returns account transactions

## Structure

`AddressTransactionsListResponse`

## Fields

| Name      | Type                                            | Tags     | Description              |
| --------- | ----------------------------------------------- | -------- | ------------------------ |
| `Limit`   | `int`                                           | Required | **Constraints**: `<= 30` |
| `Offset`  | `int`                                           | Required | -                        |
| `Total`   | `int`                                           | Required | -                        |
| `Results` | [`[]models.Result`](../../doc/models/result.md) | Required | -                        |

## Example (as JSON)

```json
{
  "limit": 212,
  "offset": 52,
  "total": 206,
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
