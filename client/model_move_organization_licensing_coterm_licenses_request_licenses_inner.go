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

// checks if the MoveOrganizationLicensingCotermLicensesRequestLicensesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MoveOrganizationLicensingCotermLicensesRequestLicensesInner{}

// MoveOrganizationLicensingCotermLicensesRequestLicensesInner struct for MoveOrganizationLicensingCotermLicensesRequestLicensesInner
type MoveOrganizationLicensingCotermLicensesRequestLicensesInner struct {
	// The license key to move counts from
	Key string `json:"key"`
	// The counts to move from the license by model type
	Counts []MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner `json:"counts"`
}

type _MoveOrganizationLicensingCotermLicensesRequestLicensesInner MoveOrganizationLicensingCotermLicensesRequestLicensesInner

// NewMoveOrganizationLicensingCotermLicensesRequestLicensesInner instantiates a new MoveOrganizationLicensingCotermLicensesRequestLicensesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveOrganizationLicensingCotermLicensesRequestLicensesInner(key string, counts []MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner) *MoveOrganizationLicensingCotermLicensesRequestLicensesInner {
	this := MoveOrganizationLicensingCotermLicensesRequestLicensesInner{}
	this.Key = key
	this.Counts = counts
	return &this
}

// NewMoveOrganizationLicensingCotermLicensesRequestLicensesInnerWithDefaults instantiates a new MoveOrganizationLicensingCotermLicensesRequestLicensesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveOrganizationLicensingCotermLicensesRequestLicensesInnerWithDefaults() *MoveOrganizationLicensingCotermLicensesRequestLicensesInner {
	this := MoveOrganizationLicensingCotermLicensesRequestLicensesInner{}
	return &this
}

// GetKey returns the Key field value
func (o *MoveOrganizationLicensingCotermLicensesRequestLicensesInner) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *MoveOrganizationLicensingCotermLicensesRequestLicensesInner) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *MoveOrganizationLicensingCotermLicensesRequestLicensesInner) SetKey(v string) {
	o.Key = v
}

// GetCounts returns the Counts field value
func (o *MoveOrganizationLicensingCotermLicensesRequestLicensesInner) GetCounts() []MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner {
	if o == nil {
		var ret []MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner
		return ret
	}

	return o.Counts
}

// GetCountsOk returns a tuple with the Counts field value
// and a boolean to check if the value has been set.
func (o *MoveOrganizationLicensingCotermLicensesRequestLicensesInner) GetCountsOk() ([]MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Counts, true
}

// SetCounts sets field value
func (o *MoveOrganizationLicensingCotermLicensesRequestLicensesInner) SetCounts(v []MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner) {
	o.Counts = v
}

func (o MoveOrganizationLicensingCotermLicensesRequestLicensesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MoveOrganizationLicensingCotermLicensesRequestLicensesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["counts"] = o.Counts
	return toSerialize, nil
}

func (o *MoveOrganizationLicensingCotermLicensesRequestLicensesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"counts",
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

	varMoveOrganizationLicensingCotermLicensesRequestLicensesInner := _MoveOrganizationLicensingCotermLicensesRequestLicensesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMoveOrganizationLicensingCotermLicensesRequestLicensesInner)

	if err != nil {
		return err
	}

	*o = MoveOrganizationLicensingCotermLicensesRequestLicensesInner(varMoveOrganizationLicensingCotermLicensesRequestLicensesInner)

	return err
}

type NullableMoveOrganizationLicensingCotermLicensesRequestLicensesInner struct {
	value *MoveOrganizationLicensingCotermLicensesRequestLicensesInner
	isSet bool
}

func (v NullableMoveOrganizationLicensingCotermLicensesRequestLicensesInner) Get() *MoveOrganizationLicensingCotermLicensesRequestLicensesInner {
	return v.value
}

func (v *NullableMoveOrganizationLicensingCotermLicensesRequestLicensesInner) Set(val *MoveOrganizationLicensingCotermLicensesRequestLicensesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveOrganizationLicensingCotermLicensesRequestLicensesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveOrganizationLicensingCotermLicensesRequestLicensesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveOrganizationLicensingCotermLicensesRequestLicensesInner(val *MoveOrganizationLicensingCotermLicensesRequestLicensesInner) *NullableMoveOrganizationLicensingCotermLicensesRequestLicensesInner {
	return &NullableMoveOrganizationLicensingCotermLicensesRequestLicensesInner{value: val, isSet: true}
}

func (v NullableMoveOrganizationLicensingCotermLicensesRequestLicensesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveOrganizationLicensingCotermLicensesRequestLicensesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


