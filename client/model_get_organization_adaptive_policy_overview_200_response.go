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

// checks if the GetOrganizationAdaptivePolicyOverview200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationAdaptivePolicyOverview200Response{}

// GetOrganizationAdaptivePolicyOverview200Response struct for GetOrganizationAdaptivePolicyOverview200Response
type GetOrganizationAdaptivePolicyOverview200Response struct {
	Counts *GetOrganizationAdaptivePolicyOverview200ResponseCounts `json:"counts,omitempty"`
	Limits *GetOrganizationAdaptivePolicyOverview200ResponseLimits `json:"limits,omitempty"`
}

// NewGetOrganizationAdaptivePolicyOverview200Response instantiates a new GetOrganizationAdaptivePolicyOverview200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationAdaptivePolicyOverview200Response() *GetOrganizationAdaptivePolicyOverview200Response {
	this := GetOrganizationAdaptivePolicyOverview200Response{}
	return &this
}

// NewGetOrganizationAdaptivePolicyOverview200ResponseWithDefaults instantiates a new GetOrganizationAdaptivePolicyOverview200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationAdaptivePolicyOverview200ResponseWithDefaults() *GetOrganizationAdaptivePolicyOverview200Response {
	this := GetOrganizationAdaptivePolicyOverview200Response{}
	return &this
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyOverview200Response) GetCounts() GetOrganizationAdaptivePolicyOverview200ResponseCounts {
	if o == nil || IsNil(o.Counts) {
		var ret GetOrganizationAdaptivePolicyOverview200ResponseCounts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyOverview200Response) GetCountsOk() (*GetOrganizationAdaptivePolicyOverview200ResponseCounts, bool) {
	if o == nil || IsNil(o.Counts) {
		return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyOverview200Response) HasCounts() bool {
	if o != nil && !IsNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given GetOrganizationAdaptivePolicyOverview200ResponseCounts and assigns it to the Counts field.
func (o *GetOrganizationAdaptivePolicyOverview200Response) SetCounts(v GetOrganizationAdaptivePolicyOverview200ResponseCounts) {
	o.Counts = &v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyOverview200Response) GetLimits() GetOrganizationAdaptivePolicyOverview200ResponseLimits {
	if o == nil || IsNil(o.Limits) {
		var ret GetOrganizationAdaptivePolicyOverview200ResponseLimits
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyOverview200Response) GetLimitsOk() (*GetOrganizationAdaptivePolicyOverview200ResponseLimits, bool) {
	if o == nil || IsNil(o.Limits) {
		return nil, false
	}
	return o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyOverview200Response) HasLimits() bool {
	if o != nil && !IsNil(o.Limits) {
		return true
	}

	return false
}

// SetLimits gets a reference to the given GetOrganizationAdaptivePolicyOverview200ResponseLimits and assigns it to the Limits field.
func (o *GetOrganizationAdaptivePolicyOverview200Response) SetLimits(v GetOrganizationAdaptivePolicyOverview200ResponseLimits) {
	o.Limits = &v
}

func (o GetOrganizationAdaptivePolicyOverview200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationAdaptivePolicyOverview200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	if !IsNil(o.Limits) {
		toSerialize["limits"] = o.Limits
	}
	return toSerialize, nil
}

type NullableGetOrganizationAdaptivePolicyOverview200Response struct {
	value *GetOrganizationAdaptivePolicyOverview200Response
	isSet bool
}

func (v NullableGetOrganizationAdaptivePolicyOverview200Response) Get() *GetOrganizationAdaptivePolicyOverview200Response {
	return v.value
}

func (v *NullableGetOrganizationAdaptivePolicyOverview200Response) Set(val *GetOrganizationAdaptivePolicyOverview200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationAdaptivePolicyOverview200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationAdaptivePolicyOverview200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationAdaptivePolicyOverview200Response(val *GetOrganizationAdaptivePolicyOverview200Response) *NullableGetOrganizationAdaptivePolicyOverview200Response {
	return &NullableGetOrganizationAdaptivePolicyOverview200Response{value: val, isSet: true}
}

func (v NullableGetOrganizationAdaptivePolicyOverview200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationAdaptivePolicyOverview200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


