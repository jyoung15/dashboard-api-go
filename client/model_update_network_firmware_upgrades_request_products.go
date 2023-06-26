/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateNetworkFirmwareUpgradesRequestProducts Contains information about the network to update
type UpdateNetworkFirmwareUpgradesRequestProducts struct {
	Wireless *UpdateNetworkFirmwareUpgradesRequestProductsWireless `json:"wireless,omitempty"`
	Appliance *UpdateNetworkFirmwareUpgradesRequestProductsWireless `json:"appliance,omitempty"`
	Switch *UpdateNetworkFirmwareUpgradesRequestProductsWireless `json:"switch,omitempty"`
	Camera *UpdateNetworkFirmwareUpgradesRequestProductsWireless `json:"camera,omitempty"`
	CellularGateway *UpdateNetworkFirmwareUpgradesRequestProductsWireless `json:"cellularGateway,omitempty"`
	Sensor *UpdateNetworkFirmwareUpgradesRequestProductsWireless `json:"sensor,omitempty"`
	CloudGateway *UpdateNetworkFirmwareUpgradesRequestProductsWireless `json:"cloudGateway,omitempty"`
	SwitchCatalyst *UpdateNetworkFirmwareUpgradesRequestProductsWireless `json:"switchCatalyst,omitempty"`
}

// NewUpdateNetworkFirmwareUpgradesRequestProducts instantiates a new UpdateNetworkFirmwareUpgradesRequestProducts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkFirmwareUpgradesRequestProducts() *UpdateNetworkFirmwareUpgradesRequestProducts {
	this := UpdateNetworkFirmwareUpgradesRequestProducts{}
	return &this
}

// NewUpdateNetworkFirmwareUpgradesRequestProductsWithDefaults instantiates a new UpdateNetworkFirmwareUpgradesRequestProducts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkFirmwareUpgradesRequestProductsWithDefaults() *UpdateNetworkFirmwareUpgradesRequestProducts {
	this := UpdateNetworkFirmwareUpgradesRequestProducts{}
	return &this
}

// GetWireless returns the Wireless field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetWireless() UpdateNetworkFirmwareUpgradesRequestProductsWireless {
	if o == nil || isNil(o.Wireless) {
		var ret UpdateNetworkFirmwareUpgradesRequestProductsWireless
		return ret
	}
	return *o.Wireless
}

// GetWirelessOk returns a tuple with the Wireless field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetWirelessOk() (*UpdateNetworkFirmwareUpgradesRequestProductsWireless, bool) {
	if o == nil || isNil(o.Wireless) {
    return nil, false
	}
	return o.Wireless, true
}

// HasWireless returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) HasWireless() bool {
	if o != nil && !isNil(o.Wireless) {
		return true
	}

	return false
}

// SetWireless gets a reference to the given UpdateNetworkFirmwareUpgradesRequestProductsWireless and assigns it to the Wireless field.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) SetWireless(v UpdateNetworkFirmwareUpgradesRequestProductsWireless) {
	o.Wireless = &v
}

// GetAppliance returns the Appliance field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetAppliance() UpdateNetworkFirmwareUpgradesRequestProductsWireless {
	if o == nil || isNil(o.Appliance) {
		var ret UpdateNetworkFirmwareUpgradesRequestProductsWireless
		return ret
	}
	return *o.Appliance
}

// GetApplianceOk returns a tuple with the Appliance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetApplianceOk() (*UpdateNetworkFirmwareUpgradesRequestProductsWireless, bool) {
	if o == nil || isNil(o.Appliance) {
    return nil, false
	}
	return o.Appliance, true
}

// HasAppliance returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) HasAppliance() bool {
	if o != nil && !isNil(o.Appliance) {
		return true
	}

	return false
}

// SetAppliance gets a reference to the given UpdateNetworkFirmwareUpgradesRequestProductsWireless and assigns it to the Appliance field.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) SetAppliance(v UpdateNetworkFirmwareUpgradesRequestProductsWireless) {
	o.Appliance = &v
}

