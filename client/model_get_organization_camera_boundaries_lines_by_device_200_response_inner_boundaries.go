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

// checks if the GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries{}

// GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries Configured line boundaries of the camera
type GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries struct {
	// The line boundary id
	Id *string `json:"id,omitempty"`
	// The line boundary type
	Type *string `json:"type,omitempty"`
	// The line boundary name
	Name *string `json:"name,omitempty"`
	// The line boundary vertices
	Vertices []GetOrganizationCameraBoundariesAreasByDevice200ResponseInnerBoundariesVerticesInner `json:"vertices,omitempty"`
	DirectionVertex *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundariesDirectionVertex `json:"directionVertex,omitempty"`
}

// NewGetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries instantiates a new GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries() *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries {
	this := GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries{}
	return &this
}

// NewGetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundariesWithDefaults instantiates a new GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundariesWithDefaults() *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries {
	this := GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) SetName(v string) {
	o.Name = &v
}

// GetVertices returns the Vertices field value if set, zero value otherwise.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) GetVertices() []GetOrganizationCameraBoundariesAreasByDevice200ResponseInnerBoundariesVerticesInner {
	if o == nil || IsNil(o.Vertices) {
		var ret []GetOrganizationCameraBoundariesAreasByDevice200ResponseInnerBoundariesVerticesInner
		return ret
	}
	return o.Vertices
}

// GetVerticesOk returns a tuple with the Vertices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) GetVerticesOk() ([]GetOrganizationCameraBoundariesAreasByDevice200ResponseInnerBoundariesVerticesInner, bool) {
	if o == nil || IsNil(o.Vertices) {
		return nil, false
	}
	return o.Vertices, true
}

// HasVertices returns a boolean if a field has been set.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) HasVertices() bool {
	if o != nil && !IsNil(o.Vertices) {
		return true
	}

	return false
}

// SetVertices gets a reference to the given []GetOrganizationCameraBoundariesAreasByDevice200ResponseInnerBoundariesVerticesInner and assigns it to the Vertices field.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) SetVertices(v []GetOrganizationCameraBoundariesAreasByDevice200ResponseInnerBoundariesVerticesInner) {
	o.Vertices = v
}

// GetDirectionVertex returns the DirectionVertex field value if set, zero value otherwise.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) GetDirectionVertex() GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundariesDirectionVertex {
	if o == nil || IsNil(o.DirectionVertex) {
		var ret GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundariesDirectionVertex
		return ret
	}
	return *o.DirectionVertex
}

// GetDirectionVertexOk returns a tuple with the DirectionVertex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) GetDirectionVertexOk() (*GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundariesDirectionVertex, bool) {
	if o == nil || IsNil(o.DirectionVertex) {
		return nil, false
	}
	return o.DirectionVertex, true
}

// HasDirectionVertex returns a boolean if a field has been set.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) HasDirectionVertex() bool {
	if o != nil && !IsNil(o.DirectionVertex) {
		return true
	}

	return false
}

// SetDirectionVertex gets a reference to the given GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundariesDirectionVertex and assigns it to the DirectionVertex field.
func (o *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) SetDirectionVertex(v GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundariesDirectionVertex) {
	o.DirectionVertex = &v
}

func (o GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Vertices) {
		toSerialize["vertices"] = o.Vertices
	}
	if !IsNil(o.DirectionVertex) {
		toSerialize["directionVertex"] = o.DirectionVertex
	}
	return toSerialize, nil
}

type NullableGetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries struct {
	value *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries
	isSet bool
}

func (v NullableGetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) Get() *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries {
	return v.value
}

func (v *NullableGetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) Set(val *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries(val *GetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) *NullableGetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries {
	return &NullableGetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries{value: val, isSet: true}
}

func (v NullableGetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationCameraBoundariesLinesByDevice200ResponseInnerBoundaries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


