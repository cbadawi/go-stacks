# Chain Tip

Current chain tip information

## Structure

`ChainTip`

## Fields

| Name                 | Type      | Tags     | Description                            |
| -------------------- | --------- | -------- | -------------------------------------- |
| `BlockHeight`        | `int`     | Required | the current block height               |
| `BlockHash`          | `string`  | Required | the current block hash                 |
| `IndexBlockHash`     | `string`  | Required | the current index block hash           |
| `MicroblockHash`     | `*string` | Optional | the current microblock hash            |
| `MicroblockSequence` | `*int`    | Optional | the current microblock sequence number |
| `BurnBlockHeight`    | `int`     | Required | the current burn chain block height    |

## Example (as JSON)

```json
{
  "block_height": 228,
  "block_hash": "block_hash4",
  "index_block_hash": "index_block_hash2",
  "microblock_hash": "microblock_hash0",
  "microblock_sequence": 26,
  "burn_block_height": 24
}
```
