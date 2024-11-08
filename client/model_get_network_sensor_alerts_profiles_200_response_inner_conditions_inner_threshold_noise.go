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

// checks if the GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise{}

// GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise Noise threshold. 'ambient' must be provided.
type GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise struct {
	Ambient GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient `json:"ambient"`
}

type _GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise(ambient GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient) *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise{}
	this.Ambient = ambient
	return &this
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseWithDefaults instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseWithDefaults() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise{}
	return &this
}

// GetAmbient returns the Ambient field value
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise) GetAmbient() GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient {
	if o == nil {
		var ret GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient
		return ret
	}

	return o.Ambient
}

// GetAmbientOk returns a tuple with the Ambient field value
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise) GetAmbientOk() (*GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ambient, true
}

// SetAmbient sets field value
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise) SetAmbient(v GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoiseAmbient) {
	o.Ambient = v
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ambient"] = o.Ambient
	return toSerialize, nil
}

func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ambient",
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

	varGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise := _GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise)

	if err != nil {
		return err
	}

	*o = GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise(varGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise)

	return err
}

type NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise struct {
	value *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise
	isSet bool
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise) Get() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise {
	return v.value
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise) Set(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise) *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise {
	return &NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise{value: val, isSet: true}
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdNoise) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


