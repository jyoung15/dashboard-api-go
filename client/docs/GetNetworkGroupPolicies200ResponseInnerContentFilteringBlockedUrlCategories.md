# GetNetworkGroupPolicies200ResponseInnerContentFilteringBlockedUrlCategories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to **string** | How URL categories are applied. Can be &#39;network default&#39;, &#39;append&#39; or &#39;override&#39;. | [optional] 
**Categories** | Pointer to **[]string** | A list of URL categories to block | [optional] 

## Methods

### NewGetNetworkGroupPolicies200ResponseInnerContentFilteringBlockedUrlCategories

`func NewGetNetworkGroupPolicies200ResponseInnerContentFilteringBlockedUrlCategories() *GetNetworkGroupPolicies200ResponseInnerContentFilteringBlockedUrlCategories`

NewGetNetworkGroupPolicies200ResponseInnerContentFilteringBlockedUrlCategories instantiates a new GetNetworkGroupPolicies200ResponseInnerContentFilteringBlockedUrlCategories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkGroupPolicies200ResponseInnerContentFilteringBlockedUrlCategoriesWithDefaults

`func NewGetNetworkGroupPolicies200ResponseInnerContentFilteringBlockedUrlCategoriesWithDefaults() *GetNetworkGroupPolicies200ResponseInnerContentFilteringBlockedUrlCategories`

NewGetNetworkGroupPolicies200ResponseInnerContentFilteringBlockedUrlCategoriesWithDefaults instantiates a new GetNetworkGroupPolicies200ResponseInnerContentFilteringBlockedUrlCategories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *GetNetworkGroupPolicies200ResponseInnerContentFilteringBlockedUrlCategories) GetSettings() string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *GetNetworkGroupPolicies200ResponseInnerContentFilteringBlockedUrlCategories) GetSettingsOk() (*string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *GetNetworkGroupPolicies200ResponseInnerContentFilteringBlockedUrlCategories) SetSettings(v string)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *GetNetworkGroupPolicies200ResponseInnerContentFilteringBlockedUrlCategories) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetCategories

`func (o *GetNetworkGroupPolicies200ResponseInnerContentFilteringBlockedUrlCategories) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *GetNetworkGroupPolicies200ResponseInnerContentFilteringBlockedUrlCategories) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *GetNetworkGroupPolicies200ResponseInnerContentFilteringBlockedUrlCategories) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *GetNetworkGroupPolicies200ResponseInnerContentFilteringBlockedUrlCategories) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


