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

// checks if the GetNetworkSmDevicePerformanceHistory200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSmDevicePerformanceHistory200ResponseInner{}

// GetNetworkSmDevicePerformanceHistory200ResponseInner struct for GetNetworkSmDevicePerformanceHistory200ResponseInner
type GetNetworkSmDevicePerformanceHistory200ResponseInner struct {
	// The percentage of CPU used as a decimal format.
	CpuPercentUsed *float32 `json:"cpuPercentUsed,omitempty"`
	// Memory that is not yet in use by the system.
	MemFree *int32 `json:"memFree,omitempty"`
	// Memory used for core OS functions on the device.
	MemWired *int32 `json:"memWired,omitempty"`
	// The active RAM on the device.
	MemActive *int32 `json:"memActive,omitempty"`
	// The inactive RAM on the device.
	MemInactive *int32 `json:"memInactive,omitempty"`
	// Network bandwith transmitted.
	NetworkSent *int32 `json:"networkSent,omitempty"`
	// Network bandwith received.
	NetworkReceived *int32 `json:"networkReceived,omitempty"`
	// The amount of space being used on the startup disk to swap unused files to and from RAM.
	SwapUsed *int32 `json:"swapUsed,omitempty"`
	DiskUsage *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage `json:"diskUsage,omitempty"`
	// The time at which the performance was measured.
	Ts *string `json:"ts,omitempty"`
}

// NewGetNetworkSmDevicePerformanceHistory200ResponseInner instantiates a new GetNetworkSmDevicePerformanceHistory200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSmDevicePerformanceHistory200ResponseInner() *GetNetworkSmDevicePerformanceHistory200ResponseInner {
	this := GetNetworkSmDevicePerformanceHistory200ResponseInner{}
	return &this
}

// NewGetNetworkSmDevicePerformanceHistory200ResponseInnerWithDefaults instantiates a new GetNetworkSmDevicePerformanceHistory200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSmDevicePerformanceHistory200ResponseInnerWithDefaults() *GetNetworkSmDevicePerformanceHistory200ResponseInner {
	this := GetNetworkSmDevicePerformanceHistory200ResponseInner{}
	return &this
}

// GetCpuPercentUsed returns the CpuPercentUsed field value if set, zero value otherwise.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) GetCpuPercentUsed() float32 {
	if o == nil || IsNil(o.CpuPercentUsed) {
		var ret float32
		return ret
	}
	return *o.CpuPercentUsed
}

// GetCpuPercentUsedOk returns a tuple with the CpuPercentUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) GetCpuPercentUsedOk() (*float32, bool) {
	if o == nil || IsNil(o.CpuPercentUsed) {
		return nil, false
	}
	return o.CpuPercentUsed, true
}

// HasCpuPercentUsed returns a boolean if a field has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) HasCpuPercentUsed() bool {
	if o != nil && !IsNil(o.CpuPercentUsed) {
		return true
	}

	return false
}

// SetCpuPercentUsed gets a reference to the given float32 and assigns it to the CpuPercentUsed field.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) SetCpuPercentUsed(v float32) {
	o.CpuPercentUsed = &v
}

// GetMemFree returns the MemFree field value if set, zero value otherwise.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) GetMemFree() int32 {
	if o == nil || IsNil(o.MemFree) {
		var ret int32
		return ret
	}
	return *o.MemFree
}

// GetMemFreeOk returns a tuple with the MemFree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) GetMemFreeOk() (*int32, bool) {
	if o == nil || IsNil(o.MemFree) {
		return nil, false
	}
	return o.MemFree, true
}

// HasMemFree returns a boolean if a field has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) HasMemFree() bool {
	if o != nil && !IsNil(o.MemFree) {
		return true
	}

	return false
}

// SetMemFree gets a reference to the given int32 and assigns it to the MemFree field.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) SetMemFree(v int32) {
	o.MemFree = &v
}

// GetMemWired returns the MemWired field value if set, zero value otherwise.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) GetMemWired() int32 {
	if o == nil || IsNil(o.MemWired) {
		var ret int32
		return ret
	}
	return *o.MemWired
}

// GetMemWiredOk returns a tuple with the MemWired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) GetMemWiredOk() (*int32, bool) {
	if o == nil || IsNil(o.MemWired) {
		return nil, false
	}
	return o.MemWired, true
}

// HasMemWired returns a boolean if a field has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) HasMemWired() bool {
	if o != nil && !IsNil(o.MemWired) {
		return true
	}

	return false
}

// SetMemWired gets a reference to the given int32 and assigns it to the MemWired field.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) SetMemWired(v int32) {
	o.MemWired = &v
}

// GetMemActive returns the MemActive field value if set, zero value otherwise.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) GetMemActive() int32 {
	if o == nil || IsNil(o.MemActive) {
		var ret int32
		return ret
	}
	return *o.MemActive
}

// GetMemActiveOk returns a tuple with the MemActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) GetMemActiveOk() (*int32, bool) {
	if o == nil || IsNil(o.MemActive) {
		return nil, false
	}
	return o.MemActive, true
}

// HasMemActive returns a boolean if a field has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) HasMemActive() bool {
	if o != nil && !IsNil(o.MemActive) {
		return true
	}

	return false
}

// SetMemActive gets a reference to the given int32 and assigns it to the MemActive field.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) SetMemActive(v int32) {
	o.MemActive = &v
}

// GetMemInactive returns the MemInactive field value if set, zero value otherwise.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) GetMemInactive() int32 {
	if o == nil || IsNil(o.MemInactive) {
		var ret int32
		return ret
	}
	return *o.MemInactive
}

// GetMemInactiveOk returns a tuple with the MemInactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) GetMemInactiveOk() (*int32, bool) {
	if o == nil || IsNil(o.MemInactive) {
		return nil, false
	}
	return o.MemInactive, true
}

