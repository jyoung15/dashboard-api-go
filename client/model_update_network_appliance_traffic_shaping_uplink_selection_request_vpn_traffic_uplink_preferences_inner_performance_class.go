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

// checks if the UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass{}

// UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass Performance class setting for this uplink preference rule
type UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass struct {
	// Type of this performance class. Must be one of: 'builtin' or 'custom'
	Type string `json:"type"`
	// Name of builtin performance class, must be present when performanceClass type is 'builtin', and value must be one of: 'VoIP'
	BuiltinPerformanceClassName *string `json:"builtinPerformanceClassName,omitempty"`
	// ID of created custom performance class, must be present when performanceClass type is 'custom'
	CustomPerformanceClassId *string `json:"customPerformanceClassId,omitempty"`
}

type _UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass

// NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass instantiates a new UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass(type_ string) *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass {
	this := UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass{}
	this.Type = type_
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClassWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClassWithDefaults() *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass {
	this := UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass{}
	return &this
}

// GetType returns the Type field value
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) SetType(v string) {
	o.Type = v
}

// GetBuiltinPerformanceClassName returns the BuiltinPerformanceClassName field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) GetBuiltinPerformanceClassName() string {
	if o == nil || IsNil(o.BuiltinPerformanceClassName) {
		var ret string
		return ret
	}
	return *o.BuiltinPerformanceClassName
}

// GetBuiltinPerformanceClassNameOk returns a tuple with the BuiltinPerformanceClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) GetBuiltinPerformanceClassNameOk() (*string, bool) {
	if o == nil || IsNil(o.BuiltinPerformanceClassName) {
		return nil, false
	}
	return o.BuiltinPerformanceClassName, true
}

// HasBuiltinPerformanceClassName returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) HasBuiltinPerformanceClassName() bool {
	if o != nil && !IsNil(o.BuiltinPerformanceClassName) {
		return true
	}

	return false
}

// SetBuiltinPerformanceClassName gets a reference to the given string and assigns it to the BuiltinPerformanceClassName field.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) SetBuiltinPerformanceClassName(v string) {
	o.BuiltinPerformanceClassName = &v
}

// GetCustomPerformanceClassId returns the CustomPerformanceClassId field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) GetCustomPerformanceClassId() string {
	if o == nil || IsNil(o.CustomPerformanceClassId) {
		var ret string
		return ret
	}
	return *o.CustomPerformanceClassId
}

// GetCustomPerformanceClassIdOk returns a tuple with the CustomPerformanceClassId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) GetCustomPerformanceClassIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomPerformanceClassId) {
		return nil, false
	}
	return o.CustomPerformanceClassId, true
}

// HasCustomPerformanceClassId returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) HasCustomPerformanceClassId() bool {
	if o != nil && !IsNil(o.CustomPerformanceClassId) {
		return true
	}

	return false
}

// SetCustomPerformanceClassId gets a reference to the given string and assigns it to the CustomPerformanceClassId field.
func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) SetCustomPerformanceClassId(v string) {
	o.CustomPerformanceClassId = &v
}

func (o UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.BuiltinPerformanceClassName) {
		toSerialize["builtinPerformanceClassName"] = o.BuiltinPerformanceClassName
	}
	if !IsNil(o.CustomPerformanceClassId) {
		toSerialize["customPerformanceClassId"] = o.CustomPerformanceClassId
	}
	return toSerialize, nil
}

func (o *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass := _UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass)

	if err != nil {
		return err
	}

	*o = UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass(varUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass)

	return err
}

type NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass struct {
	value *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) Get() *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) Set(val *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass(val *UpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) *NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass {
	return &NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkSelectionRequestVpnTrafficUplinkPreferencesInnerPerformanceClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


