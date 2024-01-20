# Core Node Info Response

GET request that core node information

## Structure

`CoreNodeInfoResponse`

## Fields

| Name                     | Type     | Tags     | Description                                                                                                                                                                                                                             |
| ------------------------ | -------- | -------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `PeerVersion`            | `int`    | Required | identifies the version number for the networking communication, this should not change while a node is running, and will only change if there's an upgrade                                                                              |
| `PoxConsensus`           | `string` | Required | is a hash used to identify the burnchain view for a node. it incorporates bitcoin chain information and PoX information. nodes that disagree on this value will appear to each other as forks. this value will change after every block |
| `BurnBlockHeight`        | `int`    | Required | latest bitcoin chain height                                                                                                                                                                                                             |
| `StablePoxConsensus`     | `string` | Required | same as burn_consensus, but evaluated at stable_burn_block_height                                                                                                                                                                       |
| `StableBurnBlockHeight`  | `int`    | Required | leftover from stacks 1.0, basically always burn_block_height - 1                                                                                                                                                                        |
| `ServerVersion`          | `string` | Required | is a version descriptor                                                                                                                                                                                                                 |
| `NetworkId`              | `int`    | Required | is similar to peer_version and will be used to differentiate between different testnets. this value will be different between mainnet and testnet. once launched, this value will not change                                            |
| `ParentNetworkId`        | `int`    | Required | same as network_id, but for bitcoin                                                                                                                                                                                                     |
| `StacksTipHeight`        | `int`    | Required | the latest Stacks chain height. Stacks forks can occur independent of the Bitcoin chain, that height doesn't increase 1-to-1 with the Bitcoin height                                                                                    |
| `StacksTip`              | `string` | Required | the best known block hash for the Stack chain (not including any pending microblocks)                                                                                                                                                   |
| `StacksTipConsensusHash` | `string` | Required | the burn chain (i.e., bitcoin) consensus hash at the time that stacks_tip was mined                                                                                                                                                     |
| `UnanchoredTip`          | `string` | Required | the latest microblock hash if any microblocks were processed. if no microblock has been processed for the current block, a 000.., hex array is returned                                                                                 |
| `ExitAtBlockHeight`      | `int`    | Required | the block height at which the testnet network will be reset. not applicable for mainnet                                                                                                                                                 |

## Example (as JSON)

```json
{
  "peer_version": 194,
  "pox_consensus": "pox_consensus0",
  "burn_block_height": 16,
  "stable_pox_consensus": "stable_pox_consensus2",
  "stable_burn_block_height": 50,
  "server_version": "server_version8",
  "network_id": 158,
  "parent_network_id": 48,
  "stacks_tip_height": 154,
  "stacks_tip": "stacks_tip6",
  "stacks_tip_consensus_hash": "stacks_tip_consensus_hash6",
  "unanchored_tip": "unanchored_tip6",
  "exit_at_block_height": 66
}
```
