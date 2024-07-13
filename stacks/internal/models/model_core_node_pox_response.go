package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the CoreNodePoxResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CoreNodePoxResponse{}

// CoreNodePoxResponse Get Proof of Transfer (PoX) information
type CoreNodePoxResponse struct {
	ContractId                 string `json:"contract_id"`
	FirstBurnchainBlockHeight  int32  `json:"first_burnchain_block_height"`
	MinAmountUstx              int32  `json:"min_amount_ustx"`
	RegistrationWindowLength   int32  `json:"registration_window_length"`
	RejectionFraction          int32  `json:"rejection_fraction"`
	RewardCycleId              int32  `json:"reward_cycle_id"`
	RewardCycleLength          int32  `json:"reward_cycle_length"`
	RejectionVotesLeftRequired int32  `json:"rejection_votes_left_required"`
	TotalLiquidSupplyUstx      int32  `json:"total_liquid_supply_ustx"`
}

type _CoreNodePoxResponse CoreNodePoxResponse

// NewCoreNodePoxResponse instantiates a new CoreNodePoxResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreNodePoxResponse(contractId string, firstBurnchainBlockHeight int32, minAmountUstx int32, registrationWindowLength int32, rejectionFraction int32, rewardCycleId int32, rewardCycleLength int32, rejectionVotesLeftRequired int32, totalLiquidSupplyUstx int32) *CoreNodePoxResponse {
	this := CoreNodePoxResponse{}
	this.ContractId = contractId
	this.FirstBurnchainBlockHeight = firstBurnchainBlockHeight
	this.MinAmountUstx = minAmountUstx
	this.RegistrationWindowLength = registrationWindowLength
	this.RejectionFraction = rejectionFraction
	this.RewardCycleId = rewardCycleId
	this.RewardCycleLength = rewardCycleLength
	this.RejectionVotesLeftRequired = rejectionVotesLeftRequired
	this.TotalLiquidSupplyUstx = totalLiquidSupplyUstx
	return &this
}

// NewCoreNodePoxResponseWithDefaults instantiates a new CoreNodePoxResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreNodePoxResponseWithDefaults() *CoreNodePoxResponse {
	this := CoreNodePoxResponse{}
	return &this
}

// GetContractId returns the ContractId field value
func (o *CoreNodePoxResponse) GetContractId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value
// and a boolean to check if the value has been set.
func (o *CoreNodePoxResponse) GetContractIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractId, true
}

// SetContractId sets field value
func (o *CoreNodePoxResponse) SetContractId(v string) {
	o.ContractId = v
}

// GetFirstBurnchainBlockHeight returns the FirstBurnchainBlockHeight field value
func (o *CoreNodePoxResponse) GetFirstBurnchainBlockHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FirstBurnchainBlockHeight
}

// GetFirstBurnchainBlockHeightOk returns a tuple with the FirstBurnchainBlockHeight field value
// and a boolean to check if the value has been set.
func (o *CoreNodePoxResponse) GetFirstBurnchainBlockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstBurnchainBlockHeight, true
}

// SetFirstBurnchainBlockHeight sets field value
func (o *CoreNodePoxResponse) SetFirstBurnchainBlockHeight(v int32) {
	o.FirstBurnchainBlockHeight = v
}

// GetMinAmountUstx returns the MinAmountUstx field value
func (o *CoreNodePoxResponse) GetMinAmountUstx() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinAmountUstx
}

// GetMinAmountUstxOk returns a tuple with the MinAmountUstx field value
// and a boolean to check if the value has been set.
func (o *CoreNodePoxResponse) GetMinAmountUstxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinAmountUstx, true
}

// SetMinAmountUstx sets field value
func (o *CoreNodePoxResponse) SetMinAmountUstx(v int32) {
	o.MinAmountUstx = v
}

// GetRegistrationWindowLength returns the RegistrationWindowLength field value
func (o *CoreNodePoxResponse) GetRegistrationWindowLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RegistrationWindowLength
}

// GetRegistrationWindowLengthOk returns a tuple with the RegistrationWindowLength field value
// and a boolean to check if the value has been set.
func (o *CoreNodePoxResponse) GetRegistrationWindowLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegistrationWindowLength, true
}

// SetRegistrationWindowLength sets field value
func (o *CoreNodePoxResponse) SetRegistrationWindowLength(v int32) {
	o.RegistrationWindowLength = v
}

// GetRejectionFraction returns the RejectionFraction field value
func (o *CoreNodePoxResponse) GetRejectionFraction() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RejectionFraction
}

