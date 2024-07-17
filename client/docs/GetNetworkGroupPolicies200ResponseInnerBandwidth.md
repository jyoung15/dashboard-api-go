# GetNetworkGroupPolicies200ResponseInnerBandwidth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to **string** | How bandwidth limits are enforced. Can be &#39;network default&#39;, &#39;ignore&#39; or &#39;custom&#39;. | [optional] 
**BandwidthLimits** | Pointer to [**GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits**](GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits.md) |  | [optional] 

## Methods

### NewGetNetworkGroupPolicies200ResponseInnerBandwidth

`func NewGetNetworkGroupPolicies200ResponseInnerBandwidth() *GetNetworkGroupPolicies200ResponseInnerBandwidth`

NewGetNetworkGroupPolicies200ResponseInnerBandwidth instantiates a new GetNetworkGroupPolicies200ResponseInnerBandwidth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkGroupPolicies200ResponseInnerBandwidthWithDefaults

`func NewGetNetworkGroupPolicies200ResponseInnerBandwidthWithDefaults() *GetNetworkGroupPolicies200ResponseInnerBandwidth`

NewGetNetworkGroupPolicies200ResponseInnerBandwidthWithDefaults instantiates a new GetNetworkGroupPolicies200ResponseInnerBandwidth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *GetNetworkGroupPolicies200ResponseInnerBandwidth) GetSettings() string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *GetNetworkGroupPolicies200ResponseInnerBandwidth) GetSettingsOk() (*string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *GetNetworkGroupPolicies200ResponseInnerBandwidth) SetSettings(v string)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *GetNetworkGroupPolicies200ResponseInnerBandwidth) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetBandwidthLimits

`func (o *GetNetworkGroupPolicies200ResponseInnerBandwidth) GetBandwidthLimits() GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits`

GetBandwidthLimits returns the BandwidthLimits field if non-nil, zero value otherwise.

### GetBandwidthLimitsOk

`func (o *GetNetworkGroupPolicies200ResponseInnerBandwidth) GetBandwidthLimitsOk() (*GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits, bool)`

GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimits

`func (o *GetNetworkGroupPolicies200ResponseInnerBandwidth) SetBandwidthLimits(v GetNetworkGroupPolicies200ResponseInnerBandwidthBandwidthLimits)`

SetBandwidthLimits sets BandwidthLimits field to given value.

### HasBandwidthLimits

`func (o *GetNetworkGroupPolicies200ResponseInnerBandwidth) HasBandwidthLimits() bool`

HasBandwidthLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


