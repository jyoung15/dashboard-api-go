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

// checks if the GetNetworkSmUserAccessDevices200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSmUserAccessDevices200ResponseInner{}

// GetNetworkSmUserAccessDevices200ResponseInner struct for GetNetworkSmUserAccessDevices200ResponseInner
type GetNetworkSmUserAccessDevices200ResponseInner struct {
	// device ID
	Id *string `json:"id,omitempty"`
	// device name
	Name *string `json:"name,omitempty"`
	// system type
	SystemType *string `json:"systemType,omitempty"`
	// mac address
	Mac *string `json:"mac,omitempty"`
	// username
	Username *string `json:"username,omitempty"`
	// user email
	Email *string `json:"email,omitempty"`
	// device tags
	Tags []string `json:"tags,omitempty"`
	// Array of trusted access configs
	TrustedAccessConnections []GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner `json:"trustedAccessConnections,omitempty"`
}

// NewGetNetworkSmUserAccessDevices200ResponseInner instantiates a new GetNetworkSmUserAccessDevices200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSmUserAccessDevices200ResponseInner() *GetNetworkSmUserAccessDevices200ResponseInner {
	this := GetNetworkSmUserAccessDevices200ResponseInner{}
	return &this
}

// NewGetNetworkSmUserAccessDevices200ResponseInnerWithDefaults instantiates a new GetNetworkSmUserAccessDevices200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSmUserAccessDevices200ResponseInnerWithDefaults() *GetNetworkSmUserAccessDevices200ResponseInner {
	this := GetNetworkSmUserAccessDevices200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetSystemType returns the SystemType field value if set, zero value otherwise.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetSystemType() string {
	if o == nil || IsNil(o.SystemType) {
		var ret string
		return ret
	}
	return *o.SystemType
}

// GetSystemTypeOk returns a tuple with the SystemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetSystemTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SystemType) {
		return nil, false
	}
	return o.SystemType, true
}

// HasSystemType returns a boolean if a field has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) HasSystemType() bool {
	if o != nil && !IsNil(o.SystemType) {
		return true
	}

	return false
}

// SetSystemType gets a reference to the given string and assigns it to the SystemType field.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) SetSystemType(v string) {
	o.SystemType = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) SetMac(v string) {
	o.Mac = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) SetUsername(v string) {
	o.Username = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) SetEmail(v string) {
	o.Email = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) SetTags(v []string) {
	o.Tags = v
}

// GetTrustedAccessConnections returns the TrustedAccessConnections field value if set, zero value otherwise.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetTrustedAccessConnections() []GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner {
	if o == nil || IsNil(o.TrustedAccessConnections) {
		var ret []GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner
		return ret
	}
	return o.TrustedAccessConnections
}

// GetTrustedAccessConnectionsOk returns a tuple with the TrustedAccessConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) GetTrustedAccessConnectionsOk() ([]GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner, bool) {
	if o == nil || IsNil(o.TrustedAccessConnections) {
		return nil, false
	}
	return o.TrustedAccessConnections, true
}

// HasTrustedAccessConnections returns a boolean if a field has been set.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) HasTrustedAccessConnections() bool {
	if o != nil && !IsNil(o.TrustedAccessConnections) {
		return true
	}

	return false
}

// SetTrustedAccessConnections gets a reference to the given []GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner and assigns it to the TrustedAccessConnections field.
func (o *GetNetworkSmUserAccessDevices200ResponseInner) SetTrustedAccessConnections(v []GetNetworkSmUserAccessDevices200ResponseInnerTrustedAccessConnectionsInner) {
	o.TrustedAccessConnections = v
}

func (o GetNetworkSmUserAccessDevices200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSmUserAccessDevices200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SystemType) {
		toSerialize["systemType"] = o.SystemType
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.TrustedAccessConnections) {
		toSerialize["trustedAccessConnections"] = o.TrustedAccessConnections
	}
	return toSerialize, nil
}

type NullableGetNetworkSmUserAccessDevices200ResponseInner struct {
	value *GetNetworkSmUserAccessDevices200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSmUserAccessDevices200ResponseInner) Get() *GetNetworkSmUserAccessDevices200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSmUserAccessDevices200ResponseInner) Set(val *GetNetworkSmUserAccessDevices200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSmUserAccessDevices200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSmUserAccessDevices200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSmUserAccessDevices200ResponseInner(val *GetNetworkSmUserAccessDevices200ResponseInner) *NullableGetNetworkSmUserAccessDevices200ResponseInner {
	return &NullableGetNetworkSmUserAccessDevices200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSmUserAccessDevices200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSmUserAccessDevices200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


