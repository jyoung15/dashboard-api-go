# DevicesSerialLiveToolsPingPostRequestMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Device serial number | [optional] 
**Target** | Pointer to **string** | IP address or FQDN to ping | [optional] 
**Count** | Pointer to **int32** | Number of pings to send | [optional] 

## Methods

### NewDevicesSerialLiveToolsPingPostRequestMessageRequest

`func NewDevicesSerialLiveToolsPingPostRequestMessageRequest() *DevicesSerialLiveToolsPingPostRequestMessageRequest`

NewDevicesSerialLiveToolsPingPostRequestMessageRequest instantiates a new DevicesSerialLiveToolsPingPostRequestMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialLiveToolsPingPostRequestMessageRequestWithDefaults

`func NewDevicesSerialLiveToolsPingPostRequestMessageRequestWithDefaults() *DevicesSerialLiveToolsPingPostRequestMessageRequest`

NewDevicesSerialLiveToolsPingPostRequestMessageRequestWithDefaults instantiates a new DevicesSerialLiveToolsPingPostRequestMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *DevicesSerialLiveToolsPingPostRequestMessageRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *DevicesSerialLiveToolsPingPostRequestMessageRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *DevicesSerialLiveToolsPingPostRequestMessageRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *DevicesSerialLiveToolsPingPostRequestMessageRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetTarget

`func (o *DevicesSerialLiveToolsPingPostRequestMessageRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DevicesSerialLiveToolsPingPostRequestMessageRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DevicesSerialLiveToolsPingPostRequestMessageRequest) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *DevicesSerialLiveToolsPingPostRequestMessageRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetCount

`func (o *DevicesSerialLiveToolsPingPostRequestMessageRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DevicesSerialLiveToolsPingPostRequestMessageRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DevicesSerialLiveToolsPingPostRequestMessageRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *DevicesSerialLiveToolsPingPostRequestMessageRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


