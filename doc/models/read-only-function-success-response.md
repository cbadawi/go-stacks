# Read Only Function Success Response

GET request to get contract source

## Structure

`ReadOnlyFunctionSuccessResponse`

## Fields

| Name     | Type      | Tags     | Description |
| -------- | --------- | -------- | ----------- |
| `Okay`   | `bool`    | Required | -           |
| `Result` | `*string` | Optional | -           |
| `Cause`  | `*string` | Optional | -           |

## Example (as JSON)

```json
{
  "okay": false,
  "result": "result2",
  "cause": "cause0"
}
```
