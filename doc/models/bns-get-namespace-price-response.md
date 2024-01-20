# Bns Get Namespace Price Response

Fetch price for namespace.

## Structure

`BnsGetNamespacePriceResponse`

## Fields

| Name     | Type     | Tags     | Description                            |
| -------- | -------- | -------- | -------------------------------------- |
| `Units`  | `string` | Required | -                                      |
| `Amount` | `string` | Required | **Constraints**: _Pattern_: `^[0-9]+$` |

## Example (as JSON)

```json
{
  "units": "units2",
  "amount": "amount8"
}
```
