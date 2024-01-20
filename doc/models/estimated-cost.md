# Estimated Cost

## Structure

`EstimatedCost`

## Fields

| Name          | Type  | Tags     | Description |
| ------------- | ----- | -------- | ----------- |
| `ReadCount`   | `int` | Required | -           |
| `ReadLength`  | `int` | Required | -           |
| `Runtime`     | `int` | Required | -           |
| `WriteCount`  | `int` | Required | -           |
| `WriteLength` | `int` | Required | -           |

## Example (as JSON)

```json
{
  "read_count": 20,
  "read_length": 28,
  "runtime": 96,
  "write_count": 44,
  "write_length": 140
}
```
