# CreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | [**CreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInnerDevices**](CreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInnerDevices.md) |  | 
**AfterAction** | **string** | What action to perform on devices.old after the device cloning is complete. &#39;remove from network&#39; will return the device to inventory, while &#39;release from organization inventory&#39; will free up the license attached to the device. | 

## Methods

### NewCreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInner

`func NewCreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInner(devices CreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInnerDevices, afterAction string, ) *CreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInner`

NewCreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInner instantiates a new CreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInnerWithDefaults

`func NewCreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInnerWithDefaults() *CreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInner`

NewCreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInnerWithDefaults instantiates a new CreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *CreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInner) GetDevices() CreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInnerDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *CreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInner) GetDevicesOk() (*CreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInnerDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *CreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInner) SetDevices(v CreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInnerDevices)`

SetDevices sets Devices field to given value.


### GetAfterAction

`func (o *CreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInner) GetAfterAction() string`

GetAfterAction returns the AfterAction field if non-nil, zero value otherwise.

### GetAfterActionOk

`func (o *CreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInner) GetAfterActionOk() (*string, bool)`

GetAfterActionOk returns a tuple with the AfterAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterAction

`func (o *CreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInner) SetAfterAction(v string)`

SetAfterAction sets AfterAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


