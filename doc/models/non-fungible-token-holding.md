# Non Fungible Token Holding

Describes the ownership of a Non-Fungible Token

## Structure

`NonFungibleTokenHolding`

## Fields

| Name              | Type                                                     | Tags     | Description                                              |
| ----------------- | -------------------------------------------------------- | -------- | -------------------------------------------------------- |
| `AssetIdentifier` | `*string`                                                | Optional | -                                                        |
| `Value`           | [`*models.Value18`](../../doc/models/value-18.md)        | Optional | Non-Fungible Token value                                 |
| `BlockHeight`     | `*float64`                                               | Optional | -                                                        |
| `TxId`            | `*string`                                                | Optional | -                                                        |
| `Tx`              | [`*models.Transaction`](../../doc/models/transaction.md) | Optional | Describes all transaction types on Stacks 2.0 blockchain |

## Example (as JSON)

```json
{
  "asset_identifier": "asset_identifier2",
  "value": {
    "hex": "hex6",
    "repr": "repr2"
  },
  "block_height": 169.96,
  "tx_id": "tx_id4",
  "tx": {
    "tx_id": "tx_id4",
    "nonce": 156,
    "fee_rate": "fee_rate2",
    "sender_address": "sender_address2",
    "sponsor_nonce": 200
  }
}
```
