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

// checks if the GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner{}

// GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner struct for GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner
type GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner struct {
	// The start time of the channel utilization interval.
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time of the channel utilization interval.
	EndTs *time.Time `json:"endTs,omitempty"`
	Network *GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork `json:"network,omitempty"`
	// Channel utilization broken down by band.
	ByBand []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner `json:"byBand,omitempty"`
}

// NewGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner instantiates a new GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner() *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner {
	this := GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner{}
	return &this
}

// NewGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInnerWithDefaults instantiates a new GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInnerWithDefaults() *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner {
	this := GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) GetStartTs() time.Time {
	if o == nil || IsNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) GetStartTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTs) {
		return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) HasStartTs() bool {
	if o != nil && !IsNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) GetEndTs() time.Time {
	if o == nil || IsNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) GetEndTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTs) {
		return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) HasEndTs() bool {
	if o != nil && !IsNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) GetNetwork() GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork {
	if o == nil || IsNil(o.Network) {
		var ret GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) GetNetworkOk() (*GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork and assigns it to the Network field.
func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) SetNetwork(v GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerNetwork) {
	o.Network = &v
}

// GetByBand returns the ByBand field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) GetByBand() []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner {
	if o == nil || IsNil(o.ByBand) {
		var ret []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner
		return ret
	}
	return o.ByBand
}

// GetByBandOk returns a tuple with the ByBand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) GetByBandOk() ([]GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner, bool) {
	if o == nil || IsNil(o.ByBand) {
		return nil, false
	}
	return o.ByBand, true
}

// HasByBand returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) HasByBand() bool {
	if o != nil && !IsNil(o.ByBand) {
		return true
	}

	return false
}

// SetByBand gets a reference to the given []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner and assigns it to the ByBand field.
func (o *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) SetByBand(v []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInnerByBandInner) {
	o.ByBand = v
}

func (o GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !IsNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.ByBand) {
		toSerialize["byBand"] = o.ByBand
	}
	return toSerialize, nil
}

type NullableGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner struct {
	value *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) Get() *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) Set(val *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner(val *GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) *NullableGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner {
	return &NullableGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


