# Rosetta Construction Preprocess Request

ConstructionPreprocessRequest is passed to the /construction/preprocess endpoint so that a Rosetta implementation can determine which metadata it needs to request for construction

## Structure

`RosettaConstructionPreprocessRequest`

## Fields

| Name                     | Type                                                                         | Tags     | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| ------------------------ | ---------------------------------------------------------------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `NetworkIdentifier`      | [`models.NetworkIdentifier`](../../doc/models/network-identifier.md)         | Required | The network_identifier specifies which network a particular object is associated with.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| `Operations`             | [`[]models.RosettaOperation`](../../doc/models/rosetta-operation.md)         | Required | -                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| `Metadata`               | `*interface{}`                                                               | Optional | -                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| `MaxFee`                 | [`[]models.RosettaMaxFeeAmount`](../../doc/models/rosetta-max-fee-amount.md) | Optional | -                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| `SuggestedFeeMultiplier` | `*int`                                                                       | Optional | The caller can also provide a suggested fee multiplier to indicate that the suggested fee should be scaled. This may be used to set higher fees for urgent transactions or to pay lower fees when there is less urgency. It is assumed that providing a very low multiplier (like 0.0001) will never lead to a transaction being created with a fee less than the minimum network fee (if applicable). In the case that the caller provides both a max fee and a suggested fee multiplier, the max fee will set an upper bound on the suggested fee (regardless of the multiplier provided). |

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
    "key1": "val1",
    "key2": "val2"
  },
  "max_fee": [
    {
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
    {
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
    }
  ],
  "suggested_fee_multiplier": 56
}
```
