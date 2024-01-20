# Rosetta Options

The options that will be sent directly to /construction/metadata by the caller.

## Structure

`RosettaOptions`

## Fields

| Name                            | Type       | Tags     | Description                                                                                                                                                                                  |
| ------------------------------- | ---------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `SenderAddress`                 | `*string`  | Optional | sender's address                                                                                                                                                                             |
| `Type`                          | `*string`  | Optional | Type of operation e.g transfer                                                                                                                                                               |
| `Status`                        | `*string`  | Optional | This value indicates the state of the operations                                                                                                                                             |
| `TokenTransferRecipientAddress` | `*string`  | Optional | Recipient's address                                                                                                                                                                          |
| `Amount`                        | `*string`  | Optional | Amount to be transfered.                                                                                                                                                                     |
| `Symbol`                        | `*string`  | Optional | Currency symbol e.g STX                                                                                                                                                                      |
| `Decimals`                      | `*int`     | Optional | Number of decimal places                                                                                                                                                                     |
| `GasLimit`                      | `*float64` | Optional | Maximum price a user is willing to pay.                                                                                                                                                      |
| `GasPrice`                      | `*float64` | Optional | Cost necessary to perform a transaction on the network                                                                                                                                       |
| `SuggestedFeeMultiplier`        | `*float64` | Optional | A suggested fee multiplier to indicate that the suggested fee should be scaled. This may be used to set higher fees for urgent transactions or to pay lower fees when there is less urgency. |
| `MaxFee`                        | `*string`  | Optional | Maximum fee user is willing to pay                                                                                                                                                           |
| `Fee`                           | `*string`  | Optional | Fee for this transaction                                                                                                                                                                     |
| `Size`                          | `*int`     | Optional | Transaction approximative size (used to calculate total fee).                                                                                                                                |
| `Memo`                          | `*string`  | Optional | STX token transfer memo.                                                                                                                                                                     |
| `NumberOfCycles`                | `*int`     | Optional | Number of cycles when stacking.                                                                                                                                                              |
| `ContractAddress`               | `*string`  | Optional | Address of the contract to call.                                                                                                                                                             |
| `ContractName`                  | `*string`  | Optional | Name of the contract to call.                                                                                                                                                                |
| `BurnBlockHeight`               | `*int`     | Optional | Set the burnchain (BTC) block for stacking lock to start.                                                                                                                                    |
| `DelegateTo`                    | `*string`  | Optional | Delegator address for when calling `delegate-stacking`.                                                                                                                                      |
| `PoxAddr`                       | `*string`  | Optional | The reward address for stacking transaction. It should be a valid Bitcoin address                                                                                                            |

## Example (as JSON)

```json
{
  "sender_address": "sender_address6",
  "type": "type6",
  "status": "status4",
  "token_transfer_recipient_address": "token_transfer_recipient_address6",
  "amount": "amount6"
}
```
