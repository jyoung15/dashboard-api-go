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

// checks if the UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection{}

// UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection Dynamic ARP Inspection settings
type UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection struct {
	// Enable or disable Dynamic ARP Inspection on the network. Default value is false.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUpdateNetworkSwitchDhcpServerPolicyRequestArpInspection instantiates a new UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchDhcpServerPolicyRequestArpInspection() *UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection {
	this := UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection{}
	return &this
}

// NewUpdateNetworkSwitchDhcpServerPolicyRequestArpInspectionWithDefaults instantiates a new UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchDhcpServerPolicyRequestArpInspectionWithDefaults() *UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection {
	this := UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchDhcpServerPolicyRequestArpInspection struct {
	value *UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection
	isSet bool
}

func (v NullableUpdateNetworkSwitchDhcpServerPolicyRequestArpInspection) Get() *UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection {
	return v.value
}

func (v *NullableUpdateNetworkSwitchDhcpServerPolicyRequestArpInspection) Set(val *UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchDhcpServerPolicyRequestArpInspection) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchDhcpServerPolicyRequestArpInspection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchDhcpServerPolicyRequestArpInspection(val *UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection) *NullableUpdateNetworkSwitchDhcpServerPolicyRequestArpInspection {
	return &NullableUpdateNetworkSwitchDhcpServerPolicyRequestArpInspection{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchDhcpServerPolicyRequestArpInspection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchDhcpServerPolicyRequestArpInspection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


