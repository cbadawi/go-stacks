# Rosetta Construction Payload Response

RosettaConstructionPayloadResponse is returned by /construction/payloads. It contains an unsigned transaction blob (that is usually needed to construct the a network transaction from a collection of signatures) and an array of payloads that must be signed by the caller.

## Structure

`RosettaConstructionPayloadResponse`

## Fields

| Name                  | Type                                                             | Tags     | Description                                                                                                                          |
| --------------------- | ---------------------------------------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------ |
| `UnsignedTransaction` | `string`                                                         | Required | This is an unsigned transaction blob (that is usually needed to construct the a network transaction from a collection of signatures) |
| `Payloads`            | [`[]models.SigningPayload`](../../doc/models/signing-payload.md) | Required | An array of payloads that must be signed by the caller                                                                               |

## Example (as JSON)

```json
{
  "unsigned_transaction": "unsigned_transaction8",
  "payloads": [
    {
      "address": "address2",
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
      "hex_bytes": "hex_bytes6",
      "signature_type": "schnorr_1"
    }
  ]
}
```