// GetRejectionFractionOk returns a tuple with the RejectionFraction field value
// and a boolean to check if the value has been set.
func (o *CoreNodePoxResponse) GetRejectionFractionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RejectionFraction, true
}

// SetRejectionFraction sets field value
func (o *CoreNodePoxResponse) SetRejectionFraction(v int32) {
	o.RejectionFraction = v
}

// GetRewardCycleId returns the RewardCycleId field value
func (o *CoreNodePoxResponse) GetRewardCycleId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RewardCycleId
}

// GetRewardCycleIdOk returns a tuple with the RewardCycleId field value
// and a boolean to check if the value has been set.
func (o *CoreNodePoxResponse) GetRewardCycleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RewardCycleId, true
}

// SetRewardCycleId sets field value
func (o *CoreNodePoxResponse) SetRewardCycleId(v int32) {
	o.RewardCycleId = v
}

// GetRewardCycleLength returns the RewardCycleLength field value
func (o *CoreNodePoxResponse) GetRewardCycleLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RewardCycleLength
}

// GetRewardCycleLengthOk returns a tuple with the RewardCycleLength field value
// and a boolean to check if the value has been set.
func (o *CoreNodePoxResponse) GetRewardCycleLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RewardCycleLength, true
}

// SetRewardCycleLength sets field value
func (o *CoreNodePoxResponse) SetRewardCycleLength(v int32) {
	o.RewardCycleLength = v
}

// GetRejectionVotesLeftRequired returns the RejectionVotesLeftRequired field value
func (o *CoreNodePoxResponse) GetRejectionVotesLeftRequired() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RejectionVotesLeftRequired
}

// GetRejectionVotesLeftRequiredOk returns a tuple with the RejectionVotesLeftRequired field value
// and a boolean to check if the value has been set.
func (o *CoreNodePoxResponse) GetRejectionVotesLeftRequiredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RejectionVotesLeftRequired, true
}

// SetRejectionVotesLeftRequired sets field value
func (o *CoreNodePoxResponse) SetRejectionVotesLeftRequired(v int32) {
	o.RejectionVotesLeftRequired = v
}

// GetTotalLiquidSupplyUstx returns the TotalLiquidSupplyUstx field value
func (o *CoreNodePoxResponse) GetTotalLiquidSupplyUstx() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalLiquidSupplyUstx
}

// GetTotalLiquidSupplyUstxOk returns a tuple with the TotalLiquidSupplyUstx field value
// and a boolean to check if the value has been set.
func (o *CoreNodePoxResponse) GetTotalLiquidSupplyUstxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalLiquidSupplyUstx, true
}

// SetTotalLiquidSupplyUstx sets field value
func (o *CoreNodePoxResponse) SetTotalLiquidSupplyUstx(v int32) {
	o.TotalLiquidSupplyUstx = v
}

func (o CoreNodePoxResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CoreNodePoxResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contract_id"] = o.ContractId
	toSerialize["first_burnchain_block_height"] = o.FirstBurnchainBlockHeight
	toSerialize["min_amount_ustx"] = o.MinAmountUstx
	toSerialize["registration_window_length"] = o.RegistrationWindowLength
	toSerialize["rejection_fraction"] = o.RejectionFraction
	toSerialize["reward_cycle_id"] = o.RewardCycleId
	toSerialize["reward_cycle_length"] = o.RewardCycleLength
	toSerialize["rejection_votes_left_required"] = o.RejectionVotesLeftRequired
	toSerialize["total_liquid_supply_ustx"] = o.TotalLiquidSupplyUstx
	return toSerialize, nil
}

func (o *CoreNodePoxResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contract_id",
		"first_burnchain_block_height",
		"min_amount_ustx",
		"registration_window_length",
		"rejection_fraction",
		"reward_cycle_id",
		"reward_cycle_length",
		"rejection_votes_left_required",
		"total_liquid_supply_ustx",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCoreNodePoxResponse := _CoreNodePoxResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCoreNodePoxResponse)

	if err != nil {
		return err
	}

	*o = CoreNodePoxResponse(varCoreNodePoxResponse)

	return err
}

type NullableCoreNodePoxResponse struct {
	value *CoreNodePoxResponse
	isSet bool
}

func (v NullableCoreNodePoxResponse) Get() *CoreNodePoxResponse {
	return v.value
}

func (v *NullableCoreNodePoxResponse) Set(val *CoreNodePoxResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreNodePoxResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreNodePoxResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreNodePoxResponse(val *CoreNodePoxResponse) *NullableCoreNodePoxResponse {
	return &NullableCoreNodePoxResponse{value: val, isSet: true}
}

func (v NullableCoreNodePoxResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreNodePoxResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
