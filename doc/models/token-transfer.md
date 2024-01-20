# Token Transfer

## Structure

`TokenTransfer`

## Fields

| Name               | Type     | Tags     | Description                                                                                   |
| ------------------ | -------- | -------- | --------------------------------------------------------------------------------------------- |
| `RecipientAddress` | `string` | Required | -                                                                                             |
| `Amount`           | `string` | Required | Transfer amount as Integer string (64-bit unsigned integer)                                   |
| `Memo`             | `string` | Required | Hex encoded arbitrary message, up to 34 bytes length (should try decoding to an ASCII string) |

## Example (as JSON)

```json
{
  "recipient_address": "recipient_address8",
  "amount": "amount4",
  "memo": "memo6"
}
```
