# Network Block Times Response

GET request that returns network target block times

## Structure

`NetworkBlockTimesResponse`

## Fields

| Name      | Type                                                              | Tags     | Description |
| --------- | ----------------------------------------------------------------- | -------- | ----------- |
| `Mainnet` | [`models.TargetBlockTime`](../../doc/models/target-block-time.md) | Required | -           |
| `Testnet` | [`models.TargetBlockTime`](../../doc/models/target-block-time.md) | Required | -           |

## Example (as JSON)

```json
{
  "mainnet": {
    "target_block_time": 116
  },
  "testnet": {
    "target_block_time": 16
  }
}
```
