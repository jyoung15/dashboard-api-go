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

// checks if the UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest{}

// UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest struct for UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest
type UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest struct {
	// The DHCP mode options for the switch stack interface         ('dhcpDisabled', 'dhcpRelay' or 'dhcpServer')
	DhcpMode *string `json:"dhcpMode,omitempty"`
	// The DHCP relay server IPs to which DHCP packets would get relayed for the switch stack interface
	DhcpRelayServerIps []string `json:"dhcpRelayServerIps,omitempty"`
	// The DHCP lease time config for the dhcp server running on switch stack interface         ('30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week')
	DhcpLeaseTime *string `json:"dhcpLeaseTime,omitempty"`
	// The DHCP name server option for the dhcp server running on the switch stack interface         ('googlePublicDns', 'openDns' or 'custom')
	DnsNameserversOption *string `json:"dnsNameserversOption,omitempty"`
	// The DHCP name server IPs when DHCP name server option is '         custom'
	DnsCustomNameservers []string `json:"dnsCustomNameservers,omitempty"`
	// Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch         stack interface
	BootOptionsEnabled *bool `json:"bootOptionsEnabled,omitempty"`
	// The PXE boot server IP for the DHCP server running on the switch stack interface
	BootNextServer *string `json:"bootNextServer,omitempty"`
	// The PXE boot server file name for the DHCP server running on the switch stack interface
	BootFileName *string `json:"bootFileName,omitempty"`
	// Array of DHCP options consisting of code, type and value for the DHCP server running on the         switch stack interface
	DhcpOptions []UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner `json:"dhcpOptions,omitempty"`
	// Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface
	ReservedIpRanges []UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner `json:"reservedIpRanges,omitempty"`
	// Array of DHCP fixed IP assignments for the DHCP server running on the switch stack interface
	FixedIpAssignments []UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner `json:"fixedIpAssignments,omitempty"`
}

// NewUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest instantiates a new UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest() *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest {
	this := UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest{}
	return &this
}

// NewUpdateNetworkSwitchStackRoutingInterfaceDhcpRequestWithDefaults instantiates a new UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchStackRoutingInterfaceDhcpRequestWithDefaults() *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest {
	this := UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest{}
	return &this
}

// GetDhcpMode returns the DhcpMode field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDhcpMode() string {
	if o == nil || IsNil(o.DhcpMode) {
		var ret string
		return ret
	}
	return *o.DhcpMode
}

// GetDhcpModeOk returns a tuple with the DhcpMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDhcpModeOk() (*string, bool) {
	if o == nil || IsNil(o.DhcpMode) {
		return nil, false
	}
	return o.DhcpMode, true
}

// HasDhcpMode returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasDhcpMode() bool {
	if o != nil && !IsNil(o.DhcpMode) {
		return true
	}

	return false
}

// SetDhcpMode gets a reference to the given string and assigns it to the DhcpMode field.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetDhcpMode(v string) {
	o.DhcpMode = &v
}

// GetDhcpRelayServerIps returns the DhcpRelayServerIps field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDhcpRelayServerIps() []string {
	if o == nil || IsNil(o.DhcpRelayServerIps) {
		var ret []string
		return ret
	}
	return o.DhcpRelayServerIps
}

// GetDhcpRelayServerIpsOk returns a tuple with the DhcpRelayServerIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDhcpRelayServerIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.DhcpRelayServerIps) {
		return nil, false
	}
	return o.DhcpRelayServerIps, true
}

// HasDhcpRelayServerIps returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasDhcpRelayServerIps() bool {
	if o != nil && !IsNil(o.DhcpRelayServerIps) {
		return true
	}

	return false
}

// SetDhcpRelayServerIps gets a reference to the given []string and assigns it to the DhcpRelayServerIps field.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetDhcpRelayServerIps(v []string) {
	o.DhcpRelayServerIps = v
}

// GetDhcpLeaseTime returns the DhcpLeaseTime field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDhcpLeaseTime() string {
	if o == nil || IsNil(o.DhcpLeaseTime) {
		var ret string
		return ret
	}
	return *o.DhcpLeaseTime
}

// GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDhcpLeaseTimeOk() (*string, bool) {
	if o == nil || IsNil(o.DhcpLeaseTime) {
		return nil, false
	}
	return o.DhcpLeaseTime, true
}

