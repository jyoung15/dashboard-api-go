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

// checks if the GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner{}

// GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner struct for GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner
type GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner struct {
	// The ID of the policy object
	Id *string `json:"id,omitempty"`
	// The name of the policy object
	Name *string `json:"name,omitempty"`
}

// NewGetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner instantiates a new GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner() *GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner {
	this := GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner{}
	return &this
}

// NewGetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInnerWithDefaults instantiates a new GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInnerWithDefaults() *GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner {
	this := GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner) SetName(v string) {
	o.Name = &v
}

func (o GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner struct {
	value *GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner
	isSet bool
}

func (v NullableGetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner) Get() *GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner {
	return v.value
}

func (v *NullableGetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner) Set(val *GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner(val *GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner) *NullableGetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner {
	return &NullableGetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner{value: val, isSet: true}
}

func (v NullableGetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


