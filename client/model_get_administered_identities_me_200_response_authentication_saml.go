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

// checks if the GetAdministeredIdentitiesMe200ResponseAuthenticationSaml type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAdministeredIdentitiesMe200ResponseAuthenticationSaml{}

// GetAdministeredIdentitiesMe200ResponseAuthenticationSaml SAML authentication
type GetAdministeredIdentitiesMe200ResponseAuthenticationSaml struct {
	// If SAML authentication is enabled for this user
	Enabled *bool `json:"enabled,omitempty"`
}

// NewGetAdministeredIdentitiesMe200ResponseAuthenticationSaml instantiates a new GetAdministeredIdentitiesMe200ResponseAuthenticationSaml object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAdministeredIdentitiesMe200ResponseAuthenticationSaml() *GetAdministeredIdentitiesMe200ResponseAuthenticationSaml {
	this := GetAdministeredIdentitiesMe200ResponseAuthenticationSaml{}
	return &this
}

// NewGetAdministeredIdentitiesMe200ResponseAuthenticationSamlWithDefaults instantiates a new GetAdministeredIdentitiesMe200ResponseAuthenticationSaml object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAdministeredIdentitiesMe200ResponseAuthenticationSamlWithDefaults() *GetAdministeredIdentitiesMe200ResponseAuthenticationSaml {
	this := GetAdministeredIdentitiesMe200ResponseAuthenticationSaml{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetAdministeredIdentitiesMe200ResponseAuthenticationSaml) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredIdentitiesMe200ResponseAuthenticationSaml) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetAdministeredIdentitiesMe200ResponseAuthenticationSaml) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetAdministeredIdentitiesMe200ResponseAuthenticationSaml) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o GetAdministeredIdentitiesMe200ResponseAuthenticationSaml) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAdministeredIdentitiesMe200ResponseAuthenticationSaml) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableGetAdministeredIdentitiesMe200ResponseAuthenticationSaml struct {
	value *GetAdministeredIdentitiesMe200ResponseAuthenticationSaml
	isSet bool
}

func (v NullableGetAdministeredIdentitiesMe200ResponseAuthenticationSaml) Get() *GetAdministeredIdentitiesMe200ResponseAuthenticationSaml {
	return v.value
}

func (v *NullableGetAdministeredIdentitiesMe200ResponseAuthenticationSaml) Set(val *GetAdministeredIdentitiesMe200ResponseAuthenticationSaml) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAdministeredIdentitiesMe200ResponseAuthenticationSaml) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAdministeredIdentitiesMe200ResponseAuthenticationSaml) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAdministeredIdentitiesMe200ResponseAuthenticationSaml(val *GetAdministeredIdentitiesMe200ResponseAuthenticationSaml) *NullableGetAdministeredIdentitiesMe200ResponseAuthenticationSaml {
	return &NullableGetAdministeredIdentitiesMe200ResponseAuthenticationSaml{value: val, isSet: true}
}

func (v NullableGetAdministeredIdentitiesMe200ResponseAuthenticationSaml) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAdministeredIdentitiesMe200ResponseAuthenticationSaml) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


