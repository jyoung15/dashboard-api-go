# GetNetworkTrafficAnalysis200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** |     The traffic analysis mode for the network. Can be one of &#39;disabled&#39; (do not collect traffic types),     &#39;basic&#39; (collect generic traffic categories), or &#39;detailed&#39; (collect destination hostnames).  | [optional] 
**CustomPieChartItems** | Pointer to [**[]GetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner**](GetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner.md) | The list of items that make up the custom pie chart for traffic reporting. | [optional] 

## Methods

### NewGetNetworkTrafficAnalysis200Response

`func NewGetNetworkTrafficAnalysis200Response() *GetNetworkTrafficAnalysis200Response`

NewGetNetworkTrafficAnalysis200Response instantiates a new GetNetworkTrafficAnalysis200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkTrafficAnalysis200ResponseWithDefaults

`func NewGetNetworkTrafficAnalysis200ResponseWithDefaults() *GetNetworkTrafficAnalysis200Response`

NewGetNetworkTrafficAnalysis200ResponseWithDefaults instantiates a new GetNetworkTrafficAnalysis200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *GetNetworkTrafficAnalysis200Response) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetNetworkTrafficAnalysis200Response) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetNetworkTrafficAnalysis200Response) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GetNetworkTrafficAnalysis200Response) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetCustomPieChartItems

`func (o *GetNetworkTrafficAnalysis200Response) GetCustomPieChartItems() []GetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner`

GetCustomPieChartItems returns the CustomPieChartItems field if non-nil, zero value otherwise.

### GetCustomPieChartItemsOk

`func (o *GetNetworkTrafficAnalysis200Response) GetCustomPieChartItemsOk() (*[]GetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner, bool)`

GetCustomPieChartItemsOk returns a tuple with the CustomPieChartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPieChartItems

`func (o *GetNetworkTrafficAnalysis200Response) SetCustomPieChartItems(v []GetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner)`

SetCustomPieChartItems sets CustomPieChartItems field to given value.

### HasCustomPieChartItems

`func (o *GetNetworkTrafficAnalysis200Response) HasCustomPieChartItems() bool`

HasCustomPieChartItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