// GetSwitch returns the Switch field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetSwitch() UpdateNetworkFirmwareUpgradesRequestProductsWireless {
	if o == nil || isNil(o.Switch) {
		var ret UpdateNetworkFirmwareUpgradesRequestProductsWireless
		return ret
	}
	return *o.Switch
}

// GetSwitchOk returns a tuple with the Switch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetSwitchOk() (*UpdateNetworkFirmwareUpgradesRequestProductsWireless, bool) {
	if o == nil || isNil(o.Switch) {
    return nil, false
	}
	return o.Switch, true
}

// HasSwitch returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) HasSwitch() bool {
	if o != nil && !isNil(o.Switch) {
		return true
	}

	return false
}

// SetSwitch gets a reference to the given UpdateNetworkFirmwareUpgradesRequestProductsWireless and assigns it to the Switch field.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) SetSwitch(v UpdateNetworkFirmwareUpgradesRequestProductsWireless) {
	o.Switch = &v
}

// GetCamera returns the Camera field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetCamera() UpdateNetworkFirmwareUpgradesRequestProductsWireless {
	if o == nil || isNil(o.Camera) {
		var ret UpdateNetworkFirmwareUpgradesRequestProductsWireless
		return ret
	}
	return *o.Camera
}

// GetCameraOk returns a tuple with the Camera field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetCameraOk() (*UpdateNetworkFirmwareUpgradesRequestProductsWireless, bool) {
	if o == nil || isNil(o.Camera) {
    return nil, false
	}
	return o.Camera, true
}

// HasCamera returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) HasCamera() bool {
	if o != nil && !isNil(o.Camera) {
		return true
	}

	return false
}

// SetCamera gets a reference to the given UpdateNetworkFirmwareUpgradesRequestProductsWireless and assigns it to the Camera field.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) SetCamera(v UpdateNetworkFirmwareUpgradesRequestProductsWireless) {
	o.Camera = &v
}

// GetCellularGateway returns the CellularGateway field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetCellularGateway() UpdateNetworkFirmwareUpgradesRequestProductsWireless {
	if o == nil || isNil(o.CellularGateway) {
		var ret UpdateNetworkFirmwareUpgradesRequestProductsWireless
		return ret
	}
	return *o.CellularGateway
}

// GetCellularGatewayOk returns a tuple with the CellularGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetCellularGatewayOk() (*UpdateNetworkFirmwareUpgradesRequestProductsWireless, bool) {
	if o == nil || isNil(o.CellularGateway) {
    return nil, false
	}
	return o.CellularGateway, true
}

// HasCellularGateway returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) HasCellularGateway() bool {
	if o != nil && !isNil(o.CellularGateway) {
		return true
	}

	return false
}

// SetCellularGateway gets a reference to the given UpdateNetworkFirmwareUpgradesRequestProductsWireless and assigns it to the CellularGateway field.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) SetCellularGateway(v UpdateNetworkFirmwareUpgradesRequestProductsWireless) {
	o.CellularGateway = &v
}

// GetSensor returns the Sensor field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetSensor() UpdateNetworkFirmwareUpgradesRequestProductsWireless {
	if o == nil || isNil(o.Sensor) {
		var ret UpdateNetworkFirmwareUpgradesRequestProductsWireless
		return ret
	}
	return *o.Sensor
}

// GetSensorOk returns a tuple with the Sensor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetSensorOk() (*UpdateNetworkFirmwareUpgradesRequestProductsWireless, bool) {
	if o == nil || isNil(o.Sensor) {
    return nil, false
	}
	return o.Sensor, true
}

// HasSensor returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) HasSensor() bool {
	if o != nil && !isNil(o.Sensor) {
		return true
	}

	return false
}

// SetSensor gets a reference to the given UpdateNetworkFirmwareUpgradesRequestProductsWireless and assigns it to the Sensor field.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) SetSensor(v UpdateNetworkFirmwareUpgradesRequestProductsWireless) {
	o.Sensor = &v
}

