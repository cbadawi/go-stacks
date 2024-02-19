package models

import (
	"encoding/json"
)

// PostCondition1 represents a PostCondition1 struct.
// Post-conditionscan limit the damage done to a user's assets
type PostCondition1 struct {
	Principal PostConditionPrincipal4 `json:"principal"`
	// A fungible condition code encodes a statement being made for either STX or a fungible token, with respect to the originating account.
	ConditionCode PostConditionFungibleConditionCodeEnum `json:"condition_code"`
	Amount        string                                 `json:"amount"`
	Type          string                                 `json:"type"`
	Asset         Asset                                  `json:"asset"`
	AssetValue    AssetValue                             `json:"asset_value"`
}

// MarshalJSON implements the json.Marshaler interface for PostCondition1.
// It customizes the JSON marshaling process for PostCondition1 objects.
func (p *PostCondition1) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the PostCondition1 object to a map representation for JSON marshaling.
func (p *PostCondition1) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["principal"] = p.Principal
	structMap["condition_code"] = p.ConditionCode
	structMap["amount"] = p.Amount
	structMap["type"] = p.Type
	structMap["asset"] = p.Asset
	structMap["asset_value"] = p.AssetValue
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PostCondition1.
// It customizes the JSON unmarshaling process for PostCondition1 objects.
func (p *PostCondition1) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		Principal     PostConditionPrincipal4                `json:"principal"`
		ConditionCode PostConditionFungibleConditionCodeEnum `json:"condition_code"`
		Amount        string                                 `json:"amount"`
		Type          string                                 `json:"type"`
		Asset         Asset                                  `json:"asset"`
		AssetValue    AssetValue                             `json:"asset_value"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	p.Principal = temp.Principal
	p.ConditionCode = temp.ConditionCode
	p.Amount = temp.Amount
	p.Type = temp.Type
	p.Asset = temp.Asset
	p.AssetValue = temp.AssetValue
	return nil
}
