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

// checks if the UpdateNetworkClientSplashAuthorizationStatusRequestSsids type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkClientSplashAuthorizationStatusRequestSsids{}

// UpdateNetworkClientSplashAuthorizationStatusRequestSsids The target SSIDs. Each SSID must be enabled and must have Click-through splash enabled. For each SSID where isAuthorized is true, the expiration time will automatically be set according to the SSID's splash frequency. Not all networks support configuring all SSIDs
type UpdateNetworkClientSplashAuthorizationStatusRequestSsids struct {
	Var0 *UpdateNetworkClientSplashAuthorizationStatusRequestSsids0 `json:"0,omitempty"`
	Var1 *UpdateNetworkClientSplashAuthorizationStatusRequestSsids1 `json:"1,omitempty"`
	Var2 *UpdateNetworkClientSplashAuthorizationStatusRequestSsids2 `json:"2,omitempty"`
	Var3 *UpdateNetworkClientSplashAuthorizationStatusRequestSsids3 `json:"3,omitempty"`
	Var4 *UpdateNetworkClientSplashAuthorizationStatusRequestSsids4 `json:"4,omitempty"`
	Var5 *UpdateNetworkClientSplashAuthorizationStatusRequestSsids5 `json:"5,omitempty"`
	Var6 *UpdateNetworkClientSplashAuthorizationStatusRequestSsids6 `json:"6,omitempty"`
	Var7 *UpdateNetworkClientSplashAuthorizationStatusRequestSsids7 `json:"7,omitempty"`
	Var8 *UpdateNetworkClientSplashAuthorizationStatusRequestSsids8 `json:"8,omitempty"`
	Var9 *UpdateNetworkClientSplashAuthorizationStatusRequestSsids9 `json:"9,omitempty"`
	Var10 *UpdateNetworkClientSplashAuthorizationStatusRequestSsids10 `json:"10,omitempty"`
	Var11 *UpdateNetworkClientSplashAuthorizationStatusRequestSsids11 `json:"11,omitempty"`
	Var12 *UpdateNetworkClientSplashAuthorizationStatusRequestSsids12 `json:"12,omitempty"`
	Var13 *UpdateNetworkClientSplashAuthorizationStatusRequestSsids13 `json:"13,omitempty"`
	Var14 *UpdateNetworkClientSplashAuthorizationStatusRequestSsids14 `json:"14,omitempty"`
}

// NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids instantiates a new UpdateNetworkClientSplashAuthorizationStatusRequestSsids object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids() *UpdateNetworkClientSplashAuthorizationStatusRequestSsids {
	this := UpdateNetworkClientSplashAuthorizationStatusRequestSsids{}
	return &this
}

// NewUpdateNetworkClientSplashAuthorizationStatusRequestSsidsWithDefaults instantiates a new UpdateNetworkClientSplashAuthorizationStatusRequestSsids object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkClientSplashAuthorizationStatusRequestSsidsWithDefaults() *UpdateNetworkClientSplashAuthorizationStatusRequestSsids {
	this := UpdateNetworkClientSplashAuthorizationStatusRequestSsids{}
	return &this
}

// GetVar0 returns the Var0 field value if set, zero value otherwise.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar0() UpdateNetworkClientSplashAuthorizationStatusRequestSsids0 {
	if o == nil || IsNil(o.Var0) {
		var ret UpdateNetworkClientSplashAuthorizationStatusRequestSsids0
		return ret
	}
	return *o.Var0
}

// GetVar0Ok returns a tuple with the Var0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar0Ok() (*UpdateNetworkClientSplashAuthorizationStatusRequestSsids0, bool) {
	if o == nil || IsNil(o.Var0) {
		return nil, false
	}
	return o.Var0, true
}

// HasVar0 returns a boolean if a field has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) HasVar0() bool {
	if o != nil && !IsNil(o.Var0) {
		return true
	}

	return false
}

// SetVar0 gets a reference to the given UpdateNetworkClientSplashAuthorizationStatusRequestSsids0 and assigns it to the Var0 field.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) SetVar0(v UpdateNetworkClientSplashAuthorizationStatusRequestSsids0) {
	o.Var0 = &v
}

