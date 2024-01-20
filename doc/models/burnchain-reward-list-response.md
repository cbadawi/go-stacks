# Burnchain Reward List Response

GET request that returns blocks

## Structure

`BurnchainRewardListResponse`

## Fields

| Name      | Type                                                               | Tags     | Description                                                                   |
| --------- | ------------------------------------------------------------------ | -------- | ----------------------------------------------------------------------------- |
| `Limit`   | `int`                                                              | Required | The number of burnchain rewards to return<br>**Constraints**: `<= 30`         |
| `Offset`  | `int`                                                              | Required | The number to burnchain rewards to skip (starting at `0`)<br>**Default**: `0` |
| `Results` | [`[]models.BurnchainReward`](../../doc/models/burnchain-reward.md) | Required | -                                                                             |

## Example (as JSON)

```json
{
  "limit": 66,
  "offset": 0,
  "results": [
    {
      "canonical": false,
      "burn_block_hash": "burn_block_hash0",
      "burn_block_height": 104,
      "burn_amount": "burn_amount4",
      "reward_recipient": "reward_recipient0",
      "reward_amount": "reward_amount0",
      "reward_index": 244
    }
  ]
}
```
