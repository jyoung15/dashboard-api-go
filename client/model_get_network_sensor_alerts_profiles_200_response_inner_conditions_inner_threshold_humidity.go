/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity Humidity threshold. One of 'relativePercentage' or 'quality' must be provided.
type GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity struct {
	// Alerting threshold in %RH.
	RelativePercentage *int32 `json:"relativePercentage,omitempty"`
	// Alerting threshold as a qualitative humidity level.
	Quality *string `json:"quality,omitempty"`
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity{}
	return &this
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidityWithDefaults instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidityWithDefaults() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity{}
	return &this
}

// GetRelativePercentage returns the RelativePercentage field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity) GetRelativePercentage() int32 {
	if o == nil || isNil(o.RelativePercentage) {
		var ret int32
		return ret
	}
	return *o.RelativePercentage
}

// GetRelativePercentageOk returns a tuple with the RelativePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity) GetRelativePercentageOk() (*int32, bool) {
	if o == nil || isNil(o.RelativePercentage) {
    return nil, false
	}
	return o.RelativePercentage, true
}

// HasRelativePercentage returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity) HasRelativePercentage() bool {
	if o != nil && !isNil(o.RelativePercentage) {
		return true
	}

	return false
}

// SetRelativePercentage gets a reference to the given int32 and assigns it to the RelativePercentage field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity) SetRelativePercentage(v int32) {
	o.RelativePercentage = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity) GetQuality() string {
	if o == nil || isNil(o.Quality) {
		var ret string
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity) GetQualityOk() (*string, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given string and assigns it to the Quality field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity) SetQuality(v string) {
	o.Quality = &v
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RelativePercentage) {
		toSerialize["relativePercentage"] = o.RelativePercentage
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity struct {
	value *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity
	isSet bool
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity) Get() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity {
	return v.value
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity) Set(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity) *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity {
	return &NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity{value: val, isSet: true}
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdHumidity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


