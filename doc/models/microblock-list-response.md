# Microblock List Response

GET request that returns microblocks

## Structure

`MicroblockListResponse`

## Fields

| Name      | Type                                                    | Tags     | Description                                                             |
| --------- | ------------------------------------------------------- | -------- | ----------------------------------------------------------------------- |
| `Limit`   | `int`                                                   | Required | The number of microblocks to return<br>**Constraints**: `<= 30`         |
| `Offset`  | `int`                                                   | Required | The number to microblocks to skip (starting at `0`)<br>**Default**: `0` |
| `Total`   | `int`                                                   | Required | The number of microblocks available                                     |
| `Results` | [`[]models.Microblock`](../../doc/models/microblock.md) | Required | -                                                                       |

## Example (as JSON)

```json
{
  "limit": 38,
  "offset": 0,
  "total": 132,
  "results": [
    {
      "canonical": false,
      "microblock_canonical": false,
      "microblock_hash": "microblock_hash4",
      "microblock_sequence": 106,
      "microblock_parent_hash": "microblock_parent_hash0",
      "block_height": 52,
      "parent_block_height": 92,
      "parent_block_hash": "parent_block_hash0",
      "parent_burn_block_hash": "parent_burn_block_hash6",
      "parent_burn_block_time": 10,
      "parent_burn_block_time_iso": "parent_burn_block_time_iso2",
      "parent_burn_block_height": 222,
      "block_hash": "block_hash8",
      "txs": ["txs5", "txs6", "txs7"]
    }
  ]
}
```
