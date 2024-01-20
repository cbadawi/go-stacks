# Rosetta Coin

If a blockchain is UTXO-based, all unspent Coins owned by an account_identifier should be returned alongside the balance. It is highly recommended to populate this field so that users of the Rosetta API implementation don't need to maintain their own indexer to track their UTXOs.

## Structure

`RosettaCoin`

## Fields

| Name             | Type                                                           | Tags     | Description                                                                                         |
| ---------------- | -------------------------------------------------------------- | -------- | --------------------------------------------------------------------------------------------------- |
| `CoinIdentifier` | [`models.CoinIdentifier`](../../doc/models/coin-identifier.md) | Required | CoinIdentifier uniquely identifies a Coin.                                                          |
| `Amount`         | [`models.RosettaAmount`](../../doc/models/rosetta-amount.md)   | Required | Amount is some Value of a Currency. It is considered invalid to specify a Value without a Currency. |

## Example (as JSON)

```json
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
```
