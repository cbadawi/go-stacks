# Rosetta Construction Derive Response

ConstructionDeriveResponse is returned by the /construction/derive endpoint.

## Structure

`RosettaConstructionDeriveResponse`

## Fields

| Name                | Type                                                                                 | Tags     | Description                                                                                                                                                                                          |
| ------------------- | ------------------------------------------------------------------------------------ | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Address`           | `*string`                                                                            | Optional | [DEPRECATED by account_identifier in v1.4.4] Address in network-specific format.                                                                                                                     |
| `AccountIdentifier` | [`*models.RosettaAccountIdentifier`](../../doc/models/rosetta-account-identifier.md) | Optional | The account_identifier uniquely identifies an account within a network. All fields in the account_identifier are utilized to determine this uniqueness (including the metadata field, if populated). |
| `Metadata`          | `*interface{}`                                                                       | Optional | -                                                                                                                                                                                                    |

## Example (as JSON)

```json
{
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
  "metadata": {
    "key1": "val1",
    "key2": "val2"
  }
}
```
