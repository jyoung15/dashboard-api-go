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

// checks if the CreateOrganizationAdaptivePolicyAclRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationAdaptivePolicyAclRequest{}

// CreateOrganizationAdaptivePolicyAclRequest struct for CreateOrganizationAdaptivePolicyAclRequest
type CreateOrganizationAdaptivePolicyAclRequest struct {
	// Name of the adaptive policy ACL
	Name string `json:"name"`
	// Description of the adaptive policy ACL
	Description *string `json:"description,omitempty"`
	// An ordered array of the adaptive policy ACL rules.
	Rules []CreateOrganizationAdaptivePolicyAclRequestRulesInner `json:"rules"`
	// IP version of adpative policy ACL. One of: 'any', 'ipv4' or 'ipv6'
	IpVersion string `json:"ipVersion"`
}

type _CreateOrganizationAdaptivePolicyAclRequest CreateOrganizationAdaptivePolicyAclRequest

// NewCreateOrganizationAdaptivePolicyAclRequest instantiates a new CreateOrganizationAdaptivePolicyAclRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationAdaptivePolicyAclRequest(name string, rules []CreateOrganizationAdaptivePolicyAclRequestRulesInner, ipVersion string) *CreateOrganizationAdaptivePolicyAclRequest {
	this := CreateOrganizationAdaptivePolicyAclRequest{}
	this.Name = name
	var description string = ""
	this.Description = &description
	this.Rules = rules
	this.IpVersion = ipVersion
	return &this
}

// NewCreateOrganizationAdaptivePolicyAclRequestWithDefaults instantiates a new CreateOrganizationAdaptivePolicyAclRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationAdaptivePolicyAclRequestWithDefaults() *CreateOrganizationAdaptivePolicyAclRequest {
	this := CreateOrganizationAdaptivePolicyAclRequest{}
	var description string = ""
	this.Description = &description
	return &this
}

// GetName returns the Name field value
func (o *CreateOrganizationAdaptivePolicyAclRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdaptivePolicyAclRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrganizationAdaptivePolicyAclRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateOrganizationAdaptivePolicyAclRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdaptivePolicyAclRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOrganizationAdaptivePolicyAclRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateOrganizationAdaptivePolicyAclRequest) SetDescription(v string) {
	o.Description = &v
}

// GetRules returns the Rules field value
func (o *CreateOrganizationAdaptivePolicyAclRequest) GetRules() []CreateOrganizationAdaptivePolicyAclRequestRulesInner {
	if o == nil {
		var ret []CreateOrganizationAdaptivePolicyAclRequestRulesInner
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdaptivePolicyAclRequest) GetRulesOk() ([]CreateOrganizationAdaptivePolicyAclRequestRulesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *CreateOrganizationAdaptivePolicyAclRequest) SetRules(v []CreateOrganizationAdaptivePolicyAclRequestRulesInner) {
	o.Rules = v
}

// GetIpVersion returns the IpVersion field value
func (o *CreateOrganizationAdaptivePolicyAclRequest) GetIpVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdaptivePolicyAclRequest) GetIpVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpVersion, true
}

// SetIpVersion sets field value
func (o *CreateOrganizationAdaptivePolicyAclRequest) SetIpVersion(v string) {
	o.IpVersion = v
}

func (o CreateOrganizationAdaptivePolicyAclRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationAdaptivePolicyAclRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["rules"] = o.Rules
	toSerialize["ipVersion"] = o.IpVersion
	return toSerialize, nil
}

func (o *CreateOrganizationAdaptivePolicyAclRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"rules",
		"ipVersion",
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

	varCreateOrganizationAdaptivePolicyAclRequest := _CreateOrganizationAdaptivePolicyAclRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrganizationAdaptivePolicyAclRequest)

	if err != nil {
		return err
	}

	*o = CreateOrganizationAdaptivePolicyAclRequest(varCreateOrganizationAdaptivePolicyAclRequest)

	return err
}

type NullableCreateOrganizationAdaptivePolicyAclRequest struct {
	value *CreateOrganizationAdaptivePolicyAclRequest
	isSet bool
}

func (v NullableCreateOrganizationAdaptivePolicyAclRequest) Get() *CreateOrganizationAdaptivePolicyAclRequest {
	return v.value
}

func (v *NullableCreateOrganizationAdaptivePolicyAclRequest) Set(val *CreateOrganizationAdaptivePolicyAclRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationAdaptivePolicyAclRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationAdaptivePolicyAclRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationAdaptivePolicyAclRequest(val *CreateOrganizationAdaptivePolicyAclRequest) *NullableCreateOrganizationAdaptivePolicyAclRequest {
	return &NullableCreateOrganizationAdaptivePolicyAclRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationAdaptivePolicyAclRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationAdaptivePolicyAclRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


