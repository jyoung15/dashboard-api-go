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

// checks if the GetNetworkWirelessRfProfiles200ResponsePerSsidSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessRfProfiles200ResponsePerSsidSettings{}

// GetNetworkWirelessRfProfiles200ResponsePerSsidSettings Per-SSID radio settings by number.
type GetNetworkWirelessRfProfiles200ResponsePerSsidSettings struct {
	Var0 *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings0 `json:"0,omitempty"`
	Var1 *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings1 `json:"1,omitempty"`
	Var2 *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings2 `json:"2,omitempty"`
	Var3 *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings3 `json:"3,omitempty"`
	Var4 *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings4 `json:"4,omitempty"`
	Var5 *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings5 `json:"5,omitempty"`
	Var6 *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings6 `json:"6,omitempty"`
	Var7 *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7 `json:"7,omitempty"`
	Var8 *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings8 `json:"8,omitempty"`
	Var9 *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings9 `json:"9,omitempty"`
	Var10 *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings10 `json:"10,omitempty"`
	Var11 *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings11 `json:"11,omitempty"`
	Var12 *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings12 `json:"12,omitempty"`
	Var13 *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings13 `json:"13,omitempty"`
	Var14 *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings14 `json:"14,omitempty"`
}

// NewGetNetworkWirelessRfProfiles200ResponsePerSsidSettings instantiates a new GetNetworkWirelessRfProfiles200ResponsePerSsidSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessRfProfiles200ResponsePerSsidSettings() *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings {
	this := GetNetworkWirelessRfProfiles200ResponsePerSsidSettings{}
	return &this
}

// NewGetNetworkWirelessRfProfiles200ResponsePerSsidSettingsWithDefaults instantiates a new GetNetworkWirelessRfProfiles200ResponsePerSsidSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessRfProfiles200ResponsePerSsidSettingsWithDefaults() *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings {
	this := GetNetworkWirelessRfProfiles200ResponsePerSsidSettings{}
	return &this
}

// GetVar0 returns the Var0 field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar0() GetNetworkWirelessRfProfiles200ResponsePerSsidSettings0 {
	if o == nil || IsNil(o.Var0) {
		var ret GetNetworkWirelessRfProfiles200ResponsePerSsidSettings0
		return ret
	}
	return *o.Var0
}

// GetVar0Ok returns a tuple with the Var0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar0Ok() (*GetNetworkWirelessRfProfiles200ResponsePerSsidSettings0, bool) {
	if o == nil || IsNil(o.Var0) {
		return nil, false
	}
	return o.Var0, true
}

// HasVar0 returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) HasVar0() bool {
	if o != nil && !IsNil(o.Var0) {
		return true
	}

	return false
}

// SetVar0 gets a reference to the given GetNetworkWirelessRfProfiles200ResponsePerSsidSettings0 and assigns it to the Var0 field.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) SetVar0(v GetNetworkWirelessRfProfiles200ResponsePerSsidSettings0) {
	o.Var0 = &v
}

// GetVar1 returns the Var1 field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar1() GetNetworkWirelessRfProfiles200ResponsePerSsidSettings1 {
	if o == nil || IsNil(o.Var1) {
		var ret GetNetworkWirelessRfProfiles200ResponsePerSsidSettings1
		return ret
	}
	return *o.Var1
}

// GetVar1Ok returns a tuple with the Var1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar1Ok() (*GetNetworkWirelessRfProfiles200ResponsePerSsidSettings1, bool) {
	if o == nil || IsNil(o.Var1) {
		return nil, false
	}
	return o.Var1, true
}

// HasVar1 returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) HasVar1() bool {
	if o != nil && !IsNil(o.Var1) {
		return true
	}

	return false
}

// SetVar1 gets a reference to the given GetNetworkWirelessRfProfiles200ResponsePerSsidSettings1 and assigns it to the Var1 field.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) SetVar1(v GetNetworkWirelessRfProfiles200ResponsePerSsidSettings1) {
	o.Var1 = &v
}

// GetVar2 returns the Var2 field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar2() GetNetworkWirelessRfProfiles200ResponsePerSsidSettings2 {
	if o == nil || IsNil(o.Var2) {
		var ret GetNetworkWirelessRfProfiles200ResponsePerSsidSettings2
		return ret
	}
	return *o.Var2
}

