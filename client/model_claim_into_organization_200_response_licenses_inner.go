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

// checks if the ClaimIntoOrganization200ResponseLicensesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClaimIntoOrganization200ResponseLicensesInner{}

// ClaimIntoOrganization200ResponseLicensesInner struct for ClaimIntoOrganization200ResponseLicensesInner
type ClaimIntoOrganization200ResponseLicensesInner struct {
	// The key of the license
	Key *string `json:"key,omitempty"`
	// The mode of the license
	Mode *string `json:"mode,omitempty"`
}

// NewClaimIntoOrganization200ResponseLicensesInner instantiates a new ClaimIntoOrganization200ResponseLicensesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClaimIntoOrganization200ResponseLicensesInner() *ClaimIntoOrganization200ResponseLicensesInner {
	this := ClaimIntoOrganization200ResponseLicensesInner{}
	return &this
}

// NewClaimIntoOrganization200ResponseLicensesInnerWithDefaults instantiates a new ClaimIntoOrganization200ResponseLicensesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClaimIntoOrganization200ResponseLicensesInnerWithDefaults() *ClaimIntoOrganization200ResponseLicensesInner {
	this := ClaimIntoOrganization200ResponseLicensesInner{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ClaimIntoOrganization200ResponseLicensesInner) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimIntoOrganization200ResponseLicensesInner) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ClaimIntoOrganization200ResponseLicensesInner) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ClaimIntoOrganization200ResponseLicensesInner) SetKey(v string) {
	o.Key = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ClaimIntoOrganization200ResponseLicensesInner) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimIntoOrganization200ResponseLicensesInner) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ClaimIntoOrganization200ResponseLicensesInner) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *ClaimIntoOrganization200ResponseLicensesInner) SetMode(v string) {
	o.Mode = &v
}

func (o ClaimIntoOrganization200ResponseLicensesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClaimIntoOrganization200ResponseLicensesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return toSerialize, nil
}

type NullableClaimIntoOrganization200ResponseLicensesInner struct {
	value *ClaimIntoOrganization200ResponseLicensesInner
	isSet bool
}

func (v NullableClaimIntoOrganization200ResponseLicensesInner) Get() *ClaimIntoOrganization200ResponseLicensesInner {
	return v.value
}

func (v *NullableClaimIntoOrganization200ResponseLicensesInner) Set(val *ClaimIntoOrganization200ResponseLicensesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimIntoOrganization200ResponseLicensesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimIntoOrganization200ResponseLicensesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimIntoOrganization200ResponseLicensesInner(val *ClaimIntoOrganization200ResponseLicensesInner) *NullableClaimIntoOrganization200ResponseLicensesInner {
	return &NullableClaimIntoOrganization200ResponseLicensesInner{value: val, isSet: true}
}

func (v NullableClaimIntoOrganization200ResponseLicensesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaimIntoOrganization200ResponseLicensesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


