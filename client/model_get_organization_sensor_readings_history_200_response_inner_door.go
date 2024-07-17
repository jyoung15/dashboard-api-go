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

// checks if the GetOrganizationSensorReadingsHistory200ResponseInnerDoor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSensorReadingsHistory200ResponseInnerDoor{}

// GetOrganizationSensorReadingsHistory200ResponseInnerDoor Reading for the 'door' metric. This will only be present if the 'metric' property equals 'door'.
type GetOrganizationSensorReadingsHistory200ResponseInnerDoor struct {
	// True if the door is open.
	Open *bool `json:"open,omitempty"`
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerDoor instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerDoor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSensorReadingsHistory200ResponseInnerDoor() *GetOrganizationSensorReadingsHistory200ResponseInnerDoor {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerDoor{}
	return &this
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerDoorWithDefaults instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerDoor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSensorReadingsHistory200ResponseInnerDoorWithDefaults() *GetOrganizationSensorReadingsHistory200ResponseInnerDoor {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerDoor{}
	return &this
}

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerDoor) GetOpen() bool {
	if o == nil || IsNil(o.Open) {
		var ret bool
		return ret
	}
	return *o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerDoor) GetOpenOk() (*bool, bool) {
	if o == nil || IsNil(o.Open) {
		return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerDoor) HasOpen() bool {
	if o != nil && !IsNil(o.Open) {
		return true
	}

	return false
}

// SetOpen gets a reference to the given bool and assigns it to the Open field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerDoor) SetOpen(v bool) {
	o.Open = &v
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerDoor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerDoor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Open) {
		toSerialize["open"] = o.Open
	}
	return toSerialize, nil
}

type NullableGetOrganizationSensorReadingsHistory200ResponseInnerDoor struct {
	value *GetOrganizationSensorReadingsHistory200ResponseInnerDoor
	isSet bool
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerDoor) Get() *GetOrganizationSensorReadingsHistory200ResponseInnerDoor {
	return v.value
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerDoor) Set(val *GetOrganizationSensorReadingsHistory200ResponseInnerDoor) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerDoor) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerDoor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSensorReadingsHistory200ResponseInnerDoor(val *GetOrganizationSensorReadingsHistory200ResponseInnerDoor) *NullableGetOrganizationSensorReadingsHistory200ResponseInnerDoor {
	return &NullableGetOrganizationSensorReadingsHistory200ResponseInnerDoor{value: val, isSet: true}
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerDoor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerDoor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


