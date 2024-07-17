# GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | **string** | id | 
**NetworkName** | **string** | Name | 
**AlertCount** | **int32** | Total Alerts | 
**SeverityCounts** | [**[]GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInnerSeverityCountsInner**](GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInnerSeverityCountsInner.md) | Alerts By Severity | 

## Methods

### NewGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner

`func NewGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner(networkId string, networkName string, alertCount int32, severityCounts []GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInnerSeverityCountsInner, ) *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner`

NewGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner instantiates a new GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInnerWithDefaults

`func NewGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInnerWithDefaults() *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner`

NewGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInnerWithDefaults instantiates a new GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetNetworkName

`func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.


### GetAlertCount

`func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner) GetAlertCount() int32`

GetAlertCount returns the AlertCount field if non-nil, zero value otherwise.

### GetAlertCountOk

`func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner) GetAlertCountOk() (*int32, bool)`

GetAlertCountOk returns a tuple with the AlertCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCount

`func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner) SetAlertCount(v int32)`

SetAlertCount sets AlertCount field to given value.


### GetSeverityCounts

`func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner) GetSeverityCounts() []GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInnerSeverityCountsInner`

GetSeverityCounts returns the SeverityCounts field if non-nil, zero value otherwise.

### GetSeverityCountsOk

`func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner) GetSeverityCountsOk() (*[]GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInnerSeverityCountsInner, bool)`

GetSeverityCountsOk returns a tuple with the SeverityCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityCounts

`func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner) SetSeverityCounts(v []GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInnerSeverityCountsInner)`

SetSeverityCounts sets SeverityCounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