// GetVar1 returns the Var1 field value if set, zero value otherwise.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar1() UpdateNetworkClientSplashAuthorizationStatusRequestSsids1 {
	if o == nil || IsNil(o.Var1) {
		var ret UpdateNetworkClientSplashAuthorizationStatusRequestSsids1
		return ret
	}
	return *o.Var1
}

// GetVar1Ok returns a tuple with the Var1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar1Ok() (*UpdateNetworkClientSplashAuthorizationStatusRequestSsids1, bool) {
	if o == nil || IsNil(o.Var1) {
		return nil, false
	}
	return o.Var1, true
}

// HasVar1 returns a boolean if a field has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) HasVar1() bool {
	if o != nil && !IsNil(o.Var1) {
		return true
	}

	return false
}

// SetVar1 gets a reference to the given UpdateNetworkClientSplashAuthorizationStatusRequestSsids1 and assigns it to the Var1 field.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) SetVar1(v UpdateNetworkClientSplashAuthorizationStatusRequestSsids1) {
	o.Var1 = &v
}

// GetVar2 returns the Var2 field value if set, zero value otherwise.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar2() UpdateNetworkClientSplashAuthorizationStatusRequestSsids2 {
	if o == nil || IsNil(o.Var2) {
		var ret UpdateNetworkClientSplashAuthorizationStatusRequestSsids2
		return ret
	}
	return *o.Var2
}

// GetVar2Ok returns a tuple with the Var2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar2Ok() (*UpdateNetworkClientSplashAuthorizationStatusRequestSsids2, bool) {
	if o == nil || IsNil(o.Var2) {
		return nil, false
	}
	return o.Var2, true
}

// HasVar2 returns a boolean if a field has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) HasVar2() bool {
	if o != nil && !IsNil(o.Var2) {
		return true
	}

	return false
}

// SetVar2 gets a reference to the given UpdateNetworkClientSplashAuthorizationStatusRequestSsids2 and assigns it to the Var2 field.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) SetVar2(v UpdateNetworkClientSplashAuthorizationStatusRequestSsids2) {
	o.Var2 = &v
}

// GetVar3 returns the Var3 field value if set, zero value otherwise.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar3() UpdateNetworkClientSplashAuthorizationStatusRequestSsids3 {
	if o == nil || IsNil(o.Var3) {
		var ret UpdateNetworkClientSplashAuthorizationStatusRequestSsids3
		return ret
	}
	return *o.Var3
}

// GetVar3Ok returns a tuple with the Var3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar3Ok() (*UpdateNetworkClientSplashAuthorizationStatusRequestSsids3, bool) {
	if o == nil || IsNil(o.Var3) {
		return nil, false
	}
	return o.Var3, true
}

// HasVar3 returns a boolean if a field has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) HasVar3() bool {
	if o != nil && !IsNil(o.Var3) {
		return true
	}

	return false
}

// SetVar3 gets a reference to the given UpdateNetworkClientSplashAuthorizationStatusRequestSsids3 and assigns it to the Var3 field.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) SetVar3(v UpdateNetworkClientSplashAuthorizationStatusRequestSsids3) {
	o.Var3 = &v
}

// GetVar4 returns the Var4 field value if set, zero value otherwise.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar4() UpdateNetworkClientSplashAuthorizationStatusRequestSsids4 {
	if o == nil || IsNil(o.Var4) {
		var ret UpdateNetworkClientSplashAuthorizationStatusRequestSsids4
		return ret
	}
	return *o.Var4
}

// GetVar4Ok returns a tuple with the Var4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar4Ok() (*UpdateNetworkClientSplashAuthorizationStatusRequestSsids4, bool) {
	if o == nil || IsNil(o.Var4) {
		return nil, false
	}
	return o.Var4, true
}

// HasVar4 returns a boolean if a field has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) HasVar4() bool {
	if o != nil && !IsNil(o.Var4) {
		return true
	}

	return false
}

// SetVar4 gets a reference to the given UpdateNetworkClientSplashAuthorizationStatusRequestSsids4 and assigns it to the Var4 field.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) SetVar4(v UpdateNetworkClientSplashAuthorizationStatusRequestSsids4) {
	o.Var4 = &v
}

