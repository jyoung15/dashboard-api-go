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

// checks if the GetOrganizationAdaptivePolicyGroups200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationAdaptivePolicyGroups200ResponseInner{}

// GetOrganizationAdaptivePolicyGroups200ResponseInner struct for GetOrganizationAdaptivePolicyGroups200ResponseInner
type GetOrganizationAdaptivePolicyGroups200ResponseInner struct {
	// The ID of the adaptive policy group
	GroupId *string `json:"groupId,omitempty"`
	// The name of the adaptive policy group
	Name *string `json:"name,omitempty"`
	// The security group tag for the adaptive policy group
	Sgt *int32 `json:"sgt,omitempty"`
	// The description for the adaptive policy group
	Description *string `json:"description,omitempty"`
	// The policy objects for the adaptive policy group
	PolicyObjects []GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner `json:"policyObjects,omitempty"`
	// Whether the adaptive policy group is the default group
	IsDefaultGroup *bool `json:"isDefaultGroup,omitempty"`
	// List of required IP mappings for the adaptive policy group
	RequiredIpMappings []string `json:"requiredIpMappings,omitempty"`
	// Created at timestamp for the adaptive policy group
	CreatedAt *string `json:"createdAt,omitempty"`
	// Updated at timestamp for the adaptive policy group
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewGetOrganizationAdaptivePolicyGroups200ResponseInner instantiates a new GetOrganizationAdaptivePolicyGroups200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationAdaptivePolicyGroups200ResponseInner() *GetOrganizationAdaptivePolicyGroups200ResponseInner {
	this := GetOrganizationAdaptivePolicyGroups200ResponseInner{}
	return &this
}

// NewGetOrganizationAdaptivePolicyGroups200ResponseInnerWithDefaults instantiates a new GetOrganizationAdaptivePolicyGroups200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationAdaptivePolicyGroups200ResponseInnerWithDefaults() *GetOrganizationAdaptivePolicyGroups200ResponseInner {
	this := GetOrganizationAdaptivePolicyGroups200ResponseInner{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) SetGroupId(v string) {
	o.GroupId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetSgt returns the Sgt field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetSgt() int32 {
	if o == nil || IsNil(o.Sgt) {
		var ret int32
		return ret
	}
	return *o.Sgt
}

// GetSgtOk returns a tuple with the Sgt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetSgtOk() (*int32, bool) {
	if o == nil || IsNil(o.Sgt) {
		return nil, false
	}
	return o.Sgt, true
}

// HasSgt returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) HasSgt() bool {
	if o != nil && !IsNil(o.Sgt) {
		return true
	}

	return false
}

// SetSgt gets a reference to the given int32 and assigns it to the Sgt field.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) SetSgt(v int32) {
	o.Sgt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) SetDescription(v string) {
	o.Description = &v
}

// GetPolicyObjects returns the PolicyObjects field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetPolicyObjects() []GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner {
	if o == nil || IsNil(o.PolicyObjects) {
		var ret []GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner
		return ret
	}
	return o.PolicyObjects
}

// GetPolicyObjectsOk returns a tuple with the PolicyObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetPolicyObjectsOk() ([]GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner, bool) {
	if o == nil || IsNil(o.PolicyObjects) {
		return nil, false
	}
	return o.PolicyObjects, true
}

// HasPolicyObjects returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) HasPolicyObjects() bool {
	if o != nil && !IsNil(o.PolicyObjects) {
		return true
	}

	return false
}

// SetPolicyObjects gets a reference to the given []GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner and assigns it to the PolicyObjects field.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) SetPolicyObjects(v []GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner) {
	o.PolicyObjects = v
}

// GetIsDefaultGroup returns the IsDefaultGroup field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetIsDefaultGroup() bool {
	if o == nil || IsNil(o.IsDefaultGroup) {
		var ret bool
		return ret
	}
	return *o.IsDefaultGroup
}

// GetIsDefaultGroupOk returns a tuple with the IsDefaultGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetIsDefaultGroupOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefaultGroup) {
		return nil, false
	}
	return o.IsDefaultGroup, true
}

// HasIsDefaultGroup returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) HasIsDefaultGroup() bool {
	if o != nil && !IsNil(o.IsDefaultGroup) {
		return true
	}

	return false
}

// SetIsDefaultGroup gets a reference to the given bool and assigns it to the IsDefaultGroup field.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) SetIsDefaultGroup(v bool) {
	o.IsDefaultGroup = &v
}

// GetRequiredIpMappings returns the RequiredIpMappings field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetRequiredIpMappings() []string {
	if o == nil || IsNil(o.RequiredIpMappings) {
		var ret []string
		return ret
	}
	return o.RequiredIpMappings
}

// GetRequiredIpMappingsOk returns a tuple with the RequiredIpMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetRequiredIpMappingsOk() ([]string, bool) {
	if o == nil || IsNil(o.RequiredIpMappings) {
		return nil, false
	}
	return o.RequiredIpMappings, true
}

// HasRequiredIpMappings returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) HasRequiredIpMappings() bool {
	if o != nil && !IsNil(o.RequiredIpMappings) {
		return true
	}

	return false
}

// SetRequiredIpMappings gets a reference to the given []string and assigns it to the RequiredIpMappings field.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) SetRequiredIpMappings(v []string) {
	o.RequiredIpMappings = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o GetOrganizationAdaptivePolicyGroups200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationAdaptivePolicyGroups200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Sgt) {
		toSerialize["sgt"] = o.Sgt
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.PolicyObjects) {
		toSerialize["policyObjects"] = o.PolicyObjects
	}
	if !IsNil(o.IsDefaultGroup) {
		toSerialize["isDefaultGroup"] = o.IsDefaultGroup
	}
	if !IsNil(o.RequiredIpMappings) {
		toSerialize["requiredIpMappings"] = o.RequiredIpMappings
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableGetOrganizationAdaptivePolicyGroups200ResponseInner struct {
	value *GetOrganizationAdaptivePolicyGroups200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationAdaptivePolicyGroups200ResponseInner) Get() *GetOrganizationAdaptivePolicyGroups200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationAdaptivePolicyGroups200ResponseInner) Set(val *GetOrganizationAdaptivePolicyGroups200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationAdaptivePolicyGroups200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationAdaptivePolicyGroups200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationAdaptivePolicyGroups200ResponseInner(val *GetOrganizationAdaptivePolicyGroups200ResponseInner) *NullableGetOrganizationAdaptivePolicyGroups200ResponseInner {
	return &NullableGetOrganizationAdaptivePolicyGroups200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationAdaptivePolicyGroups200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationAdaptivePolicyGroups200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


