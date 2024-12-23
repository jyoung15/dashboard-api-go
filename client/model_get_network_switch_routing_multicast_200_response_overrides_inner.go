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

// checks if the GetNetworkSwitchRoutingMulticast200ResponseOverridesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchRoutingMulticast200ResponseOverridesInner{}

// GetNetworkSwitchRoutingMulticast200ResponseOverridesInner struct for GetNetworkSwitchRoutingMulticast200ResponseOverridesInner
type GetNetworkSwitchRoutingMulticast200ResponseOverridesInner struct {
	// (optional) List of switch serials for non-template network
	Switches []string `json:"switches,omitempty"`
	// (optional) List of switch stack ids for non-template network
	Stacks []string `json:"stacks,omitempty"`
	// (optional) List of switch templates ids for template network
	SwitchProfiles []string `json:"switchProfiles,omitempty"`
	// IGMP snooping enabled for switches, switch stacks or switch templates
	IgmpSnoopingEnabled *bool `json:"igmpSnoopingEnabled,omitempty"`
	// Flood unknown multicast traffic enabled for switches, switch stacks or switch templates
	FloodUnknownMulticastTrafficEnabled *bool `json:"floodUnknownMulticastTrafficEnabled,omitempty"`
}

// NewGetNetworkSwitchRoutingMulticast200ResponseOverridesInner instantiates a new GetNetworkSwitchRoutingMulticast200ResponseOverridesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchRoutingMulticast200ResponseOverridesInner() *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner {
	this := GetNetworkSwitchRoutingMulticast200ResponseOverridesInner{}
	return &this
}

// NewGetNetworkSwitchRoutingMulticast200ResponseOverridesInnerWithDefaults instantiates a new GetNetworkSwitchRoutingMulticast200ResponseOverridesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchRoutingMulticast200ResponseOverridesInnerWithDefaults() *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner {
	this := GetNetworkSwitchRoutingMulticast200ResponseOverridesInner{}
	return &this
}

// GetSwitches returns the Switches field value if set, zero value otherwise.
func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) GetSwitches() []string {
	if o == nil || IsNil(o.Switches) {
		var ret []string
		return ret
	}
	return o.Switches
}

// GetSwitchesOk returns a tuple with the Switches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) GetSwitchesOk() ([]string, bool) {
	if o == nil || IsNil(o.Switches) {
		return nil, false
	}
	return o.Switches, true
}

// HasSwitches returns a boolean if a field has been set.
func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) HasSwitches() bool {
	if o != nil && !IsNil(o.Switches) {
		return true
	}

	return false
}

// SetSwitches gets a reference to the given []string and assigns it to the Switches field.
func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) SetSwitches(v []string) {
	o.Switches = v
}

// GetStacks returns the Stacks field value if set, zero value otherwise.
func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) GetStacks() []string {
	if o == nil || IsNil(o.Stacks) {
		var ret []string
		return ret
	}
	return o.Stacks
}

// GetStacksOk returns a tuple with the Stacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) GetStacksOk() ([]string, bool) {
	if o == nil || IsNil(o.Stacks) {
		return nil, false
	}
	return o.Stacks, true
}

// HasStacks returns a boolean if a field has been set.
func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) HasStacks() bool {
	if o != nil && !IsNil(o.Stacks) {
		return true
	}

	return false
}

// SetStacks gets a reference to the given []string and assigns it to the Stacks field.
func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) SetStacks(v []string) {
	o.Stacks = v
}

// GetSwitchProfiles returns the SwitchProfiles field value if set, zero value otherwise.
func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) GetSwitchProfiles() []string {
	if o == nil || IsNil(o.SwitchProfiles) {
		var ret []string
		return ret
	}
	return o.SwitchProfiles
}

// GetSwitchProfilesOk returns a tuple with the SwitchProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) GetSwitchProfilesOk() ([]string, bool) {
	if o == nil || IsNil(o.SwitchProfiles) {
		return nil, false
	}
	return o.SwitchProfiles, true
}

// HasSwitchProfiles returns a boolean if a field has been set.
func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) HasSwitchProfiles() bool {
	if o != nil && !IsNil(o.SwitchProfiles) {
		return true
	}

	return false
}