// GetVar5 returns the Var5 field value if set, zero value otherwise.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar5() UpdateNetworkClientSplashAuthorizationStatusRequestSsids5 {
	if o == nil || IsNil(o.Var5) {
		var ret UpdateNetworkClientSplashAuthorizationStatusRequestSsids5
		return ret
	}
	return *o.Var5
}

// GetVar5Ok returns a tuple with the Var5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar5Ok() (*UpdateNetworkClientSplashAuthorizationStatusRequestSsids5, bool) {
	if o == nil || IsNil(o.Var5) {
		return nil, false
	}
	return o.Var5, true
}

// HasVar5 returns a boolean if a field has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) HasVar5() bool {
	if o != nil && !IsNil(o.Var5) {
		return true
	}

	return false
}

// SetVar5 gets a reference to the given UpdateNetworkClientSplashAuthorizationStatusRequestSsids5 and assigns it to the Var5 field.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) SetVar5(v UpdateNetworkClientSplashAuthorizationStatusRequestSsids5) {
	o.Var5 = &v
}

// GetVar6 returns the Var6 field value if set, zero value otherwise.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar6() UpdateNetworkClientSplashAuthorizationStatusRequestSsids6 {
	if o == nil || IsNil(o.Var6) {
		var ret UpdateNetworkClientSplashAuthorizationStatusRequestSsids6
		return ret
	}
	return *o.Var6
}

// GetVar6Ok returns a tuple with the Var6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar6Ok() (*UpdateNetworkClientSplashAuthorizationStatusRequestSsids6, bool) {
	if o == nil || IsNil(o.Var6) {
		return nil, false
	}
	return o.Var6, true
}

// HasVar6 returns a boolean if a field has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) HasVar6() bool {
	if o != nil && !IsNil(o.Var6) {
		return true
	}

	return false
}

// SetVar6 gets a reference to the given UpdateNetworkClientSplashAuthorizationStatusRequestSsids6 and assigns it to the Var6 field.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) SetVar6(v UpdateNetworkClientSplashAuthorizationStatusRequestSsids6) {
	o.Var6 = &v
}

// GetVar7 returns the Var7 field value if set, zero value otherwise.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar7() UpdateNetworkClientSplashAuthorizationStatusRequestSsids7 {
	if o == nil || IsNil(o.Var7) {
		var ret UpdateNetworkClientSplashAuthorizationStatusRequestSsids7
		return ret
	}
	return *o.Var7
}

// GetVar7Ok returns a tuple with the Var7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar7Ok() (*UpdateNetworkClientSplashAuthorizationStatusRequestSsids7, bool) {
	if o == nil || IsNil(o.Var7) {
		return nil, false
	}
	return o.Var7, true
}

// HasVar7 returns a boolean if a field has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) HasVar7() bool {
	if o != nil && !IsNil(o.Var7) {
		return true
	}

	return false
}

// SetVar7 gets a reference to the given UpdateNetworkClientSplashAuthorizationStatusRequestSsids7 and assigns it to the Var7 field.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) SetVar7(v UpdateNetworkClientSplashAuthorizationStatusRequestSsids7) {
	o.Var7 = &v
}

// GetVar8 returns the Var8 field value if set, zero value otherwise.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar8() UpdateNetworkClientSplashAuthorizationStatusRequestSsids8 {
	if o == nil || IsNil(o.Var8) {
		var ret UpdateNetworkClientSplashAuthorizationStatusRequestSsids8
		return ret
	}
	return *o.Var8
}

// GetVar8Ok returns a tuple with the Var8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar8Ok() (*UpdateNetworkClientSplashAuthorizationStatusRequestSsids8, bool) {
	if o == nil || IsNil(o.Var8) {
		return nil, false
	}
	return o.Var8, true
}

// HasVar8 returns a boolean if a field has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) HasVar8() bool {
	if o != nil && !IsNil(o.Var8) {
		return true
	}

	return false
}

// SetVar8 gets a reference to the given UpdateNetworkClientSplashAuthorizationStatusRequestSsids8 and assigns it to the Var8 field.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) SetVar8(v UpdateNetworkClientSplashAuthorizationStatusRequestSsids8) {
	o.Var8 = &v
}

