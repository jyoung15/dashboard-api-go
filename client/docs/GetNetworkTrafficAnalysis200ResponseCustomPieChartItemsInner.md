# GetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the custom pie chart item. | 
**Type** | **string** |     The signature type for the custom pie chart item. Can be one of &#39;host&#39;, &#39;port&#39; or &#39;ipRange&#39;.  | 
**Value** | **string** |     The value of the custom pie chart item. Valid syntax depends on the signature type of the chart item     (see sample request/response for more details).  | 

## Methods

### NewGetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner

`func NewGetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner(name string, type_ string, value string, ) *GetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner`

NewGetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner instantiates a new GetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInnerWithDefaults

`func NewGetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInnerWithDefaults() *GetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner`

NewGetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInnerWithDefaults instantiates a new GetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *GetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *GetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetNetworkTrafficAnalysis200ResponseCustomPieChartItemsInner) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


