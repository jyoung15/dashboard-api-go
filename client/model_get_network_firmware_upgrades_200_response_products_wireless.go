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

// checks if the GetNetworkFirmwareUpgrades200ResponseProductsWireless type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkFirmwareUpgrades200ResponseProductsWireless{}

// GetNetworkFirmwareUpgrades200ResponseProductsWireless The network device to be updated
type GetNetworkFirmwareUpgrades200ResponseProductsWireless struct {
	CurrentVersion *GetNetworkFirmwareUpgrades200ResponseProductsWirelessCurrentVersion `json:"currentVersion,omitempty"`
	LastUpgrade *GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade `json:"lastUpgrade,omitempty"`
	NextUpgrade *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade `json:"nextUpgrade,omitempty"`
	// Firmware versions available for upgrade
	AvailableVersions []GetNetworkFirmwareUpgrades200ResponseProductsWirelessAvailableVersionsInner `json:"availableVersions,omitempty"`
	// Whether or not the network wants beta firmware
	ParticipateInNextBetaRelease *bool `json:"participateInNextBetaRelease,omitempty"`
}

// NewGetNetworkFirmwareUpgrades200ResponseProductsWireless instantiates a new GetNetworkFirmwareUpgrades200ResponseProductsWireless object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFirmwareUpgrades200ResponseProductsWireless() *GetNetworkFirmwareUpgrades200ResponseProductsWireless {
	this := GetNetworkFirmwareUpgrades200ResponseProductsWireless{}
	return &this
}

// NewGetNetworkFirmwareUpgrades200ResponseProductsWirelessWithDefaults instantiates a new GetNetworkFirmwareUpgrades200ResponseProductsWireless object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFirmwareUpgrades200ResponseProductsWirelessWithDefaults() *GetNetworkFirmwareUpgrades200ResponseProductsWireless {
	this := GetNetworkFirmwareUpgrades200ResponseProductsWireless{}
	return &this
}

// GetCurrentVersion returns the CurrentVersion field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) GetCurrentVersion() GetNetworkFirmwareUpgrades200ResponseProductsWirelessCurrentVersion {
	if o == nil || IsNil(o.CurrentVersion) {
		var ret GetNetworkFirmwareUpgrades200ResponseProductsWirelessCurrentVersion
		return ret
	}
	return *o.CurrentVersion
}

// GetCurrentVersionOk returns a tuple with the CurrentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) GetCurrentVersionOk() (*GetNetworkFirmwareUpgrades200ResponseProductsWirelessCurrentVersion, bool) {
	if o == nil || IsNil(o.CurrentVersion) {
		return nil, false
	}
	return o.CurrentVersion, true
}

// HasCurrentVersion returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) HasCurrentVersion() bool {
	if o != nil && !IsNil(o.CurrentVersion) {
		return true
	}

	return false
}

// SetCurrentVersion gets a reference to the given GetNetworkFirmwareUpgrades200ResponseProductsWirelessCurrentVersion and assigns it to the CurrentVersion field.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) SetCurrentVersion(v GetNetworkFirmwareUpgrades200ResponseProductsWirelessCurrentVersion) {
	o.CurrentVersion = &v
}

// GetLastUpgrade returns the LastUpgrade field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) GetLastUpgrade() GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade {
	if o == nil || IsNil(o.LastUpgrade) {
		var ret GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade
		return ret
	}
	return *o.LastUpgrade
}

// GetLastUpgradeOk returns a tuple with the LastUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) GetLastUpgradeOk() (*GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade, bool) {
	if o == nil || IsNil(o.LastUpgrade) {
		return nil, false
	}
	return o.LastUpgrade, true
}

// HasLastUpgrade returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) HasLastUpgrade() bool {
	if o != nil && !IsNil(o.LastUpgrade) {
		return true
	}

	return false
}

// SetLastUpgrade gets a reference to the given GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade and assigns it to the LastUpgrade field.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) SetLastUpgrade(v GetNetworkFirmwareUpgrades200ResponseProductsWirelessLastUpgrade) {
	o.LastUpgrade = &v
}

// GetNextUpgrade returns the NextUpgrade field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) GetNextUpgrade() GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade {
	if o == nil || IsNil(o.NextUpgrade) {
		var ret GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade
		return ret
	}
	return *o.NextUpgrade
}

// GetNextUpgradeOk returns a tuple with the NextUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) GetNextUpgradeOk() (*GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade, bool) {
	if o == nil || IsNil(o.NextUpgrade) {
		return nil, false
	}
	return o.NextUpgrade, true
}

