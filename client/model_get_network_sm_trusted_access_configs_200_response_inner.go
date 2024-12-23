/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetNetworkSmTrustedAccessConfigs200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSmTrustedAccessConfigs200ResponseInner{}

// GetNetworkSmTrustedAccessConfigs200ResponseInner struct for GetNetworkSmTrustedAccessConfigs200ResponseInner
type GetNetworkSmTrustedAccessConfigs200ResponseInner struct {
	// device ID
	Id *string `json:"id,omitempty"`
	// SSID name
	SsidName *string `json:"ssidName,omitempty"`
	// device name
	Name *string `json:"name,omitempty"`
	// scope
	Scope *string `json:"scope,omitempty"`
	// device tags
	Tags []string `json:"tags,omitempty"`
	// type of access period, either a static range or a dynamic period
	TimeboundType *string `json:"timeboundType,omitempty"`
	// Send Email Notifications
	SendExpirationEmails *bool `json:"sendExpirationEmails,omitempty"`
	// Time before access expiration reminder email sends
	NotifyTimeBeforeAccessEnds *int32 `json:"notifyTimeBeforeAccessEnds,omitempty"`
	// Optional email text
	AdditionalEmailText *string `json:"additionalEmailText,omitempty"`
	// time that access starts
	AccessStartAt *time.Time `json:"accessStartAt,omitempty"`
	// time that access ends
	AccessEndAt *time.Time `json:"accessEndAt,omitempty"`
}

// NewGetNetworkSmTrustedAccessConfigs200ResponseInner instantiates a new GetNetworkSmTrustedAccessConfigs200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSmTrustedAccessConfigs200ResponseInner() *GetNetworkSmTrustedAccessConfigs200ResponseInner {
	this := GetNetworkSmTrustedAccessConfigs200ResponseInner{}
	return &this
}

// NewGetNetworkSmTrustedAccessConfigs200ResponseInnerWithDefaults instantiates a new GetNetworkSmTrustedAccessConfigs200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSmTrustedAccessConfigs200ResponseInnerWithDefaults() *GetNetworkSmTrustedAccessConfigs200ResponseInner {
	this := GetNetworkSmTrustedAccessConfigs200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetSsidName returns the SsidName field value if set, zero value otherwise.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetSsidName() string {
	if o == nil || IsNil(o.SsidName) {
		var ret string
		return ret
	}
	return *o.SsidName
}

// GetSsidNameOk returns a tuple with the SsidName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetSsidNameOk() (*string, bool) {
	if o == nil || IsNil(o.SsidName) {
		return nil, false
	}
	return o.SsidName, true
}

// HasSsidName returns a boolean if a field has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) HasSsidName() bool {
	if o != nil && !IsNil(o.SsidName) {
		return true
	}

	return false
}

// SetSsidName gets a reference to the given string and assigns it to the SsidName field.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) SetSsidName(v string) {
	o.SsidName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) SetScope(v string) {
	o.Scope = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) SetTags(v []string) {
	o.Tags = v
}

// GetTimeboundType returns the TimeboundType field value if set, zero value otherwise.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetTimeboundType() string {
	if o == nil || IsNil(o.TimeboundType) {
		var ret string
		return ret
	}
	return *o.TimeboundType
}

// GetTimeboundTypeOk returns a tuple with the TimeboundType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetTimeboundTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TimeboundType) {
		return nil, false
	}
	return o.TimeboundType, true
}

// HasTimeboundType returns a boolean if a field has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) HasTimeboundType() bool {
	if o != nil && !IsNil(o.TimeboundType) {
		return true
	}

	return false
}

// SetTimeboundType gets a reference to the given string and assigns it to the TimeboundType field.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) SetTimeboundType(v string) {
	o.TimeboundType = &v
}

// GetSendExpirationEmails returns the SendExpirationEmails field value if set, zero value otherwise.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetSendExpirationEmails() bool {
	if o == nil || IsNil(o.SendExpirationEmails) {
		var ret bool
		return ret
	}
	return *o.SendExpirationEmails
}

// GetSendExpirationEmailsOk returns a tuple with the SendExpirationEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetSendExpirationEmailsOk() (*bool, bool) {
	if o == nil || IsNil(o.SendExpirationEmails) {
		return nil, false
	}
	return o.SendExpirationEmails, true
}

// HasSendExpirationEmails returns a boolean if a field has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) HasSendExpirationEmails() bool {
	if o != nil && !IsNil(o.SendExpirationEmails) {
		return true
	}

	return false
}

// SetSendExpirationEmails gets a reference to the given bool and assigns it to the SendExpirationEmails field.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) SetSendExpirationEmails(v bool) {
	o.SendExpirationEmails = &v
}

// GetNotifyTimeBeforeAccessEnds returns the NotifyTimeBeforeAccessEnds field value if set, zero value otherwise.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetNotifyTimeBeforeAccessEnds() int32 {
	if o == nil || IsNil(o.NotifyTimeBeforeAccessEnds) {
		var ret int32
		return ret
	}
	return *o.NotifyTimeBeforeAccessEnds
}

