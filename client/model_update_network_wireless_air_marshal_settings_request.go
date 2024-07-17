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

// checks if the UpdateNetworkWirelessAirMarshalSettingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessAirMarshalSettingsRequest{}

// UpdateNetworkWirelessAirMarshalSettingsRequest struct for UpdateNetworkWirelessAirMarshalSettingsRequest
type UpdateNetworkWirelessAirMarshalSettingsRequest struct {
	// Allows clients to access rogue networks. Blocked by default.
	DefaultPolicy string `json:"defaultPolicy"`
}

type _UpdateNetworkWirelessAirMarshalSettingsRequest UpdateNetworkWirelessAirMarshalSettingsRequest

// NewUpdateNetworkWirelessAirMarshalSettingsRequest instantiates a new UpdateNetworkWirelessAirMarshalSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessAirMarshalSettingsRequest(defaultPolicy string) *UpdateNetworkWirelessAirMarshalSettingsRequest {
	this := UpdateNetworkWirelessAirMarshalSettingsRequest{}
	this.DefaultPolicy = defaultPolicy
	return &this
}

// NewUpdateNetworkWirelessAirMarshalSettingsRequestWithDefaults instantiates a new UpdateNetworkWirelessAirMarshalSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessAirMarshalSettingsRequestWithDefaults() *UpdateNetworkWirelessAirMarshalSettingsRequest {
	this := UpdateNetworkWirelessAirMarshalSettingsRequest{}
	return &this
}

// GetDefaultPolicy returns the DefaultPolicy field value
func (o *UpdateNetworkWirelessAirMarshalSettingsRequest) GetDefaultPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultPolicy
}

// GetDefaultPolicyOk returns a tuple with the DefaultPolicy field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessAirMarshalSettingsRequest) GetDefaultPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultPolicy, true
}

// SetDefaultPolicy sets field value
func (o *UpdateNetworkWirelessAirMarshalSettingsRequest) SetDefaultPolicy(v string) {
	o.DefaultPolicy = v
}

func (o UpdateNetworkWirelessAirMarshalSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessAirMarshalSettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["defaultPolicy"] = o.DefaultPolicy
	return toSerialize, nil
}

func (o *UpdateNetworkWirelessAirMarshalSettingsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"defaultPolicy",
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

	varUpdateNetworkWirelessAirMarshalSettingsRequest := _UpdateNetworkWirelessAirMarshalSettingsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateNetworkWirelessAirMarshalSettingsRequest)

	if err != nil {
		return err
	}

	*o = UpdateNetworkWirelessAirMarshalSettingsRequest(varUpdateNetworkWirelessAirMarshalSettingsRequest)

	return err
}

type NullableUpdateNetworkWirelessAirMarshalSettingsRequest struct {
	value *UpdateNetworkWirelessAirMarshalSettingsRequest
	isSet bool
}

func (v NullableUpdateNetworkWirelessAirMarshalSettingsRequest) Get() *UpdateNetworkWirelessAirMarshalSettingsRequest {
	return v.value
}

func (v *NullableUpdateNetworkWirelessAirMarshalSettingsRequest) Set(val *UpdateNetworkWirelessAirMarshalSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessAirMarshalSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessAirMarshalSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessAirMarshalSettingsRequest(val *UpdateNetworkWirelessAirMarshalSettingsRequest) *NullableUpdateNetworkWirelessAirMarshalSettingsRequest {
	return &NullableUpdateNetworkWirelessAirMarshalSettingsRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessAirMarshalSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessAirMarshalSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


