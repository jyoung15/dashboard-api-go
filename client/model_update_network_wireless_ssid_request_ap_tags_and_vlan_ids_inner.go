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

// checks if the UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner{}

// UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner struct for UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner
type UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner struct {
	// Array of AP tags
	Tags []string `json:"tags,omitempty"`
	// Numerical identifier that is assigned to the VLAN
	VlanId *int32 `json:"vlanId,omitempty"`
}

// NewUpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner instantiates a new UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner() *UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner {
	this := UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner{}
	return &this
}

// NewUpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInnerWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInnerWithDefaults() *UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner {
	this := UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner) SetTags(v []string) {
	o.Tags = v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner) GetVlanId() int32 {
	if o == nil || IsNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner) GetVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner) SetVlanId(v int32) {
	o.VlanId = &v
}

func (o UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner struct {
	value *UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner) Get() *UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner) Set(val *UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner(val *UpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner) *NullableUpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner {
	return &NullableUpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestApTagsAndVlanIdsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