// GetNotifyTimeBeforeAccessEndsOk returns a tuple with the NotifyTimeBeforeAccessEnds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetNotifyTimeBeforeAccessEndsOk() (*int32, bool) {
	if o == nil || IsNil(o.NotifyTimeBeforeAccessEnds) {
		return nil, false
	}
	return o.NotifyTimeBeforeAccessEnds, true
}

// HasNotifyTimeBeforeAccessEnds returns a boolean if a field has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) HasNotifyTimeBeforeAccessEnds() bool {
	if o != nil && !IsNil(o.NotifyTimeBeforeAccessEnds) {
		return true
	}

	return false
}

// SetNotifyTimeBeforeAccessEnds gets a reference to the given int32 and assigns it to the NotifyTimeBeforeAccessEnds field.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) SetNotifyTimeBeforeAccessEnds(v int32) {
	o.NotifyTimeBeforeAccessEnds = &v
}

// GetAdditionalEmailText returns the AdditionalEmailText field value if set, zero value otherwise.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetAdditionalEmailText() string {
	if o == nil || IsNil(o.AdditionalEmailText) {
		var ret string
		return ret
	}
	return *o.AdditionalEmailText
}

// GetAdditionalEmailTextOk returns a tuple with the AdditionalEmailText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetAdditionalEmailTextOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalEmailText) {
		return nil, false
	}
	return o.AdditionalEmailText, true
}

// HasAdditionalEmailText returns a boolean if a field has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) HasAdditionalEmailText() bool {
	if o != nil && !IsNil(o.AdditionalEmailText) {
		return true
	}

	return false
}

// SetAdditionalEmailText gets a reference to the given string and assigns it to the AdditionalEmailText field.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) SetAdditionalEmailText(v string) {
	o.AdditionalEmailText = &v
}

// GetAccessStartAt returns the AccessStartAt field value if set, zero value otherwise.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetAccessStartAt() time.Time {
	if o == nil || IsNil(o.AccessStartAt) {
		var ret time.Time
		return ret
	}
	return *o.AccessStartAt
}

// GetAccessStartAtOk returns a tuple with the AccessStartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetAccessStartAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AccessStartAt) {
		return nil, false
	}
	return o.AccessStartAt, true
}

// HasAccessStartAt returns a boolean if a field has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) HasAccessStartAt() bool {
	if o != nil && !IsNil(o.AccessStartAt) {
		return true
	}

	return false
}

// SetAccessStartAt gets a reference to the given time.Time and assigns it to the AccessStartAt field.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) SetAccessStartAt(v time.Time) {
	o.AccessStartAt = &v
}

// GetAccessEndAt returns the AccessEndAt field value if set, zero value otherwise.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetAccessEndAt() time.Time {
	if o == nil || IsNil(o.AccessEndAt) {
		var ret time.Time
		return ret
	}
	return *o.AccessEndAt
}

// GetAccessEndAtOk returns a tuple with the AccessEndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) GetAccessEndAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AccessEndAt) {
		return nil, false
	}
	return o.AccessEndAt, true
}

// HasAccessEndAt returns a boolean if a field has been set.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) HasAccessEndAt() bool {
	if o != nil && !IsNil(o.AccessEndAt) {
		return true
	}

	return false
}

// SetAccessEndAt gets a reference to the given time.Time and assigns it to the AccessEndAt field.
func (o *GetNetworkSmTrustedAccessConfigs200ResponseInner) SetAccessEndAt(v time.Time) {
	o.AccessEndAt = &v
}

func (o GetNetworkSmTrustedAccessConfigs200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSmTrustedAccessConfigs200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SsidName) {
		toSerialize["ssidName"] = o.SsidName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.TimeboundType) {
		toSerialize["timeboundType"] = o.TimeboundType
	}
	if !IsNil(o.SendExpirationEmails) {
		toSerialize["sendExpirationEmails"] = o.SendExpirationEmails
	}
	if !IsNil(o.NotifyTimeBeforeAccessEnds) {
		toSerialize["notifyTimeBeforeAccessEnds"] = o.NotifyTimeBeforeAccessEnds
	}
	if !IsNil(o.AdditionalEmailText) {
		toSerialize["additionalEmailText"] = o.AdditionalEmailText
	}
	if !IsNil(o.AccessStartAt) {
		toSerialize["accessStartAt"] = o.AccessStartAt
	}
	if !IsNil(o.AccessEndAt) {
		toSerialize["accessEndAt"] = o.AccessEndAt
	}
	return toSerialize, nil
}

type NullableGetNetworkSmTrustedAccessConfigs200ResponseInner struct {
	value *GetNetworkSmTrustedAccessConfigs200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSmTrustedAccessConfigs200ResponseInner) Get() *GetNetworkSmTrustedAccessConfigs200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSmTrustedAccessConfigs200ResponseInner) Set(val *GetNetworkSmTrustedAccessConfigs200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSmTrustedAccessConfigs200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSmTrustedAccessConfigs200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSmTrustedAccessConfigs200ResponseInner(val *GetNetworkSmTrustedAccessConfigs200ResponseInner) *NullableGetNetworkSmTrustedAccessConfigs200ResponseInner {
	return &NullableGetNetworkSmTrustedAccessConfigs200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSmTrustedAccessConfigs200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSmTrustedAccessConfigs200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


