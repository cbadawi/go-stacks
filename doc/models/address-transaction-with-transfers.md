# Address Transaction With Transfers

Transaction with STX transfers for a given address

## Structure

`AddressTransactionWithTransfers`

## Fields

| Name           | Type                                                       | Tags     | Description                                                                                 |
| -------------- | ---------------------------------------------------------- | -------- | ------------------------------------------------------------------------------------------- |
| `Tx`           | [`models.Transaction`](../../doc/models/transaction.md)    | Required | Describes all transaction types on Stacks 2.0 blockchain                                    |
| `StxSent`      | `string`                                                   | Required | Total sent from the given address, including the tx fee, in micro-STX as an integer string. |
| `StxReceived`  | `string`                                                   | Required | Total received by the given address in micro-STX as an integer string.                      |
| `StxTransfers` | [`[]models.StxTransfer`](../../doc/models/stx-transfer.md) | Required | -                                                                                           |
| `FtTransfers`  | [`[]models.FtTransfer`](../../doc/models/ft-transfer.md)   | Optional | -                                                                                           |
| `NftTransfers` | [`[]models.NftTransfer`](../../doc/models/nft-transfer.md) | Optional | -                                                                                           |

## Example (as JSON)

```json
{
  "tx": {
    "tx_id": "tx_id4",
    "nonce": 156,
    "fee_rate": "fee_rate2",
    "sender_address": "sender_address2",
    "sponsor_nonce": 200
  },
  "stx_sent": "stx_sent6",
  "stx_received": "stx_received4",
  "stx_transfers": [
    {
      "amount": "amount4",
      "sender": "sender4",
      "recipient": "recipient4"
    }
  ],
  "ft_transfers": [
    {
      "asset_identifier": "asset_identifier8",
      "amount": "amount8",
      "sender": "sender8",
      "recipient": "recipient8"
    },
    {
      "asset_identifier": "asset_identifier8",
      "amount": "amount8",
      "sender": "sender8",
      "recipient": "recipient8"
    }
  ],
  "nft_transfers": [
    {
      "asset_identifier": "asset_identifier8",
      "value": {
        "hex": "hex6",
        "repr": "repr2"
      },
      "sender": "sender2",
      "recipient": "recipient2"
    },
    {
      "asset_identifier": "asset_identifier8",
      "value": {
        "hex": "hex6",
        "repr": "repr2"
      },
      "sender": "sender2",
      "recipient": "recipient2"
    }
  ]
}
```
