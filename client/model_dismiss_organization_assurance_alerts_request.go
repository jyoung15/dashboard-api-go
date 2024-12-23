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

// checks if the DismissOrganizationAssuranceAlertsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DismissOrganizationAssuranceAlertsRequest{}

// DismissOrganizationAssuranceAlertsRequest struct for DismissOrganizationAssuranceAlertsRequest
type DismissOrganizationAssuranceAlertsRequest struct {
	// Array of alert IDs to dismiss
	AlertIds []string `json:"alertIds"`
}

type _DismissOrganizationAssuranceAlertsRequest DismissOrganizationAssuranceAlertsRequest

// NewDismissOrganizationAssuranceAlertsRequest instantiates a new DismissOrganizationAssuranceAlertsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDismissOrganizationAssuranceAlertsRequest(alertIds []string) *DismissOrganizationAssuranceAlertsRequest {
	this := DismissOrganizationAssuranceAlertsRequest{}
	this.AlertIds = alertIds
	return &this
}

// NewDismissOrganizationAssuranceAlertsRequestWithDefaults instantiates a new DismissOrganizationAssuranceAlertsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDismissOrganizationAssuranceAlertsRequestWithDefaults() *DismissOrganizationAssuranceAlertsRequest {
	this := DismissOrganizationAssuranceAlertsRequest{}
	return &this
}

// GetAlertIds returns the AlertIds field value
func (o *DismissOrganizationAssuranceAlertsRequest) GetAlertIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AlertIds
}

// GetAlertIdsOk returns a tuple with the AlertIds field value
// and a boolean to check if the value has been set.
func (o *DismissOrganizationAssuranceAlertsRequest) GetAlertIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlertIds, true
}

// SetAlertIds sets field value
func (o *DismissOrganizationAssuranceAlertsRequest) SetAlertIds(v []string) {
	o.AlertIds = v
}

func (o DismissOrganizationAssuranceAlertsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DismissOrganizationAssuranceAlertsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alertIds"] = o.AlertIds
	return toSerialize, nil
}

func (o *DismissOrganizationAssuranceAlertsRequest) UnmarshalJSON(data []byte) (err error) {
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

	varDismissOrganizationAssuranceAlertsRequest := _DismissOrganizationAssuranceAlertsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDismissOrganizationAssuranceAlertsRequest)

	if err != nil {
		return err
	}

	*o = DismissOrganizationAssuranceAlertsRequest(varDismissOrganizationAssuranceAlertsRequest)

	return err
}

type NullableDismissOrganizationAssuranceAlertsRequest struct {
	value *DismissOrganizationAssuranceAlertsRequest
	isSet bool
}

func (v NullableDismissOrganizationAssuranceAlertsRequest) Get() *DismissOrganizationAssuranceAlertsRequest {
	return v.value
}

func (v *NullableDismissOrganizationAssuranceAlertsRequest) Set(val *DismissOrganizationAssuranceAlertsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDismissOrganizationAssuranceAlertsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDismissOrganizationAssuranceAlertsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDismissOrganizationAssuranceAlertsRequest(val *DismissOrganizationAssuranceAlertsRequest) *NullableDismissOrganizationAssuranceAlertsRequest {
	return &NullableDismissOrganizationAssuranceAlertsRequest{value: val, isSet: true}
}

func (v NullableDismissOrganizationAssuranceAlertsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDismissOrganizationAssuranceAlertsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


