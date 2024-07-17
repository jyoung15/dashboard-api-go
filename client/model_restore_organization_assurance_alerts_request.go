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

// checks if the RestoreOrganizationAssuranceAlertsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RestoreOrganizationAssuranceAlertsRequest{}

// RestoreOrganizationAssuranceAlertsRequest struct for RestoreOrganizationAssuranceAlertsRequest
type RestoreOrganizationAssuranceAlertsRequest struct {
	// Array of alert IDs to restore
	AlertIds []string `json:"alertIds"`
}

type _RestoreOrganizationAssuranceAlertsRequest RestoreOrganizationAssuranceAlertsRequest

// NewRestoreOrganizationAssuranceAlertsRequest instantiates a new RestoreOrganizationAssuranceAlertsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestoreOrganizationAssuranceAlertsRequest(alertIds []string) *RestoreOrganizationAssuranceAlertsRequest {
	this := RestoreOrganizationAssuranceAlertsRequest{}
	this.AlertIds = alertIds
	return &this
}

// NewRestoreOrganizationAssuranceAlertsRequestWithDefaults instantiates a new RestoreOrganizationAssuranceAlertsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestoreOrganizationAssuranceAlertsRequestWithDefaults() *RestoreOrganizationAssuranceAlertsRequest {
	this := RestoreOrganizationAssuranceAlertsRequest{}
	return &this
}

// GetAlertIds returns the AlertIds field value
func (o *RestoreOrganizationAssuranceAlertsRequest) GetAlertIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AlertIds
}

// GetAlertIdsOk returns a tuple with the AlertIds field value
// and a boolean to check if the value has been set.
func (o *RestoreOrganizationAssuranceAlertsRequest) GetAlertIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlertIds, true
}

// SetAlertIds sets field value
func (o *RestoreOrganizationAssuranceAlertsRequest) SetAlertIds(v []string) {
	o.AlertIds = v
}

func (o RestoreOrganizationAssuranceAlertsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestoreOrganizationAssuranceAlertsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alertIds"] = o.AlertIds
	return toSerialize, nil
}

func (o *RestoreOrganizationAssuranceAlertsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alertIds",
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

	varRestoreOrganizationAssuranceAlertsRequest := _RestoreOrganizationAssuranceAlertsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRestoreOrganizationAssuranceAlertsRequest)

	if err != nil {
		return err
	}

	*o = RestoreOrganizationAssuranceAlertsRequest(varRestoreOrganizationAssuranceAlertsRequest)

	return err
}

type NullableRestoreOrganizationAssuranceAlertsRequest struct {
	value *RestoreOrganizationAssuranceAlertsRequest
	isSet bool
}

func (v NullableRestoreOrganizationAssuranceAlertsRequest) Get() *RestoreOrganizationAssuranceAlertsRequest {
	return v.value
}

func (v *NullableRestoreOrganizationAssuranceAlertsRequest) Set(val *RestoreOrganizationAssuranceAlertsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRestoreOrganizationAssuranceAlertsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRestoreOrganizationAssuranceAlertsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestoreOrganizationAssuranceAlertsRequest(val *RestoreOrganizationAssuranceAlertsRequest) *NullableRestoreOrganizationAssuranceAlertsRequest {
	return &NullableRestoreOrganizationAssuranceAlertsRequest{value: val, isSet: true}
}

func (v NullableRestoreOrganizationAssuranceAlertsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestoreOrganizationAssuranceAlertsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


