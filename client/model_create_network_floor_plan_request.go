/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateNetworkFloorPlanRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkFloorPlanRequest{}

// CreateNetworkFloorPlanRequest struct for CreateNetworkFloorPlanRequest
type CreateNetworkFloorPlanRequest struct {
	// The name of your floor plan.
	Name string `json:"name"`
	Center *GetNetworkFloorPlans200ResponseInnerCenter `json:"center,omitempty"`
	BottomLeftCorner *GetNetworkFloorPlans200ResponseInnerBottomLeftCorner `json:"bottomLeftCorner,omitempty"`
	BottomRightCorner *GetNetworkFloorPlans200ResponseInnerBottomRightCorner `json:"bottomRightCorner,omitempty"`
	TopLeftCorner *GetNetworkFloorPlans200ResponseInnerTopLeftCorner `json:"topLeftCorner,omitempty"`
	TopRightCorner *GetNetworkFloorPlans200ResponseInnerTopRightCorner `json:"topRightCorner,omitempty"`
	// The file contents (a base 64 encoded string) of your image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in.
	ImageContents string `json:"imageContents"`
}

type _CreateNetworkFloorPlanRequest CreateNetworkFloorPlanRequest

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
func (o *CreateNetworkFloorPlanRequest) GetCenter() GetNetworkFloorPlans200ResponseInnerCenter {
	if o == nil || IsNil(o.Center) {
		var ret GetNetworkFloorPlans200ResponseInnerCenter
		return ret
	}
	return *o.Center
}

// GetCenterOk returns a tuple with the Center field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFloorPlanRequest) GetCenterOk() (*GetNetworkFloorPlans200ResponseInnerCenter, bool) {
	if o == nil || IsNil(o.Center) {
		return nil, false
	}
	return o.Center, true
}

// HasCenter returns a boolean if a field has been set.
func (o *CreateNetworkFloorPlanRequest) HasCenter() bool {
	if o != nil && !IsNil(o.Center) {
		return true
	}

	return false
}

// SetCenter gets a reference to the given GetNetworkFloorPlans200ResponseInnerCenter and assigns it to the Center field.
func (o *CreateNetworkFloorPlanRequest) SetCenter(v GetNetworkFloorPlans200ResponseInnerCenter) {
	o.Center = &v
}

// GetBottomLeftCorner returns the BottomLeftCorner field value if set, zero value otherwise.
func (o *CreateNetworkFloorPlanRequest) GetBottomLeftCorner() GetNetworkFloorPlans200ResponseInnerBottomLeftCorner {
	if o == nil || IsNil(o.BottomLeftCorner) {
		var ret GetNetworkFloorPlans200ResponseInnerBottomLeftCorner
		return ret
	}
	return *o.BottomLeftCorner
}

// GetBottomLeftCornerOk returns a tuple with the BottomLeftCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFloorPlanRequest) GetBottomLeftCornerOk() (*GetNetworkFloorPlans200ResponseInnerBottomLeftCorner, bool) {
	if o == nil || IsNil(o.BottomLeftCorner) {
		return nil, false
	}
	return o.BottomLeftCorner, true
}

// HasBottomLeftCorner returns a boolean if a field has been set.
func (o *CreateNetworkFloorPlanRequest) HasBottomLeftCorner() bool {
	if o != nil && !IsNil(o.BottomLeftCorner) {
		return true
	}

	return false
}

// SetBottomLeftCorner gets a reference to the given GetNetworkFloorPlans200ResponseInnerBottomLeftCorner and assigns it to the BottomLeftCorner field.
func (o *CreateNetworkFloorPlanRequest) SetBottomLeftCorner(v GetNetworkFloorPlans200ResponseInnerBottomLeftCorner) {
	o.BottomLeftCorner = &v
}

// GetBottomRightCorner returns the BottomRightCorner field value if set, zero value otherwise.
func (o *CreateNetworkFloorPlanRequest) GetBottomRightCorner() GetNetworkFloorPlans200ResponseInnerBottomRightCorner {
	if o == nil || IsNil(o.BottomRightCorner) {
		var ret GetNetworkFloorPlans200ResponseInnerBottomRightCorner
		return ret
	}
	return *o.BottomRightCorner
}

