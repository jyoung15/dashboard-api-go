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

// checks if the GetOrganizationSmAdminsRoles200ResponseMetaCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSmAdminsRoles200ResponseMetaCounts{}

// GetOrganizationSmAdminsRoles200ResponseMetaCounts Counts relating to the paginated dataset
type GetOrganizationSmAdminsRoles200ResponseMetaCounts struct {
	Items *GetOrganizationSmAdminsRoles200ResponseMetaCountsItems `json:"items,omitempty"`
}

// NewGetOrganizationSmAdminsRoles200ResponseMetaCounts instantiates a new GetOrganizationSmAdminsRoles200ResponseMetaCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSmAdminsRoles200ResponseMetaCounts() *GetOrganizationSmAdminsRoles200ResponseMetaCounts {
	this := GetOrganizationSmAdminsRoles200ResponseMetaCounts{}
	return &this
}

// NewGetOrganizationSmAdminsRoles200ResponseMetaCountsWithDefaults instantiates a new GetOrganizationSmAdminsRoles200ResponseMetaCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSmAdminsRoles200ResponseMetaCountsWithDefaults() *GetOrganizationSmAdminsRoles200ResponseMetaCounts {
	this := GetOrganizationSmAdminsRoles200ResponseMetaCounts{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GetOrganizationSmAdminsRoles200ResponseMetaCounts) GetItems() GetOrganizationSmAdminsRoles200ResponseMetaCountsItems {
	if o == nil || IsNil(o.Items) {
		var ret GetOrganizationSmAdminsRoles200ResponseMetaCountsItems
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmAdminsRoles200ResponseMetaCounts) GetItemsOk() (*GetOrganizationSmAdminsRoles200ResponseMetaCountsItems, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GetOrganizationSmAdminsRoles200ResponseMetaCounts) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given GetOrganizationSmAdminsRoles200ResponseMetaCountsItems and assigns it to the Items field.
func (o *GetOrganizationSmAdminsRoles200ResponseMetaCounts) SetItems(v GetOrganizationSmAdminsRoles200ResponseMetaCountsItems) {
	o.Items = &v
}

func (o GetOrganizationSmAdminsRoles200ResponseMetaCounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSmAdminsRoles200ResponseMetaCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableGetOrganizationSmAdminsRoles200ResponseMetaCounts struct {
	value *GetOrganizationSmAdminsRoles200ResponseMetaCounts
	isSet bool
}

func (v NullableGetOrganizationSmAdminsRoles200ResponseMetaCounts) Get() *GetOrganizationSmAdminsRoles200ResponseMetaCounts {
	return v.value
}

func (v *NullableGetOrganizationSmAdminsRoles200ResponseMetaCounts) Set(val *GetOrganizationSmAdminsRoles200ResponseMetaCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSmAdminsRoles200ResponseMetaCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSmAdminsRoles200ResponseMetaCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSmAdminsRoles200ResponseMetaCounts(val *GetOrganizationSmAdminsRoles200ResponseMetaCounts) *NullableGetOrganizationSmAdminsRoles200ResponseMetaCounts {
	return &NullableGetOrganizationSmAdminsRoles200ResponseMetaCounts{value: val, isSet: true}
}

func (v NullableGetOrganizationSmAdminsRoles200ResponseMetaCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSmAdminsRoles200ResponseMetaCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


