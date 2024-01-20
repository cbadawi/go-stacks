# Rosetta Network List Response

A NetworkListResponse contains all NetworkIdentifiers that the node can serve information for.

## Structure

`RosettaNetworkListResponse`

## Fields

| Name                 | Type                                                                   | Tags     | Description                                                                            |
| -------------------- | ---------------------------------------------------------------------- | -------- | -------------------------------------------------------------------------------------- |
| `NetworkIdentifiers` | [`[]models.NetworkIdentifier`](../../doc/models/network-identifier.md) | Required | The network_identifier specifies which network a particular object is associated with. |

## Example (as JSON)

```json
{
  "network_identifiers": [
    {
      "blockchain": "blockchain0",
      "network": "network4",
      "sub_network_identifier": {
        "network": "network6",
        "metadata": {
          "producer": "producer4"
        }
      }
    }
  ]
}
```
