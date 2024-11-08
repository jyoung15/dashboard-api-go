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

// checks if the UpdateDeviceManagementInterfaceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceManagementInterfaceRequest{}

// UpdateDeviceManagementInterfaceRequest struct for UpdateDeviceManagementInterfaceRequest
type UpdateDeviceManagementInterfaceRequest struct {
	Wan1 *UpdateDeviceManagementInterfaceRequestWan1 `json:"wan1,omitempty"`
	Wan2 *UpdateDeviceManagementInterfaceRequestWan2 `json:"wan2,omitempty"`
}

// NewUpdateDeviceManagementInterfaceRequest instantiates a new UpdateDeviceManagementInterfaceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceManagementInterfaceRequest() *UpdateDeviceManagementInterfaceRequest {
	this := UpdateDeviceManagementInterfaceRequest{}
	return &this
}

// NewUpdateDeviceManagementInterfaceRequestWithDefaults instantiates a new UpdateDeviceManagementInterfaceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceManagementInterfaceRequestWithDefaults() *UpdateDeviceManagementInterfaceRequest {
	this := UpdateDeviceManagementInterfaceRequest{}
	return &this
}

// GetWan1 returns the Wan1 field value if set, zero value otherwise.
func (o *UpdateDeviceManagementInterfaceRequest) GetWan1() UpdateDeviceManagementInterfaceRequestWan1 {
	if o == nil || IsNil(o.Wan1) {
		var ret UpdateDeviceManagementInterfaceRequestWan1
		return ret
	}
	return *o.Wan1
}

// GetWan1Ok returns a tuple with the Wan1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceManagementInterfaceRequest) GetWan1Ok() (*UpdateDeviceManagementInterfaceRequestWan1, bool) {
	if o == nil || IsNil(o.Wan1) {
		return nil, false
	}
	return o.Wan1, true
}

// HasWan1 returns a boolean if a field has been set.
func (o *UpdateDeviceManagementInterfaceRequest) HasWan1() bool {
	if o != nil && !IsNil(o.Wan1) {
		return true
	}

	return false
}

// SetWan1 gets a reference to the given UpdateDeviceManagementInterfaceRequestWan1 and assigns it to the Wan1 field.
func (o *UpdateDeviceManagementInterfaceRequest) SetWan1(v UpdateDeviceManagementInterfaceRequestWan1) {
	o.Wan1 = &v
}

// GetWan2 returns the Wan2 field value if set, zero value otherwise.
func (o *UpdateDeviceManagementInterfaceRequest) GetWan2() UpdateDeviceManagementInterfaceRequestWan2 {
	if o == nil || IsNil(o.Wan2) {
		var ret UpdateDeviceManagementInterfaceRequestWan2
		return ret
	}
	return *o.Wan2
}

// GetWan2Ok returns a tuple with the Wan2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceManagementInterfaceRequest) GetWan2Ok() (*UpdateDeviceManagementInterfaceRequestWan2, bool) {
	if o == nil || IsNil(o.Wan2) {
		return nil, false
	}
	return o.Wan2, true
}

// HasWan2 returns a boolean if a field has been set.
func (o *UpdateDeviceManagementInterfaceRequest) HasWan2() bool {
	if o != nil && !IsNil(o.Wan2) {
		return true
	}

	return false
}

// SetWan2 gets a reference to the given UpdateDeviceManagementInterfaceRequestWan2 and assigns it to the Wan2 field.
func (o *UpdateDeviceManagementInterfaceRequest) SetWan2(v UpdateDeviceManagementInterfaceRequestWan2) {
	o.Wan2 = &v
}

func (o UpdateDeviceManagementInterfaceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceManagementInterfaceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Wan1) {
		toSerialize["wan1"] = o.Wan1
	}
	if !IsNil(o.Wan2) {
		toSerialize["wan2"] = o.Wan2
	}
	return toSerialize, nil
}

type NullableUpdateDeviceManagementInterfaceRequest struct {
	value *UpdateDeviceManagementInterfaceRequest
	isSet bool
}

func (v NullableUpdateDeviceManagementInterfaceRequest) Get() *UpdateDeviceManagementInterfaceRequest {
	return v.value
}

func (v *NullableUpdateDeviceManagementInterfaceRequest) Set(val *UpdateDeviceManagementInterfaceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceManagementInterfaceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceManagementInterfaceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceManagementInterfaceRequest(val *UpdateDeviceManagementInterfaceRequest) *NullableUpdateDeviceManagementInterfaceRequest {
	return &NullableUpdateDeviceManagementInterfaceRequest{value: val, isSet: true}
}

func (v NullableUpdateDeviceManagementInterfaceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceManagementInterfaceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


