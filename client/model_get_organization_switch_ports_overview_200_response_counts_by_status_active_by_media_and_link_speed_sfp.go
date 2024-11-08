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

// checks if the GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp{}

// GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp The count data for SFP ports, indexed by speed in Mb
type GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp struct {
	// The total number of active SFP ports
	Total *int32 `json:"total,omitempty"`
	// The number of active 100 Mbps SFP ports
	Var100 *int32 `json:"100,omitempty"`
	// The number of active 1 Gbps SFP ports
	Var1000 *int32 `json:"1000,omitempty"`
	// The number of active 10 Gbps SFP ports
	Var10000 *int32 `json:"10000,omitempty"`
	// The number of active 20 Gbps SFP ports
	Var20000 *int32 `json:"20000,omitempty"`
	// The number of active 25 Gbps SFP ports
	Var25000 *int32 `json:"25000,omitempty"`
	// The number of active 40 Gbps SFP ports
	Var40000 *int32 `json:"40000,omitempty"`
	// The number of active 50 Gbps SFP ports
	Var50000 *int32 `json:"50000,omitempty"`
	// The number of active 100 Gbps SFP ports
	Var100000 *int32 `json:"100000,omitempty"`
}

// NewGetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp instantiates a new GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp() *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp {
	this := GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp{}
	return &this
}

// NewGetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfpWithDefaults instantiates a new GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfpWithDefaults() *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp {
	this := GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) SetTotal(v int32) {
	o.Total = &v
}

// GetVar100 returns the Var100 field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) GetVar100() int32 {
	if o == nil || IsNil(o.Var100) {
		var ret int32
		return ret
	}
	return *o.Var100
}

// GetVar100Ok returns a tuple with the Var100 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) GetVar100Ok() (*int32, bool) {
	if o == nil || IsNil(o.Var100) {
		return nil, false
	}
	return o.Var100, true
}

// HasVar100 returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) HasVar100() bool {
	if o != nil && !IsNil(o.Var100) {
		return true
	}

	return false
}

// SetVar100 gets a reference to the given int32 and assigns it to the Var100 field.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) SetVar100(v int32) {
	o.Var100 = &v
}

// GetVar1000 returns the Var1000 field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) GetVar1000() int32 {
	if o == nil || IsNil(o.Var1000) {
		var ret int32
		return ret
	}
	return *o.Var1000
}

// GetVar1000Ok returns a tuple with the Var1000 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) GetVar1000Ok() (*int32, bool) {
	if o == nil || IsNil(o.Var1000) {
		return nil, false
	}
	return o.Var1000, true
}

// HasVar1000 returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) HasVar1000() bool {
	if o != nil && !IsNil(o.Var1000) {
		return true
	}

	return false
}

// SetVar1000 gets a reference to the given int32 and assigns it to the Var1000 field.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) SetVar1000(v int32) {
	o.Var1000 = &v
}

// GetVar10000 returns the Var10000 field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) GetVar10000() int32 {
	if o == nil || IsNil(o.Var10000) {
		var ret int32
		return ret
	}
	return *o.Var10000
}

// GetVar10000Ok returns a tuple with the Var10000 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) GetVar10000Ok() (*int32, bool) {
	if o == nil || IsNil(o.Var10000) {
		return nil, false
	}
	return o.Var10000, true
}

// HasVar10000 returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) HasVar10000() bool {
	if o != nil && !IsNil(o.Var10000) {
		return true
	}

	return false
}

// SetVar10000 gets a reference to the given int32 and assigns it to the Var10000 field.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) SetVar10000(v int32) {
	o.Var10000 = &v
}

// GetVar20000 returns the Var20000 field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) GetVar20000() int32 {
	if o == nil || IsNil(o.Var20000) {
		var ret int32
		return ret
	}
	return *o.Var20000
}

// GetVar20000Ok returns a tuple with the Var20000 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) GetVar20000Ok() (*int32, bool) {
	if o == nil || IsNil(o.Var20000) {
		return nil, false
	}
	return o.Var20000, true
}

// HasVar20000 returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) HasVar20000() bool {
	if o != nil && !IsNil(o.Var20000) {
		return true
	}

	return false
}

