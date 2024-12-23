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

// checks if the GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts{}

// GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts Counts
type GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts struct {
	// Total Segments
	Items int32 `json:"items"`
}

type _GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts

// NewGetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts instantiates a new GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts(items int32) *GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts {
	this := GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts{}
	this.Items = items
	return &this
}

// NewGetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCountsWithDefaults instantiates a new GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCountsWithDefaults() *GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts {
	this := GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts{}
	return &this
}

// GetItems returns the Items field value
func (o *GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts) GetItems() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts) GetItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts) SetItems(v int32) {
	o.Items = v
}

func (o GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
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

	varGetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts := _GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts)

	if err != nil {
		return err
	}

	*o = GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts(varGetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts)

	return err
}

type NullableGetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts struct {
	value *GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts
	isSet bool
}

func (v NullableGetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts) Get() *GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts {
	return v.value
}

func (v *NullableGetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts) Set(val *GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts(val *GetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts) *NullableGetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts {
	return &NullableGetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts{value: val, isSet: true}
}

func (v NullableGetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationAssuranceAlertsOverviewHistorical200ResponseMetaCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


