# Tenure Change Payload

## Structure

`TenureChangePayload`

## Fields

| Name                      | Type                                                 | Tags     | Description                                                                                               |
| ------------------------- | ---------------------------------------------------- | -------- | --------------------------------------------------------------------------------------------------------- |
| `TenureConsensusHash`     | `string`                                             | Required | Consensus hash of this tenure. Corresponds to the sortition in which the miner of this block was chosen.  |
| `PrevTenureConsensusHash` | `string`                                             | Required | Consensus hash of the previous tenure. Corresponds to the sortition of the previous winning block-commit. |
| `BurnViewConsensusHash`   | `string`                                             | Required | Current consensus hash on the underlying burnchain. Corresponds to the last-seen sortition.               |
| `PreviousTenureEnd`       | `string`                                             | Required | (Hex string) Stacks Block hash                                                                            |
| `PreviousTenureBlocks`    | `float64`                                            | Required | The number of blocks produced in the previous tenure.                                                     |
| `Cause`                   | [`models.CauseEnum`](../../doc/models/cause-enum.md) | Required | Cause of change in mining tenure. Depending on cause, tenure can be ended or extended.                    |
| `PubkeyHash`              | `string`                                             | Required | (Hex string) The ECDSA public key hash of the current tenure.                                             |
| `Signature`               | `string`                                             | Required | (Hex string) A Schnorr signature from the Stackers.                                                       |
| `Signers`                 | `string`                                             | Required | (Hex string) A bitmap of which Stackers signed.                                                           |

## Example (as JSON)

```json
{
  "tenure_consensus_hash": "tenure_consensus_hash2",
  "prev_tenure_consensus_hash": "prev_tenure_consensus_hash2",
  "burn_view_consensus_hash": "burn_view_consensus_hash2",
  "previous_tenure_end": "previous_tenure_end8",
  "previous_tenure_blocks": 105.18,
  "cause": "block_found",
  "pubkey_hash": "pubkey_hash4",
  "signature": "signature2",
  "signers": "signers2"
}
```
