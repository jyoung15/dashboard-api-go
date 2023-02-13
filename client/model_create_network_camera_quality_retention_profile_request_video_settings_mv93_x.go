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

// CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X Quality and resolution for MV93X camera models.
type CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X struct {
	// Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Quality string `json:"quality"`
	// Resolution of the camera. Can be one of '1080x1080', '2112x2112' or '2880x2880'.
	Resolution string `json:"resolution"`
}

// NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X instantiates a new CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X(quality string, resolution string) *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X {
	this := CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X{}
	this.Quality = quality
	this.Resolution = resolution
	return &this
}

// NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93XWithDefaults instantiates a new CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93XWithDefaults() *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X {
	this := CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X{}
	return &this
}

// GetQuality returns the Quality field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X) GetQuality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quality
}

// GetQualityOk returns a tuple with the Quality field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X) GetQualityOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Quality, true
}

// SetQuality sets field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X) SetQuality(v string) {
	o.Quality = v
}

// GetResolution returns the Resolution field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X) GetResolution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X) GetResolutionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Resolution, true
}

// SetResolution sets field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X) SetResolution(v string) {
	o.Resolution = v
}

func (o CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["quality"] = o.Quality
	}
	if true {
		toSerialize["resolution"] = o.Resolution
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X struct {
	value *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X
	isSet bool
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X) Get() *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X {
	return v.value
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X) Set(val *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X(val *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X) *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X {
	return &NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X{value: val, isSet: true}
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV93X) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

