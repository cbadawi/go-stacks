# Burnchain Reward Slot Holder

Reward slot holder on the burnchain

## Structure

`BurnchainRewardSlotHolder`

## Fields

| Name              | Type     | Tags     | Description                                                                                                                       |
| ----------------- | -------- | -------- | --------------------------------------------------------------------------------------------------------------------------------- |
| `Canonical`       | `bool`   | Required | Set to `true` if block corresponds to the canonical burchchain tip                                                                |
| `BurnBlockHash`   | `string` | Required | The hash representing the burnchain block                                                                                         |
| `BurnBlockHeight` | `int`    | Required | Height of the burnchain block                                                                                                     |
| `Address`         | `string` | Required | The recipient address that validly received PoX commitments, in the format native to the burnchain (e.g. B58 encoded for Bitcoin) |
| `SlotIndex`       | `int`    | Required | The index position of the reward entry, useful for ordering when there's more than one slot per burnchain block                   |

## Example (as JSON)

```json
{
  "canonical": false,
  "burn_block_hash": "burn_block_hash4",
  "burn_block_height": 254,
  "address": "address6",
  "slot_index": 126
}
```
