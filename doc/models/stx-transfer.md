# Stx Transfer

## Structure

`StxTransfer`

## Fields

| Name        | Type      | Tags     | Description                                                              |
| ----------- | --------- | -------- | ------------------------------------------------------------------------ |
| `Amount`    | `string`  | Required | Amount transferred in micro-STX as an integer string.                    |
| `Sender`    | `*string` | Optional | Principal that sent STX. This is unspecified if the STX were minted.     |
| `Recipient` | `*string` | Optional | Principal that received STX. This is unspecified if the STX were burned. |

## Example (as JSON)

```json
{
  "amount": "amount4",
  "sender": "sender4",
  "recipient": "recipient4"
}
```
