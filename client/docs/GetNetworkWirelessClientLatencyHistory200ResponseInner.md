# GetNetworkWirelessClientLatencyHistory200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**T0** | Pointer to **int32** | The latency history bucket start time in seconds | [optional] 
**T1** | Pointer to **int32** | The latency history bucket end time in seconds | [optional] 
**LatencyBinsByCategory** | Pointer to [**GetNetworkWirelessClientLatencyHistory200ResponseInnerLatencyBinsByCategory**](GetNetworkWirelessClientLatencyHistory200ResponseInnerLatencyBinsByCategory.md) |  | [optional] 

## Methods

### NewGetNetworkWirelessClientLatencyHistory200ResponseInner

`func NewGetNetworkWirelessClientLatencyHistory200ResponseInner() *GetNetworkWirelessClientLatencyHistory200ResponseInner`

NewGetNetworkWirelessClientLatencyHistory200ResponseInner instantiates a new GetNetworkWirelessClientLatencyHistory200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWirelessClientLatencyHistory200ResponseInnerWithDefaults

`func NewGetNetworkWirelessClientLatencyHistory200ResponseInnerWithDefaults() *GetNetworkWirelessClientLatencyHistory200ResponseInner`

NewGetNetworkWirelessClientLatencyHistory200ResponseInnerWithDefaults instantiates a new GetNetworkWirelessClientLatencyHistory200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetT0

`func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) GetT0() int32`

GetT0 returns the T0 field if non-nil, zero value otherwise.

### GetT0Ok

`func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) GetT0Ok() (*int32, bool)`

GetT0Ok returns a tuple with the T0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT0

`func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) SetT0(v int32)`

SetT0 sets T0 field to given value.

### HasT0

`func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) HasT0() bool`

HasT0 returns a boolean if a field has been set.

### GetT1

`func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) GetT1() int32`

GetT1 returns the T1 field if non-nil, zero value otherwise.

### GetT1Ok

`func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) GetT1Ok() (*int32, bool)`

GetT1Ok returns a tuple with the T1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT1

`func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) SetT1(v int32)`

SetT1 sets T1 field to given value.

### HasT1

`func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) HasT1() bool`

HasT1 returns a boolean if a field has been set.

### GetLatencyBinsByCategory

`func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) GetLatencyBinsByCategory() GetNetworkWirelessClientLatencyHistory200ResponseInnerLatencyBinsByCategory`

GetLatencyBinsByCategory returns the LatencyBinsByCategory field if non-nil, zero value otherwise.

### GetLatencyBinsByCategoryOk

`func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) GetLatencyBinsByCategoryOk() (*GetNetworkWirelessClientLatencyHistory200ResponseInnerLatencyBinsByCategory, bool)`

GetLatencyBinsByCategoryOk returns a tuple with the LatencyBinsByCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyBinsByCategory

`func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) SetLatencyBinsByCategory(v GetNetworkWirelessClientLatencyHistory200ResponseInnerLatencyBinsByCategory)`

SetLatencyBinsByCategory sets LatencyBinsByCategory field to given value.

### HasLatencyBinsByCategory

`func (o *GetNetworkWirelessClientLatencyHistory200ResponseInner) HasLatencyBinsByCategory() bool`

HasLatencyBinsByCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


