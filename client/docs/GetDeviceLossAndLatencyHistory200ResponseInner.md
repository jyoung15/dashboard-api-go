# GetDeviceLossAndLatencyHistory200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **time.Time** | Start time of the sample | [optional] 
**EndTime** | Pointer to **time.Time** | End time of the sample | [optional] 
**LossPercent** | Pointer to **float32** | Percentage of packets lost | [optional] 
**LatencyMs** | Pointer to **float32** | Latency in milliseconds | [optional] 
**Goodput** | Pointer to **int32** | Number of useful information bits delivered | [optional] 
**Jitter** | Pointer to **float32** | Jitter, in milliseconds | [optional] 

## Methods

### NewGetDeviceLossAndLatencyHistory200ResponseInner

`func NewGetDeviceLossAndLatencyHistory200ResponseInner() *GetDeviceLossAndLatencyHistory200ResponseInner`

NewGetDeviceLossAndLatencyHistory200ResponseInner instantiates a new GetDeviceLossAndLatencyHistory200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceLossAndLatencyHistory200ResponseInnerWithDefaults

`func NewGetDeviceLossAndLatencyHistory200ResponseInnerWithDefaults() *GetDeviceLossAndLatencyHistory200ResponseInner`

NewGetDeviceLossAndLatencyHistory200ResponseInnerWithDefaults instantiates a new GetDeviceLossAndLatencyHistory200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetLossPercent

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetLossPercent() float32`

GetLossPercent returns the LossPercent field if non-nil, zero value otherwise.

### GetLossPercentOk

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetLossPercentOk() (*float32, bool)`

GetLossPercentOk returns a tuple with the LossPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLossPercent

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) SetLossPercent(v float32)`

SetLossPercent sets LossPercent field to given value.

### HasLossPercent

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) HasLossPercent() bool`

HasLossPercent returns a boolean if a field has been set.

### GetLatencyMs

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetLatencyMs() float32`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetLatencyMsOk() (*float32, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) SetLatencyMs(v float32)`

SetLatencyMs sets LatencyMs field to given value.

### HasLatencyMs

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) HasLatencyMs() bool`

HasLatencyMs returns a boolean if a field has been set.

### GetGoodput

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetGoodput() int32`

GetGoodput returns the Goodput field if non-nil, zero value otherwise.

### GetGoodputOk

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetGoodputOk() (*int32, bool)`

GetGoodputOk returns a tuple with the Goodput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodput

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) SetGoodput(v int32)`

SetGoodput sets Goodput field to given value.

### HasGoodput

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) HasGoodput() bool`

HasGoodput returns a boolean if a field has been set.

### GetJitter

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetJitter() float32`

GetJitter returns the Jitter field if non-nil, zero value otherwise.

### GetJitterOk

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetJitterOk() (*float32, bool)`

GetJitterOk returns a tuple with the Jitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitter

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) SetJitter(v float32)`

SetJitter sets Jitter field to given value.

### HasJitter

`func (o *GetDeviceLossAndLatencyHistory200ResponseInner) HasJitter() bool`

HasJitter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


