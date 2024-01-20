# Rosetta Network Status Response

NetworkStatusResponse contains basic information about the node's view of a blockchain network. It is assumed that any BlockIdentifier.Index less than or equal to CurrentBlockIdentifier.Index can be queried. If a Rosetta implementation prunes historical state, it should populate the optional oldest_block_identifier field with the oldest block available to query. If this is not populated, it is assumed that the genesis_block_identifier is the oldest queryable block. If a Rosetta implementation performs some pre-sync before it is possible to query blocks, sync_status should be populated so that clients can still monitor healthiness. Without this field, it may appear that the implementation is stuck syncing and needs to be terminated.

## Structure

`RosettaNetworkStatusResponse`

## Fields

| Name                     | Type                                                                                           | Tags     | Description                                                                                                                                                                                                                                                                                      |
| ------------------------ | ---------------------------------------------------------------------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `CurrentBlockIdentifier` | [`models.RosettaBlockIdentifier`](../../doc/models/rosetta-block-identifier.md)                | Required | The block_identifier uniquely identifies a block in a particular network.                                                                                                                                                                                                                        |
| `CurrentBlockTimestamp`  | `int`                                                                                          | Required | The timestamp of the block in milliseconds since the Unix Epoch. The timestamp is stored in milliseconds because some blockchains produce blocks more often than once a second.                                                                                                                  |
| `GenesisBlockIdentifier` | [`models.RosettaGenesisBlockIdentifier`](../../doc/models/rosetta-genesis-block-identifier.md) | Required | The block_identifier uniquely identifies a block in a particular network.                                                                                                                                                                                                                        |
| `OldestBlockIdentifier`  | [`*models.RosettaOldestBlockIdentifier`](../../doc/models/rosetta-oldest-block-identifier.md)  | Optional | The block_identifier uniquely identifies a block in a particular network.                                                                                                                                                                                                                        |
| `SyncStatus`             | [`*models.RosettaSyncStatus`](../../doc/models/rosetta-sync-status.md)                         | Optional | SyncStatus is used to provide additional context about an implementation's sync status. It is often used to indicate that an implementation is healthy when it cannot be queried until some sync phase occurs. If an implementation is immediately queryable, this model is often not populated. |
| `Peers`                  | [`[]models.RosettaPeers`](../../doc/models/rosetta-peers.md)                                   | Required | Peers information                                                                                                                                                                                                                                                                                |

## Example (as JSON)

```json
{
  "current_block_identifier": {
    "hash": "hash4",
    "index": 156
  },
  "current_block_timestamp": 38,
  "genesis_block_identifier": {
    "index": 200,
    "hash": "hash6"
  },
  "peers": [
    {
      "peer_id": "peer_id4",
      "metadata": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "oldest_block_identifier": {
    "index": 210,
    "hash": "hash2"
  },
  "sync_status": {
    "current_index": 52,
    "target_index": 176,
    "stage": "stage2",
    "synced": false
  }
}
```
