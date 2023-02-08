/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateOrganizationActionBatch201Response struct for CreateOrganizationActionBatch201Response
type CreateOrganizationActionBatch201Response struct {
	// ID of the action batch. Can be used to check the status of the action batch at /organizations/{organizationId}/actionBatches/{actionBatchId}
	Id *string `json:"id,omitempty"`
	// ID of the organization this action batch belongs to
	OrganizationId *string `json:"organizationId,omitempty"`
	// Flag describing whether the action should be previewed before executing or not
	Confirmed *bool `json:"confirmed,omitempty"`
	// Flag describing whether actions should run synchronously or asynchronously
	Synchronous *bool `json:"synchronous,omitempty"`
	Status *CreateOrganizationActionBatch201ResponseStatus `json:"status,omitempty"`
	// A set of changes made as part of this action (<a href='https://developer.cisco.com/meraki/api/#/rest/guides/action-batches/'>more details</a>)
	Actions []CreateOrganizationActionBatch201ResponseActionsInner `json:"actions"`
}

// NewCreateOrganizationActionBatch201Response instantiates a new CreateOrganizationActionBatch201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationActionBatch201Response(actions []CreateOrganizationActionBatch201ResponseActionsInner) *CreateOrganizationActionBatch201Response {
	this := CreateOrganizationActionBatch201Response{}
	this.Actions = actions
	return &this
}

// NewCreateOrganizationActionBatch201ResponseWithDefaults instantiates a new CreateOrganizationActionBatch201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationActionBatch201ResponseWithDefaults() *CreateOrganizationActionBatch201Response {
	this := CreateOrganizationActionBatch201Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateOrganizationActionBatch201Response) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationActionBatch201Response) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateOrganizationActionBatch201Response) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateOrganizationActionBatch201Response) SetId(v string) {
	o.Id = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *CreateOrganizationActionBatch201Response) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationActionBatch201Response) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *CreateOrganizationActionBatch201Response) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *CreateOrganizationActionBatch201Response) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetConfirmed returns the Confirmed field value if set, zero value otherwise.
func (o *CreateOrganizationActionBatch201Response) GetConfirmed() bool {
	if o == nil || isNil(o.Confirmed) {
		var ret bool
		return ret
	}
	return *o.Confirmed
}

// GetConfirmedOk returns a tuple with the Confirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationActionBatch201Response) GetConfirmedOk() (*bool, bool) {
	if o == nil || isNil(o.Confirmed) {
    return nil, false
	}
	return o.Confirmed, true
}

// HasConfirmed returns a boolean if a field has been set.
func (o *CreateOrganizationActionBatch201Response) HasConfirmed() bool {
	if o != nil && !isNil(o.Confirmed) {
		return true
	}

	return false
}

// SetConfirmed gets a reference to the given bool and assigns it to the Confirmed field.
func (o *CreateOrganizationActionBatch201Response) SetConfirmed(v bool) {
	o.Confirmed = &v
}

// GetSynchronous returns the Synchronous field value if set, zero value otherwise.
func (o *CreateOrganizationActionBatch201Response) GetSynchronous() bool {
	if o == nil || isNil(o.Synchronous) {
		var ret bool
		return ret
	}
	return *o.Synchronous
}

// GetSynchronousOk returns a tuple with the Synchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationActionBatch201Response) GetSynchronousOk() (*bool, bool) {
	if o == nil || isNil(o.Synchronous) {
    return nil, false
	}
	return o.Synchronous, true
}

// HasSynchronous returns a boolean if a field has been set.
func (o *CreateOrganizationActionBatch201Response) HasSynchronous() bool {
	if o != nil && !isNil(o.Synchronous) {
		return true
	}

	return false
}

// SetSynchronous gets a reference to the given bool and assigns it to the Synchronous field.
func (o *CreateOrganizationActionBatch201Response) SetSynchronous(v bool) {
	o.Synchronous = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateOrganizationActionBatch201Response) GetStatus() CreateOrganizationActionBatch201ResponseStatus {
	if o == nil || isNil(o.Status) {
		var ret CreateOrganizationActionBatch201ResponseStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationActionBatch201Response) GetStatusOk() (*CreateOrganizationActionBatch201ResponseStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateOrganizationActionBatch201Response) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CreateOrganizationActionBatch201ResponseStatus and assigns it to the Status field.
func (o *CreateOrganizationActionBatch201Response) SetStatus(v CreateOrganizationActionBatch201ResponseStatus) {
	o.Status = &v
}

// GetActions returns the Actions field value
func (o *CreateOrganizationActionBatch201Response) GetActions() []CreateOrganizationActionBatch201ResponseActionsInner {
	if o == nil {
		var ret []CreateOrganizationActionBatch201ResponseActionsInner
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationActionBatch201Response) GetActionsOk() ([]CreateOrganizationActionBatch201ResponseActionsInner, bool) {
	if o == nil {
    return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *CreateOrganizationActionBatch201Response) SetActions(v []CreateOrganizationActionBatch201ResponseActionsInner) {
	o.Actions = v
}

func (o CreateOrganizationActionBatch201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !isNil(o.Confirmed) {
		toSerialize["confirmed"] = o.Confirmed
	}
	if !isNil(o.Synchronous) {
		toSerialize["synchronous"] = o.Synchronous
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["actions"] = o.Actions
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationActionBatch201Response struct {
	value *CreateOrganizationActionBatch201Response
	isSet bool
}

func (v NullableCreateOrganizationActionBatch201Response) Get() *CreateOrganizationActionBatch201Response {
	return v.value
}

func (v *NullableCreateOrganizationActionBatch201Response) Set(val *CreateOrganizationActionBatch201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationActionBatch201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationActionBatch201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationActionBatch201Response(val *CreateOrganizationActionBatch201Response) *NullableCreateOrganizationActionBatch201Response {
	return &NullableCreateOrganizationActionBatch201Response{value: val, isSet: true}
}

func (v NullableCreateOrganizationActionBatch201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationActionBatch201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


