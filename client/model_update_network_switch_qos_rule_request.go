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

// UpdateNetworkSwitchQosRuleRequest struct for UpdateNetworkSwitchQosRuleRequest
type UpdateNetworkSwitchQosRuleRequest struct {
	// The VLAN of the incoming packet. A null value will match any VLAN.
	Vlan *int32 `json:"vlan,omitempty"`
	// The protocol of the incoming packet. Can be one of \"ANY\", \"TCP\" or \"UDP\". Default value is \"ANY\".
	Protocol *string `json:"protocol,omitempty"`
	// The source port of the incoming packet. Applicable only if protocol is TCP or UDP.
	SrcPort *int32 `json:"srcPort,omitempty"`
	// The source port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	SrcPortRange *string `json:"srcPortRange,omitempty"`
	// The destination port of the incoming packet. Applicable only if protocol is TCP or UDP.
	DstPort *int32 `json:"dstPort,omitempty"`
	// The destination port range of the incoming packet. Applicable only if protocol is set to TCP or UDP. Example: 70-80
	DstPortRange *string `json:"dstPortRange,omitempty"`
	// DSCP tag that should be assigned to incoming packet. Set this to -1 to trust incoming DSCP. Default value is 0.
	Dscp *int32 `json:"dscp,omitempty"`
}

// NewUpdateNetworkSwitchQosRuleRequest instantiates a new UpdateNetworkSwitchQosRuleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchQosRuleRequest() *UpdateNetworkSwitchQosRuleRequest {
	this := UpdateNetworkSwitchQosRuleRequest{}
	return &this
}

// NewUpdateNetworkSwitchQosRuleRequestWithDefaults instantiates a new UpdateNetworkSwitchQosRuleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchQosRuleRequestWithDefaults() *UpdateNetworkSwitchQosRuleRequest {
	this := UpdateNetworkSwitchQosRuleRequest{}
	return &this
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchQosRuleRequest) GetVlan() int32 {
	if o == nil || isNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchQosRuleRequest) GetVlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vlan) {
    return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchQosRuleRequest) HasVlan() bool {
	if o != nil && !isNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *UpdateNetworkSwitchQosRuleRequest) SetVlan(v int32) {
	o.Vlan = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchQosRuleRequest) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchQosRuleRequest) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchQosRuleRequest) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *UpdateNetworkSwitchQosRuleRequest) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSrcPort returns the SrcPort field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchQosRuleRequest) GetSrcPort() int32 {
	if o == nil || isNil(o.SrcPort) {
		var ret int32
		return ret
	}
	return *o.SrcPort
}

// GetSrcPortOk returns a tuple with the SrcPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchQosRuleRequest) GetSrcPortOk() (*int32, bool) {
	if o == nil || isNil(o.SrcPort) {
    return nil, false
	}
	return o.SrcPort, true
}

// HasSrcPort returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchQosRuleRequest) HasSrcPort() bool {
	if o != nil && !isNil(o.SrcPort) {
		return true
	}

	return false
}

// SetSrcPort gets a reference to the given int32 and assigns it to the SrcPort field.
func (o *UpdateNetworkSwitchQosRuleRequest) SetSrcPort(v int32) {
	o.SrcPort = &v
}

// GetSrcPortRange returns the SrcPortRange field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchQosRuleRequest) GetSrcPortRange() string {
	if o == nil || isNil(o.SrcPortRange) {
		var ret string
		return ret
	}
	return *o.SrcPortRange
}

// GetSrcPortRangeOk returns a tuple with the SrcPortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchQosRuleRequest) GetSrcPortRangeOk() (*string, bool) {
	if o == nil || isNil(o.SrcPortRange) {
    return nil, false
	}
	return o.SrcPortRange, true
}

// HasSrcPortRange returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchQosRuleRequest) HasSrcPortRange() bool {
	if o != nil && !isNil(o.SrcPortRange) {
		return true
	}

	return false
}

