# CreateDeviceLiveToolsPingDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count parameter to pass to ping. [1..5], default 5 | [optional] 
**Callback** | Pointer to [**CreateDeviceLiveToolsArpTableRequestCallback**](CreateDeviceLiveToolsArpTableRequestCallback.md) |  | [optional] 

## Methods

### NewCreateDeviceLiveToolsPingDeviceRequest

`func NewCreateDeviceLiveToolsPingDeviceRequest() *CreateDeviceLiveToolsPingDeviceRequest`

NewCreateDeviceLiveToolsPingDeviceRequest instantiates a new CreateDeviceLiveToolsPingDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceLiveToolsPingDeviceRequestWithDefaults

`func NewCreateDeviceLiveToolsPingDeviceRequestWithDefaults() *CreateDeviceLiveToolsPingDeviceRequest`

NewCreateDeviceLiveToolsPingDeviceRequestWithDefaults instantiates a new CreateDeviceLiveToolsPingDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CreateDeviceLiveToolsPingDeviceRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CreateDeviceLiveToolsPingDeviceRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CreateDeviceLiveToolsPingDeviceRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CreateDeviceLiveToolsPingDeviceRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCallback

`func (o *CreateDeviceLiveToolsPingDeviceRequest) GetCallback() CreateDeviceLiveToolsArpTableRequestCallback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *CreateDeviceLiveToolsPingDeviceRequest) GetCallbackOk() (*CreateDeviceLiveToolsArpTableRequestCallback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *CreateDeviceLiveToolsPingDeviceRequest) SetCallback(v CreateDeviceLiveToolsArpTableRequestCallback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *CreateDeviceLiveToolsPingDeviceRequest) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