// SetVar20000 gets a reference to the given int32 and assigns it to the Var20000 field.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) SetVar20000(v int32) {
	o.Var20000 = &v
}

// GetVar25000 returns the Var25000 field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) GetVar25000() int32 {
	if o == nil || IsNil(o.Var25000) {
		var ret int32
		return ret
	}
	return *o.Var25000
}

// GetVar25000Ok returns a tuple with the Var25000 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) GetVar25000Ok() (*int32, bool) {
	if o == nil || IsNil(o.Var25000) {
		return nil, false
	}
	return o.Var25000, true
}

// HasVar25000 returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) HasVar25000() bool {
	if o != nil && !IsNil(o.Var25000) {
		return true
	}

	return false
}

// SetVar25000 gets a reference to the given int32 and assigns it to the Var25000 field.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) SetVar25000(v int32) {
	o.Var25000 = &v
}

// GetVar40000 returns the Var40000 field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) GetVar40000() int32 {
	if o == nil || IsNil(o.Var40000) {
		var ret int32
		return ret
	}
	return *o.Var40000
}

// GetVar40000Ok returns a tuple with the Var40000 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) GetVar40000Ok() (*int32, bool) {
	if o == nil || IsNil(o.Var40000) {
		return nil, false
	}
	return o.Var40000, true
}

// HasVar40000 returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) HasVar40000() bool {
	if o != nil && !IsNil(o.Var40000) {
		return true
	}

	return false
}

// SetVar40000 gets a reference to the given int32 and assigns it to the Var40000 field.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) SetVar40000(v int32) {
	o.Var40000 = &v
}

// GetVar50000 returns the Var50000 field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) GetVar50000() int32 {
	if o == nil || IsNil(o.Var50000) {
		var ret int32
		return ret
	}
	return *o.Var50000
}

// GetVar50000Ok returns a tuple with the Var50000 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) GetVar50000Ok() (*int32, bool) {
	if o == nil || IsNil(o.Var50000) {
		return nil, false
	}
	return o.Var50000, true
}

// HasVar50000 returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) HasVar50000() bool {
	if o != nil && !IsNil(o.Var50000) {
		return true
	}

	return false
}

// SetVar50000 gets a reference to the given int32 and assigns it to the Var50000 field.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) SetVar50000(v int32) {
	o.Var50000 = &v
}

// GetVar100000 returns the Var100000 field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) GetVar100000() int32 {
	if o == nil || IsNil(o.Var100000) {
		var ret int32
		return ret
	}
	return *o.Var100000
}

// GetVar100000Ok returns a tuple with the Var100000 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) GetVar100000Ok() (*int32, bool) {
	if o == nil || IsNil(o.Var100000) {
		return nil, false
	}
	return o.Var100000, true
}

// HasVar100000 returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) HasVar100000() bool {
	if o != nil && !IsNil(o.Var100000) {
		return true
	}

	return false
}

// SetVar100000 gets a reference to the given int32 and assigns it to the Var100000 field.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) SetVar100000(v int32) {
	o.Var100000 = &v
}

func (o GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Var100) {
		toSerialize["100"] = o.Var100
	}
	if !IsNil(o.Var1000) {
		toSerialize["1000"] = o.Var1000
	}
	if !IsNil(o.Var10000) {
		toSerialize["10000"] = o.Var10000
	}
	if !IsNil(o.Var20000) {
		toSerialize["20000"] = o.Var20000
	}
	if !IsNil(o.Var25000) {
		toSerialize["25000"] = o.Var25000
	}
	if !IsNil(o.Var40000) {
		toSerialize["40000"] = o.Var40000
	}
	if !IsNil(o.Var50000) {
		toSerialize["50000"] = o.Var50000
	}
	if !IsNil(o.Var100000) {
		toSerialize["100000"] = o.Var100000
	}
	return toSerialize, nil
}

type NullableGetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp struct {
	value *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp
	isSet bool
}

func (v NullableGetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) Get() *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp {
	return v.value
}

func (v *NullableGetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) Set(val *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp(val *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) *NullableGetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp {
	return &NullableGetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp{value: val, isSet: true}
}

func (v NullableGetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSwitchPortsOverview200ResponseCountsByStatusActiveByMediaAndLinkSpeedSfp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


