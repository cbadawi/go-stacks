# Rosetta Related Operation

Restrict referenced related_operations to identifier indexes < the current operation_identifier.index. This ensures there exists a clear DAG-structure of relations. Since operations are one-sided, one could imagine relating operations in a single transfer or linking operations in a call tree.

## Structure

`RosettaRelatedOperation`

## Fields

| Name           | Type   | Tags     | Description                                                                                                                                                                                                                 |
| -------------- | ------ | -------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Index`        | `int`  | Required | Describes the index of related operation.                                                                                                                                                                                   |
| `NetworkIndex` | `*int` | Optional | Some blockchains specify an operation index that is essential for client use. network_index should not be populated if there is no notion of an operation index in a blockchain (typically most account-based blockchains). |

## Example (as JSON)

```json
{
  "index": 102,
  "network_index": 18
}
```
