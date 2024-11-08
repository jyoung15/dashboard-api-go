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

// checks if the CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE{}

// CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE Quality and resolution for MV12WE camera models.
type CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE struct {
	// Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Quality string `json:"quality"`
	// Resolution of the camera. Can be one of '1280x720' or '1920x1080'.
	Resolution string `json:"resolution"`
}

type _CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE

// NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE instantiates a new CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE(quality string, resolution string) *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE {
	this := CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE{}
	this.Quality = quality
	this.Resolution = resolution
	return &this
}

// NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WEWithDefaults instantiates a new CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WEWithDefaults() *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE {
	this := CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE{}
	return &this
}

// GetQuality returns the Quality field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE) GetQuality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quality
}

// GetQualityOk returns a tuple with the Quality field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE) GetQualityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quality, true
}

// SetQuality sets field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE) SetQuality(v string) {
	o.Quality = v
}

// GetResolution returns the Resolution field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE) GetResolution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE) GetResolutionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resolution, true
}

// SetResolution sets field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE) SetResolution(v string) {
	o.Resolution = v
}

func (o CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["quality"] = o.Quality
	toSerialize["resolution"] = o.Resolution
	return toSerialize, nil
}

func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"quality",
		"resolution",
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

	varCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE := _CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE)

	if err != nil {
		return err
	}

	*o = CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE(varCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE)

	return err
}

type NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE struct {
	value *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE
	isSet bool
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE) Get() *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE {
	return v.value
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE) Set(val *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE(val *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE) *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE {
	return &NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE{value: val, isSet: true}
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV12WE) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


