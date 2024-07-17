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

// checks if the CreateNetworkApplianceRfProfileRequestPerSsidSettings1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkApplianceRfProfileRequestPerSsidSettings1{}

// CreateNetworkApplianceRfProfileRequestPerSsidSettings1 Settings for SSID 1
type CreateNetworkApplianceRfProfileRequestPerSsidSettings1 struct {
	// Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandOperationMode *string `json:"bandOperationMode,omitempty"`
	// Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	BandSteeringEnabled *bool `json:"bandSteeringEnabled,omitempty"`
}

// NewCreateNetworkApplianceRfProfileRequestPerSsidSettings1 instantiates a new CreateNetworkApplianceRfProfileRequestPerSsidSettings1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkApplianceRfProfileRequestPerSsidSettings1() *CreateNetworkApplianceRfProfileRequestPerSsidSettings1 {
	this := CreateNetworkApplianceRfProfileRequestPerSsidSettings1{}
	return &this
}

// NewCreateNetworkApplianceRfProfileRequestPerSsidSettings1WithDefaults instantiates a new CreateNetworkApplianceRfProfileRequestPerSsidSettings1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkApplianceRfProfileRequestPerSsidSettings1WithDefaults() *CreateNetworkApplianceRfProfileRequestPerSsidSettings1 {
	this := CreateNetworkApplianceRfProfileRequestPerSsidSettings1{}
	return &this
}

// GetBandOperationMode returns the BandOperationMode field value if set, zero value otherwise.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings1) GetBandOperationMode() string {
	if o == nil || IsNil(o.BandOperationMode) {
		var ret string
		return ret
	}
	return *o.BandOperationMode
}

// GetBandOperationModeOk returns a tuple with the BandOperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings1) GetBandOperationModeOk() (*string, bool) {
	if o == nil || IsNil(o.BandOperationMode) {
		return nil, false
	}
	return o.BandOperationMode, true
}

// HasBandOperationMode returns a boolean if a field has been set.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings1) HasBandOperationMode() bool {
	if o != nil && !IsNil(o.BandOperationMode) {
		return true
	}

	return false
}

// SetBandOperationMode gets a reference to the given string and assigns it to the BandOperationMode field.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings1) SetBandOperationMode(v string) {
	o.BandOperationMode = &v
}

// GetBandSteeringEnabled returns the BandSteeringEnabled field value if set, zero value otherwise.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings1) GetBandSteeringEnabled() bool {
	if o == nil || IsNil(o.BandSteeringEnabled) {
		var ret bool
		return ret
	}
	return *o.BandSteeringEnabled
}

// GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings1) GetBandSteeringEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BandSteeringEnabled) {
		return nil, false
	}
	return o.BandSteeringEnabled, true
}

// HasBandSteeringEnabled returns a boolean if a field has been set.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings1) HasBandSteeringEnabled() bool {
	if o != nil && !IsNil(o.BandSteeringEnabled) {
		return true
	}

	return false
}

// SetBandSteeringEnabled gets a reference to the given bool and assigns it to the BandSteeringEnabled field.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings1) SetBandSteeringEnabled(v bool) {
	o.BandSteeringEnabled = &v
}

func (o CreateNetworkApplianceRfProfileRequestPerSsidSettings1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkApplianceRfProfileRequestPerSsidSettings1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BandOperationMode) {
		toSerialize["bandOperationMode"] = o.BandOperationMode
	}
	if !IsNil(o.BandSteeringEnabled) {
		toSerialize["bandSteeringEnabled"] = o.BandSteeringEnabled
	}
	return toSerialize, nil
}

type NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings1 struct {
	value *CreateNetworkApplianceRfProfileRequestPerSsidSettings1
	isSet bool
}

func (v NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings1) Get() *CreateNetworkApplianceRfProfileRequestPerSsidSettings1 {
	return v.value
}

func (v *NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings1) Set(val *CreateNetworkApplianceRfProfileRequestPerSsidSettings1) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings1) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkApplianceRfProfileRequestPerSsidSettings1(val *CreateNetworkApplianceRfProfileRequestPerSsidSettings1) *NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings1 {
	return &NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings1{value: val, isSet: true}
}

func (v NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


