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

// checks if the CreateNetworkApplianceRfProfileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkApplianceRfProfileRequest{}

// CreateNetworkApplianceRfProfileRequest struct for CreateNetworkApplianceRfProfileRequest
type CreateNetworkApplianceRfProfileRequest struct {
	// The name of the new profile. Must be unique. This param is required on creation.
	Name string `json:"name"`
	TwoFourGhzSettings *CreateNetworkApplianceRfProfileRequestTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`
	FiveGhzSettings *CreateNetworkApplianceRfProfileRequestFiveGhzSettings `json:"fiveGhzSettings,omitempty"`
	PerSsidSettings *CreateNetworkApplianceRfProfileRequestPerSsidSettings `json:"perSsidSettings,omitempty"`
}

type _CreateNetworkApplianceRfProfileRequest CreateNetworkApplianceRfProfileRequest

// NewCreateNetworkApplianceRfProfileRequest instantiates a new CreateNetworkApplianceRfProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkApplianceRfProfileRequest(name string) *CreateNetworkApplianceRfProfileRequest {
	this := CreateNetworkApplianceRfProfileRequest{}
	this.Name = name
	return &this
}

// NewCreateNetworkApplianceRfProfileRequestWithDefaults instantiates a new CreateNetworkApplianceRfProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkApplianceRfProfileRequestWithDefaults() *CreateNetworkApplianceRfProfileRequest {
	this := CreateNetworkApplianceRfProfileRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNetworkApplianceRfProfileRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceRfProfileRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkApplianceRfProfileRequest) SetName(v string) {
	o.Name = v
}

// GetTwoFourGhzSettings returns the TwoFourGhzSettings field value if set, zero value otherwise.
func (o *CreateNetworkApplianceRfProfileRequest) GetTwoFourGhzSettings() CreateNetworkApplianceRfProfileRequestTwoFourGhzSettings {
	if o == nil || IsNil(o.TwoFourGhzSettings) {
		var ret CreateNetworkApplianceRfProfileRequestTwoFourGhzSettings
		return ret
	}
	return *o.TwoFourGhzSettings
}

// GetTwoFourGhzSettingsOk returns a tuple with the TwoFourGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceRfProfileRequest) GetTwoFourGhzSettingsOk() (*CreateNetworkApplianceRfProfileRequestTwoFourGhzSettings, bool) {
	if o == nil || IsNil(o.TwoFourGhzSettings) {
		return nil, false
	}
	return o.TwoFourGhzSettings, true
}

// HasTwoFourGhzSettings returns a boolean if a field has been set.
func (o *CreateNetworkApplianceRfProfileRequest) HasTwoFourGhzSettings() bool {
	if o != nil && !IsNil(o.TwoFourGhzSettings) {
		return true
	}

	return false
}

// SetTwoFourGhzSettings gets a reference to the given CreateNetworkApplianceRfProfileRequestTwoFourGhzSettings and assigns it to the TwoFourGhzSettings field.
func (o *CreateNetworkApplianceRfProfileRequest) SetTwoFourGhzSettings(v CreateNetworkApplianceRfProfileRequestTwoFourGhzSettings) {
	o.TwoFourGhzSettings = &v
}

// GetFiveGhzSettings returns the FiveGhzSettings field value if set, zero value otherwise.
func (o *CreateNetworkApplianceRfProfileRequest) GetFiveGhzSettings() CreateNetworkApplianceRfProfileRequestFiveGhzSettings {
	if o == nil || IsNil(o.FiveGhzSettings) {
		var ret CreateNetworkApplianceRfProfileRequestFiveGhzSettings
		return ret
	}
	return *o.FiveGhzSettings
}

// GetFiveGhzSettingsOk returns a tuple with the FiveGhzSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceRfProfileRequest) GetFiveGhzSettingsOk() (*CreateNetworkApplianceRfProfileRequestFiveGhzSettings, bool) {
	if o == nil || IsNil(o.FiveGhzSettings) {
		return nil, false
	}
	return o.FiveGhzSettings, true
}

// HasFiveGhzSettings returns a boolean if a field has been set.
func (o *CreateNetworkApplianceRfProfileRequest) HasFiveGhzSettings() bool {
	if o != nil && !IsNil(o.FiveGhzSettings) {
		return true
	}

	return false
}

// SetFiveGhzSettings gets a reference to the given CreateNetworkApplianceRfProfileRequestFiveGhzSettings and assigns it to the FiveGhzSettings field.
func (o *CreateNetworkApplianceRfProfileRequest) SetFiveGhzSettings(v CreateNetworkApplianceRfProfileRequestFiveGhzSettings) {
	o.FiveGhzSettings = &v
}

// GetPerSsidSettings returns the PerSsidSettings field value if set, zero value otherwise.
func (o *CreateNetworkApplianceRfProfileRequest) GetPerSsidSettings() CreateNetworkApplianceRfProfileRequestPerSsidSettings {
	if o == nil || IsNil(o.PerSsidSettings) {
		var ret CreateNetworkApplianceRfProfileRequestPerSsidSettings
		return ret
	}
	return *o.PerSsidSettings
}

// GetPerSsidSettingsOk returns a tuple with the PerSsidSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceRfProfileRequest) GetPerSsidSettingsOk() (*CreateNetworkApplianceRfProfileRequestPerSsidSettings, bool) {
	if o == nil || IsNil(o.PerSsidSettings) {
		return nil, false
	}
	return o.PerSsidSettings, true
}

// HasPerSsidSettings returns a boolean if a field has been set.
func (o *CreateNetworkApplianceRfProfileRequest) HasPerSsidSettings() bool {
	if o != nil && !IsNil(o.PerSsidSettings) {
		return true
	}

	return false
}

// SetPerSsidSettings gets a reference to the given CreateNetworkApplianceRfProfileRequestPerSsidSettings and assigns it to the PerSsidSettings field.
func (o *CreateNetworkApplianceRfProfileRequest) SetPerSsidSettings(v CreateNetworkApplianceRfProfileRequestPerSsidSettings) {
	o.PerSsidSettings = &v
}

func (o CreateNetworkApplianceRfProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkApplianceRfProfileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
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

func (o *CreateNetworkApplianceRfProfileRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCreateNetworkApplianceRfProfileRequest := _CreateNetworkApplianceRfProfileRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNetworkApplianceRfProfileRequest)

	if err != nil {
		return err
	}

	*o = CreateNetworkApplianceRfProfileRequest(varCreateNetworkApplianceRfProfileRequest)

	return err
}

type NullableCreateNetworkApplianceRfProfileRequest struct {
	value *CreateNetworkApplianceRfProfileRequest
	isSet bool
}

func (v NullableCreateNetworkApplianceRfProfileRequest) Get() *CreateNetworkApplianceRfProfileRequest {
	return v.value
}

func (v *NullableCreateNetworkApplianceRfProfileRequest) Set(val *CreateNetworkApplianceRfProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkApplianceRfProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkApplianceRfProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkApplianceRfProfileRequest(val *CreateNetworkApplianceRfProfileRequest) *NullableCreateNetworkApplianceRfProfileRequest {
	return &NullableCreateNetworkApplianceRfProfileRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkApplianceRfProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkApplianceRfProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


