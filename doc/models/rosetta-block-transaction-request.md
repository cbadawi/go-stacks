# Rosetta Block Transaction Request

A BlockTransactionRequest is used to fetch a Transaction included in a block that is not returned in a BlockResponse.

## Structure

`RosettaBlockTransactionRequest`

## Fields

| Name                    | Type                                                                                           | Tags     | Description                                                                                                                                                                                            |
| ----------------------- | ---------------------------------------------------------------------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `NetworkIdentifier`     | [`models.NetworkIdentifier`](../../doc/models/network-identifier.md)                           | Required | The network_identifier specifies which network a particular object is associated with.                                                                                                                 |
| `BlockIdentifier`       | [`models.RosettaPartialBlockIdentifier`](../../doc/models/rosetta-partial-block-identifier.md) | Required | When fetching data by BlockIdentifier, it may be possible to only specify the index or hash. If neither property is specified, it is assumed that the client is making a request at the current block. |
| `TransactionIdentifier` | [`models.TransactionIdentifier`](../../doc/models/transaction-identifier.md)                   | Required | The transaction_identifier uniquely identifies a transaction in a particular network and block or in the mempool.                                                                                      |

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
  "block_identifier": {
    "hash": "hash6",
    "index": 54
  },
  "transaction_identifier": {
    "hash": "hash6"
  }
}
```