// HasDhcpLeaseTime returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasDhcpLeaseTime() bool {
	if o != nil && !IsNil(o.DhcpLeaseTime) {
		return true
	}

	return false
}

// SetDhcpLeaseTime gets a reference to the given string and assigns it to the DhcpLeaseTime field.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetDhcpLeaseTime(v string) {
	o.DhcpLeaseTime = &v
}

// GetDnsNameserversOption returns the DnsNameserversOption field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDnsNameserversOption() string {
	if o == nil || IsNil(o.DnsNameserversOption) {
		var ret string
		return ret
	}
	return *o.DnsNameserversOption
}

// GetDnsNameserversOptionOk returns a tuple with the DnsNameserversOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDnsNameserversOptionOk() (*string, bool) {
	if o == nil || IsNil(o.DnsNameserversOption) {
		return nil, false
	}
	return o.DnsNameserversOption, true
}

// HasDnsNameserversOption returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasDnsNameserversOption() bool {
	if o != nil && !IsNil(o.DnsNameserversOption) {
		return true
	}

	return false
}

// SetDnsNameserversOption gets a reference to the given string and assigns it to the DnsNameserversOption field.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetDnsNameserversOption(v string) {
	o.DnsNameserversOption = &v
}

// GetDnsCustomNameservers returns the DnsCustomNameservers field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDnsCustomNameservers() []string {
	if o == nil || IsNil(o.DnsCustomNameservers) {
		var ret []string
		return ret
	}
	return o.DnsCustomNameservers
}

// GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDnsCustomNameserversOk() ([]string, bool) {
	if o == nil || IsNil(o.DnsCustomNameservers) {
		return nil, false
	}
	return o.DnsCustomNameservers, true
}

// HasDnsCustomNameservers returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasDnsCustomNameservers() bool {
	if o != nil && !IsNil(o.DnsCustomNameservers) {
		return true
	}

	return false
}

// SetDnsCustomNameservers gets a reference to the given []string and assigns it to the DnsCustomNameservers field.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetDnsCustomNameservers(v []string) {
	o.DnsCustomNameservers = v
}

// GetBootOptionsEnabled returns the BootOptionsEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetBootOptionsEnabled() bool {
	if o == nil || IsNil(o.BootOptionsEnabled) {
		var ret bool
		return ret
	}
	return *o.BootOptionsEnabled
}

// GetBootOptionsEnabledOk returns a tuple with the BootOptionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetBootOptionsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BootOptionsEnabled) {
		return nil, false
	}
	return o.BootOptionsEnabled, true
}

// HasBootOptionsEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasBootOptionsEnabled() bool {
	if o != nil && !IsNil(o.BootOptionsEnabled) {
		return true
	}

	return false
}

// SetBootOptionsEnabled gets a reference to the given bool and assigns it to the BootOptionsEnabled field.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetBootOptionsEnabled(v bool) {
	o.BootOptionsEnabled = &v
}

// GetBootNextServer returns the BootNextServer field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetBootNextServer() string {
	if o == nil || IsNil(o.BootNextServer) {
		var ret string
		return ret
	}
	return *o.BootNextServer
}

// GetBootNextServerOk returns a tuple with the BootNextServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetBootNextServerOk() (*string, bool) {
	if o == nil || IsNil(o.BootNextServer) {
		return nil, false
	}
	return o.BootNextServer, true
}

// HasBootNextServer returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasBootNextServer() bool {
	if o != nil && !IsNil(o.BootNextServer) {
		return true
	}

	return false
}

// SetBootNextServer gets a reference to the given string and assigns it to the BootNextServer field.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetBootNextServer(v string) {
	o.BootNextServer = &v
}

// GetBootFileName returns the BootFileName field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetBootFileName() string {
	if o == nil || IsNil(o.BootFileName) {
		var ret string
		return ret
	}
	return *o.BootFileName
}

// GetBootFileNameOk returns a tuple with the BootFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetBootFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.BootFileName) {
		return nil, false
	}
	return o.BootFileName, true
}

// HasBootFileName returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasBootFileName() bool {
	if o != nil && !IsNil(o.BootFileName) {
		return true
	}

	return false
}

// SetBootFileName gets a reference to the given string and assigns it to the BootFileName field.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetBootFileName(v string) {
	o.BootFileName = &v
}

// GetDhcpOptions returns the DhcpOptions field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDhcpOptions() []UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner {
	if o == nil || IsNil(o.DhcpOptions) {
		var ret []UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner
		return ret
	}
	return o.DhcpOptions
}

