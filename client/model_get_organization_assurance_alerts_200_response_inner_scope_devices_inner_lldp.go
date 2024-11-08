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

// checks if the GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp{}

// GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp Port of affected device
type GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp struct {
	// Port of affect device
	Port *string `json:"port,omitempty"`
	// Port ID of affect device
	PortId *string `json:"portId,omitempty"`
}

// NewGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp instantiates a new GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp() *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp {
	this := GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp{}
	return &this
}

// NewGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldpWithDefaults instantiates a new GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldpWithDefaults() *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp {
	this := GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp) GetPort() string {
	if o == nil || IsNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp) GetPortOk() (*string, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp) SetPort(v string) {
	o.Port = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp) GetPortId() string {
	if o == nil || IsNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp) GetPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortId) {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp) HasPortId() bool {
	if o != nil && !IsNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp) SetPortId(v string) {
	o.PortId = &v
}

func (o GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.PortId) {
		toSerialize["portId"] = o.PortId
	}
	return toSerialize, nil
}

type NullableGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp struct {
	value *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp
	isSet bool
}

func (v NullableGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp) Get() *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp {
	return v.value
}

func (v *NullableGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp) Set(val *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp(val *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp) *NullableGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp {
	return &NullableGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp{value: val, isSet: true}
}

func (v NullableGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


