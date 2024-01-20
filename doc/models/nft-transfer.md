# Nft Transfer

## Structure

`NftTransfer`

## Fields

| Name              | Type                                             | Tags     | Description                          |
| ----------------- | ------------------------------------------------ | -------- | ------------------------------------ |
| `AssetIdentifier` | `string`                                         | Required | Non Fungible Token asset identifier. |
| `Value`           | [`models.Value11`](../../doc/models/value-11.md) | Required | Non Fungible Token asset value.      |
| `Sender`          | `*string`                                        | Optional | Principal that sent the asset.       |
| `Recipient`       | `*string`                                        | Optional | Principal that received the asset.   |

## Example (as JSON)

```json
{
  "asset_identifier": "asset_identifier2",
  "value": {
    "hex": "hex6",
    "repr": "repr2"
  },
  "sender": "sender2",
  "recipient": "recipient2"
}
```
