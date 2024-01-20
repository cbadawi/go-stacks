# Rosetta Construction Combine Response

RosettaConstructionCombineResponse is returned by /construction/combine. The network payload will be sent directly to the construction/submit endpoint.

## Structure

`RosettaConstructionCombineResponse`

## Fields

| Name                | Type     | Tags     | Description                     |
| ------------------- | -------- | -------- | ------------------------------- |
| `SignedTransaction` | `string` | Required | Signed transaction bytes in hex |

## Example (as JSON)

```json
{
  "signed_transaction": "signed_transaction2"
}
```
