# CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Swap Request ID | 
**Devices** | [**CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices**](CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices.md) |  | 
**Status** | **string** | The current status of the swap job. | 
**AfterAction** | **string** | An action to perform on the devices.old object after swap is complete. | 
**CreatedAt** | **string** | An iso8601 timestamp for the creation of the swap request. | 
**CompletedAt** | Pointer to **string** | An iso8601 timestamp for when the swap completed or failed. | [optional] 
**Errors** | Pointer to **[]string** | A list of error messages for why a swap failed. | [optional] 

## Methods

### NewCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner

`func NewCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner(id string, devices CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices, status string, afterAction string, createdAt string, ) *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner`

NewCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner instantiates a new CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerWithDefaults

`func NewCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerWithDefaults() *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner`

NewCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerWithDefaults instantiates a new CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) SetId(v string)`

SetId sets Id field to given value.


### GetDevices

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) GetDevices() CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) GetDevicesOk() (*CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) SetDevices(v CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices)`

SetDevices sets Devices field to given value.


### GetStatus

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAfterAction

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) GetAfterAction() string`

GetAfterAction returns the AfterAction field if non-nil, zero value otherwise.

### GetAfterActionOk

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) GetAfterActionOk() (*string, bool)`

GetAfterActionOk returns a tuple with the AfterAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterAction

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) SetAfterAction(v string)`

SetAfterAction sets AfterAction field to given value.


### GetCreatedAt

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCompletedAt

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetErrors

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


