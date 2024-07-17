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

// checks if the GetNetworkApplianceVlans200ResponseInnerIpv6 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceVlans200ResponseInnerIpv6{}

// GetNetworkApplianceVlans200ResponseInnerIpv6 IPv6 configuration on the VLAN
type GetNetworkApplianceVlans200ResponseInnerIpv6 struct {
	// Enable IPv6 on VLAN
	Enabled *bool `json:"enabled,omitempty"`
	// Prefix assignments on the VLAN
	PrefixAssignments []GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner `json:"prefixAssignments,omitempty"`
}

// NewGetNetworkApplianceVlans200ResponseInnerIpv6 instantiates a new GetNetworkApplianceVlans200ResponseInnerIpv6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceVlans200ResponseInnerIpv6() *GetNetworkApplianceVlans200ResponseInnerIpv6 {
	this := GetNetworkApplianceVlans200ResponseInnerIpv6{}
	return &this
}

// NewGetNetworkApplianceVlans200ResponseInnerIpv6WithDefaults instantiates a new GetNetworkApplianceVlans200ResponseInnerIpv6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceVlans200ResponseInnerIpv6WithDefaults() *GetNetworkApplianceVlans200ResponseInnerIpv6 {
	this := GetNetworkApplianceVlans200ResponseInnerIpv6{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkApplianceVlans200ResponseInnerIpv6) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceVlans200ResponseInnerIpv6) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetNetworkApplianceVlans200ResponseInnerIpv6) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkApplianceVlans200ResponseInnerIpv6) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPrefixAssignments returns the PrefixAssignments field value if set, zero value otherwise.
func (o *GetNetworkApplianceVlans200ResponseInnerIpv6) GetPrefixAssignments() []GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner {
	if o == nil || IsNil(o.PrefixAssignments) {
		var ret []GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner
		return ret
	}
	return o.PrefixAssignments
}

// GetPrefixAssignmentsOk returns a tuple with the PrefixAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceVlans200ResponseInnerIpv6) GetPrefixAssignmentsOk() ([]GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner, bool) {
	if o == nil || IsNil(o.PrefixAssignments) {
		return nil, false
	}
	return o.PrefixAssignments, true
}

// HasPrefixAssignments returns a boolean if a field has been set.
func (o *GetNetworkApplianceVlans200ResponseInnerIpv6) HasPrefixAssignments() bool {
	if o != nil && !IsNil(o.PrefixAssignments) {
		return true
	}

	return false
}

// SetPrefixAssignments gets a reference to the given []GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner and assigns it to the PrefixAssignments field.
func (o *GetNetworkApplianceVlans200ResponseInnerIpv6) SetPrefixAssignments(v []GetNetworkApplianceVlans200ResponseInnerIpv6PrefixAssignmentsInner) {
	o.PrefixAssignments = v
}

func (o GetNetworkApplianceVlans200ResponseInnerIpv6) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceVlans200ResponseInnerIpv6) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.PrefixAssignments) {
		toSerialize["prefixAssignments"] = o.PrefixAssignments
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceVlans200ResponseInnerIpv6 struct {
	value *GetNetworkApplianceVlans200ResponseInnerIpv6
	isSet bool
}

func (v NullableGetNetworkApplianceVlans200ResponseInnerIpv6) Get() *GetNetworkApplianceVlans200ResponseInnerIpv6 {
	return v.value
}

func (v *NullableGetNetworkApplianceVlans200ResponseInnerIpv6) Set(val *GetNetworkApplianceVlans200ResponseInnerIpv6) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceVlans200ResponseInnerIpv6) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceVlans200ResponseInnerIpv6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceVlans200ResponseInnerIpv6(val *GetNetworkApplianceVlans200ResponseInnerIpv6) *NullableGetNetworkApplianceVlans200ResponseInnerIpv6 {
	return &NullableGetNetworkApplianceVlans200ResponseInnerIpv6{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceVlans200ResponseInnerIpv6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceVlans200ResponseInnerIpv6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


