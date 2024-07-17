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

// checks if the UpdateNetworkFloorPlanRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkFloorPlanRequest{}

// UpdateNetworkFloorPlanRequest struct for UpdateNetworkFloorPlanRequest
type UpdateNetworkFloorPlanRequest struct {
	// The name of your floor plan.
	Name *string `json:"name,omitempty"`
	Center *UpdateNetworkFloorPlanRequestCenter `json:"center,omitempty"`
	BottomLeftCorner *GetNetworkFloorPlans200ResponseInnerBottomLeftCorner `json:"bottomLeftCorner,omitempty"`
	BottomRightCorner *GetNetworkFloorPlans200ResponseInnerBottomRightCorner `json:"bottomRightCorner,omitempty"`
	TopLeftCorner *GetNetworkFloorPlans200ResponseInnerTopLeftCorner `json:"topLeftCorner,omitempty"`
	TopRightCorner *GetNetworkFloorPlans200ResponseInnerTopRightCorner `json:"topRightCorner,omitempty"`
	// The file contents (a base 64 encoded string) of your new image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in. If you upload a new image, and you do NOT specify any new geolocation fields ('center, 'topLeftCorner', etc), the floor plan will be recentered with no rotation in order to maintain the aspect ratio of your new image.
	ImageContents *string `json:"imageContents,omitempty"`
}

// NewUpdateNetworkFloorPlanRequest instantiates a new UpdateNetworkFloorPlanRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkFloorPlanRequest() *UpdateNetworkFloorPlanRequest {
	this := UpdateNetworkFloorPlanRequest{}
	return &this
}

// NewUpdateNetworkFloorPlanRequestWithDefaults instantiates a new UpdateNetworkFloorPlanRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkFloorPlanRequestWithDefaults() *UpdateNetworkFloorPlanRequest {
	this := UpdateNetworkFloorPlanRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkFloorPlanRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFloorPlanRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkFloorPlanRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkFloorPlanRequest) SetName(v string) {
	o.Name = &v
}

// GetCenter returns the Center field value if set, zero value otherwise.
func (o *UpdateNetworkFloorPlanRequest) GetCenter() UpdateNetworkFloorPlanRequestCenter {
	if o == nil || IsNil(o.Center) {
		var ret UpdateNetworkFloorPlanRequestCenter
		return ret
	}
	return *o.Center
}

// GetCenterOk returns a tuple with the Center field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFloorPlanRequest) GetCenterOk() (*UpdateNetworkFloorPlanRequestCenter, bool) {
	if o == nil || IsNil(o.Center) {
		return nil, false
	}
	return o.Center, true
}

// HasCenter returns a boolean if a field has been set.
func (o *UpdateNetworkFloorPlanRequest) HasCenter() bool {
	if o != nil && !IsNil(o.Center) {
		return true
	}

	return false
}

// SetCenter gets a reference to the given UpdateNetworkFloorPlanRequestCenter and assigns it to the Center field.
func (o *UpdateNetworkFloorPlanRequest) SetCenter(v UpdateNetworkFloorPlanRequestCenter) {
	o.Center = &v
}

// GetBottomLeftCorner returns the BottomLeftCorner field value if set, zero value otherwise.
func (o *UpdateNetworkFloorPlanRequest) GetBottomLeftCorner() GetNetworkFloorPlans200ResponseInnerBottomLeftCorner {
	if o == nil || IsNil(o.BottomLeftCorner) {
		var ret GetNetworkFloorPlans200ResponseInnerBottomLeftCorner
		return ret
	}
	return *o.BottomLeftCorner
}

// GetBottomLeftCornerOk returns a tuple with the BottomLeftCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFloorPlanRequest) GetBottomLeftCornerOk() (*GetNetworkFloorPlans200ResponseInnerBottomLeftCorner, bool) {
	if o == nil || IsNil(o.BottomLeftCorner) {
		return nil, false
	}
	return o.BottomLeftCorner, true
}

// HasBottomLeftCorner returns a boolean if a field has been set.
func (o *UpdateNetworkFloorPlanRequest) HasBottomLeftCorner() bool {
	if o != nil && !IsNil(o.BottomLeftCorner) {
		return true
	}

	return false
}

// SetBottomLeftCorner gets a reference to the given GetNetworkFloorPlans200ResponseInnerBottomLeftCorner and assigns it to the BottomLeftCorner field.
func (o *UpdateNetworkFloorPlanRequest) SetBottomLeftCorner(v GetNetworkFloorPlans200ResponseInnerBottomLeftCorner) {
	o.BottomLeftCorner = &v
}

// GetBottomRightCorner returns the BottomRightCorner field value if set, zero value otherwise.
func (o *UpdateNetworkFloorPlanRequest) GetBottomRightCorner() GetNetworkFloorPlans200ResponseInnerBottomRightCorner {
	if o == nil || IsNil(o.BottomRightCorner) {
		var ret GetNetworkFloorPlans200ResponseInnerBottomRightCorner
		return ret
	}
	return *o.BottomRightCorner
}

// GetBottomRightCornerOk returns a tuple with the BottomRightCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFloorPlanRequest) GetBottomRightCornerOk() (*GetNetworkFloorPlans200ResponseInnerBottomRightCorner, bool) {
	if o == nil || IsNil(o.BottomRightCorner) {
		return nil, false
	}
	return o.BottomRightCorner, true
}

// HasBottomRightCorner returns a boolean if a field has been set.
func (o *UpdateNetworkFloorPlanRequest) HasBottomRightCorner() bool {
	if o != nil && !IsNil(o.BottomRightCorner) {
		return true
	}

	return false
}

