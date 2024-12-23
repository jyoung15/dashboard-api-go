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

// checks if the GetOrganizationAdaptivePolicyPolicies200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationAdaptivePolicyPolicies200ResponseInner{}

// GetOrganizationAdaptivePolicyPolicies200ResponseInner struct for GetOrganizationAdaptivePolicyPolicies200ResponseInner
type GetOrganizationAdaptivePolicyPolicies200ResponseInner struct {
	// The ID for the adaptive policy
	AdaptivePolicyId *string `json:"adaptivePolicyId,omitempty"`
	SourceGroup *GetOrganizationAdaptivePolicyPolicies200ResponseInnerSourceGroup `json:"sourceGroup,omitempty"`
	DestinationGroup *GetOrganizationAdaptivePolicyPolicies200ResponseInnerDestinationGroup `json:"destinationGroup,omitempty"`
	// The access control lists for the adaptive policy
	Acls []GetOrganizationAdaptivePolicyPolicies200ResponseInnerAclsInner `json:"acls,omitempty"`
	// The rule to apply if there is no matching ACL
	LastEntryRule *string `json:"lastEntryRule,omitempty"`
	// The created at timestamp for the adaptive policy
	CreatedAt *string `json:"createdAt,omitempty"`
	// The updated at timestamp for the adaptive policy
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewGetOrganizationAdaptivePolicyPolicies200ResponseInner instantiates a new GetOrganizationAdaptivePolicyPolicies200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationAdaptivePolicyPolicies200ResponseInner() *GetOrganizationAdaptivePolicyPolicies200ResponseInner {
	this := GetOrganizationAdaptivePolicyPolicies200ResponseInner{}
	return &this
}

// NewGetOrganizationAdaptivePolicyPolicies200ResponseInnerWithDefaults instantiates a new GetOrganizationAdaptivePolicyPolicies200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationAdaptivePolicyPolicies200ResponseInnerWithDefaults() *GetOrganizationAdaptivePolicyPolicies200ResponseInner {
	this := GetOrganizationAdaptivePolicyPolicies200ResponseInner{}
	return &this
}

// GetAdaptivePolicyId returns the AdaptivePolicyId field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetAdaptivePolicyId() string {
	if o == nil || IsNil(o.AdaptivePolicyId) {
		var ret string
		return ret
	}
	return *o.AdaptivePolicyId
}

// GetAdaptivePolicyIdOk returns a tuple with the AdaptivePolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetAdaptivePolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdaptivePolicyId) {
		return nil, false
	}
	return o.AdaptivePolicyId, true
}

// HasAdaptivePolicyId returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) HasAdaptivePolicyId() bool {
	if o != nil && !IsNil(o.AdaptivePolicyId) {
		return true
	}

	return false
}

// SetAdaptivePolicyId gets a reference to the given string and assigns it to the AdaptivePolicyId field.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) SetAdaptivePolicyId(v string) {
	o.AdaptivePolicyId = &v
}

// GetSourceGroup returns the SourceGroup field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetSourceGroup() GetOrganizationAdaptivePolicyPolicies200ResponseInnerSourceGroup {
	if o == nil || IsNil(o.SourceGroup) {
		var ret GetOrganizationAdaptivePolicyPolicies200ResponseInnerSourceGroup
		return ret
	}
	return *o.SourceGroup
}

// GetSourceGroupOk returns a tuple with the SourceGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetSourceGroupOk() (*GetOrganizationAdaptivePolicyPolicies200ResponseInnerSourceGroup, bool) {
	if o == nil || IsNil(o.SourceGroup) {
		return nil, false
	}
	return o.SourceGroup, true
}

// HasSourceGroup returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) HasSourceGroup() bool {
	if o != nil && !IsNil(o.SourceGroup) {
		return true
	}

	return false
}

// SetSourceGroup gets a reference to the given GetOrganizationAdaptivePolicyPolicies200ResponseInnerSourceGroup and assigns it to the SourceGroup field.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) SetSourceGroup(v GetOrganizationAdaptivePolicyPolicies200ResponseInnerSourceGroup) {
	o.SourceGroup = &v
}

// GetDestinationGroup returns the DestinationGroup field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetDestinationGroup() GetOrganizationAdaptivePolicyPolicies200ResponseInnerDestinationGroup {
	if o == nil || IsNil(o.DestinationGroup) {
		var ret GetOrganizationAdaptivePolicyPolicies200ResponseInnerDestinationGroup
		return ret
	}
	return *o.DestinationGroup
}

