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

// checks if the GetNetworkNetflow200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkNetflow200Response{}

// GetNetworkNetflow200Response struct for GetNetworkNetflow200Response
type GetNetworkNetflow200Response struct {
	// Boolean indicating whether NetFlow traffic reporting is enabled (true) or disabled (false).
	ReportingEnabled *bool `json:"reportingEnabled,omitempty"`
	// The IPv4 address of the NetFlow collector.
	CollectorIp *string `json:"collectorIp,omitempty"`
	// The port that the NetFlow collector will be listening on.
	CollectorPort *int32 `json:"collectorPort,omitempty"`
	// Boolean indicating whether Encrypted Traffic Analytics is enabled (true) or disabled (false).
	EtaEnabled *bool `json:"etaEnabled,omitempty"`
	// The port that the Encrypted Traffic Analytics collector will be listening on.
	EtaDstPort *int32 `json:"etaDstPort,omitempty"`
}

// NewGetNetworkNetflow200Response instantiates a new GetNetworkNetflow200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkNetflow200Response() *GetNetworkNetflow200Response {
	this := GetNetworkNetflow200Response{}
	return &this
}

// NewGetNetworkNetflow200ResponseWithDefaults instantiates a new GetNetworkNetflow200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkNetflow200ResponseWithDefaults() *GetNetworkNetflow200Response {
	this := GetNetworkNetflow200Response{}
	return &this
}

// GetReportingEnabled returns the ReportingEnabled field value if set, zero value otherwise.
func (o *GetNetworkNetflow200Response) GetReportingEnabled() bool {
	if o == nil || IsNil(o.ReportingEnabled) {
		var ret bool
		return ret
	}
	return *o.ReportingEnabled
}

// GetReportingEnabledOk returns a tuple with the ReportingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkNetflow200Response) GetReportingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ReportingEnabled) {
		return nil, false
	}
	return o.ReportingEnabled, true
}

// HasReportingEnabled returns a boolean if a field has been set.
func (o *GetNetworkNetflow200Response) HasReportingEnabled() bool {
	if o != nil && !IsNil(o.ReportingEnabled) {
		return true
	}

	return false
}

// SetReportingEnabled gets a reference to the given bool and assigns it to the ReportingEnabled field.
func (o *GetNetworkNetflow200Response) SetReportingEnabled(v bool) {
	o.ReportingEnabled = &v
}

// GetCollectorIp returns the CollectorIp field value if set, zero value otherwise.
func (o *GetNetworkNetflow200Response) GetCollectorIp() string {
	if o == nil || IsNil(o.CollectorIp) {
		var ret string
		return ret
	}
	return *o.CollectorIp
}

// GetCollectorIpOk returns a tuple with the CollectorIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkNetflow200Response) GetCollectorIpOk() (*string, bool) {
	if o == nil || IsNil(o.CollectorIp) {
		return nil, false
	}
	return o.CollectorIp, true
}

// HasCollectorIp returns a boolean if a field has been set.
func (o *GetNetworkNetflow200Response) HasCollectorIp() bool {
	if o != nil && !IsNil(o.CollectorIp) {
		return true
	}

	return false
}

// SetCollectorIp gets a reference to the given string and assigns it to the CollectorIp field.
func (o *GetNetworkNetflow200Response) SetCollectorIp(v string) {
	o.CollectorIp = &v
}

// GetCollectorPort returns the CollectorPort field value if set, zero value otherwise.
func (o *GetNetworkNetflow200Response) GetCollectorPort() int32 {
	if o == nil || IsNil(o.CollectorPort) {
		var ret int32
		return ret
	}
	return *o.CollectorPort
}

// GetCollectorPortOk returns a tuple with the CollectorPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkNetflow200Response) GetCollectorPortOk() (*int32, bool) {
	if o == nil || IsNil(o.CollectorPort) {
		return nil, false
	}
	return o.CollectorPort, true
}

// HasCollectorPort returns a boolean if a field has been set.
func (o *GetNetworkNetflow200Response) HasCollectorPort() bool {
	if o != nil && !IsNil(o.CollectorPort) {
		return true
	}

	return false
}

// SetCollectorPort gets a reference to the given int32 and assigns it to the CollectorPort field.
func (o *GetNetworkNetflow200Response) SetCollectorPort(v int32) {
	o.CollectorPort = &v
}

// GetEtaEnabled returns the EtaEnabled field value if set, zero value otherwise.
func (o *GetNetworkNetflow200Response) GetEtaEnabled() bool {
	if o == nil || IsNil(o.EtaEnabled) {
		var ret bool
		return ret
	}
	return *o.EtaEnabled
}

// GetEtaEnabledOk returns a tuple with the EtaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkNetflow200Response) GetEtaEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EtaEnabled) {
		return nil, false
	}
	return o.EtaEnabled, true
}

// HasEtaEnabled returns a boolean if a field has been set.
func (o *GetNetworkNetflow200Response) HasEtaEnabled() bool {
	if o != nil && !IsNil(o.EtaEnabled) {
		return true
	}

	return false
}

// SetEtaEnabled gets a reference to the given bool and assigns it to the EtaEnabled field.
func (o *GetNetworkNetflow200Response) SetEtaEnabled(v bool) {
	o.EtaEnabled = &v
}

// GetEtaDstPort returns the EtaDstPort field value if set, zero value otherwise.
func (o *GetNetworkNetflow200Response) GetEtaDstPort() int32 {
	if o == nil || IsNil(o.EtaDstPort) {
		var ret int32
		return ret
	}
	return *o.EtaDstPort
}

// GetEtaDstPortOk returns a tuple with the EtaDstPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkNetflow200Response) GetEtaDstPortOk() (*int32, bool) {
	if o == nil || IsNil(o.EtaDstPort) {
		return nil, false
	}
	return o.EtaDstPort, true
}

// HasEtaDstPort returns a boolean if a field has been set.
func (o *GetNetworkNetflow200Response) HasEtaDstPort() bool {
	if o != nil && !IsNil(o.EtaDstPort) {
		return true
	}

	return false
}

// SetEtaDstPort gets a reference to the given int32 and assigns it to the EtaDstPort field.
func (o *GetNetworkNetflow200Response) SetEtaDstPort(v int32) {
	o.EtaDstPort = &v
}

func (o GetNetworkNetflow200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkNetflow200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReportingEnabled) {
		toSerialize["reportingEnabled"] = o.ReportingEnabled
	}
	if !IsNil(o.CollectorIp) {
		toSerialize["collectorIp"] = o.CollectorIp
	}
	if !IsNil(o.CollectorPort) {
		toSerialize["collectorPort"] = o.CollectorPort
	}
	if !IsNil(o.EtaEnabled) {
		toSerialize["etaEnabled"] = o.EtaEnabled
	}
	if !IsNil(o.EtaDstPort) {
		toSerialize["etaDstPort"] = o.EtaDstPort
	}
	return toSerialize, nil
}

type NullableGetNetworkNetflow200Response struct {
	value *GetNetworkNetflow200Response
	isSet bool
}

func (v NullableGetNetworkNetflow200Response) Get() *GetNetworkNetflow200Response {
	return v.value
}

func (v *NullableGetNetworkNetflow200Response) Set(val *GetNetworkNetflow200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkNetflow200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkNetflow200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkNetflow200Response(val *GetNetworkNetflow200Response) *NullableGetNetworkNetflow200Response {
	return &NullableGetNetworkNetflow200Response{value: val, isSet: true}
}

func (v NullableGetNetworkNetflow200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkNetflow200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


