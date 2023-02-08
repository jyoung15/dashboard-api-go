/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateOrganizationPolicyObjectRequest struct for UpdateOrganizationPolicyObjectRequest
type UpdateOrganizationPolicyObjectRequest struct {
	// Name of a policy object, unique within the organization (alphanumeric, space, dash, or underscore characters only)
	Name *string `json:"name,omitempty"`
	// CIDR Value of a policy object (e.g. 10.11.12.1/24\")
	Cidr *string `json:"cidr,omitempty"`
	// Fully qualified domain name of policy object (e.g. \"example.com\")
	Fqdn *string `json:"fqdn,omitempty"`
	// Mask of a policy object (e.g. \"255.255.0.0\")
	Mask *string `json:"mask,omitempty"`
	// IP Address of a policy object (e.g. \"1.2.3.4\")
	Ip *string `json:"ip,omitempty"`
	// The IDs of policy object groups the policy object belongs to
	GroupIds []int32 `json:"groupIds,omitempty"`
}

// NewUpdateOrganizationPolicyObjectRequest instantiates a new UpdateOrganizationPolicyObjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationPolicyObjectRequest() *UpdateOrganizationPolicyObjectRequest {
	this := UpdateOrganizationPolicyObjectRequest{}
	return &this
}

// NewUpdateOrganizationPolicyObjectRequestWithDefaults instantiates a new UpdateOrganizationPolicyObjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationPolicyObjectRequestWithDefaults() *UpdateOrganizationPolicyObjectRequest {
	this := UpdateOrganizationPolicyObjectRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateOrganizationPolicyObjectRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationPolicyObjectRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateOrganizationPolicyObjectRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateOrganizationPolicyObjectRequest) SetName(v string) {
	o.Name = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *UpdateOrganizationPolicyObjectRequest) GetCidr() string {
	if o == nil || isNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationPolicyObjectRequest) GetCidrOk() (*string, bool) {
	if o == nil || isNil(o.Cidr) {
    return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *UpdateOrganizationPolicyObjectRequest) HasCidr() bool {
	if o != nil && !isNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *UpdateOrganizationPolicyObjectRequest) SetCidr(v string) {
	o.Cidr = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *UpdateOrganizationPolicyObjectRequest) GetFqdn() string {
	if o == nil || isNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationPolicyObjectRequest) GetFqdnOk() (*string, bool) {
	if o == nil || isNil(o.Fqdn) {
    return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *UpdateOrganizationPolicyObjectRequest) HasFqdn() bool {
	if o != nil && !isNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *UpdateOrganizationPolicyObjectRequest) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *UpdateOrganizationPolicyObjectRequest) GetMask() string {
	if o == nil || isNil(o.Mask) {
		var ret string
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationPolicyObjectRequest) GetMaskOk() (*string, bool) {
	if o == nil || isNil(o.Mask) {
    return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *UpdateOrganizationPolicyObjectRequest) HasMask() bool {
	if o != nil && !isNil(o.Mask) {
		return true
	}

	return false
}

// SetMask gets a reference to the given string and assigns it to the Mask field.
func (o *UpdateOrganizationPolicyObjectRequest) SetMask(v string) {
	o.Mask = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *UpdateOrganizationPolicyObjectRequest) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationPolicyObjectRequest) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *UpdateOrganizationPolicyObjectRequest) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *UpdateOrganizationPolicyObjectRequest) SetIp(v string) {
	o.Ip = &v
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise.
func (o *UpdateOrganizationPolicyObjectRequest) GetGroupIds() []int32 {
	if o == nil || isNil(o.GroupIds) {
		var ret []int32
		return ret
	}
	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationPolicyObjectRequest) GetGroupIdsOk() ([]int32, bool) {
	if o == nil || isNil(o.GroupIds) {
    return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *UpdateOrganizationPolicyObjectRequest) HasGroupIds() bool {
	if o != nil && !isNil(o.GroupIds) {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []int32 and assigns it to the GroupIds field.
func (o *UpdateOrganizationPolicyObjectRequest) SetGroupIds(v []int32) {
	o.GroupIds = v
}

func (o UpdateOrganizationPolicyObjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !isNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !isNil(o.Mask) {
		toSerialize["mask"] = o.Mask
	}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.GroupIds) {
		toSerialize["groupIds"] = o.GroupIds
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateOrganizationPolicyObjectRequest struct {
	value *UpdateOrganizationPolicyObjectRequest
	isSet bool
}

func (v NullableUpdateOrganizationPolicyObjectRequest) Get() *UpdateOrganizationPolicyObjectRequest {
	return v.value
}

func (v *NullableUpdateOrganizationPolicyObjectRequest) Set(val *UpdateOrganizationPolicyObjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationPolicyObjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationPolicyObjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationPolicyObjectRequest(val *UpdateOrganizationPolicyObjectRequest) *NullableUpdateOrganizationPolicyObjectRequest {
	return &NullableUpdateOrganizationPolicyObjectRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationPolicyObjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationPolicyObjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


