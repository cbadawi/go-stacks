# Rosetta Construction Submit Response

TransactionIdentifier contains the transaction_identifier of a transaction that was submitted to either /construction/submit.

## Structure

`RosettaConstructionSubmitResponse`

## Fields

| Name                    | Type                                                                         | Tags     | Description                                                                                                       |
| ----------------------- | ---------------------------------------------------------------------------- | -------- | ----------------------------------------------------------------------------------------------------------------- |
| `TransactionIdentifier` | [`models.TransactionIdentifier`](../../doc/models/transaction-identifier.md) | Required | The transaction_identifier uniquely identifies a transaction in a particular network and block or in the mempool. |
| `Metadata`              | `*interface{}`                                                               | Optional | -                                                                                                                 |

## Example (as JSON)

```json
{
  "transaction_identifier": {
    "hash": "hash6"
  },
  "metadata": {
    "key1": "val1",
    "key2": "val2"
  }
}
```
