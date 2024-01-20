# Rosetta Construction Parse Request

Parse is called on both unsigned and signed transactions to understand the intent of the formulated transaction. This is run as a sanity check before signing (after /construction/payloads) and before broadcast (after /construction/combine).

## Structure

`RosettaConstructionParseRequest`

## Fields

| Name                | Type                                                                 | Tags     | Description                                                                                                                                            |
| ------------------- | -------------------------------------------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `NetworkIdentifier` | [`models.NetworkIdentifier`](../../doc/models/network-identifier.md) | Required | The network_identifier specifies which network a particular object is associated with.                                                                 |
| `Signed`            | `bool`                                                               | Required | Signed is a boolean indicating whether the transaction is signed.                                                                                      |
| `Transaction`       | `string`                                                             | Required | This must be either the unsigned transaction blob returned by /construction/payloads or the signed transaction blob returned by /construction/combine. |

## Example (as JSON)

```json
{
  "network_identifier": {
    "blockchain": "blockchain2",
    "network": "network6",
    "sub_network_identifier": {
      "network": "network6",
      "metadata": {
        "producer": "producer4"
      }
    }
  },
  "signed": false,
  "transaction": "transaction6"
}
```
