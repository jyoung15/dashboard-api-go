/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateNetworkWebhooksWebhookTestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkWebhooksWebhookTestRequest{}

// CreateNetworkWebhooksWebhookTestRequest struct for CreateNetworkWebhooksWebhookTestRequest
type CreateNetworkWebhooksWebhookTestRequest struct {
	// The URL where the test webhook will be sent
	Url string `json:"url"`
	// The shared secret the test webhook will send. Optional. Defaults to an empty string.
	SharedSecret *string `json:"sharedSecret,omitempty"`
	// The ID of the payload template of the test webhook. Defaults to the HTTP server's template ID if one exists for the given URL, or Generic template ID otherwise
	PayloadTemplateId *string `json:"payloadTemplateId,omitempty"`
	// The name of the payload template.
	PayloadTemplateName *string `json:"payloadTemplateName,omitempty"`
	// The type of alert which the test webhook will send. Optional. Defaults to power_supply_down.
	AlertTypeId *string `json:"alertTypeId,omitempty"`
}

type _CreateNetworkWebhooksWebhookTestRequest CreateNetworkWebhooksWebhookTestRequest

// NewCreateNetworkWebhooksWebhookTestRequest instantiates a new CreateNetworkWebhooksWebhookTestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWebhooksWebhookTestRequest(url string) *CreateNetworkWebhooksWebhookTestRequest {
	this := CreateNetworkWebhooksWebhookTestRequest{}
	this.Url = url
	var sharedSecret string = ""
	this.SharedSecret = &sharedSecret
	var alertTypeId string = "power_supply_down"
	this.AlertTypeId = &alertTypeId
	return &this
}

// NewCreateNetworkWebhooksWebhookTestRequestWithDefaults instantiates a new CreateNetworkWebhooksWebhookTestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWebhooksWebhookTestRequestWithDefaults() *CreateNetworkWebhooksWebhookTestRequest {
	this := CreateNetworkWebhooksWebhookTestRequest{}
	var sharedSecret string = ""
	this.SharedSecret = &sharedSecret
	var alertTypeId string = "power_supply_down"
	this.AlertTypeId = &alertTypeId
	return &this
}

// GetUrl returns the Url field value
func (o *CreateNetworkWebhooksWebhookTestRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksWebhookTestRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CreateNetworkWebhooksWebhookTestRequest) SetUrl(v string) {
	o.Url = v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *CreateNetworkWebhooksWebhookTestRequest) GetSharedSecret() string {
	if o == nil || IsNil(o.SharedSecret) {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksWebhookTestRequest) GetSharedSecretOk() (*string, bool) {
	if o == nil || IsNil(o.SharedSecret) {
		return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *CreateNetworkWebhooksWebhookTestRequest) HasSharedSecret() bool {
	if o != nil && !IsNil(o.SharedSecret) {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *CreateNetworkWebhooksWebhookTestRequest) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetPayloadTemplateId returns the PayloadTemplateId field value if set, zero value otherwise.
func (o *CreateNetworkWebhooksWebhookTestRequest) GetPayloadTemplateId() string {
	if o == nil || IsNil(o.PayloadTemplateId) {
		var ret string
		return ret
	}
	return *o.PayloadTemplateId
}

// GetPayloadTemplateIdOk returns a tuple with the PayloadTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksWebhookTestRequest) GetPayloadTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.PayloadTemplateId) {
		return nil, false
	}
	return o.PayloadTemplateId, true
}

// HasPayloadTemplateId returns a boolean if a field has been set.
func (o *CreateNetworkWebhooksWebhookTestRequest) HasPayloadTemplateId() bool {
	if o != nil && !IsNil(o.PayloadTemplateId) {
		return true
	}

	return false
}

// SetPayloadTemplateId gets a reference to the given string and assigns it to the PayloadTemplateId field.
func (o *CreateNetworkWebhooksWebhookTestRequest) SetPayloadTemplateId(v string) {
	o.PayloadTemplateId = &v
}

// GetPayloadTemplateName returns the PayloadTemplateName field value if set, zero value otherwise.
func (o *CreateNetworkWebhooksWebhookTestRequest) GetPayloadTemplateName() string {
	if o == nil || IsNil(o.PayloadTemplateName) {
		var ret string
		return ret
	}
	return *o.PayloadTemplateName
}

// GetPayloadTemplateNameOk returns a tuple with the PayloadTemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksWebhookTestRequest) GetPayloadTemplateNameOk() (*string, bool) {
	if o == nil || IsNil(o.PayloadTemplateName) {
		return nil, false
	}
	return o.PayloadTemplateName, true
}

