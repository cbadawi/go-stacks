# Details

Often times it is useful to return context specific to the request that caused the error (i.e. a sample of the stack trace or impacted account) in addition to the standard error message.

## Structure

`Details`

## Fields

| Name      | Type      | Tags     | Description |
| --------- | --------- | -------- | ----------- |
| `Address` | `*string` | Optional | -           |
| `Error`   | `*string` | Optional | -           |

## Example (as JSON)

```json
{
  "address": "address4",
  "error": "error2"
}
```
