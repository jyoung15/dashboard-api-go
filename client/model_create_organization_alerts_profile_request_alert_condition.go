/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CreateOrganizationAlertsProfileRequestAlertCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationAlertsProfileRequestAlertCondition{}

// CreateOrganizationAlertsProfileRequestAlertCondition The conditions that determine if the alert triggers
type CreateOrganizationAlertsProfileRequestAlertCondition struct {
	// The total duration in seconds that the threshold should be crossed before alerting
	Duration *int32 `json:"duration,omitempty"`
	// The look back period in seconds for sensing the alert
	Window *int32 `json:"window,omitempty"`
	// The threshold the metric must cross to be valid for alerting. Used only for WAN Utilization alerts.
	BitRateBps *int32 `json:"bit_rate_bps,omitempty"`
	// The threshold the metric must cross to be valid for alerting. Used only for Packet Loss alerts.
	LossRatio *float32 `json:"loss_ratio,omitempty"`
	// The threshold the metric must cross to be valid for alerting. Used only for WAN Latency alerts.
	LatencyMs *int32 `json:"latency_ms,omitempty"`
	// The threshold the metric must cross to be valid for alerting. Used only for VoIP Jitter alerts.
	JitterMs *int32 `json:"jitter_ms,omitempty"`
	// The threshold the metric must drop below to be valid for alerting. Used only for VoIP MOS alerts.
	Mos *float32 `json:"mos,omitempty"`
	// The uplink observed for the alert.  interface must be one of the following: wan1, wan2, wan3, cellular
	Interface *string `json:"interface,omitempty"`
}

// NewCreateOrganizationAlertsProfileRequestAlertCondition instantiates a new CreateOrganizationAlertsProfileRequestAlertCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationAlertsProfileRequestAlertCondition() *CreateOrganizationAlertsProfileRequestAlertCondition {
	this := CreateOrganizationAlertsProfileRequestAlertCondition{}
	return &this
}

// NewCreateOrganizationAlertsProfileRequestAlertConditionWithDefaults instantiates a new CreateOrganizationAlertsProfileRequestAlertCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationAlertsProfileRequestAlertConditionWithDefaults() *CreateOrganizationAlertsProfileRequestAlertCondition {
	this := CreateOrganizationAlertsProfileRequestAlertCondition{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) SetDuration(v int32) {
	o.Duration = &v
}

// GetWindow returns the Window field value if set, zero value otherwise.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetWindow() int32 {
	if o == nil || IsNil(o.Window) {
		var ret int32
		return ret
	}
	return *o.Window
}

// GetWindowOk returns a tuple with the Window field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetWindowOk() (*int32, bool) {
	if o == nil || IsNil(o.Window) {
		return nil, false
	}
	return o.Window, true
}

// HasWindow returns a boolean if a field has been set.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) HasWindow() bool {
	if o != nil && !IsNil(o.Window) {
		return true
	}

	return false
}

// SetWindow gets a reference to the given int32 and assigns it to the Window field.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) SetWindow(v int32) {
	o.Window = &v
}

// GetBitRateBps returns the BitRateBps field value if set, zero value otherwise.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetBitRateBps() int32 {
	if o == nil || IsNil(o.BitRateBps) {
		var ret int32
		return ret
	}
	return *o.BitRateBps
}

// GetBitRateBpsOk returns a tuple with the BitRateBps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetBitRateBpsOk() (*int32, bool) {
	if o == nil || IsNil(o.BitRateBps) {
		return nil, false
	}
	return o.BitRateBps, true
}

// HasBitRateBps returns a boolean if a field has been set.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) HasBitRateBps() bool {
	if o != nil && !IsNil(o.BitRateBps) {
		return true
	}

	return false
}

// SetBitRateBps gets a reference to the given int32 and assigns it to the BitRateBps field.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) SetBitRateBps(v int32) {
	o.BitRateBps = &v
}

// GetLossRatio returns the LossRatio field value if set, zero value otherwise.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetLossRatio() float32 {
	if o == nil || IsNil(o.LossRatio) {
		var ret float32
		return ret
	}
	return *o.LossRatio
}

// GetLossRatioOk returns a tuple with the LossRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetLossRatioOk() (*float32, bool) {
	if o == nil || IsNil(o.LossRatio) {
		return nil, false
	}
	return o.LossRatio, true
}

// HasLossRatio returns a boolean if a field has been set.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) HasLossRatio() bool {
	if o != nil && !IsNil(o.LossRatio) {
		return true
	}

	return false
}

