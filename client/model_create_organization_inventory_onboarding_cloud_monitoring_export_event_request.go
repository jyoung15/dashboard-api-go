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

// checks if the CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest{}

// CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest struct for CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest
type CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest struct {
	// The type of log event this is recording, e.g. download or opening a banner
	LogEvent string `json:"logEvent"`
	// A JavaScript UTC datetime stamp for when the even occurred
	Timestamp int32 `json:"timestamp"`
	// The name of the onboarding distro being downloaded
	TargetOS *string `json:"targetOS,omitempty"`
	// Used to describe if this event was the result of a redirect. E.g. a query param if an info banner is being used
	Request *string `json:"request,omitempty"`
}

type _CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest

// NewCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest(logEvent string, timestamp int32) *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest{}
	this.LogEvent = logEvent
	this.Timestamp = timestamp
	return &this
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequestWithDefaults instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequestWithDefaults() *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest{}
	return &this
}

// GetLogEvent returns the LogEvent field value
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) GetLogEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogEvent
}

// GetLogEventOk returns a tuple with the LogEvent field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) GetLogEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogEvent, true
}

// SetLogEvent sets field value
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) SetLogEvent(v string) {
	o.LogEvent = v
}

// GetTimestamp returns the Timestamp field value
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetTargetOS returns the TargetOS field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) GetTargetOS() string {
	if o == nil || IsNil(o.TargetOS) {
		var ret string
		return ret
	}
	return *o.TargetOS
}

// GetTargetOSOk returns a tuple with the TargetOS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) GetTargetOSOk() (*string, bool) {
	if o == nil || IsNil(o.TargetOS) {
		return nil, false
	}
	return o.TargetOS, true
}

// HasTargetOS returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) HasTargetOS() bool {
	if o != nil && !IsNil(o.TargetOS) {
		return true
	}

	return false
}

// SetTargetOS gets a reference to the given string and assigns it to the TargetOS field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) SetTargetOS(v string) {
	o.TargetOS = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) GetRequest() string {
	if o == nil || IsNil(o.Request) {
		var ret string
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) GetRequestOk() (*string, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given string and assigns it to the Request field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) SetRequest(v string) {
	o.Request = &v
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["logEvent"] = o.LogEvent
	toSerialize["timestamp"] = o.Timestamp
	if !IsNil(o.TargetOS) {
		toSerialize["targetOS"] = o.TargetOS
	}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	return toSerialize, nil
}

func (o *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"logEvent",
		"timestamp",
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

	varCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest := _CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest)

	if err != nil {
		return err
	}

	*o = CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest(varCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest)

	return err
}

type NullableCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest struct {
	value *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest
	isSet bool
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) Get() *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest {
	return v.value
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) Set(val *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest(val *CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) *NullableCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest {
	return &NullableCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


