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

// checks if the ProvisionNetworkClientsRequestPoliciesBySecurityAppliance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisionNetworkClientsRequestPoliciesBySecurityAppliance{}

// ProvisionNetworkClientsRequestPoliciesBySecurityAppliance An object, describing what the policy-connection association is for the security appliance. (Only relevant if the security appliance is actually within the network)
type ProvisionNetworkClientsRequestPoliciesBySecurityAppliance struct {
	// The policy to apply to the specified client. Can be 'Allowed', 'Blocked' or 'Normal'. Required.
	DevicePolicy *string `json:"devicePolicy,omitempty"`
}

// NewProvisionNetworkClientsRequestPoliciesBySecurityAppliance instantiates a new ProvisionNetworkClientsRequestPoliciesBySecurityAppliance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionNetworkClientsRequestPoliciesBySecurityAppliance() *ProvisionNetworkClientsRequestPoliciesBySecurityAppliance {
	this := ProvisionNetworkClientsRequestPoliciesBySecurityAppliance{}
	return &this
}

// NewProvisionNetworkClientsRequestPoliciesBySecurityApplianceWithDefaults instantiates a new ProvisionNetworkClientsRequestPoliciesBySecurityAppliance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionNetworkClientsRequestPoliciesBySecurityApplianceWithDefaults() *ProvisionNetworkClientsRequestPoliciesBySecurityAppliance {
	this := ProvisionNetworkClientsRequestPoliciesBySecurityAppliance{}
	return &this
}

// GetDevicePolicy returns the DevicePolicy field value if set, zero value otherwise.
func (o *ProvisionNetworkClientsRequestPoliciesBySecurityAppliance) GetDevicePolicy() string {
	if o == nil || IsNil(o.DevicePolicy) {
		var ret string
		return ret
	}
	return *o.DevicePolicy
}

// GetDevicePolicyOk returns a tuple with the DevicePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySecurityAppliance) GetDevicePolicyOk() (*string, bool) {
	if o == nil || IsNil(o.DevicePolicy) {
		return nil, false
	}
	return o.DevicePolicy, true
}

// HasDevicePolicy returns a boolean if a field has been set.
func (o *ProvisionNetworkClientsRequestPoliciesBySecurityAppliance) HasDevicePolicy() bool {
	if o != nil && !IsNil(o.DevicePolicy) {
		return true
	}

	return false
}

// SetDevicePolicy gets a reference to the given string and assigns it to the DevicePolicy field.
func (o *ProvisionNetworkClientsRequestPoliciesBySecurityAppliance) SetDevicePolicy(v string) {
	o.DevicePolicy = &v
}

func (o ProvisionNetworkClientsRequestPoliciesBySecurityAppliance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisionNetworkClientsRequestPoliciesBySecurityAppliance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DevicePolicy) {
		toSerialize["devicePolicy"] = o.DevicePolicy
	}
	return toSerialize, nil
}

type NullableProvisionNetworkClientsRequestPoliciesBySecurityAppliance struct {
	value *ProvisionNetworkClientsRequestPoliciesBySecurityAppliance
	isSet bool
}

func (v NullableProvisionNetworkClientsRequestPoliciesBySecurityAppliance) Get() *ProvisionNetworkClientsRequestPoliciesBySecurityAppliance {
	return v.value
}

func (v *NullableProvisionNetworkClientsRequestPoliciesBySecurityAppliance) Set(val *ProvisionNetworkClientsRequestPoliciesBySecurityAppliance) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionNetworkClientsRequestPoliciesBySecurityAppliance) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionNetworkClientsRequestPoliciesBySecurityAppliance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionNetworkClientsRequestPoliciesBySecurityAppliance(val *ProvisionNetworkClientsRequestPoliciesBySecurityAppliance) *NullableProvisionNetworkClientsRequestPoliciesBySecurityAppliance {
	return &NullableProvisionNetworkClientsRequestPoliciesBySecurityAppliance{value: val, isSet: true}
}

func (v NullableProvisionNetworkClientsRequestPoliciesBySecurityAppliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionNetworkClientsRequestPoliciesBySecurityAppliance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


