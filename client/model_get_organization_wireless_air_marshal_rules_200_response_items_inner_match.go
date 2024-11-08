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

// checks if the GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch{}

// GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch Indicates whether or not clients are allowed to        connect to rogue SSIDs by default. (blocked by default)
type GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch struct {
	// Indicates whether or not clients are allowed to        connect to rogue SSIDs by default. (blocked by default)
	String *string `json:"string,omitempty"`
	// Indicates whether or not clients are allowed to        connect to rogue SSIDs by default. (blocked by default)
	Type *string `json:"type,omitempty"`
}

// NewGetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch instantiates a new GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch() *GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch {
	this := GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch{}
	return &this
}

// NewGetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatchWithDefaults instantiates a new GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatchWithDefaults() *GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch {
	this := GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch{}
	return &this
}

// GetString returns the String field value if set, zero value otherwise.
func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch) GetString() string {
	if o == nil || IsNil(o.String) {
		var ret string
		return ret
	}
	return *o.String
}

// GetStringOk returns a tuple with the String field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch) GetStringOk() (*string, bool) {
	if o == nil || IsNil(o.String) {
		return nil, false
	}
	return o.String, true
}

// HasString returns a boolean if a field has been set.
func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch) HasString() bool {
	if o != nil && !IsNil(o.String) {
		return true
	}

	return false
}

// SetString gets a reference to the given string and assigns it to the String field.
func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch) SetString(v string) {
	o.String = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch) SetType(v string) {
	o.Type = &v
}

func (o GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.String) {
		toSerialize["string"] = o.String
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableGetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch struct {
	value *GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch
	isSet bool
}

func (v NullableGetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch) Get() *GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch {
	return v.value
}

func (v *NullableGetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch) Set(val *GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch(val *GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch) *NullableGetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch {
	return &NullableGetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


