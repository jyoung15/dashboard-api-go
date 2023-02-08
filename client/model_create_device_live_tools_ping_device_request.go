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

// CreateDeviceLiveToolsPingDeviceRequest struct for CreateDeviceLiveToolsPingDeviceRequest
type CreateDeviceLiveToolsPingDeviceRequest struct {
	// Count parameter to pass to ping. [1..5], default 5
	Count *int32 `json:"count,omitempty"`
}

// NewCreateDeviceLiveToolsPingDeviceRequest instantiates a new CreateDeviceLiveToolsPingDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceLiveToolsPingDeviceRequest() *CreateDeviceLiveToolsPingDeviceRequest {
	this := CreateDeviceLiveToolsPingDeviceRequest{}
	return &this
}

// NewCreateDeviceLiveToolsPingDeviceRequestWithDefaults instantiates a new CreateDeviceLiveToolsPingDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceLiveToolsPingDeviceRequestWithDefaults() *CreateDeviceLiveToolsPingDeviceRequest {
	this := CreateDeviceLiveToolsPingDeviceRequest{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsPingDeviceRequest) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsPingDeviceRequest) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsPingDeviceRequest) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CreateDeviceLiveToolsPingDeviceRequest) SetCount(v int32) {
	o.Count = &v
}

func (o CreateDeviceLiveToolsPingDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDeviceLiveToolsPingDeviceRequest struct {
	value *CreateDeviceLiveToolsPingDeviceRequest
	isSet bool
}

func (v NullableCreateDeviceLiveToolsPingDeviceRequest) Get() *CreateDeviceLiveToolsPingDeviceRequest {
	return v.value
}

func (v *NullableCreateDeviceLiveToolsPingDeviceRequest) Set(val *CreateDeviceLiveToolsPingDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceLiveToolsPingDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceLiveToolsPingDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceLiveToolsPingDeviceRequest(val *CreateDeviceLiveToolsPingDeviceRequest) *NullableCreateDeviceLiveToolsPingDeviceRequest {
	return &NullableCreateDeviceLiveToolsPingDeviceRequest{value: val, isSet: true}
}

func (v NullableCreateDeviceLiveToolsPingDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceLiveToolsPingDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


