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

// CreateNetworkFloorPlanRequest struct for CreateNetworkFloorPlanRequest
type CreateNetworkFloorPlanRequest struct {
	// The name of your floor plan.
	Name string `json:"name"`
	Center *CreateNetworkFloorPlanRequestCenter `json:"center,omitempty"`
	BottomLeftCorner *CreateNetworkFloorPlanRequestBottomLeftCorner `json:"bottomLeftCorner,omitempty"`
	BottomRightCorner *CreateNetworkFloorPlanRequestBottomRightCorner `json:"bottomRightCorner,omitempty"`
	TopLeftCorner *CreateNetworkFloorPlanRequestTopLeftCorner `json:"topLeftCorner,omitempty"`
	TopRightCorner *CreateNetworkFloorPlanRequestTopRightCorner `json:"topRightCorner,omitempty"`
	// The file contents (a base 64 encoded string) of your image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in.
	ImageContents string `json:"imageContents"`
}

// NewCreateNetworkFloorPlanRequest instantiates a new CreateNetworkFloorPlanRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkFloorPlanRequest(name string, imageContents string) *CreateNetworkFloorPlanRequest {
	this := CreateNetworkFloorPlanRequest{}
	this.Name = name
	this.ImageContents = imageContents
	return &this
}

// NewCreateNetworkFloorPlanRequestWithDefaults instantiates a new CreateNetworkFloorPlanRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkFloorPlanRequestWithDefaults() *CreateNetworkFloorPlanRequest {
	this := CreateNetworkFloorPlanRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNetworkFloorPlanRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkFloorPlanRequest) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkFloorPlanRequest) SetName(v string) {
	o.Name = v
}

// GetCenter returns the Center field value if set, zero value otherwise.
func (o *CreateNetworkFloorPlanRequest) GetCenter() CreateNetworkFloorPlanRequestCenter {
	if o == nil || isNil(o.Center) {
		var ret CreateNetworkFloorPlanRequestCenter
		return ret
	}
	return *o.Center
}

// GetCenterOk returns a tuple with the Center field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFloorPlanRequest) GetCenterOk() (*CreateNetworkFloorPlanRequestCenter, bool) {
	if o == nil || isNil(o.Center) {
    return nil, false
	}
	return o.Center, true
}

// HasCenter returns a boolean if a field has been set.
func (o *CreateNetworkFloorPlanRequest) HasCenter() bool {
	if o != nil && !isNil(o.Center) {
		return true
	}

	return false
}

// SetCenter gets a reference to the given CreateNetworkFloorPlanRequestCenter and assigns it to the Center field.
func (o *CreateNetworkFloorPlanRequest) SetCenter(v CreateNetworkFloorPlanRequestCenter) {
	o.Center = &v
}

// GetBottomLeftCorner returns the BottomLeftCorner field value if set, zero value otherwise.
func (o *CreateNetworkFloorPlanRequest) GetBottomLeftCorner() CreateNetworkFloorPlanRequestBottomLeftCorner {
	if o == nil || isNil(o.BottomLeftCorner) {
		var ret CreateNetworkFloorPlanRequestBottomLeftCorner
		return ret
	}
	return *o.BottomLeftCorner
}

// GetBottomLeftCornerOk returns a tuple with the BottomLeftCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFloorPlanRequest) GetBottomLeftCornerOk() (*CreateNetworkFloorPlanRequestBottomLeftCorner, bool) {
	if o == nil || isNil(o.BottomLeftCorner) {
    return nil, false
	}
	return o.BottomLeftCorner, true
}

// HasBottomLeftCorner returns a boolean if a field has been set.
func (o *CreateNetworkFloorPlanRequest) HasBottomLeftCorner() bool {
	if o != nil && !isNil(o.BottomLeftCorner) {
		return true
	}

	return false
}

// SetBottomLeftCorner gets a reference to the given CreateNetworkFloorPlanRequestBottomLeftCorner and assigns it to the BottomLeftCorner field.
func (o *CreateNetworkFloorPlanRequest) SetBottomLeftCorner(v CreateNetworkFloorPlanRequestBottomLeftCorner) {
	o.BottomLeftCorner = &v
}

// GetBottomRightCorner returns the BottomRightCorner field value if set, zero value otherwise.
func (o *CreateNetworkFloorPlanRequest) GetBottomRightCorner() CreateNetworkFloorPlanRequestBottomRightCorner {
	if o == nil || isNil(o.BottomRightCorner) {
		var ret CreateNetworkFloorPlanRequestBottomRightCorner
		return ret
	}
	return *o.BottomRightCorner
}

// GetBottomRightCornerOk returns a tuple with the BottomRightCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFloorPlanRequest) GetBottomRightCornerOk() (*CreateNetworkFloorPlanRequestBottomRightCorner, bool) {
	if o == nil || isNil(o.BottomRightCorner) {
    return nil, false
	}
	return o.BottomRightCorner, true
}

