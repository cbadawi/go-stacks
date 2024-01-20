# Non Fungible Token Mint

Describes the minting of a Non-Fungible Token

## Structure

`NonFungibleTokenMint`

## Fields

| Name         | Type                                                     | Tags     | Description                                              |
| ------------ | -------------------------------------------------------- | -------- | -------------------------------------------------------- |
| `Recipient`  | `*string`                                                | Optional | -                                                        |
| `EventIndex` | `*int`                                                   | Optional | -                                                        |
| `Value`      | [`*models.Value18`](../../doc/models/value-18.md)        | Optional | Non-Fungible Token value                                 |
| `TxId`       | `*string`                                                | Optional | -                                                        |
| `Tx`         | [`*models.Transaction`](../../doc/models/transaction.md) | Optional | Describes all transaction types on Stacks 2.0 blockchain |

## Example (as JSON)

```json
{
  "recipient": "recipient6",
  "event_index": 156,
  "value": {
    "hex": "hex6",
    "repr": "repr2"
  },
  "tx_id": "tx_id8",
  "tx": {
    "tx_id": "tx_id4",
    "nonce": 156,
    "fee_rate": "fee_rate2",
    "sender_address": "sender_address2",
    "sponsor_nonce": 200
  }
}
```
