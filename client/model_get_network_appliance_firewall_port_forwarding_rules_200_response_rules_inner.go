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

// checks if the GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner{}

// GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner struct for GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner
type GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner struct {
	// IP address of the device subject to port forwarding
	LanIp *string `json:"lanIp,omitempty"`
	// An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges (or any)
	AllowedIps []string `json:"allowedIps,omitempty"`
	// Name of the rule
	Name *string `json:"name,omitempty"`
	// Protocol the rule applies to
	Protocol *string `json:"protocol,omitempty"`
	// The port or port range forwarded to the host on the LAN
	PublicPort *string `json:"publicPort,omitempty"`
	// The port or port range that receives forwarded traffic from the WAN
	LocalPort *string `json:"localPort,omitempty"`
	// The physical WAN interface on which the traffic arrives
	Uplink *string `json:"uplink,omitempty"`
}

// NewGetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner instantiates a new GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner() *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner {
	this := GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner{}
	return &this
}

// NewGetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInnerWithDefaults instantiates a new GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInnerWithDefaults() *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner {
	this := GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner{}
	return &this
}

// GetLanIp returns the LanIp field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) GetLanIp() string {
	if o == nil || IsNil(o.LanIp) {
		var ret string
		return ret
	}
	return *o.LanIp
}

// GetLanIpOk returns a tuple with the LanIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) GetLanIpOk() (*string, bool) {
	if o == nil || IsNil(o.LanIp) {
		return nil, false
	}
	return o.LanIp, true
}

// HasLanIp returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) HasLanIp() bool {
	if o != nil && !IsNil(o.LanIp) {
		return true
	}

	return false
}

// SetLanIp gets a reference to the given string and assigns it to the LanIp field.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) SetLanIp(v string) {
	o.LanIp = &v
}

// GetAllowedIps returns the AllowedIps field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) GetAllowedIps() []string {
	if o == nil || IsNil(o.AllowedIps) {
		var ret []string
		return ret
	}
	return o.AllowedIps
}

// GetAllowedIpsOk returns a tuple with the AllowedIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) GetAllowedIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedIps) {
		return nil, false
	}
	return o.AllowedIps, true
}

// HasAllowedIps returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) HasAllowedIps() bool {
	if o != nil && !IsNil(o.AllowedIps) {
		return true
	}

	return false
}

// SetAllowedIps gets a reference to the given []string and assigns it to the AllowedIps field.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) SetAllowedIps(v []string) {
	o.AllowedIps = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) SetName(v string) {
	o.Name = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) SetProtocol(v string) {
	o.Protocol = &v
}

// GetPublicPort returns the PublicPort field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) GetPublicPort() string {
	if o == nil || IsNil(o.PublicPort) {
		var ret string
		return ret
	}
	return *o.PublicPort
}

// GetPublicPortOk returns a tuple with the PublicPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) GetPublicPortOk() (*string, bool) {
	if o == nil || IsNil(o.PublicPort) {
		return nil, false
	}
	return o.PublicPort, true
}

// HasPublicPort returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) HasPublicPort() bool {
	if o != nil && !IsNil(o.PublicPort) {
		return true
	}

	return false
}

// SetPublicPort gets a reference to the given string and assigns it to the PublicPort field.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) SetPublicPort(v string) {
	o.PublicPort = &v
}

// GetLocalPort returns the LocalPort field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) GetLocalPort() string {
	if o == nil || IsNil(o.LocalPort) {
		var ret string
		return ret
	}
	return *o.LocalPort
}

// GetLocalPortOk returns a tuple with the LocalPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) GetLocalPortOk() (*string, bool) {
	if o == nil || IsNil(o.LocalPort) {
		return nil, false
	}
	return o.LocalPort, true
}

// HasLocalPort returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) HasLocalPort() bool {
	if o != nil && !IsNil(o.LocalPort) {
		return true
	}

	return false
}

// SetLocalPort gets a reference to the given string and assigns it to the LocalPort field.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) SetLocalPort(v string) {
	o.LocalPort = &v
}

// GetUplink returns the Uplink field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) GetUplink() string {
	if o == nil || IsNil(o.Uplink) {
		var ret string
		return ret
	}
	return *o.Uplink
}

// GetUplinkOk returns a tuple with the Uplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) GetUplinkOk() (*string, bool) {
	if o == nil || IsNil(o.Uplink) {
		return nil, false
	}
	return o.Uplink, true
}

// HasUplink returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) HasUplink() bool {
	if o != nil && !IsNil(o.Uplink) {
		return true
	}

	return false
}

// SetUplink gets a reference to the given string and assigns it to the Uplink field.
func (o *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) SetUplink(v string) {
	o.Uplink = &v
}

func (o GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LanIp) {
		toSerialize["lanIp"] = o.LanIp
	}
	if !IsNil(o.AllowedIps) {
		toSerialize["allowedIps"] = o.AllowedIps
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.PublicPort) {
		toSerialize["publicPort"] = o.PublicPort
	}
	if !IsNil(o.LocalPort) {
		toSerialize["localPort"] = o.LocalPort
	}
	if !IsNil(o.Uplink) {
		toSerialize["uplink"] = o.Uplink
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner struct {
	value *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner
	isSet bool
}

func (v NullableGetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) Get() *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner {
	return v.value
}

func (v *NullableGetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) Set(val *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner(val *GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) *NullableGetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner {
	return &NullableGetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


