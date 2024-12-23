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

// checks if the GetOrganizationAdaptivePolicyOverview200ResponseCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationAdaptivePolicyOverview200ResponseCounts{}

// GetOrganizationAdaptivePolicyOverview200ResponseCounts The current amount of various adaptive policy objects.
type GetOrganizationAdaptivePolicyOverview200ResponseCounts struct {
	// Number of adaptive policy groups currently in the organization.
	Groups *int32 `json:"groups,omitempty"`
	// Number of user-created adaptive policy groups currently in the organization.
	CustomGroups *int32 `json:"customGroups,omitempty"`
	// Number of user-created adaptive policy ACLs currently in the organization.
	CustomAcls *int32 `json:"customAcls,omitempty"`
	// Number of adaptive policies currently in the organization.
	Policies *int32 `json:"policies,omitempty"`
	// Number of adaptive policies currently in the organization that deny all traffic.
	DenyPolicies *int32 `json:"denyPolicies,omitempty"`
	// Number of adaptive policies currently in the organization that allow all traffic.
	AllowPolicies *int32 `json:"allowPolicies,omitempty"`
	// Number of policy objects (with the adaptive policy type) currently in the organization.
	PolicyObjects *int32 `json:"policyObjects,omitempty"`
}

// NewGetOrganizationAdaptivePolicyOverview200ResponseCounts instantiates a new GetOrganizationAdaptivePolicyOverview200ResponseCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationAdaptivePolicyOverview200ResponseCounts() *GetOrganizationAdaptivePolicyOverview200ResponseCounts {
	this := GetOrganizationAdaptivePolicyOverview200ResponseCounts{}
	return &this
}

// NewGetOrganizationAdaptivePolicyOverview200ResponseCountsWithDefaults instantiates a new GetOrganizationAdaptivePolicyOverview200ResponseCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationAdaptivePolicyOverview200ResponseCountsWithDefaults() *GetOrganizationAdaptivePolicyOverview200ResponseCounts {
	this := GetOrganizationAdaptivePolicyOverview200ResponseCounts{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetGroups() int32 {
	if o == nil || IsNil(o.Groups) {
		var ret int32
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetGroupsOk() (*int32, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given int32 and assigns it to the Groups field.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) SetGroups(v int32) {
	o.Groups = &v
}

// GetCustomGroups returns the CustomGroups field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetCustomGroups() int32 {
	if o == nil || IsNil(o.CustomGroups) {
		var ret int32
		return ret
	}
	return *o.CustomGroups
}

// GetCustomGroupsOk returns a tuple with the CustomGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetCustomGroupsOk() (*int32, bool) {
	if o == nil || IsNil(o.CustomGroups) {
		return nil, false
	}
	return o.CustomGroups, true
}

// HasCustomGroups returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) HasCustomGroups() bool {
	if o != nil && !IsNil(o.CustomGroups) {
		return true
	}

	return false
}

// SetCustomGroups gets a reference to the given int32 and assigns it to the CustomGroups field.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) SetCustomGroups(v int32) {
	o.CustomGroups = &v
}

// GetCustomAcls returns the CustomAcls field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetCustomAcls() int32 {
	if o == nil || IsNil(o.CustomAcls) {
		var ret int32
		return ret
	}
	return *o.CustomAcls
}

// GetCustomAclsOk returns a tuple with the CustomAcls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetCustomAclsOk() (*int32, bool) {
	if o == nil || IsNil(o.CustomAcls) {
		return nil, false
	}
	return o.CustomAcls, true
}

// HasCustomAcls returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) HasCustomAcls() bool {
	if o != nil && !IsNil(o.CustomAcls) {
		return true
	}

	return false
}

// SetCustomAcls gets a reference to the given int32 and assigns it to the CustomAcls field.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) SetCustomAcls(v int32) {
	o.CustomAcls = &v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetPolicies() int32 {
	if o == nil || IsNil(o.Policies) {
		var ret int32
		return ret
	}
	return *o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetPoliciesOk() (*int32, bool) {
	if o == nil || IsNil(o.Policies) {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) HasPolicies() bool {
	if o != nil && !IsNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given int32 and assigns it to the Policies field.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) SetPolicies(v int32) {
	o.Policies = &v
}

// GetDenyPolicies returns the DenyPolicies field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetDenyPolicies() int32 {
	if o == nil || IsNil(o.DenyPolicies) {
		var ret int32
		return ret
	}
	return *o.DenyPolicies
}

// GetDenyPoliciesOk returns a tuple with the DenyPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetDenyPoliciesOk() (*int32, bool) {
	if o == nil || IsNil(o.DenyPolicies) {
		return nil, false
	}
	return o.DenyPolicies, true
}

