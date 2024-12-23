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

// checks if the GetOrganizationLicenses200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationLicenses200ResponseInner{}

// GetOrganizationLicenses200ResponseInner struct for GetOrganizationLicenses200ResponseInner
type GetOrganizationLicenses200ResponseInner struct {
	// License ID
	Id *string `json:"id,omitempty"`
	// License type
	LicenseType *string `json:"licenseType,omitempty"`
	// License key
	LicenseKey *string `json:"licenseKey,omitempty"`
	// Order number
	OrderNumber *string `json:"orderNumber,omitempty"`
	// Serial number of the device the license is assigned to
	DeviceSerial *string `json:"deviceSerial,omitempty"`
	// ID of the network the license is assigned to
	NetworkId *string `json:"networkId,omitempty"`
	// The state of the license. All queued licenses have a status of `recentlyQueued`.
	State *string `json:"state,omitempty"`
	// The number of seats of the license. Only applicable to SM licenses.
	SeatCount *int32 `json:"seatCount,omitempty"`
	// The duration of the license plus all permanently queued licenses associated with it
	TotalDurationInDays *int32 `json:"totalDurationInDays,omitempty"`
	// The duration of the individual license
	DurationInDays *int32 `json:"durationInDays,omitempty"`
	// DEPRECATED List of permanently queued licenses attached to the license. Instead, use /organizations/{organizationId}/licenses?deviceSerial= to retrieved queued licenses for a given device.
	PermanentlyQueuedLicenses []GetOrganizationLicenses200ResponseInnerPermanentlyQueuedLicensesInner `json:"permanentlyQueuedLicenses,omitempty"`
	// The date the license was claimed into the organization
	ClaimDate *string `json:"claimDate,omitempty"`
	// The date the license started burning
	ActivationDate *string `json:"activationDate,omitempty"`
	// The date the license will expire
	ExpirationDate *string `json:"expirationDate,omitempty"`
	// The id of the head license this license is queued behind. If there is no head license, it returns nil.
	HeadLicenseId *string `json:"headLicenseId,omitempty"`
}

// NewGetOrganizationLicenses200ResponseInner instantiates a new GetOrganizationLicenses200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationLicenses200ResponseInner() *GetOrganizationLicenses200ResponseInner {
	this := GetOrganizationLicenses200ResponseInner{}
	return &this
}

// NewGetOrganizationLicenses200ResponseInnerWithDefaults instantiates a new GetOrganizationLicenses200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationLicenses200ResponseInnerWithDefaults() *GetOrganizationLicenses200ResponseInner {
	this := GetOrganizationLicenses200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationLicenses200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicenses200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationLicenses200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationLicenses200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *GetOrganizationLicenses200ResponseInner) GetLicenseType() string {
	if o == nil || IsNil(o.LicenseType) {
		var ret string
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicenses200ResponseInner) GetLicenseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseType) {
		return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *GetOrganizationLicenses200ResponseInner) HasLicenseType() bool {
	if o != nil && !IsNil(o.LicenseType) {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given string and assigns it to the LicenseType field.
func (o *GetOrganizationLicenses200ResponseInner) SetLicenseType(v string) {
	o.LicenseType = &v
}

// GetLicenseKey returns the LicenseKey field value if set, zero value otherwise.
func (o *GetOrganizationLicenses200ResponseInner) GetLicenseKey() string {
	if o == nil || IsNil(o.LicenseKey) {
		var ret string
		return ret
	}
	return *o.LicenseKey
}

// GetLicenseKeyOk returns a tuple with the LicenseKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicenses200ResponseInner) GetLicenseKeyOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseKey) {
		return nil, false
	}
	return o.LicenseKey, true
}

// HasLicenseKey returns a boolean if a field has been set.
func (o *GetOrganizationLicenses200ResponseInner) HasLicenseKey() bool {
	if o != nil && !IsNil(o.LicenseKey) {
		return true
	}

	return false
}

// SetLicenseKey gets a reference to the given string and assigns it to the LicenseKey field.
func (o *GetOrganizationLicenses200ResponseInner) SetLicenseKey(v string) {
	o.LicenseKey = &v
}

