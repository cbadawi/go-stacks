# Rosetta Peers

A Peer is a representation of a node's peer.

## Structure

`RosettaPeers`

## Fields

| Name       | Type           | Tags     | Description |
| ---------- | -------------- | -------- | ----------- |
| `PeerId`   | `string`       | Required | peer id     |
| `Metadata` | `*interface{}` | Optional | meta data   |

## Example (as JSON)

```json
{
  "peer_id": "peer_id2",
  "metadata": {
    "key1": "val1",
    "key2": "val2"
  }
}
```
