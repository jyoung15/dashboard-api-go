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

// checks if the GetDeviceWirelessElectronicShelfLabel200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceWirelessElectronicShelfLabel200Response{}

// GetDeviceWirelessElectronicShelfLabel200Response struct for GetDeviceWirelessElectronicShelfLabel200Response
type GetDeviceWirelessElectronicShelfLabel200Response struct {
	// An identifier for the device used by the ESL system
	ApEslId *int32 `json:"apEslId,omitempty"`
	// The serial number of the device
	Serial *string `json:"serial,omitempty"`
	// Desired ESL channel for the device, or 'Auto' (case insensitive) to use the recommended channel
	Channel *string `json:"channel,omitempty"`
	// Turn ESL features on and off for this device
	Enabled *bool `json:"enabled,omitempty"`
	// The identifier for the device's network
	NetworkId *string `json:"networkId,omitempty"`
	// Hostname of the ESL management service
	Hostname *string `json:"hostname,omitempty"`
	// The service providing ESL functionality
	Provider *string `json:"provider,omitempty"`
}

// NewGetDeviceWirelessElectronicShelfLabel200Response instantiates a new GetDeviceWirelessElectronicShelfLabel200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceWirelessElectronicShelfLabel200Response() *GetDeviceWirelessElectronicShelfLabel200Response {
	this := GetDeviceWirelessElectronicShelfLabel200Response{}
	return &this
}

// NewGetDeviceWirelessElectronicShelfLabel200ResponseWithDefaults instantiates a new GetDeviceWirelessElectronicShelfLabel200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceWirelessElectronicShelfLabel200ResponseWithDefaults() *GetDeviceWirelessElectronicShelfLabel200Response {
	this := GetDeviceWirelessElectronicShelfLabel200Response{}
	return &this
}

// GetApEslId returns the ApEslId field value if set, zero value otherwise.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetApEslId() int32 {
	if o == nil || IsNil(o.ApEslId) {
		var ret int32
		return ret
	}
	return *o.ApEslId
}

// GetApEslIdOk returns a tuple with the ApEslId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetApEslIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ApEslId) {
		return nil, false
	}
	return o.ApEslId, true
}

// HasApEslId returns a boolean if a field has been set.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) HasApEslId() bool {
	if o != nil && !IsNil(o.ApEslId) {
		return true
	}

	return false
}

// SetApEslId gets a reference to the given int32 and assigns it to the ApEslId field.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) SetApEslId(v int32) {
	o.ApEslId = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) SetSerial(v string) {
	o.Serial = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) SetChannel(v string) {
	o.Channel = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetNetworkId() string {
	if o == nil || IsNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) SetHostname(v string) {
	o.Hostname = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *GetDeviceWirelessElectronicShelfLabel200Response) SetProvider(v string) {
	o.Provider = &v
}

func (o GetDeviceWirelessElectronicShelfLabel200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceWirelessElectronicShelfLabel200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApEslId) {
		toSerialize["apEslId"] = o.ApEslId
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	return toSerialize, nil
}

type NullableGetDeviceWirelessElectronicShelfLabel200Response struct {
	value *GetDeviceWirelessElectronicShelfLabel200Response
	isSet bool
}

func (v NullableGetDeviceWirelessElectronicShelfLabel200Response) Get() *GetDeviceWirelessElectronicShelfLabel200Response {
	return v.value
}

func (v *NullableGetDeviceWirelessElectronicShelfLabel200Response) Set(val *GetDeviceWirelessElectronicShelfLabel200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceWirelessElectronicShelfLabel200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceWirelessElectronicShelfLabel200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceWirelessElectronicShelfLabel200Response(val *GetDeviceWirelessElectronicShelfLabel200Response) *NullableGetDeviceWirelessElectronicShelfLabel200Response {
	return &NullableGetDeviceWirelessElectronicShelfLabel200Response{value: val, isSet: true}
}

func (v NullableGetDeviceWirelessElectronicShelfLabel200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceWirelessElectronicShelfLabel200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

