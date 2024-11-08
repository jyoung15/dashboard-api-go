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

// checks if the CreateNetworkWirelessRfProfileRequestFlexRadios type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkWirelessRfProfileRequestFlexRadios{}

// CreateNetworkWirelessRfProfileRequestFlexRadios Flex radio settings.
type CreateNetworkWirelessRfProfileRequestFlexRadios struct {
	// Flex radios by model.
	ByModel []CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner `json:"byModel,omitempty"`
}

// NewCreateNetworkWirelessRfProfileRequestFlexRadios instantiates a new CreateNetworkWirelessRfProfileRequestFlexRadios object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWirelessRfProfileRequestFlexRadios() *CreateNetworkWirelessRfProfileRequestFlexRadios {
	this := CreateNetworkWirelessRfProfileRequestFlexRadios{}
	return &this
}

// NewCreateNetworkWirelessRfProfileRequestFlexRadiosWithDefaults instantiates a new CreateNetworkWirelessRfProfileRequestFlexRadios object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWirelessRfProfileRequestFlexRadiosWithDefaults() *CreateNetworkWirelessRfProfileRequestFlexRadios {
	this := CreateNetworkWirelessRfProfileRequestFlexRadios{}
	return &this
}

// GetByModel returns the ByModel field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequestFlexRadios) GetByModel() []CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner {
	if o == nil || IsNil(o.ByModel) {
		var ret []CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner
		return ret
	}
	return o.ByModel
}

// GetByModelOk returns a tuple with the ByModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequestFlexRadios) GetByModelOk() ([]CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner, bool) {
	if o == nil || IsNil(o.ByModel) {
		return nil, false
	}
	return o.ByModel, true
}

// HasByModel returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequestFlexRadios) HasByModel() bool {
	if o != nil && !IsNil(o.ByModel) {
		return true
	}

	return false
}

// SetByModel gets a reference to the given []CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner and assigns it to the ByModel field.
func (o *CreateNetworkWirelessRfProfileRequestFlexRadios) SetByModel(v []CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) {
	o.ByModel = v
}

func (o CreateNetworkWirelessRfProfileRequestFlexRadios) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkWirelessRfProfileRequestFlexRadios) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ByModel) {
		toSerialize["byModel"] = o.ByModel
	}
	return toSerialize, nil
}

type NullableCreateNetworkWirelessRfProfileRequestFlexRadios struct {
	value *CreateNetworkWirelessRfProfileRequestFlexRadios
	isSet bool
}

func (v NullableCreateNetworkWirelessRfProfileRequestFlexRadios) Get() *CreateNetworkWirelessRfProfileRequestFlexRadios {
	return v.value
}

func (v *NullableCreateNetworkWirelessRfProfileRequestFlexRadios) Set(val *CreateNetworkWirelessRfProfileRequestFlexRadios) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWirelessRfProfileRequestFlexRadios) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWirelessRfProfileRequestFlexRadios) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWirelessRfProfileRequestFlexRadios(val *CreateNetworkWirelessRfProfileRequestFlexRadios) *NullableCreateNetworkWirelessRfProfileRequestFlexRadios {
	return &NullableCreateNetworkWirelessRfProfileRequestFlexRadios{value: val, isSet: true}
}

func (v NullableCreateNetworkWirelessRfProfileRequestFlexRadios) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWirelessRfProfileRequestFlexRadios) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


