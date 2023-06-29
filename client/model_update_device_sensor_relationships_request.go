/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateDeviceSensorRelationshipsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceSensorRelationshipsRequest{}

// UpdateDeviceSensorRelationshipsRequest struct for UpdateDeviceSensorRelationshipsRequest
type UpdateDeviceSensorRelationshipsRequest struct {
	Livestream *UpdateDeviceSensorRelationshipsRequestLivestream `json:"livestream,omitempty"`
}

// NewUpdateDeviceSensorRelationshipsRequest instantiates a new UpdateDeviceSensorRelationshipsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceSensorRelationshipsRequest() *UpdateDeviceSensorRelationshipsRequest {
	this := UpdateDeviceSensorRelationshipsRequest{}
	return &this
}

// NewUpdateDeviceSensorRelationshipsRequestWithDefaults instantiates a new UpdateDeviceSensorRelationshipsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceSensorRelationshipsRequestWithDefaults() *UpdateDeviceSensorRelationshipsRequest {
	this := UpdateDeviceSensorRelationshipsRequest{}
	return &this
}

// GetLivestream returns the Livestream field value if set, zero value otherwise.
func (o *UpdateDeviceSensorRelationshipsRequest) GetLivestream() UpdateDeviceSensorRelationshipsRequestLivestream {
	if o == nil || IsNil(o.Livestream) {
		var ret UpdateDeviceSensorRelationshipsRequestLivestream
		return ret
	}
	return *o.Livestream
}

// GetLivestreamOk returns a tuple with the Livestream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSensorRelationshipsRequest) GetLivestreamOk() (*UpdateDeviceSensorRelationshipsRequestLivestream, bool) {
	if o == nil || IsNil(o.Livestream) {
		return nil, false
	}
	return o.Livestream, true
}

// HasLivestream returns a boolean if a field has been set.
func (o *UpdateDeviceSensorRelationshipsRequest) HasLivestream() bool {
	if o != nil && !IsNil(o.Livestream) {
		return true
	}

	return false
}

// SetLivestream gets a reference to the given UpdateDeviceSensorRelationshipsRequestLivestream and assigns it to the Livestream field.
func (o *UpdateDeviceSensorRelationshipsRequest) SetLivestream(v UpdateDeviceSensorRelationshipsRequestLivestream) {
	o.Livestream = &v
}

func (o UpdateDeviceSensorRelationshipsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceSensorRelationshipsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Livestream) {
		toSerialize["livestream"] = o.Livestream
	}
	return toSerialize, nil
}

type NullableUpdateDeviceSensorRelationshipsRequest struct {
	value *UpdateDeviceSensorRelationshipsRequest
	isSet bool
}

func (v NullableUpdateDeviceSensorRelationshipsRequest) Get() *UpdateDeviceSensorRelationshipsRequest {
	return v.value
}

func (v *NullableUpdateDeviceSensorRelationshipsRequest) Set(val *UpdateDeviceSensorRelationshipsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceSensorRelationshipsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceSensorRelationshipsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceSensorRelationshipsRequest(val *UpdateDeviceSensorRelationshipsRequest) *NullableUpdateDeviceSensorRelationshipsRequest {
	return &NullableUpdateDeviceSensorRelationshipsRequest{value: val, isSet: true}
}

func (v NullableUpdateDeviceSensorRelationshipsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceSensorRelationshipsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


