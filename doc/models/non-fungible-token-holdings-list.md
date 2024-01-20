# Non Fungible Token Holdings List

List of Non-Fungible Token holdings

## Structure

`NonFungibleTokenHoldingsList`

## Fields

| Name      | Type                                                                                 | Tags     | Description                                                                      |
| --------- | ------------------------------------------------------------------------------------ | -------- | -------------------------------------------------------------------------------- |
| `Limit`   | `int`                                                                                | Required | The number of Non-Fungible Token holdings to return<br>**Constraints**: `<= 200` |
| `Offset`  | `int`                                                                                | Required | The number to Non-Fungible Token holdings to skip (starting at `0`)              |
| `Total`   | `int`                                                                                | Required | The number of Non-Fungible Token holdings available                              |
| `Results` | [`[]models.NonFungibleTokenHolding`](../../doc/models/non-fungible-token-holding.md) | Required | -                                                                                |

## Example (as JSON)

```json
{
  "limit": 118,
  "offset": 214,
  "total": 212,
  "results": [
    {
      "asset_identifier": "asset_identifier8",
      "value": {
        "hex": "hex6",
        "repr": "repr2"
      },
      "block_height": 13.32,
      "tx_id": "tx_id0",
      "tx": {
        "tx_id": "tx_id4",
        "nonce": 156,
        "fee_rate": "fee_rate2",
        "sender_address": "sender_address2",
        "sponsor_nonce": 200
      }
    }
  ]
}
```
