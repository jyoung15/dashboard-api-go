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

// checks if the CreateDeviceLiveToolsCableTest201ResponseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDeviceLiveToolsCableTest201ResponseRequest{}

// CreateDeviceLiveToolsCableTest201ResponseRequest Cable test request parameters
type CreateDeviceLiveToolsCableTest201ResponseRequest struct {
	// Device serial number
	Serial *string `json:"serial,omitempty"`
	// A list of ports for which to perform the cable test.
	Ports []string `json:"ports,omitempty"`
}

// NewCreateDeviceLiveToolsCableTest201ResponseRequest instantiates a new CreateDeviceLiveToolsCableTest201ResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceLiveToolsCableTest201ResponseRequest() *CreateDeviceLiveToolsCableTest201ResponseRequest {
	this := CreateDeviceLiveToolsCableTest201ResponseRequest{}
	return &this
}

// NewCreateDeviceLiveToolsCableTest201ResponseRequestWithDefaults instantiates a new CreateDeviceLiveToolsCableTest201ResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceLiveToolsCableTest201ResponseRequestWithDefaults() *CreateDeviceLiveToolsCableTest201ResponseRequest {
	this := CreateDeviceLiveToolsCableTest201ResponseRequest{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsCableTest201ResponseRequest) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsCableTest201ResponseRequest) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsCableTest201ResponseRequest) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *CreateDeviceLiveToolsCableTest201ResponseRequest) SetSerial(v string) {
	o.Serial = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsCableTest201ResponseRequest) GetPorts() []string {
	if o == nil || IsNil(o.Ports) {
		var ret []string
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsCableTest201ResponseRequest) GetPortsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsCableTest201ResponseRequest) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []string and assigns it to the Ports field.
func (o *CreateDeviceLiveToolsCableTest201ResponseRequest) SetPorts(v []string) {
	o.Ports = v
}

func (o CreateDeviceLiveToolsCableTest201ResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDeviceLiveToolsCableTest201ResponseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	return toSerialize, nil
}

type NullableCreateDeviceLiveToolsCableTest201ResponseRequest struct {
	value *CreateDeviceLiveToolsCableTest201ResponseRequest
	isSet bool
}

func (v NullableCreateDeviceLiveToolsCableTest201ResponseRequest) Get() *CreateDeviceLiveToolsCableTest201ResponseRequest {
	return v.value
}

func (v *NullableCreateDeviceLiveToolsCableTest201ResponseRequest) Set(val *CreateDeviceLiveToolsCableTest201ResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceLiveToolsCableTest201ResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceLiveToolsCableTest201ResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceLiveToolsCableTest201ResponseRequest(val *CreateDeviceLiveToolsCableTest201ResponseRequest) *NullableCreateDeviceLiveToolsCableTest201ResponseRequest {
	return &NullableCreateDeviceLiveToolsCableTest201ResponseRequest{value: val, isSet: true}
}

func (v NullableCreateDeviceLiveToolsCableTest201ResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceLiveToolsCableTest201ResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

