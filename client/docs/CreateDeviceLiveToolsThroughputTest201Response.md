# CreateDeviceLiveToolsThroughputTest201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThroughputTestId** | Pointer to **string** | ID of throughput test job | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your throughput test request | [optional] 
**Status** | Pointer to **string** | Status of the throughput test request | [optional] 
**Result** | Pointer to [**CreateDeviceLiveToolsThroughputTest201ResponseResult**](CreateDeviceLiveToolsThroughputTest201ResponseResult.md) |  | [optional] 
**Request** | Pointer to [**CreateDeviceLiveToolsThroughputTest201ResponseRequest**](CreateDeviceLiveToolsThroughputTest201ResponseRequest.md) |  | [optional] 
**Error** | Pointer to **string** | Description of the error. | [optional] 
**Callback** | Pointer to [**CreateDeviceLiveToolsArpTable201ResponseCallback**](CreateDeviceLiveToolsArpTable201ResponseCallback.md) |  | [optional] 

## Methods

### NewCreateDeviceLiveToolsThroughputTest201Response

`func NewCreateDeviceLiveToolsThroughputTest201Response() *CreateDeviceLiveToolsThroughputTest201Response`

NewCreateDeviceLiveToolsThroughputTest201Response instantiates a new CreateDeviceLiveToolsThroughputTest201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceLiveToolsThroughputTest201ResponseWithDefaults

`func NewCreateDeviceLiveToolsThroughputTest201ResponseWithDefaults() *CreateDeviceLiveToolsThroughputTest201Response`

NewCreateDeviceLiveToolsThroughputTest201ResponseWithDefaults instantiates a new CreateDeviceLiveToolsThroughputTest201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThroughputTestId

`func (o *CreateDeviceLiveToolsThroughputTest201Response) GetThroughputTestId() string`

GetThroughputTestId returns the ThroughputTestId field if non-nil, zero value otherwise.

### GetThroughputTestIdOk

`func (o *CreateDeviceLiveToolsThroughputTest201Response) GetThroughputTestIdOk() (*string, bool)`

GetThroughputTestIdOk returns a tuple with the ThroughputTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputTestId

`func (o *CreateDeviceLiveToolsThroughputTest201Response) SetThroughputTestId(v string)`

SetThroughputTestId sets ThroughputTestId field to given value.

### HasThroughputTestId

`func (o *CreateDeviceLiveToolsThroughputTest201Response) HasThroughputTestId() bool`

HasThroughputTestId returns a boolean if a field has been set.

### GetUrl

`func (o *CreateDeviceLiveToolsThroughputTest201Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateDeviceLiveToolsThroughputTest201Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateDeviceLiveToolsThroughputTest201Response) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateDeviceLiveToolsThroughputTest201Response) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *CreateDeviceLiveToolsThroughputTest201Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateDeviceLiveToolsThroughputTest201Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateDeviceLiveToolsThroughputTest201Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateDeviceLiveToolsThroughputTest201Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *CreateDeviceLiveToolsThroughputTest201Response) GetResult() CreateDeviceLiveToolsThroughputTest201ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CreateDeviceLiveToolsThroughputTest201Response) GetResultOk() (*CreateDeviceLiveToolsThroughputTest201ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CreateDeviceLiveToolsThroughputTest201Response) SetResult(v CreateDeviceLiveToolsThroughputTest201ResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *CreateDeviceLiveToolsThroughputTest201Response) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRequest

`func (o *CreateDeviceLiveToolsThroughputTest201Response) GetRequest() CreateDeviceLiveToolsThroughputTest201ResponseRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *CreateDeviceLiveToolsThroughputTest201Response) GetRequestOk() (*CreateDeviceLiveToolsThroughputTest201ResponseRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *CreateDeviceLiveToolsThroughputTest201Response) SetRequest(v CreateDeviceLiveToolsThroughputTest201ResponseRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *CreateDeviceLiveToolsThroughputTest201Response) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetError

`func (o *CreateDeviceLiveToolsThroughputTest201Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CreateDeviceLiveToolsThroughputTest201Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CreateDeviceLiveToolsThroughputTest201Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CreateDeviceLiveToolsThroughputTest201Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCallback

`func (o *CreateDeviceLiveToolsThroughputTest201Response) GetCallback() CreateDeviceLiveToolsArpTable201ResponseCallback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *CreateDeviceLiveToolsThroughputTest201Response) GetCallbackOk() (*CreateDeviceLiveToolsArpTable201ResponseCallback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *CreateDeviceLiveToolsThroughputTest201Response) SetCallback(v CreateDeviceLiveToolsArpTable201ResponseCallback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *CreateDeviceLiveToolsThroughputTest201Response) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