// GetOrderNumber returns the OrderNumber field value if set, zero value otherwise.
func (o *GetOrganizationLicenses200ResponseInner) GetOrderNumber() string {
	if o == nil || IsNil(o.OrderNumber) {
		var ret string
		return ret
	}
	return *o.OrderNumber
}

// GetOrderNumberOk returns a tuple with the OrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicenses200ResponseInner) GetOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.OrderNumber) {
		return nil, false
	}
	return o.OrderNumber, true
}

// HasOrderNumber returns a boolean if a field has been set.
func (o *GetOrganizationLicenses200ResponseInner) HasOrderNumber() bool {
	if o != nil && !IsNil(o.OrderNumber) {
		return true
	}

	return false
}

// SetOrderNumber gets a reference to the given string and assigns it to the OrderNumber field.
func (o *GetOrganizationLicenses200ResponseInner) SetOrderNumber(v string) {
	o.OrderNumber = &v
}

// GetDeviceSerial returns the DeviceSerial field value if set, zero value otherwise.
func (o *GetOrganizationLicenses200ResponseInner) GetDeviceSerial() string {
	if o == nil || IsNil(o.DeviceSerial) {
		var ret string
		return ret
	}
	return *o.DeviceSerial
}

// GetDeviceSerialOk returns a tuple with the DeviceSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicenses200ResponseInner) GetDeviceSerialOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceSerial) {
		return nil, false
	}
	return o.DeviceSerial, true
}

// HasDeviceSerial returns a boolean if a field has been set.
func (o *GetOrganizationLicenses200ResponseInner) HasDeviceSerial() bool {
	if o != nil && !IsNil(o.DeviceSerial) {
		return true
	}

	return false
}

// SetDeviceSerial gets a reference to the given string and assigns it to the DeviceSerial field.
func (o *GetOrganizationLicenses200ResponseInner) SetDeviceSerial(v string) {
	o.DeviceSerial = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *GetOrganizationLicenses200ResponseInner) GetNetworkId() string {
	if o == nil || IsNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicenses200ResponseInner) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *GetOrganizationLicenses200ResponseInner) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *GetOrganizationLicenses200ResponseInner) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GetOrganizationLicenses200ResponseInner) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicenses200ResponseInner) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GetOrganizationLicenses200ResponseInner) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GetOrganizationLicenses200ResponseInner) SetState(v string) {
	o.State = &v
}

// GetSeatCount returns the SeatCount field value if set, zero value otherwise.
func (o *GetOrganizationLicenses200ResponseInner) GetSeatCount() int32 {
	if o == nil || IsNil(o.SeatCount) {
		var ret int32
		return ret
	}
	return *o.SeatCount
}

// GetSeatCountOk returns a tuple with the SeatCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicenses200ResponseInner) GetSeatCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SeatCount) {
		return nil, false
	}
	return o.SeatCount, true
}

// HasSeatCount returns a boolean if a field has been set.
func (o *GetOrganizationLicenses200ResponseInner) HasSeatCount() bool {
	if o != nil && !IsNil(o.SeatCount) {
		return true
	}

	return false
}

// SetSeatCount gets a reference to the given int32 and assigns it to the SeatCount field.
func (o *GetOrganizationLicenses200ResponseInner) SetSeatCount(v int32) {
	o.SeatCount = &v
}

// GetTotalDurationInDays returns the TotalDurationInDays field value if set, zero value otherwise.
func (o *GetOrganizationLicenses200ResponseInner) GetTotalDurationInDays() int32 {
	if o == nil || IsNil(o.TotalDurationInDays) {
		var ret int32
		return ret
	}
	return *o.TotalDurationInDays
}

// GetTotalDurationInDaysOk returns a tuple with the TotalDurationInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicenses200ResponseInner) GetTotalDurationInDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalDurationInDays) {
		return nil, false
	}
	return o.TotalDurationInDays, true
}

// HasTotalDurationInDays returns a boolean if a field has been set.
func (o *GetOrganizationLicenses200ResponseInner) HasTotalDurationInDays() bool {
	if o != nil && !IsNil(o.TotalDurationInDays) {
		return true
	}

	return false
}

