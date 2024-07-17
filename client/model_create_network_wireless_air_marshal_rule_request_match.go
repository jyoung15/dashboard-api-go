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

// checks if the CreateNetworkWirelessAirMarshalRuleRequestMatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkWirelessAirMarshalRuleRequestMatch{}

// CreateNetworkWirelessAirMarshalRuleRequestMatch Object describing the rule specification.
type CreateNetworkWirelessAirMarshalRuleRequestMatch struct {
	// The type of match.
	Type *string `json:"type,omitempty"`
	// The string used to match.
	String *string `json:"string,omitempty"`
}

// NewCreateNetworkWirelessAirMarshalRuleRequestMatch instantiates a new CreateNetworkWirelessAirMarshalRuleRequestMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWirelessAirMarshalRuleRequestMatch() *CreateNetworkWirelessAirMarshalRuleRequestMatch {
	this := CreateNetworkWirelessAirMarshalRuleRequestMatch{}
	return &this
}

// NewCreateNetworkWirelessAirMarshalRuleRequestMatchWithDefaults instantiates a new CreateNetworkWirelessAirMarshalRuleRequestMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWirelessAirMarshalRuleRequestMatchWithDefaults() *CreateNetworkWirelessAirMarshalRuleRequestMatch {
	this := CreateNetworkWirelessAirMarshalRuleRequestMatch{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateNetworkWirelessAirMarshalRuleRequestMatch) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessAirMarshalRuleRequestMatch) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateNetworkWirelessAirMarshalRuleRequestMatch) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateNetworkWirelessAirMarshalRuleRequestMatch) SetType(v string) {
	o.Type = &v
}

// GetString returns the String field value if set, zero value otherwise.
func (o *CreateNetworkWirelessAirMarshalRuleRequestMatch) GetString() string {
	if o == nil || IsNil(o.String) {
		var ret string
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessAirMarshalRuleRequestMatch) GetStringOk() (*string, bool) {
	if o == nil || IsNil(o.String) {
		return nil, false
	}
	return o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *CreateNetworkWirelessAirMarshalRuleRequestMatch) HasString() bool {
	if o != nil && !IsNil(o.String) {
		return true
	}

	return false
}

// SetString gets a reference to the given string and assigns it to the String field.
func (o *CreateNetworkWirelessAirMarshalRuleRequestMatch) SetString(v string) {
	o.String = &v
}

func (o CreateNetworkWirelessAirMarshalRuleRequestMatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkWirelessAirMarshalRuleRequestMatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.String) {
		toSerialize["string"] = o.String
	}
	return toSerialize, nil
}

type NullableCreateNetworkWirelessAirMarshalRuleRequestMatch struct {
	value *CreateNetworkWirelessAirMarshalRuleRequestMatch
	isSet bool
}

func (v NullableCreateNetworkWirelessAirMarshalRuleRequestMatch) Get() *CreateNetworkWirelessAirMarshalRuleRequestMatch {
	return v.value
}

func (v *NullableCreateNetworkWirelessAirMarshalRuleRequestMatch) Set(val *CreateNetworkWirelessAirMarshalRuleRequestMatch) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWirelessAirMarshalRuleRequestMatch) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWirelessAirMarshalRuleRequestMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWirelessAirMarshalRuleRequestMatch(val *CreateNetworkWirelessAirMarshalRuleRequestMatch) *NullableCreateNetworkWirelessAirMarshalRuleRequestMatch {
	return &NullableCreateNetworkWirelessAirMarshalRuleRequestMatch{value: val, isSet: true}
}

func (v NullableCreateNetworkWirelessAirMarshalRuleRequestMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWirelessAirMarshalRuleRequestMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


