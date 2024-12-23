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

// checks if the GetNetworkSwitchSettings200ResponsePowerExceptionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchSettings200ResponsePowerExceptionsInner{}

// GetNetworkSwitchSettings200ResponsePowerExceptionsInner struct for GetNetworkSwitchSettings200ResponsePowerExceptionsInner
type GetNetworkSwitchSettings200ResponsePowerExceptionsInner struct {
	// Serial number of the switch
	Serial *string `json:"serial,omitempty"`
	// Per switch exception (combined, redundant, useNetworkSetting)
	PowerType *string `json:"powerType,omitempty"`
}

// NewGetNetworkSwitchSettings200ResponsePowerExceptionsInner instantiates a new GetNetworkSwitchSettings200ResponsePowerExceptionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchSettings200ResponsePowerExceptionsInner() *GetNetworkSwitchSettings200ResponsePowerExceptionsInner {
	this := GetNetworkSwitchSettings200ResponsePowerExceptionsInner{}
	return &this
}

// NewGetNetworkSwitchSettings200ResponsePowerExceptionsInnerWithDefaults instantiates a new GetNetworkSwitchSettings200ResponsePowerExceptionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchSettings200ResponsePowerExceptionsInnerWithDefaults() *GetNetworkSwitchSettings200ResponsePowerExceptionsInner {
	this := GetNetworkSwitchSettings200ResponsePowerExceptionsInner{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetNetworkSwitchSettings200ResponsePowerExceptionsInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchSettings200ResponsePowerExceptionsInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetNetworkSwitchSettings200ResponsePowerExceptionsInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetNetworkSwitchSettings200ResponsePowerExceptionsInner) SetSerial(v string) {
	o.Serial = &v
}

// GetPowerType returns the PowerType field value if set, zero value otherwise.
func (o *GetNetworkSwitchSettings200ResponsePowerExceptionsInner) GetPowerType() string {
	if o == nil || IsNil(o.PowerType) {
		var ret string
		return ret
	}
	return *o.PowerType
}

// GetPowerTypeOk returns a tuple with the PowerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchSettings200ResponsePowerExceptionsInner) GetPowerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PowerType) {
		return nil, false
	}
	return o.PowerType, true
}

// HasPowerType returns a boolean if a field has been set.
func (o *GetNetworkSwitchSettings200ResponsePowerExceptionsInner) HasPowerType() bool {
	if o != nil && !IsNil(o.PowerType) {
		return true
	}

	return false
}

// SetPowerType gets a reference to the given string and assigns it to the PowerType field.
func (o *GetNetworkSwitchSettings200ResponsePowerExceptionsInner) SetPowerType(v string) {
	o.PowerType = &v
}

func (o GetNetworkSwitchSettings200ResponsePowerExceptionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchSettings200ResponsePowerExceptionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.PowerType) {
		toSerialize["powerType"] = o.PowerType
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchSettings200ResponsePowerExceptionsInner struct {
	value *GetNetworkSwitchSettings200ResponsePowerExceptionsInner
	isSet bool
}

func (v NullableGetNetworkSwitchSettings200ResponsePowerExceptionsInner) Get() *GetNetworkSwitchSettings200ResponsePowerExceptionsInner {
	return v.value
}

func (v *NullableGetNetworkSwitchSettings200ResponsePowerExceptionsInner) Set(val *GetNetworkSwitchSettings200ResponsePowerExceptionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchSettings200ResponsePowerExceptionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchSettings200ResponsePowerExceptionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchSettings200ResponsePowerExceptionsInner(val *GetNetworkSwitchSettings200ResponsePowerExceptionsInner) *NullableGetNetworkSwitchSettings200ResponsePowerExceptionsInner {
	return &NullableGetNetworkSwitchSettings200ResponsePowerExceptionsInner{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchSettings200ResponsePowerExceptionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchSettings200ResponsePowerExceptionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


