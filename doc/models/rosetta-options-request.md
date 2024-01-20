# Rosetta Options Request

This endpoint returns the version information and allowed network-specific types for a NetworkIdentifier. Any NetworkIdentifier returned by /network/list should be accessible here. Because options are retrievable in the context of a NetworkIdentifier, it is possible to define unique options for each network.

## Structure

`RosettaOptionsRequest`

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
