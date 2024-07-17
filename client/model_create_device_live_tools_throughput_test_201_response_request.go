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

// checks if the CreateDeviceLiveToolsThroughputTest201ResponseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDeviceLiveToolsThroughputTest201ResponseRequest{}

// CreateDeviceLiveToolsThroughputTest201ResponseRequest The parameters of the throughput test request
type CreateDeviceLiveToolsThroughputTest201ResponseRequest struct {
	// Device serial number
	Serial *string `json:"serial,omitempty"`
}

// NewCreateDeviceLiveToolsThroughputTest201ResponseRequest instantiates a new CreateDeviceLiveToolsThroughputTest201ResponseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceLiveToolsThroughputTest201ResponseRequest() *CreateDeviceLiveToolsThroughputTest201ResponseRequest {
	this := CreateDeviceLiveToolsThroughputTest201ResponseRequest{}
	return &this
}

// NewCreateDeviceLiveToolsThroughputTest201ResponseRequestWithDefaults instantiates a new CreateDeviceLiveToolsThroughputTest201ResponseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceLiveToolsThroughputTest201ResponseRequestWithDefaults() *CreateDeviceLiveToolsThroughputTest201ResponseRequest {
	this := CreateDeviceLiveToolsThroughputTest201ResponseRequest{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsThroughputTest201ResponseRequest) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsThroughputTest201ResponseRequest) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsThroughputTest201ResponseRequest) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *CreateDeviceLiveToolsThroughputTest201ResponseRequest) SetSerial(v string) {
	o.Serial = &v
}

func (o CreateDeviceLiveToolsThroughputTest201ResponseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDeviceLiveToolsThroughputTest201ResponseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	return toSerialize, nil
}

type NullableCreateDeviceLiveToolsThroughputTest201ResponseRequest struct {
	value *CreateDeviceLiveToolsThroughputTest201ResponseRequest
	isSet bool
}

func (v NullableCreateDeviceLiveToolsThroughputTest201ResponseRequest) Get() *CreateDeviceLiveToolsThroughputTest201ResponseRequest {
	return v.value
}

func (v *NullableCreateDeviceLiveToolsThroughputTest201ResponseRequest) Set(val *CreateDeviceLiveToolsThroughputTest201ResponseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceLiveToolsThroughputTest201ResponseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceLiveToolsThroughputTest201ResponseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceLiveToolsThroughputTest201ResponseRequest(val *CreateDeviceLiveToolsThroughputTest201ResponseRequest) *NullableCreateDeviceLiveToolsThroughputTest201ResponseRequest {
	return &NullableCreateDeviceLiveToolsThroughputTest201ResponseRequest{value: val, isSet: true}
}

func (v NullableCreateDeviceLiveToolsThroughputTest201ResponseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceLiveToolsThroughputTest201ResponseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


