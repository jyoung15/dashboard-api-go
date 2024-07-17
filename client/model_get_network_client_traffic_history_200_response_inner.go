/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetNetworkClientTrafficHistory200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkClientTrafficHistory200ResponseInner{}

// GetNetworkClientTrafficHistory200ResponseInner struct for GetNetworkClientTrafficHistory200ResponseInner
type GetNetworkClientTrafficHistory200ResponseInner struct {
	// The timestamp of when the client was connected to an application
	Ts *time.Time `json:"ts,omitempty"`
	// The name of the application the client is connected to
	Application *string `json:"application,omitempty"`
	// The IP or web address the client is connected to
	Destination *string `json:"destination,omitempty"`
	// The client protocol
	Protocol *string `json:"protocol,omitempty"`
	// The port the client is connected to
	Port *int32 `json:"port,omitempty"`
	// Usage received by the client
	Recv *float32 `json:"recv,omitempty"`
	// Usage sent by the client
	Sent *float32 `json:"sent,omitempty"`
	// The number of flows the client has
	NumFlows *int32 `json:"numFlows,omitempty"`
	// The amount of seconds the client was active
	ActiveSeconds *int32 `json:"activeSeconds,omitempty"`
}

// NewGetNetworkClientTrafficHistory200ResponseInner instantiates a new GetNetworkClientTrafficHistory200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkClientTrafficHistory200ResponseInner() *GetNetworkClientTrafficHistory200ResponseInner {
	this := GetNetworkClientTrafficHistory200ResponseInner{}
	return &this
}

// NewGetNetworkClientTrafficHistory200ResponseInnerWithDefaults instantiates a new GetNetworkClientTrafficHistory200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkClientTrafficHistory200ResponseInnerWithDefaults() *GetNetworkClientTrafficHistory200ResponseInner {
	this := GetNetworkClientTrafficHistory200ResponseInner{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetNetworkClientTrafficHistory200ResponseInner) GetTs() time.Time {
	if o == nil || IsNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClientTrafficHistory200ResponseInner) GetTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetNetworkClientTrafficHistory200ResponseInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *GetNetworkClientTrafficHistory200ResponseInner) SetTs(v time.Time) {
	o.Ts = &v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *GetNetworkClientTrafficHistory200ResponseInner) GetApplication() string {
	if o == nil || IsNil(o.Application) {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClientTrafficHistory200ResponseInner) GetApplicationOk() (*string, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *GetNetworkClientTrafficHistory200ResponseInner) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *GetNetworkClientTrafficHistory200ResponseInner) SetApplication(v string) {
	o.Application = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *GetNetworkClientTrafficHistory200ResponseInner) GetDestination() string {
	if o == nil || IsNil(o.Destination) {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClientTrafficHistory200ResponseInner) GetDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *GetNetworkClientTrafficHistory200ResponseInner) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *GetNetworkClientTrafficHistory200ResponseInner) SetDestination(v string) {
	o.Destination = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *GetNetworkClientTrafficHistory200ResponseInner) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClientTrafficHistory200ResponseInner) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *GetNetworkClientTrafficHistory200ResponseInner) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *GetNetworkClientTrafficHistory200ResponseInner) SetProtocol(v string) {
	o.Protocol = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *GetNetworkClientTrafficHistory200ResponseInner) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClientTrafficHistory200ResponseInner) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *GetNetworkClientTrafficHistory200ResponseInner) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *GetNetworkClientTrafficHistory200ResponseInner) SetPort(v int32) {
	o.Port = &v
}

// GetRecv returns the Recv field value if set, zero value otherwise.
func (o *GetNetworkClientTrafficHistory200ResponseInner) GetRecv() float32 {
	if o == nil || IsNil(o.Recv) {
		var ret float32
		return ret
	}
	return *o.Recv
}

// GetRecvOk returns a tuple with the Recv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClientTrafficHistory200ResponseInner) GetRecvOk() (*float32, bool) {
	if o == nil || IsNil(o.Recv) {
		return nil, false
	}
	return o.Recv, true
}

// HasRecv returns a boolean if a field has been set.
func (o *GetNetworkClientTrafficHistory200ResponseInner) HasRecv() bool {
	if o != nil && !IsNil(o.Recv) {
		return true
	}

	return false
}

