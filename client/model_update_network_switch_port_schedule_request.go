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

// checks if the UpdateNetworkSwitchPortScheduleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchPortScheduleRequest{}

// UpdateNetworkSwitchPortScheduleRequest struct for UpdateNetworkSwitchPortScheduleRequest
type UpdateNetworkSwitchPortScheduleRequest struct {
	// The name for your port schedule.
	Name *string `json:"name,omitempty"`
	PortSchedule *CreateNetworkSwitchPortScheduleRequestPortSchedule `json:"portSchedule,omitempty"`
}

// NewUpdateNetworkSwitchPortScheduleRequest instantiates a new UpdateNetworkSwitchPortScheduleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchPortScheduleRequest() *UpdateNetworkSwitchPortScheduleRequest {
	this := UpdateNetworkSwitchPortScheduleRequest{}
	return &this
}

// NewUpdateNetworkSwitchPortScheduleRequestWithDefaults instantiates a new UpdateNetworkSwitchPortScheduleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchPortScheduleRequestWithDefaults() *UpdateNetworkSwitchPortScheduleRequest {
	this := UpdateNetworkSwitchPortScheduleRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchPortScheduleRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchPortScheduleRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchPortScheduleRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkSwitchPortScheduleRequest) SetName(v string) {
	o.Name = &v
}

// GetPortSchedule returns the PortSchedule field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchPortScheduleRequest) GetPortSchedule() CreateNetworkSwitchPortScheduleRequestPortSchedule {
	if o == nil || IsNil(o.PortSchedule) {
		var ret CreateNetworkSwitchPortScheduleRequestPortSchedule
		return ret
	}
	return *o.PortSchedule
}

// GetPortScheduleOk returns a tuple with the PortSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchPortScheduleRequest) GetPortScheduleOk() (*CreateNetworkSwitchPortScheduleRequestPortSchedule, bool) {
	if o == nil || IsNil(o.PortSchedule) {
		return nil, false
	}
	return o.PortSchedule, true
}

// HasPortSchedule returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchPortScheduleRequest) HasPortSchedule() bool {
	if o != nil && !IsNil(o.PortSchedule) {
		return true
	}

	return false
}

// SetPortSchedule gets a reference to the given CreateNetworkSwitchPortScheduleRequestPortSchedule and assigns it to the PortSchedule field.
func (o *UpdateNetworkSwitchPortScheduleRequest) SetPortSchedule(v CreateNetworkSwitchPortScheduleRequestPortSchedule) {
	o.PortSchedule = &v
}

func (o UpdateNetworkSwitchPortScheduleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchPortScheduleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PortSchedule) {
		toSerialize["portSchedule"] = o.PortSchedule
	}
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchPortScheduleRequest struct {
	value *UpdateNetworkSwitchPortScheduleRequest
	isSet bool
}

func (v NullableUpdateNetworkSwitchPortScheduleRequest) Get() *UpdateNetworkSwitchPortScheduleRequest {
	return v.value
}

func (v *NullableUpdateNetworkSwitchPortScheduleRequest) Set(val *UpdateNetworkSwitchPortScheduleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchPortScheduleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchPortScheduleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchPortScheduleRequest(val *UpdateNetworkSwitchPortScheduleRequest) *NullableUpdateNetworkSwitchPortScheduleRequest {
	return &NullableUpdateNetworkSwitchPortScheduleRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchPortScheduleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchPortScheduleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


