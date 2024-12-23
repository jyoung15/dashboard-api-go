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

// checks if the GetOrganizationWebhooksCallbacksStatus200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWebhooksCallbacksStatus200Response{}

// GetOrganizationWebhooksCallbacksStatus200Response struct for GetOrganizationWebhooksCallbacksStatus200Response
type GetOrganizationWebhooksCallbacksStatus200Response struct {
	// The ID of the callback
	CallbackId *string `json:"callbackId,omitempty"`
	// The status of the callback
	Status *string `json:"status,omitempty"`
	// The errors returned by the callback
	Errors []string `json:"errors,omitempty"`
	CreatedBy *GetOrganizationWebhooksCallbacksStatus200ResponseCreatedBy `json:"createdBy,omitempty"`
	Webhook *GetOrganizationWebhooksCallbacksStatus200ResponseWebhook `json:"webhook,omitempty"`
}

// NewGetOrganizationWebhooksCallbacksStatus200Response instantiates a new GetOrganizationWebhooksCallbacksStatus200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWebhooksCallbacksStatus200Response() *GetOrganizationWebhooksCallbacksStatus200Response {
	this := GetOrganizationWebhooksCallbacksStatus200Response{}
	return &this
}

// NewGetOrganizationWebhooksCallbacksStatus200ResponseWithDefaults instantiates a new GetOrganizationWebhooksCallbacksStatus200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWebhooksCallbacksStatus200ResponseWithDefaults() *GetOrganizationWebhooksCallbacksStatus200Response {
	this := GetOrganizationWebhooksCallbacksStatus200Response{}
	return &this
}

// GetCallbackId returns the CallbackId field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksCallbacksStatus200Response) GetCallbackId() string {
	if o == nil || IsNil(o.CallbackId) {
		var ret string
		return ret
	}
	return *o.CallbackId
}

// GetCallbackIdOk returns a tuple with the CallbackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200Response) GetCallbackIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackId) {
		return nil, false
	}
	return o.CallbackId, true
}

// HasCallbackId returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200Response) HasCallbackId() bool {
	if o != nil && !IsNil(o.CallbackId) {
		return true
	}

	return false
}

// SetCallbackId gets a reference to the given string and assigns it to the CallbackId field.
func (o *GetOrganizationWebhooksCallbacksStatus200Response) SetCallbackId(v string) {
	o.CallbackId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksCallbacksStatus200Response) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200Response) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetOrganizationWebhooksCallbacksStatus200Response) SetStatus(v string) {
	o.Status = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksCallbacksStatus200Response) GetErrors() []string {
	if o == nil || IsNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200Response) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200Response) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *GetOrganizationWebhooksCallbacksStatus200Response) SetErrors(v []string) {
	o.Errors = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksCallbacksStatus200Response) GetCreatedBy() GetOrganizationWebhooksCallbacksStatus200ResponseCreatedBy {
	if o == nil || IsNil(o.CreatedBy) {
		var ret GetOrganizationWebhooksCallbacksStatus200ResponseCreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200Response) GetCreatedByOk() (*GetOrganizationWebhooksCallbacksStatus200ResponseCreatedBy, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200Response) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given GetOrganizationWebhooksCallbacksStatus200ResponseCreatedBy and assigns it to the CreatedBy field.
func (o *GetOrganizationWebhooksCallbacksStatus200Response) SetCreatedBy(v GetOrganizationWebhooksCallbacksStatus200ResponseCreatedBy) {
	o.CreatedBy = &v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksCallbacksStatus200Response) GetWebhook() GetOrganizationWebhooksCallbacksStatus200ResponseWebhook {
	if o == nil || IsNil(o.Webhook) {
		var ret GetOrganizationWebhooksCallbacksStatus200ResponseWebhook
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200Response) GetWebhookOk() (*GetOrganizationWebhooksCallbacksStatus200ResponseWebhook, bool) {
	if o == nil || IsNil(o.Webhook) {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200Response) HasWebhook() bool {
	if o != nil && !IsNil(o.Webhook) {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given GetOrganizationWebhooksCallbacksStatus200ResponseWebhook and assigns it to the Webhook field.
func (o *GetOrganizationWebhooksCallbacksStatus200Response) SetWebhook(v GetOrganizationWebhooksCallbacksStatus200ResponseWebhook) {
	o.Webhook = &v
}

func (o GetOrganizationWebhooksCallbacksStatus200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWebhooksCallbacksStatus200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallbackId) {
		toSerialize["callbackId"] = o.CallbackId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.Webhook) {
		toSerialize["webhook"] = o.Webhook
	}
	return toSerialize, nil
}

type NullableGetOrganizationWebhooksCallbacksStatus200Response struct {
	value *GetOrganizationWebhooksCallbacksStatus200Response
	isSet bool
}

func (v NullableGetOrganizationWebhooksCallbacksStatus200Response) Get() *GetOrganizationWebhooksCallbacksStatus200Response {
	return v.value
}

func (v *NullableGetOrganizationWebhooksCallbacksStatus200Response) Set(val *GetOrganizationWebhooksCallbacksStatus200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWebhooksCallbacksStatus200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWebhooksCallbacksStatus200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWebhooksCallbacksStatus200Response(val *GetOrganizationWebhooksCallbacksStatus200Response) *NullableGetOrganizationWebhooksCallbacksStatus200Response {
	return &NullableGetOrganizationWebhooksCallbacksStatus200Response{value: val, isSet: true}
}

func (v NullableGetOrganizationWebhooksCallbacksStatus200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWebhooksCallbacksStatus200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