// GetVar2Ok returns a tuple with the Var2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar2Ok() (*GetNetworkWirelessRfProfiles200ResponsePerSsidSettings2, bool) {
	if o == nil || IsNil(o.Var2) {
		return nil, false
	}
	return o.Var2, true
}

// HasVar2 returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) HasVar2() bool {
	if o != nil && !IsNil(o.Var2) {
		return true
	}

	return false
}

// SetVar2 gets a reference to the given GetNetworkWirelessRfProfiles200ResponsePerSsidSettings2 and assigns it to the Var2 field.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) SetVar2(v GetNetworkWirelessRfProfiles200ResponsePerSsidSettings2) {
	o.Var2 = &v
}

// GetVar3 returns the Var3 field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar3() GetNetworkWirelessRfProfiles200ResponsePerSsidSettings3 {
	if o == nil || IsNil(o.Var3) {
		var ret GetNetworkWirelessRfProfiles200ResponsePerSsidSettings3
		return ret
	}
	return *o.Var3
}

// GetVar3Ok returns a tuple with the Var3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar3Ok() (*GetNetworkWirelessRfProfiles200ResponsePerSsidSettings3, bool) {
	if o == nil || IsNil(o.Var3) {
		return nil, false
	}
	return o.Var3, true
}

// HasVar3 returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) HasVar3() bool {
	if o != nil && !IsNil(o.Var3) {
		return true
	}

	return false
}

// SetVar3 gets a reference to the given GetNetworkWirelessRfProfiles200ResponsePerSsidSettings3 and assigns it to the Var3 field.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) SetVar3(v GetNetworkWirelessRfProfiles200ResponsePerSsidSettings3) {
	o.Var3 = &v
}

// GetVar4 returns the Var4 field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar4() GetNetworkWirelessRfProfiles200ResponsePerSsidSettings4 {
	if o == nil || IsNil(o.Var4) {
		var ret GetNetworkWirelessRfProfiles200ResponsePerSsidSettings4
		return ret
	}
	return *o.Var4
}

// GetVar4Ok returns a tuple with the Var4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar4Ok() (*GetNetworkWirelessRfProfiles200ResponsePerSsidSettings4, bool) {
	if o == nil || IsNil(o.Var4) {
		return nil, false
	}
	return o.Var4, true
}

// HasVar4 returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) HasVar4() bool {
	if o != nil && !IsNil(o.Var4) {
		return true
	}

	return false
}

// SetVar4 gets a reference to the given GetNetworkWirelessRfProfiles200ResponsePerSsidSettings4 and assigns it to the Var4 field.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) SetVar4(v GetNetworkWirelessRfProfiles200ResponsePerSsidSettings4) {
	o.Var4 = &v
}

// GetVar5 returns the Var5 field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar5() GetNetworkWirelessRfProfiles200ResponsePerSsidSettings5 {
	if o == nil || IsNil(o.Var5) {
		var ret GetNetworkWirelessRfProfiles200ResponsePerSsidSettings5
		return ret
	}
	return *o.Var5
}

// GetVar5Ok returns a tuple with the Var5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar5Ok() (*GetNetworkWirelessRfProfiles200ResponsePerSsidSettings5, bool) {
	if o == nil || IsNil(o.Var5) {
		return nil, false
	}
	return o.Var5, true
}

// HasVar5 returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) HasVar5() bool {
	if o != nil && !IsNil(o.Var5) {
		return true
	}

	return false
}

// SetVar5 gets a reference to the given GetNetworkWirelessRfProfiles200ResponsePerSsidSettings5 and assigns it to the Var5 field.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) SetVar5(v GetNetworkWirelessRfProfiles200ResponsePerSsidSettings5) {
	o.Var5 = &v
}

// GetVar6 returns the Var6 field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar6() GetNetworkWirelessRfProfiles200ResponsePerSsidSettings6 {
	if o == nil || IsNil(o.Var6) {
		var ret GetNetworkWirelessRfProfiles200ResponsePerSsidSettings6
		return ret
	}
	return *o.Var6
}

// GetVar6Ok returns a tuple with the Var6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar6Ok() (*GetNetworkWirelessRfProfiles200ResponsePerSsidSettings6, bool) {
	if o == nil || IsNil(o.Var6) {
		return nil, false
	}
	return o.Var6, true
}

// HasVar6 returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) HasVar6() bool {
	if o != nil && !IsNil(o.Var6) {
		return true
	}

	return false
}

