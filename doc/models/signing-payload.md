# Signing Payload

SigningPayload is signed by the client with the keypair associated with an address using the specified SignatureType. SignatureType can be optionally populated if there is a restriction on the signature scheme that can be used to sign the payload.

## Structure

`SigningPayload`

## Fields

| Name                | Type                                                                   | Tags     | Description                                                                                                                                                                                          |
| ------------------- | ---------------------------------------------------------------------- | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Address`           | `*string`                                                              | Optional | [DEPRECATED by account_identifier in v1.4.4] The network-specific address of the account that should sign the payload.                                                                               |
| `AccountIdentifier` | [`*models.RosettaAccount`](../../doc/models/rosetta-account.md)        | Optional | The account_identifier uniquely identifies an account within a network. All fields in the account_identifier are utilized to determine this uniqueness (including the metadata field, if populated). |
| `HexBytes`          | `string`                                                               | Required | -                                                                                                                                                                                                    |
| `SignatureType`     | [`*models.SignatureTypeEnum`](../../doc/models/signature-type-enum.md) | Optional | SignatureType is the type of a cryptographic signature.                                                                                                                                              |

## Example (as JSON)

```json
{
  "address": "address6",
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
  "hex_bytes": "hex_bytes2",
  "signature_type": "ed25519"
}
```
