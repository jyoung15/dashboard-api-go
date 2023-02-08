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

// GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat Tower Signal Status
type GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat struct {
	// Reference Signal Received Power
	Rsrp *string `json:"rsrp,omitempty"`
	// Reference Signal Received Quality
	Rsrq *string `json:"rsrq,omitempty"`
}

// NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat instantiates a new GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat() *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat {
	this := GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat{}
	return &this
}

// NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStatWithDefaults instantiates a new GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStatWithDefaults() *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat {
	this := GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat{}
	return &this
}

// GetRsrp returns the Rsrp field value if set, zero value otherwise.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat) GetRsrp() string {
	if o == nil || isNil(o.Rsrp) {
		var ret string
		return ret
	}
	return *o.Rsrp
}

// GetRsrpOk returns a tuple with the Rsrp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat) GetRsrpOk() (*string, bool) {
	if o == nil || isNil(o.Rsrp) {
    return nil, false
	}
	return o.Rsrp, true
}

// HasRsrp returns a boolean if a field has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat) HasRsrp() bool {
	if o != nil && !isNil(o.Rsrp) {
		return true
	}

	return false
}

// SetRsrp gets a reference to the given string and assigns it to the Rsrp field.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat) SetRsrp(v string) {
	o.Rsrp = &v
}

// GetRsrq returns the Rsrq field value if set, zero value otherwise.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat) GetRsrq() string {
	if o == nil || isNil(o.Rsrq) {
		var ret string
		return ret
	}
	return *o.Rsrq
}

// GetRsrqOk returns a tuple with the Rsrq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat) GetRsrqOk() (*string, bool) {
	if o == nil || isNil(o.Rsrq) {
    return nil, false
	}
	return o.Rsrq, true
}

// HasRsrq returns a boolean if a field has been set.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat) HasRsrq() bool {
	if o != nil && !isNil(o.Rsrq) {
		return true
	}

	return false
}

// SetRsrq gets a reference to the given string and assigns it to the Rsrq field.
func (o *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat) SetRsrq(v string) {
	o.Rsrq = &v
}

func (o GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rsrp) {
		toSerialize["rsrp"] = o.Rsrp
	}
	if !isNil(o.Rsrq) {
		toSerialize["rsrq"] = o.Rsrq
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat struct {
	value *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat
	isSet bool
}

func (v NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat) Get() *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat {
	return v.value
}

func (v *NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat) Set(val *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat(val *GetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat) *NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat {
	return &NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat{value: val, isSet: true}
}

func (v NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationCellularGatewayUplinkStatuses200ResponseInnerUplinksInnerSignalStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