// HasNextUpgrade returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) HasNextUpgrade() bool {
	if o != nil && !IsNil(o.NextUpgrade) {
		return true
	}

	return false
}

// SetNextUpgrade gets a reference to the given GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade and assigns it to the NextUpgrade field.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) SetNextUpgrade(v GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgrade) {
	o.NextUpgrade = &v
}

// GetAvailableVersions returns the AvailableVersions field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) GetAvailableVersions() []GetNetworkFirmwareUpgrades200ResponseProductsWirelessAvailableVersionsInner {
	if o == nil || IsNil(o.AvailableVersions) {
		var ret []GetNetworkFirmwareUpgrades200ResponseProductsWirelessAvailableVersionsInner
		return ret
	}
	return o.AvailableVersions
}

// GetAvailableVersionsOk returns a tuple with the AvailableVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) GetAvailableVersionsOk() ([]GetNetworkFirmwareUpgrades200ResponseProductsWirelessAvailableVersionsInner, bool) {
	if o == nil || IsNil(o.AvailableVersions) {
		return nil, false
	}
	return o.AvailableVersions, true
}

// HasAvailableVersions returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) HasAvailableVersions() bool {
	if o != nil && !IsNil(o.AvailableVersions) {
		return true
	}

	return false
}

// SetAvailableVersions gets a reference to the given []GetNetworkFirmwareUpgrades200ResponseProductsWirelessAvailableVersionsInner and assigns it to the AvailableVersions field.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) SetAvailableVersions(v []GetNetworkFirmwareUpgrades200ResponseProductsWirelessAvailableVersionsInner) {
	o.AvailableVersions = v
}

// GetParticipateInNextBetaRelease returns the ParticipateInNextBetaRelease field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) GetParticipateInNextBetaRelease() bool {
	if o == nil || IsNil(o.ParticipateInNextBetaRelease) {
		var ret bool
		return ret
	}
	return *o.ParticipateInNextBetaRelease
}

// GetParticipateInNextBetaReleaseOk returns a tuple with the ParticipateInNextBetaRelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) GetParticipateInNextBetaReleaseOk() (*bool, bool) {
	if o == nil || IsNil(o.ParticipateInNextBetaRelease) {
		return nil, false
	}
	return o.ParticipateInNextBetaRelease, true
}

// HasParticipateInNextBetaRelease returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) HasParticipateInNextBetaRelease() bool {
	if o != nil && !IsNil(o.ParticipateInNextBetaRelease) {
		return true
	}

	return false
}

// SetParticipateInNextBetaRelease gets a reference to the given bool and assigns it to the ParticipateInNextBetaRelease field.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWireless) SetParticipateInNextBetaRelease(v bool) {
	o.ParticipateInNextBetaRelease = &v
}

func (o GetNetworkFirmwareUpgrades200ResponseProductsWireless) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkFirmwareUpgrades200ResponseProductsWireless) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentVersion) {
		toSerialize["currentVersion"] = o.CurrentVersion
	}
	if !IsNil(o.LastUpgrade) {
		toSerialize["lastUpgrade"] = o.LastUpgrade
	}
	if !IsNil(o.NextUpgrade) {
		toSerialize["nextUpgrade"] = o.NextUpgrade
	}
	if !IsNil(o.AvailableVersions) {
		toSerialize["availableVersions"] = o.AvailableVersions
	}
	if !IsNil(o.ParticipateInNextBetaRelease) {
		toSerialize["participateInNextBetaRelease"] = o.ParticipateInNextBetaRelease
	}
	return toSerialize, nil
}

type NullableGetNetworkFirmwareUpgrades200ResponseProductsWireless struct {
	value *GetNetworkFirmwareUpgrades200ResponseProductsWireless
	isSet bool
}

func (v NullableGetNetworkFirmwareUpgrades200ResponseProductsWireless) Get() *GetNetworkFirmwareUpgrades200ResponseProductsWireless {
	return v.value
}

func (v *NullableGetNetworkFirmwareUpgrades200ResponseProductsWireless) Set(val *GetNetworkFirmwareUpgrades200ResponseProductsWireless) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFirmwareUpgrades200ResponseProductsWireless) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFirmwareUpgrades200ResponseProductsWireless) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFirmwareUpgrades200ResponseProductsWireless(val *GetNetworkFirmwareUpgrades200ResponseProductsWireless) *NullableGetNetworkFirmwareUpgrades200ResponseProductsWireless {
	return &NullableGetNetworkFirmwareUpgrades200ResponseProductsWireless{value: val, isSet: true}
}

func (v NullableGetNetworkFirmwareUpgrades200ResponseProductsWireless) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFirmwareUpgrades200ResponseProductsWireless) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


