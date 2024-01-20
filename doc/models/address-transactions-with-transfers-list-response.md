# Address Transactions With Transfers List Response

GET request that returns account transactions

## Structure

`AddressTransactionsWithTransfersListResponse`

## Fields

| Name      | Type                                                                                                 | Tags     | Description              |
| --------- | ---------------------------------------------------------------------------------------------------- | -------- | ------------------------ |
| `Limit`   | `int`                                                                                                | Required | **Constraints**: `<= 30` |
| `Offset`  | `int`                                                                                                | Required | -                        |
| `Total`   | `int`                                                                                                | Required | -                        |
| `Results` | [`[]models.AddressTransactionWithTransfers`](../../doc/models/address-transaction-with-transfers.md) | Required | -                        |

## Example (as JSON)

```json
{
  "limit": 200,
  "offset": 152,
  "total": 106,
  "results": [
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
  ]
}
```
