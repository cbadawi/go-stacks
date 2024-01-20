# Rosetta Transaction

Transactions contain an array of Operations that are attributable to the same TransactionIdentifier.

## Structure

`RosettaTransaction`

## Fields

| Name                    | Type                                                                         | Tags     | Description                                                                                                                                                          |
| ----------------------- | ---------------------------------------------------------------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `TransactionIdentifier` | [`models.TransactionIdentifier`](../../doc/models/transaction-identifier.md) | Required | The transaction_identifier uniquely identifies a transaction in a particular network and block or in the mempool.                                                    |
| `Operations`            | [`[]models.RosettaOperation`](../../doc/models/rosetta-operation.md)         | Required | List of operations                                                                                                                                                   |
| `Metadata`              | [`*models.Metadata3`](../../doc/models/metadata-3.md)                        | Optional | Transactions that are related to other transactions (like a cross-shard transaction) should include the tranaction_identifier of these transactions in the metadata. |

## Example (as JSON)

```json
{
  "transaction_identifier": {
    "hash": "hash6"
  },
  "operations": [
    {
      "operation_identifier": {
        "index": 254,
        "network_index": 122
      },
      "related_operations": [
        {
          "index": 246,
          "network_index": 130
        }
      ],
      "type": "type8",
      "status": "status6",
      "account": {
        "address": "address6",
        "sub_account": {
          "address": "address4",
          "metadata": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        "metadata": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "amount": {
        "value": "value0",
        "currency": {
          "symbol": "symbol8",
          "decimals": 194,
          "metadata": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        "metadata": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "coin_change": {
        "coin_identifier": {
          "identifier": "identifier8"
        },
        "coin_action": "coin_created"
      }
    }
  ],
  "metadata": {
    "memo": "memo0",
    "size": 176,
    "lockTime": 36
  }
}
```
