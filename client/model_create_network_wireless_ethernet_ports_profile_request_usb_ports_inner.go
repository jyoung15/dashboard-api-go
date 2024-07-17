/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner{}

// CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner struct for CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner
type CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner struct {
	// AP usb port name
	Name string `json:"name"`
	// AP usb port enabled
	Enabled *bool `json:"enabled,omitempty"`
	// AP usb port ssid number
	Ssid *int32 `json:"ssid,omitempty"`
}

type _CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner

// NewCreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner instantiates a new CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner(name string) *CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner {
	this := CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner{}
	this.Name = name
	return &this
}

// NewCreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInnerWithDefaults instantiates a new CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInnerWithDefaults() *CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner {
	this := CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) GetSsid() int32 {
	if o == nil || IsNil(o.Ssid) {
		var ret int32
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) GetSsidOk() (*int32, bool) {
	if o == nil || IsNil(o.Ssid) {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) HasSsid() bool {
	if o != nil && !IsNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given int32 and assigns it to the Ssid field.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) SetSsid(v int32) {
	o.Ssid = &v
}

func (o CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	return toSerialize, nil
}

func (o *CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner := _CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner)

	if err != nil {
		return err
	}

	*o = CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner(varCreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner)

	return err
}

type NullableCreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner struct {
	value *CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner
	isSet bool
}

func (v NullableCreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) Get() *CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner {
	return v.value
}

func (v *NullableCreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) Set(val *CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner(val *CreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) *NullableCreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner {
	return &NullableCreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner{value: val, isSet: true}
}

func (v NullableCreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWirelessEthernetPortsProfileRequestUsbPortsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


