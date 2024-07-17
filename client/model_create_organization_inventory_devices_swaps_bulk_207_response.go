/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CreateOrganizationInventoryDevicesSwapsBulk207Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationInventoryDevicesSwapsBulk207Response{}

// CreateOrganizationInventoryDevicesSwapsBulk207Response struct for CreateOrganizationInventoryDevicesSwapsBulk207Response
type CreateOrganizationInventoryDevicesSwapsBulk207Response struct {
	// The ID of the job that was used to create all of the device swaps.
	JobId *string `json:"jobId,omitempty"`
	// An array of recent swap requests and their statuses.
	Swaps []CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner `json:"swaps,omitempty"`
}

// NewCreateOrganizationInventoryDevicesSwapsBulk207Response instantiates a new CreateOrganizationInventoryDevicesSwapsBulk207Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInventoryDevicesSwapsBulk207Response() *CreateOrganizationInventoryDevicesSwapsBulk207Response {
	this := CreateOrganizationInventoryDevicesSwapsBulk207Response{}
	return &this
}

// NewCreateOrganizationInventoryDevicesSwapsBulk207ResponseWithDefaults instantiates a new CreateOrganizationInventoryDevicesSwapsBulk207Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInventoryDevicesSwapsBulk207ResponseWithDefaults() *CreateOrganizationInventoryDevicesSwapsBulk207Response {
	this := CreateOrganizationInventoryDevicesSwapsBulk207Response{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryDevicesSwapsBulk207Response) GetJobId() string {
	if o == nil || IsNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryDevicesSwapsBulk207Response) GetJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryDevicesSwapsBulk207Response) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *CreateOrganizationInventoryDevicesSwapsBulk207Response) SetJobId(v string) {
	o.JobId = &v
}

// GetSwaps returns the Swaps field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryDevicesSwapsBulk207Response) GetSwaps() []CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner {
	if o == nil || IsNil(o.Swaps) {
		var ret []CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner
		return ret
	}
	return o.Swaps
}

// GetSwapsOk returns a tuple with the Swaps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryDevicesSwapsBulk207Response) GetSwapsOk() ([]CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner, bool) {
	if o == nil || IsNil(o.Swaps) {
		return nil, false
	}
	return o.Swaps, true
}

// HasSwaps returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryDevicesSwapsBulk207Response) HasSwaps() bool {
	if o != nil && !IsNil(o.Swaps) {
		return true
	}

	return false
}

// SetSwaps gets a reference to the given []CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner and assigns it to the Swaps field.
func (o *CreateOrganizationInventoryDevicesSwapsBulk207Response) SetSwaps(v []CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) {
	o.Swaps = v
}

func (o CreateOrganizationInventoryDevicesSwapsBulk207Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationInventoryDevicesSwapsBulk207Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JobId) {
		toSerialize["jobId"] = o.JobId
	}
	if !IsNil(o.Swaps) {
		toSerialize["swaps"] = o.Swaps
	}
	return toSerialize, nil
}

type NullableCreateOrganizationInventoryDevicesSwapsBulk207Response struct {
	value *CreateOrganizationInventoryDevicesSwapsBulk207Response
	isSet bool
}

func (v NullableCreateOrganizationInventoryDevicesSwapsBulk207Response) Get() *CreateOrganizationInventoryDevicesSwapsBulk207Response {
	return v.value
}

func (v *NullableCreateOrganizationInventoryDevicesSwapsBulk207Response) Set(val *CreateOrganizationInventoryDevicesSwapsBulk207Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInventoryDevicesSwapsBulk207Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInventoryDevicesSwapsBulk207Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInventoryDevicesSwapsBulk207Response(val *CreateOrganizationInventoryDevicesSwapsBulk207Response) *NullableCreateOrganizationInventoryDevicesSwapsBulk207Response {
	return &NullableCreateOrganizationInventoryDevicesSwapsBulk207Response{value: val, isSet: true}
}

func (v NullableCreateOrganizationInventoryDevicesSwapsBulk207Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInventoryDevicesSwapsBulk207Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


