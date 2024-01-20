# Map Entry Response

Response of get data map entry request

## Structure

`MapEntryResponse`

## Fields

| Name    | Type      | Tags     | Description                                                          |
| ------- | --------- | -------- | -------------------------------------------------------------------- |
| `Data`  | `string`  | Required | Hex-encoded string of clarity value. It is always an optional tuple. |
| `Proof` | `*string` | Optional | Hex-encoded string of the MARF proof for the data                    |

## Example (as JSON)

```json
{
  "data": "data6",
  "proof": "proof0"
}
```
