# GetNetworkNetworkHealthChannelUtilization200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Device serial | [optional] 
**Model** | Pointer to **string** | Device model. | [optional] 
**Tags** | Pointer to **string** | Device tags. | [optional] 
**Wifi0** | Pointer to [**[]GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner**](GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner.md) | Channel utilization for first wifi radio of device. | [optional] 
**Wifi1** | Pointer to [**[]GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner**](GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner.md) | Channel utilization for second wifi radio of device. | [optional] 

## Methods

### NewGetNetworkNetworkHealthChannelUtilization200ResponseInner

`func NewGetNetworkNetworkHealthChannelUtilization200ResponseInner() *GetNetworkNetworkHealthChannelUtilization200ResponseInner`

NewGetNetworkNetworkHealthChannelUtilization200ResponseInner instantiates a new GetNetworkNetworkHealthChannelUtilization200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkNetworkHealthChannelUtilization200ResponseInnerWithDefaults

`func NewGetNetworkNetworkHealthChannelUtilization200ResponseInnerWithDefaults() *GetNetworkNetworkHealthChannelUtilization200ResponseInner`

NewGetNetworkNetworkHealthChannelUtilization200ResponseInnerWithDefaults instantiates a new GetNetworkNetworkHealthChannelUtilization200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetModel

`func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTags

`func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInner) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInner) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInner) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetWifi0

`func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInner) GetWifi0() []GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner`

GetWifi0 returns the Wifi0 field if non-nil, zero value otherwise.

### GetWifi0Ok

`func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInner) GetWifi0Ok() (*[]GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner, bool)`

GetWifi0Ok returns a tuple with the Wifi0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifi0

`func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInner) SetWifi0(v []GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner)`

SetWifi0 sets Wifi0 field to given value.

### HasWifi0

`func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInner) HasWifi0() bool`

HasWifi0 returns a boolean if a field has been set.

### GetWifi1

`func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInner) GetWifi1() []GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner`

GetWifi1 returns the Wifi1 field if non-nil, zero value otherwise.

### GetWifi1Ok

`func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInner) GetWifi1Ok() (*[]GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner, bool)`

GetWifi1Ok returns a tuple with the Wifi1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifi1

`func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInner) SetWifi1(v []GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner)`

SetWifi1 sets Wifi1 field to given value.

### HasWifi1

`func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInner) HasWifi1() bool`

HasWifi1 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


