# Transaction Event Asset 1

## Structure

`TransactionEventAsset1`

## Fields

| Name             | Type                                                                                            | Tags     | Description |
| ---------------- | ----------------------------------------------------------------------------------------------- | -------- | ----------- |
| `AssetEventType` | [`models.TransactionEventAssetTypeEnum`](../../doc/models/transaction-event-asset-type-enum.md) | Required | -           |
| `AssetId`        | `string`                                                                                        | Required | -           |
| `Sender`         | `string`                                                                                        | Required | -           |
| `Recipient`      | `string`                                                                                        | Required | -           |
| `Amount`         | `string`                                                                                        | Required | -           |
| `Value`          | [`models.Value`](../../doc/models/value.md)                                                     | Required | -           |
| `Memo`           | `*string`                                                                                       | Optional | -           |

## Example (as JSON)

```json
{
  "asset_event_type": "mint",
  "asset_id": "asset_id8",
  "sender": "sender4",
  "recipient": "recipient4",
  "amount": "amount4",
  "value": {
    "hex": "hex6",
    "repr": "repr2"
  },
  "memo": "memo6"
}
```