// SetVar6 gets a reference to the given GetNetworkWirelessRfProfiles200ResponsePerSsidSettings6 and assigns it to the Var6 field.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) SetVar6(v GetNetworkWirelessRfProfiles200ResponsePerSsidSettings6) {
	o.Var6 = &v
}

// GetVar7 returns the Var7 field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar7() GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7 {
	if o == nil || IsNil(o.Var7) {
		var ret GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7
		return ret
	}
	return *o.Var7
}

// GetVar7Ok returns a tuple with the Var7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar7Ok() (*GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7, bool) {
	if o == nil || IsNil(o.Var7) {
		return nil, false
	}
	return o.Var7, true
}

// HasVar7 returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) HasVar7() bool {
	if o != nil && !IsNil(o.Var7) {
		return true
	}

	return false
}

// SetVar7 gets a reference to the given GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7 and assigns it to the Var7 field.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) SetVar7(v GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) {
	o.Var7 = &v
}

// GetVar8 returns the Var8 field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar8() GetNetworkWirelessRfProfiles200ResponsePerSsidSettings8 {
	if o == nil || IsNil(o.Var8) {
		var ret GetNetworkWirelessRfProfiles200ResponsePerSsidSettings8
		return ret
	}
	return *o.Var8
}

// GetVar8Ok returns a tuple with the Var8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar8Ok() (*GetNetworkWirelessRfProfiles200ResponsePerSsidSettings8, bool) {
	if o == nil || IsNil(o.Var8) {
		return nil, false
	}
	return o.Var8, true
}

// HasVar8 returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) HasVar8() bool {
	if o != nil && !IsNil(o.Var8) {
		return true
	}

	return false
}

// SetVar8 gets a reference to the given GetNetworkWirelessRfProfiles200ResponsePerSsidSettings8 and assigns it to the Var8 field.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) SetVar8(v GetNetworkWirelessRfProfiles200ResponsePerSsidSettings8) {
	o.Var8 = &v
}

// GetVar9 returns the Var9 field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar9() GetNetworkWirelessRfProfiles200ResponsePerSsidSettings9 {
	if o == nil || IsNil(o.Var9) {
		var ret GetNetworkWirelessRfProfiles200ResponsePerSsidSettings9
		return ret
	}
	return *o.Var9
}

// GetVar9Ok returns a tuple with the Var9 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar9Ok() (*GetNetworkWirelessRfProfiles200ResponsePerSsidSettings9, bool) {
	if o == nil || IsNil(o.Var9) {
		return nil, false
	}
	return o.Var9, true
}

// HasVar9 returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) HasVar9() bool {
	if o != nil && !IsNil(o.Var9) {
		return true
	}

	return false
}

// SetVar9 gets a reference to the given GetNetworkWirelessRfProfiles200ResponsePerSsidSettings9 and assigns it to the Var9 field.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) SetVar9(v GetNetworkWirelessRfProfiles200ResponsePerSsidSettings9) {
	o.Var9 = &v
}

// GetVar10 returns the Var10 field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar10() GetNetworkWirelessRfProfiles200ResponsePerSsidSettings10 {
	if o == nil || IsNil(o.Var10) {
		var ret GetNetworkWirelessRfProfiles200ResponsePerSsidSettings10
		return ret
	}
	return *o.Var10
}

// GetVar10Ok returns a tuple with the Var10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar10Ok() (*GetNetworkWirelessRfProfiles200ResponsePerSsidSettings10, bool) {
	if o == nil || IsNil(o.Var10) {
		return nil, false
	}
	return o.Var10, true
}

// HasVar10 returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) HasVar10() bool {
	if o != nil && !IsNil(o.Var10) {
		return true
	}

	return false
}

// SetVar10 gets a reference to the given GetNetworkWirelessRfProfiles200ResponsePerSsidSettings10 and assigns it to the Var10 field.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) SetVar10(v GetNetworkWirelessRfProfiles200ResponsePerSsidSettings10) {
	o.Var10 = &v
}

// GetVar11 returns the Var11 field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar11() GetNetworkWirelessRfProfiles200ResponsePerSsidSettings11 {
	if o == nil || IsNil(o.Var11) {
		var ret GetNetworkWirelessRfProfiles200ResponsePerSsidSettings11
		return ret
	}
	return *o.Var11
}

// GetVar11Ok returns a tuple with the Var11 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar11Ok() (*GetNetworkWirelessRfProfiles200ResponsePerSsidSettings11, bool) {
	if o == nil || IsNil(o.Var11) {
		return nil, false
	}
	return o.Var11, true
}

