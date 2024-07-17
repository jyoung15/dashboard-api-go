# GetNetworkPiiRequests200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The network or organization identifier | [optional] 
**OrganizationWide** | Pointer to **bool** | If the data returned is organization-wide. False indicates the data is network-wide. | [optional] 
**NetworkId** | Pointer to **string** | The network identifier | [optional] 
**Type** | Pointer to **string** | The type of PII request | [optional] 
**Mac** | Pointer to **string** | The MAC address of the PII request | [optional] 
**Datasets** | Pointer to **string** | The stringified array of datasets related to the provided key that should be deleted. | [optional] 
**Status** | Pointer to **string** | The status of the PII request | [optional] 
**CreatedAt** | Pointer to **int32** | The request&#39;s creation time | [optional] 
**CompletedAt** | Pointer to **int32** | The request&#39;s completion time | [optional] 

## Methods

### NewGetNetworkPiiRequests200ResponseInner

`func NewGetNetworkPiiRequests200ResponseInner() *GetNetworkPiiRequests200ResponseInner`

NewGetNetworkPiiRequests200ResponseInner instantiates a new GetNetworkPiiRequests200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkPiiRequests200ResponseInnerWithDefaults

`func NewGetNetworkPiiRequests200ResponseInnerWithDefaults() *GetNetworkPiiRequests200ResponseInner`

NewGetNetworkPiiRequests200ResponseInnerWithDefaults instantiates a new GetNetworkPiiRequests200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkPiiRequests200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkPiiRequests200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkPiiRequests200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkPiiRequests200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationWide

`func (o *GetNetworkPiiRequests200ResponseInner) GetOrganizationWide() bool`

GetOrganizationWide returns the OrganizationWide field if non-nil, zero value otherwise.

### GetOrganizationWideOk

`func (o *GetNetworkPiiRequests200ResponseInner) GetOrganizationWideOk() (*bool, bool)`

GetOrganizationWideOk returns a tuple with the OrganizationWide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationWide

`func (o *GetNetworkPiiRequests200ResponseInner) SetOrganizationWide(v bool)`

SetOrganizationWide sets OrganizationWide field to given value.

### HasOrganizationWide

`func (o *GetNetworkPiiRequests200ResponseInner) HasOrganizationWide() bool`

HasOrganizationWide returns a boolean if a field has been set.

### GetNetworkId

`func (o *GetNetworkPiiRequests200ResponseInner) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *GetNetworkPiiRequests200ResponseInner) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *GetNetworkPiiRequests200ResponseInner) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *GetNetworkPiiRequests200ResponseInner) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetType

`func (o *GetNetworkPiiRequests200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetworkPiiRequests200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetworkPiiRequests200ResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetNetworkPiiRequests200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMac

`func (o *GetNetworkPiiRequests200ResponseInner) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetNetworkPiiRequests200ResponseInner) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetNetworkPiiRequests200ResponseInner) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetNetworkPiiRequests200ResponseInner) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetDatasets

`func (o *GetNetworkPiiRequests200ResponseInner) GetDatasets() string`

GetDatasets returns the Datasets field if non-nil, zero value otherwise.

### GetDatasetsOk

`func (o *GetNetworkPiiRequests200ResponseInner) GetDatasetsOk() (*string, bool)`

GetDatasetsOk returns a tuple with the Datasets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasets

`func (o *GetNetworkPiiRequests200ResponseInner) SetDatasets(v string)`

SetDatasets sets Datasets field to given value.

### HasDatasets

`func (o *GetNetworkPiiRequests200ResponseInner) HasDatasets() bool`

HasDatasets returns a boolean if a field has been set.

### GetStatus

`func (o *GetNetworkPiiRequests200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetNetworkPiiRequests200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetNetworkPiiRequests200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetNetworkPiiRequests200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetNetworkPiiRequests200ResponseInner) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetNetworkPiiRequests200ResponseInner) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetNetworkPiiRequests200ResponseInner) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetNetworkPiiRequests200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *GetNetworkPiiRequests200ResponseInner) GetCompletedAt() int32`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GetNetworkPiiRequests200ResponseInner) GetCompletedAtOk() (*int32, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GetNetworkPiiRequests200ResponseInner) SetCompletedAt(v int32)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GetNetworkPiiRequests200ResponseInner) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


