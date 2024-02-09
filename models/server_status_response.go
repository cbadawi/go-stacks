package models

import (
	"encoding/json"
)

// ServerStatusResponse represents a ServerStatusResponse struct.
// GET blockchain API status
type ServerStatusResponse struct {
	// the server version that is currently running
	ServerVersion *string `json:"server_version,omitempty"`
	// the current server status
	Status            string        `json:"status"`
	PoxV1UnlockHeight Optional[int] `json:"pox_v1_unlock_height"`
	PoxV2UnlockHeight Optional[int] `json:"pox_v2_unlock_height"`
	PoxV3UnlockHeight Optional[int] `json:"pox_v3_unlock_height"`
	// Current chain tip information
	ChainTip *ChainTip `json:"chain_tip,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ServerStatusResponse.
// It customizes the JSON marshaling process for ServerStatusResponse objects.
func (s *ServerStatusResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the ServerStatusResponse object to a map representation for JSON marshaling.
func (s *ServerStatusResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if s.ServerVersion != nil {
		structMap["server_version"] = s.ServerVersion
	}
	structMap["status"] = s.Status
	if s.PoxV1UnlockHeight.IsValueSet() {
		structMap["pox_v1_unlock_height"] = s.PoxV1UnlockHeight.Value()
	}
	if s.PoxV2UnlockHeight.IsValueSet() {
		structMap["pox_v2_unlock_height"] = s.PoxV2UnlockHeight.Value()
	}
	if s.PoxV3UnlockHeight.IsValueSet() {
		structMap["pox_v3_unlock_height"] = s.PoxV3UnlockHeight.Value()
	}
	if s.ChainTip != nil {
		structMap["chain_tip"] = s.ChainTip
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServerStatusResponse.
// It customizes the JSON unmarshaling process for ServerStatusResponse objects.
func (s *ServerStatusResponse) UnmarshalJSON(input []byte) error {
	if input[0] == '"' {
		return unmarshalResponseString(input)
	}
	temp := &struct {
		ServerVersion     *string       `json:"server_version,omitempty"`
		Status            string        `json:"status"`
		PoxV1UnlockHeight Optional[int] `json:"pox_v1_unlock_height"`
		PoxV2UnlockHeight Optional[int] `json:"pox_v2_unlock_height"`
		PoxV3UnlockHeight Optional[int] `json:"pox_v3_unlock_height"`
		ChainTip          *ChainTip     `json:"chain_tip,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.ServerVersion = temp.ServerVersion
	s.Status = temp.Status
	s.PoxV1UnlockHeight = temp.PoxV1UnlockHeight
	s.PoxV2UnlockHeight = temp.PoxV2UnlockHeight
	s.PoxV3UnlockHeight = temp.PoxV3UnlockHeight
	s.ChainTip = temp.ChainTip
	return nil
}
