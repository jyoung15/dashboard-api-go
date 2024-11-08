/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetOrganizationWebhooksCallbacksStatus200ResponseWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWebhooksCallbacksStatus200ResponseWebhook{}

// GetOrganizationWebhooksCallbacksStatus200ResponseWebhook The webhook receiver used by the callback to send results
type GetOrganizationWebhooksCallbacksStatus200ResponseWebhook struct {
	// The webhook receiver URL where the callback will be sent
	Url *string `json:"url,omitempty"`
	HttpServer *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer `json:"httpServer,omitempty"`
	PayloadTemplate *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate `json:"payloadTemplate,omitempty"`
	// The timestamp the callback was sent to the webhook receiver
	SentAt *time.Time `json:"sentAt,omitempty"`
}

// NewGetOrganizationWebhooksCallbacksStatus200ResponseWebhook instantiates a new GetOrganizationWebhooksCallbacksStatus200ResponseWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWebhooksCallbacksStatus200ResponseWebhook() *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook {
	this := GetOrganizationWebhooksCallbacksStatus200ResponseWebhook{}
	return &this
}

// NewGetOrganizationWebhooksCallbacksStatus200ResponseWebhookWithDefaults instantiates a new GetOrganizationWebhooksCallbacksStatus200ResponseWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWebhooksCallbacksStatus200ResponseWebhookWithDefaults() *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook {
	this := GetOrganizationWebhooksCallbacksStatus200ResponseWebhook{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) SetUrl(v string) {
	o.Url = &v
}

// GetHttpServer returns the HttpServer field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) GetHttpServer() GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer {
	if o == nil || IsNil(o.HttpServer) {
		var ret GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer
		return ret
	}
	return *o.HttpServer
}

// GetHttpServerOk returns a tuple with the HttpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) GetHttpServerOk() (*GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer, bool) {
	if o == nil || IsNil(o.HttpServer) {
		return nil, false
	}
	return o.HttpServer, true
}

// HasHttpServer returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) HasHttpServer() bool {
	if o != nil && !IsNil(o.HttpServer) {
		return true
	}

	return false
}

// SetHttpServer gets a reference to the given GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer and assigns it to the HttpServer field.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) SetHttpServer(v GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer) {
	o.HttpServer = &v
}

// GetPayloadTemplate returns the PayloadTemplate field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) GetPayloadTemplate() GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate {
	if o == nil || IsNil(o.PayloadTemplate) {
		var ret GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate
		return ret
	}
	return *o.PayloadTemplate
}

// GetPayloadTemplateOk returns a tuple with the PayloadTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) GetPayloadTemplateOk() (*GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate, bool) {
	if o == nil || IsNil(o.PayloadTemplate) {
		return nil, false
	}
	return o.PayloadTemplate, true
}

// HasPayloadTemplate returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) HasPayloadTemplate() bool {
	if o != nil && !IsNil(o.PayloadTemplate) {
		return true
	}

	return false
}

// SetPayloadTemplate gets a reference to the given GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate and assigns it to the PayloadTemplate field.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) SetPayloadTemplate(v GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate) {
	o.PayloadTemplate = &v
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) GetSentAt() time.Time {
	if o == nil || IsNil(o.SentAt) {
		var ret time.Time
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) GetSentAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SentAt) {
		return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) HasSentAt() bool {
	if o != nil && !IsNil(o.SentAt) {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given time.Time and assigns it to the SentAt field.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) SetSentAt(v time.Time) {
	o.SentAt = &v
}

func (o GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.HttpServer) {
		toSerialize["httpServer"] = o.HttpServer
	}
	if !IsNil(o.PayloadTemplate) {
		toSerialize["payloadTemplate"] = o.PayloadTemplate
	}
	if !IsNil(o.SentAt) {
		toSerialize["sentAt"] = o.SentAt
	}
	return toSerialize, nil
}

type NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhook struct {
	value *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook
	isSet bool
}

func (v NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhook) Get() *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook {
	return v.value
}

func (v *NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhook) Set(val *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhook(val *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) *NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhook {
	return &NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhook{value: val, isSet: true}
}

func (v NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


