# Rosetta Sync Status

SyncStatus is used to provide additional context about an implementation's sync status. It is often used to indicate that an implementation is healthy when it cannot be queried until some sync phase occurs. If an implementation is immediately queryable, this model is often not populated.

## Structure

`RosettaSyncStatus`

## Fields

| Name           | Type      | Tags     | Description                                                                                                  |
| -------------- | --------- | -------- | ------------------------------------------------------------------------------------------------------------ |
| `CurrentIndex` | `int`     | Required | CurrentIndex is the index of the last synced block in the current stage.                                     |
| `TargetIndex`  | `*int`    | Optional | TargetIndex is the index of the block that the implementation is attempting to sync to in the current stage. |
| `Stage`        | `*string` | Optional | Stage is the phase of the sync process.                                                                      |
| `Synced`       | `*bool`   | Optional | Synced indicates if an implementation has synced up to the most recent block.                                |

## Example (as JSON)

```json
{
  "current_index": 28,
  "target_index": 152,
  "stage": "stage2",
  "synced": false
}
```
