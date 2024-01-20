# Rosetta Mempool Response

A MempoolResponse contains all transaction identifiers in the mempool for a particular network_identifier.

## Structure

`RosettaMempoolResponse`

## Fields

| Name                     | Type                                                                           | Tags     | Description |
| ------------------------ | ------------------------------------------------------------------------------ | -------- | ----------- |
| `TransactionIdentifiers` | [`[]models.TransactionIdentifier`](../../doc/models/transaction-identifier.md) | Required | -           |
| `Metadata`               | `*interface{}`                                                                 | Optional | -           |

## Example (as JSON)

```json
{
  "transaction_identifiers": [
    {
      "hash": "hash6"
    }
  ],
  "metadata": {
    "key1": "val1",
    "key2": "val2"
  }
}
```
