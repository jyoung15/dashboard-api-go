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

// checks if the DevicesSerialLiveToolsWakeOnLanPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevicesSerialLiveToolsWakeOnLanPostRequest{}

// DevicesSerialLiveToolsWakeOnLanPostRequest struct for DevicesSerialLiveToolsWakeOnLanPostRequest
type DevicesSerialLiveToolsWakeOnLanPostRequest struct {
	Organization *DevicesSerialLiveToolsArpTablePostRequestOrganization `json:"organization,omitempty"`
	Network *DevicesSerialLiveToolsArpTablePostRequestOrganization `json:"network,omitempty"`
	SentAt *string `json:"sentAt,omitempty"`
	CallbackId *string `json:"callbackId,omitempty"`
	Message *DevicesSerialLiveToolsWakeOnLanPostRequestMessage `json:"message,omitempty"`
}

// NewDevicesSerialLiveToolsWakeOnLanPostRequest instantiates a new DevicesSerialLiveToolsWakeOnLanPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialLiveToolsWakeOnLanPostRequest() *DevicesSerialLiveToolsWakeOnLanPostRequest {
	this := DevicesSerialLiveToolsWakeOnLanPostRequest{}
	return &this
}

// NewDevicesSerialLiveToolsWakeOnLanPostRequestWithDefaults instantiates a new DevicesSerialLiveToolsWakeOnLanPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialLiveToolsWakeOnLanPostRequestWithDefaults() *DevicesSerialLiveToolsWakeOnLanPostRequest {
	this := DevicesSerialLiveToolsWakeOnLanPostRequest{}
	return &this
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsWakeOnLanPostRequest) GetOrganization() DevicesSerialLiveToolsArpTablePostRequestOrganization {
	if o == nil || IsNil(o.Organization) {
		var ret DevicesSerialLiveToolsArpTablePostRequestOrganization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsWakeOnLanPostRequest) GetOrganizationOk() (*DevicesSerialLiveToolsArpTablePostRequestOrganization, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsWakeOnLanPostRequest) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given DevicesSerialLiveToolsArpTablePostRequestOrganization and assigns it to the Organization field.
func (o *DevicesSerialLiveToolsWakeOnLanPostRequest) SetOrganization(v DevicesSerialLiveToolsArpTablePostRequestOrganization) {
	o.Organization = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsWakeOnLanPostRequest) GetNetwork() DevicesSerialLiveToolsArpTablePostRequestOrganization {
	if o == nil || IsNil(o.Network) {
		var ret DevicesSerialLiveToolsArpTablePostRequestOrganization
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsWakeOnLanPostRequest) GetNetworkOk() (*DevicesSerialLiveToolsArpTablePostRequestOrganization, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsWakeOnLanPostRequest) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given DevicesSerialLiveToolsArpTablePostRequestOrganization and assigns it to the Network field.
func (o *DevicesSerialLiveToolsWakeOnLanPostRequest) SetNetwork(v DevicesSerialLiveToolsArpTablePostRequestOrganization) {
	o.Network = &v
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsWakeOnLanPostRequest) GetSentAt() string {
	if o == nil || IsNil(o.SentAt) {
		var ret string
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsWakeOnLanPostRequest) GetSentAtOk() (*string, bool) {
	if o == nil || IsNil(o.SentAt) {
		return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsWakeOnLanPostRequest) HasSentAt() bool {
	if o != nil && !IsNil(o.SentAt) {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given string and assigns it to the SentAt field.
func (o *DevicesSerialLiveToolsWakeOnLanPostRequest) SetSentAt(v string) {
	o.SentAt = &v
}

// GetCallbackId returns the CallbackId field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsWakeOnLanPostRequest) GetCallbackId() string {
	if o == nil || IsNil(o.CallbackId) {
		var ret string
		return ret
	}
	return *o.CallbackId
}

// GetCallbackIdOk returns a tuple with the CallbackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsWakeOnLanPostRequest) GetCallbackIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackId) {
		return nil, false
	}
	return o.CallbackId, true
}

// HasCallbackId returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsWakeOnLanPostRequest) HasCallbackId() bool {
	if o != nil && !IsNil(o.CallbackId) {
		return true
	}

	return false
}

// SetCallbackId gets a reference to the given string and assigns it to the CallbackId field.
func (o *DevicesSerialLiveToolsWakeOnLanPostRequest) SetCallbackId(v string) {
	o.CallbackId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsWakeOnLanPostRequest) GetMessage() DevicesSerialLiveToolsWakeOnLanPostRequestMessage {
	if o == nil || IsNil(o.Message) {
		var ret DevicesSerialLiveToolsWakeOnLanPostRequestMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsWakeOnLanPostRequest) GetMessageOk() (*DevicesSerialLiveToolsWakeOnLanPostRequestMessage, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsWakeOnLanPostRequest) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given DevicesSerialLiveToolsWakeOnLanPostRequestMessage and assigns it to the Message field.
func (o *DevicesSerialLiveToolsWakeOnLanPostRequest) SetMessage(v DevicesSerialLiveToolsWakeOnLanPostRequestMessage) {
	o.Message = &v
}

func (o DevicesSerialLiveToolsWakeOnLanPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevicesSerialLiveToolsWakeOnLanPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.SentAt) {
		toSerialize["sentAt"] = o.SentAt
	}
	if !IsNil(o.CallbackId) {
		toSerialize["callbackId"] = o.CallbackId
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableDevicesSerialLiveToolsWakeOnLanPostRequest struct {
	value *DevicesSerialLiveToolsWakeOnLanPostRequest
	isSet bool
}

func (v NullableDevicesSerialLiveToolsWakeOnLanPostRequest) Get() *DevicesSerialLiveToolsWakeOnLanPostRequest {
	return v.value
}

func (v *NullableDevicesSerialLiveToolsWakeOnLanPostRequest) Set(val *DevicesSerialLiveToolsWakeOnLanPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialLiveToolsWakeOnLanPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialLiveToolsWakeOnLanPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialLiveToolsWakeOnLanPostRequest(val *DevicesSerialLiveToolsWakeOnLanPostRequest) *NullableDevicesSerialLiveToolsWakeOnLanPostRequest {
	return &NullableDevicesSerialLiveToolsWakeOnLanPostRequest{value: val, isSet: true}
}

func (v NullableDevicesSerialLiveToolsWakeOnLanPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialLiveToolsWakeOnLanPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