// SetTotalDurationInDays gets a reference to the given int32 and assigns it to the TotalDurationInDays field.
func (o *GetOrganizationLicenses200ResponseInner) SetTotalDurationInDays(v int32) {
	o.TotalDurationInDays = &v
}

// GetDurationInDays returns the DurationInDays field value if set, zero value otherwise.
func (o *GetOrganizationLicenses200ResponseInner) GetDurationInDays() int32 {
	if o == nil || IsNil(o.DurationInDays) {
		var ret int32
		return ret
	}
	return *o.DurationInDays
}

// GetDurationInDaysOk returns a tuple with the DurationInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicenses200ResponseInner) GetDurationInDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.DurationInDays) {
		return nil, false
	}
	return o.DurationInDays, true
}

// HasDurationInDays returns a boolean if a field has been set.
func (o *GetOrganizationLicenses200ResponseInner) HasDurationInDays() bool {
	if o != nil && !IsNil(o.DurationInDays) {
		return true
	}

	return false
}

// SetDurationInDays gets a reference to the given int32 and assigns it to the DurationInDays field.
func (o *GetOrganizationLicenses200ResponseInner) SetDurationInDays(v int32) {
	o.DurationInDays = &v
}

// GetPermanentlyQueuedLicenses returns the PermanentlyQueuedLicenses field value if set, zero value otherwise.
func (o *GetOrganizationLicenses200ResponseInner) GetPermanentlyQueuedLicenses() []GetOrganizationLicenses200ResponseInnerPermanentlyQueuedLicensesInner {
	if o == nil || IsNil(o.PermanentlyQueuedLicenses) {
		var ret []GetOrganizationLicenses200ResponseInnerPermanentlyQueuedLicensesInner
		return ret
	}
	return o.PermanentlyQueuedLicenses
}

// GetPermanentlyQueuedLicensesOk returns a tuple with the PermanentlyQueuedLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicenses200ResponseInner) GetPermanentlyQueuedLicensesOk() ([]GetOrganizationLicenses200ResponseInnerPermanentlyQueuedLicensesInner, bool) {
	if o == nil || IsNil(o.PermanentlyQueuedLicenses) {
		return nil, false
	}
	return o.PermanentlyQueuedLicenses, true
}

// HasPermanentlyQueuedLicenses returns a boolean if a field has been set.
func (o *GetOrganizationLicenses200ResponseInner) HasPermanentlyQueuedLicenses() bool {
	if o != nil && !IsNil(o.PermanentlyQueuedLicenses) {
		return true
	}

	return false
}

// SetPermanentlyQueuedLicenses gets a reference to the given []GetOrganizationLicenses200ResponseInnerPermanentlyQueuedLicensesInner and assigns it to the PermanentlyQueuedLicenses field.
func (o *GetOrganizationLicenses200ResponseInner) SetPermanentlyQueuedLicenses(v []GetOrganizationLicenses200ResponseInnerPermanentlyQueuedLicensesInner) {
	o.PermanentlyQueuedLicenses = v
}

// GetClaimDate returns the ClaimDate field value if set, zero value otherwise.
func (o *GetOrganizationLicenses200ResponseInner) GetClaimDate() string {
	if o == nil || IsNil(o.ClaimDate) {
		var ret string
		return ret
	}
	return *o.ClaimDate
}

// GetClaimDateOk returns a tuple with the ClaimDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicenses200ResponseInner) GetClaimDateOk() (*string, bool) {
	if o == nil || IsNil(o.ClaimDate) {
		return nil, false
	}
	return o.ClaimDate, true
}

// HasClaimDate returns a boolean if a field has been set.
func (o *GetOrganizationLicenses200ResponseInner) HasClaimDate() bool {
	if o != nil && !IsNil(o.ClaimDate) {
		return true
	}

	return false
}

// SetClaimDate gets a reference to the given string and assigns it to the ClaimDate field.
func (o *GetOrganizationLicenses200ResponseInner) SetClaimDate(v string) {
	o.ClaimDate = &v
}

