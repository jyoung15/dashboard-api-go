# UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **string** | E.g.: \&quot;any\&quot;, \&quot;0\&quot; (also means \&quot;any\&quot;), \&quot;8080\&quot;, \&quot;1-1024\&quot; | [optional] 
**Cidr** | Pointer to **string** | CIDR format address (e.g.\&quot;192.168.10.1\&quot;, which is the same as \&quot;192.168.10.1/32\&quot;), or \&quot;any\&quot; | [optional] 
**Applications** | Pointer to [**[]UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestinationApplicationsInner**](UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestinationApplicationsInner.md) | list of application objects (either majorApplication or nbar) | [optional] 

## Methods

### NewUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination

`func NewUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination() *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination`

NewUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination instantiates a new UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestinationWithDefaults

`func NewUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestinationWithDefaults() *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination`

NewUpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestinationWithDefaults instantiates a new UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetCidr

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetApplications

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination) GetApplications() []UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestinationApplicationsInner`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination) GetApplicationsOk() (*[]UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestinationApplicationsInner, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination) SetApplications(v []UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestinationApplicationsInner)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *UpdateNetworkApplianceSdwanInternetPoliciesRequestWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