// GetDestinationGroupOk returns a tuple with the DestinationGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetDestinationGroupOk() (*GetOrganizationAdaptivePolicyPolicies200ResponseInnerDestinationGroup, bool) {
	if o == nil || IsNil(o.DestinationGroup) {
		return nil, false
	}
	return o.DestinationGroup, true
}

// HasDestinationGroup returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) HasDestinationGroup() bool {
	if o != nil && !IsNil(o.DestinationGroup) {
		return true
	}

	return false
}

// SetDestinationGroup gets a reference to the given GetOrganizationAdaptivePolicyPolicies200ResponseInnerDestinationGroup and assigns it to the DestinationGroup field.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) SetDestinationGroup(v GetOrganizationAdaptivePolicyPolicies200ResponseInnerDestinationGroup) {
	o.DestinationGroup = &v
}

// GetAcls returns the Acls field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetAcls() []GetOrganizationAdaptivePolicyPolicies200ResponseInnerAclsInner {
	if o == nil || IsNil(o.Acls) {
		var ret []GetOrganizationAdaptivePolicyPolicies200ResponseInnerAclsInner
		return ret
	}
	return o.Acls
}

// GetAclsOk returns a tuple with the Acls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetAclsOk() ([]GetOrganizationAdaptivePolicyPolicies200ResponseInnerAclsInner, bool) {
	if o == nil || IsNil(o.Acls) {
		return nil, false
	}
	return o.Acls, true
}

// HasAcls returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) HasAcls() bool {
	if o != nil && !IsNil(o.Acls) {
		return true
	}

	return false
}

// SetAcls gets a reference to the given []GetOrganizationAdaptivePolicyPolicies200ResponseInnerAclsInner and assigns it to the Acls field.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) SetAcls(v []GetOrganizationAdaptivePolicyPolicies200ResponseInnerAclsInner) {
	o.Acls = v
}

// GetLastEntryRule returns the LastEntryRule field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetLastEntryRule() string {
	if o == nil || IsNil(o.LastEntryRule) {
		var ret string
		return ret
	}
	return *o.LastEntryRule
}

// GetLastEntryRuleOk returns a tuple with the LastEntryRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetLastEntryRuleOk() (*string, bool) {
	if o == nil || IsNil(o.LastEntryRule) {
		return nil, false
	}
	return o.LastEntryRule, true
}

// HasLastEntryRule returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) HasLastEntryRule() bool {
	if o != nil && !IsNil(o.LastEntryRule) {
		return true
	}

	return false
}

// SetLastEntryRule gets a reference to the given string and assigns it to the LastEntryRule field.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) SetLastEntryRule(v string) {
	o.LastEntryRule = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o GetOrganizationAdaptivePolicyPolicies200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationAdaptivePolicyPolicies200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdaptivePolicyId) {
		toSerialize["adaptivePolicyId"] = o.AdaptivePolicyId
	}
	if !IsNil(o.SourceGroup) {
		toSerialize["sourceGroup"] = o.SourceGroup
	}
	if !IsNil(o.DestinationGroup) {
		toSerialize["destinationGroup"] = o.DestinationGroup
	}
	if !IsNil(o.Acls) {
		toSerialize["acls"] = o.Acls
	}
	if !IsNil(o.LastEntryRule) {
		toSerialize["lastEntryRule"] = o.LastEntryRule
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableGetOrganizationAdaptivePolicyPolicies200ResponseInner struct {
	value *GetOrganizationAdaptivePolicyPolicies200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationAdaptivePolicyPolicies200ResponseInner) Get() *GetOrganizationAdaptivePolicyPolicies200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationAdaptivePolicyPolicies200ResponseInner) Set(val *GetOrganizationAdaptivePolicyPolicies200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationAdaptivePolicyPolicies200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationAdaptivePolicyPolicies200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationAdaptivePolicyPolicies200ResponseInner(val *GetOrganizationAdaptivePolicyPolicies200ResponseInner) *NullableGetOrganizationAdaptivePolicyPolicies200ResponseInner {
	return &NullableGetOrganizationAdaptivePolicyPolicies200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationAdaptivePolicyPolicies200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationAdaptivePolicyPolicies200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


