# Core Node Pox Response

Get Proof of Transfer (PoX) information

## Structure

`CoreNodePoxResponse`

## Fields

| Name                         | Type     | Tags     | Description |
| ---------------------------- | -------- | -------- | ----------- |
| `ContractId`                 | `string` | Required | -           |
| `FirstBurnchainBlockHeight`  | `int`    | Required | -           |
| `MinAmountUstx`              | `int`    | Required | -           |
| `RegistrationWindowLength`   | `int`    | Required | -           |
| `RejectionFraction`          | `int`    | Required | -           |
| `RewardCycleId`              | `int`    | Required | -           |
| `RewardCycleLength`          | `int`    | Required | -           |
| `RejectionVotesLeftRequired` | `int`    | Required | -           |
| `TotalLiquidSupplyUstx`      | `int`    | Required | -           |

## Example (as JSON)

```json
{
  "contract_id": "contract_id6",
  "first_burnchain_block_height": 142,
  "min_amount_ustx": 222,
  "registration_window_length": 44,
  "rejection_fraction": 160,
  "reward_cycle_id": 206,
  "reward_cycle_length": 164,
  "rejection_votes_left_required": 148,
  "total_liquid_supply_ustx": 104
}
```
