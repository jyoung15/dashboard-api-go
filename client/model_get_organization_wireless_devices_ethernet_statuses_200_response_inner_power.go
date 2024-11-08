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

// checks if the GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower{}

// GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower Power details object
type GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower struct {
	// The PoE power mode for the AP. Can be 'full' or 'low'
	Mode *string `json:"mode,omitempty"`
	Ac *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc `json:"ac,omitempty"`
	Poe *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerPoe `json:"poe,omitempty"`
}

// NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower instantiates a new GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower {
	this := GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower{}
	return &this
}

// NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerWithDefaults instantiates a new GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerWithDefaults() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower {
	this := GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) SetMode(v string) {
	o.Mode = &v
}

// GetAc returns the Ac field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) GetAc() GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc {
	if o == nil || IsNil(o.Ac) {
		var ret GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc
		return ret
	}
	return *o.Ac
}

// GetAcOk returns a tuple with the Ac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) GetAcOk() (*GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc, bool) {
	if o == nil || IsNil(o.Ac) {
		return nil, false
	}
	return o.Ac, true
}

// HasAc returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) HasAc() bool {
	if o != nil && !IsNil(o.Ac) {
		return true
	}

	return false
}

// SetAc gets a reference to the given GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc and assigns it to the Ac field.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) SetAc(v GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc) {
	o.Ac = &v
}

// GetPoe returns the Poe field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) GetPoe() GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerPoe {
	if o == nil || IsNil(o.Poe) {
		var ret GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerPoe
		return ret
	}
	return *o.Poe
}

// GetPoeOk returns a tuple with the Poe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) GetPoeOk() (*GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerPoe, bool) {
	if o == nil || IsNil(o.Poe) {
		return nil, false
	}
	return o.Poe, true
}

// HasPoe returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) HasPoe() bool {
	if o != nil && !IsNil(o.Poe) {
		return true
	}

	return false
}

// SetPoe gets a reference to the given GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerPoe and assigns it to the Poe field.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) SetPoe(v GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerPoe) {
	o.Poe = &v
}

func (o GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Ac) {
		toSerialize["ac"] = o.Ac
	}
	if !IsNil(o.Poe) {
		toSerialize["poe"] = o.Poe
	}
	return toSerialize, nil
}

type NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower struct {
	value *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower
	isSet bool
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) Get() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower {
	return v.value
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) Set(val *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower(val *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower {
	return &NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPower) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


