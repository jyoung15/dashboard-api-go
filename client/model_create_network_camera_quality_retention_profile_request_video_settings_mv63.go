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

// checks if the CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63{}

// CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63 Quality and resolution for MV63 camera models.
type CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63 struct {
	// Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Quality string `json:"quality"`
	// Resolution of the camera. Can be one of '1920x1080' or '2688x1512'.
	Resolution string `json:"resolution"`
}

type _CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63 CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63

// NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63 instantiates a new CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63(quality string, resolution string) *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63 {
	this := CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63{}
	this.Quality = quality
	this.Resolution = resolution
	return &this
}

// NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63WithDefaults instantiates a new CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63WithDefaults() *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63 {
	this := CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63{}
	return &this
}

// GetQuality returns the Quality field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63) GetQuality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quality
}

// GetQualityOk returns a tuple with the Quality field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63) GetQualityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quality, true
}

// SetQuality sets field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63) SetQuality(v string) {
	o.Quality = v
}

// GetResolution returns the Resolution field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63) GetResolution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63) GetResolutionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resolution, true
}

// SetResolution sets field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63) SetResolution(v string) {
	o.Resolution = v
}

func (o CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["quality"] = o.Quality
	toSerialize["resolution"] = o.Resolution
	return toSerialize, nil
}

func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63) UnmarshalJSON(data []byte) (err error) {
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

	varCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63 := _CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63)

	if err != nil {
		return err
	}

	*o = CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63(varCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63)

	return err
}

type NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63 struct {
	value *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63
	isSet bool
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63) Get() *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63 {
	return v.value
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63) Set(val *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63(val *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63) *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63 {
	return &NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63{value: val, isSet: true}
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


