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

// checks if the GetDeviceLiveToolsPingDevice200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceLiveToolsPingDevice200Response{}

// GetDeviceLiveToolsPingDevice200Response struct for GetDeviceLiveToolsPingDevice200Response
type GetDeviceLiveToolsPingDevice200Response struct {
	// Id to check the status of your ping request.
	PingId *string `json:"pingId,omitempty"`
	// GET this url to check the status of your ping request.
	Url *string `json:"url,omitempty"`
	Request *CreateDeviceLiveToolsPing201ResponseRequest `json:"request,omitempty"`
	// Status of the ping request.
	Status *string `json:"status,omitempty"`
	Results *DevicesSerialLiveToolsPingPostRequestMessageResults `json:"results,omitempty"`
	Callback *CreateDeviceLiveToolsArpTable201ResponseCallback `json:"callback,omitempty"`
}

// NewGetDeviceLiveToolsPingDevice200Response instantiates a new GetDeviceLiveToolsPingDevice200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceLiveToolsPingDevice200Response() *GetDeviceLiveToolsPingDevice200Response {
	this := GetDeviceLiveToolsPingDevice200Response{}
	return &this
}

// NewGetDeviceLiveToolsPingDevice200ResponseWithDefaults instantiates a new GetDeviceLiveToolsPingDevice200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceLiveToolsPingDevice200ResponseWithDefaults() *GetDeviceLiveToolsPingDevice200Response {
	this := GetDeviceLiveToolsPingDevice200Response{}
	return &this
}

// GetPingId returns the PingId field value if set, zero value otherwise.
func (o *GetDeviceLiveToolsPingDevice200Response) GetPingId() string {
	if o == nil || IsNil(o.PingId) {
		var ret string
		return ret
	}
	return *o.PingId
}

// GetPingIdOk returns a tuple with the PingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceLiveToolsPingDevice200Response) GetPingIdOk() (*string, bool) {
	if o == nil || IsNil(o.PingId) {
		return nil, false
	}
	return o.PingId, true
}

// HasPingId returns a boolean if a field has been set.
func (o *GetDeviceLiveToolsPingDevice200Response) HasPingId() bool {
	if o != nil && !IsNil(o.PingId) {
		return true
	}

	return false
}

// SetPingId gets a reference to the given string and assigns it to the PingId field.
func (o *GetDeviceLiveToolsPingDevice200Response) SetPingId(v string) {
	o.PingId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetDeviceLiveToolsPingDevice200Response) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceLiveToolsPingDevice200Response) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetDeviceLiveToolsPingDevice200Response) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetDeviceLiveToolsPingDevice200Response) SetUrl(v string) {
	o.Url = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *GetDeviceLiveToolsPingDevice200Response) GetRequest() CreateDeviceLiveToolsPing201ResponseRequest {
	if o == nil || IsNil(o.Request) {
		var ret CreateDeviceLiveToolsPing201ResponseRequest
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceLiveToolsPingDevice200Response) GetRequestOk() (*CreateDeviceLiveToolsPing201ResponseRequest, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *GetDeviceLiveToolsPingDevice200Response) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given CreateDeviceLiveToolsPing201ResponseRequest and assigns it to the Request field.
func (o *GetDeviceLiveToolsPingDevice200Response) SetRequest(v CreateDeviceLiveToolsPing201ResponseRequest) {
	o.Request = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetDeviceLiveToolsPingDevice200Response) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceLiveToolsPingDevice200Response) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetDeviceLiveToolsPingDevice200Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetDeviceLiveToolsPingDevice200Response) SetStatus(v string) {
	o.Status = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *GetDeviceLiveToolsPingDevice200Response) GetResults() DevicesSerialLiveToolsPingPostRequestMessageResults {
	if o == nil || IsNil(o.Results) {
		var ret DevicesSerialLiveToolsPingPostRequestMessageResults
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceLiveToolsPingDevice200Response) GetResultsOk() (*DevicesSerialLiveToolsPingPostRequestMessageResults, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *GetDeviceLiveToolsPingDevice200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given DevicesSerialLiveToolsPingPostRequestMessageResults and assigns it to the Results field.
func (o *GetDeviceLiveToolsPingDevice200Response) SetResults(v DevicesSerialLiveToolsPingPostRequestMessageResults) {
	o.Results = &v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *GetDeviceLiveToolsPingDevice200Response) GetCallback() CreateDeviceLiveToolsArpTable201ResponseCallback {
	if o == nil || IsNil(o.Callback) {
		var ret CreateDeviceLiveToolsArpTable201ResponseCallback
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceLiveToolsPingDevice200Response) GetCallbackOk() (*CreateDeviceLiveToolsArpTable201ResponseCallback, bool) {
	if o == nil || IsNil(o.Callback) {
		return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *GetDeviceLiveToolsPingDevice200Response) HasCallback() bool {
	if o != nil && !IsNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given CreateDeviceLiveToolsArpTable201ResponseCallback and assigns it to the Callback field.
func (o *GetDeviceLiveToolsPingDevice200Response) SetCallback(v CreateDeviceLiveToolsArpTable201ResponseCallback) {
	o.Callback = &v
}

func (o GetDeviceLiveToolsPingDevice200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceLiveToolsPingDevice200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PingId) {
		toSerialize["pingId"] = o.PingId
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
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.Callback) {
		toSerialize["callback"] = o.Callback
	}
	return toSerialize, nil
}

type NullableGetDeviceLiveToolsPingDevice200Response struct {
	value *GetDeviceLiveToolsPingDevice200Response
	isSet bool
}

func (v NullableGetDeviceLiveToolsPingDevice200Response) Get() *GetDeviceLiveToolsPingDevice200Response {
	return v.value
}

func (v *NullableGetDeviceLiveToolsPingDevice200Response) Set(val *GetDeviceLiveToolsPingDevice200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceLiveToolsPingDevice200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceLiveToolsPingDevice200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceLiveToolsPingDevice200Response(val *GetDeviceLiveToolsPingDevice200Response) *NullableGetDeviceLiveToolsPingDevice200Response {
	return &NullableGetDeviceLiveToolsPingDevice200Response{value: val, isSet: true}
}

func (v NullableGetDeviceLiveToolsPingDevice200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceLiveToolsPingDevice200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


