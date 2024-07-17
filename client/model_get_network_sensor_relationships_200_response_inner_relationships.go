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

// checks if the GetNetworkSensorRelationships200ResponseInnerRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSensorRelationships200ResponseInnerRelationships{}

// GetNetworkSensorRelationships200ResponseInnerRelationships An object describing the relationships defined between the device and other devices
type GetNetworkSensorRelationships200ResponseInnerRelationships struct {
	Livestream *GetDeviceSensorRelationships200ResponseLivestream `json:"livestream,omitempty"`
}

// NewGetNetworkSensorRelationships200ResponseInnerRelationships instantiates a new GetNetworkSensorRelationships200ResponseInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSensorRelationships200ResponseInnerRelationships() *GetNetworkSensorRelationships200ResponseInnerRelationships {
	this := GetNetworkSensorRelationships200ResponseInnerRelationships{}
	return &this
}

// NewGetNetworkSensorRelationships200ResponseInnerRelationshipsWithDefaults instantiates a new GetNetworkSensorRelationships200ResponseInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSensorRelationships200ResponseInnerRelationshipsWithDefaults() *GetNetworkSensorRelationships200ResponseInnerRelationships {
	this := GetNetworkSensorRelationships200ResponseInnerRelationships{}
	return &this
}

// GetLivestream returns the Livestream field value if set, zero value otherwise.
func (o *GetNetworkSensorRelationships200ResponseInnerRelationships) GetLivestream() GetDeviceSensorRelationships200ResponseLivestream {
	if o == nil || IsNil(o.Livestream) {
		var ret GetDeviceSensorRelationships200ResponseLivestream
		return ret
	}
	return *o.Livestream
}

// GetLivestreamOk returns a tuple with the Livestream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorRelationships200ResponseInnerRelationships) GetLivestreamOk() (*GetDeviceSensorRelationships200ResponseLivestream, bool) {
	if o == nil || IsNil(o.Livestream) {
		return nil, false
	}
	return o.Livestream, true
}

// HasLivestream returns a boolean if a field has been set.
func (o *GetNetworkSensorRelationships200ResponseInnerRelationships) HasLivestream() bool {
	if o != nil && !IsNil(o.Livestream) {
		return true
	}

	return false
}

// SetLivestream gets a reference to the given GetDeviceSensorRelationships200ResponseLivestream and assigns it to the Livestream field.
func (o *GetNetworkSensorRelationships200ResponseInnerRelationships) SetLivestream(v GetDeviceSensorRelationships200ResponseLivestream) {
	o.Livestream = &v
}

func (o GetNetworkSensorRelationships200ResponseInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSensorRelationships200ResponseInnerRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Livestream) {
		toSerialize["livestream"] = o.Livestream
	}
	return toSerialize, nil
}

type NullableGetNetworkSensorRelationships200ResponseInnerRelationships struct {
	value *GetNetworkSensorRelationships200ResponseInnerRelationships
	isSet bool
}

func (v NullableGetNetworkSensorRelationships200ResponseInnerRelationships) Get() *GetNetworkSensorRelationships200ResponseInnerRelationships {
	return v.value
}

func (v *NullableGetNetworkSensorRelationships200ResponseInnerRelationships) Set(val *GetNetworkSensorRelationships200ResponseInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSensorRelationships200ResponseInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSensorRelationships200ResponseInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSensorRelationships200ResponseInnerRelationships(val *GetNetworkSensorRelationships200ResponseInnerRelationships) *NullableGetNetworkSensorRelationships200ResponseInnerRelationships {
	return &NullableGetNetworkSensorRelationships200ResponseInnerRelationships{value: val, isSet: true}
}

func (v NullableGetNetworkSensorRelationships200ResponseInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSensorRelationships200ResponseInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


