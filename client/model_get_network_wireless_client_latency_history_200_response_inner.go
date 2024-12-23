/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkWirelessClientLatencyHistory200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessClientLatencyHistory200ResponseInner{}

// GetNetworkWirelessClientLatencyHistory200ResponseInner struct for GetNetworkWirelessClientLatencyHistory200ResponseInner
type GetNetworkWirelessClientLatencyHistory200ResponseInner struct {
	// The latency history bucket start time in seconds
	T0 *int32 `json:"t0,omitempty"`
	// The latency history bucket end time in seconds
	T1 *int32 `json:"t1,omitempty"`
	LatencyBinsByCategory *GetNetworkWirelessClientLatencyHistory200ResponseInnerLatencyBinsByCategory `json:"latencyBinsByCategory,omitempty"`
}

// NewGetNetworkWirelessClientLatencyHistory200ResponseInner instantiates a new GetNetworkWirelessClientLatencyHistory200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessClientLatencyHistory200ResponseInner() *GetNetworkWirelessClientLatencyHistory200ResponseInner {
	this := GetNetworkWirelessClientLatencyHistory200ResponseInner{}
	return &this
}

// NewGetNetworkWirelessClientLatencyHistory200ResponseInnerWithDefaults instantiates a new GetNetworkWirelessClientLatencyHistory200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessClientLatencyHistory200ResponseInnerWithDefaults() *GetNetworkWirelessClientLatencyHistory200ResponseInner {
	this := GetNetworkWirelessClientLatencyHistory200ResponseInner{}
	return &this
}

// GetT0 returns the T0 field value if set, zero value otherwise.
func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) GetT0() int32 {
	if o == nil || IsNil(o.T0) {
		var ret int32
		return ret
	}
	return *o.T0
}

// GetT0Ok returns a tuple with the T0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) GetT0Ok() (*int32, bool) {
	if o == nil || IsNil(o.T0) {
		return nil, false
	}
	return o.T0, true
}

// HasT0 returns a boolean if a field has been set.
func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) HasT0() bool {
	if o != nil && !IsNil(o.T0) {
		return true
	}

	return false
}

// SetT0 gets a reference to the given int32 and assigns it to the T0 field.
func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) SetT0(v int32) {
	o.T0 = &v
}

// GetT1 returns the T1 field value if set, zero value otherwise.
func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) GetT1() int32 {
	if o == nil || IsNil(o.T1) {
		var ret int32
		return ret
	}
	return *o.T1
}

// GetT1Ok returns a tuple with the T1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) GetT1Ok() (*int32, bool) {
	if o == nil || IsNil(o.T1) {
		return nil, false
	}
	return o.T1, true
}

// HasT1 returns a boolean if a field has been set.
func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) HasT1() bool {
	if o != nil && !IsNil(o.T1) {
		return true
	}

	return false
}

// SetT1 gets a reference to the given int32 and assigns it to the T1 field.
func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) SetT1(v int32) {
	o.T1 = &v
}

// GetLatencyBinsByCategory returns the LatencyBinsByCategory field value if set, zero value otherwise.
func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) GetLatencyBinsByCategory() GetNetworkWirelessClientLatencyHistory200ResponseInnerLatencyBinsByCategory {
	if o == nil || IsNil(o.LatencyBinsByCategory) {
		var ret GetNetworkWirelessClientLatencyHistory200ResponseInnerLatencyBinsByCategory
		return ret
	}
	return *o.LatencyBinsByCategory
}

// GetLatencyBinsByCategoryOk returns a tuple with the LatencyBinsByCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) GetLatencyBinsByCategoryOk() (*GetNetworkWirelessClientLatencyHistory200ResponseInnerLatencyBinsByCategory, bool) {
	if o == nil || IsNil(o.LatencyBinsByCategory) {
		return nil, false
	}
	return o.LatencyBinsByCategory, true
}

// HasLatencyBinsByCategory returns a boolean if a field has been set.
func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) HasLatencyBinsByCategory() bool {
	if o != nil && !IsNil(o.LatencyBinsByCategory) {
		return true
	}

	return false
}

// SetLatencyBinsByCategory gets a reference to the given GetNetworkWirelessClientLatencyHistory200ResponseInnerLatencyBinsByCategory and assigns it to the LatencyBinsByCategory field.
func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) SetLatencyBinsByCategory(v GetNetworkWirelessClientLatencyHistory200ResponseInnerLatencyBinsByCategory) {
	o.LatencyBinsByCategory = &v
}

func (o GetNetworkWirelessClientLatencyHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessClientLatencyHistory200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.T0) {
		toSerialize["t0"] = o.T0
	}
	if !IsNil(o.T1) {
		toSerialize["t1"] = o.T1
	}
	if !IsNil(o.LatencyBinsByCategory) {
		toSerialize["latencyBinsByCategory"] = o.LatencyBinsByCategory
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessClientLatencyHistory200ResponseInner struct {
	value *GetNetworkWirelessClientLatencyHistory200ResponseInner
	isSet bool
}

func (v NullableGetNetworkWirelessClientLatencyHistory200ResponseInner) Get() *GetNetworkWirelessClientLatencyHistory200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkWirelessClientLatencyHistory200ResponseInner) Set(val *GetNetworkWirelessClientLatencyHistory200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessClientLatencyHistory200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessClientLatencyHistory200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessClientLatencyHistory200ResponseInner(val *GetNetworkWirelessClientLatencyHistory200ResponseInner) *NullableGetNetworkWirelessClientLatencyHistory200ResponseInner {
	return &NullableGetNetworkWirelessClientLatencyHistory200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessClientLatencyHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessClientLatencyHistory200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


