# Metadata 2

Account-based blockchains that utilize a nonce or sequence number should include that number in the metadata. This number could be unique to the identifier or global across the account address.

## Structure

`Metadata2`

## Fields

| Name             | Type  | Tags     | Description |
| ---------------- | ----- | -------- | ----------- |
| `SequenceNumber` | `int` | Required | -           |

## Example (as JSON)

```json
{
  "sequence_number": 78
}
```
