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

// checks if the GetOrganizationFirmwareUpgradesByDevice200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationFirmwareUpgradesByDevice200ResponseInner{}

// GetOrganizationFirmwareUpgradesByDevice200ResponseInner struct for GetOrganizationFirmwareUpgradesByDevice200ResponseInner
type GetOrganizationFirmwareUpgradesByDevice200ResponseInner struct {
	// Serial of the device
	Serial *string `json:"serial,omitempty"`
	// Name assigned to the device
	Name *string `json:"name,omitempty"`
	// Status of the device upgrade
	DeviceStatus *string `json:"deviceStatus,omitempty"`
	Upgrade *GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade `json:"upgrade,omitempty"`
}

// NewGetOrganizationFirmwareUpgradesByDevice200ResponseInner instantiates a new GetOrganizationFirmwareUpgradesByDevice200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationFirmwareUpgradesByDevice200ResponseInner() *GetOrganizationFirmwareUpgradesByDevice200ResponseInner {
	this := GetOrganizationFirmwareUpgradesByDevice200ResponseInner{}
	return &this
}

// NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerWithDefaults instantiates a new GetOrganizationFirmwareUpgradesByDevice200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationFirmwareUpgradesByDevice200ResponseInnerWithDefaults() *GetOrganizationFirmwareUpgradesByDevice200ResponseInner {
	this := GetOrganizationFirmwareUpgradesByDevice200ResponseInner{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) SetSerial(v string) {
	o.Serial = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetDeviceStatus returns the DeviceStatus field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) GetDeviceStatus() string {
	if o == nil || IsNil(o.DeviceStatus) {
		var ret string
		return ret
	}
	return *o.DeviceStatus
}

// GetDeviceStatusOk returns a tuple with the DeviceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) GetDeviceStatusOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceStatus) {
		return nil, false
	}
	return o.DeviceStatus, true
}

// HasDeviceStatus returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) HasDeviceStatus() bool {
	if o != nil && !IsNil(o.DeviceStatus) {
		return true
	}

	return false
}

// SetDeviceStatus gets a reference to the given string and assigns it to the DeviceStatus field.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) SetDeviceStatus(v string) {
	o.DeviceStatus = &v
}

// GetUpgrade returns the Upgrade field value if set, zero value otherwise.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) GetUpgrade() GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade {
	if o == nil || IsNil(o.Upgrade) {
		var ret GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade
		return ret
	}
	return *o.Upgrade
}

// GetUpgradeOk returns a tuple with the Upgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) GetUpgradeOk() (*GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade, bool) {
	if o == nil || IsNil(o.Upgrade) {
		return nil, false
	}
	return o.Upgrade, true
}

// HasUpgrade returns a boolean if a field has been set.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) HasUpgrade() bool {
	if o != nil && !IsNil(o.Upgrade) {
		return true
	}

	return false
}

// SetUpgrade gets a reference to the given GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade and assigns it to the Upgrade field.
func (o *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) SetUpgrade(v GetOrganizationFirmwareUpgradesByDevice200ResponseInnerUpgrade) {
	o.Upgrade = &v
}

func (o GetOrganizationFirmwareUpgradesByDevice200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationFirmwareUpgradesByDevice200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DeviceStatus) {
		toSerialize["deviceStatus"] = o.DeviceStatus
	}
	if !IsNil(o.Upgrade) {
		toSerialize["upgrade"] = o.Upgrade
	}
	return toSerialize, nil
}

type NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInner struct {
	value *GetOrganizationFirmwareUpgradesByDevice200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInner) Get() *GetOrganizationFirmwareUpgradesByDevice200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInner) Set(val *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationFirmwareUpgradesByDevice200ResponseInner(val *GetOrganizationFirmwareUpgradesByDevice200ResponseInner) *NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInner {
	return &NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationFirmwareUpgradesByDevice200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


