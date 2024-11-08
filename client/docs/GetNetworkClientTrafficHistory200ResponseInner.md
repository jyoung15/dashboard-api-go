# GetNetworkClientTrafficHistory200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | The timestamp of when the client was connected to an application | [optional] 
**Application** | Pointer to **string** | The name of the application the client is connected to | [optional] 
**Destination** | Pointer to **string** | The IP or web address the client is connected to | [optional] 
**Protocol** | Pointer to **string** | The client protocol | [optional] 
**Port** | Pointer to **int32** | The port the client is connected to | [optional] 
**Recv** | Pointer to **float32** | Usage received by the client | [optional] 
**Sent** | Pointer to **float32** | Usage sent by the client | [optional] 
**NumFlows** | Pointer to **int32** | The number of flows the client has | [optional] 
**ActiveSeconds** | Pointer to **int32** | The amount of seconds the client was active | [optional] 

## Methods

### NewGetNetworkClientTrafficHistory200ResponseInner

`func NewGetNetworkClientTrafficHistory200ResponseInner() *GetNetworkClientTrafficHistory200ResponseInner`

NewGetNetworkClientTrafficHistory200ResponseInner instantiates a new GetNetworkClientTrafficHistory200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkClientTrafficHistory200ResponseInnerWithDefaults

`func NewGetNetworkClientTrafficHistory200ResponseInnerWithDefaults() *GetNetworkClientTrafficHistory200ResponseInner`

NewGetNetworkClientTrafficHistory200ResponseInnerWithDefaults instantiates a new GetNetworkClientTrafficHistory200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *GetNetworkClientTrafficHistory200ResponseInner) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *GetNetworkClientTrafficHistory200ResponseInner) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *GetNetworkClientTrafficHistory200ResponseInner) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *GetNetworkClientTrafficHistory200ResponseInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetApplication

`func (o *GetNetworkClientTrafficHistory200ResponseInner) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *GetNetworkClientTrafficHistory200ResponseInner) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *GetNetworkClientTrafficHistory200ResponseInner) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *GetNetworkClientTrafficHistory200ResponseInner) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetDestination

`func (o *GetNetworkClientTrafficHistory200ResponseInner) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *GetNetworkClientTrafficHistory200ResponseInner) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *GetNetworkClientTrafficHistory200ResponseInner) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *GetNetworkClientTrafficHistory200ResponseInner) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetProtocol

`func (o *GetNetworkClientTrafficHistory200ResponseInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetNetworkClientTrafficHistory200ResponseInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetNetworkClientTrafficHistory200ResponseInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetNetworkClientTrafficHistory200ResponseInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPort

`func (o *GetNetworkClientTrafficHistory200ResponseInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetNetworkClientTrafficHistory200ResponseInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetNetworkClientTrafficHistory200ResponseInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetNetworkClientTrafficHistory200ResponseInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRecv

`func (o *GetNetworkClientTrafficHistory200ResponseInner) GetRecv() float32`

GetRecv returns the Recv field if non-nil, zero value otherwise.

### GetRecvOk

`func (o *GetNetworkClientTrafficHistory200ResponseInner) GetRecvOk() (*float32, bool)`

GetRecvOk returns a tuple with the Recv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecv

`func (o *GetNetworkClientTrafficHistory200ResponseInner) SetRecv(v float32)`

SetRecv sets Recv field to given value.

### HasRecv

`func (o *GetNetworkClientTrafficHistory200ResponseInner) HasRecv() bool`

HasRecv returns a boolean if a field has been set.

### GetSent

`func (o *GetNetworkClientTrafficHistory200ResponseInner) GetSent() float32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *GetNetworkClientTrafficHistory200ResponseInner) GetSentOk() (*float32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *GetNetworkClientTrafficHistory200ResponseInner) SetSent(v float32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *GetNetworkClientTrafficHistory200ResponseInner) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetNumFlows

`func (o *GetNetworkClientTrafficHistory200ResponseInner) GetNumFlows() int32`

GetNumFlows returns the NumFlows field if non-nil, zero value otherwise.

### GetNumFlowsOk

`func (o *GetNetworkClientTrafficHistory200ResponseInner) GetNumFlowsOk() (*int32, bool)`

GetNumFlowsOk returns a tuple with the NumFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFlows

`func (o *GetNetworkClientTrafficHistory200ResponseInner) SetNumFlows(v int32)`

SetNumFlows sets NumFlows field to given value.

### HasNumFlows

`func (o *GetNetworkClientTrafficHistory200ResponseInner) HasNumFlows() bool`

HasNumFlows returns a boolean if a field has been set.

### GetActiveSeconds

`func (o *GetNetworkClientTrafficHistory200ResponseInner) GetActiveSeconds() int32`

GetActiveSeconds returns the ActiveSeconds field if non-nil, zero value otherwise.

### GetActiveSecondsOk

`func (o *GetNetworkClientTrafficHistory200ResponseInner) GetActiveSecondsOk() (*int32, bool)`

GetActiveSecondsOk returns a tuple with the ActiveSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSeconds

`func (o *GetNetworkClientTrafficHistory200ResponseInner) SetActiveSeconds(v int32)`

SetActiveSeconds sets ActiveSeconds field to given value.

### HasActiveSeconds

`func (o *GetNetworkClientTrafficHistory200ResponseInner) HasActiveSeconds() bool`

HasActiveSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


