/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateNetworkApplianceSingleLanRequestIpv6 IPv6 configuration on the VLAN
type UpdateNetworkApplianceSingleLanRequestIpv6 struct {
	// Enable IPv6 on VLAN.
	Enabled *bool `json:"enabled,omitempty"`
	// Prefix assignments on the VLAN
	PrefixAssignments []UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner `json:"prefixAssignments,omitempty"`
}

// NewUpdateNetworkApplianceSingleLanRequestIpv6 instantiates a new UpdateNetworkApplianceSingleLanRequestIpv6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceSingleLanRequestIpv6() *UpdateNetworkApplianceSingleLanRequestIpv6 {
	this := UpdateNetworkApplianceSingleLanRequestIpv6{}
	return &this
}

// NewUpdateNetworkApplianceSingleLanRequestIpv6WithDefaults instantiates a new UpdateNetworkApplianceSingleLanRequestIpv6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceSingleLanRequestIpv6WithDefaults() *UpdateNetworkApplianceSingleLanRequestIpv6 {
	this := UpdateNetworkApplianceSingleLanRequestIpv6{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPrefixAssignments returns the PrefixAssignments field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6) GetPrefixAssignments() []UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner {
	if o == nil || isNil(o.PrefixAssignments) {
		var ret []UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner
		return ret
	}
	return o.PrefixAssignments
}

// GetPrefixAssignmentsOk returns a tuple with the PrefixAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6) GetPrefixAssignmentsOk() ([]UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner, bool) {
	if o == nil || isNil(o.PrefixAssignments) {
    return nil, false
	}
	return o.PrefixAssignments, true
}

// HasPrefixAssignments returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6) HasPrefixAssignments() bool {
	if o != nil && !isNil(o.PrefixAssignments) {
		return true
	}

	return false
}

// SetPrefixAssignments gets a reference to the given []UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner and assigns it to the PrefixAssignments field.
func (o *UpdateNetworkApplianceSingleLanRequestIpv6) SetPrefixAssignments(v []UpdateNetworkApplianceSingleLanRequestIpv6PrefixAssignmentsInner) {
	o.PrefixAssignments = v
}

func (o UpdateNetworkApplianceSingleLanRequestIpv6) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.PrefixAssignments) {
		toSerialize["prefixAssignments"] = o.PrefixAssignments
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceSingleLanRequestIpv6 struct {
	value *UpdateNetworkApplianceSingleLanRequestIpv6
	isSet bool
}

func (v NullableUpdateNetworkApplianceSingleLanRequestIpv6) Get() *UpdateNetworkApplianceSingleLanRequestIpv6 {
	return v.value
}

func (v *NullableUpdateNetworkApplianceSingleLanRequestIpv6) Set(val *UpdateNetworkApplianceSingleLanRequestIpv6) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceSingleLanRequestIpv6) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceSingleLanRequestIpv6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceSingleLanRequestIpv6(val *UpdateNetworkApplianceSingleLanRequestIpv6) *NullableUpdateNetworkApplianceSingleLanRequestIpv6 {
	return &NullableUpdateNetworkApplianceSingleLanRequestIpv6{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceSingleLanRequestIpv6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceSingleLanRequestIpv6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

