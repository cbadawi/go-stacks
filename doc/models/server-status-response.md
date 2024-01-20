# Server Status Response

GET blockchain API status

## Structure

`ServerStatusResponse`

## Fields

| Name                | Type                                                | Tags     | Description                                  |
| ------------------- | --------------------------------------------------- | -------- | -------------------------------------------- |
| `ServerVersion`     | `*string`                                           | Optional | the server version that is currently running |
| `Status`            | `string`                                            | Required | the current server status                    |
| `PoxV1UnlockHeight` | `Optional[int]`                                     | Optional | -                                            |
| `PoxV2UnlockHeight` | `Optional[int]`                                     | Optional | -                                            |
| `PoxV3UnlockHeight` | `Optional[int]`                                     | Optional | -                                            |
| `ChainTip`          | [`*models.ChainTip`](../../doc/models/chain-tip.md) | Optional | Current chain tip information                |

## Example (as JSON)

```json
{
  "server_version": "server_version2",
  "status": "status2",
  "pox_v1_unlock_height": 108,
  "pox_v2_unlock_height": 198,
  "pox_v3_unlock_height": 160,
  "chain_tip": {
    "block_height": 114,
    "block_hash": "block_hash0",
    "index_block_hash": "index_block_hash8",
    "microblock_hash": "microblock_hash6",
    "microblock_sequence": 196,
    "burn_block_height": 194
  }
}
```
