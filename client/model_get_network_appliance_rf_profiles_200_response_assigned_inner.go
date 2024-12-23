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

// checks if the GetNetworkApplianceRfProfiles200ResponseAssignedInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceRfProfiles200ResponseAssignedInner{}

// GetNetworkApplianceRfProfiles200ResponseAssignedInner struct for GetNetworkApplianceRfProfiles200ResponseAssignedInner
type GetNetworkApplianceRfProfiles200ResponseAssignedInner struct {
	// ID of the RF Profile.
	Id *string `json:"id,omitempty"`
	// ID of network this RF Profile belongs in.
	NetworkId *string `json:"networkId,omitempty"`
	// The name of the profile.
	Name *string `json:"name,omitempty"`
	TwoFourGhzSettings *GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`
	FiveGhzSettings *GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings `json:"fiveGhzSettings,omitempty"`
	PerSsidSettings *GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings `json:"perSsidSettings,omitempty"`
}

// NewGetNetworkApplianceRfProfiles200ResponseAssignedInner instantiates a new GetNetworkApplianceRfProfiles200ResponseAssignedInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceRfProfiles200ResponseAssignedInner() *GetNetworkApplianceRfProfiles200ResponseAssignedInner {
	this := GetNetworkApplianceRfProfiles200ResponseAssignedInner{}
	return &this
}

// NewGetNetworkApplianceRfProfiles200ResponseAssignedInnerWithDefaults instantiates a new GetNetworkApplianceRfProfiles200ResponseAssignedInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceRfProfiles200ResponseAssignedInnerWithDefaults() *GetNetworkApplianceRfProfiles200ResponseAssignedInner {
	this := GetNetworkApplianceRfProfiles200ResponseAssignedInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) SetId(v string) {
	o.Id = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetNetworkId() string {
	if o == nil || IsNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) SetName(v string) {
	o.Name = &v
}

// GetTwoFourGhzSettings returns the TwoFourGhzSettings field value if set, zero value otherwise.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetTwoFourGhzSettings() GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings {
	if o == nil || IsNil(o.TwoFourGhzSettings) {
		var ret GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings
		return ret
	}
	return *o.TwoFourGhzSettings
}

// GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetTwoFourGhzSettingsOk() (*GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings, bool) {
	if o == nil || IsNil(o.TwoFourGhzSettings) {
		return nil, false
	}
	return o.TwoFourGhzSettings, true
}

// HasTwoFourGhzSettings returns a boolean if a field has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) HasTwoFourGhzSettings() bool {
	if o != nil && !IsNil(o.TwoFourGhzSettings) {
		return true
	}

	return false
}

// SetTwoFourGhzSettings gets a reference to the given GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings and assigns it to the TwoFourGhzSettings field.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) SetTwoFourGhzSettings(v GetNetworkApplianceRfProfiles200ResponseAssignedInnerTwoFourGhzSettings) {
	o.TwoFourGhzSettings = &v
}

// GetFiveGhzSettings returns the FiveGhzSettings field value if set, zero value otherwise.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetFiveGhzSettings() GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings {
	if o == nil || IsNil(o.FiveGhzSettings) {
		var ret GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings
		return ret
	}
	return *o.FiveGhzSettings
}

// GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetFiveGhzSettingsOk() (*GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings, bool) {
	if o == nil || IsNil(o.FiveGhzSettings) {
		return nil, false
	}
	return o.FiveGhzSettings, true
}

// HasFiveGhzSettings returns a boolean if a field has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) HasFiveGhzSettings() bool {
	if o != nil && !IsNil(o.FiveGhzSettings) {
		return true
	}

	return false
}

// SetFiveGhzSettings gets a reference to the given GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings and assigns it to the FiveGhzSettings field.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) SetFiveGhzSettings(v GetNetworkApplianceRfProfiles200ResponseAssignedInnerFiveGhzSettings) {
	o.FiveGhzSettings = &v
}

// GetPerSsidSettings returns the PerSsidSettings field value if set, zero value otherwise.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetPerSsidSettings() GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings {
	if o == nil || IsNil(o.PerSsidSettings) {
		var ret GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings
		return ret
	}
	return *o.PerSsidSettings
}

// GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) GetPerSsidSettingsOk() (*GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings, bool) {
	if o == nil || IsNil(o.PerSsidSettings) {
		return nil, false
	}
	return o.PerSsidSettings, true
}

// HasPerSsidSettings returns a boolean if a field has been set.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) HasPerSsidSettings() bool {
	if o != nil && !IsNil(o.PerSsidSettings) {
		return true
	}

	return false
}

// SetPerSsidSettings gets a reference to the given GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings and assigns it to the PerSsidSettings field.
func (o *GetNetworkApplianceRfProfiles200ResponseAssignedInner) SetPerSsidSettings(v GetNetworkApplianceRfProfiles200ResponseAssignedInnerPerSsidSettings) {
	o.PerSsidSettings = &v
}

func (o GetNetworkApplianceRfProfiles200ResponseAssignedInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceRfProfiles200ResponseAssignedInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TwoFourGhzSettings) {
		toSerialize["twoFourGhzSettings"] = o.TwoFourGhzSettings
	}
	if !IsNil(o.FiveGhzSettings) {
		toSerialize["fiveGhzSettings"] = o.FiveGhzSettings
	}
	if !IsNil(o.PerSsidSettings) {
		toSerialize["perSsidSettings"] = o.PerSsidSettings
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceRfProfiles200ResponseAssignedInner struct {
	value *GetNetworkApplianceRfProfiles200ResponseAssignedInner
	isSet bool
}

func (v NullableGetNetworkApplianceRfProfiles200ResponseAssignedInner) Get() *GetNetworkApplianceRfProfiles200ResponseAssignedInner {
	return v.value
}

func (v *NullableGetNetworkApplianceRfProfiles200ResponseAssignedInner) Set(val *GetNetworkApplianceRfProfiles200ResponseAssignedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceRfProfiles200ResponseAssignedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceRfProfiles200ResponseAssignedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceRfProfiles200ResponseAssignedInner(val *GetNetworkApplianceRfProfiles200ResponseAssignedInner) *NullableGetNetworkApplianceRfProfiles200ResponseAssignedInner {
	return &NullableGetNetworkApplianceRfProfiles200ResponseAssignedInner{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceRfProfiles200ResponseAssignedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceRfProfiles200ResponseAssignedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


