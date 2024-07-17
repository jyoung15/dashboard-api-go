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

// checks if the CreateDeviceLiveToolsArpTableRequestCallback type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDeviceLiveToolsArpTableRequestCallback{}

// CreateDeviceLiveToolsArpTableRequestCallback Details for the callback. Please include either an httpServerId OR url and sharedSecret
type CreateDeviceLiveToolsArpTableRequestCallback struct {
	// The callback URL for the webhook target. If using this field, please also specify a sharedSecret.
	Url *string `json:"url,omitempty"`
	// A shared secret that will be included in the requests sent to the callback URL. It can be used to verify that the request was sent by Meraki. If using this field, please also specify an url.
	SharedSecret *string `json:"sharedSecret,omitempty"`
	HttpServer *CreateDeviceLiveToolsArpTableRequestCallbackHttpServer `json:"httpServer,omitempty"`
	PayloadTemplate *CreateDeviceLiveToolsArpTableRequestCallbackPayloadTemplate `json:"payloadTemplate,omitempty"`
}

// NewCreateDeviceLiveToolsArpTableRequestCallback instantiates a new CreateDeviceLiveToolsArpTableRequestCallback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceLiveToolsArpTableRequestCallback() *CreateDeviceLiveToolsArpTableRequestCallback {
	this := CreateDeviceLiveToolsArpTableRequestCallback{}
	return &this
}

// NewCreateDeviceLiveToolsArpTableRequestCallbackWithDefaults instantiates a new CreateDeviceLiveToolsArpTableRequestCallback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceLiveToolsArpTableRequestCallbackWithDefaults() *CreateDeviceLiveToolsArpTableRequestCallback {
	this := CreateDeviceLiveToolsArpTableRequestCallback{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsArpTableRequestCallback) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsArpTableRequestCallback) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsArpTableRequestCallback) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CreateDeviceLiveToolsArpTableRequestCallback) SetUrl(v string) {
	o.Url = &v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsArpTableRequestCallback) GetSharedSecret() string {
	if o == nil || IsNil(o.SharedSecret) {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsArpTableRequestCallback) GetSharedSecretOk() (*string, bool) {
	if o == nil || IsNil(o.SharedSecret) {
		return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsArpTableRequestCallback) HasSharedSecret() bool {
	if o != nil && !IsNil(o.SharedSecret) {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *CreateDeviceLiveToolsArpTableRequestCallback) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetHttpServer returns the HttpServer field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsArpTableRequestCallback) GetHttpServer() CreateDeviceLiveToolsArpTableRequestCallbackHttpServer {
	if o == nil || IsNil(o.HttpServer) {
		var ret CreateDeviceLiveToolsArpTableRequestCallbackHttpServer
		return ret
	}
	return *o.HttpServer
}

// GetHttpServerOk returns a tuple with the HttpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsArpTableRequestCallback) GetHttpServerOk() (*CreateDeviceLiveToolsArpTableRequestCallbackHttpServer, bool) {
	if o == nil || IsNil(o.HttpServer) {
		return nil, false
	}
	return o.HttpServer, true
}

// HasHttpServer returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsArpTableRequestCallback) HasHttpServer() bool {
	if o != nil && !IsNil(o.HttpServer) {
		return true
	}

	return false
}

// SetHttpServer gets a reference to the given CreateDeviceLiveToolsArpTableRequestCallbackHttpServer and assigns it to the HttpServer field.
func (o *CreateDeviceLiveToolsArpTableRequestCallback) SetHttpServer(v CreateDeviceLiveToolsArpTableRequestCallbackHttpServer) {
	o.HttpServer = &v
}

// GetPayloadTemplate returns the PayloadTemplate field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsArpTableRequestCallback) GetPayloadTemplate() CreateDeviceLiveToolsArpTableRequestCallbackPayloadTemplate {
	if o == nil || IsNil(o.PayloadTemplate) {
		var ret CreateDeviceLiveToolsArpTableRequestCallbackPayloadTemplate
		return ret
	}
	return *o.PayloadTemplate
}

// GetPayloadTemplateOk returns a tuple with the PayloadTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsArpTableRequestCallback) GetPayloadTemplateOk() (*CreateDeviceLiveToolsArpTableRequestCallbackPayloadTemplate, bool) {
	if o == nil || IsNil(o.PayloadTemplate) {
		return nil, false
	}
	return o.PayloadTemplate, true
}

// HasPayloadTemplate returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsArpTableRequestCallback) HasPayloadTemplate() bool {
	if o != nil && !IsNil(o.PayloadTemplate) {
		return true
	}

	return false
}

// SetPayloadTemplate gets a reference to the given CreateDeviceLiveToolsArpTableRequestCallbackPayloadTemplate and assigns it to the PayloadTemplate field.
func (o *CreateDeviceLiveToolsArpTableRequestCallback) SetPayloadTemplate(v CreateDeviceLiveToolsArpTableRequestCallbackPayloadTemplate) {
	o.PayloadTemplate = &v
}

func (o CreateDeviceLiveToolsArpTableRequestCallback) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDeviceLiveToolsArpTableRequestCallback) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.SharedSecret) {
		toSerialize["sharedSecret"] = o.SharedSecret
	}
	if !IsNil(o.HttpServer) {
		toSerialize["httpServer"] = o.HttpServer
	}
	if !IsNil(o.PayloadTemplate) {
		toSerialize["payloadTemplate"] = o.PayloadTemplate
	}
	return toSerialize, nil
}

type NullableCreateDeviceLiveToolsArpTableRequestCallback struct {
	value *CreateDeviceLiveToolsArpTableRequestCallback
	isSet bool
}

func (v NullableCreateDeviceLiveToolsArpTableRequestCallback) Get() *CreateDeviceLiveToolsArpTableRequestCallback {
	return v.value
}

func (v *NullableCreateDeviceLiveToolsArpTableRequestCallback) Set(val *CreateDeviceLiveToolsArpTableRequestCallback) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceLiveToolsArpTableRequestCallback) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceLiveToolsArpTableRequestCallback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceLiveToolsArpTableRequestCallback(val *CreateDeviceLiveToolsArpTableRequestCallback) *NullableCreateDeviceLiveToolsArpTableRequestCallback {
	return &NullableCreateDeviceLiveToolsArpTableRequestCallback{value: val, isSet: true}
}

func (v NullableCreateDeviceLiveToolsArpTableRequestCallback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceLiveToolsArpTableRequestCallback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

