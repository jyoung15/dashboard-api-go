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

// checks if the ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile{}

// ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile The VLAN Profile
type ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile struct {
	// IName of the VLAN Profile
	Iname *string `json:"iname,omitempty"`
	// Name of the VLAN Profile
	Name *string `json:"name,omitempty"`
}

// NewReassignNetworkVlanProfilesAssignments200ResponseVlanProfile instantiates a new ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReassignNetworkVlanProfilesAssignments200ResponseVlanProfile() *ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile {
	this := ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile{}
	return &this
}

// NewReassignNetworkVlanProfilesAssignments200ResponseVlanProfileWithDefaults instantiates a new ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReassignNetworkVlanProfilesAssignments200ResponseVlanProfileWithDefaults() *ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile {
	this := ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile{}
	return &this
}

// GetIname returns the Iname field value if set, zero value otherwise.
func (o *ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile) GetIname() string {
	if o == nil || IsNil(o.Iname) {
		var ret string
		return ret
	}
	return *o.Iname
}

// GetInameOk returns a tuple with the Iname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile) GetInameOk() (*string, bool) {
	if o == nil || IsNil(o.Iname) {
		return nil, false
	}
	return o.Iname, true
}

// HasIname returns a boolean if a field has been set.
func (o *ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile) HasIname() bool {
	if o != nil && !IsNil(o.Iname) {
		return true
	}

	return false
}

// SetIname gets a reference to the given string and assigns it to the Iname field.
func (o *ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile) SetIname(v string) {
	o.Iname = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile) SetName(v string) {
	o.Name = &v
}

func (o ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Iname) {
		toSerialize["iname"] = o.Iname
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableReassignNetworkVlanProfilesAssignments200ResponseVlanProfile struct {
	value *ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile
	isSet bool
}

func (v NullableReassignNetworkVlanProfilesAssignments200ResponseVlanProfile) Get() *ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile {
	return v.value
}

func (v *NullableReassignNetworkVlanProfilesAssignments200ResponseVlanProfile) Set(val *ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableReassignNetworkVlanProfilesAssignments200ResponseVlanProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableReassignNetworkVlanProfilesAssignments200ResponseVlanProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReassignNetworkVlanProfilesAssignments200ResponseVlanProfile(val *ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile) *NullableReassignNetworkVlanProfilesAssignments200ResponseVlanProfile {
	return &NullableReassignNetworkVlanProfilesAssignments200ResponseVlanProfile{value: val, isSet: true}
}

func (v NullableReassignNetworkVlanProfilesAssignments200ResponseVlanProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReassignNetworkVlanProfilesAssignments200ResponseVlanProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

