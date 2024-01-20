# Ft Transfer

## Structure

`FtTransfer`

## Fields

| Name              | Type      | Tags     | Description                                                                                         |
| ----------------- | --------- | -------- | --------------------------------------------------------------------------------------------------- |
| `AssetIdentifier` | `string`  | Required | Fungible Token asset identifier.                                                                    |
| `Amount`          | `string`  | Required | Amount transferred as an integer string. This balance does not factor in possible SIP-010 decimals. |
| `Sender`          | `*string` | Optional | Principal that sent the asset.                                                                      |
| `Recipient`       | `*string` | Optional | Principal that received the asset.                                                                  |

## Example (as JSON)

```json
{
  "asset_identifier": "asset_identifier0",
  "amount": "amount0",
  "sender": "sender0",
  "recipient": "recipient0"
}
```
