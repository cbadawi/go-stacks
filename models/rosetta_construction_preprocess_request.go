package models

import (
	"encoding/json"
)

// RosettaConstructionPreprocessRequest represents a RosettaConstructionPreprocessRequest struct.
// ConstructionPreprocessRequest is passed to the /construction/preprocess endpoint so that a Rosetta implementation can determine which metadata it needs to request for construction
type RosettaConstructionPreprocessRequest struct {
	// The network_identifier specifies which network a particular object is associated with.
	NetworkIdentifier NetworkIdentifier     `json:"network_identifier"`
	Operations        []RosettaOperation    `json:"operations"`
	Metadata          *interface{}          `json:"metadata,omitempty"`
	MaxFee            []RosettaMaxFeeAmount `json:"max_fee,omitempty"`
	// The caller can also provide a suggested fee multiplier to indicate that the suggested fee should be scaled. This may be used to set higher fees for urgent transactions or to pay lower fees when there is less urgency. It is assumed that providing a very low multiplier (like 0.0001) will never lead to a transaction being created with a fee less than the minimum network fee (if applicable). In the case that the caller provides both a max fee and a suggested fee multiplier, the max fee will set an upper bound on the suggested fee (regardless of the multiplier provided).
	SuggestedFeeMultiplier *int `json:"suggested_fee_multiplier,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaConstructionPreprocessRequest.
// It customizes the JSON marshaling process for RosettaConstructionPreprocessRequest objects.
func (r *RosettaConstructionPreprocessRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaConstructionPreprocessRequest object to a map representation for JSON marshaling.
func (r *RosettaConstructionPreprocessRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["network_identifier"] = r.NetworkIdentifier
	structMap["operations"] = r.Operations
	if r.Metadata != nil {
		structMap["metadata"] = r.Metadata
	}
	if r.MaxFee != nil {
		structMap["max_fee"] = r.MaxFee
	}
	if r.SuggestedFeeMultiplier != nil {
		structMap["suggested_fee_multiplier"] = r.SuggestedFeeMultiplier
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaConstructionPreprocessRequest.
// It customizes the JSON unmarshaling process for RosettaConstructionPreprocessRequest objects.
func (r *RosettaConstructionPreprocessRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		NetworkIdentifier      NetworkIdentifier     `json:"network_identifier"`
		Operations             []RosettaOperation    `json:"operations"`
		Metadata               *interface{}          `json:"metadata,omitempty"`
		MaxFee                 []RosettaMaxFeeAmount `json:"max_fee,omitempty"`
		SuggestedFeeMultiplier *int                  `json:"suggested_fee_multiplier,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.NetworkIdentifier = temp.NetworkIdentifier
	r.Operations = temp.Operations
	r.Metadata = temp.Metadata
	r.MaxFee = temp.MaxFee
	r.SuggestedFeeMultiplier = temp.SuggestedFeeMultiplier
	return nil
}
