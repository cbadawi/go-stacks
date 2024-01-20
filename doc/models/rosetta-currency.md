# Rosetta Currency

Currency is composed of a canonical Symbol and Decimals. This Decimals value is used to convert an Amount.Value from atomic units (Satoshis) to standard units (Bitcoins).

## Structure

`RosettaCurrency`

## Fields

| Name       | Type           | Tags     | Description                                                                                                                                                                                                            |
| ---------- | -------------- | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Symbol`   | `string`       | Required | Canonical symbol associated with a currency.                                                                                                                                                                           |
| `Decimals` | `int`          | Required | Number of decimal places in the standard unit representation of the amount. For example, BTC has 8 decimals. Note that it is not possible to represent the value of some currency in atomic units that is not base 10. |
| `Metadata` | `*interface{}` | Optional | Any additional information related to the currency itself. For example, it would be useful to populate this object with the contract address of an ERC-20 token.                                                       |

## Example (as JSON)

```json
{
  "symbol": "symbol4",
  "decimals": 148,
  "metadata": {
    "key1": "val1",
    "key2": "val2"
  }
}
```
