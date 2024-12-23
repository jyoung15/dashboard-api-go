/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade{}

// GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade Details of the next firmware upgrade on the device
type GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade struct {
	// Timestamp of the next scheduled firmware upgrade
	Time *time.Time `json:"time,omitempty"`
	ToVersion *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion `json:"toVersion,omitempty"`
}

// NewGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade instantiates a new GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade() *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade {
	this := GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade{}
	return &this
}

// NewGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeWithDefaults instantiates a new GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeWithDefaults() *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade {
	this := GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade) SetTime(v time.Time) {
	o.Time = &v
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade) GetToVersion() GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion {
	if o == nil || IsNil(o.ToVersion) {
		var ret GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade) GetToVersionOk() (*GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion, bool) {
	if o == nil || IsNil(o.ToVersion) {
		return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade) HasToVersion() bool {
	if o != nil && !IsNil(o.ToVersion) {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion and assigns it to the ToVersion field.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade) SetToVersion(v GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) {
	o.ToVersion = &v
}

func (o GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.ToVersion) {
		toSerialize["toVersion"] = o.ToVersion
	}
	return toSerialize, nil
}

type NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade struct {
	value *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade
	isSet bool
}

func (v NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade) Get() *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade {
	return v.value
}

func (v *NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade) Set(val *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade(val *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade) *NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade {
	return &NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade{value: val, isSet: true}
}

func (v NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


