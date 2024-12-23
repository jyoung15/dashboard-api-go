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

// checks if the GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion{}

// GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion Details of the version the device will upgrade to
type GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion struct {
	// Id of the Version being upgraded to
	Id *string `json:"id,omitempty"`
	// Firmware version short name
	ShortName *string `json:"shortName,omitempty"`
}

// NewGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion instantiates a new GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion() *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion {
	this := GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion{}
	return &this
}

// NewGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersionWithDefaults instantiates a new GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersionWithDefaults() *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion {
	this := GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion) SetId(v string) {
	o.Id = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion) SetShortName(v string) {
	o.ShortName = &v
}

func (o GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	return toSerialize, nil
}

type NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion struct {
	value *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion
	isSet bool
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion) Get() *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion {
	return v.value
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion) Set(val *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion(val *GetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion) *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion {
	return &NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion{value: val, isSet: true}
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseProductsSwitchNextUpgradeToVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


