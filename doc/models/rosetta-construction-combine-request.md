# Rosetta Construction Combine Request

RosettaConstructionCombineRequest is the input to the /construction/combine endpoint. It contains the unsigned transaction blob returned by /construction/payloads and all required signatures to create a network transaction.

## Structure

`RosettaConstructionCombineRequest`

## Fields

| Name                  | Type                                                                 | Tags     | Description                                                                            |
| --------------------- | -------------------------------------------------------------------- | -------- | -------------------------------------------------------------------------------------- |
| `NetworkIdentifier`   | [`models.NetworkIdentifier`](../../doc/models/network-identifier.md) | Required | The network_identifier specifies which network a particular object is associated with. |
| `UnsignedTransaction` | `string`                                                             | Required | -                                                                                      |
| `Signatures`          | [`[]models.RosettaSignature`](../../doc/models/rosetta-signature.md) | Required | -                                                                                      |

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
  "unsigned_transaction": "unsigned_transaction2",
  "signatures": [
    {
      "signing_payload": {
        "address": "address4",
        "account_identifier": {
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
        "hex_bytes": "hex_bytes4",
        "signature_type": "ecdsa"
      },
      "public_key": {
        "hex_bytes": "hex_bytes4",
        "curve_type": "secp256k1"
      },
      "signature_type": "schnorr_poseidon",
      "hex_bytes": "hex_bytes0"
    }
  ]
}
```