// SetRecv gets a reference to the given float32 and assigns it to the Recv field.
func (o *GetNetworkClientTrafficHistory200ResponseInner) SetRecv(v float32) {
	o.Recv = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *GetNetworkClientTrafficHistory200ResponseInner) GetSent() float32 {
	if o == nil || IsNil(o.Sent) {
		var ret float32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClientTrafficHistory200ResponseInner) GetSentOk() (*float32, bool) {
	if o == nil || IsNil(o.Sent) {
		return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *GetNetworkClientTrafficHistory200ResponseInner) HasSent() bool {
	if o != nil && !IsNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given float32 and assigns it to the Sent field.
func (o *GetNetworkClientTrafficHistory200ResponseInner) SetSent(v float32) {
	o.Sent = &v
}

// GetNumFlows returns the NumFlows field value if set, zero value otherwise.
func (o *GetNetworkClientTrafficHistory200ResponseInner) GetNumFlows() int32 {
	if o == nil || IsNil(o.NumFlows) {
		var ret int32
		return ret
	}
	return *o.NumFlows
}

// GetNumFlowsOk returns a tuple with the NumFlows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClientTrafficHistory200ResponseInner) GetNumFlowsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumFlows) {
		return nil, false
	}
	return o.NumFlows, true
}

// HasNumFlows returns a boolean if a field has been set.
func (o *GetNetworkClientTrafficHistory200ResponseInner) HasNumFlows() bool {
	if o != nil && !IsNil(o.NumFlows) {
		return true
	}

	return false
}

// SetNumFlows gets a reference to the given int32 and assigns it to the NumFlows field.
func (o *GetNetworkClientTrafficHistory200ResponseInner) SetNumFlows(v int32) {
	o.NumFlows = &v
}

// GetActiveSeconds returns the ActiveSeconds field value if set, zero value otherwise.
func (o *GetNetworkClientTrafficHistory200ResponseInner) GetActiveSeconds() int32 {
	if o == nil || IsNil(o.ActiveSeconds) {
		var ret int32
		return ret
	}
	return *o.ActiveSeconds
}

// GetActiveSecondsOk returns a tuple with the ActiveSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClientTrafficHistory200ResponseInner) GetActiveSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.ActiveSeconds) {
		return nil, false
	}
	return o.ActiveSeconds, true
}

// HasActiveSeconds returns a boolean if a field has been set.
func (o *GetNetworkClientTrafficHistory200ResponseInner) HasActiveSeconds() bool {
	if o != nil && !IsNil(o.ActiveSeconds) {
		return true
	}

	return false
}

// SetActiveSeconds gets a reference to the given int32 and assigns it to the ActiveSeconds field.
func (o *GetNetworkClientTrafficHistory200ResponseInner) SetActiveSeconds(v int32) {
	o.ActiveSeconds = &v
}

func (o GetNetworkClientTrafficHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkClientTrafficHistory200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Recv) {
		toSerialize["recv"] = o.Recv
	}
	if !IsNil(o.Sent) {
		toSerialize["sent"] = o.Sent
	}
	if !IsNil(o.NumFlows) {
		toSerialize["numFlows"] = o.NumFlows
	}
	if !IsNil(o.ActiveSeconds) {
		toSerialize["activeSeconds"] = o.ActiveSeconds
	}
	return toSerialize, nil
}

type NullableGetNetworkClientTrafficHistory200ResponseInner struct {
	value *GetNetworkClientTrafficHistory200ResponseInner
	isSet bool
}

func (v NullableGetNetworkClientTrafficHistory200ResponseInner) Get() *GetNetworkClientTrafficHistory200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkClientTrafficHistory200ResponseInner) Set(val *GetNetworkClientTrafficHistory200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkClientTrafficHistory200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkClientTrafficHistory200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkClientTrafficHistory200ResponseInner(val *GetNetworkClientTrafficHistory200ResponseInner) *NullableGetNetworkClientTrafficHistory200ResponseInner {
	return &NullableGetNetworkClientTrafficHistory200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkClientTrafficHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkClientTrafficHistory200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


