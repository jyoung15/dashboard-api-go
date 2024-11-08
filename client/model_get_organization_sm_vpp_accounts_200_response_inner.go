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

// checks if the GetOrganizationSmVppAccounts200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSmVppAccounts200ResponseInner{}

// GetOrganizationSmVppAccounts200ResponseInner struct for GetOrganizationSmVppAccounts200ResponseInner
type GetOrganizationSmVppAccounts200ResponseInner struct {
	// The id of the VPP Account
	VppAccountId *string `json:"vppAccountId,omitempty"`
	// The VPP service token
	ContentToken *string `json:"contentToken,omitempty"`
	// The email address associated with the VPP account
	Email *string `json:"email,omitempty"`
	// The name of the VPP account
	Name *string `json:"name,omitempty"`
	// The allowed admins for the VPP account
	AllowedAdmins *string `json:"allowedAdmins,omitempty"`
	// The network IDs of the admins for the VPP account
	NetworkIdAdmins *string `json:"networkIdAdmins,omitempty"`
	// The assignable networks for the VPP account
	AssignableNetworks *string `json:"assignableNetworks,omitempty"`
	// The network IDs of the assignable networks for the VPP account
	AssignableNetworkIds []string `json:"assignableNetworkIds,omitempty"`
	// The VPP location ID
	VppLocationId *string `json:"vppLocationId,omitempty"`
	// The VPP location name
	VppLocationName *string `json:"vppLocationName,omitempty"`
	// The last time the VPP account was synced
	LastSyncedAt *string `json:"lastSyncedAt,omitempty"`
	// The last time the VPP account was force synced
	LastForceSyncedAt *string `json:"lastForceSyncedAt,omitempty"`
	ParsedToken *GetOrganizationSmVppAccounts200ResponseInnerParsedToken `json:"parsedToken,omitempty"`
	// The id of the VPP Account
	Id *string `json:"id,omitempty"`
	// The VPP Account's Service Token
	VppServiceToken *string `json:"vppServiceToken,omitempty"`
}

// NewGetOrganizationSmVppAccounts200ResponseInner instantiates a new GetOrganizationSmVppAccounts200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSmVppAccounts200ResponseInner() *GetOrganizationSmVppAccounts200ResponseInner {
	this := GetOrganizationSmVppAccounts200ResponseInner{}
	return &this
}

// NewGetOrganizationSmVppAccounts200ResponseInnerWithDefaults instantiates a new GetOrganizationSmVppAccounts200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSmVppAccounts200ResponseInnerWithDefaults() *GetOrganizationSmVppAccounts200ResponseInner {
	this := GetOrganizationSmVppAccounts200ResponseInner{}
	return &this
}

// GetVppAccountId returns the VppAccountId field value if set, zero value otherwise.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetVppAccountId() string {
	if o == nil || IsNil(o.VppAccountId) {
		var ret string
		return ret
	}
	return *o.VppAccountId
}

// GetVppAccountIdOk returns a tuple with the VppAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetVppAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.VppAccountId) {
		return nil, false
	}
	return o.VppAccountId, true
}

// HasVppAccountId returns a boolean if a field has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) HasVppAccountId() bool {
	if o != nil && !IsNil(o.VppAccountId) {
		return true
	}

	return false
}

// SetVppAccountId gets a reference to the given string and assigns it to the VppAccountId field.
func (o *GetOrganizationSmVppAccounts200ResponseInner) SetVppAccountId(v string) {
	o.VppAccountId = &v
}

// GetContentToken returns the ContentToken field value if set, zero value otherwise.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetContentToken() string {
	if o == nil || IsNil(o.ContentToken) {
		var ret string
		return ret
	}
	return *o.ContentToken
}

// GetContentTokenOk returns a tuple with the ContentToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetContentTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ContentToken) {
		return nil, false
	}
	return o.ContentToken, true
}

// HasContentToken returns a boolean if a field has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) HasContentToken() bool {
	if o != nil && !IsNil(o.ContentToken) {
		return true
	}

	return false
}