// HasMemInactive returns a boolean if a field has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) HasMemInactive() bool {
	if o != nil && !IsNil(o.MemInactive) {
		return true
	}

	return false
}

// SetMemInactive gets a reference to the given int32 and assigns it to the MemInactive field.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) SetMemInactive(v int32) {
	o.MemInactive = &v
}

// GetNetworkSent returns the NetworkSent field value if set, zero value otherwise.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) GetNetworkSent() int32 {
	if o == nil || IsNil(o.NetworkSent) {
		var ret int32
		return ret
	}
	return *o.NetworkSent
}

// GetNetworkSentOk returns a tuple with the NetworkSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) GetNetworkSentOk() (*int32, bool) {
	if o == nil || IsNil(o.NetworkSent) {
		return nil, false
	}
	return o.NetworkSent, true
}

// HasNetworkSent returns a boolean if a field has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) HasNetworkSent() bool {
	if o != nil && !IsNil(o.NetworkSent) {
		return true
	}

	return false
}

// SetNetworkSent gets a reference to the given int32 and assigns it to the NetworkSent field.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) SetNetworkSent(v int32) {
	o.NetworkSent = &v
}

// GetNetworkReceived returns the NetworkReceived field value if set, zero value otherwise.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) GetNetworkReceived() int32 {
	if o == nil || IsNil(o.NetworkReceived) {
		var ret int32
		return ret
	}
	return *o.NetworkReceived
}

// GetNetworkReceivedOk returns a tuple with the NetworkReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) GetNetworkReceivedOk() (*int32, bool) {
	if o == nil || IsNil(o.NetworkReceived) {
		return nil, false
	}
	return o.NetworkReceived, true
}

// HasNetworkReceived returns a boolean if a field has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) HasNetworkReceived() bool {
	if o != nil && !IsNil(o.NetworkReceived) {
		return true
	}

	return false
}

// SetNetworkReceived gets a reference to the given int32 and assigns it to the NetworkReceived field.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) SetNetworkReceived(v int32) {
	o.NetworkReceived = &v
}

// GetSwapUsed returns the SwapUsed field value if set, zero value otherwise.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) GetSwapUsed() int32 {
	if o == nil || IsNil(o.SwapUsed) {
		var ret int32
		return ret
	}
	return *o.SwapUsed
}

// GetSwapUsedOk returns a tuple with the SwapUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) GetSwapUsedOk() (*int32, bool) {
	if o == nil || IsNil(o.SwapUsed) {
		return nil, false
	}
	return o.SwapUsed, true
}

// HasSwapUsed returns a boolean if a field has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) HasSwapUsed() bool {
	if o != nil && !IsNil(o.SwapUsed) {
		return true
	}

	return false
}

// SetSwapUsed gets a reference to the given int32 and assigns it to the SwapUsed field.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) SetSwapUsed(v int32) {
	o.SwapUsed = &v
}

// GetDiskUsage returns the DiskUsage field value if set, zero value otherwise.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) GetDiskUsage() GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage {
	if o == nil || IsNil(o.DiskUsage) {
		var ret GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage
		return ret
	}
	return *o.DiskUsage
}

// GetDiskUsageOk returns a tuple with the DiskUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) GetDiskUsageOk() (*GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage, bool) {
	if o == nil || IsNil(o.DiskUsage) {
		return nil, false
	}
	return o.DiskUsage, true
}

// HasDiskUsage returns a boolean if a field has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) HasDiskUsage() bool {
	if o != nil && !IsNil(o.DiskUsage) {
		return true
	}

	return false
}

// SetDiskUsage gets a reference to the given GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage and assigns it to the DiskUsage field.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) SetDiskUsage(v GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsage) {
	o.DiskUsage = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInner) SetTs(v string) {
	o.Ts = &v
}

func (o GetNetworkSmDevicePerformanceHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSmDevicePerformanceHistory200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CpuPercentUsed) {
		toSerialize["cpuPercentUsed"] = o.CpuPercentUsed
	}
	if !IsNil(o.MemFree) {
		toSerialize["memFree"] = o.MemFree
	}
	if !IsNil(o.MemWired) {
		toSerialize["memWired"] = o.MemWired
	}
	if !IsNil(o.MemActive) {
		toSerialize["memActive"] = o.MemActive
	}
	if !IsNil(o.MemInactive) {
		toSerialize["memInactive"] = o.MemInactive
	}
	if !IsNil(o.NetworkSent) {
		toSerialize["networkSent"] = o.NetworkSent
	}
	if !IsNil(o.NetworkReceived) {
		toSerialize["networkReceived"] = o.NetworkReceived
	}
	if !IsNil(o.SwapUsed) {
		toSerialize["swapUsed"] = o.SwapUsed
	}
	if !IsNil(o.DiskUsage) {
		toSerialize["diskUsage"] = o.DiskUsage
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableGetNetworkSmDevicePerformanceHistory200ResponseInner struct {
	value *GetNetworkSmDevicePerformanceHistory200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSmDevicePerformanceHistory200ResponseInner) Get() *GetNetworkSmDevicePerformanceHistory200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSmDevicePerformanceHistory200ResponseInner) Set(val *GetNetworkSmDevicePerformanceHistory200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSmDevicePerformanceHistory200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSmDevicePerformanceHistory200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSmDevicePerformanceHistory200ResponseInner(val *GetNetworkSmDevicePerformanceHistory200ResponseInner) *NullableGetNetworkSmDevicePerformanceHistory200ResponseInner {
	return &NullableGetNetworkSmDevicePerformanceHistory200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSmDevicePerformanceHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSmDevicePerformanceHistory200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


