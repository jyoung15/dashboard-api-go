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

// checks if the ReassignNetworkVlanProfilesAssignments200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReassignNetworkVlanProfilesAssignments200Response{}

// ReassignNetworkVlanProfilesAssignments200Response struct for ReassignNetworkVlanProfilesAssignments200Response
type ReassignNetworkVlanProfilesAssignments200Response struct {
	VlanProfile *ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile `json:"vlanProfile,omitempty"`
	// Array of Device Serials
	Serials []string `json:"serials,omitempty"`
	// Array of Switch Stack IDs
	StackIds []string `json:"stackIds,omitempty"`
}

// NewReassignNetworkVlanProfilesAssignments200Response instantiates a new ReassignNetworkVlanProfilesAssignments200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReassignNetworkVlanProfilesAssignments200Response() *ReassignNetworkVlanProfilesAssignments200Response {
	this := ReassignNetworkVlanProfilesAssignments200Response{}
	return &this
}

// NewReassignNetworkVlanProfilesAssignments200ResponseWithDefaults instantiates a new ReassignNetworkVlanProfilesAssignments200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReassignNetworkVlanProfilesAssignments200ResponseWithDefaults() *ReassignNetworkVlanProfilesAssignments200Response {
	this := ReassignNetworkVlanProfilesAssignments200Response{}
	return &this
}

// GetVlanProfile returns the VlanProfile field value if set, zero value otherwise.
func (o *ReassignNetworkVlanProfilesAssignments200Response) GetVlanProfile() ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile {
	if o == nil || IsNil(o.VlanProfile) {
		var ret ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile
		return ret
	}
	return *o.VlanProfile
}

// GetVlanProfileOk returns a tuple with the VlanProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReassignNetworkVlanProfilesAssignments200Response) GetVlanProfileOk() (*ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile, bool) {
	if o == nil || IsNil(o.VlanProfile) {
		return nil, false
	}
	return o.VlanProfile, true
}

// HasVlanProfile returns a boolean if a field has been set.
func (o *ReassignNetworkVlanProfilesAssignments200Response) HasVlanProfile() bool {
	if o != nil && !IsNil(o.VlanProfile) {
		return true
	}

	return false
}

// SetVlanProfile gets a reference to the given ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile and assigns it to the VlanProfile field.
func (o *ReassignNetworkVlanProfilesAssignments200Response) SetVlanProfile(v ReassignNetworkVlanProfilesAssignments200ResponseVlanProfile) {
	o.VlanProfile = &v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *ReassignNetworkVlanProfilesAssignments200Response) GetSerials() []string {
	if o == nil || IsNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReassignNetworkVlanProfilesAssignments200Response) GetSerialsOk() ([]string, bool) {
	if o == nil || IsNil(o.Serials) {
		return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *ReassignNetworkVlanProfilesAssignments200Response) HasSerials() bool {
	if o != nil && !IsNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *ReassignNetworkVlanProfilesAssignments200Response) SetSerials(v []string) {
	o.Serials = v
}

// GetStackIds returns the StackIds field value if set, zero value otherwise.
func (o *ReassignNetworkVlanProfilesAssignments200Response) GetStackIds() []string {
	if o == nil || IsNil(o.StackIds) {
		var ret []string
		return ret
	}
	return o.StackIds
}

// GetStackIdsOk returns a tuple with the StackIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReassignNetworkVlanProfilesAssignments200Response) GetStackIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.StackIds) {
		return nil, false
	}
	return o.StackIds, true
}

// HasStackIds returns a boolean if a field has been set.
func (o *ReassignNetworkVlanProfilesAssignments200Response) HasStackIds() bool {
	if o != nil && !IsNil(o.StackIds) {
		return true
	}

	return false
}

// SetStackIds gets a reference to the given []string and assigns it to the StackIds field.
func (o *ReassignNetworkVlanProfilesAssignments200Response) SetStackIds(v []string) {
	o.StackIds = v
}

func (o ReassignNetworkVlanProfilesAssignments200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReassignNetworkVlanProfilesAssignments200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VlanProfile) {
		toSerialize["vlanProfile"] = o.VlanProfile
	}
	if !IsNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	if !IsNil(o.StackIds) {
		toSerialize["stackIds"] = o.StackIds
	}
	return toSerialize, nil
}

type NullableReassignNetworkVlanProfilesAssignments200Response struct {
	value *ReassignNetworkVlanProfilesAssignments200Response
	isSet bool
}

func (v NullableReassignNetworkVlanProfilesAssignments200Response) Get() *ReassignNetworkVlanProfilesAssignments200Response {
	return v.value
}

func (v *NullableReassignNetworkVlanProfilesAssignments200Response) Set(val *ReassignNetworkVlanProfilesAssignments200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReassignNetworkVlanProfilesAssignments200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReassignNetworkVlanProfilesAssignments200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReassignNetworkVlanProfilesAssignments200Response(val *ReassignNetworkVlanProfilesAssignments200Response) *NullableReassignNetworkVlanProfilesAssignments200Response {
	return &NullableReassignNetworkVlanProfilesAssignments200Response{value: val, isSet: true}
}

func (v NullableReassignNetworkVlanProfilesAssignments200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReassignNetworkVlanProfilesAssignments200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


