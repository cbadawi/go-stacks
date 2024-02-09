package models

import (
	"encoding/json"
)

// TenureChangePayload represents a TenureChangePayload struct.
type TenureChangePayload struct {
	// Consensus hash of this tenure. Corresponds to the sortition in which the miner of this block was chosen.
	TenureConsensusHash string `json:"tenure_consensus_hash"`
	// Consensus hash of the previous tenure. Corresponds to the sortition of the previous winning block-commit.
	PrevTenureConsensusHash string `json:"prev_tenure_consensus_hash"`
	// Current consensus hash on the underlying burnchain. Corresponds to the last-seen sortition.
	BurnViewConsensusHash string `json:"burn_view_consensus_hash"`
	// (Hex string) Stacks Block hash
	PreviousTenureEnd string `json:"previous_tenure_end"`
	// The number of blocks produced in the previous tenure.
	PreviousTenureBlocks float64 `json:"previous_tenure_blocks"`
	// Cause of change in mining tenure. Depending on cause, tenure can be ended or extended.
	Cause CauseEnum `json:"cause"`
	// (Hex string) The ECDSA public key hash of the current tenure.
	PubkeyHash string `json:"pubkey_hash"`
	// (Hex string) A Schnorr signature from the Stackers.
	Signature string `json:"signature"`
	// (Hex string) A bitmap of which Stackers signed.
	Signers string `json:"signers"`
}

// MarshalJSON implements the json.Marshaler interface for TenureChangePayload.
// It customizes the JSON marshaling process for TenureChangePayload objects.
func (t *TenureChangePayload) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TenureChangePayload object to a map representation for JSON marshaling.
func (t *TenureChangePayload) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["tenure_consensus_hash"] = t.TenureConsensusHash
	structMap["prev_tenure_consensus_hash"] = t.PrevTenureConsensusHash
	structMap["burn_view_consensus_hash"] = t.BurnViewConsensusHash
	structMap["previous_tenure_end"] = t.PreviousTenureEnd
	structMap["previous_tenure_blocks"] = t.PreviousTenureBlocks
	structMap["cause"] = t.Cause
	structMap["pubkey_hash"] = t.PubkeyHash
	structMap["signature"] = t.Signature
	structMap["signers"] = t.Signers
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TenureChangePayload.
// It customizes the JSON unmarshaling process for TenureChangePayload objects.
func (t *TenureChangePayload) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		TenureConsensusHash     string    `json:"tenure_consensus_hash"`
		PrevTenureConsensusHash string    `json:"prev_tenure_consensus_hash"`
		BurnViewConsensusHash   string    `json:"burn_view_consensus_hash"`
		PreviousTenureEnd       string    `json:"previous_tenure_end"`
		PreviousTenureBlocks    float64   `json:"previous_tenure_blocks"`
		Cause                   CauseEnum `json:"cause"`
		PubkeyHash              string    `json:"pubkey_hash"`
		Signature               string    `json:"signature"`
		Signers                 string    `json:"signers"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	t.TenureConsensusHash = temp.TenureConsensusHash
	t.PrevTenureConsensusHash = temp.PrevTenureConsensusHash
	t.BurnViewConsensusHash = temp.BurnViewConsensusHash
	t.PreviousTenureEnd = temp.PreviousTenureEnd
	t.PreviousTenureBlocks = temp.PreviousTenureBlocks
	t.Cause = temp.Cause
	t.PubkeyHash = temp.PubkeyHash
	t.Signature = temp.Signature
	t.Signers = temp.Signers
	return nil
}
