# Nakamoto Block List Response

GET request that returns blocks

## Structure

`NakamotoBlockListResponse`

## Fields

| Name      | Type                                                           | Tags     | Description                                                        |
| --------- | -------------------------------------------------------------- | -------- | ------------------------------------------------------------------ |
| `Limit`   | `int`                                                          | Required | The number of blocks to return<br>**Constraints**: `<= 30`         |
| `Offset`  | `int`                                                          | Required | The number to blocks to skip (starting at `0`)<br>**Default**: `0` |
| `Total`   | `int`                                                          | Required | The number of blocks available                                     |
| `Results` | [`[]models.NakamotoBlock`](../../doc/models/nakamoto-block.md) | Required | -                                                                  |

## Example (as JSON)

```json
{
  "limit": 202,
  "offset": 0,
  "total": 216,
  "results": [
    {
      "canonical": false,
      "height": 194,
      "hash": "hash2",
      "index_block_hash": "index_block_hash6",
      "parent_block_hash": "parent_block_hash0",
      "parent_index_block_hash": "parent_index_block_hash0",
      "burn_block_time": 88.74,
      "burn_block_time_iso": "burn_block_time_iso4",
      "burn_block_hash": "burn_block_hash0",
      "burn_block_height": 104,
      "miner_txid": "miner_txid8",
      "tx_count": 104,
      "execution_cost_read_count": 10,
      "execution_cost_read_length": 118,
      "execution_cost_runtime": 26,
      "execution_cost_write_count": 2,
      "execution_cost_write_length": 194
    }
  ]
}
```
