# Rosetta Account Balance Response

An AccountBalanceResponse is returned on the /account/balance endpoint. If an account has a balance for each AccountIdentifier describing it (ex: an ERC-20 token balance on a few smart contracts), an account balance request must be made with each AccountIdentifier.

## Structure

`RosettaAccountBalanceResponse`

## Fields

| Name              | Type                                                                            | Tags     | Description                                                                                                                                                                                                                                                                              |
| ----------------- | ------------------------------------------------------------------------------- | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `BlockIdentifier` | [`models.RosettaBlockIdentifier`](../../doc/models/rosetta-block-identifier.md) | Required | The block_identifier uniquely identifies a block in a particular network.                                                                                                                                                                                                                |
| `Balances`        | [`[]models.RosettaAmount`](../../doc/models/rosetta-amount.md)                  | Required | A single account balance may have multiple currencies                                                                                                                                                                                                                                    |
| `Coins`           | [`[]models.RosettaCoin`](../../doc/models/rosetta-coin.md)                      | Optional | If a blockchain is UTXO-based, all unspent Coins owned by an account_identifier should be returned alongside the balance. It is highly recommended to populate this field so that users of the Rosetta API implementation don't need to maintain their own indexer to track their UTXOs. |
| `Metadata`        | [`*models.Metadata2`](../../doc/models/metadata-2.md)                           | Optional | Account-based blockchains that utilize a nonce or sequence number should include that number in the metadata. This number could be unique to the identifier or global across the account address.                                                                                        |

## Example (as JSON)

```json
{
  "block_identifier": {
    "hash": "hash6",
    "index": 54
  },
  "balances": [
    {
      "value": "value2",
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
  ],
  "coins": [
    {
      "coin_identifier": {
        "identifier": "identifier8"
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
      }
    }
  ],
  "metadata": {
    "sequence_number": 86
  }
}
```
