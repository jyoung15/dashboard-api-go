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

// checks if the UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass{}

// UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass Performance class setting for uplink preference rule
type UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass struct {
	// Type of this performance class. Must be one of: 'builtin' or 'custom'
	Type *string `json:"type,omitempty"`
	// Name of builtin performance class. Must be present when performanceClass type is 'builtin' and value must be one of: 'VoIP'
	BuiltinPerformanceClassName *string `json:"builtinPerformanceClassName,omitempty"`
	// ID of created custom performance class, must be present when performanceClass type is \"custom\"
	CustomPerformanceClassId *string `json:"customPerformanceClassId,omitempty"`
}

// NewUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass instantiates a new UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass() *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass {
	this := UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass{}
	return &this
}

// NewUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClassWithDefaults instantiates a new UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClassWithDefaults() *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass {
	this := UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) SetType(v string) {
	o.Type = &v
}

// GetBuiltinPerformanceClassName returns the BuiltinPerformanceClassName field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) GetBuiltinPerformanceClassName() string {
	if o == nil || IsNil(o.BuiltinPerformanceClassName) {
		var ret string
		return ret
	}
	return *o.BuiltinPerformanceClassName
}

// GetBuiltinPerformanceClassNameOk returns a tuple with the BuiltinPerformanceClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) GetBuiltinPerformanceClassNameOk() (*string, bool) {
	if o == nil || IsNil(o.BuiltinPerformanceClassName) {
		return nil, false
	}
	return o.BuiltinPerformanceClassName, true
}

// HasBuiltinPerformanceClassName returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) HasBuiltinPerformanceClassName() bool {
	if o != nil && !IsNil(o.BuiltinPerformanceClassName) {
		return true
	}

	return false
}

// SetBuiltinPerformanceClassName gets a reference to the given string and assigns it to the BuiltinPerformanceClassName field.
func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) SetBuiltinPerformanceClassName(v string) {
	o.BuiltinPerformanceClassName = &v
}

// GetCustomPerformanceClassId returns the CustomPerformanceClassId field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) GetCustomPerformanceClassId() string {
	if o == nil || IsNil(o.CustomPerformanceClassId) {
		var ret string
		return ret
	}
	return *o.CustomPerformanceClassId
}

// GetCustomPerformanceClassIdOk returns a tuple with the CustomPerformanceClassId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) GetCustomPerformanceClassIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomPerformanceClassId) {
		return nil, false
	}
	return o.CustomPerformanceClassId, true
}

// HasCustomPerformanceClassId returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) HasCustomPerformanceClassId() bool {
	if o != nil && !IsNil(o.CustomPerformanceClassId) {
		return true
	}

	return false
}

// SetCustomPerformanceClassId gets a reference to the given string and assigns it to the CustomPerformanceClassId field.
func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) SetCustomPerformanceClassId(v string) {
	o.CustomPerformanceClassId = &v
}

func (o UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.BuiltinPerformanceClassName) {
		toSerialize["builtinPerformanceClassName"] = o.BuiltinPerformanceClassName
	}
	if !IsNil(o.CustomPerformanceClassId) {
		toSerialize["customPerformanceClassId"] = o.CustomPerformanceClassId
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass struct {
	value *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass
	isSet bool
}

func (v NullableUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) Get() *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass {
	return v.value
}

func (v *NullableUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) Set(val *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass(val *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) *NullableUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass {
	return &NullableUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


