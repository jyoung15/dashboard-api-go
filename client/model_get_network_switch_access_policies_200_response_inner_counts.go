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

// checks if the GetNetworkSwitchAccessPolicies200ResponseInnerCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchAccessPolicies200ResponseInnerCounts{}

// GetNetworkSwitchAccessPolicies200ResponseInnerCounts Counts associated with the access policy
type GetNetworkSwitchAccessPolicies200ResponseInnerCounts struct {
	Ports *GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts `json:"ports,omitempty"`
}

// NewGetNetworkSwitchAccessPolicies200ResponseInnerCounts instantiates a new GetNetworkSwitchAccessPolicies200ResponseInnerCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchAccessPolicies200ResponseInnerCounts() *GetNetworkSwitchAccessPolicies200ResponseInnerCounts {
	this := GetNetworkSwitchAccessPolicies200ResponseInnerCounts{}
	return &this
}

// NewGetNetworkSwitchAccessPolicies200ResponseInnerCountsWithDefaults instantiates a new GetNetworkSwitchAccessPolicies200ResponseInnerCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchAccessPolicies200ResponseInnerCountsWithDefaults() *GetNetworkSwitchAccessPolicies200ResponseInnerCounts {
	this := GetNetworkSwitchAccessPolicies200ResponseInnerCounts{}
	return &this
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerCounts) GetPorts() GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts {
	if o == nil || IsNil(o.Ports) {
		var ret GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts
		return ret
	}
	return *o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerCounts) GetPortsOk() (*GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerCounts) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts and assigns it to the Ports field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerCounts) SetPorts(v GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts) {
	o.Ports = &v
}

func (o GetNetworkSwitchAccessPolicies200ResponseInnerCounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchAccessPolicies200ResponseInnerCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchAccessPolicies200ResponseInnerCounts struct {
	value *GetNetworkSwitchAccessPolicies200ResponseInnerCounts
	isSet bool
}

func (v NullableGetNetworkSwitchAccessPolicies200ResponseInnerCounts) Get() *GetNetworkSwitchAccessPolicies200ResponseInnerCounts {
	return v.value
}

func (v *NullableGetNetworkSwitchAccessPolicies200ResponseInnerCounts) Set(val *GetNetworkSwitchAccessPolicies200ResponseInnerCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchAccessPolicies200ResponseInnerCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchAccessPolicies200ResponseInnerCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchAccessPolicies200ResponseInnerCounts(val *GetNetworkSwitchAccessPolicies200ResponseInnerCounts) *NullableGetNetworkSwitchAccessPolicies200ResponseInnerCounts {
	return &NullableGetNetworkSwitchAccessPolicies200ResponseInnerCounts{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchAccessPolicies200ResponseInnerCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchAccessPolicies200ResponseInnerCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


