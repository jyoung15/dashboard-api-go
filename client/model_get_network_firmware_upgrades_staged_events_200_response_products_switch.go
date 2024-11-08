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

// checks if the GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch{}

// GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch The Switch network to be updated
type GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch struct {
	NextUpgrade *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade `json:"nextUpgrade,omitempty"`
}

// NewGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch instantiates a new GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch() *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch {
	this := GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch{}
	return &this
}

// NewGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchWithDefaults instantiates a new GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchWithDefaults() *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch {
	this := GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch{}
	return &this
}

// GetNextUpgrade returns the NextUpgrade field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch) GetNextUpgrade() GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade {
	if o == nil || IsNil(o.NextUpgrade) {
		var ret GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade
		return ret
	}
	return *o.NextUpgrade
}

// GetNextUpgradeOk returns a tuple with the NextUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch) GetNextUpgradeOk() (*GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade, bool) {
	if o == nil || IsNil(o.NextUpgrade) {
		return nil, false
	}
	return o.NextUpgrade, true
}

// HasNextUpgrade returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch) HasNextUpgrade() bool {
	if o != nil && !IsNil(o.NextUpgrade) {
		return true
	}

	return false
}

// SetNextUpgrade gets a reference to the given GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade and assigns it to the NextUpgrade field.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch) SetNextUpgrade(v GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade) {
	o.NextUpgrade = &v
}

func (o GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextUpgrade) {
		toSerialize["nextUpgrade"] = o.NextUpgrade
	}
	return toSerialize, nil
}

type NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch struct {
	value *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch
	isSet bool
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch) Get() *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch {
	return v.value
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch) Set(val *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch(val *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch) *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch {
	return &NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch{value: val, isSet: true}
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


