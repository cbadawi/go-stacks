# Rosetta Coin Change

CoinChange is used to represent a change in state of a some coin identified by a coin_identifier. This object is part of the Operation model and must be populated for UTXO-based blockchains. Coincidentally, this abstraction of UTXOs allows for supporting both account-based transfers and UTXO-based transfers on the same blockchain (when a transfer is account-based, don't populate this model).

## Structure

`RosettaCoinChange`

## Fields

| Name             | Type                                                            | Tags     | Description                                                                                                                                                                                                                      |
| ---------------- | --------------------------------------------------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `CoinIdentifier` | [`models.CoinIdentifier`](../../doc/models/coin-identifier.md)  | Required | CoinIdentifier uniquely identifies a Coin.                                                                                                                                                                                       |
| `CoinAction`     | [`models.CoinActionEnum`](../../doc/models/coin-action-enum.md) | Required | CoinActions are different state changes that a Coin can undergo. When a Coin is created, it is coin_created. When a Coin is spent, it is coin_spent. It is assumed that a single Coin cannot be created or spent more than once. |

## Example (as JSON)

```json
{
  "coin_identifier": {
    "identifier": "identifier8"
  },
  "coin_action": "coin_created"
}
```
