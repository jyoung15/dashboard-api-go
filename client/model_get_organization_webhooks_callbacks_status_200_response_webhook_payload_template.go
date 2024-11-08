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

// checks if the GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate{}

// GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate The payload template of the webhook used for the callback
type GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate struct {
	// The ID of the payload template
	Id *string `json:"id,omitempty"`
}

// NewGetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate instantiates a new GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate() *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate {
	this := GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate{}
	return &this
}

// NewGetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplateWithDefaults instantiates a new GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplateWithDefaults() *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate {
	this := GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate) SetId(v string) {
	o.Id = &v
}

func (o GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate struct {
	value *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate
	isSet bool
}

func (v NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate) Get() *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate {
	return v.value
}

func (v *NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate) Set(val *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate(val *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate) *NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate {
	return &NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate{value: val, isSet: true}
}

func (v NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhookPayloadTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


