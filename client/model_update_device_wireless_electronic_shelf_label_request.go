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

// checks if the UpdateDeviceWirelessElectronicShelfLabelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceWirelessElectronicShelfLabelRequest{}

// UpdateDeviceWirelessElectronicShelfLabelRequest struct for UpdateDeviceWirelessElectronicShelfLabelRequest
type UpdateDeviceWirelessElectronicShelfLabelRequest struct {
	// Desired ESL channel for the device, or 'Auto' (case insensitive) to use the recommended channel
	Channel *string `json:"channel,omitempty"`
	// Turn ESL features on and off for this device
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUpdateDeviceWirelessElectronicShelfLabelRequest instantiates a new UpdateDeviceWirelessElectronicShelfLabelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceWirelessElectronicShelfLabelRequest() *UpdateDeviceWirelessElectronicShelfLabelRequest {
	this := UpdateDeviceWirelessElectronicShelfLabelRequest{}
	return &this
}

// NewUpdateDeviceWirelessElectronicShelfLabelRequestWithDefaults instantiates a new UpdateDeviceWirelessElectronicShelfLabelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceWirelessElectronicShelfLabelRequestWithDefaults() *UpdateDeviceWirelessElectronicShelfLabelRequest {
	this := UpdateDeviceWirelessElectronicShelfLabelRequest{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *UpdateDeviceWirelessElectronicShelfLabelRequest) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceWirelessElectronicShelfLabelRequest) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *UpdateDeviceWirelessElectronicShelfLabelRequest) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *UpdateDeviceWirelessElectronicShelfLabelRequest) SetChannel(v string) {
	o.Channel = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateDeviceWirelessElectronicShelfLabelRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceWirelessElectronicShelfLabelRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateDeviceWirelessElectronicShelfLabelRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateDeviceWirelessElectronicShelfLabelRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UpdateDeviceWirelessElectronicShelfLabelRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceWirelessElectronicShelfLabelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableUpdateDeviceWirelessElectronicShelfLabelRequest struct {
	value *UpdateDeviceWirelessElectronicShelfLabelRequest
	isSet bool
}

func (v NullableUpdateDeviceWirelessElectronicShelfLabelRequest) Get() *UpdateDeviceWirelessElectronicShelfLabelRequest {
	return v.value
}

func (v *NullableUpdateDeviceWirelessElectronicShelfLabelRequest) Set(val *UpdateDeviceWirelessElectronicShelfLabelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceWirelessElectronicShelfLabelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceWirelessElectronicShelfLabelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceWirelessElectronicShelfLabelRequest(val *UpdateDeviceWirelessElectronicShelfLabelRequest) *NullableUpdateDeviceWirelessElectronicShelfLabelRequest {
	return &NullableUpdateDeviceWirelessElectronicShelfLabelRequest{value: val, isSet: true}
}

func (v NullableUpdateDeviceWirelessElectronicShelfLabelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceWirelessElectronicShelfLabelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

