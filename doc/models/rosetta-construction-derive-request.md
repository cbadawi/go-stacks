# Rosetta Construction Derive Request

Network is provided in the request because some blockchains have different address formats for different networks

## Structure

`RosettaConstructionDeriveRequest`

## Fields

| Name                | Type                                                                 | Tags     | Description                                                                                                                                                                      |
| ------------------- | -------------------------------------------------------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `NetworkIdentifier` | [`models.NetworkIdentifier`](../../doc/models/network-identifier.md) | Required | The network_identifier specifies which network a particular object is associated with.                                                                                           |
| `PublicKey`         | [`models.RosettaPublicKey`](../../doc/models/rosetta-public-key.md)  | Required | PublicKey contains a public key byte array for a particular CurveType encoded in hex. Note that there is no PrivateKey struct as this is NEVER the concern of an implementation. |
| `Metadata`          | `*interface{}`                                                       | Optional | -                                                                                                                                                                                |

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
  "public_key": {
    "hex_bytes": "hex_bytes4",
    "curve_type": "secp256k1"
  },
  "metadata": {
    "key1": "val1",
    "key2": "val2"
  }
}
```
