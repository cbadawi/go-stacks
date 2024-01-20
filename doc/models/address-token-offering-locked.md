# Address Token Offering Locked

Token Offering Locked

## Structure

`AddressTokenOfferingLocked`

## Fields

| Name             | Type                                                                            | Tags     | Description                                            |
| ---------------- | ------------------------------------------------------------------------------- | -------- | ------------------------------------------------------ |
| `TotalLocked`    | `string`                                                                        | Required | Micro-STX amount still locked at current block height. |
| `TotalUnlocked`  | `string`                                                                        | Required | Micro-STX amount unlocked at current block height.     |
| `UnlockSchedule` | [`[]models.AddressUnlockSchedule`](../../doc/models/address-unlock-schedule.md) | Required | -                                                      |

## Example (as JSON)

```json
{
  "total_locked": "total_locked4",
  "total_unlocked": "total_unlocked6",
  "unlock_schedule": [
    {
      "amount": "amount4",
      "block_height": 184.68
    }
  ]
}
```
