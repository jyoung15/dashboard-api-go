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

// checks if the UpdateNetworkWirelessRfProfileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessRfProfileRequest{}

// UpdateNetworkWirelessRfProfileRequest struct for UpdateNetworkWirelessRfProfileRequest
type UpdateNetworkWirelessRfProfileRequest struct {
	// The name of the new profile. Must be unique.
	Name *string `json:"name,omitempty"`
	// Steers client to best available access point. Can be either true or false.
	ClientBalancingEnabled *bool `json:"clientBalancingEnabled,omitempty"`
	// Minimum bitrate can be set to either 'band' or 'ssid'.
	MinBitrateType *string `json:"minBitrateType,omitempty"`
	// Band selection can be set to either 'ssid' or 'ap'.
	BandSelectionType *string `json:"bandSelectionType,omitempty"`
	ApBandSettings *UpdateNetworkWirelessRfProfileRequestApBandSettings `json:"apBandSettings,omitempty"`
	TwoFourGhzSettings *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`
	FiveGhzSettings *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings `json:"fiveGhzSettings,omitempty"`
	SixGhzSettings *UpdateNetworkWirelessRfProfileRequestSixGhzSettings `json:"sixGhzSettings,omitempty"`
	Transmission *GetNetworkWirelessRfProfiles200ResponseTransmission `json:"transmission,omitempty"`
	PerSsidSettings *CreateNetworkWirelessRfProfileRequestPerSsidSettings `json:"perSsidSettings,omitempty"`
	FlexRadios *CreateNetworkWirelessRfProfileRequestFlexRadios `json:"flexRadios,omitempty"`
}

// NewUpdateNetworkWirelessRfProfileRequest instantiates a new UpdateNetworkWirelessRfProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessRfProfileRequest() *UpdateNetworkWirelessRfProfileRequest {
	this := UpdateNetworkWirelessRfProfileRequest{}
	return &this
}

// NewUpdateNetworkWirelessRfProfileRequestWithDefaults instantiates a new UpdateNetworkWirelessRfProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessRfProfileRequestWithDefaults() *UpdateNetworkWirelessRfProfileRequest {
	this := UpdateNetworkWirelessRfProfileRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessRfProfileRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkWirelessRfProfileRequest) SetName(v string) {
	o.Name = &v
}

// GetClientBalancingEnabled returns the ClientBalancingEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessRfProfileRequest) GetClientBalancingEnabled() bool {
	if o == nil || IsNil(o.ClientBalancingEnabled) {
		var ret bool
		return ret
	}
	return *o.ClientBalancingEnabled
}

// GetClientBalancingEnabledOk returns a tuple with the ClientBalancingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) GetClientBalancingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientBalancingEnabled) {
		return nil, false
	}
	return o.ClientBalancingEnabled, true
}

// HasClientBalancingEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) HasClientBalancingEnabled() bool {
	if o != nil && !IsNil(o.ClientBalancingEnabled) {
		return true
	}

	return false
}

// SetClientBalancingEnabled gets a reference to the given bool and assigns it to the ClientBalancingEnabled field.
func (o *UpdateNetworkWirelessRfProfileRequest) SetClientBalancingEnabled(v bool) {
	o.ClientBalancingEnabled = &v
}

// GetMinBitrateType returns the MinBitrateType field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessRfProfileRequest) GetMinBitrateType() string {
	if o == nil || IsNil(o.MinBitrateType) {
		var ret string
		return ret
	}
	return *o.MinBitrateType
}

// GetMinBitrateTypeOk returns a tuple with the MinBitrateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) GetMinBitrateTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MinBitrateType) {
		return nil, false
	}
	return o.MinBitrateType, true
}

// HasMinBitrateType returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) HasMinBitrateType() bool {
	if o != nil && !IsNil(o.MinBitrateType) {
		return true
	}

	return false
}

// SetMinBitrateType gets a reference to the given string and assigns it to the MinBitrateType field.
func (o *UpdateNetworkWirelessRfProfileRequest) SetMinBitrateType(v string) {
	o.MinBitrateType = &v
}

// GetBandSelectionType returns the BandSelectionType field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessRfProfileRequest) GetBandSelectionType() string {
	if o == nil || IsNil(o.BandSelectionType) {
		var ret string
		return ret
	}
	return *o.BandSelectionType
}

// GetBandSelectionTypeOk returns a tuple with the BandSelectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) GetBandSelectionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BandSelectionType) {
		return nil, false
	}
	return o.BandSelectionType, true
}

// HasBandSelectionType returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) HasBandSelectionType() bool {
	if o != nil && !IsNil(o.BandSelectionType) {
		return true
	}

	return false
}

// SetBandSelectionType gets a reference to the given string and assigns it to the BandSelectionType field.
func (o *UpdateNetworkWirelessRfProfileRequest) SetBandSelectionType(v string) {
	o.BandSelectionType = &v
}

// GetApBandSettings returns the ApBandSettings field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessRfProfileRequest) GetApBandSettings() UpdateNetworkWirelessRfProfileRequestApBandSettings {
	if o == nil || IsNil(o.ApBandSettings) {
		var ret UpdateNetworkWirelessRfProfileRequestApBandSettings
		return ret
	}
	return *o.ApBandSettings
}

// GetApBandSettingsOk returns a tuple with the ApBandSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) GetApBandSettingsOk() (*UpdateNetworkWirelessRfProfileRequestApBandSettings, bool) {
	if o == nil || IsNil(o.ApBandSettings) {
		return nil, false
	}
	return o.ApBandSettings, true
}

// HasApBandSettings returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) HasApBandSettings() bool {
	if o != nil && !IsNil(o.ApBandSettings) {
		return true
	}

	return false
}

// SetApBandSettings gets a reference to the given UpdateNetworkWirelessRfProfileRequestApBandSettings and assigns it to the ApBandSettings field.
func (o *UpdateNetworkWirelessRfProfileRequest) SetApBandSettings(v UpdateNetworkWirelessRfProfileRequestApBandSettings) {
	o.ApBandSettings = &v
}

// GetTwoFourGhzSettings returns the TwoFourGhzSettings field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessRfProfileRequest) GetTwoFourGhzSettings() UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings {
	if o == nil || IsNil(o.TwoFourGhzSettings) {
		var ret UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings
		return ret
	}
	return *o.TwoFourGhzSettings
}

// GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) GetTwoFourGhzSettingsOk() (*UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings, bool) {
	if o == nil || IsNil(o.TwoFourGhzSettings) {
		return nil, false
	}
	return o.TwoFourGhzSettings, true
}

// HasTwoFourGhzSettings returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) HasTwoFourGhzSettings() bool {
	if o != nil && !IsNil(o.TwoFourGhzSettings) {
		return true
	}

	return false
}

// SetTwoFourGhzSettings gets a reference to the given UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings and assigns it to the TwoFourGhzSettings field.
func (o *UpdateNetworkWirelessRfProfileRequest) SetTwoFourGhzSettings(v UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) {
	o.TwoFourGhzSettings = &v
}

// GetFiveGhzSettings returns the FiveGhzSettings field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessRfProfileRequest) GetFiveGhzSettings() UpdateNetworkWirelessRfProfileRequestFiveGhzSettings {
	if o == nil || IsNil(o.FiveGhzSettings) {
		var ret UpdateNetworkWirelessRfProfileRequestFiveGhzSettings
		return ret
	}
	return *o.FiveGhzSettings
}

// GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) GetFiveGhzSettingsOk() (*UpdateNetworkWirelessRfProfileRequestFiveGhzSettings, bool) {
	if o == nil || IsNil(o.FiveGhzSettings) {
		return nil, false
	}
	return o.FiveGhzSettings, true
}

// HasFiveGhzSettings returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) HasFiveGhzSettings() bool {
	if o != nil && !IsNil(o.FiveGhzSettings) {
		return true
	}

	return false
}

// SetFiveGhzSettings gets a reference to the given UpdateNetworkWirelessRfProfileRequestFiveGhzSettings and assigns it to the FiveGhzSettings field.
func (o *UpdateNetworkWirelessRfProfileRequest) SetFiveGhzSettings(v UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) {
	o.FiveGhzSettings = &v
}

// GetSixGhzSettings returns the SixGhzSettings field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessRfProfileRequest) GetSixGhzSettings() UpdateNetworkWirelessRfProfileRequestSixGhzSettings {
	if o == nil || IsNil(o.SixGhzSettings) {
		var ret UpdateNetworkWirelessRfProfileRequestSixGhzSettings
		return ret
	}
	return *o.SixGhzSettings
}

// GetSixGhzSettingsOk returns a tuple with the SixGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) GetSixGhzSettingsOk() (*UpdateNetworkWirelessRfProfileRequestSixGhzSettings, bool) {
	if o == nil || IsNil(o.SixGhzSettings) {
		return nil, false
	}
	return o.SixGhzSettings, true
}

// HasSixGhzSettings returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) HasSixGhzSettings() bool {
	if o != nil && !IsNil(o.SixGhzSettings) {
		return true
	}

	return false
}

// SetSixGhzSettings gets a reference to the given UpdateNetworkWirelessRfProfileRequestSixGhzSettings and assigns it to the SixGhzSettings field.
func (o *UpdateNetworkWirelessRfProfileRequest) SetSixGhzSettings(v UpdateNetworkWirelessRfProfileRequestSixGhzSettings) {
	o.SixGhzSettings = &v
}

// GetTransmission returns the Transmission field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessRfProfileRequest) GetTransmission() GetNetworkWirelessRfProfiles200ResponseTransmission {
	if o == nil || IsNil(o.Transmission) {
		var ret GetNetworkWirelessRfProfiles200ResponseTransmission
		return ret
	}
	return *o.Transmission
}

// GetTransmissionOk returns a tuple with the Transmission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) GetTransmissionOk() (*GetNetworkWirelessRfProfiles200ResponseTransmission, bool) {
	if o == nil || IsNil(o.Transmission) {
		return nil, false
	}
	return o.Transmission, true
}

// HasTransmission returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) HasTransmission() bool {
	if o != nil && !IsNil(o.Transmission) {
		return true
	}

	return false
}

// SetTransmission gets a reference to the given GetNetworkWirelessRfProfiles200ResponseTransmission and assigns it to the Transmission field.
func (o *UpdateNetworkWirelessRfProfileRequest) SetTransmission(v GetNetworkWirelessRfProfiles200ResponseTransmission) {
	o.Transmission = &v
}

// GetPerSsidSettings returns the PerSsidSettings field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessRfProfileRequest) GetPerSsidSettings() CreateNetworkWirelessRfProfileRequestPerSsidSettings {
	if o == nil || IsNil(o.PerSsidSettings) {
		var ret CreateNetworkWirelessRfProfileRequestPerSsidSettings
		return ret
	}
	return *o.PerSsidSettings
}

// GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) GetPerSsidSettingsOk() (*CreateNetworkWirelessRfProfileRequestPerSsidSettings, bool) {
	if o == nil || IsNil(o.PerSsidSettings) {
		return nil, false
	}
	return o.PerSsidSettings, true
}

// HasPerSsidSettings returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) HasPerSsidSettings() bool {
	if o != nil && !IsNil(o.PerSsidSettings) {
		return true
	}

	return false
}

// SetPerSsidSettings gets a reference to the given CreateNetworkWirelessRfProfileRequestPerSsidSettings and assigns it to the PerSsidSettings field.
func (o *UpdateNetworkWirelessRfProfileRequest) SetPerSsidSettings(v CreateNetworkWirelessRfProfileRequestPerSsidSettings) {
	o.PerSsidSettings = &v
}

// GetFlexRadios returns the FlexRadios field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessRfProfileRequest) GetFlexRadios() CreateNetworkWirelessRfProfileRequestFlexRadios {
	if o == nil || IsNil(o.FlexRadios) {
		var ret CreateNetworkWirelessRfProfileRequestFlexRadios
		return ret
	}
	return *o.FlexRadios
}

// GetFlexRadiosOk returns a tuple with the FlexRadios field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) GetFlexRadiosOk() (*CreateNetworkWirelessRfProfileRequestFlexRadios, bool) {
	if o == nil || IsNil(o.FlexRadios) {
		return nil, false
	}
	return o.FlexRadios, true
}

// HasFlexRadios returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessRfProfileRequest) HasFlexRadios() bool {
	if o != nil && !IsNil(o.FlexRadios) {
		return true
	}

	return false
}

// SetFlexRadios gets a reference to the given CreateNetworkWirelessRfProfileRequestFlexRadios and assigns it to the FlexRadios field.
func (o *UpdateNetworkWirelessRfProfileRequest) SetFlexRadios(v CreateNetworkWirelessRfProfileRequestFlexRadios) {
	o.FlexRadios = &v
}

func (o UpdateNetworkWirelessRfProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessRfProfileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ClientBalancingEnabled) {
		toSerialize["clientBalancingEnabled"] = o.ClientBalancingEnabled
	}
	if !IsNil(o.MinBitrateType) {
		toSerialize["minBitrateType"] = o.MinBitrateType
	}
	if !IsNil(o.BandSelectionType) {
		toSerialize["bandSelectionType"] = o.BandSelectionType
	}
	if !IsNil(o.ApBandSettings) {
		toSerialize["apBandSettings"] = o.ApBandSettings
	}
	if !IsNil(o.TwoFourGhzSettings) {
		toSerialize["twoFourGhzSettings"] = o.TwoFourGhzSettings
	}
	if !IsNil(o.FiveGhzSettings) {
		toSerialize["fiveGhzSettings"] = o.FiveGhzSettings
	}
	if !IsNil(o.SixGhzSettings) {
		toSerialize["sixGhzSettings"] = o.SixGhzSettings
	}
	if !IsNil(o.Transmission) {
		toSerialize["transmission"] = o.Transmission
	}
	if !IsNil(o.PerSsidSettings) {
		toSerialize["perSsidSettings"] = o.PerSsidSettings
	}
	if !IsNil(o.FlexRadios) {
		toSerialize["flexRadios"] = o.FlexRadios
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessRfProfileRequest struct {
	value *UpdateNetworkWirelessRfProfileRequest
	isSet bool
}

func (v NullableUpdateNetworkWirelessRfProfileRequest) Get() *UpdateNetworkWirelessRfProfileRequest {
	return v.value
}

func (v *NullableUpdateNetworkWirelessRfProfileRequest) Set(val *UpdateNetworkWirelessRfProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessRfProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessRfProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessRfProfileRequest(val *UpdateNetworkWirelessRfProfileRequest) *NullableUpdateNetworkWirelessRfProfileRequest {
	return &NullableUpdateNetworkWirelessRfProfileRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessRfProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessRfProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


