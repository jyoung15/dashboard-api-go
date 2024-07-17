# GetNetworkAlertsSettings200ResponseAlertsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of alert | 
**Enabled** | Pointer to **bool** | A boolean depicting if the alert is turned on or off | [optional] 
**AlertDestinations** | Pointer to [**GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations**](GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations.md) |  | [optional] 
**Filters** | Pointer to [**GetNetworkAlertsSettings200ResponseAlertsInnerFilters**](GetNetworkAlertsSettings200ResponseAlertsInnerFilters.md) |  | [optional] 

## Methods

### NewGetNetworkAlertsSettings200ResponseAlertsInner

`func NewGetNetworkAlertsSettings200ResponseAlertsInner(type_ string, ) *GetNetworkAlertsSettings200ResponseAlertsInner`

NewGetNetworkAlertsSettings200ResponseAlertsInner instantiates a new GetNetworkAlertsSettings200ResponseAlertsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkAlertsSettings200ResponseAlertsInnerWithDefaults

`func NewGetNetworkAlertsSettings200ResponseAlertsInnerWithDefaults() *GetNetworkAlertsSettings200ResponseAlertsInner`

NewGetNetworkAlertsSettings200ResponseAlertsInnerWithDefaults instantiates a new GetNetworkAlertsSettings200ResponseAlertsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetNetworkAlertsSettings200ResponseAlertsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetworkAlertsSettings200ResponseAlertsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetworkAlertsSettings200ResponseAlertsInner) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *GetNetworkAlertsSettings200ResponseAlertsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkAlertsSettings200ResponseAlertsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkAlertsSettings200ResponseAlertsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkAlertsSettings200ResponseAlertsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAlertDestinations

`func (o *GetNetworkAlertsSettings200ResponseAlertsInner) GetAlertDestinations() GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations`

GetAlertDestinations returns the AlertDestinations field if non-nil, zero value otherwise.

### GetAlertDestinationsOk

`func (o *GetNetworkAlertsSettings200ResponseAlertsInner) GetAlertDestinationsOk() (*GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations, bool)`

GetAlertDestinationsOk returns a tuple with the AlertDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertDestinations

`func (o *GetNetworkAlertsSettings200ResponseAlertsInner) SetAlertDestinations(v GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations)`

SetAlertDestinations sets AlertDestinations field to given value.

### HasAlertDestinations

`func (o *GetNetworkAlertsSettings200ResponseAlertsInner) HasAlertDestinations() bool`

HasAlertDestinations returns a boolean if a field has been set.

### GetFilters

`func (o *GetNetworkAlertsSettings200ResponseAlertsInner) GetFilters() GetNetworkAlertsSettings200ResponseAlertsInnerFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GetNetworkAlertsSettings200ResponseAlertsInner) GetFiltersOk() (*GetNetworkAlertsSettings200ResponseAlertsInnerFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GetNetworkAlertsSettings200ResponseAlertsInner) SetFilters(v GetNetworkAlertsSettings200ResponseAlertsInnerFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GetNetworkAlertsSettings200ResponseAlertsInner) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


