# GetDeviceSensorCommands200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandId** | Pointer to **string** | ID to check the status of the command request | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the command was triggered | [optional] 
**CompletedAt** | Pointer to **time.Time** | Time when the command was completed | [optional] 
**CreatedBy** | Pointer to [**GetDeviceSensorCommands200ResponseInnerCreatedBy**](GetDeviceSensorCommands200ResponseInnerCreatedBy.md) |  | [optional] 
**Operation** | Pointer to **string** | Operation run on the sensor | [optional] 
**Status** | Pointer to **string** | Status of the command request | [optional] 
**Errors** | Pointer to **[]string** | Array of errors if failed | [optional] 

## Methods

### NewGetDeviceSensorCommands200ResponseInner

`func NewGetDeviceSensorCommands200ResponseInner() *GetDeviceSensorCommands200ResponseInner`

NewGetDeviceSensorCommands200ResponseInner instantiates a new GetDeviceSensorCommands200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceSensorCommands200ResponseInnerWithDefaults

`func NewGetDeviceSensorCommands200ResponseInnerWithDefaults() *GetDeviceSensorCommands200ResponseInner`

NewGetDeviceSensorCommands200ResponseInnerWithDefaults instantiates a new GetDeviceSensorCommands200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandId

`func (o *GetDeviceSensorCommands200ResponseInner) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *GetDeviceSensorCommands200ResponseInner) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *GetDeviceSensorCommands200ResponseInner) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *GetDeviceSensorCommands200ResponseInner) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetDeviceSensorCommands200ResponseInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetDeviceSensorCommands200ResponseInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetDeviceSensorCommands200ResponseInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetDeviceSensorCommands200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *GetDeviceSensorCommands200ResponseInner) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GetDeviceSensorCommands200ResponseInner) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GetDeviceSensorCommands200ResponseInner) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GetDeviceSensorCommands200ResponseInner) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetDeviceSensorCommands200ResponseInner) GetCreatedBy() GetDeviceSensorCommands200ResponseInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetDeviceSensorCommands200ResponseInner) GetCreatedByOk() (*GetDeviceSensorCommands200ResponseInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetDeviceSensorCommands200ResponseInner) SetCreatedBy(v GetDeviceSensorCommands200ResponseInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetDeviceSensorCommands200ResponseInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetOperation

`func (o *GetDeviceSensorCommands200ResponseInner) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *GetDeviceSensorCommands200ResponseInner) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *GetDeviceSensorCommands200ResponseInner) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *GetDeviceSensorCommands200ResponseInner) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetStatus

`func (o *GetDeviceSensorCommands200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDeviceSensorCommands200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDeviceSensorCommands200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetDeviceSensorCommands200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrors

`func (o *GetDeviceSensorCommands200ResponseInner) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GetDeviceSensorCommands200ResponseInner) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GetDeviceSensorCommands200ResponseInner) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GetDeviceSensorCommands200ResponseInner) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


