/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateNetworkWirelessRfProfileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkWirelessRfProfileRequest{}

// CreateNetworkWirelessRfProfileRequest struct for CreateNetworkWirelessRfProfileRequest
type CreateNetworkWirelessRfProfileRequest struct {
	// The name of the new profile. Must be unique. This param is required on creation.
	Name string `json:"name"`
	// Steers client to best available access point. Can be either true or false. Defaults to true.
	ClientBalancingEnabled *bool `json:"clientBalancingEnabled,omitempty"`
	// Minimum bitrate can be set to either 'band' or 'ssid'. Defaults to band.
	MinBitrateType *string `json:"minBitrateType,omitempty"`
	// Band selection can be set to either 'ssid' or 'ap'. This param is required on creation.
	BandSelectionType string `json:"bandSelectionType"`
	ApBandSettings *CreateNetworkWirelessRfProfileRequestApBandSettings `json:"apBandSettings,omitempty"`
	TwoFourGhzSettings *GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`
	FiveGhzSettings *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings `json:"fiveGhzSettings,omitempty"`
	SixGhzSettings *CreateNetworkWirelessRfProfileRequestSixGhzSettings `json:"sixGhzSettings,omitempty"`
	Transmission *GetNetworkWirelessRfProfiles200ResponseTransmission `json:"transmission,omitempty"`
	PerSsidSettings *CreateNetworkWirelessRfProfileRequestPerSsidSettings `json:"perSsidSettings,omitempty"`
	FlexRadios *CreateNetworkWirelessRfProfileRequestFlexRadios `json:"flexRadios,omitempty"`
}

type _CreateNetworkWirelessRfProfileRequest CreateNetworkWirelessRfProfileRequest

// NewCreateNetworkWirelessRfProfileRequest instantiates a new CreateNetworkWirelessRfProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWirelessRfProfileRequest(name string, bandSelectionType string) *CreateNetworkWirelessRfProfileRequest {
	this := CreateNetworkWirelessRfProfileRequest{}
	this.Name = name
	this.BandSelectionType = bandSelectionType
	return &this
}

// NewCreateNetworkWirelessRfProfileRequestWithDefaults instantiates a new CreateNetworkWirelessRfProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWirelessRfProfileRequestWithDefaults() *CreateNetworkWirelessRfProfileRequest {
	this := CreateNetworkWirelessRfProfileRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNetworkWirelessRfProfileRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkWirelessRfProfileRequest) SetName(v string) {
	o.Name = v
}

// GetClientBalancingEnabled returns the ClientBalancingEnabled field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequest) GetClientBalancingEnabled() bool {
	if o == nil || IsNil(o.ClientBalancingEnabled) {
		var ret bool
		return ret
	}
	return *o.ClientBalancingEnabled
}

// GetClientBalancingEnabledOk returns a tuple with the ClientBalancingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequest) GetClientBalancingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientBalancingEnabled) {
		return nil, false
	}
	return o.ClientBalancingEnabled, true
}

// HasClientBalancingEnabled returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequest) HasClientBalancingEnabled() bool {
	if o != nil && !IsNil(o.ClientBalancingEnabled) {
		return true
	}

	return false
}

// SetClientBalancingEnabled gets a reference to the given bool and assigns it to the ClientBalancingEnabled field.
func (o *CreateNetworkWirelessRfProfileRequest) SetClientBalancingEnabled(v bool) {
	o.ClientBalancingEnabled = &v
}

// GetMinBitrateType returns the MinBitrateType field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequest) GetMinBitrateType() string {
	if o == nil || IsNil(o.MinBitrateType) {
		var ret string
		return ret
	}
	return *o.MinBitrateType
}

// GetMinBitrateTypeOk returns a tuple with the MinBitrateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequest) GetMinBitrateTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MinBitrateType) {
		return nil, false
	}
	return o.MinBitrateType, true
}

// HasMinBitrateType returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequest) HasMinBitrateType() bool {
	if o != nil && !IsNil(o.MinBitrateType) {
		return true
	}

	return false
}

// SetMinBitrateType gets a reference to the given string and assigns it to the MinBitrateType field.
func (o *CreateNetworkWirelessRfProfileRequest) SetMinBitrateType(v string) {
	o.MinBitrateType = &v
}

// GetBandSelectionType returns the BandSelectionType field value
func (o *CreateNetworkWirelessRfProfileRequest) GetBandSelectionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BandSelectionType
}

// GetBandSelectionTypeOk returns a tuple with the BandSelectionType field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequest) GetBandSelectionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BandSelectionType, true
}

// SetBandSelectionType sets field value
func (o *CreateNetworkWirelessRfProfileRequest) SetBandSelectionType(v string) {
	o.BandSelectionType = v
}

// GetApBandSettings returns the ApBandSettings field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequest) GetApBandSettings() CreateNetworkWirelessRfProfileRequestApBandSettings {
	if o == nil || IsNil(o.ApBandSettings) {
		var ret CreateNetworkWirelessRfProfileRequestApBandSettings
		return ret
	}
	return *o.ApBandSettings
}

// GetApBandSettingsOk returns a tuple with the ApBandSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequest) GetApBandSettingsOk() (*CreateNetworkWirelessRfProfileRequestApBandSettings, bool) {
	if o == nil || IsNil(o.ApBandSettings) {
		return nil, false
	}
	return o.ApBandSettings, true
}

// HasApBandSettings returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequest) HasApBandSettings() bool {
	if o != nil && !IsNil(o.ApBandSettings) {
		return true
	}

	return false
}

// SetApBandSettings gets a reference to the given CreateNetworkWirelessRfProfileRequestApBandSettings and assigns it to the ApBandSettings field.
func (o *CreateNetworkWirelessRfProfileRequest) SetApBandSettings(v CreateNetworkWirelessRfProfileRequestApBandSettings) {
	o.ApBandSettings = &v
}

// GetTwoFourGhzSettings returns the TwoFourGhzSettings field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequest) GetTwoFourGhzSettings() GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings {
	if o == nil || IsNil(o.TwoFourGhzSettings) {
		var ret GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings
		return ret
	}
	return *o.TwoFourGhzSettings
}

// GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequest) GetTwoFourGhzSettingsOk() (*GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings, bool) {
	if o == nil || IsNil(o.TwoFourGhzSettings) {
		return nil, false
	}
	return o.TwoFourGhzSettings, true
}

// HasTwoFourGhzSettings returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequest) HasTwoFourGhzSettings() bool {
	if o != nil && !IsNil(o.TwoFourGhzSettings) {
		return true
	}

	return false
}

// SetTwoFourGhzSettings gets a reference to the given GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings and assigns it to the TwoFourGhzSettings field.
func (o *CreateNetworkWirelessRfProfileRequest) SetTwoFourGhzSettings(v GetNetworkWirelessRfProfiles200ResponseTwoFourGhzSettings) {
	o.TwoFourGhzSettings = &v
}

// GetFiveGhzSettings returns the FiveGhzSettings field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequest) GetFiveGhzSettings() GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings {
	if o == nil || IsNil(o.FiveGhzSettings) {
		var ret GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings
		return ret
	}
	return *o.FiveGhzSettings
}

// GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequest) GetFiveGhzSettingsOk() (*GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings, bool) {
	if o == nil || IsNil(o.FiveGhzSettings) {
		return nil, false
	}
	return o.FiveGhzSettings, true
}

// HasFiveGhzSettings returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequest) HasFiveGhzSettings() bool {
	if o != nil && !IsNil(o.FiveGhzSettings) {
		return true
	}

	return false
}

// SetFiveGhzSettings gets a reference to the given GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings and assigns it to the FiveGhzSettings field.
func (o *CreateNetworkWirelessRfProfileRequest) SetFiveGhzSettings(v GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) {
	o.FiveGhzSettings = &v
}

// GetSixGhzSettings returns the SixGhzSettings field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequest) GetSixGhzSettings() CreateNetworkWirelessRfProfileRequestSixGhzSettings {
	if o == nil || IsNil(o.SixGhzSettings) {
		var ret CreateNetworkWirelessRfProfileRequestSixGhzSettings
		return ret
	}
	return *o.SixGhzSettings
}

// GetSixGhzSettingsOk returns a tuple with the SixGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequest) GetSixGhzSettingsOk() (*CreateNetworkWirelessRfProfileRequestSixGhzSettings, bool) {
	if o == nil || IsNil(o.SixGhzSettings) {
		return nil, false
	}
	return o.SixGhzSettings, true
}

// HasSixGhzSettings returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequest) HasSixGhzSettings() bool {
	if o != nil && !IsNil(o.SixGhzSettings) {
		return true
	}

	return false
}

// SetSixGhzSettings gets a reference to the given CreateNetworkWirelessRfProfileRequestSixGhzSettings and assigns it to the SixGhzSettings field.
func (o *CreateNetworkWirelessRfProfileRequest) SetSixGhzSettings(v CreateNetworkWirelessRfProfileRequestSixGhzSettings) {
	o.SixGhzSettings = &v
}

// GetTransmission returns the Transmission field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequest) GetTransmission() GetNetworkWirelessRfProfiles200ResponseTransmission {
	if o == nil || IsNil(o.Transmission) {
		var ret GetNetworkWirelessRfProfiles200ResponseTransmission
		return ret
	}
	return *o.Transmission
}

// GetTransmissionOk returns a tuple with the Transmission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequest) GetTransmissionOk() (*GetNetworkWirelessRfProfiles200ResponseTransmission, bool) {
	if o == nil || IsNil(o.Transmission) {
		return nil, false
	}
	return o.Transmission, true
}

// HasTransmission returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequest) HasTransmission() bool {
	if o != nil && !IsNil(o.Transmission) {
		return true
	}

	return false
}

// SetTransmission gets a reference to the given GetNetworkWirelessRfProfiles200ResponseTransmission and assigns it to the Transmission field.
func (o *CreateNetworkWirelessRfProfileRequest) SetTransmission(v GetNetworkWirelessRfProfiles200ResponseTransmission) {
	o.Transmission = &v
}

// GetPerSsidSettings returns the PerSsidSettings field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequest) GetPerSsidSettings() CreateNetworkWirelessRfProfileRequestPerSsidSettings {
	if o == nil || IsNil(o.PerSsidSettings) {
		var ret CreateNetworkWirelessRfProfileRequestPerSsidSettings
		return ret
	}
	return *o.PerSsidSettings
}

// GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequest) GetPerSsidSettingsOk() (*CreateNetworkWirelessRfProfileRequestPerSsidSettings, bool) {
	if o == nil || IsNil(o.PerSsidSettings) {
		return nil, false
	}
	return o.PerSsidSettings, true
}

// HasPerSsidSettings returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequest) HasPerSsidSettings() bool {
	if o != nil && !IsNil(o.PerSsidSettings) {
		return true
	}

	return false
}

// SetPerSsidSettings gets a reference to the given CreateNetworkWirelessRfProfileRequestPerSsidSettings and assigns it to the PerSsidSettings field.
func (o *CreateNetworkWirelessRfProfileRequest) SetPerSsidSettings(v CreateNetworkWirelessRfProfileRequestPerSsidSettings) {
	o.PerSsidSettings = &v
}

// GetFlexRadios returns the FlexRadios field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequest) GetFlexRadios() CreateNetworkWirelessRfProfileRequestFlexRadios {
	if o == nil || IsNil(o.FlexRadios) {
		var ret CreateNetworkWirelessRfProfileRequestFlexRadios
		return ret
	}
	return *o.FlexRadios
}

// GetFlexRadiosOk returns a tuple with the FlexRadios field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequest) GetFlexRadiosOk() (*CreateNetworkWirelessRfProfileRequestFlexRadios, bool) {
	if o == nil || IsNil(o.FlexRadios) {
		return nil, false
	}
	return o.FlexRadios, true
}

// HasFlexRadios returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequest) HasFlexRadios() bool {
	if o != nil && !IsNil(o.FlexRadios) {
		return true
	}

	return false
}

// SetFlexRadios gets a reference to the given CreateNetworkWirelessRfProfileRequestFlexRadios and assigns it to the FlexRadios field.
func (o *CreateNetworkWirelessRfProfileRequest) SetFlexRadios(v CreateNetworkWirelessRfProfileRequestFlexRadios) {
	o.FlexRadios = &v
}

func (o CreateNetworkWirelessRfProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkWirelessRfProfileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.ClientBalancingEnabled) {
		toSerialize["clientBalancingEnabled"] = o.ClientBalancingEnabled
	}
	if !IsNil(o.MinBitrateType) {
		toSerialize["minBitrateType"] = o.MinBitrateType
	}
	toSerialize["bandSelectionType"] = o.BandSelectionType
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

func (o *CreateNetworkWirelessRfProfileRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"bandSelectionType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateNetworkWirelessRfProfileRequest := _CreateNetworkWirelessRfProfileRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNetworkWirelessRfProfileRequest)

	if err != nil {
		return err
	}

	*o = CreateNetworkWirelessRfProfileRequest(varCreateNetworkWirelessRfProfileRequest)

	return err
}

type NullableCreateNetworkWirelessRfProfileRequest struct {
	value *CreateNetworkWirelessRfProfileRequest
	isSet bool
}

func (v NullableCreateNetworkWirelessRfProfileRequest) Get() *CreateNetworkWirelessRfProfileRequest {
	return v.value
}

func (v *NullableCreateNetworkWirelessRfProfileRequest) Set(val *CreateNetworkWirelessRfProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWirelessRfProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWirelessRfProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWirelessRfProfileRequest(val *CreateNetworkWirelessRfProfileRequest) *NullableCreateNetworkWirelessRfProfileRequest {
	return &NullableCreateNetworkWirelessRfProfileRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkWirelessRfProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWirelessRfProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


