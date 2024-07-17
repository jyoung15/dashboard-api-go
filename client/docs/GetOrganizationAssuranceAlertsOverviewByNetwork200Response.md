# GetOrganizationAssuranceAlertsOverviewByNetwork200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner**](GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner.md) | Alert Counts by Network | 
**Meta** | [**GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMeta**](GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMeta.md) |  | 

## Methods

### NewGetOrganizationAssuranceAlertsOverviewByNetwork200Response

`func NewGetOrganizationAssuranceAlertsOverviewByNetwork200Response(items []GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner, meta GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMeta, ) *GetOrganizationAssuranceAlertsOverviewByNetwork200Response`

NewGetOrganizationAssuranceAlertsOverviewByNetwork200Response instantiates a new GetOrganizationAssuranceAlertsOverviewByNetwork200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseWithDefaults

`func NewGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseWithDefaults() *GetOrganizationAssuranceAlertsOverviewByNetwork200Response`

NewGetOrganizationAssuranceAlertsOverviewByNetwork200ResponseWithDefaults instantiates a new GetOrganizationAssuranceAlertsOverviewByNetwork200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200Response) GetItems() []GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200Response) GetItemsOk() (*[]GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200Response) SetItems(v []GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseItemsInner)`

SetItems sets Items field to given value.


### GetMeta

`func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200Response) GetMeta() GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200Response) GetMetaOk() (*GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetOrganizationAssuranceAlertsOverviewByNetwork200Response) SetMeta(v GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