// HasPayloadTemplateName returns a boolean if a field has been set.
func (o *CreateNetworkWebhooksWebhookTestRequest) HasPayloadTemplateName() bool {
	if o != nil && !IsNil(o.PayloadTemplateName) {
		return true
	}

	return false
}

// SetPayloadTemplateName gets a reference to the given string and assigns it to the PayloadTemplateName field.
func (o *CreateNetworkWebhooksWebhookTestRequest) SetPayloadTemplateName(v string) {
	o.PayloadTemplateName = &v
}

// GetAlertTypeId returns the AlertTypeId field value if set, zero value otherwise.
func (o *CreateNetworkWebhooksWebhookTestRequest) GetAlertTypeId() string {
	if o == nil || IsNil(o.AlertTypeId) {
		var ret string
		return ret
	}
	return *o.AlertTypeId
}

// GetAlertTypeIdOk returns a tuple with the AlertTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksWebhookTestRequest) GetAlertTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlertTypeId) {
		return nil, false
	}
	return o.AlertTypeId, true
}

// HasAlertTypeId returns a boolean if a field has been set.
func (o *CreateNetworkWebhooksWebhookTestRequest) HasAlertTypeId() bool {
	if o != nil && !IsNil(o.AlertTypeId) {
		return true
	}

	return false
}

// SetAlertTypeId gets a reference to the given string and assigns it to the AlertTypeId field.
func (o *CreateNetworkWebhooksWebhookTestRequest) SetAlertTypeId(v string) {
	o.AlertTypeId = &v
}

func (o CreateNetworkWebhooksWebhookTestRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkWebhooksWebhookTestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	if !IsNil(o.SharedSecret) {
		toSerialize["sharedSecret"] = o.SharedSecret
	}
	if !IsNil(o.PayloadTemplateId) {
		toSerialize["payloadTemplateId"] = o.PayloadTemplateId
	}
	if !IsNil(o.PayloadTemplateName) {
		toSerialize["payloadTemplateName"] = o.PayloadTemplateName
	}
	if !IsNil(o.AlertTypeId) {
		toSerialize["alertTypeId"] = o.AlertTypeId
	}
	return toSerialize, nil
}

func (o *CreateNetworkWebhooksWebhookTestRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateNetworkWebhooksWebhookTestRequest := _CreateNetworkWebhooksWebhookTestRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNetworkWebhooksWebhookTestRequest)

	if err != nil {
		return err
	}

	*o = CreateNetworkWebhooksWebhookTestRequest(varCreateNetworkWebhooksWebhookTestRequest)

	return err
}

type NullableCreateNetworkWebhooksWebhookTestRequest struct {
	value *CreateNetworkWebhooksWebhookTestRequest
	isSet bool
}

func (v NullableCreateNetworkWebhooksWebhookTestRequest) Get() *CreateNetworkWebhooksWebhookTestRequest {
	return v.value
}

func (v *NullableCreateNetworkWebhooksWebhookTestRequest) Set(val *CreateNetworkWebhooksWebhookTestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWebhooksWebhookTestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWebhooksWebhookTestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWebhooksWebhookTestRequest(val *CreateNetworkWebhooksWebhookTestRequest) *NullableCreateNetworkWebhooksWebhookTestRequest {
	return &NullableCreateNetworkWebhooksWebhookTestRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkWebhooksWebhookTestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWebhooksWebhookTestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


