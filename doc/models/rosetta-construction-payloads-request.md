# Rosetta Construction Payloads Request

ConstructionPayloadsRequest is the request to /construction/payloads. It contains the network, a slice of operations, and arbitrary metadata that was returned by the call to /construction/metadata. Optionally, the request can also include an array of PublicKeys associated with the AccountIdentifiers returned in ConstructionPreprocessResponse.

## Structure

`RosettaConstructionPayloadsRequest`

## Fields

| Name                | Type                                                                  | Tags     | Description                                                                            |
| ------------------- | --------------------------------------------------------------------- | -------- | -------------------------------------------------------------------------------------- |
| `NetworkIdentifier` | [`models.NetworkIdentifier`](../../doc/models/network-identifier.md)  | Required | The network_identifier specifies which network a particular object is associated with. |
| `Operations`        | [`[]models.RosettaOperation`](../../doc/models/rosetta-operation.md)  | Required | -                                                                                      |
| `PublicKeys`        | [`[]models.RosettaPublicKey`](../../doc/models/rosetta-public-key.md) | Optional | -                                                                                      |
| `Metadata`          | [`*models.Metadata5`](../../doc/models/metadata-5.md)                 | Optional | -                                                                                      |

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
  "public_keys": [
    {
      "hex_bytes": "hex_bytes8",
      "curve_type": "secp256k1"
    },
    {
      "hex_bytes": "hex_bytes8",
      "curve_type": "secp256k1"
    },
    {
      "hex_bytes": "hex_bytes8",
      "curve_type": "secp256k1"
    }
  ],
  "metadata": {
    "account_sequence": 208,
    "recent_block_hash": "recent_block_hash0"
  }
}
```