// SetSrcPortRange gets a reference to the given string and assigns it to the SrcPortRange field.
func (o *UpdateNetworkSwitchQosRuleRequest) SetSrcPortRange(v string) {
	o.SrcPortRange = &v
}

// GetDstPort returns the DstPort field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchQosRuleRequest) GetDstPort() int32 {
	if o == nil || isNil(o.DstPort) {
		var ret int32
		return ret
	}
	return *o.DstPort
}

// GetDstPortOk returns a tuple with the DstPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchQosRuleRequest) GetDstPortOk() (*int32, bool) {
	if o == nil || isNil(o.DstPort) {
    return nil, false
	}
	return o.DstPort, true
}

// HasDstPort returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchQosRuleRequest) HasDstPort() bool {
	if o != nil && !isNil(o.DstPort) {
		return true
	}

	return false
}

// SetDstPort gets a reference to the given int32 and assigns it to the DstPort field.
func (o *UpdateNetworkSwitchQosRuleRequest) SetDstPort(v int32) {
	o.DstPort = &v
}

// GetDstPortRange returns the DstPortRange field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchQosRuleRequest) GetDstPortRange() string {
	if o == nil || isNil(o.DstPortRange) {
		var ret string
		return ret
	}
	return *o.DstPortRange
}

// GetDstPortRangeOk returns a tuple with the DstPortRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchQosRuleRequest) GetDstPortRangeOk() (*string, bool) {
	if o == nil || isNil(o.DstPortRange) {
    return nil, false
	}
	return o.DstPortRange, true
}

// HasDstPortRange returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchQosRuleRequest) HasDstPortRange() bool {
	if o != nil && !isNil(o.DstPortRange) {
		return true
	}

	return false
}

// SetDstPortRange gets a reference to the given string and assigns it to the DstPortRange field.
func (o *UpdateNetworkSwitchQosRuleRequest) SetDstPortRange(v string) {
	o.DstPortRange = &v
}

// GetDscp returns the Dscp field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchQosRuleRequest) GetDscp() int32 {
	if o == nil || isNil(o.Dscp) {
		var ret int32
		return ret
	}
	return *o.Dscp
}

// GetDscpOk returns a tuple with the Dscp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchQosRuleRequest) GetDscpOk() (*int32, bool) {
	if o == nil || isNil(o.Dscp) {
    return nil, false
	}
	return o.Dscp, true
}

// HasDscp returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchQosRuleRequest) HasDscp() bool {
	if o != nil && !isNil(o.Dscp) {
		return true
	}

	return false
}

// SetDscp gets a reference to the given int32 and assigns it to the Dscp field.
func (o *UpdateNetworkSwitchQosRuleRequest) SetDscp(v int32) {
	o.Dscp = &v
}

func (o UpdateNetworkSwitchQosRuleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.SrcPort) {
		toSerialize["srcPort"] = o.SrcPort
	}
	if !isNil(o.SrcPortRange) {
		toSerialize["srcPortRange"] = o.SrcPortRange
	}
	if !isNil(o.DstPort) {
		toSerialize["dstPort"] = o.DstPort
	}
	if !isNil(o.DstPortRange) {
		toSerialize["dstPortRange"] = o.DstPortRange
	}
	if !isNil(o.Dscp) {
		toSerialize["dscp"] = o.Dscp
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkSwitchQosRuleRequest struct {
	value *UpdateNetworkSwitchQosRuleRequest
	isSet bool
}

func (v NullableUpdateNetworkSwitchQosRuleRequest) Get() *UpdateNetworkSwitchQosRuleRequest {
	return v.value
}

func (v *NullableUpdateNetworkSwitchQosRuleRequest) Set(val *UpdateNetworkSwitchQosRuleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchQosRuleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchQosRuleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchQosRuleRequest(val *UpdateNetworkSwitchQosRuleRequest) *NullableUpdateNetworkSwitchQosRuleRequest {
	return &NullableUpdateNetworkSwitchQosRuleRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchQosRuleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchQosRuleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


