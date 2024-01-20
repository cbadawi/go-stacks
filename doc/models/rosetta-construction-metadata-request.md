# Rosetta Construction Metadata Request

A ConstructionMetadataRequest is utilized to get information required to construct a transaction. The Options object used to specify which metadata to return is left purposely unstructured to allow flexibility for implementers. Optionally, the request can also include an array of PublicKeys associated with the AccountIdentifiers returned in ConstructionPreprocessResponse.

## Structure

`RosettaConstructionMetadataRequest`

## Fields

| Name                | Type                                                                  | Tags     | Description                                                                            |
| ------------------- | --------------------------------------------------------------------- | -------- | -------------------------------------------------------------------------------------- |
| `NetworkIdentifier` | [`models.NetworkIdentifier`](../../doc/models/network-identifier.md)  | Required | The network_identifier specifies which network a particular object is associated with. |
| `Options`           | [`models.RosettaOptions`](../../doc/models/rosetta-options.md)        | Required | The options that will be sent directly to /construction/metadata by the caller.        |
| `PublicKeys`        | [`[]models.RosettaPublicKey`](../../doc/models/rosetta-public-key.md) | Optional | -                                                                                      |

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
  "options": {
    "sender_address": "sender_address4",
    "type": "type8",
    "status": "status6",
    "token_transfer_recipient_address": "token_transfer_recipient_address8",
    "amount": "amount4"
  },
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
  ]
}
```
