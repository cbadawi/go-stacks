# Rosetta Block Request

A BlockRequest is utilized to make a block request on the /block endpoint.

## Structure

`RosettaBlockRequest`

## Fields

| Name                | Type                                                                                           | Tags     | Description                                                                                                                                                                                            |
| ------------------- | ---------------------------------------------------------------------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `NetworkIdentifier` | [`models.NetworkIdentifier`](../../doc/models/network-identifier.md)                           | Required | The network_identifier specifies which network a particular object is associated with.                                                                                                                 |
| `BlockIdentifier`   | [`models.RosettaPartialBlockIdentifier`](../../doc/models/rosetta-partial-block-identifier.md) | Required | When fetching data by BlockIdentifier, it may be possible to only specify the index or hash. If neither property is specified, it is assumed that the client is making a request at the current block. |

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
  "block_identifier": {
    "hash": "hash6",
    "index": 54
  }
}
```