// GetBottomRightCornerOk returns a tuple with the BottomRightCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFloorPlanRequest) GetBottomRightCornerOk() (*GetNetworkFloorPlans200ResponseInnerBottomRightCorner, bool) {
	if o == nil || IsNil(o.BottomRightCorner) {
		return nil, false
	}
	return o.BottomRightCorner, true
}

// HasBottomRightCorner returns a boolean if a field has been set.
func (o *CreateNetworkFloorPlanRequest) HasBottomRightCorner() bool {
	if o != nil && !IsNil(o.BottomRightCorner) {
		return true
	}

	return false
}

// SetBottomRightCorner gets a reference to the given GetNetworkFloorPlans200ResponseInnerBottomRightCorner and assigns it to the BottomRightCorner field.
func (o *CreateNetworkFloorPlanRequest) SetBottomRightCorner(v GetNetworkFloorPlans200ResponseInnerBottomRightCorner) {
	o.BottomRightCorner = &v
}

// GetTopLeftCorner returns the TopLeftCorner field value if set, zero value otherwise.
func (o *CreateNetworkFloorPlanRequest) GetTopLeftCorner() GetNetworkFloorPlans200ResponseInnerTopLeftCorner {
	if o == nil || IsNil(o.TopLeftCorner) {
		var ret GetNetworkFloorPlans200ResponseInnerTopLeftCorner
		return ret
	}
	return *o.TopLeftCorner
}

// GetTopLeftCornerOk returns a tuple with the TopLeftCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFloorPlanRequest) GetTopLeftCornerOk() (*GetNetworkFloorPlans200ResponseInnerTopLeftCorner, bool) {
	if o == nil || IsNil(o.TopLeftCorner) {
		return nil, false
	}
	return o.TopLeftCorner, true
}

// HasTopLeftCorner returns a boolean if a field has been set.
func (o *CreateNetworkFloorPlanRequest) HasTopLeftCorner() bool {
	if o != nil && !IsNil(o.TopLeftCorner) {
		return true
	}

	return false
}

// SetTopLeftCorner gets a reference to the given GetNetworkFloorPlans200ResponseInnerTopLeftCorner and assigns it to the TopLeftCorner field.
func (o *CreateNetworkFloorPlanRequest) SetTopLeftCorner(v GetNetworkFloorPlans200ResponseInnerTopLeftCorner) {
	o.TopLeftCorner = &v
}

// GetTopRightCorner returns the TopRightCorner field value if set, zero value otherwise.
func (o *CreateNetworkFloorPlanRequest) GetTopRightCorner() GetNetworkFloorPlans200ResponseInnerTopRightCorner {
	if o == nil || IsNil(o.TopRightCorner) {
		var ret GetNetworkFloorPlans200ResponseInnerTopRightCorner
		return ret
	}
	return *o.TopRightCorner
}

// GetTopRightCornerOk returns a tuple with the TopRightCorner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFloorPlanRequest) GetTopRightCornerOk() (*GetNetworkFloorPlans200ResponseInnerTopRightCorner, bool) {
	if o == nil || IsNil(o.TopRightCorner) {
		return nil, false
	}
	return o.TopRightCorner, true
}

// HasTopRightCorner returns a boolean if a field has been set.
func (o *CreateNetworkFloorPlanRequest) HasTopRightCorner() bool {
	if o != nil && !IsNil(o.TopRightCorner) {
		return true
	}

	return false
}

// SetTopRightCorner gets a reference to the given GetNetworkFloorPlans200ResponseInnerTopRightCorner and assigns it to the TopRightCorner field.
func (o *CreateNetworkFloorPlanRequest) SetTopRightCorner(v GetNetworkFloorPlans200ResponseInnerTopRightCorner) {
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
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkFloorPlanRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
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
	toSerialize["imageContents"] = o.ImageContents
	return toSerialize, nil
}

func (o *CreateNetworkFloorPlanRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"imageContents",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateNetworkFloorPlanRequest := _CreateNetworkFloorPlanRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNetworkFloorPlanRequest)

	if err != nil {
		return err
	}

	*o = CreateNetworkFloorPlanRequest(varCreateNetworkFloorPlanRequest)

	return err
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


