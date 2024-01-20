# Burnchain Reward

Reward payment made on the burnchain

## Structure

`BurnchainReward`

## Fields

| Name              | Type     | Tags     | Description                                                                                                                     |
| ----------------- | -------- | -------- | ------------------------------------------------------------------------------------------------------------------------------- |
| `Canonical`       | `bool`   | Required | Set to `true` if block corresponds to the canonical burchchain tip                                                              |
| `BurnBlockHash`   | `string` | Required | The hash representing the burnchain block                                                                                       |
| `BurnBlockHeight` | `int`    | Required | Height of the burnchain block                                                                                                   |
| `BurnAmount`      | `string` | Required | The total amount of burnchain tokens burned for this burnchain block, in the smallest unit (e.g. satoshis for Bitcoin)          |
| `RewardRecipient` | `string` | Required | The recipient address that received the burnchain rewards, in the format native to the burnchain (e.g. B58 encoded for Bitcoin) |
| `RewardAmount`    | `string` | Required | The amount of burnchain tokens rewarded to the recipient, in the smallest unit (e.g. satoshis for Bitcoin)                      |
| `RewardIndex`     | `int`    | Required | The index position of the reward entry, useful for ordering when there's more than one recipient per burnchain block            |

## Example (as JSON)

```json
{
  "canonical": false,
  "burn_block_hash": "burn_block_hash0",
  "burn_block_height": 194,
  "burn_amount": "burn_amount4",
  "reward_recipient": "reward_recipient0",
  "reward_amount": "reward_amount0",
  "reward_index": 78
}
```
