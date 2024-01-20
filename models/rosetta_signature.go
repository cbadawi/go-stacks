package models

import (
	"encoding/json"
)

// RosettaSignature represents a RosettaSignature struct.
// Signature contains the payload that was signed, the public keys of the keypairs used to produce the signature, the signature (encoded in hex), and the SignatureType. PublicKey is often times not known during construction of the signing payloads but may be needed to combine signatures properly.
type RosettaSignature struct {
	// SigningPayload is signed by the client with the keypair associated with an address using the specified SignatureType. SignatureType can be optionally populated if there is a restriction on the signature scheme that can be used to sign the payload.
	SigningPayload SigningPayload `json:"signing_payload"`
	// PublicKey contains a public key byte array for a particular CurveType encoded in hex. Note that there is no PrivateKey struct as this is NEVER the concern of an implementation.
	PublicKey RosettaPublicKey `json:"public_key"`
	// SignatureType is the type of a cryptographic signature.
	SignatureType SignatureTypeEnum `json:"signature_type"`
	HexBytes      string            `json:"hex_bytes"`
}

// MarshalJSON implements the json.Marshaler interface for RosettaSignature.
// It customizes the JSON marshaling process for RosettaSignature objects.
func (r *RosettaSignature) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RosettaSignature object to a map representation for JSON marshaling.
func (r *RosettaSignature) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["signing_payload"] = r.SigningPayload
	structMap["public_key"] = r.PublicKey
	structMap["signature_type"] = r.SignatureType
	structMap["hex_bytes"] = r.HexBytes
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RosettaSignature.
// It customizes the JSON unmarshaling process for RosettaSignature objects.
func (r *RosettaSignature) UnmarshalJSON(input []byte) error {
	temp := &struct {
		SigningPayload SigningPayload    `json:"signing_payload"`
		PublicKey      RosettaPublicKey  `json:"public_key"`
		SignatureType  SignatureTypeEnum `json:"signature_type"`
		HexBytes       string            `json:"hex_bytes"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.SigningPayload = temp.SigningPayload
	r.PublicKey = temp.PublicKey
	r.SignatureType = temp.SignatureType
	r.HexBytes = temp.HexBytes
	return nil
}
