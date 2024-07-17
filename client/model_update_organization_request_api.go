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

// checks if the UpdateOrganizationRequestApi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationRequestApi{}

// UpdateOrganizationRequestApi API-specific settings
type UpdateOrganizationRequestApi struct {
	// If true, enable the access to the Cisco Meraki Dashboard API
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUpdateOrganizationRequestApi instantiates a new UpdateOrganizationRequestApi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationRequestApi() *UpdateOrganizationRequestApi {
	this := UpdateOrganizationRequestApi{}
	return &this
}

// NewUpdateOrganizationRequestApiWithDefaults instantiates a new UpdateOrganizationRequestApi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationRequestApiWithDefaults() *UpdateOrganizationRequestApi {
	this := UpdateOrganizationRequestApi{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateOrganizationRequestApi) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationRequestApi) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateOrganizationRequestApi) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateOrganizationRequestApi) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UpdateOrganizationRequestApi) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationRequestApi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableUpdateOrganizationRequestApi struct {
	value *UpdateOrganizationRequestApi
	isSet bool
}

func (v NullableUpdateOrganizationRequestApi) Get() *UpdateOrganizationRequestApi {
	return v.value
}

func (v *NullableUpdateOrganizationRequestApi) Set(val *UpdateOrganizationRequestApi) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationRequestApi) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationRequestApi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationRequestApi(val *UpdateOrganizationRequestApi) *NullableUpdateOrganizationRequestApi {
	return &NullableUpdateOrganizationRequestApi{value: val, isSet: true}
}

func (v NullableUpdateOrganizationRequestApi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationRequestApi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


