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

// checks if the CreateDeviceLiveToolsArpTable201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDeviceLiveToolsArpTable201Response{}

// CreateDeviceLiveToolsArpTable201Response struct for CreateDeviceLiveToolsArpTable201Response
type CreateDeviceLiveToolsArpTable201Response struct {
	// Id of the ARP table request. Used to check the status of the request.
	ArpTableId *string `json:"arpTableId,omitempty"`
	// GET this url to check the status of your ARP table request.
	Url *string `json:"url,omitempty"`
	Request *CreateDeviceLiveToolsArpTable201ResponseRequest `json:"request,omitempty"`
	// Status of the ARP table request.
	Status *string `json:"status,omitempty"`
	Callback *CreateDeviceLiveToolsArpTable201ResponseCallback `json:"callback,omitempty"`
}

// NewCreateDeviceLiveToolsArpTable201Response instantiates a new CreateDeviceLiveToolsArpTable201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceLiveToolsArpTable201Response() *CreateDeviceLiveToolsArpTable201Response {
	this := CreateDeviceLiveToolsArpTable201Response{}
	return &this
}

// NewCreateDeviceLiveToolsArpTable201ResponseWithDefaults instantiates a new CreateDeviceLiveToolsArpTable201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceLiveToolsArpTable201ResponseWithDefaults() *CreateDeviceLiveToolsArpTable201Response {
	this := CreateDeviceLiveToolsArpTable201Response{}
	return &this
}

// GetArpTableId returns the ArpTableId field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsArpTable201Response) GetArpTableId() string {
	if o == nil || IsNil(o.ArpTableId) {
		var ret string
		return ret
	}
	return *o.ArpTableId
}

// GetArpTableIdOk returns a tuple with the ArpTableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsArpTable201Response) GetArpTableIdOk() (*string, bool) {
	if o == nil || IsNil(o.ArpTableId) {
		return nil, false
	}
	return o.ArpTableId, true
}

// HasArpTableId returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsArpTable201Response) HasArpTableId() bool {
	if o != nil && !IsNil(o.ArpTableId) {
		return true
	}

	return false
}

// SetArpTableId gets a reference to the given string and assigns it to the ArpTableId field.
func (o *CreateDeviceLiveToolsArpTable201Response) SetArpTableId(v string) {
	o.ArpTableId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsArpTable201Response) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsArpTable201Response) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsArpTable201Response) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CreateDeviceLiveToolsArpTable201Response) SetUrl(v string) {
	o.Url = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsArpTable201Response) GetRequest() CreateDeviceLiveToolsArpTable201ResponseRequest {
	if o == nil || IsNil(o.Request) {
		var ret CreateDeviceLiveToolsArpTable201ResponseRequest
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsArpTable201Response) GetRequestOk() (*CreateDeviceLiveToolsArpTable201ResponseRequest, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsArpTable201Response) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given CreateDeviceLiveToolsArpTable201ResponseRequest and assigns it to the Request field.
func (o *CreateDeviceLiveToolsArpTable201Response) SetRequest(v CreateDeviceLiveToolsArpTable201ResponseRequest) {
	o.Request = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsArpTable201Response) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsArpTable201Response) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsArpTable201Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateDeviceLiveToolsArpTable201Response) SetStatus(v string) {
	o.Status = &v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsArpTable201Response) GetCallback() CreateDeviceLiveToolsArpTable201ResponseCallback {
	if o == nil || IsNil(o.Callback) {
		var ret CreateDeviceLiveToolsArpTable201ResponseCallback
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsArpTable201Response) GetCallbackOk() (*CreateDeviceLiveToolsArpTable201ResponseCallback, bool) {
	if o == nil || IsNil(o.Callback) {
		return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsArpTable201Response) HasCallback() bool {
	if o != nil && !IsNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given CreateDeviceLiveToolsArpTable201ResponseCallback and assigns it to the Callback field.
func (o *CreateDeviceLiveToolsArpTable201Response) SetCallback(v CreateDeviceLiveToolsArpTable201ResponseCallback) {
	o.Callback = &v
}

func (o CreateDeviceLiveToolsArpTable201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDeviceLiveToolsArpTable201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ArpTableId) {
		toSerialize["arpTableId"] = o.ArpTableId
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Callback) {
		toSerialize["callback"] = o.Callback
	}
	return toSerialize, nil
}

type NullableCreateDeviceLiveToolsArpTable201Response struct {
	value *CreateDeviceLiveToolsArpTable201Response
	isSet bool
}

func (v NullableCreateDeviceLiveToolsArpTable201Response) Get() *CreateDeviceLiveToolsArpTable201Response {
	return v.value
}

func (v *NullableCreateDeviceLiveToolsArpTable201Response) Set(val *CreateDeviceLiveToolsArpTable201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceLiveToolsArpTable201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceLiveToolsArpTable201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceLiveToolsArpTable201Response(val *CreateDeviceLiveToolsArpTable201Response) *NullableCreateDeviceLiveToolsArpTable201Response {
	return &NullableCreateDeviceLiveToolsArpTable201Response{value: val, isSet: true}
}

func (v NullableCreateDeviceLiveToolsArpTable201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceLiveToolsArpTable201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


