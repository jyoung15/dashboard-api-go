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

// checks if the GetNetworkEvents200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkEvents200Response{}

// GetNetworkEvents200Response struct for GetNetworkEvents200Response
type GetNetworkEvents200Response struct {
	// A message regarding the events sent. Usually 'null' unless there are no events
	Message *string `json:"message,omitempty"`
	// An UTC ISO8601 string of the earliest occured at time of the listed events of the page.
	PageStartAt *string `json:"pageStartAt,omitempty"`
	// An UTC ISO8601 string of the latest occured at time of the listed events of the page.
	PageEndAt *string `json:"pageEndAt,omitempty"`
	// An array of events that took place in the network.
	Events []GetNetworkEvents200ResponseEventsInner `json:"events,omitempty"`
}

// NewGetNetworkEvents200Response instantiates a new GetNetworkEvents200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkEvents200Response() *GetNetworkEvents200Response {
	this := GetNetworkEvents200Response{}
	return &this
}

// NewGetNetworkEvents200ResponseWithDefaults instantiates a new GetNetworkEvents200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkEvents200ResponseWithDefaults() *GetNetworkEvents200Response {
	this := GetNetworkEvents200Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetNetworkEvents200Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkEvents200Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetNetworkEvents200Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetNetworkEvents200Response) SetMessage(v string) {
	o.Message = &v
}

// GetPageStartAt returns the PageStartAt field value if set, zero value otherwise.
func (o *GetNetworkEvents200Response) GetPageStartAt() string {
	if o == nil || IsNil(o.PageStartAt) {
		var ret string
		return ret
	}
	return *o.PageStartAt
}

// GetPageStartAtOk returns a tuple with the PageStartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkEvents200Response) GetPageStartAtOk() (*string, bool) {
	if o == nil || IsNil(o.PageStartAt) {
		return nil, false
	}
	return o.PageStartAt, true
}

// HasPageStartAt returns a boolean if a field has been set.
func (o *GetNetworkEvents200Response) HasPageStartAt() bool {
	if o != nil && !IsNil(o.PageStartAt) {
		return true
	}

	return false
}

// SetPageStartAt gets a reference to the given string and assigns it to the PageStartAt field.
func (o *GetNetworkEvents200Response) SetPageStartAt(v string) {
	o.PageStartAt = &v
}

// GetPageEndAt returns the PageEndAt field value if set, zero value otherwise.
func (o *GetNetworkEvents200Response) GetPageEndAt() string {
	if o == nil || IsNil(o.PageEndAt) {
		var ret string
		return ret
	}
	return *o.PageEndAt
}

// GetPageEndAtOk returns a tuple with the PageEndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkEvents200Response) GetPageEndAtOk() (*string, bool) {
	if o == nil || IsNil(o.PageEndAt) {
		return nil, false
	}
	return o.PageEndAt, true
}

// HasPageEndAt returns a boolean if a field has been set.
func (o *GetNetworkEvents200Response) HasPageEndAt() bool {
	if o != nil && !IsNil(o.PageEndAt) {
		return true
	}

	return false
}

// SetPageEndAt gets a reference to the given string and assigns it to the PageEndAt field.
func (o *GetNetworkEvents200Response) SetPageEndAt(v string) {
	o.PageEndAt = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *GetNetworkEvents200Response) GetEvents() []GetNetworkEvents200ResponseEventsInner {
	if o == nil || IsNil(o.Events) {
		var ret []GetNetworkEvents200ResponseEventsInner
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkEvents200Response) GetEventsOk() ([]GetNetworkEvents200ResponseEventsInner, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *GetNetworkEvents200Response) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []GetNetworkEvents200ResponseEventsInner and assigns it to the Events field.
func (o *GetNetworkEvents200Response) SetEvents(v []GetNetworkEvents200ResponseEventsInner) {
	o.Events = v
}

func (o GetNetworkEvents200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkEvents200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.PageStartAt) {
		toSerialize["pageStartAt"] = o.PageStartAt
	}
	if !IsNil(o.PageEndAt) {
		toSerialize["pageEndAt"] = o.PageEndAt
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	return toSerialize, nil
}

type NullableGetNetworkEvents200Response struct {
	value *GetNetworkEvents200Response
	isSet bool
}

func (v NullableGetNetworkEvents200Response) Get() *GetNetworkEvents200Response {
	return v.value
}

func (v *NullableGetNetworkEvents200Response) Set(val *GetNetworkEvents200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkEvents200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkEvents200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkEvents200Response(val *GetNetworkEvents200Response) *NullableGetNetworkEvents200Response {
	return &NullableGetNetworkEvents200Response{value: val, isSet: true}
}

func (v NullableGetNetworkEvents200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkEvents200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


