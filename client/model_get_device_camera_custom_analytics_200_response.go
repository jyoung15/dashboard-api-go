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

// checks if the GetDeviceCameraCustomAnalytics200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceCameraCustomAnalytics200Response{}

// GetDeviceCameraCustomAnalytics200Response struct for GetDeviceCameraCustomAnalytics200Response
type GetDeviceCameraCustomAnalytics200Response struct {
	// Whether custom analytics is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Custom analytics artifact ID
	ArtifactId *string `json:"artifactId,omitempty"`
	// Parameters for the custom analytics workload
	Parameters []GetDeviceCameraCustomAnalytics200ResponseParametersInner `json:"parameters,omitempty"`
}

// NewGetDeviceCameraCustomAnalytics200Response instantiates a new GetDeviceCameraCustomAnalytics200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceCameraCustomAnalytics200Response() *GetDeviceCameraCustomAnalytics200Response {
	this := GetDeviceCameraCustomAnalytics200Response{}
	return &this
}

// NewGetDeviceCameraCustomAnalytics200ResponseWithDefaults instantiates a new GetDeviceCameraCustomAnalytics200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceCameraCustomAnalytics200ResponseWithDefaults() *GetDeviceCameraCustomAnalytics200Response {
	this := GetDeviceCameraCustomAnalytics200Response{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetDeviceCameraCustomAnalytics200Response) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceCameraCustomAnalytics200Response) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetDeviceCameraCustomAnalytics200Response) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetDeviceCameraCustomAnalytics200Response) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetArtifactId returns the ArtifactId field value if set, zero value otherwise.
func (o *GetDeviceCameraCustomAnalytics200Response) GetArtifactId() string {
	if o == nil || IsNil(o.ArtifactId) {
		var ret string
		return ret
	}
	return *o.ArtifactId
}

// GetArtifactIdOk returns a tuple with the ArtifactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceCameraCustomAnalytics200Response) GetArtifactIdOk() (*string, bool) {
	if o == nil || IsNil(o.ArtifactId) {
		return nil, false
	}
	return o.ArtifactId, true
}

// HasArtifactId returns a boolean if a field has been set.
func (o *GetDeviceCameraCustomAnalytics200Response) HasArtifactId() bool {
	if o != nil && !IsNil(o.ArtifactId) {
		return true
	}

	return false
}

// SetArtifactId gets a reference to the given string and assigns it to the ArtifactId field.
func (o *GetDeviceCameraCustomAnalytics200Response) SetArtifactId(v string) {
	o.ArtifactId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *GetDeviceCameraCustomAnalytics200Response) GetParameters() []GetDeviceCameraCustomAnalytics200ResponseParametersInner {
	if o == nil || IsNil(o.Parameters) {
		var ret []GetDeviceCameraCustomAnalytics200ResponseParametersInner
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceCameraCustomAnalytics200Response) GetParametersOk() ([]GetDeviceCameraCustomAnalytics200ResponseParametersInner, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *GetDeviceCameraCustomAnalytics200Response) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []GetDeviceCameraCustomAnalytics200ResponseParametersInner and assigns it to the Parameters field.
func (o *GetDeviceCameraCustomAnalytics200Response) SetParameters(v []GetDeviceCameraCustomAnalytics200ResponseParametersInner) {
	o.Parameters = v
}

func (o GetDeviceCameraCustomAnalytics200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceCameraCustomAnalytics200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ArtifactId) {
		toSerialize["artifactId"] = o.ArtifactId
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	return toSerialize, nil
}

type NullableGetDeviceCameraCustomAnalytics200Response struct {
	value *GetDeviceCameraCustomAnalytics200Response
	isSet bool
}

func (v NullableGetDeviceCameraCustomAnalytics200Response) Get() *GetDeviceCameraCustomAnalytics200Response {
	return v.value
}

func (v *NullableGetDeviceCameraCustomAnalytics200Response) Set(val *GetDeviceCameraCustomAnalytics200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceCameraCustomAnalytics200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceCameraCustomAnalytics200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceCameraCustomAnalytics200Response(val *GetDeviceCameraCustomAnalytics200Response) *NullableGetDeviceCameraCustomAnalytics200Response {
	return &NullableGetDeviceCameraCustomAnalytics200Response{value: val, isSet: true}
}

func (v NullableGetDeviceCameraCustomAnalytics200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceCameraCustomAnalytics200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


