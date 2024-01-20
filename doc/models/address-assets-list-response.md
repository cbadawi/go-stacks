# Address Assets List Response

GET request that returns address assets

## Structure

`AddressAssetsListResponse`

## Fields

| Name      | Type                                                                 | Tags     | Description              |
| --------- | -------------------------------------------------------------------- | -------- | ------------------------ |
| `Limit`   | `int`                                                                | Required | **Constraints**: `<= 30` |
| `Offset`  | `int`                                                                | Required | -                        |
| `Total`   | `int`                                                                | Required | -                        |
| `Results` | [`[]models.TransactionEvent`](../../doc/models/transaction-event.md) | Required | -                        |

## Example (as JSON)

```json
{
  "limit": 158,
  "offset": 62,
  "total": 64,
  "results": [
    {
      "event_index": 124,
      "event_type": "smart_contract_log",
      "tx_id": "tx_id0",
      "contract_log": {
        "contract_id": "contract_id0",
        "topic": "topic4",
        "value": {
          "hex": "hex6",
          "repr": "repr2"
        }
      },
      "stx_lock_event": {
        "locked_amount": "locked_amount0",
        "unlock_height": 234,
        "locked_address": "locked_address2"
      }
    }
  ]
}
```
