# Transaction Fee Estimate Request

POST request for estimated fee

## Structure

`TransactionFeeEstimateRequest`

## Fields

| Name                 | Type     | Tags     | Description |
| -------------------- | -------- | -------- | ----------- |
| `TransactionPayload` | `string` | Required | -           |
| `EstimatedLen`       | `*int`   | Optional | -           |

## Example (as JSON)

```json
{
  "transaction_payload": "transaction_payload4",
  "estimated_len": 202
}
```
