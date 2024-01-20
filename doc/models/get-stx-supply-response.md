# Get Stx Supply Response

GET request that returns network target block times

## Structure

`GetStxSupplyResponse`

## Fields

| Name              | Type     | Tags     | Description                                                              |
| ----------------- | -------- | -------- | ------------------------------------------------------------------------ |
| `UnlockedPercent` | `string` | Required | String quoted decimal number of the percentage of STX that have unlocked |
| `TotalStx`        | `string` | Required | String quoted decimal number of the total possible number of STX         |
| `UnlockedStx`     | `string` | Required | String quoted decimal number of the STX that have been mined or unlocked |
| `BlockHeight`     | `int`    | Required | The block height at which this information was queried                   |

## Example (as JSON)

```json
{
  "unlocked_percent": "unlocked_percent2",
  "total_stx": "total_stx4",
  "unlocked_stx": "unlocked_stx2",
  "block_height": 18
}
```
