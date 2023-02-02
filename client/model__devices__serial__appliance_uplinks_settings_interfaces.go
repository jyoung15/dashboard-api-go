/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DevicesSerialApplianceUplinksSettingsInterfaces Interface settings.
type DevicesSerialApplianceUplinksSettingsInterfaces struct {
	Wan1 *DevicesSerialApplianceUplinksSettingsInterfacesWan1 `json:"wan1,omitempty"`
	Wan2 *DevicesSerialApplianceUplinksSettingsInterfacesWan2 `json:"wan2,omitempty"`
}

// NewDevicesSerialApplianceUplinksSettingsInterfaces instantiates a new DevicesSerialApplianceUplinksSettingsInterfaces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialApplianceUplinksSettingsInterfaces() *DevicesSerialApplianceUplinksSettingsInterfaces {
	this := DevicesSerialApplianceUplinksSettingsInterfaces{}
	return &this
}

// NewDevicesSerialApplianceUplinksSettingsInterfacesWithDefaults instantiates a new DevicesSerialApplianceUplinksSettingsInterfaces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialApplianceUplinksSettingsInterfacesWithDefaults() *DevicesSerialApplianceUplinksSettingsInterfaces {
	this := DevicesSerialApplianceUplinksSettingsInterfaces{}
	return &this
}

// GetWan1 returns the Wan1 field value if set, zero value otherwise.
func (o *DevicesSerialApplianceUplinksSettingsInterfaces) GetWan1() DevicesSerialApplianceUplinksSettingsInterfacesWan1 {
	if o == nil || isNil(o.Wan1) {
		var ret DevicesSerialApplianceUplinksSettingsInterfacesWan1
		return ret
	}
	return *o.Wan1
}

// GetWan1Ok returns a tuple with the Wan1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialApplianceUplinksSettingsInterfaces) GetWan1Ok() (*DevicesSerialApplianceUplinksSettingsInterfacesWan1, bool) {
	if o == nil || isNil(o.Wan1) {
    return nil, false
	}
	return o.Wan1, true
}

// HasWan1 returns a boolean if a field has been set.
func (o *DevicesSerialApplianceUplinksSettingsInterfaces) HasWan1() bool {
	if o != nil && !isNil(o.Wan1) {
		return true
	}

	return false
}

// SetWan1 gets a reference to the given DevicesSerialApplianceUplinksSettingsInterfacesWan1 and assigns it to the Wan1 field.
func (o *DevicesSerialApplianceUplinksSettingsInterfaces) SetWan1(v DevicesSerialApplianceUplinksSettingsInterfacesWan1) {
	o.Wan1 = &v
}

// GetWan2 returns the Wan2 field value if set, zero value otherwise.
func (o *DevicesSerialApplianceUplinksSettingsInterfaces) GetWan2() DevicesSerialApplianceUplinksSettingsInterfacesWan2 {
	if o == nil || isNil(o.Wan2) {
		var ret DevicesSerialApplianceUplinksSettingsInterfacesWan2
		return ret
	}
	return *o.Wan2
}

// GetWan2Ok returns a tuple with the Wan2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialApplianceUplinksSettingsInterfaces) GetWan2Ok() (*DevicesSerialApplianceUplinksSettingsInterfacesWan2, bool) {
	if o == nil || isNil(o.Wan2) {
    return nil, false
	}
	return o.Wan2, true
}

// HasWan2 returns a boolean if a field has been set.
func (o *DevicesSerialApplianceUplinksSettingsInterfaces) HasWan2() bool {
	if o != nil && !isNil(o.Wan2) {
		return true
	}

	return false
}

// SetWan2 gets a reference to the given DevicesSerialApplianceUplinksSettingsInterfacesWan2 and assigns it to the Wan2 field.
func (o *DevicesSerialApplianceUplinksSettingsInterfaces) SetWan2(v DevicesSerialApplianceUplinksSettingsInterfacesWan2) {
	o.Wan2 = &v
}

func (o DevicesSerialApplianceUplinksSettingsInterfaces) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Wan1) {
		toSerialize["wan1"] = o.Wan1
	}
	if !isNil(o.Wan2) {
		toSerialize["wan2"] = o.Wan2
	}
	return json.Marshal(toSerialize)
}

type NullableDevicesSerialApplianceUplinksSettingsInterfaces struct {
	value *DevicesSerialApplianceUplinksSettingsInterfaces
	isSet bool
}

func (v NullableDevicesSerialApplianceUplinksSettingsInterfaces) Get() *DevicesSerialApplianceUplinksSettingsInterfaces {
	return v.value
}

func (v *NullableDevicesSerialApplianceUplinksSettingsInterfaces) Set(val *DevicesSerialApplianceUplinksSettingsInterfaces) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialApplianceUplinksSettingsInterfaces) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialApplianceUplinksSettingsInterfaces) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialApplianceUplinksSettingsInterfaces(val *DevicesSerialApplianceUplinksSettingsInterfaces) *NullableDevicesSerialApplianceUplinksSettingsInterfaces {
	return &NullableDevicesSerialApplianceUplinksSettingsInterfaces{value: val, isSet: true}
}

func (v NullableDevicesSerialApplianceUplinksSettingsInterfaces) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialApplianceUplinksSettingsInterfaces) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

