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

// UpdateNetworkWirelessSsidEapOverrideRequest struct for UpdateNetworkWirelessSsidEapOverrideRequest
type UpdateNetworkWirelessSsidEapOverrideRequest struct {
	// General EAP timeout in seconds.
	Timeout *int32 `json:"timeout,omitempty"`
	Identity *UpdateNetworkWirelessSsidEapOverrideRequestIdentity `json:"identity,omitempty"`
	// Maximum number of general EAP retries.
	MaxRetries *int32 `json:"maxRetries,omitempty"`
	EapolKey *UpdateNetworkWirelessSsidEapOverrideRequestEapolKey `json:"eapolKey,omitempty"`
}

// NewUpdateNetworkWirelessSsidEapOverrideRequest instantiates a new UpdateNetworkWirelessSsidEapOverrideRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidEapOverrideRequest() *UpdateNetworkWirelessSsidEapOverrideRequest {
	this := UpdateNetworkWirelessSsidEapOverrideRequest{}
	return &this
}

// NewUpdateNetworkWirelessSsidEapOverrideRequestWithDefaults instantiates a new UpdateNetworkWirelessSsidEapOverrideRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidEapOverrideRequestWithDefaults() *UpdateNetworkWirelessSsidEapOverrideRequest {
	this := UpdateNetworkWirelessSsidEapOverrideRequest{}
	return &this
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidEapOverrideRequest) GetTimeout() int32 {
	if o == nil || isNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidEapOverrideRequest) GetTimeoutOk() (*int32, bool) {
	if o == nil || isNil(o.Timeout) {
    return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidEapOverrideRequest) HasTimeout() bool {
	if o != nil && !isNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *UpdateNetworkWirelessSsidEapOverrideRequest) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidEapOverrideRequest) GetIdentity() UpdateNetworkWirelessSsidEapOverrideRequestIdentity {
	if o == nil || isNil(o.Identity) {
		var ret UpdateNetworkWirelessSsidEapOverrideRequestIdentity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidEapOverrideRequest) GetIdentityOk() (*UpdateNetworkWirelessSsidEapOverrideRequestIdentity, bool) {
	if o == nil || isNil(o.Identity) {
    return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidEapOverrideRequest) HasIdentity() bool {
	if o != nil && !isNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given UpdateNetworkWirelessSsidEapOverrideRequestIdentity and assigns it to the Identity field.
func (o *UpdateNetworkWirelessSsidEapOverrideRequest) SetIdentity(v UpdateNetworkWirelessSsidEapOverrideRequestIdentity) {
	o.Identity = &v
}

// GetMaxRetries returns the MaxRetries field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidEapOverrideRequest) GetMaxRetries() int32 {
	if o == nil || isNil(o.MaxRetries) {
		var ret int32
		return ret
	}
	return *o.MaxRetries
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidEapOverrideRequest) GetMaxRetriesOk() (*int32, bool) {
	if o == nil || isNil(o.MaxRetries) {
    return nil, false
	}
	return o.MaxRetries, true
}

// HasMaxRetries returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidEapOverrideRequest) HasMaxRetries() bool {
	if o != nil && !isNil(o.MaxRetries) {
		return true
	}

	return false
}

// SetMaxRetries gets a reference to the given int32 and assigns it to the MaxRetries field.
func (o *UpdateNetworkWirelessSsidEapOverrideRequest) SetMaxRetries(v int32) {
	o.MaxRetries = &v
}

// GetEapolKey returns the EapolKey field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidEapOverrideRequest) GetEapolKey() UpdateNetworkWirelessSsidEapOverrideRequestEapolKey {
	if o == nil || isNil(o.EapolKey) {
		var ret UpdateNetworkWirelessSsidEapOverrideRequestEapolKey
		return ret
	}
	return *o.EapolKey
}

// GetEapolKeyOk returns a tuple with the EapolKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidEapOverrideRequest) GetEapolKeyOk() (*UpdateNetworkWirelessSsidEapOverrideRequestEapolKey, bool) {
	if o == nil || isNil(o.EapolKey) {
    return nil, false
	}
	return o.EapolKey, true
}

// HasEapolKey returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidEapOverrideRequest) HasEapolKey() bool {
	if o != nil && !isNil(o.EapolKey) {
		return true
	}

	return false
}

// SetEapolKey gets a reference to the given UpdateNetworkWirelessSsidEapOverrideRequestEapolKey and assigns it to the EapolKey field.
func (o *UpdateNetworkWirelessSsidEapOverrideRequest) SetEapolKey(v UpdateNetworkWirelessSsidEapOverrideRequestEapolKey) {
	o.EapolKey = &v
}

func (o UpdateNetworkWirelessSsidEapOverrideRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !isNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !isNil(o.MaxRetries) {
		toSerialize["maxRetries"] = o.MaxRetries
	}
	if !isNil(o.EapolKey) {
		toSerialize["eapolKey"] = o.EapolKey
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkWirelessSsidEapOverrideRequest struct {
	value *UpdateNetworkWirelessSsidEapOverrideRequest
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidEapOverrideRequest) Get() *UpdateNetworkWirelessSsidEapOverrideRequest {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidEapOverrideRequest) Set(val *UpdateNetworkWirelessSsidEapOverrideRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidEapOverrideRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidEapOverrideRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidEapOverrideRequest(val *UpdateNetworkWirelessSsidEapOverrideRequest) *NullableUpdateNetworkWirelessSsidEapOverrideRequest {
	return &NullableUpdateNetworkWirelessSsidEapOverrideRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidEapOverrideRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidEapOverrideRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


