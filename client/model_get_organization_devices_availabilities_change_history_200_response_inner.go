/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner{}

// GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner struct for GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner
type GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner struct {
	// Timestamp, in iso8601 format, at which the event happened
	Ts *time.Time `json:"ts,omitempty"`
	Device *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDevice `json:"device,omitempty"`
	Details *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails `json:"details,omitempty"`
	Network *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerNetwork `json:"network,omitempty"`
}

// NewGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner instantiates a new GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner() *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner {
	this := GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner{}
	return &this
}

// NewGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerWithDefaults instantiates a new GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerWithDefaults() *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner {
	this := GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) GetTs() time.Time {
	if o == nil || IsNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) GetTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) SetTs(v time.Time) {
	o.Ts = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) GetDevice() GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDevice {
	if o == nil || IsNil(o.Device) {
		var ret GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) GetDeviceOk() (*GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDevice, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDevice and assigns it to the Device field.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) SetDevice(v GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDevice) {
	o.Device = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) GetDetails() GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails {
	if o == nil || IsNil(o.Details) {
		var ret GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) GetDetailsOk() (*GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails and assigns it to the Details field.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) SetDetails(v GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerDetails) {
	o.Details = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) GetNetwork() GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerNetwork {
	if o == nil || IsNil(o.Network) {
		var ret GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) GetNetworkOk() (*GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerNetwork and assigns it to the Network field.
func (o *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) SetNetwork(v GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInnerNetwork) {
	o.Network = &v
}

func (o GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	return toSerialize, nil
}

type NullableGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner struct {
	value *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) Get() *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) Set(val *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner(val *GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) *NullableGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner {
	return &NullableGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

