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

// checks if the CreateDeviceLiveToolsThroughputTest201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDeviceLiveToolsThroughputTest201Response{}

// CreateDeviceLiveToolsThroughputTest201Response struct for CreateDeviceLiveToolsThroughputTest201Response
type CreateDeviceLiveToolsThroughputTest201Response struct {
	// ID of throughput test job
	ThroughputTestId *string `json:"throughputTestId,omitempty"`
	// GET this url to check the status of your throughput test request
	Url *string `json:"url,omitempty"`
	// Status of the throughput test request
	Status *string `json:"status,omitempty"`
	Result *CreateDeviceLiveToolsThroughputTest201ResponseResult `json:"result,omitempty"`
	Request *CreateDeviceLiveToolsThroughputTest201ResponseRequest `json:"request,omitempty"`
	// Description of the error.
	Error *string `json:"error,omitempty"`
	Callback *CreateDeviceLiveToolsArpTable201ResponseCallback `json:"callback,omitempty"`
}

// NewCreateDeviceLiveToolsThroughputTest201Response instantiates a new CreateDeviceLiveToolsThroughputTest201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceLiveToolsThroughputTest201Response() *CreateDeviceLiveToolsThroughputTest201Response {
	this := CreateDeviceLiveToolsThroughputTest201Response{}
	return &this
}

// NewCreateDeviceLiveToolsThroughputTest201ResponseWithDefaults instantiates a new CreateDeviceLiveToolsThroughputTest201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceLiveToolsThroughputTest201ResponseWithDefaults() *CreateDeviceLiveToolsThroughputTest201Response {
	this := CreateDeviceLiveToolsThroughputTest201Response{}
	return &this
}

// GetThroughputTestId returns the ThroughputTestId field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsThroughputTest201Response) GetThroughputTestId() string {
	if o == nil || IsNil(o.ThroughputTestId) {
		var ret string
		return ret
	}
	return *o.ThroughputTestId
}

// GetThroughputTestIdOk returns a tuple with the ThroughputTestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsThroughputTest201Response) GetThroughputTestIdOk() (*string, bool) {
	if o == nil || IsNil(o.ThroughputTestId) {
		return nil, false
	}
	return o.ThroughputTestId, true
}

// HasThroughputTestId returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsThroughputTest201Response) HasThroughputTestId() bool {
	if o != nil && !IsNil(o.ThroughputTestId) {
		return true
	}

	return false
}

// SetThroughputTestId gets a reference to the given string and assigns it to the ThroughputTestId field.
func (o *CreateDeviceLiveToolsThroughputTest201Response) SetThroughputTestId(v string) {
	o.ThroughputTestId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsThroughputTest201Response) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsThroughputTest201Response) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsThroughputTest201Response) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CreateDeviceLiveToolsThroughputTest201Response) SetUrl(v string) {
	o.Url = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsThroughputTest201Response) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsThroughputTest201Response) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsThroughputTest201Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateDeviceLiveToolsThroughputTest201Response) SetStatus(v string) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsThroughputTest201Response) GetResult() CreateDeviceLiveToolsThroughputTest201ResponseResult {
	if o == nil || IsNil(o.Result) {
		var ret CreateDeviceLiveToolsThroughputTest201ResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsThroughputTest201Response) GetResultOk() (*CreateDeviceLiveToolsThroughputTest201ResponseResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsThroughputTest201Response) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given CreateDeviceLiveToolsThroughputTest201ResponseResult and assigns it to the Result field.
func (o *CreateDeviceLiveToolsThroughputTest201Response) SetResult(v CreateDeviceLiveToolsThroughputTest201ResponseResult) {
	o.Result = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsThroughputTest201Response) GetRequest() CreateDeviceLiveToolsThroughputTest201ResponseRequest {
	if o == nil || IsNil(o.Request) {
		var ret CreateDeviceLiveToolsThroughputTest201ResponseRequest
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsThroughputTest201Response) GetRequestOk() (*CreateDeviceLiveToolsThroughputTest201ResponseRequest, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsThroughputTest201Response) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given CreateDeviceLiveToolsThroughputTest201ResponseRequest and assigns it to the Request field.
func (o *CreateDeviceLiveToolsThroughputTest201Response) SetRequest(v CreateDeviceLiveToolsThroughputTest201ResponseRequest) {
	o.Request = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsThroughputTest201Response) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsThroughputTest201Response) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsThroughputTest201Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *CreateDeviceLiveToolsThroughputTest201Response) SetError(v string) {
	o.Error = &v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsThroughputTest201Response) GetCallback() CreateDeviceLiveToolsArpTable201ResponseCallback {
	if o == nil || IsNil(o.Callback) {
		var ret CreateDeviceLiveToolsArpTable201ResponseCallback
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsThroughputTest201Response) GetCallbackOk() (*CreateDeviceLiveToolsArpTable201ResponseCallback, bool) {
	if o == nil || IsNil(o.Callback) {
		return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsThroughputTest201Response) HasCallback() bool {
	if o != nil && !IsNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given CreateDeviceLiveToolsArpTable201ResponseCallback and assigns it to the Callback field.
func (o *CreateDeviceLiveToolsThroughputTest201Response) SetCallback(v CreateDeviceLiveToolsArpTable201ResponseCallback) {
	o.Callback = &v
}

func (o CreateDeviceLiveToolsThroughputTest201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDeviceLiveToolsThroughputTest201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ThroughputTestId) {
		toSerialize["throughputTestId"] = o.ThroughputTestId
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Callback) {
		toSerialize["callback"] = o.Callback
	}
	return toSerialize, nil
}

type NullableCreateDeviceLiveToolsThroughputTest201Response struct {
	value *CreateDeviceLiveToolsThroughputTest201Response
	isSet bool
}

func (v NullableCreateDeviceLiveToolsThroughputTest201Response) Get() *CreateDeviceLiveToolsThroughputTest201Response {
	return v.value
}

func (v *NullableCreateDeviceLiveToolsThroughputTest201Response) Set(val *CreateDeviceLiveToolsThroughputTest201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceLiveToolsThroughputTest201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceLiveToolsThroughputTest201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceLiveToolsThroughputTest201Response(val *CreateDeviceLiveToolsThroughputTest201Response) *NullableCreateDeviceLiveToolsThroughputTest201Response {
	return &NullableCreateDeviceLiveToolsThroughputTest201Response{value: val, isSet: true}
}

func (v NullableCreateDeviceLiveToolsThroughputTest201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceLiveToolsThroughputTest201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


