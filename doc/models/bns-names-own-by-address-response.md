# Bns Names Own by Address Response

Retrieves a list of names owned by the address provided.

## Structure

`BnsNamesOwnByAddressResponse`

## Fields

| Name    | Type       | Tags     | Description                                          |
| ------- | ---------- | -------- | ---------------------------------------------------- |
| `Names` | `[]string` | Required | **Constraints**: _Pattern_: `^([a-z0-9-_.+]{3,37})$` |

## Example (as JSON)

```json
{
  "names": ["names2", "names1", "names0"]
}
```
