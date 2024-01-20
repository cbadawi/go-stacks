# Get Stx Supply Legacy Format Response

GET request that returns network target block times

## Structure

`GetStxSupplyLegacyFormatResponse`

## Fields

| Name                      | Type     | Tags     | Description                                                              |
| ------------------------- | -------- | -------- | ------------------------------------------------------------------------ |
| `UnlockedPercent`         | `string` | Required | String quoted decimal number of the percentage of STX that have unlocked |
| `TotalStacks`             | `string` | Required | String quoted decimal number of the total possible number of STX         |
| `TotalStacksFormatted`    | `string` | Required | Same as `totalStacks` but formatted with comma thousands separators      |
| `UnlockedSupply`          | `string` | Required | String quoted decimal number of the STX that have been mined or unlocked |
| `UnlockedSupplyFormatted` | `string` | Required | Same as `unlockedSupply` but formatted with comma thousands separators   |
| `BlockHeight`             | `string` | Required | The block height at which this information was queried                   |

## Example (as JSON)

```json
{
  "unlockedPercent": "unlockedPercent0",
  "totalStacks": "totalStacks8",
  "totalStacksFormatted": "totalStacksFormatted6",
  "unlockedSupply": "unlockedSupply4",
  "unlockedSupplyFormatted": "unlockedSupplyFormatted4",
  "blockHeight": "blockHeight4"
}
```
