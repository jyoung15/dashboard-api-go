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

// checks if the GetDeviceSensorRelationships200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceSensorRelationships200Response{}

// GetDeviceSensorRelationships200Response struct for GetDeviceSensorRelationships200Response
type GetDeviceSensorRelationships200Response struct {
	Livestream *GetDeviceSensorRelationships200ResponseLivestream `json:"livestream,omitempty"`
}

// NewGetDeviceSensorRelationships200Response instantiates a new GetDeviceSensorRelationships200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceSensorRelationships200Response() *GetDeviceSensorRelationships200Response {
	this := GetDeviceSensorRelationships200Response{}
	return &this
}

// NewGetDeviceSensorRelationships200ResponseWithDefaults instantiates a new GetDeviceSensorRelationships200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceSensorRelationships200ResponseWithDefaults() *GetDeviceSensorRelationships200Response {
	this := GetDeviceSensorRelationships200Response{}
	return &this
}

// GetLivestream returns the Livestream field value if set, zero value otherwise.
func (o *GetDeviceSensorRelationships200Response) GetLivestream() GetDeviceSensorRelationships200ResponseLivestream {
	if o == nil || IsNil(o.Livestream) {
		var ret GetDeviceSensorRelationships200ResponseLivestream
		return ret
	}
	return *o.Livestream
}

// GetLivestreamOk returns a tuple with the Livestream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSensorRelationships200Response) GetLivestreamOk() (*GetDeviceSensorRelationships200ResponseLivestream, bool) {
	if o == nil || IsNil(o.Livestream) {
		return nil, false
	}
	return o.Livestream, true
}

// HasLivestream returns a boolean if a field has been set.
func (o *GetDeviceSensorRelationships200Response) HasLivestream() bool {
	if o != nil && !IsNil(o.Livestream) {
		return true
	}

	return false
}

// SetLivestream gets a reference to the given GetDeviceSensorRelationships200ResponseLivestream and assigns it to the Livestream field.
func (o *GetDeviceSensorRelationships200Response) SetLivestream(v GetDeviceSensorRelationships200ResponseLivestream) {
	o.Livestream = &v
}

func (o GetDeviceSensorRelationships200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceSensorRelationships200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Livestream) {
		toSerialize["livestream"] = o.Livestream
	}
	return toSerialize, nil
}

type NullableGetDeviceSensorRelationships200Response struct {
	value *GetDeviceSensorRelationships200Response
	isSet bool
}

func (v NullableGetDeviceSensorRelationships200Response) Get() *GetDeviceSensorRelationships200Response {
	return v.value
}

func (v *NullableGetDeviceSensorRelationships200Response) Set(val *GetDeviceSensorRelationships200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceSensorRelationships200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceSensorRelationships200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceSensorRelationships200Response(val *GetDeviceSensorRelationships200Response) *NullableGetDeviceSensorRelationships200Response {
	return &NullableGetDeviceSensorRelationships200Response{value: val, isSet: true}
}

func (v NullableGetDeviceSensorRelationships200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceSensorRelationships200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


