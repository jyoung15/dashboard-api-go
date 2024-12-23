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

// checks if the ReleaseFromOrganizationInventoryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReleaseFromOrganizationInventoryRequest{}

// ReleaseFromOrganizationInventoryRequest struct for ReleaseFromOrganizationInventoryRequest
type ReleaseFromOrganizationInventoryRequest struct {
	// Serials of the devices that should be released
	Serials []string `json:"serials,omitempty"`
}

// NewReleaseFromOrganizationInventoryRequest instantiates a new ReleaseFromOrganizationInventoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseFromOrganizationInventoryRequest() *ReleaseFromOrganizationInventoryRequest {
	this := ReleaseFromOrganizationInventoryRequest{}
	return &this
}

// NewReleaseFromOrganizationInventoryRequestWithDefaults instantiates a new ReleaseFromOrganizationInventoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseFromOrganizationInventoryRequestWithDefaults() *ReleaseFromOrganizationInventoryRequest {
	this := ReleaseFromOrganizationInventoryRequest{}
	return &this
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *ReleaseFromOrganizationInventoryRequest) GetSerials() []string {
	if o == nil || IsNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseFromOrganizationInventoryRequest) GetSerialsOk() ([]string, bool) {
	if o == nil || IsNil(o.Serials) {
		return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *ReleaseFromOrganizationInventoryRequest) HasSerials() bool {
	if o != nil && !IsNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *ReleaseFromOrganizationInventoryRequest) SetSerials(v []string) {
	o.Serials = v
}

func (o ReleaseFromOrganizationInventoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReleaseFromOrganizationInventoryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	return toSerialize, nil
}

type NullableReleaseFromOrganizationInventoryRequest struct {
	value *ReleaseFromOrganizationInventoryRequest
	isSet bool
}

func (v NullableReleaseFromOrganizationInventoryRequest) Get() *ReleaseFromOrganizationInventoryRequest {
	return v.value
}

func (v *NullableReleaseFromOrganizationInventoryRequest) Set(val *ReleaseFromOrganizationInventoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseFromOrganizationInventoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseFromOrganizationInventoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseFromOrganizationInventoryRequest(val *ReleaseFromOrganizationInventoryRequest) *NullableReleaseFromOrganizationInventoryRequest {
	return &NullableReleaseFromOrganizationInventoryRequest{value: val, isSet: true}
}

func (v NullableReleaseFromOrganizationInventoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseFromOrganizationInventoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


