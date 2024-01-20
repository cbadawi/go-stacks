# Block

A block

## Structure

`Block`

## Fields

| Name                       | Type                 | Tags     | Description                                                                                                                                                                                                                                                                                                                                    |
| -------------------------- | -------------------- | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Canonical`                | `bool`               | Required | Set to `true` if block corresponds to the canonical chain tip                                                                                                                                                                                                                                                                                  |
| `Height`                   | `int`                | Required | Height of the block                                                                                                                                                                                                                                                                                                                            |
| `Hash`                     | `string`             | Required | Hash representing the block                                                                                                                                                                                                                                                                                                                    |
| `IndexBlockHash`           | `string`             | Required | The only hash that can uniquely identify an anchored block or an unconfirmed state trie                                                                                                                                                                                                                                                        |
| `ParentBlockHash`          | `string`             | Required | Hash of the parent block                                                                                                                                                                                                                                                                                                                       |
| `BurnBlockTime`            | `float64`            | Required | Unix timestamp (in seconds) indicating when this block was mined.                                                                                                                                                                                                                                                                              |
| `BurnBlockTimeIso`         | `string`             | Required | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined.                                                                                                                                                                                                                                                                   |
| `BurnBlockHash`            | `string`             | Required | Hash of the anchor chain block                                                                                                                                                                                                                                                                                                                 |
| `BurnBlockHeight`          | `int`                | Required | Height of the anchor chain block                                                                                                                                                                                                                                                                                                               |
| `MinerTxid`                | `string`             | Required | Anchor chain transaction ID                                                                                                                                                                                                                                                                                                                    |
| `ParentMicroblockHash`     | `string`             | Required | The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1. |
| `ParentMicroblockSequence` | `int`                | Required | The hash of the last streamed block that precedes this block to which this block is to be appended. Not every anchored block will have a parent microblock stream. An anchored block that does not have a parent microblock stream has the parent microblock hash set to an empty string, and the parent microblock sequence number set to -1. |
| `Txs`                      | `[]string`           | Required | List of transactions included in the block                                                                                                                                                                                                                                                                                                     |
| `MicroblocksAccepted`      | `[]string`           | Required | List of microblocks that were accepted in this anchor block. Not every anchored block will have a accepted all (or any) of the previously streamed microblocks. Microblocks that were orphaned are not included in this list.                                                                                                                  |
| `MicroblocksStreamed`      | `[]string`           | Required | List of microblocks that were streamed/produced by this anchor block's miner. This list only includes microblocks that were accepted in the following anchor block. Microblocks that were orphaned are not included in this list.                                                                                                              |
| `ExecutionCostReadCount`   | `int`                | Required | Execution cost read count.                                                                                                                                                                                                                                                                                                                     |
| `ExecutionCostReadLength`  | `int`                | Required | Execution cost read length.                                                                                                                                                                                                                                                                                                                    |
| `ExecutionCostRuntime`     | `int`                | Required | Execution cost runtime.                                                                                                                                                                                                                                                                                                                        |
| `ExecutionCostWriteCount`  | `int`                | Required | Execution cost write count.                                                                                                                                                                                                                                                                                                                    |
| `ExecutionCostWriteLength` | `int`                | Required | Execution cost write length.                                                                                                                                                                                                                                                                                                                   |
| `MicroblockTxCount`        | `map[string]float64` | Required | List of txs counts in each accepted microblock                                                                                                                                                                                                                                                                                                 |

## Example (as JSON)

```json
{
  "canonical": false,
  "height": 166,
  "hash": "hash2",
  "index_block_hash": "index_block_hash6",
  "parent_block_hash": "parent_block_hash0",
  "burn_block_time": 81.34,
  "burn_block_time_iso": "burn_block_time_iso4",
  "burn_block_hash": "burn_block_hash0",
  "burn_block_height": 76,
  "miner_txid": "miner_txid8",
  "parent_microblock_hash": "parent_microblock_hash8",
  "parent_microblock_sequence": 196,
  "txs": ["txs5"],
  "microblocks_accepted": ["microblocks_accepted8", "microblocks_accepted9"],
  "microblocks_streamed": ["microblocks_streamed8"],
  "execution_cost_read_count": 238,
  "execution_cost_read_length": 90,
  "execution_cost_runtime": 54,
  "execution_cost_write_count": 30,
  "execution_cost_write_length": 222,
  "microblock_tx_count": {
    "key0": 5.56,
    "key1": 5.55
  }
}
```
