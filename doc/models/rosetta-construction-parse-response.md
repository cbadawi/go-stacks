# Rosetta Construction Parse Response

RosettaConstructionParseResponse contains an array of operations that occur in a transaction blob. This should match the array of operations provided to /construction/preprocess and /construction/payloads.

## Structure

`RosettaConstructionParseResponse`

## Fields

| Name                       | Type                                                                                  | Tags     | Description                                                                                                                                                   |
| -------------------------- | ------------------------------------------------------------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Operations`               | [`[]models.RosettaOperation`](../../doc/models/rosetta-operation.md)                  | Required | -                                                                                                                                                             |
| `Signers`                  | `[]string`                                                                            | Optional | [DEPRECATED by account_identifier_signers in v1.4.4] All signers (addresses) of a particular transaction. If the transaction is unsigned, it should be empty. |
| `AccountIdentifierSigners` | [`[]models.RosettaAccountIdentifier`](../../doc/models/rosetta-account-identifier.md) | Optional | -                                                                                                                                                             |
| `Metadata`                 | `*interface{}`                                                                        | Optional | -                                                                                                                                                             |

## Example (as JSON)

```json
{
  "operations": [
    {
      "operation_identifier": {
        "index": 254,
        "network_index": 122
      },
      "related_operations": [
        {
          "index": 246,
          "network_index": 130
        }
      ],
      "type": "type8",
      "status": "status6",
      "account": {
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
      "amount": {
        "value": "value0",
        "currency": {
          "symbol": "symbol8",
          "decimals": 194,
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
      "coin_change": {
        "coin_identifier": {
          "identifier": "identifier8"
        },
        "coin_action": "coin_created"
      }
    }
  ],
  "signers": ["signers2"],
  "account_identifier_signers": [
    {
      "address": "address4",
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
  ],
  "metadata": {
    "key1": "val1",
    "key2": "val2"
  }
}
```
