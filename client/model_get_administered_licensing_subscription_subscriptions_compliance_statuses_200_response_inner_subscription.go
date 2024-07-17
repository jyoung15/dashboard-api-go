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

// checks if the GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription{}

// GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription Subscription details
type GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription struct {
	// Subscription's ID
	Id *string `json:"id,omitempty"`
	// Friendly name to identify the subscription
	Name *string `json:"name,omitempty"`
	// One of the following: \"inactive\" | \"active\" | \"out_of_compliance\" | \"expired\" | \"canceled\"
	Status *string `json:"status,omitempty"`
}

// NewGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription instantiates a new GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription() *GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription {
	this := GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription{}
	return &this
}

// NewGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscriptionWithDefaults instantiates a new GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscriptionWithDefaults() *GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription {
	this := GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) SetStatus(v string) {
	o.Status = &v
}

func (o GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription struct {
	value *GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription
	isSet bool
}

func (v NullableGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) Get() *GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription {
	return v.value
}

func (v *NullableGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) Set(val *GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription(val *GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) *NullableGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription {
	return &NullableGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription{value: val, isSet: true}
}

func (v NullableGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInnerSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


