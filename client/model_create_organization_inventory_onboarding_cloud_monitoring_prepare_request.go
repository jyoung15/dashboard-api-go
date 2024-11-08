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

// checks if the CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest{}

// CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest struct for CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest
type CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest struct {
	// A set of devices to import (or update)
	Devices []CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner `json:"devices"`
}

type _CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest(devices []CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest{}
	this.Devices = devices
	return &this
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestWithDefaults instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestWithDefaults() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest{}
	return &this
}

// GetDevices returns the Devices field value
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) GetDevices() []CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner {
	if o == nil {
		var ret []CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner
		return ret
	}

	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) GetDevicesOk() ([]CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Devices, true
}

// SetDevices sets field value
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) SetDevices(v []CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner) {
	o.Devices = v
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["devices"] = o.Devices
	return toSerialize, nil
}

func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"devices",
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

	varCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest := _CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest)

	if err != nil {
		return err
	}

	*o = CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest(varCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest)

	return err
}

type NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest struct {
	value *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest
	isSet bool
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) Get() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest {
	return v.value
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) Set(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest {
	return &NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


