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

// checks if the UpdateNetworkSwitchStpRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchStpRequest{}

// UpdateNetworkSwitchStpRequest struct for UpdateNetworkSwitchStpRequest
type UpdateNetworkSwitchStpRequest struct {
	// The spanning tree protocol status in network
	RstpEnabled *bool `json:"rstpEnabled,omitempty"`
	// STP bridge priority for switches/stacks or switch templates. An empty array will clear the STP bridge priority settings.
	StpBridgePriority []UpdateNetworkSwitchStpRequestStpBridgePriorityInner `json:"stpBridgePriority,omitempty"`
}

// NewUpdateNetworkSwitchStpRequest instantiates a new UpdateNetworkSwitchStpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchStpRequest() *UpdateNetworkSwitchStpRequest {
	this := UpdateNetworkSwitchStpRequest{}
	return &this
}

// NewUpdateNetworkSwitchStpRequestWithDefaults instantiates a new UpdateNetworkSwitchStpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchStpRequestWithDefaults() *UpdateNetworkSwitchStpRequest {
	this := UpdateNetworkSwitchStpRequest{}
	return &this
}

// GetRstpEnabled returns the RstpEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStpRequest) GetRstpEnabled() bool {
	if o == nil || IsNil(o.RstpEnabled) {
		var ret bool
		return ret
	}
	return *o.RstpEnabled
}

// GetRstpEnabledOk returns a tuple with the RstpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStpRequest) GetRstpEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.RstpEnabled) {
		return nil, false
	}
	return o.RstpEnabled, true
}

// HasRstpEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStpRequest) HasRstpEnabled() bool {
	if o != nil && !IsNil(o.RstpEnabled) {
		return true
	}

	return false
}

// SetRstpEnabled gets a reference to the given bool and assigns it to the RstpEnabled field.
func (o *UpdateNetworkSwitchStpRequest) SetRstpEnabled(v bool) {
	o.RstpEnabled = &v
}

// GetStpBridgePriority returns the StpBridgePriority field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStpRequest) GetStpBridgePriority() []UpdateNetworkSwitchStpRequestStpBridgePriorityInner {
	if o == nil || IsNil(o.StpBridgePriority) {
		var ret []UpdateNetworkSwitchStpRequestStpBridgePriorityInner
		return ret
	}
	return o.StpBridgePriority
}

// GetStpBridgePriorityOk returns a tuple with the StpBridgePriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStpRequest) GetStpBridgePriorityOk() ([]UpdateNetworkSwitchStpRequestStpBridgePriorityInner, bool) {
	if o == nil || IsNil(o.StpBridgePriority) {
		return nil, false
	}
	return o.StpBridgePriority, true
}

// HasStpBridgePriority returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStpRequest) HasStpBridgePriority() bool {
	if o != nil && !IsNil(o.StpBridgePriority) {
		return true
	}

	return false
}

// SetStpBridgePriority gets a reference to the given []UpdateNetworkSwitchStpRequestStpBridgePriorityInner and assigns it to the StpBridgePriority field.
func (o *UpdateNetworkSwitchStpRequest) SetStpBridgePriority(v []UpdateNetworkSwitchStpRequestStpBridgePriorityInner) {
	o.StpBridgePriority = v
}

func (o UpdateNetworkSwitchStpRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchStpRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RstpEnabled) {
		toSerialize["rstpEnabled"] = o.RstpEnabled
	}
	if !IsNil(o.StpBridgePriority) {
		toSerialize["stpBridgePriority"] = o.StpBridgePriority
	}
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchStpRequest struct {
	value *UpdateNetworkSwitchStpRequest
	isSet bool
}

func (v NullableUpdateNetworkSwitchStpRequest) Get() *UpdateNetworkSwitchStpRequest {
	return v.value
}

func (v *NullableUpdateNetworkSwitchStpRequest) Set(val *UpdateNetworkSwitchStpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchStpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchStpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchStpRequest(val *UpdateNetworkSwitchStpRequest) *NullableUpdateNetworkSwitchStpRequest {
	return &NullableUpdateNetworkSwitchStpRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchStpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchStpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


