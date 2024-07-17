# UpdateDeviceWirelessElectronicShelfLabelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Desired ESL channel for the device, or &#39;Auto&#39; (case insensitive) to use the recommended channel | [optional] 
**Enabled** | Pointer to **bool** | Turn ESL features on and off for this device | [optional] 

## Methods

### NewUpdateDeviceWirelessElectronicShelfLabelRequest

`func NewUpdateDeviceWirelessElectronicShelfLabelRequest() *UpdateDeviceWirelessElectronicShelfLabelRequest`

NewUpdateDeviceWirelessElectronicShelfLabelRequest instantiates a new UpdateDeviceWirelessElectronicShelfLabelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceWirelessElectronicShelfLabelRequestWithDefaults

`func NewUpdateDeviceWirelessElectronicShelfLabelRequestWithDefaults() *UpdateDeviceWirelessElectronicShelfLabelRequest`

NewUpdateDeviceWirelessElectronicShelfLabelRequestWithDefaults instantiates a new UpdateDeviceWirelessElectronicShelfLabelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *UpdateDeviceWirelessElectronicShelfLabelRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *UpdateDeviceWirelessElectronicShelfLabelRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *UpdateDeviceWirelessElectronicShelfLabelRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *UpdateDeviceWirelessElectronicShelfLabelRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateDeviceWirelessElectronicShelfLabelRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateDeviceWirelessElectronicShelfLabelRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateDeviceWirelessElectronicShelfLabelRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateDeviceWirelessElectronicShelfLabelRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


