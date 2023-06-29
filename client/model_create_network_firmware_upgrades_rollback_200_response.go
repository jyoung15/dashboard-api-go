/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the CreateNetworkFirmwareUpgradesRollback200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkFirmwareUpgradesRollback200Response{}

// CreateNetworkFirmwareUpgradesRollback200Response struct for CreateNetworkFirmwareUpgradesRollback200Response
type CreateNetworkFirmwareUpgradesRollback200Response struct {
	// Product type to rollback (if the network is a combined network)
	Product *string `json:"product,omitempty"`
	// Status of the rollback
	Status *string `json:"status,omitempty"`
	// Batch ID of the firmware rollback
	UpgradeBatchId *string `json:"upgradeBatchId,omitempty"`
	// Scheduled time for the rollback
	Time *time.Time `json:"time,omitempty"`
	ToVersion *CreateNetworkFirmwareUpgradesRollback200ResponseToVersion `json:"toVersion,omitempty"`
	// Reasons for the rollback
	Reasons []CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner `json:"reasons,omitempty"`
}

// NewCreateNetworkFirmwareUpgradesRollback200Response instantiates a new CreateNetworkFirmwareUpgradesRollback200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkFirmwareUpgradesRollback200Response() *CreateNetworkFirmwareUpgradesRollback200Response {
	this := CreateNetworkFirmwareUpgradesRollback200Response{}
	return &this
}

// NewCreateNetworkFirmwareUpgradesRollback200ResponseWithDefaults instantiates a new CreateNetworkFirmwareUpgradesRollback200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkFirmwareUpgradesRollback200ResponseWithDefaults() *CreateNetworkFirmwareUpgradesRollback200Response {
	this := CreateNetworkFirmwareUpgradesRollback200Response{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) SetProduct(v string) {
	o.Product = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) SetStatus(v string) {
	o.Status = &v
}

// GetUpgradeBatchId returns the UpgradeBatchId field value if set, zero value otherwise.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetUpgradeBatchId() string {
	if o == nil || IsNil(o.UpgradeBatchId) {
		var ret string
		return ret
	}
	return *o.UpgradeBatchId
}

// GetUpgradeBatchIdOk returns a tuple with the UpgradeBatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetUpgradeBatchIdOk() (*string, bool) {
	if o == nil || IsNil(o.UpgradeBatchId) {
		return nil, false
	}
	return o.UpgradeBatchId, true
}

// HasUpgradeBatchId returns a boolean if a field has been set.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) HasUpgradeBatchId() bool {
	if o != nil && !IsNil(o.UpgradeBatchId) {
		return true
	}

	return false
}

// SetUpgradeBatchId gets a reference to the given string and assigns it to the UpgradeBatchId field.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) SetUpgradeBatchId(v string) {
	o.UpgradeBatchId = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) SetTime(v time.Time) {
	o.Time = &v
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetToVersion() CreateNetworkFirmwareUpgradesRollback200ResponseToVersion {
	if o == nil || IsNil(o.ToVersion) {
		var ret CreateNetworkFirmwareUpgradesRollback200ResponseToVersion
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetToVersionOk() (*CreateNetworkFirmwareUpgradesRollback200ResponseToVersion, bool) {
	if o == nil || IsNil(o.ToVersion) {
		return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) HasToVersion() bool {
	if o != nil && !IsNil(o.ToVersion) {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given CreateNetworkFirmwareUpgradesRollback200ResponseToVersion and assigns it to the ToVersion field.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) SetToVersion(v CreateNetworkFirmwareUpgradesRollback200ResponseToVersion) {
	o.ToVersion = &v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetReasons() []CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner {
	if o == nil || IsNil(o.Reasons) {
		var ret []CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) GetReasonsOk() ([]CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) HasReasons() bool {
	if o != nil && !IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner and assigns it to the Reasons field.
func (o *CreateNetworkFirmwareUpgradesRollback200Response) SetReasons(v []CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner) {
	o.Reasons = v
}

func (o CreateNetworkFirmwareUpgradesRollback200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkFirmwareUpgradesRollback200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UpgradeBatchId) {
		toSerialize["upgradeBatchId"] = o.UpgradeBatchId
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.ToVersion) {
		toSerialize["toVersion"] = o.ToVersion
	}
	if !IsNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	return toSerialize, nil
}

type NullableCreateNetworkFirmwareUpgradesRollback200Response struct {
	value *CreateNetworkFirmwareUpgradesRollback200Response
	isSet bool
}

func (v NullableCreateNetworkFirmwareUpgradesRollback200Response) Get() *CreateNetworkFirmwareUpgradesRollback200Response {
	return v.value
}

func (v *NullableCreateNetworkFirmwareUpgradesRollback200Response) Set(val *CreateNetworkFirmwareUpgradesRollback200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkFirmwareUpgradesRollback200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkFirmwareUpgradesRollback200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkFirmwareUpgradesRollback200Response(val *CreateNetworkFirmwareUpgradesRollback200Response) *NullableCreateNetworkFirmwareUpgradesRollback200Response {
	return &NullableCreateNetworkFirmwareUpgradesRollback200Response{value: val, isSet: true}
}

func (v NullableCreateNetworkFirmwareUpgradesRollback200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkFirmwareUpgradesRollback200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