// GetCloudGateway returns the CloudGateway field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetCloudGateway() UpdateNetworkFirmwareUpgradesRequestProductsWireless {
	if o == nil || isNil(o.CloudGateway) {
		var ret UpdateNetworkFirmwareUpgradesRequestProductsWireless
		return ret
	}
	return *o.CloudGateway
}

// GetCloudGatewayOk returns a tuple with the CloudGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetCloudGatewayOk() (*UpdateNetworkFirmwareUpgradesRequestProductsWireless, bool) {
	if o == nil || isNil(o.CloudGateway) {
    return nil, false
	}
	return o.CloudGateway, true
}

// HasCloudGateway returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) HasCloudGateway() bool {
	if o != nil && !isNil(o.CloudGateway) {
		return true
	}

	return false
}

// SetCloudGateway gets a reference to the given UpdateNetworkFirmwareUpgradesRequestProductsWireless and assigns it to the CloudGateway field.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) SetCloudGateway(v UpdateNetworkFirmwareUpgradesRequestProductsWireless) {
	o.CloudGateway = &v
}

// GetSwitchCatalyst returns the SwitchCatalyst field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetSwitchCatalyst() UpdateNetworkFirmwareUpgradesRequestProductsWireless {
	if o == nil || isNil(o.SwitchCatalyst) {
		var ret UpdateNetworkFirmwareUpgradesRequestProductsWireless
		return ret
	}
	return *o.SwitchCatalyst
}

// GetSwitchCatalystOk returns a tuple with the SwitchCatalyst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) GetSwitchCatalystOk() (*UpdateNetworkFirmwareUpgradesRequestProductsWireless, bool) {
	if o == nil || isNil(o.SwitchCatalyst) {
    return nil, false
	}
	return o.SwitchCatalyst, true
}

// HasSwitchCatalyst returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) HasSwitchCatalyst() bool {
	if o != nil && !isNil(o.SwitchCatalyst) {
		return true
	}

	return false
}

// SetSwitchCatalyst gets a reference to the given UpdateNetworkFirmwareUpgradesRequestProductsWireless and assigns it to the SwitchCatalyst field.
func (o *UpdateNetworkFirmwareUpgradesRequestProducts) SetSwitchCatalyst(v UpdateNetworkFirmwareUpgradesRequestProductsWireless) {
	o.SwitchCatalyst = &v
}

func (o UpdateNetworkFirmwareUpgradesRequestProducts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Wireless) {
		toSerialize["wireless"] = o.Wireless
	}
	if !isNil(o.Appliance) {
		toSerialize["appliance"] = o.Appliance
	}
	if !isNil(o.Switch) {
		toSerialize["switch"] = o.Switch
	}
	if !isNil(o.Camera) {
		toSerialize["camera"] = o.Camera
	}
	if !isNil(o.CellularGateway) {
		toSerialize["cellularGateway"] = o.CellularGateway
	}
	if !isNil(o.Sensor) {
		toSerialize["sensor"] = o.Sensor
	}
	if !isNil(o.CloudGateway) {
		toSerialize["cloudGateway"] = o.CloudGateway
	}
	if !isNil(o.SwitchCatalyst) {
		toSerialize["switchCatalyst"] = o.SwitchCatalyst
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkFirmwareUpgradesRequestProducts struct {
	value *UpdateNetworkFirmwareUpgradesRequestProducts
	isSet bool
}

func (v NullableUpdateNetworkFirmwareUpgradesRequestProducts) Get() *UpdateNetworkFirmwareUpgradesRequestProducts {
	return v.value
}

func (v *NullableUpdateNetworkFirmwareUpgradesRequestProducts) Set(val *UpdateNetworkFirmwareUpgradesRequestProducts) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkFirmwareUpgradesRequestProducts) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkFirmwareUpgradesRequestProducts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkFirmwareUpgradesRequestProducts(val *UpdateNetworkFirmwareUpgradesRequestProducts) *NullableUpdateNetworkFirmwareUpgradesRequestProducts {
	return &NullableUpdateNetworkFirmwareUpgradesRequestProducts{value: val, isSet: true}
}

func (v NullableUpdateNetworkFirmwareUpgradesRequestProducts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkFirmwareUpgradesRequestProducts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