// SetContentToken gets a reference to the given string and assigns it to the ContentToken field.
func (o *GetOrganizationSmVppAccounts200ResponseInner) SetContentToken(v string) {
	o.ContentToken = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetOrganizationSmVppAccounts200ResponseInner) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationSmVppAccounts200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetAllowedAdmins returns the AllowedAdmins field value if set, zero value otherwise.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetAllowedAdmins() string {
	if o == nil || IsNil(o.AllowedAdmins) {
		var ret string
		return ret
	}
	return *o.AllowedAdmins
}

// GetAllowedAdminsOk returns a tuple with the AllowedAdmins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetAllowedAdminsOk() (*string, bool) {
	if o == nil || IsNil(o.AllowedAdmins) {
		return nil, false
	}
	return o.AllowedAdmins, true
}

// HasAllowedAdmins returns a boolean if a field has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) HasAllowedAdmins() bool {
	if o != nil && !IsNil(o.AllowedAdmins) {
		return true
	}

	return false
}

// SetAllowedAdmins gets a reference to the given string and assigns it to the AllowedAdmins field.
func (o *GetOrganizationSmVppAccounts200ResponseInner) SetAllowedAdmins(v string) {
	o.AllowedAdmins = &v
}

// GetNetworkIdAdmins returns the NetworkIdAdmins field value if set, zero value otherwise.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetNetworkIdAdmins() string {
	if o == nil || IsNil(o.NetworkIdAdmins) {
		var ret string
		return ret
	}
	return *o.NetworkIdAdmins
}

// GetNetworkIdAdminsOk returns a tuple with the NetworkIdAdmins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetNetworkIdAdminsOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkIdAdmins) {
		return nil, false
	}
	return o.NetworkIdAdmins, true
}

// HasNetworkIdAdmins returns a boolean if a field has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) HasNetworkIdAdmins() bool {
	if o != nil && !IsNil(o.NetworkIdAdmins) {
		return true
	}

	return false
}

// SetNetworkIdAdmins gets a reference to the given string and assigns it to the NetworkIdAdmins field.
func (o *GetOrganizationSmVppAccounts200ResponseInner) SetNetworkIdAdmins(v string) {
	o.NetworkIdAdmins = &v
}

// GetAssignableNetworks returns the AssignableNetworks field value if set, zero value otherwise.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetAssignableNetworks() string {
	if o == nil || IsNil(o.AssignableNetworks) {
		var ret string
		return ret
	}
	return *o.AssignableNetworks
}

// GetAssignableNetworksOk returns a tuple with the AssignableNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetAssignableNetworksOk() (*string, bool) {
	if o == nil || IsNil(o.AssignableNetworks) {
		return nil, false
	}
	return o.AssignableNetworks, true
}

// HasAssignableNetworks returns a boolean if a field has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) HasAssignableNetworks() bool {
	if o != nil && !IsNil(o.AssignableNetworks) {
		return true
	}

	return false
}

// SetAssignableNetworks gets a reference to the given string and assigns it to the AssignableNetworks field.
func (o *GetOrganizationSmVppAccounts200ResponseInner) SetAssignableNetworks(v string) {
	o.AssignableNetworks = &v
}

// GetAssignableNetworkIds returns the AssignableNetworkIds field value if set, zero value otherwise.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetAssignableNetworkIds() []string {
	if o == nil || IsNil(o.AssignableNetworkIds) {
		var ret []string
		return ret
	}
	return o.AssignableNetworkIds
}

// GetAssignableNetworkIdsOk returns a tuple with the AssignableNetworkIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetAssignableNetworkIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AssignableNetworkIds) {
		return nil, false
	}
	return o.AssignableNetworkIds, true
}

// HasAssignableNetworkIds returns a boolean if a field has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) HasAssignableNetworkIds() bool {
	if o != nil && !IsNil(o.AssignableNetworkIds) {
		return true
	}

	return false
}

// SetAssignableNetworkIds gets a reference to the given []string and assigns it to the AssignableNetworkIds field.
func (o *GetOrganizationSmVppAccounts200ResponseInner) SetAssignableNetworkIds(v []string) {
	o.AssignableNetworkIds = v
}

