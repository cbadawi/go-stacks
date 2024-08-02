package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cbadawi/go-stacks/stacks/utils"
)

// checks if the BnsGetNameInfoResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BnsGetNameInfoResponse{}

// BnsGetNameInfoResponse Get name details
type BnsGetNameInfoResponse struct {
	Address      string  `json:"address"`
	Blockchain   string  `json:"blockchain" validate:"regexp=^stacks$"`
	ExpireBlock  *int32  `json:"expire_block,omitempty"`
	GracePeriod  *int32  `json:"grace_period,omitempty"`
	LastTxid     string  `json:"last_txid"`
	Resolver     *string `json:"resolver,omitempty"`
	Status       string  `json:"status"`
	Zonefile     string  `json:"zonefile"`
	ZonefileHash string  `json:"zonefile_hash"`
}

type _BnsGetNameInfoResponse BnsGetNameInfoResponse

// NewBnsGetNameInfoResponse instantiates a new BnsGetNameInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBnsGetNameInfoResponse(address string, blockchain string, lastTxid string, status string, zonefile string, zonefileHash string) *BnsGetNameInfoResponse {
	this := BnsGetNameInfoResponse{}
	this.Address = address
	this.Blockchain = blockchain
	this.LastTxid = lastTxid
	this.Status = status
	this.Zonefile = zonefile
	this.ZonefileHash = zonefileHash
	return &this
}

// NewBnsGetNameInfoResponseWithDefaults instantiates a new BnsGetNameInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBnsGetNameInfoResponseWithDefaults() *BnsGetNameInfoResponse {
	this := BnsGetNameInfoResponse{}
	return &this
}

// GetAddress returns the Address field value
func (o *BnsGetNameInfoResponse) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *BnsGetNameInfoResponse) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *BnsGetNameInfoResponse) SetAddress(v string) {
	o.Address = v
}

// GetBlockchain returns the Blockchain field value
func (o *BnsGetNameInfoResponse) GetBlockchain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Blockchain
}

// GetBlockchainOk returns a tuple with the Blockchain field value
// and a boolean to check if the value has been set.
func (o *BnsGetNameInfoResponse) GetBlockchainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blockchain, true
}

// SetBlockchain sets field value
func (o *BnsGetNameInfoResponse) SetBlockchain(v string) {
	o.Blockchain = v
}

// GetExpireBlock returns the ExpireBlock field value if set, zero value otherwise.
func (o *BnsGetNameInfoResponse) GetExpireBlock() int32 {
	if o == nil || utils.IsNil(o.ExpireBlock) {
		var ret int32
		return ret
	}
	return *o.ExpireBlock
}

// GetExpireBlockOk returns a tuple with the ExpireBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BnsGetNameInfoResponse) GetExpireBlockOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.ExpireBlock) {
		return nil, false
	}
	return o.ExpireBlock, true
}

// HasExpireBlock returns a boolean if a field has been set.
func (o *BnsGetNameInfoResponse) HasExpireBlock() bool {
	if o != nil && !utils.IsNil(o.ExpireBlock) {
		return true
	}

	return false
}

// SetExpireBlock gets a reference to the given int32 and assigns it to the ExpireBlock field.
func (o *BnsGetNameInfoResponse) SetExpireBlock(v int32) {
	o.ExpireBlock = &v
}

// GetGracePeriod returns the GracePeriod field value if set, zero value otherwise.
func (o *BnsGetNameInfoResponse) GetGracePeriod() int32 {
	if o == nil || utils.IsNil(o.GracePeriod) {
		var ret int32
		return ret
	}
	return *o.GracePeriod
}

// GetGracePeriodOk returns a tuple with the GracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BnsGetNameInfoResponse) GetGracePeriodOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.GracePeriod) {
		return nil, false
	}
	return o.GracePeriod, true
}

// HasGracePeriod returns a boolean if a field has been set.
func (o *BnsGetNameInfoResponse) HasGracePeriod() bool {
	if o != nil && !utils.IsNil(o.GracePeriod) {
		return true
	}

	return false
}

// SetGracePeriod gets a reference to the given int32 and assigns it to the GracePeriod field.
func (o *BnsGetNameInfoResponse) SetGracePeriod(v int32) {
	o.GracePeriod = &v
}

