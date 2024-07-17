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

// checks if the GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks{}

// GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks Networks included in and excluded from the detection engine
type GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks struct {
	// Whether special IPv4 addresses should be used (see: https://tools.ietf.org/html/rfc5735)
	UseDefault *bool `json:"useDefault,omitempty"`
	// List of IP addresses or subnets being protected
	IncludedCidr []string `json:"includedCidr,omitempty"`
	// List of IP addresses or subnets being excluded from protection
	ExcludedCidr []string `json:"excludedCidr,omitempty"`
}

// NewGetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks instantiates a new GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks() *GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks {
	this := GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks{}
	return &this
}

// NewGetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworksWithDefaults instantiates a new GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworksWithDefaults() *GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks {
	this := GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks{}
	return &this
}

// GetUseDefault returns the UseDefault field value if set, zero value otherwise.
func (o *GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) GetUseDefault() bool {
	if o == nil || IsNil(o.UseDefault) {
		var ret bool
		return ret
	}
	return *o.UseDefault
}

// GetUseDefaultOk returns a tuple with the UseDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) GetUseDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.UseDefault) {
		return nil, false
	}
	return o.UseDefault, true
}

// HasUseDefault returns a boolean if a field has been set.
func (o *GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) HasUseDefault() bool {
	if o != nil && !IsNil(o.UseDefault) {
		return true
	}

	return false
}

// SetUseDefault gets a reference to the given bool and assigns it to the UseDefault field.
func (o *GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) SetUseDefault(v bool) {
	o.UseDefault = &v
}

// GetIncludedCidr returns the IncludedCidr field value if set, zero value otherwise.
func (o *GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) GetIncludedCidr() []string {
	if o == nil || IsNil(o.IncludedCidr) {
		var ret []string
		return ret
	}
	return o.IncludedCidr
}

// GetIncludedCidrOk returns a tuple with the IncludedCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) GetIncludedCidrOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludedCidr) {
		return nil, false
	}
	return o.IncludedCidr, true
}

// HasIncludedCidr returns a boolean if a field has been set.
func (o *GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) HasIncludedCidr() bool {
	if o != nil && !IsNil(o.IncludedCidr) {
		return true
	}

	return false
}

// SetIncludedCidr gets a reference to the given []string and assigns it to the IncludedCidr field.
func (o *GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) SetIncludedCidr(v []string) {
	o.IncludedCidr = v
}

// GetExcludedCidr returns the ExcludedCidr field value if set, zero value otherwise.
func (o *GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) GetExcludedCidr() []string {
	if o == nil || IsNil(o.ExcludedCidr) {
		var ret []string
		return ret
	}
	return o.ExcludedCidr
}

// GetExcludedCidrOk returns a tuple with the ExcludedCidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) GetExcludedCidrOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludedCidr) {
		return nil, false
	}
	return o.ExcludedCidr, true
}

// HasExcludedCidr returns a boolean if a field has been set.
func (o *GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) HasExcludedCidr() bool {
	if o != nil && !IsNil(o.ExcludedCidr) {
		return true
	}

	return false
}

// SetExcludedCidr gets a reference to the given []string and assigns it to the ExcludedCidr field.
func (o *GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) SetExcludedCidr(v []string) {
	o.ExcludedCidr = v
}

func (o GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UseDefault) {
		toSerialize["useDefault"] = o.UseDefault
	}
	if !IsNil(o.IncludedCidr) {
		toSerialize["includedCidr"] = o.IncludedCidr
	}
	if !IsNil(o.ExcludedCidr) {
		toSerialize["excludedCidr"] = o.ExcludedCidr
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks struct {
	value *GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks
	isSet bool
}

func (v NullableGetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) Get() *GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks {
	return v.value
}

func (v *NullableGetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) Set(val *GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks(val *GetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) *NullableGetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks {
	return &NullableGetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceSecurityIntrusion200ResponseProtectedNetworks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

