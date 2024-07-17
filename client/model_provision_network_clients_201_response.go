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

// checks if the ProvisionNetworkClients201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisionNetworkClients201Response{}

// ProvisionNetworkClients201Response struct for ProvisionNetworkClients201Response
type ProvisionNetworkClients201Response struct {
	// The list of clients to provision
	Clients []ProvisionNetworkClients201ResponseClientsInner `json:"clients,omitempty"`
	// The name of the client's policy
	DevicePolicy *string `json:"devicePolicy,omitempty"`
	// The group policy identifier of the client
	GroupPolicyId *string `json:"groupPolicyId,omitempty"`
}

// NewProvisionNetworkClients201Response instantiates a new ProvisionNetworkClients201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionNetworkClients201Response() *ProvisionNetworkClients201Response {
	this := ProvisionNetworkClients201Response{}
	return &this
}

// NewProvisionNetworkClients201ResponseWithDefaults instantiates a new ProvisionNetworkClients201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionNetworkClients201ResponseWithDefaults() *ProvisionNetworkClients201Response {
	this := ProvisionNetworkClients201Response{}
	return &this
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *ProvisionNetworkClients201Response) GetClients() []ProvisionNetworkClients201ResponseClientsInner {
	if o == nil || IsNil(o.Clients) {
		var ret []ProvisionNetworkClients201ResponseClientsInner
		return ret
	}
	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClients201Response) GetClientsOk() ([]ProvisionNetworkClients201ResponseClientsInner, bool) {
	if o == nil || IsNil(o.Clients) {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *ProvisionNetworkClients201Response) HasClients() bool {
	if o != nil && !IsNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given []ProvisionNetworkClients201ResponseClientsInner and assigns it to the Clients field.
func (o *ProvisionNetworkClients201Response) SetClients(v []ProvisionNetworkClients201ResponseClientsInner) {
	o.Clients = v
}

// GetDevicePolicy returns the DevicePolicy field value if set, zero value otherwise.
func (o *ProvisionNetworkClients201Response) GetDevicePolicy() string {
	if o == nil || IsNil(o.DevicePolicy) {
		var ret string
		return ret
	}
	return *o.DevicePolicy
}

// GetDevicePolicyOk returns a tuple with the DevicePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClients201Response) GetDevicePolicyOk() (*string, bool) {
	if o == nil || IsNil(o.DevicePolicy) {
		return nil, false
	}
	return o.DevicePolicy, true
}

// HasDevicePolicy returns a boolean if a field has been set.
func (o *ProvisionNetworkClients201Response) HasDevicePolicy() bool {
	if o != nil && !IsNil(o.DevicePolicy) {
		return true
	}

	return false
}

// SetDevicePolicy gets a reference to the given string and assigns it to the DevicePolicy field.
func (o *ProvisionNetworkClients201Response) SetDevicePolicy(v string) {
	o.DevicePolicy = &v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *ProvisionNetworkClients201Response) GetGroupPolicyId() string {
	if o == nil || IsNil(o.GroupPolicyId) {
		var ret string
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClients201Response) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupPolicyId) {
		return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *ProvisionNetworkClients201Response) HasGroupPolicyId() bool {
	if o != nil && !IsNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given string and assigns it to the GroupPolicyId field.
func (o *ProvisionNetworkClients201Response) SetGroupPolicyId(v string) {
	o.GroupPolicyId = &v
}

func (o ProvisionNetworkClients201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisionNetworkClients201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}
	if !IsNil(o.DevicePolicy) {
		toSerialize["devicePolicy"] = o.DevicePolicy
	}
	if !IsNil(o.GroupPolicyId) {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	return toSerialize, nil
}

type NullableProvisionNetworkClients201Response struct {
	value *ProvisionNetworkClients201Response
	isSet bool
}

func (v NullableProvisionNetworkClients201Response) Get() *ProvisionNetworkClients201Response {
	return v.value
}

func (v *NullableProvisionNetworkClients201Response) Set(val *ProvisionNetworkClients201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionNetworkClients201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionNetworkClients201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionNetworkClients201Response(val *ProvisionNetworkClients201Response) *NullableProvisionNetworkClients201Response {
	return &NullableProvisionNetworkClients201Response{value: val, isSet: true}
}

func (v NullableProvisionNetworkClients201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionNetworkClients201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

