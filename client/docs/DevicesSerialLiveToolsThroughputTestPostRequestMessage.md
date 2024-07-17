# DevicesSerialLiveToolsThroughputTestPostRequestMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThroughputTestId** | Pointer to **string** | ID of throughput test job | [optional] 
**Url** | Pointer to **string** | GET this url to check the status of your throughput test request | [optional] 
**Status** | Pointer to **string** | Status of the throughput test request | [optional] 
**Result** | Pointer to [**CreateDeviceLiveToolsThroughputTest201ResponseResult**](CreateDeviceLiveToolsThroughputTest201ResponseResult.md) |  | [optional] 
**Request** | Pointer to [**CreateDeviceLiveToolsThroughputTest201ResponseRequest**](CreateDeviceLiveToolsThroughputTest201ResponseRequest.md) |  | [optional] 
**Error** | Pointer to **string** | Description of the error. | [optional] 

## Methods

### NewDevicesSerialLiveToolsThroughputTestPostRequestMessage

`func NewDevicesSerialLiveToolsThroughputTestPostRequestMessage() *DevicesSerialLiveToolsThroughputTestPostRequestMessage`

NewDevicesSerialLiveToolsThroughputTestPostRequestMessage instantiates a new DevicesSerialLiveToolsThroughputTestPostRequestMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialLiveToolsThroughputTestPostRequestMessageWithDefaults

`func NewDevicesSerialLiveToolsThroughputTestPostRequestMessageWithDefaults() *DevicesSerialLiveToolsThroughputTestPostRequestMessage`

NewDevicesSerialLiveToolsThroughputTestPostRequestMessageWithDefaults instantiates a new DevicesSerialLiveToolsThroughputTestPostRequestMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThroughputTestId

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetThroughputTestId() string`

GetThroughputTestId returns the ThroughputTestId field if non-nil, zero value otherwise.

### GetThroughputTestIdOk

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetThroughputTestIdOk() (*string, bool)`

GetThroughputTestIdOk returns a tuple with the ThroughputTestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputTestId

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) SetThroughputTestId(v string)`

SetThroughputTestId sets ThroughputTestId field to given value.

### HasThroughputTestId

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) HasThroughputTestId() bool`

HasThroughputTestId returns a boolean if a field has been set.

### GetUrl

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetStatus

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetResult() CreateDeviceLiveToolsThroughputTest201ResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetResultOk() (*CreateDeviceLiveToolsThroughputTest201ResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) SetResult(v CreateDeviceLiveToolsThroughputTest201ResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetRequest

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetRequest() CreateDeviceLiveToolsThroughputTest201ResponseRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetRequestOk() (*CreateDeviceLiveToolsThroughputTest201ResponseRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) SetRequest(v CreateDeviceLiveToolsThroughputTest201ResponseRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetError

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *DevicesSerialLiveToolsThroughputTestPostRequestMessage) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


