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

// checks if the UpdateDeviceCameraVideoSettingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceCameraVideoSettingsRequest{}

// UpdateDeviceCameraVideoSettingsRequest struct for UpdateDeviceCameraVideoSettingsRequest
type UpdateDeviceCameraVideoSettingsRequest struct {
	// Boolean indicating if external rtsp stream is exposed
	ExternalRtspEnabled *bool `json:"externalRtspEnabled,omitempty"`
}

// NewUpdateDeviceCameraVideoSettingsRequest instantiates a new UpdateDeviceCameraVideoSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceCameraVideoSettingsRequest() *UpdateDeviceCameraVideoSettingsRequest {
	this := UpdateDeviceCameraVideoSettingsRequest{}
	return &this
}

// NewUpdateDeviceCameraVideoSettingsRequestWithDefaults instantiates a new UpdateDeviceCameraVideoSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceCameraVideoSettingsRequestWithDefaults() *UpdateDeviceCameraVideoSettingsRequest {
	this := UpdateDeviceCameraVideoSettingsRequest{}
	return &this
}

// GetExternalRtspEnabled returns the ExternalRtspEnabled field value if set, zero value otherwise.
func (o *UpdateDeviceCameraVideoSettingsRequest) GetExternalRtspEnabled() bool {
	if o == nil || IsNil(o.ExternalRtspEnabled) {
		var ret bool
		return ret
	}
	return *o.ExternalRtspEnabled
}

// GetExternalRtspEnabledOk returns a tuple with the ExternalRtspEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCameraVideoSettingsRequest) GetExternalRtspEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ExternalRtspEnabled) {
		return nil, false
	}
	return o.ExternalRtspEnabled, true
}

// HasExternalRtspEnabled returns a boolean if a field has been set.
func (o *UpdateDeviceCameraVideoSettingsRequest) HasExternalRtspEnabled() bool {
	if o != nil && !IsNil(o.ExternalRtspEnabled) {
		return true
	}

	return false
}

// SetExternalRtspEnabled gets a reference to the given bool and assigns it to the ExternalRtspEnabled field.
func (o *UpdateDeviceCameraVideoSettingsRequest) SetExternalRtspEnabled(v bool) {
	o.ExternalRtspEnabled = &v
}

func (o UpdateDeviceCameraVideoSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceCameraVideoSettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalRtspEnabled) {
		toSerialize["externalRtspEnabled"] = o.ExternalRtspEnabled
	}
	return toSerialize, nil
}

type NullableUpdateDeviceCameraVideoSettingsRequest struct {
	value *UpdateDeviceCameraVideoSettingsRequest
	isSet bool
}

func (v NullableUpdateDeviceCameraVideoSettingsRequest) Get() *UpdateDeviceCameraVideoSettingsRequest {
	return v.value
}

func (v *NullableUpdateDeviceCameraVideoSettingsRequest) Set(val *UpdateDeviceCameraVideoSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceCameraVideoSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceCameraVideoSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceCameraVideoSettingsRequest(val *UpdateDeviceCameraVideoSettingsRequest) *NullableUpdateDeviceCameraVideoSettingsRequest {
	return &NullableUpdateDeviceCameraVideoSettingsRequest{value: val, isSet: true}
}

func (v NullableUpdateDeviceCameraVideoSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceCameraVideoSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


