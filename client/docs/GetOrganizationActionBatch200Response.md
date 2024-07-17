# GetOrganizationActionBatch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the action batch. Can be used to check the status of the action batch at /organizations/{organizationId}/actionBatches/{actionBatchId} | [optional] 
**OrganizationId** | Pointer to **string** | ID of the organization this action batch belongs to | [optional] 
**Confirmed** | Pointer to **bool** | Flag describing whether the action should be previewed before executing or not | [optional] 
**Synchronous** | Pointer to **bool** | Flag describing whether actions should run synchronously or asynchronously | [optional] 
**Status** | Pointer to [**GetOrganizationActionBatches200ResponseInnerStatus**](GetOrganizationActionBatches200ResponseInnerStatus.md) |  | [optional] 
**Actions** | [**[]GetOrganizationActionBatches200ResponseInnerActionsInner**](GetOrganizationActionBatches200ResponseInnerActionsInner.md) | A set of changes made as part of this action (&lt;a href&#x3D;&#39;https://developer.cisco.com/meraki/api/#/rest/guides/action-batches/&#39;&gt;more details&lt;/a&gt;) | 
**Callback** | Pointer to [**GetDeviceLiveToolsPingDevice200ResponseCallback**](GetDeviceLiveToolsPingDevice200ResponseCallback.md) |  | [optional] 

## Methods

### NewGetOrganizationActionBatch200Response

`func NewGetOrganizationActionBatch200Response(actions []GetOrganizationActionBatches200ResponseInnerActionsInner, ) *GetOrganizationActionBatch200Response`

NewGetOrganizationActionBatch200Response instantiates a new GetOrganizationActionBatch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationActionBatch200ResponseWithDefaults

`func NewGetOrganizationActionBatch200ResponseWithDefaults() *GetOrganizationActionBatch200Response`

NewGetOrganizationActionBatch200ResponseWithDefaults instantiates a new GetOrganizationActionBatch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOrganizationActionBatch200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrganizationActionBatch200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrganizationActionBatch200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetOrganizationActionBatch200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *GetOrganizationActionBatch200Response) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetOrganizationActionBatch200Response) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetOrganizationActionBatch200Response) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *GetOrganizationActionBatch200Response) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetConfirmed

`func (o *GetOrganizationActionBatch200Response) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *GetOrganizationActionBatch200Response) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *GetOrganizationActionBatch200Response) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *GetOrganizationActionBatch200Response) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetSynchronous

`func (o *GetOrganizationActionBatch200Response) GetSynchronous() bool`

GetSynchronous returns the Synchronous field if non-nil, zero value otherwise.

### GetSynchronousOk

`func (o *GetOrganizationActionBatch200Response) GetSynchronousOk() (*bool, bool)`

GetSynchronousOk returns a tuple with the Synchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronous

`func (o *GetOrganizationActionBatch200Response) SetSynchronous(v bool)`

SetSynchronous sets Synchronous field to given value.

### HasSynchronous

`func (o *GetOrganizationActionBatch200Response) HasSynchronous() bool`

HasSynchronous returns a boolean if a field has been set.

### GetStatus

`func (o *GetOrganizationActionBatch200Response) GetStatus() GetOrganizationActionBatches200ResponseInnerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrganizationActionBatch200Response) GetStatusOk() (*GetOrganizationActionBatches200ResponseInnerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrganizationActionBatch200Response) SetStatus(v GetOrganizationActionBatches200ResponseInnerStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetOrganizationActionBatch200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetActions

`func (o *GetOrganizationActionBatch200Response) GetActions() []GetOrganizationActionBatches200ResponseInnerActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *GetOrganizationActionBatch200Response) GetActionsOk() (*[]GetOrganizationActionBatches200ResponseInnerActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *GetOrganizationActionBatch200Response) SetActions(v []GetOrganizationActionBatches200ResponseInnerActionsInner)`

SetActions sets Actions field to given value.


### GetCallback

`func (o *GetOrganizationActionBatch200Response) GetCallback() GetDeviceLiveToolsPingDevice200ResponseCallback`

GetCallback returns the Callback field if non-nil, zero value otherwise.

### GetCallbackOk

`func (o *GetOrganizationActionBatch200Response) GetCallbackOk() (*GetDeviceLiveToolsPingDevice200ResponseCallback, bool)`

GetCallbackOk returns a tuple with the Callback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallback

`func (o *GetOrganizationActionBatch200Response) SetCallback(v GetDeviceLiveToolsPingDevice200ResponseCallback)`

SetCallback sets Callback field to given value.

### HasCallback

`func (o *GetOrganizationActionBatch200Response) HasCallback() bool`

HasCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


