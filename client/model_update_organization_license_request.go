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

// UpdateOrganizationLicenseRequest struct for UpdateOrganizationLicenseRequest
type UpdateOrganizationLicenseRequest struct {
	// The serial number of the device to assign this license to. Set this to  null to unassign the license. If a different license is already active on the device, this parameter will control queueing/dequeuing this license.
	DeviceSerial *string `json:"deviceSerial,omitempty"`
}

// NewUpdateOrganizationLicenseRequest instantiates a new UpdateOrganizationLicenseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationLicenseRequest() *UpdateOrganizationLicenseRequest {
	this := UpdateOrganizationLicenseRequest{}
	return &this
}

// NewUpdateOrganizationLicenseRequestWithDefaults instantiates a new UpdateOrganizationLicenseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationLicenseRequestWithDefaults() *UpdateOrganizationLicenseRequest {
	this := UpdateOrganizationLicenseRequest{}
	return &this
}

// GetDeviceSerial returns the DeviceSerial field value if set, zero value otherwise.
func (o *UpdateOrganizationLicenseRequest) GetDeviceSerial() string {
	if o == nil || isNil(o.DeviceSerial) {
		var ret string
		return ret
	}
	return *o.DeviceSerial
}

// GetDeviceSerialOk returns a tuple with the DeviceSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationLicenseRequest) GetDeviceSerialOk() (*string, bool) {
	if o == nil || isNil(o.DeviceSerial) {
    return nil, false
	}
	return o.DeviceSerial, true
}

// HasDeviceSerial returns a boolean if a field has been set.
func (o *UpdateOrganizationLicenseRequest) HasDeviceSerial() bool {
	if o != nil && !isNil(o.DeviceSerial) {
		return true
	}

	return false
}

// SetDeviceSerial gets a reference to the given string and assigns it to the DeviceSerial field.
func (o *UpdateOrganizationLicenseRequest) SetDeviceSerial(v string) {
	o.DeviceSerial = &v
}

func (o UpdateOrganizationLicenseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DeviceSerial) {
		toSerialize["deviceSerial"] = o.DeviceSerial
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateOrganizationLicenseRequest struct {
	value *UpdateOrganizationLicenseRequest
	isSet bool
}

func (v NullableUpdateOrganizationLicenseRequest) Get() *UpdateOrganizationLicenseRequest {
	return v.value
}

func (v *NullableUpdateOrganizationLicenseRequest) Set(val *UpdateOrganizationLicenseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationLicenseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationLicenseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationLicenseRequest(val *UpdateOrganizationLicenseRequest) *NullableUpdateOrganizationLicenseRequest {
	return &NullableUpdateOrganizationLicenseRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationLicenseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationLicenseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


