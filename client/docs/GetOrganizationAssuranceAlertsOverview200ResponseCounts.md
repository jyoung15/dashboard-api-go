# GetOrganizationAssuranceAlertsOverview200ResponseCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Total number of alerts on the organization | 
**BySeverity** | [**[]GetOrganizationAssuranceAlertsOverview200ResponseCountsBySeverityInner**](GetOrganizationAssuranceAlertsOverview200ResponseCountsBySeverityInner.md) | Counts of alerts on organization by severity | 

## Methods

### NewGetOrganizationAssuranceAlertsOverview200ResponseCounts

`func NewGetOrganizationAssuranceAlertsOverview200ResponseCounts(total int32, bySeverity []GetOrganizationAssuranceAlertsOverview200ResponseCountsBySeverityInner, ) *GetOrganizationAssuranceAlertsOverview200ResponseCounts`

NewGetOrganizationAssuranceAlertsOverview200ResponseCounts instantiates a new GetOrganizationAssuranceAlertsOverview200ResponseCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationAssuranceAlertsOverview200ResponseCountsWithDefaults

`func NewGetOrganizationAssuranceAlertsOverview200ResponseCountsWithDefaults() *GetOrganizationAssuranceAlertsOverview200ResponseCounts`

NewGetOrganizationAssuranceAlertsOverview200ResponseCountsWithDefaults instantiates a new GetOrganizationAssuranceAlertsOverview200ResponseCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetOrganizationAssuranceAlertsOverview200ResponseCounts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetOrganizationAssuranceAlertsOverview200ResponseCounts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetOrganizationAssuranceAlertsOverview200ResponseCounts) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetBySeverity

`func (o *GetOrganizationAssuranceAlertsOverview200ResponseCounts) GetBySeverity() []GetOrganizationAssuranceAlertsOverview200ResponseCountsBySeverityInner`

GetBySeverity returns the BySeverity field if non-nil, zero value otherwise.

### GetBySeverityOk

`func (o *GetOrganizationAssuranceAlertsOverview200ResponseCounts) GetBySeverityOk() (*[]GetOrganizationAssuranceAlertsOverview200ResponseCountsBySeverityInner, bool)`

GetBySeverityOk returns a tuple with the BySeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySeverity

`func (o *GetOrganizationAssuranceAlertsOverview200ResponseCounts) SetBySeverity(v []GetOrganizationAssuranceAlertsOverview200ResponseCountsBySeverityInner)`

SetBySeverity sets BySeverity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


