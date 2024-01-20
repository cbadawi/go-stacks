# Transaction Event

## Structure

`TransactionEvent`

## Fields

| Name           | Type                                                                              | Tags     | Description |
| -------------- | --------------------------------------------------------------------------------- | -------- | ----------- |
| `EventIndex`   | `*int`                                                                            | Optional | -           |
| `EventType`    | [`*models.EventTypeEnum`](../../doc/models/event-type-enum.md)                    | Optional | -           |
| `TxId`         | `*string`                                                                         | Optional | -           |
| `ContractLog`  | [`*models.ContractLog`](../../doc/models/contract-log.md)                         | Optional | -           |
| `StxLockEvent` | [`*models.StxLockEvent`](../../doc/models/stx-lock-event.md)                      | Optional | -           |
| `Asset`        | [`*models.TransactionEventAsset1`](../../doc/models/transaction-event-asset-1.md) | Optional | -           |

## Example (as JSON)

```json
{
  "event_index": 154,
  "event_type": "smart_contract_log",
  "tx_id": "tx_id6",
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
```
