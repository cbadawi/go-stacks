# Coin Identifier

CoinIdentifier uniquely identifies a Coin.

## Structure

`CoinIdentifier`

## Fields

| Name         | Type     | Tags     | Description                                                                                                                              |
| ------------ | -------- | -------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `Identifier` | `string` | Required | Identifier should be populated with a globally unique identifier of a Coin. In Bitcoin, this identifier would be transaction_hash:index. |

## Example (as JSON)

```json
{
  "identifier": "identifier4"
}
```
