# Rosetta Block

Blocks contain an array of Transactions that occurred at a particular BlockIdentifier. A hard requirement for blocks returned by Rosetta implementations is that they MUST be inalterable: once a client has requested and received a block identified by a specific BlockIndentifier, all future calls for that same BlockIdentifier must return the same block contents.

## Structure

`RosettaBlock`

## Fields

| Name                    | Type                                                                                         | Tags     | Description                                                                                                                                                                     |
| ----------------------- | -------------------------------------------------------------------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `BlockIdentifier`       | [`models.RosettaBlockIdentifier`](../../doc/models/rosetta-block-identifier.md)              | Required | The block_identifier uniquely identifies a block in a particular network.                                                                                                       |
| `ParentBlockIdentifier` | [`models.RosettaParentBlockIdentifier`](../../doc/models/rosetta-parent-block-identifier.md) | Required | The block_identifier uniquely identifies a block in a particular network.                                                                                                       |
| `Timestamp`             | `int`                                                                                        | Required | The timestamp of the block in milliseconds since the Unix Epoch. The timestamp is stored in milliseconds because some blockchains produce blocks more often than once a second. |
| `Transactions`          | [`[]models.RosettaTransaction`](../../doc/models/rosetta-transaction.md)                     | Required | All the transactions in the block                                                                                                                                               |
| `Metadata`              | [`*models.Metadata4`](../../doc/models/metadata-4.md)                                        | Optional | meta data                                                                                                                                                                       |

## Example (as JSON)

```json
{
  "block_identifier": {
    "hash": "hash6",
    "index": 54
  },
  "parent_block_identifier": {
    "index": 192,
    "hash": "hash8"
  },
  "timestamp": 2,
  "transactions": [
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
  ],
  "metadata": {
    "transactions_root": "transactions_root0",
    "difficulty": "difficulty6"
  }
}
```
