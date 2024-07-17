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

// checks if the UpdateNetworkWirelessSsidRequestDnsRewrite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidRequestDnsRewrite{}

// UpdateNetworkWirelessSsidRequestDnsRewrite DNS servers rewrite settings
type UpdateNetworkWirelessSsidRequestDnsRewrite struct {
	// Boolean indicating whether or not DNS server rewrite is enabled. If disabled, upstream DNS will be used
	Enabled *bool `json:"enabled,omitempty"`
	// User specified DNS servers (up to two servers)
	DnsCustomNameservers []string `json:"dnsCustomNameservers,omitempty"`
}

// NewUpdateNetworkWirelessSsidRequestDnsRewrite instantiates a new UpdateNetworkWirelessSsidRequestDnsRewrite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestDnsRewrite() *UpdateNetworkWirelessSsidRequestDnsRewrite {
	this := UpdateNetworkWirelessSsidRequestDnsRewrite{}
	return &this
}

// NewUpdateNetworkWirelessSsidRequestDnsRewriteWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestDnsRewrite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestDnsRewriteWithDefaults() *UpdateNetworkWirelessSsidRequestDnsRewrite {
	this := UpdateNetworkWirelessSsidRequestDnsRewrite{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestDnsRewrite) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestDnsRewrite) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestDnsRewrite) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkWirelessSsidRequestDnsRewrite) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDnsCustomNameservers returns the DnsCustomNameservers field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestDnsRewrite) GetDnsCustomNameservers() []string {
	if o == nil || IsNil(o.DnsCustomNameservers) {
		var ret []string
		return ret
	}
	return o.DnsCustomNameservers
}

// GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestDnsRewrite) GetDnsCustomNameserversOk() ([]string, bool) {
	if o == nil || IsNil(o.DnsCustomNameservers) {
		return nil, false
	}
	return o.DnsCustomNameservers, true
}

// HasDnsCustomNameservers returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestDnsRewrite) HasDnsCustomNameservers() bool {
	if o != nil && !IsNil(o.DnsCustomNameservers) {
		return true
	}

	return false
}

// SetDnsCustomNameservers gets a reference to the given []string and assigns it to the DnsCustomNameservers field.
func (o *UpdateNetworkWirelessSsidRequestDnsRewrite) SetDnsCustomNameservers(v []string) {
	o.DnsCustomNameservers = v
}

func (o UpdateNetworkWirelessSsidRequestDnsRewrite) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidRequestDnsRewrite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.DnsCustomNameservers) {
		toSerialize["dnsCustomNameservers"] = o.DnsCustomNameservers
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidRequestDnsRewrite struct {
	value *UpdateNetworkWirelessSsidRequestDnsRewrite
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestDnsRewrite) Get() *UpdateNetworkWirelessSsidRequestDnsRewrite {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestDnsRewrite) Set(val *UpdateNetworkWirelessSsidRequestDnsRewrite) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestDnsRewrite) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestDnsRewrite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestDnsRewrite(val *UpdateNetworkWirelessSsidRequestDnsRewrite) *NullableUpdateNetworkWirelessSsidRequestDnsRewrite {
	return &NullableUpdateNetworkWirelessSsidRequestDnsRewrite{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestDnsRewrite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestDnsRewrite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


