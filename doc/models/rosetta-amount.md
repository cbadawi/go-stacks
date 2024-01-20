# Rosetta Amount

Amount is some Value of a Currency. It is considered invalid to specify a Value without a Currency.

## Structure

`RosettaAmount`

## Fields

| Name       | Type                                                             | Tags     | Description                                                                                                                                                                |
| ---------- | ---------------------------------------------------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Value`    | `string`                                                         | Required | Value of the transaction in atomic units represented as an arbitrary-sized signed integer. For example, 1 BTC would be represented by a value of 100000000.                |
| `Currency` | [`models.RosettaCurrency`](../../doc/models/rosetta-currency.md) | Required | Currency is composed of a canonical Symbol and Decimals. This Decimals value is used to convert an Amount.Value from atomic units (Satoshis) to standard units (Bitcoins). |
| `Metadata` | `*interface{}`                                                   | Optional | -                                                                                                                                                                          |

## Example (as JSON)

```json
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
```
