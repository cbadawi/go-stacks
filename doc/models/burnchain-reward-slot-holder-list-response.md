# Burnchain Reward Slot Holder List Response

GET request that returns reward slot holders

## Structure

`BurnchainRewardSlotHolderListResponse`

## Fields

| Name      | Type                                                                                     | Tags     | Description                                                       |
| --------- | ---------------------------------------------------------------------------------------- | -------- | ----------------------------------------------------------------- |
| `Limit`   | `int`                                                                                    | Required | The number of items to return<br>**Constraints**: `<= 30`         |
| `Offset`  | `int`                                                                                    | Required | The number of items to skip (starting at `0`)<br>**Default**: `0` |
| `Total`   | `int`                                                                                    | Required | Total number of available items                                   |
| `Results` | [`[]models.BurnchainRewardSlotHolder`](../../doc/models/burnchain-reward-slot-holder.md) | Required | -                                                                 |

## Example (as JSON)

```json
{
  "limit": 152,
  "offset": 0,
  "total": 246,
  "results": [
    {
      "canonical": false,
      "burn_block_hash": "burn_block_hash0",
      "burn_block_height": 104,
      "address": "address2",
      "slot_index": 232
    }
  ]
}
```
