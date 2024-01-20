# Transaction Event 1

## Structure

`TransactionEvent1`

## Fields

| Name           | Type                                                                             | Tags               | Description                         |
| -------------- | -------------------------------------------------------------------------------- | ------------------ | ----------------------------------- |
| `EventIndex`   | `int`                                                                            | Required           | -                                   |
| `EventType`    | `string`                                                                         | Required, Constant | **Default**: `"smart_contract_log"` |
| `TxId`         | `string`                                                                         | Required           | -                                   |
| `ContractLog`  | [`models.ContractLog`](../../doc/models/contract-log.md)                         | Required           | -                                   |
| `StxLockEvent` | [`models.StxLockEvent`](../../doc/models/stx-lock-event.md)                      | Required           | -                                   |
| `Asset`        | [`models.TransactionEventAsset1`](../../doc/models/transaction-event-asset-1.md) | Required           | -                                   |

## Example (as JSON)

```json
{
  "event_index": 220,
  "event_type": "smart_contract_log",
  "tx_id": "tx_id8",
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
  },
  "asset": {
    "asset_event_type": "burn",
    "asset_id": "asset_id6",
    "sender": "sender2",
    "recipient": "recipient2",
    "amount": "amount2",
    "value": {
      "hex": "hex6",
      "repr": "repr2"
    },
    "memo": "memo4"
  }
}
```
