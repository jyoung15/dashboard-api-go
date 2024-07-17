# CreateOrganizationInventoryDevicesSwapsBulk207Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **string** | The ID of the job that was used to create all of the device swaps. | [optional] 
**Swaps** | Pointer to [**[]CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner**](CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner.md) | An array of recent swap requests and their statuses. | [optional] 

## Methods

### NewCreateOrganizationInventoryDevicesSwapsBulk207Response

`func NewCreateOrganizationInventoryDevicesSwapsBulk207Response() *CreateOrganizationInventoryDevicesSwapsBulk207Response`

NewCreateOrganizationInventoryDevicesSwapsBulk207Response instantiates a new CreateOrganizationInventoryDevicesSwapsBulk207Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationInventoryDevicesSwapsBulk207ResponseWithDefaults

`func NewCreateOrganizationInventoryDevicesSwapsBulk207ResponseWithDefaults() *CreateOrganizationInventoryDevicesSwapsBulk207Response`

NewCreateOrganizationInventoryDevicesSwapsBulk207ResponseWithDefaults instantiates a new CreateOrganizationInventoryDevicesSwapsBulk207Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207Response) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207Response) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207Response) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207Response) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetSwaps

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207Response) GetSwaps() []CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner`

GetSwaps returns the Swaps field if non-nil, zero value otherwise.

### GetSwapsOk

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207Response) GetSwapsOk() (*[]CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner, bool)`

GetSwapsOk returns a tuple with the Swaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwaps

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207Response) SetSwaps(v []CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInner)`

SetSwaps sets Swaps field to given value.

### HasSwaps

`func (o *CreateOrganizationInventoryDevicesSwapsBulk207Response) HasSwaps() bool`

HasSwaps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


