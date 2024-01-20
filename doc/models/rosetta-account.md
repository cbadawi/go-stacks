# Rosetta Account

The account_identifier uniquely identifies an account within a network. All fields in the account_identifier are utilized to determine this uniqueness (including the metadata field, if populated).

## Structure

`RosettaAccount`

## Fields

| Name         | Type                                                                   | Tags     | Description                                                                                                                                                                                                           |
| ------------ | ---------------------------------------------------------------------- | -------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Address`    | `string`                                                               | Required | The address may be a cryptographic public key (or some encoding of it) or a provided username.                                                                                                                        |
| `SubAccount` | [`*models.RosettaSubAccount`](../../doc/models/rosetta-sub-account.md) | Optional | An account may have state specific to a contract address (ERC-20 token) and/or a stake (delegated balance). The sub_account_identifier should specify which state (if applicable) an account instantiation refers to. |
| `Metadata`   | `*interface{}`                                                         | Optional | Blockchains that utilize a username model (where the address is not a derivative of a cryptographic public key) should specify the public key(s) owned by the address in metadata.                                    |

## Example (as JSON)

```json
{
  "address": "address0",
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
```
