# Rosetta Public Key

PublicKey contains a public key byte array for a particular CurveType encoded in hex. Note that there is no PrivateKey struct as this is NEVER the concern of an implementation.

## Structure

`RosettaPublicKey`

## Fields

| Name        | Type                                                          | Tags     | Description                                                               |
| ----------- | ------------------------------------------------------------- | -------- | ------------------------------------------------------------------------- |
| `HexBytes`  | `string`                                                      | Required | Hex-encoded public key bytes in the format specified by the CurveType.    |
| `CurveType` | [`models.CurveTypeEnum`](../../doc/models/curve-type-enum.md) | Required | CurveType is the type of cryptographic curve associated with a PublicKey. |

## Example (as JSON)

```json
{
  "hex_bytes": "hex_bytes2",
  "curve_type": "secp256k1"
}
```
