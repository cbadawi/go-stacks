# Address Stx Inbound List Response

GET request that returns a list of inbound STX transfers with a memo

## Structure

`AddressStxInboundListResponse`

## Fields

| Name      | Type                                                                      | Tags     | Description              |
| --------- | ------------------------------------------------------------------------- | -------- | ------------------------ |
| `Limit`   | `int`                                                                     | Required | **Constraints**: `<= 30` |
| `Offset`  | `int`                                                                     | Required | -                        |
| `Total`   | `int`                                                                     | Required | -                        |
| `Results` | [`[]models.InboundStxTransfer`](../../doc/models/inbound-stx-transfer.md) | Required | -                        |

## Example (as JSON)

```json
{
  "limit": 52,
  "offset": 148,
  "total": 146,
  "results": [
    {
      "sender": "sender2",
      "amount": "amount8",
      "memo": "memo0",
      "block_height": 13.32,
      "tx_id": "tx_id0",
      "transfer_type": "stx-transfer",
      "tx_index": 98.14
    }
  ]
}
```
