# Rosetta Network Options Response

NetworkOptionsResponse contains information about the versioning of the node and the allowed operation statuses, operation types, and errors.

## Structure

`RosettaNetworkOptionsResponse`

## Fields

| Name      | Type                                            | Tags     | Description                                                                                                                                                                                                                                                                                                                                          |
| --------- | ----------------------------------------------- | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Version` | [`models.Version`](../../doc/models/version.md) | Required | The Version object is utilized to inform the client of the versions of different components of the Rosetta implementation.                                                                                                                                                                                                                           |
| `Allow`   | [`models.Allow`](../../doc/models/allow.md)     | Required | Allow specifies supported Operation status, Operation types, and all possible error statuses. This Allow object is used by clients to validate the correctness of a Rosetta Server implementation. It is expected that these clients will error if they receive some response that contains any of the above information that is not specified here. |

## Example (as JSON)

```json
{
  "version": {
    "rosetta_version": "rosetta_version4",
    "node_version": "node_version8",
    "middleware_version": "middleware_version8",
    "metadata": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "allow": {
    "operation_statuses": [
      {
        "status": "status0",
        "successful": false
      }
    ],
    "operation_types": ["operation_types7"],
    "errors": [
      {
        "code": 144,
        "message": "message0",
        "retriable": false
      }
    ],
    "historical_balance_lookup": false
  }
}
```
