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

// checks if the CreateNetworkWebhooksHttpServerRequestPayloadTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkWebhooksHttpServerRequestPayloadTemplate{}

// CreateNetworkWebhooksHttpServerRequestPayloadTemplate The payload template to use when posting data to the HTTP server.
type CreateNetworkWebhooksHttpServerRequestPayloadTemplate struct {
	// The ID of the payload template. Defaults to 'wpt_00001' for the Meraki template. For Meraki-included templates: for the Webex (included) template use 'wpt_00002'; for the Slack (included) template use 'wpt_00003'; for the Microsoft Teams (included) template use 'wpt_00004'; for the ServiceNow (included) template use 'wpt_00006'
	PayloadTemplateId *string `json:"payloadTemplateId,omitempty"`
	// The name of the payload template.
	Name *string `json:"name,omitempty"`
}

// NewCreateNetworkWebhooksHttpServerRequestPayloadTemplate instantiates a new CreateNetworkWebhooksHttpServerRequestPayloadTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWebhooksHttpServerRequestPayloadTemplate() *CreateNetworkWebhooksHttpServerRequestPayloadTemplate {
	this := CreateNetworkWebhooksHttpServerRequestPayloadTemplate{}
	return &this
}

// NewCreateNetworkWebhooksHttpServerRequestPayloadTemplateWithDefaults instantiates a new CreateNetworkWebhooksHttpServerRequestPayloadTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWebhooksHttpServerRequestPayloadTemplateWithDefaults() *CreateNetworkWebhooksHttpServerRequestPayloadTemplate {
	this := CreateNetworkWebhooksHttpServerRequestPayloadTemplate{}
	return &this
}

// GetPayloadTemplateId returns the PayloadTemplateId field value if set, zero value otherwise.
func (o *CreateNetworkWebhooksHttpServerRequestPayloadTemplate) GetPayloadTemplateId() string {
	if o == nil || IsNil(o.PayloadTemplateId) {
		var ret string
		return ret
	}
	return *o.PayloadTemplateId
}

// GetPayloadTemplateIdOk returns a tuple with the PayloadTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksHttpServerRequestPayloadTemplate) GetPayloadTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayloadTemplateId) {
		return nil, false
	}
	return o.PayloadTemplateId, true
}

// HasPayloadTemplateId returns a boolean if a field has been set.
func (o *CreateNetworkWebhooksHttpServerRequestPayloadTemplate) HasPayloadTemplateId() bool {
	if o != nil && !IsNil(o.PayloadTemplateId) {
		return true
	}

	return false
}

// SetPayloadTemplateId gets a reference to the given string and assigns it to the PayloadTemplateId field.
func (o *CreateNetworkWebhooksHttpServerRequestPayloadTemplate) SetPayloadTemplateId(v string) {
	o.PayloadTemplateId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateNetworkWebhooksHttpServerRequestPayloadTemplate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksHttpServerRequestPayloadTemplate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateNetworkWebhooksHttpServerRequestPayloadTemplate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateNetworkWebhooksHttpServerRequestPayloadTemplate) SetName(v string) {
	o.Name = &v
}

func (o CreateNetworkWebhooksHttpServerRequestPayloadTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkWebhooksHttpServerRequestPayloadTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PayloadTemplateId) {
		toSerialize["payloadTemplateId"] = o.PayloadTemplateId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCreateNetworkWebhooksHttpServerRequestPayloadTemplate struct {
	value *CreateNetworkWebhooksHttpServerRequestPayloadTemplate
	isSet bool
}

func (v NullableCreateNetworkWebhooksHttpServerRequestPayloadTemplate) Get() *CreateNetworkWebhooksHttpServerRequestPayloadTemplate {
	return v.value
}

func (v *NullableCreateNetworkWebhooksHttpServerRequestPayloadTemplate) Set(val *CreateNetworkWebhooksHttpServerRequestPayloadTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWebhooksHttpServerRequestPayloadTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWebhooksHttpServerRequestPayloadTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWebhooksHttpServerRequestPayloadTemplate(val *CreateNetworkWebhooksHttpServerRequestPayloadTemplate) *NullableCreateNetworkWebhooksHttpServerRequestPayloadTemplate {
	return &NullableCreateNetworkWebhooksHttpServerRequestPayloadTemplate{value: val, isSet: true}
}

func (v NullableCreateNetworkWebhooksHttpServerRequestPayloadTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWebhooksHttpServerRequestPayloadTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


