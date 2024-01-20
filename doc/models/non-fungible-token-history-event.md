# Non Fungible Token History Event

Describes an event from the history of a Non-Fungible Token

## Structure

`NonFungibleTokenHistoryEvent`

## Fields

| Name             | Type                                                     | Tags     | Description                                              |
| ---------------- | -------------------------------------------------------- | -------- | -------------------------------------------------------- |
| `Sender`         | `Optional[string]`                                       | Optional | -                                                        |
| `Recipient`      | `*string`                                                | Optional | -                                                        |
| `EventIndex`     | `*int`                                                   | Optional | -                                                        |
| `AssetEventType` | `*string`                                                | Optional | -                                                        |
| `TxId`           | `*string`                                                | Optional | -                                                        |
| `Tx`             | [`*models.Transaction`](../../doc/models/transaction.md) | Optional | Describes all transaction types on Stacks 2.0 blockchain |

## Example (as JSON)

```json
{
  "sender": "sender4",
  "recipient": "recipient4",
  "event_index": 0,
  "asset_event_type": "asset_event_type4",
  "tx_id": "tx_id6"
}
```