// SetSwitchProfiles gets a reference to the given []string and assigns it to the SwitchProfiles field.
func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) SetSwitchProfiles(v []string) {
	o.SwitchProfiles = v
}

// GetIgmpSnoopingEnabled returns the IgmpSnoopingEnabled field value if set, zero value otherwise.
func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) GetIgmpSnoopingEnabled() bool {
	if o == nil || IsNil(o.IgmpSnoopingEnabled) {
		var ret bool
		return ret
	}
	return *o.IgmpSnoopingEnabled
}

// GetIgmpSnoopingEnabledOk returns a tuple with the IgmpSnoopingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) GetIgmpSnoopingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IgmpSnoopingEnabled) {
		return nil, false
	}
	return o.IgmpSnoopingEnabled, true
}

// HasIgmpSnoopingEnabled returns a boolean if a field has been set.
func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) HasIgmpSnoopingEnabled() bool {
	if o != nil && !IsNil(o.IgmpSnoopingEnabled) {
		return true
	}

	return false
}

// SetIgmpSnoopingEnabled gets a reference to the given bool and assigns it to the IgmpSnoopingEnabled field.
func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) SetIgmpSnoopingEnabled(v bool) {
	o.IgmpSnoopingEnabled = &v
}

// GetFloodUnknownMulticastTrafficEnabled returns the FloodUnknownMulticastTrafficEnabled field value if set, zero value otherwise.
func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) GetFloodUnknownMulticastTrafficEnabled() bool {
	if o == nil || IsNil(o.FloodUnknownMulticastTrafficEnabled) {
		var ret bool
		return ret
	}
	return *o.FloodUnknownMulticastTrafficEnabled
}

// GetFloodUnknownMulticastTrafficEnabledOk returns a tuple with the FloodUnknownMulticastTrafficEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) GetFloodUnknownMulticastTrafficEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FloodUnknownMulticastTrafficEnabled) {
		return nil, false
	}
	return o.FloodUnknownMulticastTrafficEnabled, true
}

// HasFloodUnknownMulticastTrafficEnabled returns a boolean if a field has been set.
func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) HasFloodUnknownMulticastTrafficEnabled() bool {
	if o != nil && !IsNil(o.FloodUnknownMulticastTrafficEnabled) {
		return true
	}

	return false
}

// SetFloodUnknownMulticastTrafficEnabled gets a reference to the given bool and assigns it to the FloodUnknownMulticastTrafficEnabled field.
func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) SetFloodUnknownMulticastTrafficEnabled(v bool) {
	o.FloodUnknownMulticastTrafficEnabled = &v
}

func (o GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Switches) {
		toSerialize["switches"] = o.Switches
	}
	if !IsNil(o.Stacks) {
		toSerialize["stacks"] = o.Stacks
	}
	if !IsNil(o.SwitchProfiles) {
		toSerialize["switchProfiles"] = o.SwitchProfiles
	}
	if !IsNil(o.IgmpSnoopingEnabled) {
		toSerialize["igmpSnoopingEnabled"] = o.IgmpSnoopingEnabled
	}
	if !IsNil(o.FloodUnknownMulticastTrafficEnabled) {
		toSerialize["floodUnknownMulticastTrafficEnabled"] = o.FloodUnknownMulticastTrafficEnabled
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchRoutingMulticast200ResponseOverridesInner struct {
	value *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner
	isSet bool
}

func (v NullableGetNetworkSwitchRoutingMulticast200ResponseOverridesInner) Get() *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner {
	return v.value
}

func (v *NullableGetNetworkSwitchRoutingMulticast200ResponseOverridesInner) Set(val *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchRoutingMulticast200ResponseOverridesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchRoutingMulticast200ResponseOverridesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchRoutingMulticast200ResponseOverridesInner(val *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) *NullableGetNetworkSwitchRoutingMulticast200ResponseOverridesInner {
	return &NullableGetNetworkSwitchRoutingMulticast200ResponseOverridesInner{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchRoutingMulticast200ResponseOverridesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchRoutingMulticast200ResponseOverridesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


