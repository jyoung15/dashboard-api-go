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

// checks if the GetOrganizationActionBatches200ResponseInnerStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationActionBatches200ResponseInnerStatus{}

// GetOrganizationActionBatches200ResponseInnerStatus Status of action batch
type GetOrganizationActionBatches200ResponseInnerStatus struct {
	// Flag describing whether all actions in the action batch have completed
	Completed *bool `json:"completed,omitempty"`
	// Flag describing whether any actions in the action batch failed
	Failed *bool `json:"failed,omitempty"`
	// List of errors encountered when running actions in the action batch
	Errors []string `json:"errors,omitempty"`
	// Resources created as a result of this action batch
	CreatedResources []GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner `json:"createdResources"`
}

type _GetOrganizationActionBatches200ResponseInnerStatus GetOrganizationActionBatches200ResponseInnerStatus

// NewGetOrganizationActionBatches200ResponseInnerStatus instantiates a new GetOrganizationActionBatches200ResponseInnerStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationActionBatches200ResponseInnerStatus(createdResources []GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner) *GetOrganizationActionBatches200ResponseInnerStatus {
	this := GetOrganizationActionBatches200ResponseInnerStatus{}
	this.CreatedResources = createdResources
	return &this
}

// NewGetOrganizationActionBatches200ResponseInnerStatusWithDefaults instantiates a new GetOrganizationActionBatches200ResponseInnerStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationActionBatches200ResponseInnerStatusWithDefaults() *GetOrganizationActionBatches200ResponseInnerStatus {
	this := GetOrganizationActionBatches200ResponseInnerStatus{}
	return &this
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *GetOrganizationActionBatches200ResponseInnerStatus) GetCompleted() bool {
	if o == nil || IsNil(o.Completed) {
		var ret bool
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationActionBatches200ResponseInnerStatus) GetCompletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *GetOrganizationActionBatches200ResponseInnerStatus) HasCompleted() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given bool and assigns it to the Completed field.
func (o *GetOrganizationActionBatches200ResponseInnerStatus) SetCompleted(v bool) {
	o.Completed = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *GetOrganizationActionBatches200ResponseInnerStatus) GetFailed() bool {
	if o == nil || IsNil(o.Failed) {
		var ret bool
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationActionBatches200ResponseInnerStatus) GetFailedOk() (*bool, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *GetOrganizationActionBatches200ResponseInnerStatus) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given bool and assigns it to the Failed field.
func (o *GetOrganizationActionBatches200ResponseInnerStatus) SetFailed(v bool) {
	o.Failed = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetOrganizationActionBatches200ResponseInnerStatus) GetErrors() []string {
	if o == nil || IsNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationActionBatches200ResponseInnerStatus) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetOrganizationActionBatches200ResponseInnerStatus) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *GetOrganizationActionBatches200ResponseInnerStatus) SetErrors(v []string) {
	o.Errors = v
}

// GetCreatedResources returns the CreatedResources field value
func (o *GetOrganizationActionBatches200ResponseInnerStatus) GetCreatedResources() []GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner {
	if o == nil {
		var ret []GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner
		return ret
	}

	return o.CreatedResources
}

// GetCreatedResourcesOk returns a tuple with the CreatedResources field value
// and a boolean to check if the value has been set.
func (o *GetOrganizationActionBatches200ResponseInnerStatus) GetCreatedResourcesOk() ([]GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedResources, true
}

// SetCreatedResources sets field value
func (o *GetOrganizationActionBatches200ResponseInnerStatus) SetCreatedResources(v []GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner) {
	o.CreatedResources = v
}

func (o GetOrganizationActionBatches200ResponseInnerStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationActionBatches200ResponseInnerStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if !IsNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	toSerialize["createdResources"] = o.CreatedResources
	return toSerialize, nil
}

func (o *GetOrganizationActionBatches200ResponseInnerStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdResources",
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

	varGetOrganizationActionBatches200ResponseInnerStatus := _GetOrganizationActionBatches200ResponseInnerStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetOrganizationActionBatches200ResponseInnerStatus)

	if err != nil {
		return err
	}

	*o = GetOrganizationActionBatches200ResponseInnerStatus(varGetOrganizationActionBatches200ResponseInnerStatus)

	return err
}

type NullableGetOrganizationActionBatches200ResponseInnerStatus struct {
	value *GetOrganizationActionBatches200ResponseInnerStatus
	isSet bool
}

func (v NullableGetOrganizationActionBatches200ResponseInnerStatus) Get() *GetOrganizationActionBatches200ResponseInnerStatus {
	return v.value
}

func (v *NullableGetOrganizationActionBatches200ResponseInnerStatus) Set(val *GetOrganizationActionBatches200ResponseInnerStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationActionBatches200ResponseInnerStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationActionBatches200ResponseInnerStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationActionBatches200ResponseInnerStatus(val *GetOrganizationActionBatches200ResponseInnerStatus) *NullableGetOrganizationActionBatches200ResponseInnerStatus {
	return &NullableGetOrganizationActionBatches200ResponseInnerStatus{value: val, isSet: true}
}

func (v NullableGetOrganizationActionBatches200ResponseInnerStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationActionBatches200ResponseInnerStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

