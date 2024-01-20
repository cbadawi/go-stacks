# Block Data

Returns basic search result information about the requested id

## Structure

`BlockData`

## Fields

| Name              | Type     | Tags     | Description                                  |
| ----------------- | -------- | -------- | -------------------------------------------- |
| `Canonical`       | `bool`   | Required | If the block lies within the canonical chain |
| `Hash`            | `string` | Required | Refers to the hash of the block              |
| `ParentBlockHash` | `string` | Required | -                                            |
| `BurnBlockTime`   | `int`    | Required | -                                            |
| `Height`          | `int`    | Required | -                                            |

## Example (as JSON)

```json
{
  "canonical": false,
  "hash": "hash6",
  "parent_block_hash": "parent_block_hash4",
  "burn_block_time": 194,
  "height": 46
}
```
