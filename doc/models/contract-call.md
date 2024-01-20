# Contract Call

## Structure

`ContractCall`

## Fields

| Name                | Type                                                       | Tags     | Description                                                                                |
| ------------------- | ---------------------------------------------------------- | -------- | ------------------------------------------------------------------------------------------ |
| `ContractId`        | `string`                                                   | Required | Contract identifier formatted as `<principaladdress>.<contract_name>`                      |
| `FunctionName`      | `string`                                                   | Required | Name of the Clarity function to be invoked                                                 |
| `FunctionSignature` | `string`                                                   | Required | Function definition, including function name and type as well as parameter names and types |
| `FunctionArgs`      | [`[]models.FunctionArg`](../../doc/models/function-arg.md) | Optional | List of arguments used to invoke the function                                              |

## Example (as JSON)

```json
{
  "contract_id": "contract_id6",
  "function_name": "function_name6",
  "function_signature": "function_signature2",
  "function_args": [
    {
      "hex": "hex8",
      "repr": "repr2",
      "name": "name8",
      "type": "type2"
    }
  ]
}
```
