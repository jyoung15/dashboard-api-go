# GetNetworkGroupPolicies200ResponseInnerContentFilteringAllowedUrlPatterns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to **string** | How URL patterns are applied. Can be &#39;network default&#39;, &#39;append&#39; or &#39;override&#39;. | [optional] 
**Patterns** | Pointer to **[]string** | A list of URL patterns that are allowed | [optional] 

## Methods

### NewGetNetworkGroupPolicies200ResponseInnerContentFilteringAllowedUrlPatterns

`func NewGetNetworkGroupPolicies200ResponseInnerContentFilteringAllowedUrlPatterns() *GetNetworkGroupPolicies200ResponseInnerContentFilteringAllowedUrlPatterns`

NewGetNetworkGroupPolicies200ResponseInnerContentFilteringAllowedUrlPatterns instantiates a new GetNetworkGroupPolicies200ResponseInnerContentFilteringAllowedUrlPatterns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkGroupPolicies200ResponseInnerContentFilteringAllowedUrlPatternsWithDefaults

`func NewGetNetworkGroupPolicies200ResponseInnerContentFilteringAllowedUrlPatternsWithDefaults() *GetNetworkGroupPolicies200ResponseInnerContentFilteringAllowedUrlPatterns`

NewGetNetworkGroupPolicies200ResponseInnerContentFilteringAllowedUrlPatternsWithDefaults instantiates a new GetNetworkGroupPolicies200ResponseInnerContentFilteringAllowedUrlPatterns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *GetNetworkGroupPolicies200ResponseInnerContentFilteringAllowedUrlPatterns) GetSettings() string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *GetNetworkGroupPolicies200ResponseInnerContentFilteringAllowedUrlPatterns) GetSettingsOk() (*string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *GetNetworkGroupPolicies200ResponseInnerContentFilteringAllowedUrlPatterns) SetSettings(v string)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *GetNetworkGroupPolicies200ResponseInnerContentFilteringAllowedUrlPatterns) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetPatterns

`func (o *GetNetworkGroupPolicies200ResponseInnerContentFilteringAllowedUrlPatterns) GetPatterns() []string`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *GetNetworkGroupPolicies200ResponseInnerContentFilteringAllowedUrlPatterns) GetPatternsOk() (*[]string, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *GetNetworkGroupPolicies200ResponseInnerContentFilteringAllowedUrlPatterns) SetPatterns(v []string)`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *GetNetworkGroupPolicies200ResponseInnerContentFilteringAllowedUrlPatterns) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


