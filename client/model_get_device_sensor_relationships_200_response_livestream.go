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

// checks if the GetDeviceSensorRelationships200ResponseLivestream type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceSensorRelationships200ResponseLivestream{}

// GetDeviceSensorRelationships200ResponseLivestream A role defined between an MT sensor and an MV camera that adds the camera's livestream to the sensor's details page. Snapshots from the camera will also appear in alert notifications that the sensor triggers.
type GetDeviceSensorRelationships200ResponseLivestream struct {
	// An array of the related devices for the role
	RelatedDevices []GetDeviceSensorRelationships200ResponseLivestreamRelatedDevicesInner `json:"relatedDevices,omitempty"`
}

// NewGetDeviceSensorRelationships200ResponseLivestream instantiates a new GetDeviceSensorRelationships200ResponseLivestream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceSensorRelationships200ResponseLivestream() *GetDeviceSensorRelationships200ResponseLivestream {
	this := GetDeviceSensorRelationships200ResponseLivestream{}
	return &this
}

// NewGetDeviceSensorRelationships200ResponseLivestreamWithDefaults instantiates a new GetDeviceSensorRelationships200ResponseLivestream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceSensorRelationships200ResponseLivestreamWithDefaults() *GetDeviceSensorRelationships200ResponseLivestream {
	this := GetDeviceSensorRelationships200ResponseLivestream{}
	return &this
}

// GetRelatedDevices returns the RelatedDevices field value if set, zero value otherwise.
func (o *GetDeviceSensorRelationships200ResponseLivestream) GetRelatedDevices() []GetDeviceSensorRelationships200ResponseLivestreamRelatedDevicesInner {
	if o == nil || IsNil(o.RelatedDevices) {
		var ret []GetDeviceSensorRelationships200ResponseLivestreamRelatedDevicesInner
		return ret
	}
	return o.RelatedDevices
}

// GetRelatedDevicesOk returns a tuple with the RelatedDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSensorRelationships200ResponseLivestream) GetRelatedDevicesOk() ([]GetDeviceSensorRelationships200ResponseLivestreamRelatedDevicesInner, bool) {
	if o == nil || IsNil(o.RelatedDevices) {
		return nil, false
	}
	return o.RelatedDevices, true
}

// HasRelatedDevices returns a boolean if a field has been set.
func (o *GetDeviceSensorRelationships200ResponseLivestream) HasRelatedDevices() bool {
	if o != nil && !IsNil(o.RelatedDevices) {
		return true
	}

	return false
}

// SetRelatedDevices gets a reference to the given []GetDeviceSensorRelationships200ResponseLivestreamRelatedDevicesInner and assigns it to the RelatedDevices field.
func (o *GetDeviceSensorRelationships200ResponseLivestream) SetRelatedDevices(v []GetDeviceSensorRelationships200ResponseLivestreamRelatedDevicesInner) {
	o.RelatedDevices = v
}

func (o GetDeviceSensorRelationships200ResponseLivestream) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceSensorRelationships200ResponseLivestream) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RelatedDevices) {
		toSerialize["relatedDevices"] = o.RelatedDevices
	}
	return toSerialize, nil
}

type NullableGetDeviceSensorRelationships200ResponseLivestream struct {
	value *GetDeviceSensorRelationships200ResponseLivestream
	isSet bool
}

func (v NullableGetDeviceSensorRelationships200ResponseLivestream) Get() *GetDeviceSensorRelationships200ResponseLivestream {
	return v.value
}

func (v *NullableGetDeviceSensorRelationships200ResponseLivestream) Set(val *GetDeviceSensorRelationships200ResponseLivestream) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceSensorRelationships200ResponseLivestream) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceSensorRelationships200ResponseLivestream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceSensorRelationships200ResponseLivestream(val *GetDeviceSensorRelationships200ResponseLivestream) *NullableGetDeviceSensorRelationships200ResponseLivestream {
	return &NullableGetDeviceSensorRelationships200ResponseLivestream{value: val, isSet: true}
}

func (v NullableGetDeviceSensorRelationships200ResponseLivestream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceSensorRelationships200ResponseLivestream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


