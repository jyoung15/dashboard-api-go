# GetNetworkAlertsSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDestinations** | Pointer to [**GetNetworkAlertsSettings200ResponseDefaultDestinations**](GetNetworkAlertsSettings200ResponseDefaultDestinations.md) |  | [optional] 
**Alerts** | Pointer to [**[]GetNetworkAlertsSettings200ResponseAlertsInner**](GetNetworkAlertsSettings200ResponseAlertsInner.md) | Alert-specific configuration for each type. Only alerts that pertain to the network can be updated. | [optional] 
**Muting** | Pointer to [**GetNetworkAlertsSettings200ResponseMuting**](GetNetworkAlertsSettings200ResponseMuting.md) |  | [optional] 

## Methods

### NewGetNetworkAlertsSettings200Response

`func NewGetNetworkAlertsSettings200Response() *GetNetworkAlertsSettings200Response`

NewGetNetworkAlertsSettings200Response instantiates a new GetNetworkAlertsSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkAlertsSettings200ResponseWithDefaults

`func NewGetNetworkAlertsSettings200ResponseWithDefaults() *GetNetworkAlertsSettings200Response`

NewGetNetworkAlertsSettings200ResponseWithDefaults instantiates a new GetNetworkAlertsSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDestinations

`func (o *GetNetworkAlertsSettings200Response) GetDefaultDestinations() GetNetworkAlertsSettings200ResponseDefaultDestinations`

GetDefaultDestinations returns the DefaultDestinations field if non-nil, zero value otherwise.

### GetDefaultDestinationsOk

`func (o *GetNetworkAlertsSettings200Response) GetDefaultDestinationsOk() (*GetNetworkAlertsSettings200ResponseDefaultDestinations, bool)`

GetDefaultDestinationsOk returns a tuple with the DefaultDestinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDestinations

`func (o *GetNetworkAlertsSettings200Response) SetDefaultDestinations(v GetNetworkAlertsSettings200ResponseDefaultDestinations)`

SetDefaultDestinations sets DefaultDestinations field to given value.

### HasDefaultDestinations

`func (o *GetNetworkAlertsSettings200Response) HasDefaultDestinations() bool`

HasDefaultDestinations returns a boolean if a field has been set.

### GetAlerts

`func (o *GetNetworkAlertsSettings200Response) GetAlerts() []GetNetworkAlertsSettings200ResponseAlertsInner`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *GetNetworkAlertsSettings200Response) GetAlertsOk() (*[]GetNetworkAlertsSettings200ResponseAlertsInner, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *GetNetworkAlertsSettings200Response) SetAlerts(v []GetNetworkAlertsSettings200ResponseAlertsInner)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *GetNetworkAlertsSettings200Response) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetMuting

`func (o *GetNetworkAlertsSettings200Response) GetMuting() GetNetworkAlertsSettings200ResponseMuting`

GetMuting returns the Muting field if non-nil, zero value otherwise.

### GetMutingOk

`func (o *GetNetworkAlertsSettings200Response) GetMutingOk() (*GetNetworkAlertsSettings200ResponseMuting, bool)`

GetMutingOk returns a tuple with the Muting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuting

`func (o *GetNetworkAlertsSettings200Response) SetMuting(v GetNetworkAlertsSettings200ResponseMuting)`

SetMuting sets Muting field to given value.

### HasMuting

`func (o *GetNetworkAlertsSettings200Response) HasMuting() bool`

HasMuting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


