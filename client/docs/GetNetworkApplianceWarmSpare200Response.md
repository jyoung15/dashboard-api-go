# GetNetworkApplianceWarmSpare200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Is the warm spare enabled | [optional] 
**PrimarySerial** | Pointer to **string** | Serial number of the primary appliance | [optional] 
**SpareSerial** | Pointer to **string** | Serial number of the warm spare appliance | [optional] 
**UplinkMode** | Pointer to **string** | Uplink mode, either virtual or public | [optional] 
**Wan1** | Pointer to [**GetNetworkApplianceWarmSpare200ResponseWan1**](GetNetworkApplianceWarmSpare200ResponseWan1.md) |  | [optional] 
**Wan2** | Pointer to [**GetNetworkApplianceWarmSpare200ResponseWan2**](GetNetworkApplianceWarmSpare200ResponseWan2.md) |  | [optional] 

## Methods

### NewGetNetworkApplianceWarmSpare200Response

`func NewGetNetworkApplianceWarmSpare200Response() *GetNetworkApplianceWarmSpare200Response`

NewGetNetworkApplianceWarmSpare200Response instantiates a new GetNetworkApplianceWarmSpare200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkApplianceWarmSpare200ResponseWithDefaults

`func NewGetNetworkApplianceWarmSpare200ResponseWithDefaults() *GetNetworkApplianceWarmSpare200Response`

NewGetNetworkApplianceWarmSpare200ResponseWithDefaults instantiates a new GetNetworkApplianceWarmSpare200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetNetworkApplianceWarmSpare200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkApplianceWarmSpare200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkApplianceWarmSpare200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkApplianceWarmSpare200Response) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPrimarySerial

`func (o *GetNetworkApplianceWarmSpare200Response) GetPrimarySerial() string`

GetPrimarySerial returns the PrimarySerial field if non-nil, zero value otherwise.

### GetPrimarySerialOk

`func (o *GetNetworkApplianceWarmSpare200Response) GetPrimarySerialOk() (*string, bool)`

GetPrimarySerialOk returns a tuple with the PrimarySerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySerial

`func (o *GetNetworkApplianceWarmSpare200Response) SetPrimarySerial(v string)`

SetPrimarySerial sets PrimarySerial field to given value.

### HasPrimarySerial

`func (o *GetNetworkApplianceWarmSpare200Response) HasPrimarySerial() bool`

HasPrimarySerial returns a boolean if a field has been set.

### GetSpareSerial

`func (o *GetNetworkApplianceWarmSpare200Response) GetSpareSerial() string`

GetSpareSerial returns the SpareSerial field if non-nil, zero value otherwise.

### GetSpareSerialOk

`func (o *GetNetworkApplianceWarmSpare200Response) GetSpareSerialOk() (*string, bool)`

GetSpareSerialOk returns a tuple with the SpareSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpareSerial

`func (o *GetNetworkApplianceWarmSpare200Response) SetSpareSerial(v string)`

SetSpareSerial sets SpareSerial field to given value.

### HasSpareSerial

`func (o *GetNetworkApplianceWarmSpare200Response) HasSpareSerial() bool`

HasSpareSerial returns a boolean if a field has been set.

### GetUplinkMode

`func (o *GetNetworkApplianceWarmSpare200Response) GetUplinkMode() string`

GetUplinkMode returns the UplinkMode field if non-nil, zero value otherwise.

### GetUplinkModeOk

`func (o *GetNetworkApplianceWarmSpare200Response) GetUplinkModeOk() (*string, bool)`

GetUplinkModeOk returns a tuple with the UplinkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkMode

`func (o *GetNetworkApplianceWarmSpare200Response) SetUplinkMode(v string)`

SetUplinkMode sets UplinkMode field to given value.

### HasUplinkMode

`func (o *GetNetworkApplianceWarmSpare200Response) HasUplinkMode() bool`

HasUplinkMode returns a boolean if a field has been set.

### GetWan1

`func (o *GetNetworkApplianceWarmSpare200Response) GetWan1() GetNetworkApplianceWarmSpare200ResponseWan1`

GetWan1 returns the Wan1 field if non-nil, zero value otherwise.

### GetWan1Ok

`func (o *GetNetworkApplianceWarmSpare200Response) GetWan1Ok() (*GetNetworkApplianceWarmSpare200ResponseWan1, bool)`

GetWan1Ok returns a tuple with the Wan1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWan1

`func (o *GetNetworkApplianceWarmSpare200Response) SetWan1(v GetNetworkApplianceWarmSpare200ResponseWan1)`

SetWan1 sets Wan1 field to given value.

### HasWan1

`func (o *GetNetworkApplianceWarmSpare200Response) HasWan1() bool`

HasWan1 returns a boolean if a field has been set.

### GetWan2

`func (o *GetNetworkApplianceWarmSpare200Response) GetWan2() GetNetworkApplianceWarmSpare200ResponseWan2`

GetWan2 returns the Wan2 field if non-nil, zero value otherwise.

### GetWan2Ok

`func (o *GetNetworkApplianceWarmSpare200Response) GetWan2Ok() (*GetNetworkApplianceWarmSpare200ResponseWan2, bool)`

GetWan2Ok returns a tuple with the Wan2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWan2

`func (o *GetNetworkApplianceWarmSpare200Response) SetWan2(v GetNetworkApplianceWarmSpare200ResponseWan2)`

SetWan2 sets Wan2 field to given value.

### HasWan2

`func (o *GetNetworkApplianceWarmSpare200Response) HasWan2() bool`

HasWan2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