// SetLossRatio gets a reference to the given float32 and assigns it to the LossRatio field.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) SetLossRatio(v float32) {
	o.LossRatio = &v
}

// GetLatencyMs returns the LatencyMs field value if set, zero value otherwise.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetLatencyMs() int32 {
	if o == nil || IsNil(o.LatencyMs) {
		var ret int32
		return ret
	}
	return *o.LatencyMs
}

// GetLatencyMsOk returns a tuple with the LatencyMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetLatencyMsOk() (*int32, bool) {
	if o == nil || IsNil(o.LatencyMs) {
		return nil, false
	}
	return o.LatencyMs, true
}

// HasLatencyMs returns a boolean if a field has been set.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) HasLatencyMs() bool {
	if o != nil && !IsNil(o.LatencyMs) {
		return true
	}

	return false
}

// SetLatencyMs gets a reference to the given int32 and assigns it to the LatencyMs field.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) SetLatencyMs(v int32) {
	o.LatencyMs = &v
}

// GetJitterMs returns the JitterMs field value if set, zero value otherwise.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetJitterMs() int32 {
	if o == nil || IsNil(o.JitterMs) {
		var ret int32
		return ret
	}
	return *o.JitterMs
}

// GetJitterMsOk returns a tuple with the JitterMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetJitterMsOk() (*int32, bool) {
	if o == nil || IsNil(o.JitterMs) {
		return nil, false
	}
	return o.JitterMs, true
}

// HasJitterMs returns a boolean if a field has been set.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) HasJitterMs() bool {
	if o != nil && !IsNil(o.JitterMs) {
		return true
	}

	return false
}

// SetJitterMs gets a reference to the given int32 and assigns it to the JitterMs field.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) SetJitterMs(v int32) {
	o.JitterMs = &v
}

// GetMos returns the Mos field value if set, zero value otherwise.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetMos() float32 {
	if o == nil || IsNil(o.Mos) {
		var ret float32
		return ret
	}
	return *o.Mos
}

// GetMosOk returns a tuple with the Mos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetMosOk() (*float32, bool) {
	if o == nil || IsNil(o.Mos) {
		return nil, false
	}
	return o.Mos, true
}

// HasMos returns a boolean if a field has been set.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) HasMos() bool {
	if o != nil && !IsNil(o.Mos) {
		return true
	}

	return false
}

// SetMos gets a reference to the given float32 and assigns it to the Mos field.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) SetMos(v float32) {
	o.Mos = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetInterface() string {
	if o == nil || IsNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) GetInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.Interface) {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) HasInterface() bool {
	if o != nil && !IsNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *CreateOrganizationAlertsProfileRequestAlertCondition) SetInterface(v string) {
	o.Interface = &v
}

func (o CreateOrganizationAlertsProfileRequestAlertCondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationAlertsProfileRequestAlertCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Window) {
		toSerialize["window"] = o.Window
	}
	if !IsNil(o.BitRateBps) {
		toSerialize["bit_rate_bps"] = o.BitRateBps
	}
	if !IsNil(o.LossRatio) {
		toSerialize["loss_ratio"] = o.LossRatio
	}
	if !IsNil(o.LatencyMs) {
		toSerialize["latency_ms"] = o.LatencyMs
	}
	if !IsNil(o.JitterMs) {
		toSerialize["jitter_ms"] = o.JitterMs
	}
	if !IsNil(o.Mos) {
		toSerialize["mos"] = o.Mos
	}
	if !IsNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	return toSerialize, nil
}

type NullableCreateOrganizationAlertsProfileRequestAlertCondition struct {
	value *CreateOrganizationAlertsProfileRequestAlertCondition
	isSet bool
}

func (v NullableCreateOrganizationAlertsProfileRequestAlertCondition) Get() *CreateOrganizationAlertsProfileRequestAlertCondition {
	return v.value
}

func (v *NullableCreateOrganizationAlertsProfileRequestAlertCondition) Set(val *CreateOrganizationAlertsProfileRequestAlertCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationAlertsProfileRequestAlertCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationAlertsProfileRequestAlertCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationAlertsProfileRequestAlertCondition(val *CreateOrganizationAlertsProfileRequestAlertCondition) *NullableCreateOrganizationAlertsProfileRequestAlertCondition {
	return &NullableCreateOrganizationAlertsProfileRequestAlertCondition{value: val, isSet: true}
}

func (v NullableCreateOrganizationAlertsProfileRequestAlertCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationAlertsProfileRequestAlertCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


