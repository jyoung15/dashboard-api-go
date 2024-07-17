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

// checks if the GetOrganizationLicensesOverview200ResponseStatesUnused type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationLicensesOverview200ResponseStatesUnused{}

// GetOrganizationLicensesOverview200ResponseStatesUnused Data for unused licenses
type GetOrganizationLicensesOverview200ResponseStatesUnused struct {
	// The number of unused licenses
	Count *int32 `json:"count,omitempty"`
	SoonestActivation *GetOrganizationLicensesOverview200ResponseStatesUnusedSoonestActivation `json:"soonestActivation,omitempty"`
}

// NewGetOrganizationLicensesOverview200ResponseStatesUnused instantiates a new GetOrganizationLicensesOverview200ResponseStatesUnused object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationLicensesOverview200ResponseStatesUnused() *GetOrganizationLicensesOverview200ResponseStatesUnused {
	this := GetOrganizationLicensesOverview200ResponseStatesUnused{}
	return &this
}

// NewGetOrganizationLicensesOverview200ResponseStatesUnusedWithDefaults instantiates a new GetOrganizationLicensesOverview200ResponseStatesUnused object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationLicensesOverview200ResponseStatesUnusedWithDefaults() *GetOrganizationLicensesOverview200ResponseStatesUnused {
	this := GetOrganizationLicensesOverview200ResponseStatesUnused{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GetOrganizationLicensesOverview200ResponseStatesUnused) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicensesOverview200ResponseStatesUnused) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GetOrganizationLicensesOverview200ResponseStatesUnused) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GetOrganizationLicensesOverview200ResponseStatesUnused) SetCount(v int32) {
	o.Count = &v
}

// GetSoonestActivation returns the SoonestActivation field value if set, zero value otherwise.
func (o *GetOrganizationLicensesOverview200ResponseStatesUnused) GetSoonestActivation() GetOrganizationLicensesOverview200ResponseStatesUnusedSoonestActivation {
	if o == nil || IsNil(o.SoonestActivation) {
		var ret GetOrganizationLicensesOverview200ResponseStatesUnusedSoonestActivation
		return ret
	}
	return *o.SoonestActivation
}

// GetSoonestActivationOk returns a tuple with the SoonestActivation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicensesOverview200ResponseStatesUnused) GetSoonestActivationOk() (*GetOrganizationLicensesOverview200ResponseStatesUnusedSoonestActivation, bool) {
	if o == nil || IsNil(o.SoonestActivation) {
		return nil, false
	}
	return o.SoonestActivation, true
}

// HasSoonestActivation returns a boolean if a field has been set.
func (o *GetOrganizationLicensesOverview200ResponseStatesUnused) HasSoonestActivation() bool {
	if o != nil && !IsNil(o.SoonestActivation) {
		return true
	}

	return false
}

// SetSoonestActivation gets a reference to the given GetOrganizationLicensesOverview200ResponseStatesUnusedSoonestActivation and assigns it to the SoonestActivation field.
func (o *GetOrganizationLicensesOverview200ResponseStatesUnused) SetSoonestActivation(v GetOrganizationLicensesOverview200ResponseStatesUnusedSoonestActivation) {
	o.SoonestActivation = &v
}

func (o GetOrganizationLicensesOverview200ResponseStatesUnused) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationLicensesOverview200ResponseStatesUnused) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.SoonestActivation) {
		toSerialize["soonestActivation"] = o.SoonestActivation
	}
	return toSerialize, nil
}

type NullableGetOrganizationLicensesOverview200ResponseStatesUnused struct {
	value *GetOrganizationLicensesOverview200ResponseStatesUnused
	isSet bool
}

func (v NullableGetOrganizationLicensesOverview200ResponseStatesUnused) Get() *GetOrganizationLicensesOverview200ResponseStatesUnused {
	return v.value
}

func (v *NullableGetOrganizationLicensesOverview200ResponseStatesUnused) Set(val *GetOrganizationLicensesOverview200ResponseStatesUnused) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationLicensesOverview200ResponseStatesUnused) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationLicensesOverview200ResponseStatesUnused) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationLicensesOverview200ResponseStatesUnused(val *GetOrganizationLicensesOverview200ResponseStatesUnused) *NullableGetOrganizationLicensesOverview200ResponseStatesUnused {
	return &NullableGetOrganizationLicensesOverview200ResponseStatesUnused{value: val, isSet: true}
}

func (v NullableGetOrganizationLicensesOverview200ResponseStatesUnused) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationLicensesOverview200ResponseStatesUnused) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


