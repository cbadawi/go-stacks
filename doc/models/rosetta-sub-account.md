# Rosetta Sub Account

An account may have state specific to a contract address (ERC-20 token) and/or a stake (delegated balance). The sub_account_identifier should specify which state (if applicable) an account instantiation refers to.

## Structure

`RosettaSubAccount`

## Fields

| Name       | Type           | Tags     | Description                                                                                                                                                                                                                                                               |
| ---------- | -------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Address`  | `string`       | Required | The address may be a cryptographic public key (or some encoding of it) or a provided username.                                                                                                                                                                            |
| `Metadata` | `*interface{}` | Optional | If the SubAccount address is not sufficient to uniquely specify a SubAccount, any other identifying information can be stored here. It is important to note that two SubAccounts with identical addresses but differing metadata will not be considered equal by clients. |

## Example (as JSON)

```json
{
  "address": "address0",
  "metadata": {
    "key1": "val1",
    "key2": "val2"
  }
}
```
