# Rosetta Signature

Signature contains the payload that was signed, the public keys of the keypairs used to produce the signature, the signature (encoded in hex), and the SignatureType. PublicKey is often times not known during construction of the signing payloads but may be needed to combine signatures properly.

## Structure

`RosettaSignature`

## Fields

| Name             | Type                                                                  | Tags     | Description                                                                                                                                                                                                                                             |
| ---------------- | --------------------------------------------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `SigningPayload` | [`models.SigningPayload`](../../doc/models/signing-payload.md)        | Required | SigningPayload is signed by the client with the keypair associated with an address using the specified SignatureType. SignatureType can be optionally populated if there is a restriction on the signature scheme that can be used to sign the payload. |
| `PublicKey`      | [`models.RosettaPublicKey`](../../doc/models/rosetta-public-key.md)   | Required | PublicKey contains a public key byte array for a particular CurveType encoded in hex. Note that there is no PrivateKey struct as this is NEVER the concern of an implementation.                                                                        |
| `SignatureType`  | [`models.SignatureTypeEnum`](../../doc/models/signature-type-enum.md) | Required | SignatureType is the type of a cryptographic signature.                                                                                                                                                                                                 |
| `HexBytes`       | `string`                                                              | Required | -                                                                                                                                                                                                                                                       |

## Example (as JSON)

```json
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
  "signature_type": "ecdsa",
  "hex_bytes": "hex_bytes4"
}
```