// GetVppLocationId returns the VppLocationId field value if set, zero value otherwise.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetVppLocationId() string {
	if o == nil || IsNil(o.VppLocationId) {
		var ret string
		return ret
	}
	return *o.VppLocationId
}

// GetVppLocationIdOk returns a tuple with the VppLocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetVppLocationIdOk() (*string, bool) {
	if o == nil || IsNil(o.VppLocationId) {
		return nil, false
	}
	return o.VppLocationId, true
}

// HasVppLocationId returns a boolean if a field has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) HasVppLocationId() bool {
	if o != nil && !IsNil(o.VppLocationId) {
		return true
	}

	return false
}

// SetVppLocationId gets a reference to the given string and assigns it to the VppLocationId field.
func (o *GetOrganizationSmVppAccounts200ResponseInner) SetVppLocationId(v string) {
	o.VppLocationId = &v
}

// GetVppLocationName returns the VppLocationName field value if set, zero value otherwise.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetVppLocationName() string {
	if o == nil || IsNil(o.VppLocationName) {
		var ret string
		return ret
	}
	return *o.VppLocationName
}

// GetVppLocationNameOk returns a tuple with the VppLocationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetVppLocationNameOk() (*string, bool) {
	if o == nil || IsNil(o.VppLocationName) {
		return nil, false
	}
	return o.VppLocationName, true
}

// HasVppLocationName returns a boolean if a field has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) HasVppLocationName() bool {
	if o != nil && !IsNil(o.VppLocationName) {
		return true
	}

	return false
}

// SetVppLocationName gets a reference to the given string and assigns it to the VppLocationName field.
func (o *GetOrganizationSmVppAccounts200ResponseInner) SetVppLocationName(v string) {
	o.VppLocationName = &v
}

// GetLastSyncedAt returns the LastSyncedAt field value if set, zero value otherwise.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetLastSyncedAt() string {
	if o == nil || IsNil(o.LastSyncedAt) {
		var ret string
		return ret
	}
	return *o.LastSyncedAt
}

// GetLastSyncedAtOk returns a tuple with the LastSyncedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetLastSyncedAtOk() (*string, bool) {
	if o == nil || IsNil(o.LastSyncedAt) {
		return nil, false
	}
	return o.LastSyncedAt, true
}

// HasLastSyncedAt returns a boolean if a field has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) HasLastSyncedAt() bool {
	if o != nil && !IsNil(o.LastSyncedAt) {
		return true
	}

	return false
}

// SetLastSyncedAt gets a reference to the given string and assigns it to the LastSyncedAt field.
func (o *GetOrganizationSmVppAccounts200ResponseInner) SetLastSyncedAt(v string) {
	o.LastSyncedAt = &v
}

// GetLastForceSyncedAt returns the LastForceSyncedAt field value if set, zero value otherwise.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetLastForceSyncedAt() string {
	if o == nil || IsNil(o.LastForceSyncedAt) {
		var ret string
		return ret
	}
	return *o.LastForceSyncedAt
}

// GetLastForceSyncedAtOk returns a tuple with the LastForceSyncedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetLastForceSyncedAtOk() (*string, bool) {
	if o == nil || IsNil(o.LastForceSyncedAt) {
		return nil, false
	}
	return o.LastForceSyncedAt, true
}

// HasLastForceSyncedAt returns a boolean if a field has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) HasLastForceSyncedAt() bool {
	if o != nil && !IsNil(o.LastForceSyncedAt) {
		return true
	}

	return false
}

// SetLastForceSyncedAt gets a reference to the given string and assigns it to the LastForceSyncedAt field.
func (o *GetOrganizationSmVppAccounts200ResponseInner) SetLastForceSyncedAt(v string) {
	o.LastForceSyncedAt = &v
}

// GetParsedToken returns the ParsedToken field value if set, zero value otherwise.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetParsedToken() GetOrganizationSmVppAccounts200ResponseInnerParsedToken {
	if o == nil || IsNil(o.ParsedToken) {
		var ret GetOrganizationSmVppAccounts200ResponseInnerParsedToken
		return ret
	}
	return *o.ParsedToken
}

