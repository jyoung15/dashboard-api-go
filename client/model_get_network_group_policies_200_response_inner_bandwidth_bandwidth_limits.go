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

// checks if the GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits{}

// GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits The bandwidth limits object, specifying upload and download speed for clients bound to the group policy. These are only enforced if 'settings' is set to 'custom'.
type GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits struct {
	// The maximum upload limit (integer, in Kbps). null indicates no limit
	LimitUp *int32 `json:"limitUp,omitempty"`
	// The maximum download limit (integer, in Kbps). null indicates no limit
	LimitDown *int32 `json:"limitDown,omitempty"`
}

// NewGetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits instantiates a new GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits() *GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits {
	this := GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits{}
	return &this
}

// NewGetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimitsWithDefaults instantiates a new GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimitsWithDefaults() *GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits {
	this := GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits{}
	return &this
}

// GetLimitUp returns the LimitUp field value if set, zero value otherwise.
func (o *GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits) GetLimitUp() int32 {
	if o == nil || IsNil(o.LimitUp) {
		var ret int32
		return ret
	}
	return *o.LimitUp
}

// GetLimitUpOk returns a tuple with the LimitUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits) GetLimitUpOk() (*int32, bool) {
	if o == nil || IsNil(o.LimitUp) {
		return nil, false
	}
	return o.LimitUp, true
}

// HasLimitUp returns a boolean if a field has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits) HasLimitUp() bool {
	if o != nil && !IsNil(o.LimitUp) {
		return true
	}

	return false
}

// SetLimitUp gets a reference to the given int32 and assigns it to the LimitUp field.
func (o *GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits) SetLimitUp(v int32) {
	o.LimitUp = &v
}

// GetLimitDown returns the LimitDown field value if set, zero value otherwise.
func (o *GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits) GetLimitDown() int32 {
	if o == nil || IsNil(o.LimitDown) {
		var ret int32
		return ret
	}
	return *o.LimitDown
}

// GetLimitDownOk returns a tuple with the LimitDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits) GetLimitDownOk() (*int32, bool) {
	if o == nil || IsNil(o.LimitDown) {
		return nil, false
	}
	return o.LimitDown, true
}

// HasLimitDown returns a boolean if a field has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits) HasLimitDown() bool {
	if o != nil && !IsNil(o.LimitDown) {
		return true
	}

	return false
}

// SetLimitDown gets a reference to the given int32 and assigns it to the LimitDown field.
func (o *GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits) SetLimitDown(v int32) {
	o.LimitDown = &v
}

func (o GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LimitUp) {
		toSerialize["limitUp"] = o.LimitUp
	}
	if !IsNil(o.LimitDown) {
		toSerialize["limitDown"] = o.LimitDown
	}
	return toSerialize, nil
}

type NullableGetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits struct {
	value *GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits
	isSet bool
}

func (v NullableGetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits) Get() *GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits {
	return v.value
}

func (v *NullableGetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits) Set(val *GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits(val *GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits) *NullableGetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits {
	return &NullableGetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits{value: val, isSet: true}
}

func (v NullableGetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


