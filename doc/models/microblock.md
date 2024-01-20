# Microblock

A microblock

## Structure

`Microblock`

## Fields

| Name                     | Type       | Tags     | Description                                                                                                                                              |
| ------------------------ | ---------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Canonical`              | `bool`     | Required | Set to `true` if the microblock corresponds to the canonical chain tip.                                                                                  |
| `MicroblockCanonical`    | `bool`     | Required | Set to `true` if the microblock was not orphaned in a following anchor block. Defaults to `true` if the following anchor block has not yet been created. |
| `MicroblockHash`         | `string`   | Required | The SHA512/256 hash of this microblock.                                                                                                                  |
| `MicroblockSequence`     | `int`      | Required | A hint to describe how to order a set of microblocks. Starts at 0.                                                                                       |
| `MicroblockParentHash`   | `string`   | Required | The SHA512/256 hash of the previous signed microblock in this stream.                                                                                    |
| `BlockHeight`            | `int`      | Required | The anchor block height that confirmed this microblock.                                                                                                  |
| `ParentBlockHeight`      | `int`      | Required | The height of the anchor block that preceded this microblock.                                                                                            |
| `ParentBlockHash`        | `string`   | Required | The hash of the anchor block that preceded this microblock.                                                                                              |
| `ParentBurnBlockHash`    | `string`   | Required | The hash of the Bitcoin block that preceded this microblock.                                                                                             |
| `ParentBurnBlockTime`    | `int`      | Required | The block timestamp of the Bitcoin block that preceded this microblock.                                                                                  |
| `ParentBurnBlockTimeIso` | `string`   | Required | The ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) formatted block time of the bitcoin block that preceded this microblock.                                         |
| `ParentBurnBlockHeight`  | `int`      | Required | The height of the Bitcoin block that preceded this microblock.                                                                                           |
| `BlockHash`              | `*string`  | Required | The hash of the anchor block that confirmed this microblock. This wil be empty for unanchored microblocks                                                |
| `Txs`                    | `[]string` | Required | List of transactions included in the microblock                                                                                                          |

## Example (as JSON)

```json
{
  "canonical": false,
  "microblock_canonical": false,
  "microblock_hash": "microblock_hash6",
  "microblock_sequence": 4,
  "microblock_parent_hash": "microblock_parent_hash2",
  "block_height": 50,
  "parent_block_height": 246,
  "parent_block_hash": "parent_block_hash2",
  "parent_burn_block_hash": "parent_burn_block_hash4",
  "parent_burn_block_time": 164,
  "parent_burn_block_time_iso": "parent_burn_block_time_iso0",
  "parent_burn_block_height": 120,
  "block_hash": "block_hash0",
  "txs": ["txs3"]
}
```
