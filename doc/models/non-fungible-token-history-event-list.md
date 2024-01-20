# Non Fungible Token History Event List

List of Non-Fungible Token history events

## Structure

`NonFungibleTokenHistoryEventList`

## Fields

| Name      | Type                                                                                            | Tags     | Description                                                 |
| --------- | ----------------------------------------------------------------------------------------------- | -------- | ----------------------------------------------------------- |
| `Limit`   | `int`                                                                                           | Required | The number of events to return<br>**Constraints**: `<= 200` |
| `Offset`  | `int`                                                                                           | Required | The number to events to skip (starting at `0`)              |
| `Total`   | `int`                                                                                           | Required | The number of events available                              |
| `Results` | [`[]models.NonFungibleTokenHistoryEvent`](../../doc/models/non-fungible-token-history-event.md) | Required | -                                                           |

## Example (as JSON)

```json
{
  "limit": 186,
  "offset": 26,
  "total": 24,
  "results": [
    {
      "sender": "sender2",
      "recipient": "recipient2",
      "event_index": 124,
      "asset_event_type": "asset_event_type8",
      "tx_id": "tx_id0"
    }
  ]
}
```
