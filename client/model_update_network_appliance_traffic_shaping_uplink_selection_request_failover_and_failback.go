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

// checks if the UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback{}

// UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback WAN failover and failback behavior
type UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback struct {
	Immediate *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate `json:"immediate,omitempty"`
}

// NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback instantiates a new UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback() *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback {
	this := UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback{}
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackWithDefaults() *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback {
	this := UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback{}
	return &this
}

// GetImmediate returns the Immediate field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback) GetImmediate() UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate {
	if o == nil || IsNil(o.Immediate) {
		var ret UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate
		return ret
	}
	return *o.Immediate
}

// GetImmediateOk returns a tuple with the Immediate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback) GetImmediateOk() (*UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate, bool) {
	if o == nil || IsNil(o.Immediate) {
		return nil, false
	}
	return o.Immediate, true
}

// HasImmediate returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback) HasImmediate() bool {
	if o != nil && !IsNil(o.Immediate) {
		return true
	}

	return false
}

// SetImmediate gets a reference to the given UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate and assigns it to the Immediate field.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback) SetImmediate(v UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailbackImmediate) {
	o.Immediate = &v
}

func (o UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Immediate) {
		toSerialize["immediate"] = o.Immediate
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback struct {
	value *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback) Get() *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback) Set(val *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback(val *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback) *NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback {
	return &NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestFailoverAndFailback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


