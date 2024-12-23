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

// checks if the UpdateNetworkSwitchDhcpServerPolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchDhcpServerPolicyRequest{}

// UpdateNetworkSwitchDhcpServerPolicyRequest struct for UpdateNetworkSwitchDhcpServerPolicyRequest
type UpdateNetworkSwitchDhcpServerPolicyRequest struct {
	Alerts *UpdateNetworkSwitchDhcpServerPolicyRequestAlerts `json:"alerts,omitempty"`
	// 'allow' or 'block' new DHCP servers. Default value is 'allow'.
	DefaultPolicy *string `json:"defaultPolicy,omitempty"`
	// List the MAC addresses of DHCP servers to permit on the network when defaultPolicy is set to block. An empty array will clear the entries.
	AllowedServers []string `json:"allowedServers,omitempty"`
	// List the MAC addresses of DHCP servers to block on the network when defaultPolicy is set to allow. An empty array will clear the entries.
	BlockedServers []string `json:"blockedServers,omitempty"`
	ArpInspection *UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection `json:"arpInspection,omitempty"`
}

// NewUpdateNetworkSwitchDhcpServerPolicyRequest instantiates a new UpdateNetworkSwitchDhcpServerPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchDhcpServerPolicyRequest() *UpdateNetworkSwitchDhcpServerPolicyRequest {
	this := UpdateNetworkSwitchDhcpServerPolicyRequest{}
	return &this
}

// NewUpdateNetworkSwitchDhcpServerPolicyRequestWithDefaults instantiates a new UpdateNetworkSwitchDhcpServerPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchDhcpServerPolicyRequestWithDefaults() *UpdateNetworkSwitchDhcpServerPolicyRequest {
	this := UpdateNetworkSwitchDhcpServerPolicyRequest{}
	return &this
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) GetAlerts() UpdateNetworkSwitchDhcpServerPolicyRequestAlerts {
	if o == nil || IsNil(o.Alerts) {
		var ret UpdateNetworkSwitchDhcpServerPolicyRequestAlerts
		return ret
	}
	return *o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) GetAlertsOk() (*UpdateNetworkSwitchDhcpServerPolicyRequestAlerts, bool) {
	if o == nil || IsNil(o.Alerts) {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) HasAlerts() bool {
	if o != nil && !IsNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given UpdateNetworkSwitchDhcpServerPolicyRequestAlerts and assigns it to the Alerts field.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) SetAlerts(v UpdateNetworkSwitchDhcpServerPolicyRequestAlerts) {
	o.Alerts = &v
}

// GetDefaultPolicy returns the DefaultPolicy field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) GetDefaultPolicy() string {
	if o == nil || IsNil(o.DefaultPolicy) {
		var ret string
		return ret
	}
	return *o.DefaultPolicy
}

// GetDefaultPolicyOk returns a tuple with the DefaultPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) GetDefaultPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultPolicy) {
		return nil, false
	}
	return o.DefaultPolicy, true
}

// HasDefaultPolicy returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) HasDefaultPolicy() bool {
	if o != nil && !IsNil(o.DefaultPolicy) {
		return true
	}

	return false
}

// SetDefaultPolicy gets a reference to the given string and assigns it to the DefaultPolicy field.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) SetDefaultPolicy(v string) {
	o.DefaultPolicy = &v
}

// GetAllowedServers returns the AllowedServers field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) GetAllowedServers() []string {
	if o == nil || IsNil(o.AllowedServers) {
		var ret []string
		return ret
	}
	return o.AllowedServers
}

// GetAllowedServersOk returns a tuple with the AllowedServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) GetAllowedServersOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedServers) {
		return nil, false
	}
	return o.AllowedServers, true
}

// HasAllowedServers returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) HasAllowedServers() bool {
	if o != nil && !IsNil(o.AllowedServers) {
		return true
	}

	return false
}

// SetAllowedServers gets a reference to the given []string and assigns it to the AllowedServers field.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) SetAllowedServers(v []string) {
	o.AllowedServers = v
}

// GetBlockedServers returns the BlockedServers field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) GetBlockedServers() []string {
	if o == nil || IsNil(o.BlockedServers) {
		var ret []string
		return ret
	}
	return o.BlockedServers
}

// GetBlockedServersOk returns a tuple with the BlockedServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) GetBlockedServersOk() ([]string, bool) {
	if o == nil || IsNil(o.BlockedServers) {
		return nil, false
	}
	return o.BlockedServers, true
}

// HasBlockedServers returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) HasBlockedServers() bool {
	if o != nil && !IsNil(o.BlockedServers) {
		return true
	}

	return false
}

// SetBlockedServers gets a reference to the given []string and assigns it to the BlockedServers field.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) SetBlockedServers(v []string) {
	o.BlockedServers = v
}

// GetArpInspection returns the ArpInspection field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) GetArpInspection() UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection {
	if o == nil || IsNil(o.ArpInspection) {
		var ret UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection
		return ret
	}
	return *o.ArpInspection
}

// GetArpInspectionOk returns a tuple with the ArpInspection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) GetArpInspectionOk() (*UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection, bool) {
	if o == nil || IsNil(o.ArpInspection) {
		return nil, false
	}
	return o.ArpInspection, true
}

// HasArpInspection returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) HasArpInspection() bool {
	if o != nil && !IsNil(o.ArpInspection) {
		return true
	}

	return false
}

// SetArpInspection gets a reference to the given UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection and assigns it to the ArpInspection field.
func (o *UpdateNetworkSwitchDhcpServerPolicyRequest) SetArpInspection(v UpdateNetworkSwitchDhcpServerPolicyRequestArpInspection) {
	o.ArpInspection = &v
}

func (o UpdateNetworkSwitchDhcpServerPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchDhcpServerPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alerts) {
		toSerialize["alerts"] = o.Alerts
	}
	if !IsNil(o.DefaultPolicy) {
		toSerialize["defaultPolicy"] = o.DefaultPolicy
	}
	if !IsNil(o.AllowedServers) {
		toSerialize["allowedServers"] = o.AllowedServers
	}
	if !IsNil(o.BlockedServers) {
		toSerialize["blockedServers"] = o.BlockedServers
	}
	if !IsNil(o.ArpInspection) {
		toSerialize["arpInspection"] = o.ArpInspection
	}
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchDhcpServerPolicyRequest struct {
	value *UpdateNetworkSwitchDhcpServerPolicyRequest
	isSet bool
}

func (v NullableUpdateNetworkSwitchDhcpServerPolicyRequest) Get() *UpdateNetworkSwitchDhcpServerPolicyRequest {
	return v.value
}

func (v *NullableUpdateNetworkSwitchDhcpServerPolicyRequest) Set(val *UpdateNetworkSwitchDhcpServerPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchDhcpServerPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchDhcpServerPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchDhcpServerPolicyRequest(val *UpdateNetworkSwitchDhcpServerPolicyRequest) *NullableUpdateNetworkSwitchDhcpServerPolicyRequest {
	return &NullableUpdateNetworkSwitchDhcpServerPolicyRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchDhcpServerPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchDhcpServerPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


