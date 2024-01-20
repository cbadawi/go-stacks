# Rosetta Construction Preprocess Response

RosettaConstructionPreprocessResponse contains options that will be sent unmodified to /construction/metadata. If it is not necessary to make a request to /construction/metadata, options should be omitted. Some blockchains require the PublicKey of particular AccountIdentifiers to construct a valid transaction. To fetch these PublicKeys, populate required_public_keys with the AccountIdentifiers associated with the desired PublicKeys. If it is not necessary to retrieve any PublicKeys for construction, required_public_keys should be omitted.

## Structure

`RosettaConstructionPreprocessResponse`

## Fields

| Name                 | Type                                                             | Tags     | Description                                                                     |
| -------------------- | ---------------------------------------------------------------- | -------- | ------------------------------------------------------------------------------- |
| `Options`            | [`*models.RosettaOptions`](../../doc/models/rosetta-options.md)  | Optional | The options that will be sent directly to /construction/metadata by the caller. |
| `RequiredPublicKeys` | [`[]models.RosettaAccount`](../../doc/models/rosetta-account.md) | Optional | -                                                                               |

## Example (as JSON)

```json
{
  "options": {
    "sender_address": "sender_address4",
    "type": "type8",
    "status": "status6",
    "token_transfer_recipient_address": "token_transfer_recipient_address8",
    "amount": "amount4"
  },
  "required_public_keys": [
    {
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
    }
  ]
}
```
