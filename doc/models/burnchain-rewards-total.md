# Burnchain Rewards Total

Total burnchain rewards made to a recipient

## Structure

`BurnchainRewardsTotal`

## Fields

| Name              | Type     | Tags     | Description                                                                                                                     |
| ----------------- | -------- | -------- | ------------------------------------------------------------------------------------------------------------------------------- |
| `RewardRecipient` | `string` | Required | The recipient address that received the burnchain rewards, in the format native to the burnchain (e.g. B58 encoded for Bitcoin) |
| `RewardAmount`    | `string` | Required | The total amount of burnchain tokens rewarded to the recipient, in the smallest unit (e.g. satoshis for Bitcoin)                |

## Example (as JSON)

```json
{
  "reward_recipient": "reward_recipient2",
  "reward_amount": "reward_amount2"
}
```
