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

// checks if the UpdateNetworkApplianceSdwanInternetPolicies200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceSdwanInternetPolicies200Response{}

// UpdateNetworkApplianceSdwanInternetPolicies200Response struct for UpdateNetworkApplianceSdwanInternetPolicies200Response
type UpdateNetworkApplianceSdwanInternetPolicies200Response struct {
	// policies with respective traffic filters for an MX network
	WanTrafficUplinkPreferences []UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner `json:"wanTrafficUplinkPreferences,omitempty"`
}

// NewUpdateNetworkApplianceSdwanInternetPolicies200Response instantiates a new UpdateNetworkApplianceSdwanInternetPolicies200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceSdwanInternetPolicies200Response() *UpdateNetworkApplianceSdwanInternetPolicies200Response {
	this := UpdateNetworkApplianceSdwanInternetPolicies200Response{}
	return &this
}

// NewUpdateNetworkApplianceSdwanInternetPolicies200ResponseWithDefaults instantiates a new UpdateNetworkApplianceSdwanInternetPolicies200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceSdwanInternetPolicies200ResponseWithDefaults() *UpdateNetworkApplianceSdwanInternetPolicies200Response {
	this := UpdateNetworkApplianceSdwanInternetPolicies200Response{}
	return &this
}

// GetWanTrafficUplinkPreferences returns the WanTrafficUplinkPreferences field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSdwanInternetPolicies200Response) GetWanTrafficUplinkPreferences() []UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner {
	if o == nil || IsNil(o.WanTrafficUplinkPreferences) {
		var ret []UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner
		return ret
	}
	return o.WanTrafficUplinkPreferences
}

// GetWanTrafficUplinkPreferencesOk returns a tuple with the WanTrafficUplinkPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSdwanInternetPolicies200Response) GetWanTrafficUplinkPreferencesOk() ([]UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner, bool) {
	if o == nil || IsNil(o.WanTrafficUplinkPreferences) {
		return nil, false
	}
	return o.WanTrafficUplinkPreferences, true
}

// HasWanTrafficUplinkPreferences returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSdwanInternetPolicies200Response) HasWanTrafficUplinkPreferences() bool {
	if o != nil && !IsNil(o.WanTrafficUplinkPreferences) {
		return true
	}

	return false
}

// SetWanTrafficUplinkPreferences gets a reference to the given []UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner and assigns it to the WanTrafficUplinkPreferences field.
func (o *UpdateNetworkApplianceSdwanInternetPolicies200Response) SetWanTrafficUplinkPreferences(v []UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner) {
	o.WanTrafficUplinkPreferences = v
}

func (o UpdateNetworkApplianceSdwanInternetPolicies200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceSdwanInternetPolicies200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WanTrafficUplinkPreferences) {
		toSerialize["wanTrafficUplinkPreferences"] = o.WanTrafficUplinkPreferences
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceSdwanInternetPolicies200Response struct {
	value *UpdateNetworkApplianceSdwanInternetPolicies200Response
	isSet bool
}

func (v NullableUpdateNetworkApplianceSdwanInternetPolicies200Response) Get() *UpdateNetworkApplianceSdwanInternetPolicies200Response {
	return v.value
}

func (v *NullableUpdateNetworkApplianceSdwanInternetPolicies200Response) Set(val *UpdateNetworkApplianceSdwanInternetPolicies200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceSdwanInternetPolicies200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceSdwanInternetPolicies200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceSdwanInternetPolicies200Response(val *UpdateNetworkApplianceSdwanInternetPolicies200Response) *NullableUpdateNetworkApplianceSdwanInternetPolicies200Response {
	return &NullableUpdateNetworkApplianceSdwanInternetPolicies200Response{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceSdwanInternetPolicies200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceSdwanInternetPolicies200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


