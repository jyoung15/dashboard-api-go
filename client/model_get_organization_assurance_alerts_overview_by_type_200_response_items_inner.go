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

// checks if the GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner{}

// GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner struct for GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner
type GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner struct {
	// Alert Type
	Type string `json:"type"`
	// Total count of the given alert type
	Count int32 `json:"count"`
}

type _GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner

// NewGetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner instantiates a new GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner(type_ string, count int32) *GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner {
	this := GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner{}
	this.Type = type_
	this.Count = count
	return &this
}

// NewGetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInnerWithDefaults instantiates a new GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInnerWithDefaults() *GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner {
	this := GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner{}
	return &this
}

// GetType returns the Type field value
func (o *GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner) SetType(v string) {
	o.Type = v
}

// GetCount returns the Count field value
func (o *GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner) SetCount(v int32) {
	o.Count = v
}

func (o GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["count"] = o.Count
	return toSerialize, nil
}

func (o *GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"count",
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

	varGetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner := _GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner)

	if err != nil {
		return err
	}

	*o = GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner(varGetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner)

	return err
}

type NullableGetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner struct {
	value *GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner
	isSet bool
}

func (v NullableGetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner) Get() *GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner {
	return v.value
}

func (v *NullableGetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner) Set(val *GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner(val *GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner) *NullableGetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner {
	return &NullableGetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner{value: val, isSet: true}
}

func (v NullableGetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


