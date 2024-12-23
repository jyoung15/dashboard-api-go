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

// checks if the GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems{}

// GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems The count metadata.
type GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems struct {
	// The total number of serials.
	Total *int32 `json:"total,omitempty"`
	// The number of serials remaining based on current pagination location within the dataset.
	Remaining *int32 `json:"remaining,omitempty"`
}

// NewGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems instantiates a new GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems() *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems {
	this := GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems{}
	return &this
}

// NewGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItemsWithDefaults instantiates a new GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItemsWithDefaults() *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems {
	this := GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems) SetTotal(v int32) {
	o.Total = &v
}

// GetRemaining returns the Remaining field value if set, zero value otherwise.
func (o *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems) GetRemaining() int32 {
	if o == nil || IsNil(o.Remaining) {
		var ret int32
		return ret
	}
	return *o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems) GetRemainingOk() (*int32, bool) {
	if o == nil || IsNil(o.Remaining) {
		return nil, false
	}
	return o.Remaining, true
}

// HasRemaining returns a boolean if a field has been set.
func (o *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems) HasRemaining() bool {
	if o != nil && !IsNil(o.Remaining) {
		return true
	}

	return false
}

// SetRemaining gets a reference to the given int32 and assigns it to the Remaining field.
func (o *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems) SetRemaining(v int32) {
	o.Remaining = &v
}

func (o GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Remaining) {
		toSerialize["remaining"] = o.Remaining
	}
	return toSerialize, nil
}

type NullableGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems struct {
	value *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems
	isSet bool
}

func (v NullableGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems) Get() *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems {
	return v.value
}

func (v *NullableGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems) Set(val *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems(val *GetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems) *NullableGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems {
	return &NullableGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessRfProfilesAssignmentsByDevice200ResponseInnerMetaCountsItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


