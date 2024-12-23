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

// checks if the GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts{}

// GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts Counts
type GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts struct {
	// Total Alerts
	Items int32 `json:"items"`
}

type _GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts

// NewGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts instantiates a new GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts(items int32) *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts {
	this := GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts{}
	this.Items = items
	return &this
}

// NewGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCountsWithDefaults instantiates a new GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCountsWithDefaults() *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts {
	this := GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts{}
	return &this
}

// GetItems returns the Items field value
func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts) GetItems() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts) GetItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts) SetItems(v int32) {
	o.Items = v
}

func (o GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts) UnmarshalJSON(data []byte) (err error) {
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

	varGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts := _GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts)

	if err != nil {
		return err
	}

	*o = GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts(varGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts)

	return err
}

type NullableGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts struct {
	value *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts
	isSet bool
}

func (v NullableGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts) Get() *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts {
	return v.value
}

func (v *NullableGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts) Set(val *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts(val *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts) *NullableGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts {
	return &NullableGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts{value: val, isSet: true}
}

func (v NullableGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMetaCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


