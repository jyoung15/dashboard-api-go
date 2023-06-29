/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CreateOrganizationAdaptivePolicyAclRequestRulesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationAdaptivePolicyAclRequestRulesInner{}

// CreateOrganizationAdaptivePolicyAclRequestRulesInner struct for CreateOrganizationAdaptivePolicyAclRequestRulesInner
type CreateOrganizationAdaptivePolicyAclRequestRulesInner struct {
	// 'allow' or 'deny' traffic specified by this rule.
	Policy string `json:"policy"`
	// The type of protocol (must be 'tcp', 'udp', 'icmp' or 'any').
	Protocol string `json:"protocol"`
	// Source port. Must be in the format of single port: '1', port list: '1,2' or port range: '1-10', and in the range of 1-65535, or 'any'. Default is 'any'.
	SrcPort *string `json:"srcPort,omitempty"`
	// Destination port. Must be in the format of single port: '1', port list: '1,2' or port range: '1-10', and in the range of 1-65535, or 'any'. Default is 'any'.
	DstPort *string `json:"dstPort,omitempty"`
}

// NewCreateOrganizationAdaptivePolicyAclRequestRulesInner instantiates a new CreateOrganizationAdaptivePolicyAclRequestRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationAdaptivePolicyAclRequestRulesInner(policy string, protocol string) *CreateOrganizationAdaptivePolicyAclRequestRulesInner {
	this := CreateOrganizationAdaptivePolicyAclRequestRulesInner{}
	this.Policy = policy
	this.Protocol = protocol
	return &this
}

// NewCreateOrganizationAdaptivePolicyAclRequestRulesInnerWithDefaults instantiates a new CreateOrganizationAdaptivePolicyAclRequestRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationAdaptivePolicyAclRequestRulesInnerWithDefaults() *CreateOrganizationAdaptivePolicyAclRequestRulesInner {
	this := CreateOrganizationAdaptivePolicyAclRequestRulesInner{}
	return &this
}

// GetPolicy returns the Policy field value
func (o *CreateOrganizationAdaptivePolicyAclRequestRulesInner) GetPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdaptivePolicyAclRequestRulesInner) GetPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Policy, true
}

// SetPolicy sets field value
func (o *CreateOrganizationAdaptivePolicyAclRequestRulesInner) SetPolicy(v string) {
	o.Policy = v
}

// GetProtocol returns the Protocol field value
func (o *CreateOrganizationAdaptivePolicyAclRequestRulesInner) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdaptivePolicyAclRequestRulesInner) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *CreateOrganizationAdaptivePolicyAclRequestRulesInner) SetProtocol(v string) {
	o.Protocol = v
}

// GetSrcPort returns the SrcPort field value if set, zero value otherwise.
func (o *CreateOrganizationAdaptivePolicyAclRequestRulesInner) GetSrcPort() string {
	if o == nil || IsNil(o.SrcPort) {
		var ret string
		return ret
	}
	return *o.SrcPort
}

// GetSrcPortOk returns a tuple with the SrcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdaptivePolicyAclRequestRulesInner) GetSrcPortOk() (*string, bool) {
	if o == nil || IsNil(o.SrcPort) {
		return nil, false
	}
	return o.SrcPort, true
}

// HasSrcPort returns a boolean if a field has been set.
func (o *CreateOrganizationAdaptivePolicyAclRequestRulesInner) HasSrcPort() bool {
	if o != nil && !IsNil(o.SrcPort) {
		return true
	}

	return false
}

// SetSrcPort gets a reference to the given string and assigns it to the SrcPort field.
func (o *CreateOrganizationAdaptivePolicyAclRequestRulesInner) SetSrcPort(v string) {
	o.SrcPort = &v
}

// GetDstPort returns the DstPort field value if set, zero value otherwise.
func (o *CreateOrganizationAdaptivePolicyAclRequestRulesInner) GetDstPort() string {
	if o == nil || IsNil(o.DstPort) {
		var ret string
		return ret
	}
	return *o.DstPort
}

// GetDstPortOk returns a tuple with the DstPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdaptivePolicyAclRequestRulesInner) GetDstPortOk() (*string, bool) {
	if o == nil || IsNil(o.DstPort) {
		return nil, false
	}
	return o.DstPort, true
}

// HasDstPort returns a boolean if a field has been set.
func (o *CreateOrganizationAdaptivePolicyAclRequestRulesInner) HasDstPort() bool {
	if o != nil && !IsNil(o.DstPort) {
		return true
	}

	return false
}

// SetDstPort gets a reference to the given string and assigns it to the DstPort field.
func (o *CreateOrganizationAdaptivePolicyAclRequestRulesInner) SetDstPort(v string) {
	o.DstPort = &v
}

func (o CreateOrganizationAdaptivePolicyAclRequestRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationAdaptivePolicyAclRequestRulesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["policy"] = o.Policy
	toSerialize["protocol"] = o.Protocol
	if !IsNil(o.SrcPort) {
		toSerialize["srcPort"] = o.SrcPort
	}
	if !IsNil(o.DstPort) {
		toSerialize["dstPort"] = o.DstPort
	}
	return toSerialize, nil
}

type NullableCreateOrganizationAdaptivePolicyAclRequestRulesInner struct {
	value *CreateOrganizationAdaptivePolicyAclRequestRulesInner
	isSet bool
}

func (v NullableCreateOrganizationAdaptivePolicyAclRequestRulesInner) Get() *CreateOrganizationAdaptivePolicyAclRequestRulesInner {
	return v.value
}

func (v *NullableCreateOrganizationAdaptivePolicyAclRequestRulesInner) Set(val *CreateOrganizationAdaptivePolicyAclRequestRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationAdaptivePolicyAclRequestRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationAdaptivePolicyAclRequestRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationAdaptivePolicyAclRequestRulesInner(val *CreateOrganizationAdaptivePolicyAclRequestRulesInner) *NullableCreateOrganizationAdaptivePolicyAclRequestRulesInner {
	return &NullableCreateOrganizationAdaptivePolicyAclRequestRulesInner{value: val, isSet: true}
}

func (v NullableCreateOrganizationAdaptivePolicyAclRequestRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationAdaptivePolicyAclRequestRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


