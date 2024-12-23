/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UpdateNetworkApplianceVpnBgpRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceVpnBgpRequest{}

// UpdateNetworkApplianceVpnBgpRequest struct for UpdateNetworkApplianceVpnBgpRequest
type UpdateNetworkApplianceVpnBgpRequest struct {
	// Boolean value to enable or disable the BGP configuration. When BGP is enabled, the asNumber (ASN) will be autopopulated with the preconfigured ASN at other Hubs or a default value if there is no ASN configured.
	Enabled bool `json:"enabled"`
	// An Autonomous System Number (ASN) is required if you are to run BGP and peer with another BGP Speaker outside of the Auto VPN domain. This ASN will be applied to the entire Auto VPN domain. The entire 4-byte ASN range is supported. So, the ASN must be an integer between 1 and 4294967295. When absent, this field is not updated. If no value exists then it defaults to 64512.
	AsNumber *int32 `json:"asNumber,omitempty"`
	// The iBGP holdtimer in seconds. The iBGP holdtimer must be an integer between 12 and 240. When absent, this field is not updated. If no value exists then it defaults to 240.
	IbgpHoldTimer *int32 `json:"ibgpHoldTimer,omitempty"`
	// List of BGP neighbors. This list replaces the existing set of neighbors. When absent, this field is not updated.
	Neighbors []UpdateNetworkApplianceVpnBgpRequestNeighborsInner `json:"neighbors,omitempty"`
}

type _UpdateNetworkApplianceVpnBgpRequest UpdateNetworkApplianceVpnBgpRequest

// NewUpdateNetworkApplianceVpnBgpRequest instantiates a new UpdateNetworkApplianceVpnBgpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceVpnBgpRequest(enabled bool) *UpdateNetworkApplianceVpnBgpRequest {
	this := UpdateNetworkApplianceVpnBgpRequest{}
	this.Enabled = enabled
	return &this
}

// NewUpdateNetworkApplianceVpnBgpRequestWithDefaults instantiates a new UpdateNetworkApplianceVpnBgpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceVpnBgpRequestWithDefaults() *UpdateNetworkApplianceVpnBgpRequest {
	this := UpdateNetworkApplianceVpnBgpRequest{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *UpdateNetworkApplianceVpnBgpRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnBgpRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *UpdateNetworkApplianceVpnBgpRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAsNumber returns the AsNumber field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceVpnBgpRequest) GetAsNumber() int32 {
	if o == nil || IsNil(o.AsNumber) {
		var ret int32
		return ret
	}
	return *o.AsNumber
}

// GetAsNumberOk returns a tuple with the AsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnBgpRequest) GetAsNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.AsNumber) {
		return nil, false
	}
	return o.AsNumber, true
}

// HasAsNumber returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceVpnBgpRequest) HasAsNumber() bool {
	if o != nil && !IsNil(o.AsNumber) {
		return true
	}

	return false
}

// SetAsNumber gets a reference to the given int32 and assigns it to the AsNumber field.
func (o *UpdateNetworkApplianceVpnBgpRequest) SetAsNumber(v int32) {
	o.AsNumber = &v
}

// GetIbgpHoldTimer returns the IbgpHoldTimer field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceVpnBgpRequest) GetIbgpHoldTimer() int32 {
	if o == nil || IsNil(o.IbgpHoldTimer) {
		var ret int32
		return ret
	}
	return *o.IbgpHoldTimer
}

// GetIbgpHoldTimerOk returns a tuple with the IbgpHoldTimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnBgpRequest) GetIbgpHoldTimerOk() (*int32, bool) {
	if o == nil || IsNil(o.IbgpHoldTimer) {
		return nil, false
	}
	return o.IbgpHoldTimer, true
}

// HasIbgpHoldTimer returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceVpnBgpRequest) HasIbgpHoldTimer() bool {
	if o != nil && !IsNil(o.IbgpHoldTimer) {
		return true
	}

	return false
}

// SetIbgpHoldTimer gets a reference to the given int32 and assigns it to the IbgpHoldTimer field.
func (o *UpdateNetworkApplianceVpnBgpRequest) SetIbgpHoldTimer(v int32) {
	o.IbgpHoldTimer = &v
}

// GetNeighbors returns the Neighbors field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceVpnBgpRequest) GetNeighbors() []UpdateNetworkApplianceVpnBgpRequestNeighborsInner {
	if o == nil || IsNil(o.Neighbors) {
		var ret []UpdateNetworkApplianceVpnBgpRequestNeighborsInner
		return ret
	}
	return o.Neighbors
}

// GetNeighborsOk returns a tuple with the Neighbors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnBgpRequest) GetNeighborsOk() ([]UpdateNetworkApplianceVpnBgpRequestNeighborsInner, bool) {
	if o == nil || IsNil(o.Neighbors) {
		return nil, false
	}
	return o.Neighbors, true
}

// HasNeighbors returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceVpnBgpRequest) HasNeighbors() bool {
	if o != nil && !IsNil(o.Neighbors) {
		return true
	}

	return false
}

// SetNeighbors gets a reference to the given []UpdateNetworkApplianceVpnBgpRequestNeighborsInner and assigns it to the Neighbors field.
func (o *UpdateNetworkApplianceVpnBgpRequest) SetNeighbors(v []UpdateNetworkApplianceVpnBgpRequestNeighborsInner) {
	o.Neighbors = v
}

func (o UpdateNetworkApplianceVpnBgpRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceVpnBgpRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.AsNumber) {
		toSerialize["asNumber"] = o.AsNumber
	}
	if !IsNil(o.IbgpHoldTimer) {
		toSerialize["ibgpHoldTimer"] = o.IbgpHoldTimer
	}
	if !IsNil(o.Neighbors) {
		toSerialize["neighbors"] = o.Neighbors
	}
	return toSerialize, nil
}

func (o *UpdateNetworkApplianceVpnBgpRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateNetworkApplianceVpnBgpRequest := _UpdateNetworkApplianceVpnBgpRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateNetworkApplianceVpnBgpRequest)

	if err != nil {
		return err
	}

	*o = UpdateNetworkApplianceVpnBgpRequest(varUpdateNetworkApplianceVpnBgpRequest)

	return err
}

type NullableUpdateNetworkApplianceVpnBgpRequest struct {
	value *UpdateNetworkApplianceVpnBgpRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceVpnBgpRequest) Get() *UpdateNetworkApplianceVpnBgpRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceVpnBgpRequest) Set(val *UpdateNetworkApplianceVpnBgpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceVpnBgpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceVpnBgpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceVpnBgpRequest(val *UpdateNetworkApplianceVpnBgpRequest) *NullableUpdateNetworkApplianceVpnBgpRequest {
	return &NullableUpdateNetworkApplianceVpnBgpRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceVpnBgpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceVpnBgpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


