# Rosetta Construction Metadata Response

The ConstructionMetadataResponse returns network-specific metadata used for transaction construction. Optionally, the implementer can return the suggested fee associated with the transaction being constructed. The caller may use this info to adjust the intent of the transaction or to create a transaction with a different account that can pay the suggested fee. Suggested fee is an array in case fee payment must occur in multiple currencies.

## Structure

`RosettaConstructionMetadataResponse`

## Fields

| Name           | Type                                                           | Tags     | Description |
| -------------- | -------------------------------------------------------------- | -------- | ----------- |
| `Metadata`     | [`models.Metadata5`](../../doc/models/metadata-5.md)           | Required | -           |
| `SuggestedFee` | [`[]models.RosettaAmount`](../../doc/models/rosetta-amount.md) | Optional | -           |

## Example (as JSON)

```json
{
  "metadata": {
    "account_sequence": 208,
    "recent_block_hash": "recent_block_hash0"
  },
  "suggested_fee": [
    {
      "value": "value6",
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
    {
      "value": "value6",
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
    {
      "value": "value6",
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
    }
  ]
}
```
