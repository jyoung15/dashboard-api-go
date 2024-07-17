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

// checks if the UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner{}

// UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner struct for UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner
type UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner struct {
	// The Id of the Network
	NetworkId string `json:"networkId"`
	// Array of Sentry Group Policies for the Network
	Policies []UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInnerPoliciesInner `json:"policies,omitempty"`
}

type _UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner

// NewUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner instantiates a new UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner(networkId string) *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner {
	this := UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner{}
	this.NetworkId = networkId
	return &this
}

// NewUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInnerWithDefaults instantiates a new UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInnerWithDefaults() *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner {
	this := UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner{}
	return &this
}

// GetNetworkId returns the NetworkId field value
func (o *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) GetNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) GetPolicies() []UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInnerPoliciesInner {
	if o == nil || IsNil(o.Policies) {
		var ret []UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInnerPoliciesInner
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) GetPoliciesOk() ([]UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInnerPoliciesInner, bool) {
	if o == nil || IsNil(o.Policies) {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) HasPolicies() bool {
	if o != nil && !IsNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInnerPoliciesInner and assigns it to the Policies field.
func (o *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) SetPolicies(v []UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInnerPoliciesInner) {
	o.Policies = v
}

func (o UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["networkId"] = o.NetworkId
	if !IsNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}
	return toSerialize, nil
}

func (o *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"networkId",
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

	varUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner := _UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner)

	if err != nil {
		return err
	}

	*o = UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner(varUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner)

	return err
}

type NullableUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner struct {
	value *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner
	isSet bool
}

func (v NullableUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) Get() *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner {
	return v.value
}

func (v *NullableUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) Set(val *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner(val *UpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) *NullableUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner {
	return &NullableUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner{value: val, isSet: true}
}

func (v NullableUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationSmSentryPoliciesAssignmentsRequestItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


