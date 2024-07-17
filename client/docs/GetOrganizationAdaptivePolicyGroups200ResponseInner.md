# GetOrganizationAdaptivePolicyGroups200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | The ID of the adaptive policy group | [optional] 
**Name** | Pointer to **string** | The name of the adaptive policy group | [optional] 
**Sgt** | Pointer to **int32** | The security group tag for the adaptive policy group | [optional] 
**Description** | Pointer to **string** | The description for the adaptive policy group | [optional] 
**PolicyObjects** | Pointer to [**[]GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner**](GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner.md) | The policy objects for the adaptive policy group | [optional] 
**IsDefaultGroup** | Pointer to **bool** | Whether the adaptive policy group is the default group | [optional] 
**RequiredIpMappings** | Pointer to **[]string** | List of required IP mappings for the adaptive policy group | [optional] 
**CreatedAt** | Pointer to **string** | Created at timestamp for the adaptive policy group | [optional] 
**UpdatedAt** | Pointer to **string** | Updated at timestamp for the adaptive policy group | [optional] 

## Methods

### NewGetOrganizationAdaptivePolicyGroups200ResponseInner

`func NewGetOrganizationAdaptivePolicyGroups200ResponseInner() *GetOrganizationAdaptivePolicyGroups200ResponseInner`

NewGetOrganizationAdaptivePolicyGroups200ResponseInner instantiates a new GetOrganizationAdaptivePolicyGroups200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationAdaptivePolicyGroups200ResponseInnerWithDefaults

`func NewGetOrganizationAdaptivePolicyGroups200ResponseInnerWithDefaults() *GetOrganizationAdaptivePolicyGroups200ResponseInner`

NewGetOrganizationAdaptivePolicyGroups200ResponseInnerWithDefaults instantiates a new GetOrganizationAdaptivePolicyGroups200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetName

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSgt

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetSgt() int32`

GetSgt returns the Sgt field if non-nil, zero value otherwise.

### GetSgtOk

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetSgtOk() (*int32, bool)`

GetSgtOk returns a tuple with the Sgt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgt

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) SetSgt(v int32)`

SetSgt sets Sgt field to given value.

### HasSgt

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) HasSgt() bool`

HasSgt returns a boolean if a field has been set.

### GetDescription

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyObjects

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetPolicyObjects() []GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner`

GetPolicyObjects returns the PolicyObjects field if non-nil, zero value otherwise.

### GetPolicyObjectsOk

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetPolicyObjectsOk() (*[]GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner, bool)`

GetPolicyObjectsOk returns a tuple with the PolicyObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyObjects

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) SetPolicyObjects(v []GetOrganizationAdaptivePolicyGroups200ResponseInnerPolicyObjectsInner)`

SetPolicyObjects sets PolicyObjects field to given value.

### HasPolicyObjects

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) HasPolicyObjects() bool`

HasPolicyObjects returns a boolean if a field has been set.

### GetIsDefaultGroup

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetIsDefaultGroup() bool`

GetIsDefaultGroup returns the IsDefaultGroup field if non-nil, zero value otherwise.

### GetIsDefaultGroupOk

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetIsDefaultGroupOk() (*bool, bool)`

GetIsDefaultGroupOk returns a tuple with the IsDefaultGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultGroup

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) SetIsDefaultGroup(v bool)`

SetIsDefaultGroup sets IsDefaultGroup field to given value.

### HasIsDefaultGroup

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) HasIsDefaultGroup() bool`

HasIsDefaultGroup returns a boolean if a field has been set.

### GetRequiredIpMappings

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetRequiredIpMappings() []string`

GetRequiredIpMappings returns the RequiredIpMappings field if non-nil, zero value otherwise.

### GetRequiredIpMappingsOk

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetRequiredIpMappingsOk() (*[]string, bool)`

GetRequiredIpMappingsOk returns a tuple with the RequiredIpMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredIpMappings

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) SetRequiredIpMappings(v []string)`

SetRequiredIpMappings sets RequiredIpMappings field to given value.

### HasRequiredIpMappings

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) HasRequiredIpMappings() bool`

HasRequiredIpMappings returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetOrganizationAdaptivePolicyGroups200ResponseInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


