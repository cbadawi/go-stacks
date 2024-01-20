# Burn Block

A burn block

## Structure

`BurnBlock`

## Fields

| Name               | Type       | Tags     | Description                                                                  |
| ------------------ | ---------- | -------- | ---------------------------------------------------------------------------- |
| `BurnBlockTime`    | `float64`  | Required | Unix timestamp (in seconds) indicating when this block was mined.            |
| `BurnBlockTimeIso` | `string`   | Required | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined. |
| `BurnBlockHash`    | `string`   | Required | Hash of the anchor chain block                                               |
| `BurnBlockHeight`  | `int`      | Required | Height of the anchor chain block                                             |
| `StacksBlocks`     | `[]string` | Required | Hashes of the Stacks blocks included in the burn block                       |

## Example (as JSON)

```json
{
  "burn_block_time": 95.04,
  "burn_block_time_iso": "burn_block_time_iso4",
  "burn_block_hash": "burn_block_hash0",
  "burn_block_height": 242,
  "stacks_blocks": ["stacks_blocks2", "stacks_blocks3"]
}
```
