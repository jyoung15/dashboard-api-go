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

// checks if the CreateNetworkWebhooksPayloadTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkWebhooksPayloadTemplateRequest{}

// CreateNetworkWebhooksPayloadTemplateRequest struct for CreateNetworkWebhooksPayloadTemplateRequest
type CreateNetworkWebhooksPayloadTemplateRequest struct {
	// The name of the new template
	Name string `json:"name"`
	// The liquid template used for the body of the webhook message. Either `body` or `bodyFile` must be specified.
	Body *string `json:"body,omitempty"`
	// The liquid template used with the webhook headers.
	Headers []CreateNetworkWebhooksPayloadTemplateRequestHeadersInner `json:"headers,omitempty"`
	// A file containing liquid template used for the body of the webhook message. Either `body` or `bodyFile` must be specified.
	BodyFile *string `json:"bodyFile,omitempty"`
	// A file containing the liquid template used with the webhook headers.
	HeadersFile *string `json:"headersFile,omitempty"`
}

type _CreateNetworkWebhooksPayloadTemplateRequest CreateNetworkWebhooksPayloadTemplateRequest

// NewCreateNetworkWebhooksPayloadTemplateRequest instantiates a new CreateNetworkWebhooksPayloadTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWebhooksPayloadTemplateRequest(name string) *CreateNetworkWebhooksPayloadTemplateRequest {
	this := CreateNetworkWebhooksPayloadTemplateRequest{}
	this.Name = name
	return &this
}

// NewCreateNetworkWebhooksPayloadTemplateRequestWithDefaults instantiates a new CreateNetworkWebhooksPayloadTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWebhooksPayloadTemplateRequestWithDefaults() *CreateNetworkWebhooksPayloadTemplateRequest {
	this := CreateNetworkWebhooksPayloadTemplateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNetworkWebhooksPayloadTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkWebhooksPayloadTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) SetBody(v string) {
	o.Body = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) GetHeaders() []CreateNetworkWebhooksPayloadTemplateRequestHeadersInner {
	if o == nil || IsNil(o.Headers) {
		var ret []CreateNetworkWebhooksPayloadTemplateRequestHeadersInner
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) GetHeadersOk() ([]CreateNetworkWebhooksPayloadTemplateRequestHeadersInner, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []CreateNetworkWebhooksPayloadTemplateRequestHeadersInner and assigns it to the Headers field.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) SetHeaders(v []CreateNetworkWebhooksPayloadTemplateRequestHeadersInner) {
	o.Headers = v
}

// GetBodyFile returns the BodyFile field value if set, zero value otherwise.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) GetBodyFile() string {
	if o == nil || IsNil(o.BodyFile) {
		var ret string
		return ret
	}
	return *o.BodyFile
}

// GetBodyFileOk returns a tuple with the BodyFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) GetBodyFileOk() (*string, bool) {
	if o == nil || IsNil(o.BodyFile) {
		return nil, false
	}
	return o.BodyFile, true
}

// HasBodyFile returns a boolean if a field has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) HasBodyFile() bool {
	if o != nil && !IsNil(o.BodyFile) {
		return true
	}

	return false
}

// SetBodyFile gets a reference to the given string and assigns it to the BodyFile field.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) SetBodyFile(v string) {
	o.BodyFile = &v
}

// GetHeadersFile returns the HeadersFile field value if set, zero value otherwise.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) GetHeadersFile() string {
	if o == nil || IsNil(o.HeadersFile) {
		var ret string
		return ret
	}
	return *o.HeadersFile
}

// GetHeadersFileOk returns a tuple with the HeadersFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) GetHeadersFileOk() (*string, bool) {
	if o == nil || IsNil(o.HeadersFile) {
		return nil, false
	}
	return o.HeadersFile, true
}

// HasHeadersFile returns a boolean if a field has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) HasHeadersFile() bool {
	if o != nil && !IsNil(o.HeadersFile) {
		return true
	}

	return false
}

// SetHeadersFile gets a reference to the given string and assigns it to the HeadersFile field.
func (o *CreateNetworkWebhooksPayloadTemplateRequest) SetHeadersFile(v string) {
	o.HeadersFile = &v
}

func (o CreateNetworkWebhooksPayloadTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkWebhooksPayloadTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.BodyFile) {
		toSerialize["bodyFile"] = o.BodyFile
	}
	if !IsNil(o.HeadersFile) {
		toSerialize["headersFile"] = o.HeadersFile
	}
	return toSerialize, nil
}

func (o *CreateNetworkWebhooksPayloadTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCreateNetworkWebhooksPayloadTemplateRequest := _CreateNetworkWebhooksPayloadTemplateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNetworkWebhooksPayloadTemplateRequest)

	if err != nil {
		return err
	}

	*o = CreateNetworkWebhooksPayloadTemplateRequest(varCreateNetworkWebhooksPayloadTemplateRequest)

	return err
}

type NullableCreateNetworkWebhooksPayloadTemplateRequest struct {
	value *CreateNetworkWebhooksPayloadTemplateRequest
	isSet bool
}

func (v NullableCreateNetworkWebhooksPayloadTemplateRequest) Get() *CreateNetworkWebhooksPayloadTemplateRequest {
	return v.value
}

func (v *NullableCreateNetworkWebhooksPayloadTemplateRequest) Set(val *CreateNetworkWebhooksPayloadTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWebhooksPayloadTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWebhooksPayloadTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWebhooksPayloadTemplateRequest(val *CreateNetworkWebhooksPayloadTemplateRequest) *NullableCreateNetworkWebhooksPayloadTemplateRequest {
	return &NullableCreateNetworkWebhooksPayloadTemplateRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkWebhooksPayloadTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWebhooksPayloadTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


