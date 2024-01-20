# Rosetta Mempool Transaction Request

A MempoolTransactionRequest is utilized to retrieve a transaction from the mempool.

## Structure

`RosettaMempoolTransactionRequest`

## Fields

| Name                    | Type                                                                         | Tags     | Description                                                                                                       |
| ----------------------- | ---------------------------------------------------------------------------- | -------- | ----------------------------------------------------------------------------------------------------------------- |
| `NetworkIdentifier`     | [`models.NetworkIdentifier`](../../doc/models/network-identifier.md)         | Required | The network_identifier specifies which network a particular object is associated with.                            |
| `TransactionIdentifier` | [`models.TransactionIdentifier`](../../doc/models/transaction-identifier.md) | Required | The transaction_identifier uniquely identifies a transaction in a particular network and block or in the mempool. |

## Example (as JSON)

```json
{
  "network_identifier": {
    "blockchain": "blockchain2",
    "network": "network6",
    "sub_network_identifier": {
      "network": "network6",
      "metadata": {
        "producer": "producer4"
      }
    }
  },
  "transaction_identifier": {
    "hash": "hash6"
  }
}
```
