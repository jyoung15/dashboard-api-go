/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ReassignNetworkVlanProfilesAssignmentsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReassignNetworkVlanProfilesAssignmentsRequest{}

// ReassignNetworkVlanProfilesAssignmentsRequest struct for ReassignNetworkVlanProfilesAssignmentsRequest
type ReassignNetworkVlanProfilesAssignmentsRequest struct {
	VlanProfile *ReassignNetworkVlanProfilesAssignmentsRequestVlanProfile `json:"vlanProfile,omitempty"`
	// Array of Device Serials
	Serials []string `json:"serials"`
	// Array of Switch Stack IDs
	StackIds []string `json:"stackIds"`
}

type _ReassignNetworkVlanProfilesAssignmentsRequest ReassignNetworkVlanProfilesAssignmentsRequest

// NewReassignNetworkVlanProfilesAssignmentsRequest instantiates a new ReassignNetworkVlanProfilesAssignmentsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReassignNetworkVlanProfilesAssignmentsRequest(serials []string, stackIds []string) *ReassignNetworkVlanProfilesAssignmentsRequest {
	this := ReassignNetworkVlanProfilesAssignmentsRequest{}
	this.Serials = serials
	this.StackIds = stackIds
	return &this
}

// NewReassignNetworkVlanProfilesAssignmentsRequestWithDefaults instantiates a new ReassignNetworkVlanProfilesAssignmentsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReassignNetworkVlanProfilesAssignmentsRequestWithDefaults() *ReassignNetworkVlanProfilesAssignmentsRequest {
	this := ReassignNetworkVlanProfilesAssignmentsRequest{}
	return &this
}

// GetVlanProfile returns the VlanProfile field value if set, zero value otherwise.
func (o *ReassignNetworkVlanProfilesAssignmentsRequest) GetVlanProfile() ReassignNetworkVlanProfilesAssignmentsRequestVlanProfile {
	if o == nil || IsNil(o.VlanProfile) {
		var ret ReassignNetworkVlanProfilesAssignmentsRequestVlanProfile
		return ret
	}
	return *o.VlanProfile
}

// GetVlanProfileOk returns a tuple with the VlanProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReassignNetworkVlanProfilesAssignmentsRequest) GetVlanProfileOk() (*ReassignNetworkVlanProfilesAssignmentsRequestVlanProfile, bool) {
	if o == nil || IsNil(o.VlanProfile) {
		return nil, false
	}
	return o.VlanProfile, true
}

// HasVlanProfile returns a boolean if a field has been set.
func (o *ReassignNetworkVlanProfilesAssignmentsRequest) HasVlanProfile() bool {
	if o != nil && !IsNil(o.VlanProfile) {
		return true
	}

	return false
}

// SetVlanProfile gets a reference to the given ReassignNetworkVlanProfilesAssignmentsRequestVlanProfile and assigns it to the VlanProfile field.
func (o *ReassignNetworkVlanProfilesAssignmentsRequest) SetVlanProfile(v ReassignNetworkVlanProfilesAssignmentsRequestVlanProfile) {
	o.VlanProfile = &v
}

// GetSerials returns the Serials field value
func (o *ReassignNetworkVlanProfilesAssignmentsRequest) GetSerials() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value
// and a boolean to check if the value has been set.
func (o *ReassignNetworkVlanProfilesAssignmentsRequest) GetSerialsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Serials, true
}

// SetSerials sets field value
func (o *ReassignNetworkVlanProfilesAssignmentsRequest) SetSerials(v []string) {
	o.Serials = v
}

// GetStackIds returns the StackIds field value
func (o *ReassignNetworkVlanProfilesAssignmentsRequest) GetStackIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.StackIds
}

// GetStackIdsOk returns a tuple with the StackIds field value
// and a boolean to check if the value has been set.
func (o *ReassignNetworkVlanProfilesAssignmentsRequest) GetStackIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StackIds, true
}

// SetStackIds sets field value
func (o *ReassignNetworkVlanProfilesAssignmentsRequest) SetStackIds(v []string) {
	o.StackIds = v
}

func (o ReassignNetworkVlanProfilesAssignmentsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReassignNetworkVlanProfilesAssignmentsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VlanProfile) {
		toSerialize["vlanProfile"] = o.VlanProfile
	}
	toSerialize["serials"] = o.Serials
	toSerialize["stackIds"] = o.StackIds
	return toSerialize, nil
}

func (o *ReassignNetworkVlanProfilesAssignmentsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"serials",
		"stackIds",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varReassignNetworkVlanProfilesAssignmentsRequest := _ReassignNetworkVlanProfilesAssignmentsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReassignNetworkVlanProfilesAssignmentsRequest)

	if err != nil {
		return err
	}

	*o = ReassignNetworkVlanProfilesAssignmentsRequest(varReassignNetworkVlanProfilesAssignmentsRequest)

	return err
}

type NullableReassignNetworkVlanProfilesAssignmentsRequest struct {
	value *ReassignNetworkVlanProfilesAssignmentsRequest
	isSet bool
}

func (v NullableReassignNetworkVlanProfilesAssignmentsRequest) Get() *ReassignNetworkVlanProfilesAssignmentsRequest {
	return v.value
}

func (v *NullableReassignNetworkVlanProfilesAssignmentsRequest) Set(val *ReassignNetworkVlanProfilesAssignmentsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReassignNetworkVlanProfilesAssignmentsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReassignNetworkVlanProfilesAssignmentsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReassignNetworkVlanProfilesAssignmentsRequest(val *ReassignNetworkVlanProfilesAssignmentsRequest) *NullableReassignNetworkVlanProfilesAssignmentsRequest {
	return &NullableReassignNetworkVlanProfilesAssignmentsRequest{value: val, isSet: true}
}

func (v NullableReassignNetworkVlanProfilesAssignmentsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReassignNetworkVlanProfilesAssignmentsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