// GetLastTxid returns the LastTxid field value
func (o *BnsGetNameInfoResponse) GetLastTxid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastTxid
}

// GetLastTxidOk returns a tuple with the LastTxid field value
// and a boolean to check if the value has been set.
func (o *BnsGetNameInfoResponse) GetLastTxidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastTxid, true
}

// SetLastTxid sets field value
func (o *BnsGetNameInfoResponse) SetLastTxid(v string) {
	o.LastTxid = v
}

// GetResolver returns the Resolver field value if set, zero value otherwise.
func (o *BnsGetNameInfoResponse) GetResolver() string {
	if o == nil || utils.IsNil(o.Resolver) {
		var ret string
		return ret
	}
	return *o.Resolver
}

// GetResolverOk returns a tuple with the Resolver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BnsGetNameInfoResponse) GetResolverOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Resolver) {
		return nil, false
	}
	return o.Resolver, true
}

// HasResolver returns a boolean if a field has been set.
func (o *BnsGetNameInfoResponse) HasResolver() bool {
	if o != nil && !utils.IsNil(o.Resolver) {
		return true
	}

	return false
}

// SetResolver gets a reference to the given string and assigns it to the Resolver field.
func (o *BnsGetNameInfoResponse) SetResolver(v string) {
	o.Resolver = &v
}

// GetStatus returns the Status field value
func (o *BnsGetNameInfoResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BnsGetNameInfoResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BnsGetNameInfoResponse) SetStatus(v string) {
	o.Status = v
}

// GetZonefile returns the Zonefile field value
func (o *BnsGetNameInfoResponse) GetZonefile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Zonefile
}

// GetZonefileOk returns a tuple with the Zonefile field value
// and a boolean to check if the value has been set.
func (o *BnsGetNameInfoResponse) GetZonefileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Zonefile, true
}

// SetZonefile sets field value
func (o *BnsGetNameInfoResponse) SetZonefile(v string) {
	o.Zonefile = v
}

// GetZonefileHash returns the ZonefileHash field value
func (o *BnsGetNameInfoResponse) GetZonefileHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ZonefileHash
}

// GetZonefileHashOk returns a tuple with the ZonefileHash field value
// and a boolean to check if the value has been set.
func (o *BnsGetNameInfoResponse) GetZonefileHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZonefileHash, true
}

// SetZonefileHash sets field value
func (o *BnsGetNameInfoResponse) SetZonefileHash(v string) {
	o.ZonefileHash = v
}

func (o BnsGetNameInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BnsGetNameInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["blockchain"] = o.Blockchain
	if !utils.IsNil(o.ExpireBlock) {
		toSerialize["expire_block"] = o.ExpireBlock
	}
	if !utils.IsNil(o.GracePeriod) {
		toSerialize["grace_period"] = o.GracePeriod
	}
	toSerialize["last_txid"] = o.LastTxid
	if !utils.IsNil(o.Resolver) {
		toSerialize["resolver"] = o.Resolver
	}
	toSerialize["status"] = o.Status
	toSerialize["zonefile"] = o.Zonefile
	toSerialize["zonefile_hash"] = o.ZonefileHash
	return toSerialize, nil
}

func (o *BnsGetNameInfoResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"blockchain",
		"last_txid",
		"status",
		"zonefile",
		"zonefile_hash",
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

	varBnsGetNameInfoResponse := _BnsGetNameInfoResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBnsGetNameInfoResponse)

	if err != nil {
		return err
	}

	*o = BnsGetNameInfoResponse(varBnsGetNameInfoResponse)

	return err
}

type NullableBnsGetNameInfoResponse struct {
	value *BnsGetNameInfoResponse
	isSet bool
}

func (v NullableBnsGetNameInfoResponse) Get() *BnsGetNameInfoResponse {
	return v.value
}

func (v *NullableBnsGetNameInfoResponse) Set(val *BnsGetNameInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBnsGetNameInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBnsGetNameInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBnsGetNameInfoResponse(val *BnsGetNameInfoResponse) *NullableBnsGetNameInfoResponse {
	return &NullableBnsGetNameInfoResponse{value: val, isSet: true}
}

func (v NullableBnsGetNameInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBnsGetNameInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
