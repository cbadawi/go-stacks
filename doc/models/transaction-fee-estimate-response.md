# Transaction Fee Estimate Response

POST response for estimated fee

## Structure

`TransactionFeeEstimateResponse`

## Fields

| Name                     | Type                                                         | Tags     | Description |
| ------------------------ | ------------------------------------------------------------ | -------- | ----------- |
| `EstimatedCostScalar`    | `int`                                                        | Required | -           |
| `CostScalarChangeByByte` | `*float64`                                                   | Optional | -           |
| `EstimatedCost`          | [`models.EstimatedCost`](../../doc/models/estimated-cost.md) | Required | -           |
| `Estimations`            | [`[]models.Estimation`](../../doc/models/estimation.md)      | Optional | -           |

## Example (as JSON)

```json
{
  "estimated_cost_scalar": 194,
  "cost_scalar_change_by_byte": 18.76,
  "estimated_cost": {
    "read_count": 204,
    "read_length": 100,
    "runtime": 168,
    "write_count": 228,
    "write_length": 68
  },
  "estimations": [
    {
      "fee_rate": 95.92,
      "fee": 76.72
    }
  ]
}
```
