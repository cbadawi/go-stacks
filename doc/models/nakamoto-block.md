# Nakamoto Block

A block

## Structure

`NakamotoBlock`

## Fields

| Name                       | Type      | Tags     | Description                                                                             |
| -------------------------- | --------- | -------- | --------------------------------------------------------------------------------------- |
| `Canonical`                | `bool`    | Required | Set to `true` if block corresponds to the canonical chain tip                           |
| `Height`                   | `int`     | Required | Height of the block                                                                     |
| `Hash`                     | `string`  | Required | Hash representing the block                                                             |
| `IndexBlockHash`           | `string`  | Required | The only hash that can uniquely identify an anchored block or an unconfirmed state trie |
| `ParentBlockHash`          | `string`  | Required | Hash of the parent block                                                                |
| `ParentIndexBlockHash`     | `string`  | Required | Index block hash of the parent block                                                    |
| `BurnBlockTime`            | `float64` | Required | Unix timestamp (in seconds) indicating when this block was mined.                       |
| `BurnBlockTimeIso`         | `string`  | Required | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) indicating when this block was mined.            |
| `BurnBlockHash`            | `string`  | Required | Hash of the anchor chain block                                                          |
| `BurnBlockHeight`          | `int`     | Required | Height of the anchor chain block                                                        |
| `MinerTxid`                | `string`  | Required | Anchor chain transaction ID                                                             |
| `TxCount`                  | `int`     | Required | Number of transactions included in the block                                            |
| `ExecutionCostReadCount`   | `int`     | Required | Execution cost read count.                                                              |
| `ExecutionCostReadLength`  | `int`     | Required | Execution cost read length.                                                             |
| `ExecutionCostRuntime`     | `int`     | Required | Execution cost runtime.                                                                 |
| `ExecutionCostWriteCount`  | `int`     | Required | Execution cost write count.                                                             |
| `ExecutionCostWriteLength` | `int`     | Required | Execution cost write length.                                                            |

## Example (as JSON)

```json
{
  "canonical": false,
  "height": 192,
  "hash": "hash6",
  "index_block_hash": "index_block_hash0",
  "parent_block_hash": "parent_block_hash4",
  "parent_index_block_hash": "parent_index_block_hash4",
  "burn_block_time": 41.8,
  "burn_block_time_iso": "burn_block_time_iso8",
  "burn_block_hash": "burn_block_hash4",
  "burn_block_height": 102,
  "miner_txid": "miner_txid2",
  "tx_count": 150,
  "execution_cost_read_count": 8,
  "execution_cost_read_length": 116,
  "execution_cost_runtime": 228,
  "execution_cost_write_count": 252,
  "execution_cost_write_length": 60
}
```
