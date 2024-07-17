# GetOrganizationAssuranceAlertsOverviewByType200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner**](GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner.md) | Organization Alert counts by type | 
**Meta** | [**GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMeta**](GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMeta.md) |  | 

## Methods

### NewGetOrganizationAssuranceAlertsOverviewByType200Response

`func NewGetOrganizationAssuranceAlertsOverviewByType200Response(items []GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner, meta GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMeta, ) *GetOrganizationAssuranceAlertsOverviewByType200Response`

NewGetOrganizationAssuranceAlertsOverviewByType200Response instantiates a new GetOrganizationAssuranceAlertsOverviewByType200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationAssuranceAlertsOverviewByType200ResponseWithDefaults

`func NewGetOrganizationAssuranceAlertsOverviewByType200ResponseWithDefaults() *GetOrganizationAssuranceAlertsOverviewByType200Response`

NewGetOrganizationAssuranceAlertsOverviewByType200ResponseWithDefaults instantiates a new GetOrganizationAssuranceAlertsOverviewByType200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *GetOrganizationAssuranceAlertsOverviewByType200Response) GetItems() []GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetOrganizationAssuranceAlertsOverviewByType200Response) GetItemsOk() (*[]GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetOrganizationAssuranceAlertsOverviewByType200Response) SetItems(v []GetOrganizationAssuranceAlertsOverviewByType200ResponseItemsInner)`

SetItems sets Items field to given value.


### GetMeta

`func (o *GetOrganizationAssuranceAlertsOverviewByType200Response) GetMeta() GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetOrganizationAssuranceAlertsOverviewByType200Response) GetMetaOk() (*GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetOrganizationAssuranceAlertsOverviewByType200Response) SetMeta(v GetOrganizationAssuranceAlertsOverviewByNetwork200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


