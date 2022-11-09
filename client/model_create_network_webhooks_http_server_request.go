/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateNetworkWebhooksHttpServerRequest struct for CreateNetworkWebhooksHttpServerRequest
type CreateNetworkWebhooksHttpServerRequest struct {
	// A name for easy reference to the HTTP server
	Name string `json:"name"`
	// The URL of the HTTP server. Once set, cannot be updated.
	Url string `json:"url"`
	// A shared secret that will be included in POSTs sent to the HTTP server. This secret can be used to verify that the request was sent by Meraki.
	SharedSecret *string `json:"sharedSecret,omitempty"`
	PayloadTemplate *CreateNetworkWebhooksHttpServerRequestPayloadTemplate `json:"payloadTemplate,omitempty"`
}

// NewCreateNetworkWebhooksHttpServerRequest instantiates a new CreateNetworkWebhooksHttpServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWebhooksHttpServerRequest(name string, url string) *CreateNetworkWebhooksHttpServerRequest {
	this := CreateNetworkWebhooksHttpServerRequest{}
	this.Name = name
	this.Url = url
	return &this
}

// NewCreateNetworkWebhooksHttpServerRequestWithDefaults instantiates a new CreateNetworkWebhooksHttpServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWebhooksHttpServerRequestWithDefaults() *CreateNetworkWebhooksHttpServerRequest {
	this := CreateNetworkWebhooksHttpServerRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNetworkWebhooksHttpServerRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksHttpServerRequest) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkWebhooksHttpServerRequest) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *CreateNetworkWebhooksHttpServerRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksHttpServerRequest) GetUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CreateNetworkWebhooksHttpServerRequest) SetUrl(v string) {
	o.Url = v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *CreateNetworkWebhooksHttpServerRequest) GetSharedSecret() string {
	if o == nil || isNil(o.SharedSecret) {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksHttpServerRequest) GetSharedSecretOk() (*string, bool) {
	if o == nil || isNil(o.SharedSecret) {
    return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *CreateNetworkWebhooksHttpServerRequest) HasSharedSecret() bool {
	if o != nil && !isNil(o.SharedSecret) {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *CreateNetworkWebhooksHttpServerRequest) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetPayloadTemplate returns the PayloadTemplate field value if set, zero value otherwise.
func (o *CreateNetworkWebhooksHttpServerRequest) GetPayloadTemplate() CreateNetworkWebhooksHttpServerRequestPayloadTemplate {
	if o == nil || isNil(o.PayloadTemplate) {
		var ret CreateNetworkWebhooksHttpServerRequestPayloadTemplate
		return ret
	}
	return *o.PayloadTemplate
}

// GetPayloadTemplateOk returns a tuple with the PayloadTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksHttpServerRequest) GetPayloadTemplateOk() (*CreateNetworkWebhooksHttpServerRequestPayloadTemplate, bool) {
	if o == nil || isNil(o.PayloadTemplate) {
    return nil, false
	}
	return o.PayloadTemplate, true
}

// HasPayloadTemplate returns a boolean if a field has been set.
func (o *CreateNetworkWebhooksHttpServerRequest) HasPayloadTemplate() bool {
	if o != nil && !isNil(o.PayloadTemplate) {
		return true
	}

	return false
}

// SetPayloadTemplate gets a reference to the given CreateNetworkWebhooksHttpServerRequestPayloadTemplate and assigns it to the PayloadTemplate field.
func (o *CreateNetworkWebhooksHttpServerRequest) SetPayloadTemplate(v CreateNetworkWebhooksHttpServerRequestPayloadTemplate) {
	o.PayloadTemplate = &v
}

func (o CreateNetworkWebhooksHttpServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.SharedSecret) {
		toSerialize["sharedSecret"] = o.SharedSecret
	}
	if !isNil(o.PayloadTemplate) {
		toSerialize["payloadTemplate"] = o.PayloadTemplate
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkWebhooksHttpServerRequest struct {
	value *CreateNetworkWebhooksHttpServerRequest
	isSet bool
}

func (v NullableCreateNetworkWebhooksHttpServerRequest) Get() *CreateNetworkWebhooksHttpServerRequest {
	return v.value
}

func (v *NullableCreateNetworkWebhooksHttpServerRequest) Set(val *CreateNetworkWebhooksHttpServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWebhooksHttpServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWebhooksHttpServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWebhooksHttpServerRequest(val *CreateNetworkWebhooksHttpServerRequest) *NullableCreateNetworkWebhooksHttpServerRequest {
	return &NullableCreateNetworkWebhooksHttpServerRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkWebhooksHttpServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWebhooksHttpServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

