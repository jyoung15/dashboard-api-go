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

// checks if the UpdateNetworkSwitchRoutingOspfRequestV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchRoutingOspfRequestV3{}

// UpdateNetworkSwitchRoutingOspfRequestV3 OSPF v3 configuration
type UpdateNetworkSwitchRoutingOspfRequestV3 struct {
	// Boolean value to enable or disable V3 OSPF routing. OSPF V3 routing is disabled by default.
	Enabled *bool `json:"enabled,omitempty"`
	// Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds.
	HelloTimerInSeconds *int32 `json:"helloTimerInSeconds,omitempty"`
	// Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535
	DeadTimerInSeconds *int32 `json:"deadTimerInSeconds,omitempty"`
	// OSPF v3 areas
	Areas []UpdateNetworkSwitchRoutingOspfRequestAreasInner `json:"areas,omitempty"`
}

// NewUpdateNetworkSwitchRoutingOspfRequestV3 instantiates a new UpdateNetworkSwitchRoutingOspfRequestV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchRoutingOspfRequestV3() *UpdateNetworkSwitchRoutingOspfRequestV3 {
	this := UpdateNetworkSwitchRoutingOspfRequestV3{}
	return &this
}

// NewUpdateNetworkSwitchRoutingOspfRequestV3WithDefaults instantiates a new UpdateNetworkSwitchRoutingOspfRequestV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchRoutingOspfRequestV3WithDefaults() *UpdateNetworkSwitchRoutingOspfRequestV3 {
	this := UpdateNetworkSwitchRoutingOspfRequestV3{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchRoutingOspfRequestV3) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequestV3) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequestV3) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkSwitchRoutingOspfRequestV3) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHelloTimerInSeconds returns the HelloTimerInSeconds field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchRoutingOspfRequestV3) GetHelloTimerInSeconds() int32 {
	if o == nil || IsNil(o.HelloTimerInSeconds) {
		var ret int32
		return ret
	}
	return *o.HelloTimerInSeconds
}

// GetHelloTimerInSecondsOk returns a tuple with the HelloTimerInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequestV3) GetHelloTimerInSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.HelloTimerInSeconds) {
		return nil, false
	}
	return o.HelloTimerInSeconds, true
}

// HasHelloTimerInSeconds returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequestV3) HasHelloTimerInSeconds() bool {
	if o != nil && !IsNil(o.HelloTimerInSeconds) {
		return true
	}

	return false
}

// SetHelloTimerInSeconds gets a reference to the given int32 and assigns it to the HelloTimerInSeconds field.
func (o *UpdateNetworkSwitchRoutingOspfRequestV3) SetHelloTimerInSeconds(v int32) {
	o.HelloTimerInSeconds = &v
}

// GetDeadTimerInSeconds returns the DeadTimerInSeconds field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchRoutingOspfRequestV3) GetDeadTimerInSeconds() int32 {
	if o == nil || IsNil(o.DeadTimerInSeconds) {
		var ret int32
		return ret
	}
	return *o.DeadTimerInSeconds
}

// GetDeadTimerInSecondsOk returns a tuple with the DeadTimerInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequestV3) GetDeadTimerInSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.DeadTimerInSeconds) {
		return nil, false
	}
	return o.DeadTimerInSeconds, true
}

// HasDeadTimerInSeconds returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequestV3) HasDeadTimerInSeconds() bool {
	if o != nil && !IsNil(o.DeadTimerInSeconds) {
		return true
	}

	return false
}

// SetDeadTimerInSeconds gets a reference to the given int32 and assigns it to the DeadTimerInSeconds field.
func (o *UpdateNetworkSwitchRoutingOspfRequestV3) SetDeadTimerInSeconds(v int32) {
	o.DeadTimerInSeconds = &v
}

// GetAreas returns the Areas field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchRoutingOspfRequestV3) GetAreas() []UpdateNetworkSwitchRoutingOspfRequestAreasInner {
	if o == nil || IsNil(o.Areas) {
		var ret []UpdateNetworkSwitchRoutingOspfRequestAreasInner
		return ret
	}
	return o.Areas
}

// GetAreasOk returns a tuple with the Areas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequestV3) GetAreasOk() ([]UpdateNetworkSwitchRoutingOspfRequestAreasInner, bool) {
	if o == nil || IsNil(o.Areas) {
		return nil, false
	}
	return o.Areas, true
}

// HasAreas returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequestV3) HasAreas() bool {
	if o != nil && !IsNil(o.Areas) {
		return true
	}

	return false
}

// SetAreas gets a reference to the given []UpdateNetworkSwitchRoutingOspfRequestAreasInner and assigns it to the Areas field.
func (o *UpdateNetworkSwitchRoutingOspfRequestV3) SetAreas(v []UpdateNetworkSwitchRoutingOspfRequestAreasInner) {
	o.Areas = v
}

func (o UpdateNetworkSwitchRoutingOspfRequestV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchRoutingOspfRequestV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.HelloTimerInSeconds) {
		toSerialize["helloTimerInSeconds"] = o.HelloTimerInSeconds
	}
	if !IsNil(o.DeadTimerInSeconds) {
		toSerialize["deadTimerInSeconds"] = o.DeadTimerInSeconds
	}
	if !IsNil(o.Areas) {
		toSerialize["areas"] = o.Areas
	}
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchRoutingOspfRequestV3 struct {
	value *UpdateNetworkSwitchRoutingOspfRequestV3
	isSet bool
}

func (v NullableUpdateNetworkSwitchRoutingOspfRequestV3) Get() *UpdateNetworkSwitchRoutingOspfRequestV3 {
	return v.value
}

func (v *NullableUpdateNetworkSwitchRoutingOspfRequestV3) Set(val *UpdateNetworkSwitchRoutingOspfRequestV3) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchRoutingOspfRequestV3) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchRoutingOspfRequestV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchRoutingOspfRequestV3(val *UpdateNetworkSwitchRoutingOspfRequestV3) *NullableUpdateNetworkSwitchRoutingOspfRequestV3 {
	return &NullableUpdateNetworkSwitchRoutingOspfRequestV3{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchRoutingOspfRequestV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchRoutingOspfRequestV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


