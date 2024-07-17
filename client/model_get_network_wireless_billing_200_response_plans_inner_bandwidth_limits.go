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

// checks if the GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits{}

// GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits The uplink bandwidth settings for the pricing plan.
type GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits struct {
	// The maximum upload limit (integer, in Kbps).
	LimitUp *int32 `json:"limitUp,omitempty"`
	// The maximum download limit (integer, in Kbps).
	LimitDown *int32 `json:"limitDown,omitempty"`
}

// NewGetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits instantiates a new GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits() *GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits {
	this := GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits{}
	return &this
}

// NewGetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimitsWithDefaults instantiates a new GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimitsWithDefaults() *GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits {
	this := GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits{}
	return &this
}

// GetLimitUp returns the LimitUp field value if set, zero value otherwise.
func (o *GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits) GetLimitUp() int32 {
	if o == nil || IsNil(o.LimitUp) {
		var ret int32
		return ret
	}
	return *o.LimitUp
}

// GetLimitUpOk returns a tuple with the LimitUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits) GetLimitUpOk() (*int32, bool) {
	if o == nil || IsNil(o.LimitUp) {
		return nil, false
	}
	return o.LimitUp, true
}

// HasLimitUp returns a boolean if a field has been set.
func (o *GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits) HasLimitUp() bool {
	if o != nil && !IsNil(o.LimitUp) {
		return true
	}

	return false
}

// SetLimitUp gets a reference to the given int32 and assigns it to the LimitUp field.
func (o *GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits) SetLimitUp(v int32) {
	o.LimitUp = &v
}

// GetLimitDown returns the LimitDown field value if set, zero value otherwise.
func (o *GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits) GetLimitDown() int32 {
	if o == nil || IsNil(o.LimitDown) {
		var ret int32
		return ret
	}
	return *o.LimitDown
}

// GetLimitDownOk returns a tuple with the LimitDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits) GetLimitDownOk() (*int32, bool) {
	if o == nil || IsNil(o.LimitDown) {
		return nil, false
	}
	return o.LimitDown, true
}

// HasLimitDown returns a boolean if a field has been set.
func (o *GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits) HasLimitDown() bool {
	if o != nil && !IsNil(o.LimitDown) {
		return true
	}

	return false
}

// SetLimitDown gets a reference to the given int32 and assigns it to the LimitDown field.
func (o *GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits) SetLimitDown(v int32) {
	o.LimitDown = &v
}

func (o GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LimitUp) {
		toSerialize["limitUp"] = o.LimitUp
	}
	if !IsNil(o.LimitDown) {
		toSerialize["limitDown"] = o.LimitDown
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits struct {
	value *GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits
	isSet bool
}

func (v NullableGetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits) Get() *GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits {
	return v.value
}

func (v *NullableGetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits) Set(val *GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits(val *GetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits) *NullableGetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits {
	return &NullableGetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessBilling200ResponsePlansInnerBandwidthLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


