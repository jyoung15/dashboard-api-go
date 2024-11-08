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

// checks if the GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta{}

// GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta Other metadata related to this result set.
type GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta struct {
	Counts *GetOrganizationWirelessSsidsStatusesByDevice200ResponseMetaCounts `json:"counts,omitempty"`
}

// NewGetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta instantiates a new GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta() *GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta {
	this := GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta{}
	return &this
}

// NewGetOrganizationWirelessSsidsStatusesByDevice200ResponseMetaWithDefaults instantiates a new GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessSsidsStatusesByDevice200ResponseMetaWithDefaults() *GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta {
	this := GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta) GetCounts() GetOrganizationWirelessSsidsStatusesByDevice200ResponseMetaCounts {
	if o == nil || IsNil(o.Counts) {
		var ret GetOrganizationWirelessSsidsStatusesByDevice200ResponseMetaCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta) GetCountsOk() (*GetOrganizationWirelessSsidsStatusesByDevice200ResponseMetaCounts, bool) {
	if o == nil || IsNil(o.Counts) {
		return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta) HasCounts() bool {
	if o != nil && !IsNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given GetOrganizationWirelessSsidsStatusesByDevice200ResponseMetaCounts and assigns it to the Counts field.
func (o *GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta) SetCounts(v GetOrganizationWirelessSsidsStatusesByDevice200ResponseMetaCounts) {
	o.Counts = &v
}

func (o GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return toSerialize, nil
}

type NullableGetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta struct {
	value *GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta
	isSet bool
}

func (v NullableGetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta) Get() *GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta {
	return v.value
}

func (v *NullableGetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta) Set(val *GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta(val *GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta) *NullableGetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta {
	return &NullableGetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