// HasBottomRightCorner returns a boolean if a field has been set.
func (o *CreateNetworkFloorPlanRequest) HasBottomRightCorner() bool {
	if o != nil && !isNil(o.BottomRightCorner) {
		return true
	}

	return false
}

// SetBottomRightCorner gets a reference to the given CreateNetworkFloorPlanRequestBottomRightCorner and assigns it to the BottomRightCorner field.
func (o *CreateNetworkFloorPlanRequest) SetBottomRightCorner(v CreateNetworkFloorPlanRequestBottomRightCorner) {
	o.BottomRightCorner = &v
}

// GetTopLeftCorner returns the TopLeftCorner field value if set, zero value otherwise.
func (o *CreateNetworkFloorPlanRequest) GetTopLeftCorner() CreateNetworkFloorPlanRequestTopLeftCorner {
	if o == nil || isNil(o.TopLeftCorner) {
		var ret CreateNetworkFloorPlanRequestTopLeftCorner
		return ret
	}
	return *o.TopLeftCorner
}

// GetTopLeftCornerOk returns a tuple with the TopLeftCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFloorPlanRequest) GetTopLeftCornerOk() (*CreateNetworkFloorPlanRequestTopLeftCorner, bool) {
	if o == nil || isNil(o.TopLeftCorner) {
    return nil, false
	}
	return o.TopLeftCorner, true
}

// HasTopLeftCorner returns a boolean if a field has been set.
func (o *CreateNetworkFloorPlanRequest) HasTopLeftCorner() bool {
	if o != nil && !isNil(o.TopLeftCorner) {
		return true
	}

	return false
}

// SetTopLeftCorner gets a reference to the given CreateNetworkFloorPlanRequestTopLeftCorner and assigns it to the TopLeftCorner field.
func (o *CreateNetworkFloorPlanRequest) SetTopLeftCorner(v CreateNetworkFloorPlanRequestTopLeftCorner) {
	o.TopLeftCorner = &v
}

// GetTopRightCorner returns the TopRightCorner field value if set, zero value otherwise.
func (o *CreateNetworkFloorPlanRequest) GetTopRightCorner() CreateNetworkFloorPlanRequestTopRightCorner {
	if o == nil || isNil(o.TopRightCorner) {
		var ret CreateNetworkFloorPlanRequestTopRightCorner
		return ret
	}
	return *o.TopRightCorner
}

// GetTopRightCornerOk returns a tuple with the TopRightCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFloorPlanRequest) GetTopRightCornerOk() (*CreateNetworkFloorPlanRequestTopRightCorner, bool) {
	if o == nil || isNil(o.TopRightCorner) {
    return nil, false
	}
	return o.TopRightCorner, true
}

// HasTopRightCorner returns a boolean if a field has been set.
func (o *CreateNetworkFloorPlanRequest) HasTopRightCorner() bool {
	if o != nil && !isNil(o.TopRightCorner) {
		return true
	}

	return false
}

// SetTopRightCorner gets a reference to the given CreateNetworkFloorPlanRequestTopRightCorner and assigns it to the TopRightCorner field.
func (o *CreateNetworkFloorPlanRequest) SetTopRightCorner(v CreateNetworkFloorPlanRequestTopRightCorner) {
	o.TopRightCorner = &v
}

// GetImageContents returns the ImageContents field value
func (o *CreateNetworkFloorPlanRequest) GetImageContents() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageContents
}

// GetImageContentsOk returns a tuple with the ImageContents field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkFloorPlanRequest) GetImageContentsOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ImageContents, true
}

// SetImageContents sets field value
func (o *CreateNetworkFloorPlanRequest) SetImageContents(v string) {
	o.ImageContents = v
}

func (o CreateNetworkFloorPlanRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Center) {
		toSerialize["center"] = o.Center
	}
	if !isNil(o.BottomLeftCorner) {
		toSerialize["bottomLeftCorner"] = o.BottomLeftCorner
	}
	if !isNil(o.BottomRightCorner) {
		toSerialize["bottomRightCorner"] = o.BottomRightCorner
	}
	if !isNil(o.TopLeftCorner) {
		toSerialize["topLeftCorner"] = o.TopLeftCorner
	}
	if !isNil(o.TopRightCorner) {
		toSerialize["topRightCorner"] = o.TopRightCorner
	}
	if true {
		toSerialize["imageContents"] = o.ImageContents
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkFloorPlanRequest struct {
	value *CreateNetworkFloorPlanRequest
	isSet bool
}

func (v NullableCreateNetworkFloorPlanRequest) Get() *CreateNetworkFloorPlanRequest {
	return v.value
}

func (v *NullableCreateNetworkFloorPlanRequest) Set(val *CreateNetworkFloorPlanRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkFloorPlanRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkFloorPlanRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkFloorPlanRequest(val *CreateNetworkFloorPlanRequest) *NullableCreateNetworkFloorPlanRequest {
	return &NullableCreateNetworkFloorPlanRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkFloorPlanRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkFloorPlanRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


