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

// checks if the GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta{}

// GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta Other metadata related to this result set.
type GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta struct {
	Counts *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCounts `json:"counts,omitempty"`
}

// NewGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta instantiates a new GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta() *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta {
	this := GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta{}
	return &this
}

// NewGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaWithDefaults instantiates a new GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaWithDefaults() *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta {
	this := GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta) GetCounts() GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCounts {
	if o == nil || IsNil(o.Counts) {
		var ret GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta) GetCountsOk() (*GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCounts, bool) {
	if o == nil || IsNil(o.Counts) {
		return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta) HasCounts() bool {
	if o != nil && !IsNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCounts and assigns it to the Counts field.
func (o *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta) SetCounts(v GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCounts) {
	o.Counts = &v
}

func (o GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return toSerialize, nil
}

type NullableGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta struct {
	value *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta
	isSet bool
}

func (v NullableGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta) Get() *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta {
	return v.value
}

func (v *NullableGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta) Set(val *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta(val *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta) *NullableGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta {
	return &NullableGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


