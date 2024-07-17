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

// checks if the CreateDeviceLiveToolsWakeOnLan201ResponseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDeviceLiveToolsWakeOnLan201ResponseRequest{}

// CreateDeviceLiveToolsWakeOnLan201ResponseRequest The parameters of the Wake-on-LAN request
type CreateDeviceLiveToolsWakeOnLan201ResponseRequest struct {
	// Device serial number
	Serial *string `json:"serial,omitempty"`
	// The target's VLAN (1 to 4094)
	VlanId *int32 `json:"vlanId,omitempty"`
	// The target's MAC address
	Mac *string `json:"mac,omitempty"`
}

// NewCreateDeviceLiveToolsWakeOnLan201ResponseRequest instantiates a new CreateDeviceLiveToolsWakeOnLan201ResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceLiveToolsWakeOnLan201ResponseRequest() *CreateDeviceLiveToolsWakeOnLan201ResponseRequest {
	this := CreateDeviceLiveToolsWakeOnLan201ResponseRequest{}
	return &this
}

// NewCreateDeviceLiveToolsWakeOnLan201ResponseRequestWithDefaults instantiates a new CreateDeviceLiveToolsWakeOnLan201ResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceLiveToolsWakeOnLan201ResponseRequestWithDefaults() *CreateDeviceLiveToolsWakeOnLan201ResponseRequest {
	this := CreateDeviceLiveToolsWakeOnLan201ResponseRequest{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsWakeOnLan201ResponseRequest) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsWakeOnLan201ResponseRequest) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsWakeOnLan201ResponseRequest) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *CreateDeviceLiveToolsWakeOnLan201ResponseRequest) SetSerial(v string) {
	o.Serial = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsWakeOnLan201ResponseRequest) GetVlanId() int32 {
	if o == nil || IsNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsWakeOnLan201ResponseRequest) GetVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsWakeOnLan201ResponseRequest) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *CreateDeviceLiveToolsWakeOnLan201ResponseRequest) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsWakeOnLan201ResponseRequest) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsWakeOnLan201ResponseRequest) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsWakeOnLan201ResponseRequest) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *CreateDeviceLiveToolsWakeOnLan201ResponseRequest) SetMac(v string) {
	o.Mac = &v
}

func (o CreateDeviceLiveToolsWakeOnLan201ResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDeviceLiveToolsWakeOnLan201ResponseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	return toSerialize, nil
}

type NullableCreateDeviceLiveToolsWakeOnLan201ResponseRequest struct {
	value *CreateDeviceLiveToolsWakeOnLan201ResponseRequest
	isSet bool
}

func (v NullableCreateDeviceLiveToolsWakeOnLan201ResponseRequest) Get() *CreateDeviceLiveToolsWakeOnLan201ResponseRequest {
	return v.value
}

func (v *NullableCreateDeviceLiveToolsWakeOnLan201ResponseRequest) Set(val *CreateDeviceLiveToolsWakeOnLan201ResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceLiveToolsWakeOnLan201ResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceLiveToolsWakeOnLan201ResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceLiveToolsWakeOnLan201ResponseRequest(val *CreateDeviceLiveToolsWakeOnLan201ResponseRequest) *NullableCreateDeviceLiveToolsWakeOnLan201ResponseRequest {
	return &NullableCreateDeviceLiveToolsWakeOnLan201ResponseRequest{value: val, isSet: true}
}

func (v NullableCreateDeviceLiveToolsWakeOnLan201ResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceLiveToolsWakeOnLan201ResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