// SetBottomRightCorner gets a reference to the given GetNetworkFloorPlans200ResponseInnerBottomRightCorner and assigns it to the BottomRightCorner field.
func (o *UpdateNetworkFloorPlanRequest) SetBottomRightCorner(v GetNetworkFloorPlans200ResponseInnerBottomRightCorner) {
	o.BottomRightCorner = &v
}

// GetTopLeftCorner returns the TopLeftCorner field value if set, zero value otherwise.
func (o *UpdateNetworkFloorPlanRequest) GetTopLeftCorner() GetNetworkFloorPlans200ResponseInnerTopLeftCorner {
	if o == nil || IsNil(o.TopLeftCorner) {
		var ret GetNetworkFloorPlans200ResponseInnerTopLeftCorner
		return ret
	}
	return *o.TopLeftCorner
}

// GetTopLeftCornerOk returns a tuple with the TopLeftCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFloorPlanRequest) GetTopLeftCornerOk() (*GetNetworkFloorPlans200ResponseInnerTopLeftCorner, bool) {
	if o == nil || IsNil(o.TopLeftCorner) {
		return nil, false
	}
	return o.TopLeftCorner, true
}

// HasTopLeftCorner returns a boolean if a field has been set.
func (o *UpdateNetworkFloorPlanRequest) HasTopLeftCorner() bool {
	if o != nil && !IsNil(o.TopLeftCorner) {
		return true
	}

	return false
}

// SetTopLeftCorner gets a reference to the given GetNetworkFloorPlans200ResponseInnerTopLeftCorner and assigns it to the TopLeftCorner field.
func (o *UpdateNetworkFloorPlanRequest) SetTopLeftCorner(v GetNetworkFloorPlans200ResponseInnerTopLeftCorner) {
	o.TopLeftCorner = &v
}

// GetTopRightCorner returns the TopRightCorner field value if set, zero value otherwise.
func (o *UpdateNetworkFloorPlanRequest) GetTopRightCorner() GetNetworkFloorPlans200ResponseInnerTopRightCorner {
	if o == nil || IsNil(o.TopRightCorner) {
		var ret GetNetworkFloorPlans200ResponseInnerTopRightCorner
		return ret
	}
	return *o.TopRightCorner
}

// GetTopRightCornerOk returns a tuple with the TopRightCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFloorPlanRequest) GetTopRightCornerOk() (*GetNetworkFloorPlans200ResponseInnerTopRightCorner, bool) {
	if o == nil || IsNil(o.TopRightCorner) {
		return nil, false
	}
	return o.TopRightCorner, true
}

// HasTopRightCorner returns a boolean if a field has been set.
func (o *UpdateNetworkFloorPlanRequest) HasTopRightCorner() bool {
	if o != nil && !IsNil(o.TopRightCorner) {
		return true
	}

	return false
}

// SetTopRightCorner gets a reference to the given GetNetworkFloorPlans200ResponseInnerTopRightCorner and assigns it to the TopRightCorner field.
func (o *UpdateNetworkFloorPlanRequest) SetTopRightCorner(v GetNetworkFloorPlans200ResponseInnerTopRightCorner) {
	o.TopRightCorner = &v
}

// GetImageContents returns the ImageContents field value if set, zero value otherwise.
func (o *UpdateNetworkFloorPlanRequest) GetImageContents() string {
	if o == nil || IsNil(o.ImageContents) {
		var ret string
		return ret
	}
	return *o.ImageContents
}

// GetImageContentsOk returns a tuple with the ImageContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFloorPlanRequest) GetImageContentsOk() (*string, bool) {
	if o == nil || IsNil(o.ImageContents) {
		return nil, false
	}
	return o.ImageContents, true
}

// HasImageContents returns a boolean if a field has been set.
func (o *UpdateNetworkFloorPlanRequest) HasImageContents() bool {
	if o != nil && !IsNil(o.ImageContents) {
		return true
	}

	return false
}

// SetImageContents gets a reference to the given string and assigns it to the ImageContents field.
func (o *UpdateNetworkFloorPlanRequest) SetImageContents(v string) {
	o.ImageContents = &v
}

func (o UpdateNetworkFloorPlanRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkFloorPlanRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Center) {
		toSerialize["center"] = o.Center
	}
	if !IsNil(o.BottomLeftCorner) {
		toSerialize["bottomLeftCorner"] = o.BottomLeftCorner
	}
	if !IsNil(o.BottomRightCorner) {
		toSerialize["bottomRightCorner"] = o.BottomRightCorner
	}
	if !IsNil(o.TopLeftCorner) {
		toSerialize["topLeftCorner"] = o.TopLeftCorner
	}
	if !IsNil(o.TopRightCorner) {
		toSerialize["topRightCorner"] = o.TopRightCorner
	}
	if !IsNil(o.ImageContents) {
		toSerialize["imageContents"] = o.ImageContents
	}
	return toSerialize, nil
}

type NullableUpdateNetworkFloorPlanRequest struct {
	value *UpdateNetworkFloorPlanRequest
	isSet bool
}

func (v NullableUpdateNetworkFloorPlanRequest) Get() *UpdateNetworkFloorPlanRequest {
	return v.value
}

func (v *NullableUpdateNetworkFloorPlanRequest) Set(val *UpdateNetworkFloorPlanRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkFloorPlanRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkFloorPlanRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkFloorPlanRequest(val *UpdateNetworkFloorPlanRequest) *NullableUpdateNetworkFloorPlanRequest {
	return &NullableUpdateNetworkFloorPlanRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkFloorPlanRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkFloorPlanRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


