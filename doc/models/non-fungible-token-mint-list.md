# Non Fungible Token Mint List

List of Non-Fungible Token mint events for an asset identifier

## Structure

`NonFungibleTokenMintList`

## Fields

| Name      | Type                                                                           | Tags     | Description                                                      |
| --------- | ------------------------------------------------------------------------------ | -------- | ---------------------------------------------------------------- |
| `Limit`   | `int`                                                                          | Required | The number of mint events to return<br>**Constraints**: `<= 200` |
| `Offset`  | `int`                                                                          | Required | The number to mint events to skip (starting at `0`)              |
| `Total`   | `int`                                                                          | Required | The number of mint events available                              |
| `Results` | [`[]models.NonFungibleTokenMint`](../../doc/models/non-fungible-token-mint.md) | Required | -                                                                |

## Example (as JSON)

```json
{
  "limit": 190,
  "offset": 162,
  "total": 96,
  "results": [
    {
      "recipient": "recipient2",
      "event_index": 124,
      "value": {
        "hex": "hex6",
        "repr": "repr2"
      },
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
