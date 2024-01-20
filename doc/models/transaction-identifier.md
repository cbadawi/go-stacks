# Transaction Identifier

The transaction_identifier uniquely identifies a transaction in a particular network and block or in the mempool.

## Structure

`TransactionIdentifier`

## Fields

| Name   | Type     | Tags     | Description                                                                                                                    |
| ------ | -------- | -------- | ------------------------------------------------------------------------------------------------------------------------------ |
| `Hash` | `string` | Required | Any transactions that are attributable only to a block (ex: a block event) should use the hash of the block as the identifier. |

## Example (as JSON)

```json
{
  "hash": "hash4"
}
```
