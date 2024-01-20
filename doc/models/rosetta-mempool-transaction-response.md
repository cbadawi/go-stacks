# Rosetta Mempool Transaction Response

A MempoolTransactionResponse contains an estimate of a mempool transaction. It may not be possible to know the full impact of a transaction in the mempool (ex: fee paid).

## Structure

`RosettaMempoolTransactionResponse`

## Fields

| Name          | Type                                                                   | Tags     | Description                                                                                          |
| ------------- | ---------------------------------------------------------------------- | -------- | ---------------------------------------------------------------------------------------------------- |
| `Transaction` | [`models.RosettaTransaction`](../../doc/models/rosetta-transaction.md) | Required | Transactions contain an array of Operations that are attributable to the same TransactionIdentifier. |
| `Metadata`    | `*interface{}`                                                         | Optional | -                                                                                                    |

## Example (as JSON)

```json
{
  "transaction": {
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
  },
  "metadata": {
    "key1": "val1",
    "key2": "val2"
  }
}
```