// GetVar9 returns the Var9 field value if set, zero value otherwise.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar9() UpdateNetworkClientSplashAuthorizationStatusRequestSsids9 {
	if o == nil || IsNil(o.Var9) {
		var ret UpdateNetworkClientSplashAuthorizationStatusRequestSsids9
		return ret
	}
	return *o.Var9
}

// GetVar9Ok returns a tuple with the Var9 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar9Ok() (*UpdateNetworkClientSplashAuthorizationStatusRequestSsids9, bool) {
	if o == nil || IsNil(o.Var9) {
		return nil, false
	}
	return o.Var9, true
}

// HasVar9 returns a boolean if a field has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) HasVar9() bool {
	if o != nil && !IsNil(o.Var9) {
		return true
	}

	return false
}

// SetVar9 gets a reference to the given UpdateNetworkClientSplashAuthorizationStatusRequestSsids9 and assigns it to the Var9 field.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) SetVar9(v UpdateNetworkClientSplashAuthorizationStatusRequestSsids9) {
	o.Var9 = &v
}

// GetVar10 returns the Var10 field value if set, zero value otherwise.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar10() UpdateNetworkClientSplashAuthorizationStatusRequestSsids10 {
	if o == nil || IsNil(o.Var10) {
		var ret UpdateNetworkClientSplashAuthorizationStatusRequestSsids10
		return ret
	}
	return *o.Var10
}

// GetVar10Ok returns a tuple with the Var10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar10Ok() (*UpdateNetworkClientSplashAuthorizationStatusRequestSsids10, bool) {
	if o == nil || IsNil(o.Var10) {
		return nil, false
	}
	return o.Var10, true
}

// HasVar10 returns a boolean if a field has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) HasVar10() bool {
	if o != nil && !IsNil(o.Var10) {
		return true
	}

	return false
}

// SetVar10 gets a reference to the given UpdateNetworkClientSplashAuthorizationStatusRequestSsids10 and assigns it to the Var10 field.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) SetVar10(v UpdateNetworkClientSplashAuthorizationStatusRequestSsids10) {
	o.Var10 = &v
}

// GetVar11 returns the Var11 field value if set, zero value otherwise.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar11() UpdateNetworkClientSplashAuthorizationStatusRequestSsids11 {
	if o == nil || IsNil(o.Var11) {
		var ret UpdateNetworkClientSplashAuthorizationStatusRequestSsids11
		return ret
	}
	return *o.Var11
}

// GetVar11Ok returns a tuple with the Var11 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar11Ok() (*UpdateNetworkClientSplashAuthorizationStatusRequestSsids11, bool) {
	if o == nil || IsNil(o.Var11) {
		return nil, false
	}
	return o.Var11, true
}

// HasVar11 returns a boolean if a field has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) HasVar11() bool {
	if o != nil && !IsNil(o.Var11) {
		return true
	}

	return false
}

// SetVar11 gets a reference to the given UpdateNetworkClientSplashAuthorizationStatusRequestSsids11 and assigns it to the Var11 field.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) SetVar11(v UpdateNetworkClientSplashAuthorizationStatusRequestSsids11) {
	o.Var11 = &v
}

// GetVar12 returns the Var12 field value if set, zero value otherwise.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar12() UpdateNetworkClientSplashAuthorizationStatusRequestSsids12 {
	if o == nil || IsNil(o.Var12) {
		var ret UpdateNetworkClientSplashAuthorizationStatusRequestSsids12
		return ret
	}
	return *o.Var12
}

// GetVar12Ok returns a tuple with the Var12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar12Ok() (*UpdateNetworkClientSplashAuthorizationStatusRequestSsids12, bool) {
	if o == nil || IsNil(o.Var12) {
		return nil, false
	}
	return o.Var12, true
}

// HasVar12 returns a boolean if a field has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) HasVar12() bool {
	if o != nil && !IsNil(o.Var12) {
		return true
	}

	return false
}

// SetVar12 gets a reference to the given UpdateNetworkClientSplashAuthorizationStatusRequestSsids12 and assigns it to the Var12 field.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) SetVar12(v UpdateNetworkClientSplashAuthorizationStatusRequestSsids12) {
	o.Var12 = &v
}

