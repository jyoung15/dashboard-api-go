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

// checks if the GetNetworkSettings200ResponseFips type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSettings200ResponseFips{}

// GetNetworkSettings200ResponseFips A hash of FIPS options applied to the Network
type GetNetworkSettings200ResponseFips struct {
	// Enables / disables FIPS on the network.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewGetNetworkSettings200ResponseFips instantiates a new GetNetworkSettings200ResponseFips object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSettings200ResponseFips() *GetNetworkSettings200ResponseFips {
	this := GetNetworkSettings200ResponseFips{}
	return &this
}

// NewGetNetworkSettings200ResponseFipsWithDefaults instantiates a new GetNetworkSettings200ResponseFips object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSettings200ResponseFipsWithDefaults() *GetNetworkSettings200ResponseFips {
	this := GetNetworkSettings200ResponseFips{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkSettings200ResponseFips) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSettings200ResponseFips) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetNetworkSettings200ResponseFips) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkSettings200ResponseFips) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o GetNetworkSettings200ResponseFips) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSettings200ResponseFips) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableGetNetworkSettings200ResponseFips struct {
	value *GetNetworkSettings200ResponseFips
	isSet bool
}

func (v NullableGetNetworkSettings200ResponseFips) Get() *GetNetworkSettings200ResponseFips {
	return v.value
}

func (v *NullableGetNetworkSettings200ResponseFips) Set(val *GetNetworkSettings200ResponseFips) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSettings200ResponseFips) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSettings200ResponseFips) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSettings200ResponseFips(val *GetNetworkSettings200ResponseFips) *NullableGetNetworkSettings200ResponseFips {
	return &NullableGetNetworkSettings200ResponseFips{value: val, isSet: true}
}

func (v NullableGetNetworkSettings200ResponseFips) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSettings200ResponseFips) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


