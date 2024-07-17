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

// checks if the GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack{}

// GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack The Switch Stack the device belongs to
type GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack struct {
	// ID of the Switch Stack
	Id *string `json:"id,omitempty"`
}

// NewGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack instantiates a new GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack() *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack {
	this := GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack{}
	return &this
}

// NewGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStackWithDefaults instantiates a new GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStackWithDefaults() *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack {
	this := GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack) SetId(v string) {
	o.Id = &v
}

func (o GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack struct {
	value *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack
	isSet bool
}

func (v NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack) Get() *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack {
	return v.value
}

func (v *NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack) Set(val *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack(val *GetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack) *NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack {
	return &NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack{value: val, isSet: true}
}

func (v NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkVlanProfilesAssignmentsByDevice200ResponseInnerStack) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

