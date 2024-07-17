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

// checks if the DevicesSerialLiveToolsThroughputTestPostRequestMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevicesSerialLiveToolsThroughputTestPostRequestMessage{}

// DevicesSerialLiveToolsThroughputTestPostRequestMessage struct for DevicesSerialLiveToolsThroughputTestPostRequestMessage
type DevicesSerialLiveToolsThroughputTestPostRequestMessage struct {
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
}

// NewDevicesSerialLiveToolsThroughputTestPostRequestMessage instantiates a new DevicesSerialLiveToolsThroughputTestPostRequestMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialLiveToolsThroughputTestPostRequestMessage() *DevicesSerialLiveToolsThroughputTestPostRequestMessage {
	this := DevicesSerialLiveToolsThroughputTestPostRequestMessage{}
	return &this
}

// NewDevicesSerialLiveToolsThroughputTestPostRequestMessageWithDefaults instantiates a new DevicesSerialLiveToolsThroughputTestPostRequestMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialLiveToolsThroughputTestPostRequestMessageWithDefaults() *DevicesSerialLiveToolsThroughputTestPostRequestMessage {
	this := DevicesSerialLiveToolsThroughputTestPostRequestMessage{}
	return &this
}

// GetThroughputTestId returns the ThroughputTestId field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetThroughputTestId() string {
	if o == nil || IsNil(o.ThroughputTestId) {
		var ret string
		return ret
	}
	return *o.ThroughputTestId
}

// GetThroughputTestIdOk returns a tuple with the ThroughputTestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetThroughputTestIdOk() (*string, bool) {
	if o == nil || IsNil(o.ThroughputTestId) {
		return nil, false
	}
	return o.ThroughputTestId, true
}

// HasThroughputTestId returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) HasThroughputTestId() bool {
	if o != nil && !IsNil(o.ThroughputTestId) {
		return true
	}

	return false
}

// SetThroughputTestId gets a reference to the given string and assigns it to the ThroughputTestId field.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) SetThroughputTestId(v string) {
	o.ThroughputTestId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) SetUrl(v string) {
	o.Url = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) SetStatus(v string) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetResult() CreateDeviceLiveToolsThroughputTest201ResponseResult {
	if o == nil || IsNil(o.Result) {
		var ret CreateDeviceLiveToolsThroughputTest201ResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetResultOk() (*CreateDeviceLiveToolsThroughputTest201ResponseResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given CreateDeviceLiveToolsThroughputTest201ResponseResult and assigns it to the Result field.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) SetResult(v CreateDeviceLiveToolsThroughputTest201ResponseResult) {
	o.Result = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetRequest() CreateDeviceLiveToolsThroughputTest201ResponseRequest {
	if o == nil || IsNil(o.Request) {
		var ret CreateDeviceLiveToolsThroughputTest201ResponseRequest
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetRequestOk() (*CreateDeviceLiveToolsThroughputTest201ResponseRequest, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given CreateDeviceLiveToolsThroughputTest201ResponseRequest and assigns it to the Request field.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) SetRequest(v CreateDeviceLiveToolsThroughputTest201ResponseRequest) {
	o.Request = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) SetError(v string) {
	o.Error = &v
}

func (o DevicesSerialLiveToolsThroughputTestPostRequestMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevicesSerialLiveToolsThroughputTestPostRequestMessage) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableDevicesSerialLiveToolsThroughputTestPostRequestMessage struct {
	value *DevicesSerialLiveToolsThroughputTestPostRequestMessage
	isSet bool
}

func (v NullableDevicesSerialLiveToolsThroughputTestPostRequestMessage) Get() *DevicesSerialLiveToolsThroughputTestPostRequestMessage {
	return v.value
}

func (v *NullableDevicesSerialLiveToolsThroughputTestPostRequestMessage) Set(val *DevicesSerialLiveToolsThroughputTestPostRequestMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialLiveToolsThroughputTestPostRequestMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialLiveToolsThroughputTestPostRequestMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialLiveToolsThroughputTestPostRequestMessage(val *DevicesSerialLiveToolsThroughputTestPostRequestMessage) *NullableDevicesSerialLiveToolsThroughputTestPostRequestMessage {
	return &NullableDevicesSerialLiveToolsThroughputTestPostRequestMessage{value: val, isSet: true}
}

func (v NullableDevicesSerialLiveToolsThroughputTestPostRequestMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialLiveToolsThroughputTestPostRequestMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

