package models

import (
	"encoding/json"
)

// RosettaPublicKey represents a RosettaPublicKey struct.
// PublicKey contains a public key byte array for a particular CurveType encoded in hex. Note that there is no PrivateKey struct as this is NEVER the concern of an implementation.
type RosettaPublicKey struct {
	// Hex-encoded public key bytes in the format specified by the CurveType.
	HexBytes string `json:"hex_bytes"`
	// CurveType is the type of cryptographic curve associated with a PublicKey.
	CurveType CurveTypeEnum `json:"curve_type"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaPublicKey.
// It customizes the JSON marshaling process for RosettaPublicKey objects.
func (r *RosettaPublicKey) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaPublicKey object to a map representation for JSON marshaling.
func (r *RosettaPublicKey) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["hex_bytes"] = r.HexBytes
	structMap["curve_type"] = r.CurveType
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaPublicKey.
// It customizes the JSON unmarshaling process for RosettaPublicKey objects.
func (r *RosettaPublicKey) UnmarshalJSON(input []byte) error {
	temp := &struct {
		HexBytes  string        `json:"hex_bytes"`
		CurveType CurveTypeEnum `json:"curve_type"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.HexBytes = temp.HexBytes
	r.CurveType = temp.CurveType
	return nil
}