// HasDenyPolicies returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) HasDenyPolicies() bool {
	if o != nil && !IsNil(o.DenyPolicies) {
		return true
	}

	return false
}

// SetDenyPolicies gets a reference to the given int32 and assigns it to the DenyPolicies field.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) SetDenyPolicies(v int32) {
	o.DenyPolicies = &v
}

// GetAllowPolicies returns the AllowPolicies field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetAllowPolicies() int32 {
	if o == nil || IsNil(o.AllowPolicies) {
		var ret int32
		return ret
	}
	return *o.AllowPolicies
}

// GetAllowPoliciesOk returns a tuple with the AllowPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetAllowPoliciesOk() (*int32, bool) {
	if o == nil || IsNil(o.AllowPolicies) {
		return nil, false
	}
	return o.AllowPolicies, true
}

// HasAllowPolicies returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) HasAllowPolicies() bool {
	if o != nil && !IsNil(o.AllowPolicies) {
		return true
	}

	return false
}

// SetAllowPolicies gets a reference to the given int32 and assigns it to the AllowPolicies field.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) SetAllowPolicies(v int32) {
	o.AllowPolicies = &v
}

// GetPolicyObjects returns the PolicyObjects field value if set, zero value otherwise.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetPolicyObjects() int32 {
	if o == nil || IsNil(o.PolicyObjects) {
		var ret int32
		return ret
	}
	return *o.PolicyObjects
}

// GetPolicyObjectsOk returns a tuple with the PolicyObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) GetPolicyObjectsOk() (*int32, bool) {
	if o == nil || IsNil(o.PolicyObjects) {
		return nil, false
	}
	return o.PolicyObjects, true
}

// HasPolicyObjects returns a boolean if a field has been set.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) HasPolicyObjects() bool {
	if o != nil && !IsNil(o.PolicyObjects) {
		return true
	}

	return false
}

// SetPolicyObjects gets a reference to the given int32 and assigns it to the PolicyObjects field.
func (o *GetOrganizationAdaptivePolicyOverview200ResponseCounts) SetPolicyObjects(v int32) {
	o.PolicyObjects = &v
}

func (o GetOrganizationAdaptivePolicyOverview200ResponseCounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationAdaptivePolicyOverview200ResponseCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.CustomGroups) {
		toSerialize["customGroups"] = o.CustomGroups
	}
	if !IsNil(o.CustomAcls) {
		toSerialize["customAcls"] = o.CustomAcls
	}
	if !IsNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}
	if !IsNil(o.DenyPolicies) {
		toSerialize["denyPolicies"] = o.DenyPolicies
	}
	if !IsNil(o.AllowPolicies) {
		toSerialize["allowPolicies"] = o.AllowPolicies
	}
	if !IsNil(o.PolicyObjects) {
		toSerialize["policyObjects"] = o.PolicyObjects
	}
	return toSerialize, nil
}

type NullableGetOrganizationAdaptivePolicyOverview200ResponseCounts struct {
	value *GetOrganizationAdaptivePolicyOverview200ResponseCounts
	isSet bool
}

func (v NullableGetOrganizationAdaptivePolicyOverview200ResponseCounts) Get() *GetOrganizationAdaptivePolicyOverview200ResponseCounts {
	return v.value
}

func (v *NullableGetOrganizationAdaptivePolicyOverview200ResponseCounts) Set(val *GetOrganizationAdaptivePolicyOverview200ResponseCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationAdaptivePolicyOverview200ResponseCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationAdaptivePolicyOverview200ResponseCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationAdaptivePolicyOverview200ResponseCounts(val *GetOrganizationAdaptivePolicyOverview200ResponseCounts) *NullableGetOrganizationAdaptivePolicyOverview200ResponseCounts {
	return &NullableGetOrganizationAdaptivePolicyOverview200ResponseCounts{value: val, isSet: true}
}

func (v NullableGetOrganizationAdaptivePolicyOverview200ResponseCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationAdaptivePolicyOverview200ResponseCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


