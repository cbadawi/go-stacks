# Contract Source Response

GET request to get contract source

## Structure

`ContractSourceResponse`

## Fields

| Name            | Type     | Tags     | Description |
| --------------- | -------- | -------- | ----------- |
| `Source`        | `string` | Required | -           |
| `PublishHeight` | `int`    | Required | -           |
| `Proof`         | `string` | Required | -           |

## Example (as JSON)

```json
{
  "source": "source6",
  "publish_height": 112,
  "proof": "proof6"
}
```
