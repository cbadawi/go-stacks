# Rosetta Construction Hash Request

TransactionHash returns the network-specific transaction hash for a signed transaction.

## Structure

`RosettaConstructionHashRequest`

## Fields

| Name                | Type                                                                 | Tags     | Description                                                                            |
| ------------------- | -------------------------------------------------------------------- | -------- | -------------------------------------------------------------------------------------- |
| `NetworkIdentifier` | [`models.NetworkIdentifier`](../../doc/models/network-identifier.md) | Required | The network_identifier specifies which network a particular object is associated with. |
| `SignedTransaction` | `string`                                                             | Required | Signed transaction                                                                     |

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
  "signed_transaction": "signed_transaction8"
}
```
