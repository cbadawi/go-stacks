# Stx Lock Event

## Structure

`StxLockEvent`

## Fields

| Name            | Type     | Tags     | Description |
| --------------- | -------- | -------- | ----------- |
| `LockedAmount`  | `string` | Required | -           |
| `UnlockHeight`  | `int`    | Required | -           |
| `LockedAddress` | `string` | Required | -           |

## Example (as JSON)

```json
{
  "locked_amount": "locked_amount0",
  "unlock_height": 230,
  "locked_address": "locked_address2"
}
```
