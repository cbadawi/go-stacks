# Version

The Version object is utilized to inform the client of the versions of different components of the Rosetta implementation.

## Structure

`Version`

## Fields

| Name                | Type           | Tags     | Description                                                                                                                                                    |
| ------------------- | -------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `RosettaVersion`    | `string`       | Required | The rosetta_version is the version of the Rosetta interface the implementation adheres to. This can be useful for clients looking to reliably parse responses. |
| `NodeVersion`       | `string`       | Required | The node_version is the canonical version of the node runtime. This can help clients manage deployments.                                                       |
| `MiddlewareVersion` | `*string`      | Optional | When a middleware server is used to adhere to the Rosetta interface, it should return its version here. This can help clients manage deployments.              |
| `Metadata`          | `*interface{}` | Optional | Any other information that may be useful about versioning of dependent services should be returned here.                                                       |

## Example (as JSON)

```json
{
  "rosetta_version": "rosetta_version6",
  "node_version": "node_version2",
  "middleware_version": "middleware_version8",
  "metadata": {
    "key1": "val1",
    "key2": "val2"
  }
}
```
