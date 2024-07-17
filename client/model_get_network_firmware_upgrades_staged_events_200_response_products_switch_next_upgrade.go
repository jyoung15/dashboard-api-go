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

// checks if the GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade{}

// GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade Details of the next firmware upgrade
type GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade struct {
	ToVersion *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion `json:"toVersion,omitempty"`
}

// NewGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade instantiates a new GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade() *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade {
	this := GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade{}
	return &this
}

// NewGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeWithDefaults instantiates a new GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeWithDefaults() *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade {
	this := GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade{}
	return &this
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade) GetToVersion() GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion {
	if o == nil || IsNil(o.ToVersion) {
		var ret GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade) GetToVersionOk() (*GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion, bool) {
	if o == nil || IsNil(o.ToVersion) {
		return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade) HasToVersion() bool {
	if o != nil && !IsNil(o.ToVersion) {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion and assigns it to the ToVersion field.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade) SetToVersion(v GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion) {
	o.ToVersion = &v
}

func (o GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ToVersion) {
		toSerialize["toVersion"] = o.ToVersion
	}
	return toSerialize, nil
}

type NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade struct {
	value *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade
	isSet bool
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade) Get() *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade {
	return v.value
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade) Set(val *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade(val *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade) *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade {
	return &NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade{value: val, isSet: true}
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


