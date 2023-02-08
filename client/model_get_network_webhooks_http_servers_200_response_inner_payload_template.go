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

// GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate The payload template to use when posting data to the HTTP server.
type GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate struct {
	// The ID of the payload template.
	PayloadTemplateId *string `json:"payloadTemplateId,omitempty"`
	// The name of the payload template.
	Name *string `json:"name,omitempty"`
}

// NewGetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate instantiates a new GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate() *GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate {
	this := GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate{}
	return &this
}

// NewGetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplateWithDefaults instantiates a new GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplateWithDefaults() *GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate {
	this := GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate{}
	return &this
}

// GetPayloadTemplateId returns the PayloadTemplateId field value if set, zero value otherwise.
func (o *GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate) GetPayloadTemplateId() string {
	if o == nil || isNil(o.PayloadTemplateId) {
		var ret string
		return ret
	}
	return *o.PayloadTemplateId
}

// GetPayloadTemplateIdOk returns a tuple with the PayloadTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate) GetPayloadTemplateIdOk() (*string, bool) {
	if o == nil || isNil(o.PayloadTemplateId) {
    return nil, false
	}
	return o.PayloadTemplateId, true
}

// HasPayloadTemplateId returns a boolean if a field has been set.
func (o *GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate) HasPayloadTemplateId() bool {
	if o != nil && !isNil(o.PayloadTemplateId) {
		return true
	}

	return false
}

// SetPayloadTemplateId gets a reference to the given string and assigns it to the PayloadTemplateId field.
func (o *GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate) SetPayloadTemplateId(v string) {
	o.PayloadTemplateId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate) SetName(v string) {
	o.Name = &v
}

func (o GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PayloadTemplateId) {
		toSerialize["payloadTemplateId"] = o.PayloadTemplateId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate struct {
	value *GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate
	isSet bool
}

func (v NullableGetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate) Get() *GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate {
	return v.value
}

func (v *NullableGetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate) Set(val *GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate(val *GetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate) *NullableGetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate {
	return &NullableGetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate{value: val, isSet: true}
}

func (v NullableGetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWebhooksHttpServers200ResponseInnerPayloadTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


