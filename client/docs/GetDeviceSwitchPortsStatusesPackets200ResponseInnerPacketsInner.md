# GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Desc** | Pointer to **string** | The type of packets being counted. | [optional] 
**Total** | Pointer to **int32** | The total count of sent and received packets. | [optional] 
**Sent** | Pointer to **int32** | The total count of packets sent by the switch during the timespan. | [optional] 
**Recv** | Pointer to **int32** | The total count of packets received by the switch during the timespan. | [optional] 
**RatePerSec** | Pointer to [**GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInnerRatePerSec**](GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInnerRatePerSec.md) |  | [optional] 

## Methods

### NewGetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner

`func NewGetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner() *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner`

NewGetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner instantiates a new GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInnerWithDefaults

`func NewGetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInnerWithDefaults() *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner`

NewGetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInnerWithDefaults instantiates a new GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesc

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetTotal

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetSent

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetRecv

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner) GetRecv() int32`

GetRecv returns the Recv field if non-nil, zero value otherwise.

### GetRecvOk

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner) GetRecvOk() (*int32, bool)`

GetRecvOk returns a tuple with the Recv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecv

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner) SetRecv(v int32)`

SetRecv sets Recv field to given value.

### HasRecv

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner) HasRecv() bool`

HasRecv returns a boolean if a field has been set.

### GetRatePerSec

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner) GetRatePerSec() GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInnerRatePerSec`

GetRatePerSec returns the RatePerSec field if non-nil, zero value otherwise.

### GetRatePerSecOk

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner) GetRatePerSecOk() (*GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInnerRatePerSec, bool)`

GetRatePerSecOk returns a tuple with the RatePerSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePerSec

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner) SetRatePerSec(v GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInnerRatePerSec)`

SetRatePerSec sets RatePerSec field to given value.

### HasRatePerSec

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner) HasRatePerSec() bool`

HasRatePerSec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


