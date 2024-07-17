# GetNetworkClientUsageHistory200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Received** | Pointer to **float32** | Usage received by the client on a given day | [optional] 
**Sent** | Pointer to **float32** | Usage sent by the client on a given day | [optional] 
**Ts** | Pointer to **time.Time** | The day&#39;s timestamp | [optional] 

## Methods

### NewGetNetworkClientUsageHistory200ResponseInner

`func NewGetNetworkClientUsageHistory200ResponseInner() *GetNetworkClientUsageHistory200ResponseInner`

NewGetNetworkClientUsageHistory200ResponseInner instantiates a new GetNetworkClientUsageHistory200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkClientUsageHistory200ResponseInnerWithDefaults

`func NewGetNetworkClientUsageHistory200ResponseInnerWithDefaults() *GetNetworkClientUsageHistory200ResponseInner`

NewGetNetworkClientUsageHistory200ResponseInnerWithDefaults instantiates a new GetNetworkClientUsageHistory200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceived

`func (o *GetNetworkClientUsageHistory200ResponseInner) GetReceived() float32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *GetNetworkClientUsageHistory200ResponseInner) GetReceivedOk() (*float32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *GetNetworkClientUsageHistory200ResponseInner) SetReceived(v float32)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *GetNetworkClientUsageHistory200ResponseInner) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetSent

`func (o *GetNetworkClientUsageHistory200ResponseInner) GetSent() float32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *GetNetworkClientUsageHistory200ResponseInner) GetSentOk() (*float32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *GetNetworkClientUsageHistory200ResponseInner) SetSent(v float32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *GetNetworkClientUsageHistory200ResponseInner) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetTs

`func (o *GetNetworkClientUsageHistory200ResponseInner) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetNetworkClientUsageHistory200ResponseInner) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetNetworkClientUsageHistory200ResponseInner) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetNetworkClientUsageHistory200ResponseInner) HasTs() bool`

HasTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


