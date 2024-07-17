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

// checks if the GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache{}

// GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache Object for RADIUS Cache Settings
type GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache struct {
	// Enable to cache authorization and authentication responses on the RADIUS server
	Enabled *bool `json:"enabled,omitempty"`
	// If RADIUS caching is enabled, this value dictates how long the cache will remain in the RADIUS server, in hours, to allow network access without authentication
	Timeout *int32 `json:"timeout,omitempty"`
}

// NewGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache instantiates a new GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache() *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache {
	this := GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache{}
	return &this
}

// NewGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCacheWithDefaults instantiates a new GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCacheWithDefaults() *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache {
	this := GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache) GetTimeout() int32 {
	if o == nil || IsNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache) GetTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache) SetTimeout(v int32) {
	o.Timeout = &v
}

func (o GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache struct {
	value *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache
	isSet bool
}

func (v NullableGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache) Get() *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache {
	return v.value
}

func (v *NullableGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache) Set(val *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache(val *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache) *NullableGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache {
	return &NullableGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCache) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

