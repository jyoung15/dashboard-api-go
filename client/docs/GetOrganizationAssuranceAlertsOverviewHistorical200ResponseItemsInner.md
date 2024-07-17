# GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentStart** | **time.Time** | Starting datetime of the segment in iso8601 format | 
**Totals** | [**GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInnerTotals**](GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInnerTotals.md) |  | 
**ByAlertType** | [**[]GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInnerByAlertTypeInner**](GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInnerByAlertTypeInner.md) | Totals by Type | 

## Methods

### NewGetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInner

`func NewGetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInner(segmentStart time.Time, totals GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInnerTotals, byAlertType []GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInnerByAlertTypeInner, ) *GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInner`

NewGetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInner instantiates a new GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInnerWithDefaults

`func NewGetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInnerWithDefaults() *GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInner`

NewGetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInnerWithDefaults instantiates a new GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentStart

`func (o *GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInner) GetSegmentStart() time.Time`

GetSegmentStart returns the SegmentStart field if non-nil, zero value otherwise.

### GetSegmentStartOk

`func (o *GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInner) GetSegmentStartOk() (*time.Time, bool)`

GetSegmentStartOk returns a tuple with the SegmentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStart

`func (o *GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInner) SetSegmentStart(v time.Time)`

SetSegmentStart sets SegmentStart field to given value.


### GetTotals

`func (o *GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInner) GetTotals() GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInnerTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInner) GetTotalsOk() (*GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInnerTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInner) SetTotals(v GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInnerTotals)`

SetTotals sets Totals field to given value.


### GetByAlertType

`func (o *GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInner) GetByAlertType() []GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInnerByAlertTypeInner`

GetByAlertType returns the ByAlertType field if non-nil, zero value otherwise.

### GetByAlertTypeOk

`func (o *GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInner) GetByAlertTypeOk() (*[]GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInnerByAlertTypeInner, bool)`

GetByAlertTypeOk returns a tuple with the ByAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByAlertType

`func (o *GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInner) SetByAlertType(v []GetOrganizationAssuranceAlertsOverviewHistorical200ResponseItemsInnerByAlertTypeInner)`

SetByAlertType sets ByAlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


