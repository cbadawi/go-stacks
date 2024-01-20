# Burn Block List Response

GET request that returns burn blocks

## Structure

`BurnBlockListResponse`

## Fields

| Name      | Type                                                   | Tags     | Description                                                             |
| --------- | ------------------------------------------------------ | -------- | ----------------------------------------------------------------------- |
| `Limit`   | `int`                                                  | Required | The number of burn blocks to return<br>**Constraints**: `<= 30`         |
| `Offset`  | `int`                                                  | Required | The number to burn blocks to skip (starting at `0`)<br>**Default**: `0` |
| `Total`   | `int`                                                  | Required | The number of burn blocks available (regardless of filter parameters)   |
| `Results` | [`[]models.BurnBlock`](../../doc/models/burn-block.md) | Required | -                                                                       |

## Example (as JSON)

```json
{
  "limit": 206,
  "offset": 0,
  "total": 44,
  "results": [
    {
      "burn_block_time": 88.74,
      "burn_block_time_iso": "burn_block_time_iso4",
      "burn_block_hash": "burn_block_hash0",
      "burn_block_height": 104,
      "stacks_blocks": ["stacks_blocks2", "stacks_blocks3"]
    }
  ]
}
```
