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

// CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner struct for CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner
type CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner struct {
	// Message related to whether or not the device was found and can be imported.
	Message *string `json:"message,omitempty"`
	// Device UDI certificate
	Udi *string `json:"udi,omitempty"`
	// Import ID from the Import operation
	DeviceId *string `json:"deviceId,omitempty"`
	// The import status of the device
	Status *string `json:"status,omitempty"`
	ConfigParams *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams `json:"configParams,omitempty"`
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner{}
	return &this
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerWithDefaults instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerWithDefaults() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) SetMessage(v string) {
	o.Message = &v
}

// GetUdi returns the Udi field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) GetUdi() string {
	if o == nil || isNil(o.Udi) {
		var ret string
		return ret
	}
	return *o.Udi
}

// GetUdiOk returns a tuple with the Udi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) GetUdiOk() (*string, bool) {
	if o == nil || isNil(o.Udi) {
    return nil, false
	}
	return o.Udi, true
}

// HasUdi returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) HasUdi() bool {
	if o != nil && !isNil(o.Udi) {
		return true
	}

	return false
}

// SetUdi gets a reference to the given string and assigns it to the Udi field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) SetUdi(v string) {
	o.Udi = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) GetDeviceId() string {
	if o == nil || isNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) GetDeviceIdOk() (*string, bool) {
	if o == nil || isNil(o.DeviceId) {
    return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) HasDeviceId() bool {
	if o != nil && !isNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) SetStatus(v string) {
	o.Status = &v
}

// GetConfigParams returns the ConfigParams field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) GetConfigParams() CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams {
	if o == nil || isNil(o.ConfigParams) {
		var ret CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams
		return ret
	}
	return *o.ConfigParams
}

// GetConfigParamsOk returns a tuple with the ConfigParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) GetConfigParamsOk() (*CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams, bool) {
	if o == nil || isNil(o.ConfigParams) {
    return nil, false
	}
	return o.ConfigParams, true
}

// HasConfigParams returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) HasConfigParams() bool {
	if o != nil && !isNil(o.ConfigParams) {
		return true
	}

	return false
}

// SetConfigParams gets a reference to the given CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams and assigns it to the ConfigParams field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) SetConfigParams(v CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInnerConfigParams) {
	o.ConfigParams = &v
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Udi) {
		toSerialize["udi"] = o.Udi
	}
	if !isNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.ConfigParams) {
		toSerialize["configParams"] = o.ConfigParams
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner struct {
	value *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner
	isSet bool
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) Get() *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner {
	return v.value
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) Set(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner(val *CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner {
	return &NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner{value: val, isSet: true}
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