// HasVar11 returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) HasVar11() bool {
	if o != nil && !IsNil(o.Var11) {
		return true
	}

	return false
}

// SetVar11 gets a reference to the given GetNetworkWirelessRfProfiles200ResponsePerSsidSettings11 and assigns it to the Var11 field.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) SetVar11(v GetNetworkWirelessRfProfiles200ResponsePerSsidSettings11) {
	o.Var11 = &v
}

// GetVar12 returns the Var12 field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar12() GetNetworkWirelessRfProfiles200ResponsePerSsidSettings12 {
	if o == nil || IsNil(o.Var12) {
		var ret GetNetworkWirelessRfProfiles200ResponsePerSsidSettings12
		return ret
	}
	return *o.Var12
}

// GetVar12Ok returns a tuple with the Var12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar12Ok() (*GetNetworkWirelessRfProfiles200ResponsePerSsidSettings12, bool) {
	if o == nil || IsNil(o.Var12) {
		return nil, false
	}
	return o.Var12, true
}

// HasVar12 returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) HasVar12() bool {
	if o != nil && !IsNil(o.Var12) {
		return true
	}

	return false
}

// SetVar12 gets a reference to the given GetNetworkWirelessRfProfiles200ResponsePerSsidSettings12 and assigns it to the Var12 field.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) SetVar12(v GetNetworkWirelessRfProfiles200ResponsePerSsidSettings12) {
	o.Var12 = &v
}

// GetVar13 returns the Var13 field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar13() GetNetworkWirelessRfProfiles200ResponsePerSsidSettings13 {
	if o == nil || IsNil(o.Var13) {
		var ret GetNetworkWirelessRfProfiles200ResponsePerSsidSettings13
		return ret
	}
	return *o.Var13
}

// GetVar13Ok returns a tuple with the Var13 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar13Ok() (*GetNetworkWirelessRfProfiles200ResponsePerSsidSettings13, bool) {
	if o == nil || IsNil(o.Var13) {
		return nil, false
	}
	return o.Var13, true
}

// HasVar13 returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) HasVar13() bool {
	if o != nil && !IsNil(o.Var13) {
		return true
	}

	return false
}

// SetVar13 gets a reference to the given GetNetworkWirelessRfProfiles200ResponsePerSsidSettings13 and assigns it to the Var13 field.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) SetVar13(v GetNetworkWirelessRfProfiles200ResponsePerSsidSettings13) {
	o.Var13 = &v
}

// GetVar14 returns the Var14 field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar14() GetNetworkWirelessRfProfiles200ResponsePerSsidSettings14 {
	if o == nil || IsNil(o.Var14) {
		var ret GetNetworkWirelessRfProfiles200ResponsePerSsidSettings14
		return ret
	}
	return *o.Var14
}

// GetVar14Ok returns a tuple with the Var14 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) GetVar14Ok() (*GetNetworkWirelessRfProfiles200ResponsePerSsidSettings14, bool) {
	if o == nil || IsNil(o.Var14) {
		return nil, false
	}
	return o.Var14, true
}

// HasVar14 returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) HasVar14() bool {
	if o != nil && !IsNil(o.Var14) {
		return true
	}

	return false
}

// SetVar14 gets a reference to the given GetNetworkWirelessRfProfiles200ResponsePerSsidSettings14 and assigns it to the Var14 field.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) SetVar14(v GetNetworkWirelessRfProfiles200ResponsePerSsidSettings14) {
	o.Var14 = &v
}

func (o GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) ToMap() (map[string]interface{}, error) {
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

type NullableGetNetworkWirelessRfProfiles200ResponsePerSsidSettings struct {
	value *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings
	isSet bool
}

func (v NullableGetNetworkWirelessRfProfiles200ResponsePerSsidSettings) Get() *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings {
	return v.value
}

func (v *NullableGetNetworkWirelessRfProfiles200ResponsePerSsidSettings) Set(val *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessRfProfiles200ResponsePerSsidSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessRfProfiles200ResponsePerSsidSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessRfProfiles200ResponsePerSsidSettings(val *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings) *NullableGetNetworkWirelessRfProfiles200ResponsePerSsidSettings {
	return &NullableGetNetworkWirelessRfProfiles200ResponsePerSsidSettings{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessRfProfiles200ResponsePerSsidSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessRfProfiles200ResponsePerSsidSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


