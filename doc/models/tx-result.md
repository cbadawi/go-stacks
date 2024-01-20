# Tx Result

Result of the transaction. For contract calls, this will show the value returned by the call. For other transaction types, this will return a boolean indicating the success of the transaction.

## Structure

`TxResult`

## Fields

| Name   | Type     | Tags     | Description                                                 |
| ------ | -------- | -------- | ----------------------------------------------------------- |
| `Hex`  | `string` | Required | Hex string representing the value fo the transaction result |
| `Repr` | `string` | Required | Readable string of the transaction result                   |

## Example (as JSON)

```json
{
  "hex": "hex4",
  "repr": "repr8"
}
```
