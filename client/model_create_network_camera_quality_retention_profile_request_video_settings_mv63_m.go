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

// checks if the CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M{}

// CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M Quality and resolution for MV63M camera models.
type CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M struct {
	// Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Quality string `json:"quality"`
	// Resolution of the camera. Can be one of '1920x1080' or '2688x1512'.
	Resolution string `json:"resolution"`
}

type _CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M

// NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M instantiates a new CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M(quality string, resolution string) *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M {
	this := CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M{}
	this.Quality = quality
	this.Resolution = resolution
	return &this
}

// NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63MWithDefaults instantiates a new CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63MWithDefaults() *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M {
	this := CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M{}
	return &this
}

// GetQuality returns the Quality field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M) GetQuality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quality
}

// GetQualityOk returns a tuple with the Quality field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M) GetQualityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quality, true
}

// SetQuality sets field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M) SetQuality(v string) {
	o.Quality = v
}

// GetResolution returns the Resolution field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M) GetResolution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M) GetResolutionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resolution, true
}

// SetResolution sets field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M) SetResolution(v string) {
	o.Resolution = v
}

func (o CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["quality"] = o.Quality
	toSerialize["resolution"] = o.Resolution
	return toSerialize, nil
}

func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M) UnmarshalJSON(data []byte) (err error) {
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

	varCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M := _CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M)

	if err != nil {
		return err
	}

	*o = CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M(varCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M)

	return err
}

type NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M struct {
	value *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M
	isSet bool
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M) Get() *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M {
	return v.value
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M) Set(val *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M(val *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M) *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M {
	return &NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M{value: val, isSet: true}
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV63M) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


