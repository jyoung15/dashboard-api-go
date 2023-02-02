/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20026ProductsSwitch The Switch network to be updated
type InlineResponse20026ProductsSwitch struct {
	NextUpgrade *InlineResponse20026ProductsSwitchNextUpgrade `json:"nextUpgrade,omitempty"`
}

// NewInlineResponse20026ProductsSwitch instantiates a new InlineResponse20026ProductsSwitch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20026ProductsSwitch() *InlineResponse20026ProductsSwitch {
	this := InlineResponse20026ProductsSwitch{}
	return &this
}

// NewInlineResponse20026ProductsSwitchWithDefaults instantiates a new InlineResponse20026ProductsSwitch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20026ProductsSwitchWithDefaults() *InlineResponse20026ProductsSwitch {
	this := InlineResponse20026ProductsSwitch{}
	return &this
}

// GetNextUpgrade returns the NextUpgrade field value if set, zero value otherwise.
func (o *InlineResponse20026ProductsSwitch) GetNextUpgrade() InlineResponse20026ProductsSwitchNextUpgrade {
	if o == nil || isNil(o.NextUpgrade) {
		var ret InlineResponse20026ProductsSwitchNextUpgrade
		return ret
	}
	return *o.NextUpgrade
}

// GetNextUpgradeOk returns a tuple with the NextUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20026ProductsSwitch) GetNextUpgradeOk() (*InlineResponse20026ProductsSwitchNextUpgrade, bool) {
	if o == nil || isNil(o.NextUpgrade) {
    return nil, false
	}
	return o.NextUpgrade, true
}

// HasNextUpgrade returns a boolean if a field has been set.
func (o *InlineResponse20026ProductsSwitch) HasNextUpgrade() bool {
	if o != nil && !isNil(o.NextUpgrade) {
		return true
	}

	return false
}

// SetNextUpgrade gets a reference to the given InlineResponse20026ProductsSwitchNextUpgrade and assigns it to the NextUpgrade field.
func (o *InlineResponse20026ProductsSwitch) SetNextUpgrade(v InlineResponse20026ProductsSwitchNextUpgrade) {
	o.NextUpgrade = &v
}

func (o InlineResponse20026ProductsSwitch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NextUpgrade) {
		toSerialize["nextUpgrade"] = o.NextUpgrade
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20026ProductsSwitch struct {
	value *InlineResponse20026ProductsSwitch
	isSet bool
}

func (v NullableInlineResponse20026ProductsSwitch) Get() *InlineResponse20026ProductsSwitch {
	return v.value
}

func (v *NullableInlineResponse20026ProductsSwitch) Set(val *InlineResponse20026ProductsSwitch) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20026ProductsSwitch) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20026ProductsSwitch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20026ProductsSwitch(val *InlineResponse20026ProductsSwitch) *NullableInlineResponse20026ProductsSwitch {
	return &NullableInlineResponse20026ProductsSwitch{value: val, isSet: true}
}

func (v NullableInlineResponse20026ProductsSwitch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20026ProductsSwitch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

