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

// checks if the GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer{}

// GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer The webhook receiver used for the callback webhook
type GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer struct {
	// The webhook receiver ID that will receive information
	Id *string `json:"id,omitempty"`
}

// NewGetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer instantiates a new GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer() *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer {
	this := GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer{}
	return &this
}

// NewGetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServerWithDefaults instantiates a new GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServerWithDefaults() *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer {
	this := GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer) SetId(v string) {
	o.Id = &v
}

func (o GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer struct {
	value *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer
	isSet bool
}

func (v NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer) Get() *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer {
	return v.value
}

func (v *NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer) Set(val *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer(val *GetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer) *NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer {
	return &NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer{value: val, isSet: true}
}

func (v NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWebhooksCallbacksStatus200ResponseWebhookHttpServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


