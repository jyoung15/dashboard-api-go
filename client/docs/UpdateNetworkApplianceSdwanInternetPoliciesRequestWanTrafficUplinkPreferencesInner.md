# UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreferredUplink** | **string** | Preferred uplink for uplink preference rule. Must be one of: &#39;wan1&#39;, &#39;wan2&#39;, &#39;bestForVoIP&#39;, &#39;loadBalancing&#39; or &#39;defaultUplink&#39; | 
**FailOverCriterion** | Pointer to **string** | WAN failover and failback behavior | [optional] 
**PerformanceClass** | Pointer to [**UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass**](UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass.md) |  | [optional] 
**TrafficFilters** | [**[]UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInner**](UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInner.md) | Traffic filters | 

## Methods

### NewUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner

`func NewUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner(preferredUplink string, trafficFilters []UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInner, ) *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner`

NewUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner instantiates a new UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerWithDefaults

`func NewUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerWithDefaults() *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner`

NewUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerWithDefaults instantiates a new UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreferredUplink

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner) GetPreferredUplink() string`

GetPreferredUplink returns the PreferredUplink field if non-nil, zero value otherwise.

### GetPreferredUplinkOk

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner) GetPreferredUplinkOk() (*string, bool)`

GetPreferredUplinkOk returns a tuple with the PreferredUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUplink

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner) SetPreferredUplink(v string)`

SetPreferredUplink sets PreferredUplink field to given value.


### GetFailOverCriterion

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner) GetFailOverCriterion() string`

GetFailOverCriterion returns the FailOverCriterion field if non-nil, zero value otherwise.

### GetFailOverCriterionOk

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner) GetFailOverCriterionOk() (*string, bool)`

GetFailOverCriterionOk returns a tuple with the FailOverCriterion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailOverCriterion

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner) SetFailOverCriterion(v string)`

SetFailOverCriterion sets FailOverCriterion field to given value.

### HasFailOverCriterion

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner) HasFailOverCriterion() bool`

HasFailOverCriterion returns a boolean if a field has been set.

### GetPerformanceClass

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner) GetPerformanceClass() UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass`

GetPerformanceClass returns the PerformanceClass field if non-nil, zero value otherwise.

### GetPerformanceClassOk

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner) GetPerformanceClassOk() (*UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass, bool)`

GetPerformanceClassOk returns a tuple with the PerformanceClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceClass

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner) SetPerformanceClass(v UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerPerformanceClass)`

SetPerformanceClass sets PerformanceClass field to given value.

### HasPerformanceClass

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner) HasPerformanceClass() bool`

HasPerformanceClass returns a boolean if a field has been set.

### GetTrafficFilters

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner) GetTrafficFilters() []UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInner`

GetTrafficFilters returns the TrafficFilters field if non-nil, zero value otherwise.

### GetTrafficFiltersOk

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner) GetTrafficFiltersOk() (*[]UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInner, bool)`

GetTrafficFiltersOk returns a tuple with the TrafficFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficFilters

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInner) SetTrafficFilters(v []UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInner)`

SetTrafficFilters sets TrafficFilters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


