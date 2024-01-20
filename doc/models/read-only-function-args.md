# Read Only Function Args

Describes representation of a Type-0 Stacks 2.0 transaction. https://github.com/blockstack/stacks-blockchain/blob/master/sip/sip-005-blocks-and-transactions.md#type-0-transferring-an-asset

## Structure

`ReadOnlyFunctionArgs`

## Fields

| Name        | Type       | Tags     | Description                               |
| ----------- | ---------- | -------- | ----------------------------------------- |
| `Sender`    | `string`   | Required | The simulated tx-sender                   |
| `Arguments` | `[]string` | Required | An array of hex serialized Clarity values |

## Example (as JSON)

```json
{
  "sender": "sender2",
  "arguments": ["arguments3", "arguments4", "arguments5"]
}
```
