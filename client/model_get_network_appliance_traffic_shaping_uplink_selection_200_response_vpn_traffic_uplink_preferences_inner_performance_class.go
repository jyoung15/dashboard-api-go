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

// GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass Performance class setting for uplink preference rule
type GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass struct {
	// Type of this performance class. Must be one of: 'builtin' or 'custom'
	Type string `json:"type"`
	// Name of builtin performance class. Must be present when performanceClass type is 'builtin' and value must be one of: 'VoIP'
	BuiltinPerformanceClassName *string `json:"builtinPerformanceClassName,omitempty"`
	// ID of created custom performance class, must be present when performanceClass type is \"custom\"
	CustomPerformanceClassId *string `json:"customPerformanceClassId,omitempty"`
}

// NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass instantiates a new GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass(type_ string) *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass {
	this := GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass{}
	this.Type = type_
	return &this
}

// NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClassWithDefaults instantiates a new GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClassWithDefaults() *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass {
	this := GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass{}
	return &this
}

// GetType returns the Type field value
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass) GetTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass) SetType(v string) {
	o.Type = v
}

// GetBuiltinPerformanceClassName returns the BuiltinPerformanceClassName field value if set, zero value otherwise.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass) GetBuiltinPerformanceClassName() string {
	if o == nil || isNil(o.BuiltinPerformanceClassName) {
		var ret string
		return ret
	}
	return *o.BuiltinPerformanceClassName
}

// GetBuiltinPerformanceClassNameOk returns a tuple with the BuiltinPerformanceClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass) GetBuiltinPerformanceClassNameOk() (*string, bool) {
	if o == nil || isNil(o.BuiltinPerformanceClassName) {
    return nil, false
	}
	return o.BuiltinPerformanceClassName, true
}

// HasBuiltinPerformanceClassName returns a boolean if a field has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass) HasBuiltinPerformanceClassName() bool {
	if o != nil && !isNil(o.BuiltinPerformanceClassName) {
		return true
	}

	return false
}

// SetBuiltinPerformanceClassName gets a reference to the given string and assigns it to the BuiltinPerformanceClassName field.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass) SetBuiltinPerformanceClassName(v string) {
	o.BuiltinPerformanceClassName = &v
}

// GetCustomPerformanceClassId returns the CustomPerformanceClassId field value if set, zero value otherwise.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass) GetCustomPerformanceClassId() string {
	if o == nil || isNil(o.CustomPerformanceClassId) {
		var ret string
		return ret
	}
	return *o.CustomPerformanceClassId
}

// GetCustomPerformanceClassIdOk returns a tuple with the CustomPerformanceClassId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass) GetCustomPerformanceClassIdOk() (*string, bool) {
	if o == nil || isNil(o.CustomPerformanceClassId) {
    return nil, false
	}
	return o.CustomPerformanceClassId, true
}

// HasCustomPerformanceClassId returns a boolean if a field has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass) HasCustomPerformanceClassId() bool {
	if o != nil && !isNil(o.CustomPerformanceClassId) {
		return true
	}

	return false
}

// SetCustomPerformanceClassId gets a reference to the given string and assigns it to the CustomPerformanceClassId field.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass) SetCustomPerformanceClassId(v string) {
	o.CustomPerformanceClassId = &v
}

func (o GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.BuiltinPerformanceClassName) {
		toSerialize["builtinPerformanceClassName"] = o.BuiltinPerformanceClassName
	}
	if !isNil(o.CustomPerformanceClassId) {
		toSerialize["customPerformanceClassId"] = o.CustomPerformanceClassId
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass struct {
	value *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass
	isSet bool
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass) Get() *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass {
	return v.value
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass) Set(val *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass(val *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass) *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass {
	return &NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInnerPerformanceClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