// GetActivationDate returns the ActivationDate field value if set, zero value otherwise.
func (o *GetOrganizationLicenses200ResponseInner) GetActivationDate() string {
	if o == nil || IsNil(o.ActivationDate) {
		var ret string
		return ret
	}
	return *o.ActivationDate
}

// GetActivationDateOk returns a tuple with the ActivationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicenses200ResponseInner) GetActivationDateOk() (*string, bool) {
	if o == nil || IsNil(o.ActivationDate) {
		return nil, false
	}
	return o.ActivationDate, true
}

// HasActivationDate returns a boolean if a field has been set.
func (o *GetOrganizationLicenses200ResponseInner) HasActivationDate() bool {
	if o != nil && !IsNil(o.ActivationDate) {
		return true
	}

	return false
}

// SetActivationDate gets a reference to the given string and assigns it to the ActivationDate field.
func (o *GetOrganizationLicenses200ResponseInner) SetActivationDate(v string) {
	o.ActivationDate = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *GetOrganizationLicenses200ResponseInner) GetExpirationDate() string {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicenses200ResponseInner) GetExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *GetOrganizationLicenses200ResponseInner) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *GetOrganizationLicenses200ResponseInner) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetHeadLicenseId returns the HeadLicenseId field value if set, zero value otherwise.
func (o *GetOrganizationLicenses200ResponseInner) GetHeadLicenseId() string {
	if o == nil || IsNil(o.HeadLicenseId) {
		var ret string
		return ret
	}
	return *o.HeadLicenseId
}

// GetHeadLicenseIdOk returns a tuple with the HeadLicenseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicenses200ResponseInner) GetHeadLicenseIdOk() (*string, bool) {
	if o == nil || IsNil(o.HeadLicenseId) {
		return nil, false
	}
	return o.HeadLicenseId, true
}

// HasHeadLicenseId returns a boolean if a field has been set.
func (o *GetOrganizationLicenses200ResponseInner) HasHeadLicenseId() bool {
	if o != nil && !IsNil(o.HeadLicenseId) {
		return true
	}

	return false
}

// SetHeadLicenseId gets a reference to the given string and assigns it to the HeadLicenseId field.
func (o *GetOrganizationLicenses200ResponseInner) SetHeadLicenseId(v string) {
	o.HeadLicenseId = &v
}

func (o GetOrganizationLicenses200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationLicenses200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LicenseType) {
		toSerialize["licenseType"] = o.LicenseType
	}
	if !IsNil(o.LicenseKey) {
		toSerialize["licenseKey"] = o.LicenseKey
	}
	if !IsNil(o.OrderNumber) {
		toSerialize["orderNumber"] = o.OrderNumber
	}
	if !IsNil(o.DeviceSerial) {
		toSerialize["deviceSerial"] = o.DeviceSerial
	}
	if !IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.SeatCount) {
		toSerialize["seatCount"] = o.SeatCount
	}
	if !IsNil(o.TotalDurationInDays) {
		toSerialize["totalDurationInDays"] = o.TotalDurationInDays
	}
	if !IsNil(o.DurationInDays) {
		toSerialize["durationInDays"] = o.DurationInDays
	}
	if !IsNil(o.PermanentlyQueuedLicenses) {
		toSerialize["permanentlyQueuedLicenses"] = o.PermanentlyQueuedLicenses
	}
	if !IsNil(o.ClaimDate) {
		toSerialize["claimDate"] = o.ClaimDate
	}
	if !IsNil(o.ActivationDate) {
		toSerialize["activationDate"] = o.ActivationDate
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !IsNil(o.HeadLicenseId) {
		toSerialize["headLicenseId"] = o.HeadLicenseId
	}
	return toSerialize, nil
}

type NullableGetOrganizationLicenses200ResponseInner struct {
	value *GetOrganizationLicenses200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationLicenses200ResponseInner) Get() *GetOrganizationLicenses200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationLicenses200ResponseInner) Set(val *GetOrganizationLicenses200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationLicenses200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationLicenses200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationLicenses200ResponseInner(val *GetOrganizationLicenses200ResponseInner) *NullableGetOrganizationLicenses200ResponseInner {
	return &NullableGetOrganizationLicenses200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationLicenses200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationLicenses200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