// GetDhcpOptionsOk returns a tuple with the DhcpOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetDhcpOptionsOk() ([]UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner, bool) {
	if o == nil || IsNil(o.DhcpOptions) {
		return nil, false
	}
	return o.DhcpOptions, true
}

// HasDhcpOptions returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasDhcpOptions() bool {
	if o != nil && !IsNil(o.DhcpOptions) {
		return true
	}

	return false
}

// SetDhcpOptions gets a reference to the given []UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner and assigns it to the DhcpOptions field.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetDhcpOptions(v []UpdateDeviceSwitchRoutingInterfaceDhcpRequestDhcpOptionsInner) {
	o.DhcpOptions = v
}

// GetReservedIpRanges returns the ReservedIpRanges field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetReservedIpRanges() []UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner {
	if o == nil || IsNil(o.ReservedIpRanges) {
		var ret []UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner
		return ret
	}
	return o.ReservedIpRanges
}

// GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetReservedIpRangesOk() ([]UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner, bool) {
	if o == nil || IsNil(o.ReservedIpRanges) {
		return nil, false
	}
	return o.ReservedIpRanges, true
}

// HasReservedIpRanges returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasReservedIpRanges() bool {
	if o != nil && !IsNil(o.ReservedIpRanges) {
		return true
	}

	return false
}

// SetReservedIpRanges gets a reference to the given []UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner and assigns it to the ReservedIpRanges field.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetReservedIpRanges(v []UpdateDeviceSwitchRoutingInterfaceDhcpRequestReservedIpRangesInner) {
	o.ReservedIpRanges = v
}

// GetFixedIpAssignments returns the FixedIpAssignments field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetFixedIpAssignments() []UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner {
	if o == nil || IsNil(o.FixedIpAssignments) {
		var ret []UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner
		return ret
	}
	return o.FixedIpAssignments
}

// GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) GetFixedIpAssignmentsOk() ([]UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner, bool) {
	if o == nil || IsNil(o.FixedIpAssignments) {
		return nil, false
	}
	return o.FixedIpAssignments, true
}

// HasFixedIpAssignments returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) HasFixedIpAssignments() bool {
	if o != nil && !IsNil(o.FixedIpAssignments) {
		return true
	}

	return false
}

// SetFixedIpAssignments gets a reference to the given []UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner and assigns it to the FixedIpAssignments field.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) SetFixedIpAssignments(v []UpdateDeviceSwitchRoutingInterfaceDhcpRequestFixedIpAssignmentsInner) {
	o.FixedIpAssignments = v
}

func (o UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DhcpMode) {
		toSerialize["dhcpMode"] = o.DhcpMode
	}
	if !IsNil(o.DhcpRelayServerIps) {
		toSerialize["dhcpRelayServerIps"] = o.DhcpRelayServerIps
	}
	if !IsNil(o.DhcpLeaseTime) {
		toSerialize["dhcpLeaseTime"] = o.DhcpLeaseTime
	}
	if !IsNil(o.DnsNameserversOption) {
		toSerialize["dnsNameserversOption"] = o.DnsNameserversOption
	}
	if !IsNil(o.DnsCustomNameservers) {
		toSerialize["dnsCustomNameservers"] = o.DnsCustomNameservers
	}
	if !IsNil(o.BootOptionsEnabled) {
		toSerialize["bootOptionsEnabled"] = o.BootOptionsEnabled
	}
	if !IsNil(o.BootNextServer) {
		toSerialize["bootNextServer"] = o.BootNextServer
	}
	if !IsNil(o.BootFileName) {
		toSerialize["bootFileName"] = o.BootFileName
	}
	if !IsNil(o.DhcpOptions) {
		toSerialize["dhcpOptions"] = o.DhcpOptions
	}
	if !IsNil(o.ReservedIpRanges) {
		toSerialize["reservedIpRanges"] = o.ReservedIpRanges
	}
	if !IsNil(o.FixedIpAssignments) {
		toSerialize["fixedIpAssignments"] = o.FixedIpAssignments
	}
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest struct {
	value *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest
	isSet bool
}

func (v NullableUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) Get() *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest {
	return v.value
}

func (v *NullableUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) Set(val *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest(val *UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) *NullableUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest {
	return &NullableUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


