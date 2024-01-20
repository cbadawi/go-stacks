# Contract Log

## Structure

`ContractLog`

## Fields

| Name         | Type                                        | Tags     | Description |
| ------------ | ------------------------------------------- | -------- | ----------- |
| `ContractId` | `string`                                    | Required | -           |
| `Topic`      | `string`                                    | Required | -           |
| `Value`      | [`models.Value`](../../doc/models/value.md) | Required | -           |

## Example (as JSON)

```json
{
  "contract_id": "contract_id8",
  "topic": "topic2",
  "value": {
    "hex": "hex6",
    "repr": "repr2"
  }
}
```
