# GetOrganizationAdaptivePolicyPolicies200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdaptivePolicyId** | Pointer to **string** | The ID for the adaptive policy | [optional] 
**SourceGroup** | Pointer to [**GetOrganizationAdaptivePolicyPolicies200ResponseInnerSourceGroup**](GetOrganizationAdaptivePolicyPolicies200ResponseInnerSourceGroup.md) |  | [optional] 
**DestinationGroup** | Pointer to [**GetOrganizationAdaptivePolicyPolicies200ResponseInnerDestinationGroup**](GetOrganizationAdaptivePolicyPolicies200ResponseInnerDestinationGroup.md) |  | [optional] 
**Acls** | Pointer to [**[]GetOrganizationAdaptivePolicyPolicies200ResponseInnerAclsInner**](GetOrganizationAdaptivePolicyPolicies200ResponseInnerAclsInner.md) | The access control lists for the adaptive policy | [optional] 
**LastEntryRule** | Pointer to **string** | The rule to apply if there is no matching ACL | [optional] 
**CreatedAt** | Pointer to **string** | The created at timestamp for the adaptive policy | [optional] 
**UpdatedAt** | Pointer to **string** | The updated at timestamp for the adaptive policy | [optional] 

## Methods

### NewGetOrganizationAdaptivePolicyPolicies200ResponseInner

`func NewGetOrganizationAdaptivePolicyPolicies200ResponseInner() *GetOrganizationAdaptivePolicyPolicies200ResponseInner`

NewGetOrganizationAdaptivePolicyPolicies200ResponseInner instantiates a new GetOrganizationAdaptivePolicyPolicies200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationAdaptivePolicyPolicies200ResponseInnerWithDefaults

`func NewGetOrganizationAdaptivePolicyPolicies200ResponseInnerWithDefaults() *GetOrganizationAdaptivePolicyPolicies200ResponseInner`

NewGetOrganizationAdaptivePolicyPolicies200ResponseInnerWithDefaults instantiates a new GetOrganizationAdaptivePolicyPolicies200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdaptivePolicyId

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetAdaptivePolicyId() string`

GetAdaptivePolicyId returns the AdaptivePolicyId field if non-nil, zero value otherwise.

### GetAdaptivePolicyIdOk

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetAdaptivePolicyIdOk() (*string, bool)`

GetAdaptivePolicyIdOk returns a tuple with the AdaptivePolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptivePolicyId

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) SetAdaptivePolicyId(v string)`

SetAdaptivePolicyId sets AdaptivePolicyId field to given value.

### HasAdaptivePolicyId

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) HasAdaptivePolicyId() bool`

HasAdaptivePolicyId returns a boolean if a field has been set.

### GetSourceGroup

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetSourceGroup() GetOrganizationAdaptivePolicyPolicies200ResponseInnerSourceGroup`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetSourceGroupOk() (*GetOrganizationAdaptivePolicyPolicies200ResponseInnerSourceGroup, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) SetSourceGroup(v GetOrganizationAdaptivePolicyPolicies200ResponseInnerSourceGroup)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### GetDestinationGroup

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetDestinationGroup() GetOrganizationAdaptivePolicyPolicies200ResponseInnerDestinationGroup`

GetDestinationGroup returns the DestinationGroup field if non-nil, zero value otherwise.

### GetDestinationGroupOk

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetDestinationGroupOk() (*GetOrganizationAdaptivePolicyPolicies200ResponseInnerDestinationGroup, bool)`

GetDestinationGroupOk returns a tuple with the DestinationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroup

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) SetDestinationGroup(v GetOrganizationAdaptivePolicyPolicies200ResponseInnerDestinationGroup)`

SetDestinationGroup sets DestinationGroup field to given value.

### HasDestinationGroup

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) HasDestinationGroup() bool`

HasDestinationGroup returns a boolean if a field has been set.

### GetAcls

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetAcls() []GetOrganizationAdaptivePolicyPolicies200ResponseInnerAclsInner`

GetAcls returns the Acls field if non-nil, zero value otherwise.

### GetAclsOk

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetAclsOk() (*[]GetOrganizationAdaptivePolicyPolicies200ResponseInnerAclsInner, bool)`

GetAclsOk returns a tuple with the Acls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcls

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) SetAcls(v []GetOrganizationAdaptivePolicyPolicies200ResponseInnerAclsInner)`

SetAcls sets Acls field to given value.

### HasAcls

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) HasAcls() bool`

HasAcls returns a boolean if a field has been set.

### GetLastEntryRule

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetLastEntryRule() string`

GetLastEntryRule returns the LastEntryRule field if non-nil, zero value otherwise.

### GetLastEntryRuleOk

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetLastEntryRuleOk() (*string, bool)`

GetLastEntryRuleOk returns a tuple with the LastEntryRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEntryRule

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) SetLastEntryRule(v string)`

SetLastEntryRule sets LastEntryRule field to given value.

### HasLastEntryRule

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) HasLastEntryRule() bool`

HasLastEntryRule returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetOrganizationAdaptivePolicyPolicies200ResponseInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


