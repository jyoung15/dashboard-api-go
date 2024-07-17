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

// checks if the OrganizationsOrganizationIdActionBatchesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationsOrganizationIdActionBatchesPostRequest{}

// OrganizationsOrganizationIdActionBatchesPostRequest struct for OrganizationsOrganizationIdActionBatchesPostRequest
type OrganizationsOrganizationIdActionBatchesPostRequest struct {
	Organization *DevicesSerialLiveToolsArpTablePostRequestOrganization `json:"organization,omitempty"`
	Network *DevicesSerialLiveToolsArpTablePostRequestOrganization `json:"network,omitempty"`
	SentAt *string `json:"sentAt,omitempty"`
	CallbackId *string `json:"callbackId,omitempty"`
	Message *GetOrganizationActionBatches200ResponseInner `json:"message,omitempty"`
}

// NewOrganizationsOrganizationIdActionBatchesPostRequest instantiates a new OrganizationsOrganizationIdActionBatchesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdActionBatchesPostRequest() *OrganizationsOrganizationIdActionBatchesPostRequest {
	this := OrganizationsOrganizationIdActionBatchesPostRequest{}
	return &this
}

// NewOrganizationsOrganizationIdActionBatchesPostRequestWithDefaults instantiates a new OrganizationsOrganizationIdActionBatchesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdActionBatchesPostRequestWithDefaults() *OrganizationsOrganizationIdActionBatchesPostRequest {
	this := OrganizationsOrganizationIdActionBatchesPostRequest{}
	return &this
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdActionBatchesPostRequest) GetOrganization() DevicesSerialLiveToolsArpTablePostRequestOrganization {
	if o == nil || IsNil(o.Organization) {
		var ret DevicesSerialLiveToolsArpTablePostRequestOrganization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesPostRequest) GetOrganizationOk() (*DevicesSerialLiveToolsArpTablePostRequestOrganization, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdActionBatchesPostRequest) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given DevicesSerialLiveToolsArpTablePostRequestOrganization and assigns it to the Organization field.
func (o *OrganizationsOrganizationIdActionBatchesPostRequest) SetOrganization(v DevicesSerialLiveToolsArpTablePostRequestOrganization) {
	o.Organization = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdActionBatchesPostRequest) GetNetwork() DevicesSerialLiveToolsArpTablePostRequestOrganization {
	if o == nil || IsNil(o.Network) {
		var ret DevicesSerialLiveToolsArpTablePostRequestOrganization
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesPostRequest) GetNetworkOk() (*DevicesSerialLiveToolsArpTablePostRequestOrganization, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdActionBatchesPostRequest) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given DevicesSerialLiveToolsArpTablePostRequestOrganization and assigns it to the Network field.
func (o *OrganizationsOrganizationIdActionBatchesPostRequest) SetNetwork(v DevicesSerialLiveToolsArpTablePostRequestOrganization) {
	o.Network = &v
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdActionBatchesPostRequest) GetSentAt() string {
	if o == nil || IsNil(o.SentAt) {
		var ret string
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesPostRequest) GetSentAtOk() (*string, bool) {
	if o == nil || IsNil(o.SentAt) {
		return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdActionBatchesPostRequest) HasSentAt() bool {
	if o != nil && !IsNil(o.SentAt) {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given string and assigns it to the SentAt field.
func (o *OrganizationsOrganizationIdActionBatchesPostRequest) SetSentAt(v string) {
	o.SentAt = &v
}

// GetCallbackId returns the CallbackId field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdActionBatchesPostRequest) GetCallbackId() string {
	if o == nil || IsNil(o.CallbackId) {
		var ret string
		return ret
	}
	return *o.CallbackId
}

// GetCallbackIdOk returns a tuple with the CallbackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesPostRequest) GetCallbackIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackId) {
		return nil, false
	}
	return o.CallbackId, true
}

// HasCallbackId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdActionBatchesPostRequest) HasCallbackId() bool {
	if o != nil && !IsNil(o.CallbackId) {
		return true
	}

	return false
}

// SetCallbackId gets a reference to the given string and assigns it to the CallbackId field.
func (o *OrganizationsOrganizationIdActionBatchesPostRequest) SetCallbackId(v string) {
	o.CallbackId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdActionBatchesPostRequest) GetMessage() GetOrganizationActionBatches200ResponseInner {
	if o == nil || IsNil(o.Message) {
		var ret GetOrganizationActionBatches200ResponseInner
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesPostRequest) GetMessageOk() (*GetOrganizationActionBatches200ResponseInner, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdActionBatchesPostRequest) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given GetOrganizationActionBatches200ResponseInner and assigns it to the Message field.
func (o *OrganizationsOrganizationIdActionBatchesPostRequest) SetMessage(v GetOrganizationActionBatches200ResponseInner) {
	o.Message = &v
}

func (o OrganizationsOrganizationIdActionBatchesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationsOrganizationIdActionBatchesPostRequest) ToMap() (map[string]interface{}, error) {
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

type NullableOrganizationsOrganizationIdActionBatchesPostRequest struct {
	value *OrganizationsOrganizationIdActionBatchesPostRequest
	isSet bool
}

func (v NullableOrganizationsOrganizationIdActionBatchesPostRequest) Get() *OrganizationsOrganizationIdActionBatchesPostRequest {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdActionBatchesPostRequest) Set(val *OrganizationsOrganizationIdActionBatchesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdActionBatchesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdActionBatchesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdActionBatchesPostRequest(val *OrganizationsOrganizationIdActionBatchesPostRequest) *NullableOrganizationsOrganizationIdActionBatchesPostRequest {
	return &NullableOrganizationsOrganizationIdActionBatchesPostRequest{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdActionBatchesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdActionBatchesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