// GetVar13 returns the Var13 field value if set, zero value otherwise.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar13() UpdateNetworkClientSplashAuthorizationStatusRequestSsids13 {
	if o == nil || IsNil(o.Var13) {
		var ret UpdateNetworkClientSplashAuthorizationStatusRequestSsids13
		return ret
	}
	return *o.Var13
}

// GetVar13Ok returns a tuple with the Var13 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar13Ok() (*UpdateNetworkClientSplashAuthorizationStatusRequestSsids13, bool) {
	if o == nil || IsNil(o.Var13) {
		return nil, false
	}
	return o.Var13, true
}

// HasVar13 returns a boolean if a field has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) HasVar13() bool {
	if o != nil && !IsNil(o.Var13) {
		return true
	}

	return false
}

// SetVar13 gets a reference to the given UpdateNetworkClientSplashAuthorizationStatusRequestSsids13 and assigns it to the Var13 field.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) SetVar13(v UpdateNetworkClientSplashAuthorizationStatusRequestSsids13) {
	o.Var13 = &v
}

// GetVar14 returns the Var14 field value if set, zero value otherwise.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar14() UpdateNetworkClientSplashAuthorizationStatusRequestSsids14 {
	if o == nil || IsNil(o.Var14) {
		var ret UpdateNetworkClientSplashAuthorizationStatusRequestSsids14
		return ret
	}
	return *o.Var14
}

// GetVar14Ok returns a tuple with the Var14 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) GetVar14Ok() (*UpdateNetworkClientSplashAuthorizationStatusRequestSsids14, bool) {
	if o == nil || IsNil(o.Var14) {
		return nil, false
	}
	return o.Var14, true
}

// HasVar14 returns a boolean if a field has been set.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) HasVar14() bool {
	if o != nil && !IsNil(o.Var14) {
		return true
	}

	return false
}

// SetVar14 gets a reference to the given UpdateNetworkClientSplashAuthorizationStatusRequestSsids14 and assigns it to the Var14 field.
func (o *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) SetVar14(v UpdateNetworkClientSplashAuthorizationStatusRequestSsids14) {
	o.Var14 = &v
}

func (o UpdateNetworkClientSplashAuthorizationStatusRequestSsids) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkClientSplashAuthorizationStatusRequestSsids) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Var0) {
		toSerialize["0"] = o.Var0
	}
	if !IsNil(o.Var1) {
		toSerialize["1"] = o.Var1
	}
	if !IsNil(o.Var2) {
		toSerialize["2"] = o.Var2
	}
	if !IsNil(o.Var3) {
		toSerialize["3"] = o.Var3
	}
	if !IsNil(o.Var4) {
		toSerialize["4"] = o.Var4
	}
	if !IsNil(o.Var5) {
		toSerialize["5"] = o.Var5
	}
	if !IsNil(o.Var6) {
		toSerialize["6"] = o.Var6
	}
	if !IsNil(o.Var7) {
		toSerialize["7"] = o.Var7
	}
	if !IsNil(o.Var8) {
		toSerialize["8"] = o.Var8
	}
	if !IsNil(o.Var9) {
		toSerialize["9"] = o.Var9
	}
	if !IsNil(o.Var10) {
		toSerialize["10"] = o.Var10
	}
	if !IsNil(o.Var11) {
		toSerialize["11"] = o.Var11
	}
	if !IsNil(o.Var12) {
		toSerialize["12"] = o.Var12
	}
	if !IsNil(o.Var13) {
		toSerialize["13"] = o.Var13
	}
	if !IsNil(o.Var14) {
		toSerialize["14"] = o.Var14
	}
	return toSerialize, nil
}

type NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids struct {
	value *UpdateNetworkClientSplashAuthorizationStatusRequestSsids
	isSet bool
}

func (v NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids) Get() *UpdateNetworkClientSplashAuthorizationStatusRequestSsids {
	return v.value
}

func (v *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids) Set(val *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids(val *UpdateNetworkClientSplashAuthorizationStatusRequestSsids) *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids {
	return &NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids{value: val, isSet: true}
}

func (v NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkClientSplashAuthorizationStatusRequestSsids) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


