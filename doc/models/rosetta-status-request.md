# Rosetta Status Request

This endpoint returns the current status of the network requested. Any NetworkIdentifier returned by /network/list should be accessible here.

## Structure

`RosettaStatusRequest`

## Fields

| Name                | Type                                                                 | Tags     | Description                                                                            |
| ------------------- | -------------------------------------------------------------------- | -------- | -------------------------------------------------------------------------------------- |
| `NetworkIdentifier` | [`models.NetworkIdentifier`](../../doc/models/network-identifier.md) | Required | The network_identifier specifies which network a particular object is associated with. |
| `Metadata`          | `*interface{}`                                                       | Optional | -                                                                                      |

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
  "metadata": {
    "key1": "val1",
    "key2": "val2"
  }
}
```
