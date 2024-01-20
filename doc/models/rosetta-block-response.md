# Rosetta Block Response

A BlockResponse includes a fully-populated block or a partially-populated block with a list of other transactions to fetch (other_transactions). As a result of the consensus algorithm of some blockchains, blocks can be omitted (i.e. certain block indexes can be skipped). If a query for one of these omitted indexes is made, the response should not include a Block object. It is VERY important to note that blocks MUST still form a canonical, connected chain of blocks where each block has a unique index. In other words, the PartialBlockIdentifier of a block after an omitted block should reference the last non-omitted block.

## Structure

`RosettaBlockResponse`

## Fields

| Name                | Type                                                                                      | Tags     | Description                                                                                                                                                                                                                                                                                                                                                                |
| ------------------- | ----------------------------------------------------------------------------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Block`             | [`*models.RosettaBlock`](../../doc/models/rosetta-block.md)                               | Optional | Blocks contain an array of Transactions that occurred at a particular BlockIdentifier. A hard requirement for blocks returned by Rosetta implementations is that they MUST be inalterable: once a client has requested and received a block identified by a specific BlockIndentifier, all future calls for that same BlockIdentifier must return the same block contents. |
| `OtherTransactions` | [`[]models.OtherTransactionIdentifier`](../../doc/models/other-transaction-identifier.md) | Optional | Some blockchains may require additional transactions to be fetched that weren't returned in the block response (ex: block only returns transaction hashes). For blockchains with a lot of transactions in each block, this can be very useful as consumers can concurrently fetch all transactions returned.                                                               |

## Example (as JSON)

```json
{
  "block": {
    "block_identifier": {
      "hash": "hash6",
      "index": 54
    },
    "parent_block_identifier": {
      "index": 192,
      "hash": "hash8"
    },
    "timestamp": 214,
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
      },
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
      },
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
  },
  "other_transactions": [
    {
      "hash": "hash6"
    }
  ]
}
```
