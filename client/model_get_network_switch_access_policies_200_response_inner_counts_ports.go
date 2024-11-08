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

// checks if the GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts{}

// GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts Counts associated with ports
type GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts struct {
	// Number of ports in the network with this policy. For template networks, this is the number of template ports (not child ports) with this policy.
	WithThisPolicy *int32 `json:"withThisPolicy,omitempty"`
}

// NewGetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts instantiates a new GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts() *GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts {
	this := GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts{}
	return &this
}

// NewGetNetworkSwitchAccessPolicies200ResponseInnerCountsPortsWithDefaults instantiates a new GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchAccessPolicies200ResponseInnerCountsPortsWithDefaults() *GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts {
	this := GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts{}
	return &this
}

// GetWithThisPolicy returns the WithThisPolicy field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts) GetWithThisPolicy() int32 {
	if o == nil || IsNil(o.WithThisPolicy) {
		var ret int32
		return ret
	}
	return *o.WithThisPolicy
}

// GetWithThisPolicyOk returns a tuple with the WithThisPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts) GetWithThisPolicyOk() (*int32, bool) {
	if o == nil || IsNil(o.WithThisPolicy) {
		return nil, false
	}
	return o.WithThisPolicy, true
}

// HasWithThisPolicy returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts) HasWithThisPolicy() bool {
	if o != nil && !IsNil(o.WithThisPolicy) {
		return true
	}

	return false
}

// SetWithThisPolicy gets a reference to the given int32 and assigns it to the WithThisPolicy field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts) SetWithThisPolicy(v int32) {
	o.WithThisPolicy = &v
}

func (o GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WithThisPolicy) {
		toSerialize["withThisPolicy"] = o.WithThisPolicy
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts struct {
	value *GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts
	isSet bool
}

func (v NullableGetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts) Get() *GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts {
	return v.value
}

func (v *NullableGetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts) Set(val *GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts(val *GetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts) *NullableGetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts {
	return &NullableGetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchAccessPolicies200ResponseInnerCountsPorts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