// GetParsedTokenOk returns a tuple with the ParsedToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetParsedTokenOk() (*GetOrganizationSmVppAccounts200ResponseInnerParsedToken, bool) {
	if o == nil || IsNil(o.ParsedToken) {
		return nil, false
	}
	return o.ParsedToken, true
}

// HasParsedToken returns a boolean if a field has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) HasParsedToken() bool {
	if o != nil && !IsNil(o.ParsedToken) {
		return true
	}

	return false
}

// SetParsedToken gets a reference to the given GetOrganizationSmVppAccounts200ResponseInnerParsedToken and assigns it to the ParsedToken field.
func (o *GetOrganizationSmVppAccounts200ResponseInner) SetParsedToken(v GetOrganizationSmVppAccounts200ResponseInnerParsedToken) {
	o.ParsedToken = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationSmVppAccounts200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetVppServiceToken returns the VppServiceToken field value if set, zero value otherwise.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetVppServiceToken() string {
	if o == nil || IsNil(o.VppServiceToken) {
		var ret string
		return ret
	}
	return *o.VppServiceToken
}

// GetVppServiceTokenOk returns a tuple with the VppServiceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) GetVppServiceTokenOk() (*string, bool) {
	if o == nil || IsNil(o.VppServiceToken) {
		return nil, false
	}
	return o.VppServiceToken, true
}

// HasVppServiceToken returns a boolean if a field has been set.
func (o *GetOrganizationSmVppAccounts200ResponseInner) HasVppServiceToken() bool {
	if o != nil && !IsNil(o.VppServiceToken) {
		return true
	}

	return false
}

// SetVppServiceToken gets a reference to the given string and assigns it to the VppServiceToken field.
func (o *GetOrganizationSmVppAccounts200ResponseInner) SetVppServiceToken(v string) {
	o.VppServiceToken = &v
}

func (o GetOrganizationSmVppAccounts200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSmVppAccounts200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VppAccountId) {
		toSerialize["vppAccountId"] = o.VppAccountId
	}
	if !IsNil(o.ContentToken) {
		toSerialize["contentToken"] = o.ContentToken
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.AllowedAdmins) {
		toSerialize["allowedAdmins"] = o.AllowedAdmins
	}
	if !IsNil(o.NetworkIdAdmins) {
		toSerialize["networkIdAdmins"] = o.NetworkIdAdmins
	}
	if !IsNil(o.AssignableNetworks) {
		toSerialize["assignableNetworks"] = o.AssignableNetworks
	}
	if !IsNil(o.AssignableNetworkIds) {
		toSerialize["assignableNetworkIds"] = o.AssignableNetworkIds
	}
	if !IsNil(o.VppLocationId) {
		toSerialize["vppLocationId"] = o.VppLocationId
	}
	if !IsNil(o.VppLocationName) {
		toSerialize["vppLocationName"] = o.VppLocationName
	}
	if !IsNil(o.LastSyncedAt) {
		toSerialize["lastSyncedAt"] = o.LastSyncedAt
	}
	if !IsNil(o.LastForceSyncedAt) {
		toSerialize["lastForceSyncedAt"] = o.LastForceSyncedAt
	}
	if !IsNil(o.ParsedToken) {
		toSerialize["parsedToken"] = o.ParsedToken
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.VppServiceToken) {
		toSerialize["vppServiceToken"] = o.VppServiceToken
	}
	return toSerialize, nil
}

type NullableGetOrganizationSmVppAccounts200ResponseInner struct {
	value *GetOrganizationSmVppAccounts200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationSmVppAccounts200ResponseInner) Get() *GetOrganizationSmVppAccounts200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationSmVppAccounts200ResponseInner) Set(val *GetOrganizationSmVppAccounts200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSmVppAccounts200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSmVppAccounts200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSmVppAccounts200ResponseInner(val *GetOrganizationSmVppAccounts200ResponseInner) *NullableGetOrganizationSmVppAccounts200ResponseInner {
	return &NullableGetOrganizationSmVppAccounts200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationSmVppAccounts200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSmVppAccounts200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


