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

// checks if the UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion{}

// UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion The version to be updated to
type UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion struct {
	// The version ID
	Id *string `json:"id,omitempty"`
}

// NewUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion instantiates a new UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion() *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion {
	this := UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion{}
	return &this
}

// NewUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersionWithDefaults instantiates a new UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersionWithDefaults() *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion {
	this := UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) SetId(v string) {
	o.Id = &v
}

func (o UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion struct {
	value *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion
	isSet bool
}

func (v NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) Get() *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion {
	return v.value
}

func (v *NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) Set(val *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion(val *UpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) *NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion {
	return &NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion{value: val, isSet: true}
}

func (v NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkFirmwareUpgradesRequestProductsWirelessNextUpgradeToVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


