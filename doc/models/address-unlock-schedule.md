# Address Unlock Schedule

Unlock schedule amount and block height

## Structure

`AddressUnlockSchedule`

## Fields

| Name          | Type      | Tags     | Description                                   |
| ------------- | --------- | -------- | --------------------------------------------- |
| `Amount`      | `string`  | Required | Micro-STX amount locked at this block height. |
| `BlockHeight` | `float64` | Required | -                                             |

## Example (as JSON)

```json
{
  "amount": "amount2",
  "block_height": 6.06
}
```
