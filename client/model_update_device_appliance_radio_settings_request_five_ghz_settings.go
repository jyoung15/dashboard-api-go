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

// checks if the UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings{}

// UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings Manual radio settings for 5 GHz.
type UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings struct {
	// Sets a manual channel for 5 GHz. Can be '36', '40', '44', '48', '52', '56', '60', '64', '100', '104', '108', '112', '116', '120', '124', '128', '132', '136', '140', '144', '149', '153', '157', '161', '165', '169', '173' or '177' or null for using auto channel.
	Channel *int32 `json:"channel,omitempty"`
	// Sets a manual channel width for 5 GHz. Can be '0', '20', '40', '80' or '160' or null for using auto channel width.
	ChannelWidth *int32 `json:"channelWidth,omitempty"`
	// Set a manual target power for 5 GHz (dBm). Enter null for using auto power range.
	TargetPower *int32 `json:"targetPower,omitempty"`
}

// NewUpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings instantiates a new UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings() *UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings {
	this := UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings{}
	return &this
}

// NewUpdateDeviceApplianceRadioSettingsRequestFiveGhzSettingsWithDefaults instantiates a new UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceApplianceRadioSettingsRequestFiveGhzSettingsWithDefaults() *UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings {
	this := UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) GetChannel() int32 {
	if o == nil || IsNil(o.Channel) {
		var ret int32
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) GetChannelOk() (*int32, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given int32 and assigns it to the Channel field.
func (o *UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) SetChannel(v int32) {
	o.Channel = &v
}

// GetChannelWidth returns the ChannelWidth field value if set, zero value otherwise.
func (o *UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) GetChannelWidth() int32 {
	if o == nil || IsNil(o.ChannelWidth) {
		var ret int32
		return ret
	}
	return *o.ChannelWidth
}

// GetChannelWidthOk returns a tuple with the ChannelWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) GetChannelWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.ChannelWidth) {
		return nil, false
	}
	return o.ChannelWidth, true
}

// HasChannelWidth returns a boolean if a field has been set.
func (o *UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) HasChannelWidth() bool {
	if o != nil && !IsNil(o.ChannelWidth) {
		return true
	}

	return false
}

// SetChannelWidth gets a reference to the given int32 and assigns it to the ChannelWidth field.
func (o *UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) SetChannelWidth(v int32) {
	o.ChannelWidth = &v
}

// GetTargetPower returns the TargetPower field value if set, zero value otherwise.
func (o *UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) GetTargetPower() int32 {
	if o == nil || IsNil(o.TargetPower) {
		var ret int32
		return ret
	}
	return *o.TargetPower
}

// GetTargetPowerOk returns a tuple with the TargetPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) GetTargetPowerOk() (*int32, bool) {
	if o == nil || IsNil(o.TargetPower) {
		return nil, false
	}
	return o.TargetPower, true
}

// HasTargetPower returns a boolean if a field has been set.
func (o *UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) HasTargetPower() bool {
	if o != nil && !IsNil(o.TargetPower) {
		return true
	}

	return false
}

// SetTargetPower gets a reference to the given int32 and assigns it to the TargetPower field.
func (o *UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) SetTargetPower(v int32) {
	o.TargetPower = &v
}

func (o UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.ChannelWidth) {
		toSerialize["channelWidth"] = o.ChannelWidth
	}
	if !IsNil(o.TargetPower) {
		toSerialize["targetPower"] = o.TargetPower
	}
	return toSerialize, nil
}

type NullableUpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings struct {
	value *UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings
	isSet bool
}

func (v NullableUpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) Get() *UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings {
	return v.value
}

func (v *NullableUpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) Set(val *UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings(val *UpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) *NullableUpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings {
	return &NullableUpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings{value: val, isSet: true}
}

func (v NullableUpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceApplianceRadioSettingsRequestFiveGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


