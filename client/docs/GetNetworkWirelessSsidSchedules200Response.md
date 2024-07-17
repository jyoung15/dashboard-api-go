# GetNetworkWirelessSsidSchedules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, the SSID outage schedule is enabled. | [optional] 
**Ranges** | Pointer to [**[]GetNetworkWirelessSsidSchedules200ResponseRangesInner**](GetNetworkWirelessSsidSchedules200ResponseRangesInner.md) | List of outage ranges. Has a start date and time, and end date and time. If this parameter is passed in along with rangesInSeconds parameter, this will take precedence. | [optional] 
**RangesInSeconds** | Pointer to [**[]GetNetworkWirelessSsidSchedules200ResponseRangesInSecondsInner**](GetNetworkWirelessSsidSchedules200ResponseRangesInSecondsInner.md) | List of outage ranges in seconds since Sunday at Midnight. Has a start and end. If this parameter is passed in along with the ranges parameter, ranges will take precedence. | [optional] 

## Methods

### NewGetNetworkWirelessSsidSchedules200Response

`func NewGetNetworkWirelessSsidSchedules200Response() *GetNetworkWirelessSsidSchedules200Response`

NewGetNetworkWirelessSsidSchedules200Response instantiates a new GetNetworkWirelessSsidSchedules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWirelessSsidSchedules200ResponseWithDefaults

`func NewGetNetworkWirelessSsidSchedules200ResponseWithDefaults() *GetNetworkWirelessSsidSchedules200Response`

NewGetNetworkWirelessSsidSchedules200ResponseWithDefaults instantiates a new GetNetworkWirelessSsidSchedules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetNetworkWirelessSsidSchedules200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkWirelessSsidSchedules200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkWirelessSsidSchedules200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkWirelessSsidSchedules200Response) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRanges

`func (o *GetNetworkWirelessSsidSchedules200Response) GetRanges() []GetNetworkWirelessSsidSchedules200ResponseRangesInner`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *GetNetworkWirelessSsidSchedules200Response) GetRangesOk() (*[]GetNetworkWirelessSsidSchedules200ResponseRangesInner, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *GetNetworkWirelessSsidSchedules200Response) SetRanges(v []GetNetworkWirelessSsidSchedules200ResponseRangesInner)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *GetNetworkWirelessSsidSchedules200Response) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetRangesInSeconds

`func (o *GetNetworkWirelessSsidSchedules200Response) GetRangesInSeconds() []GetNetworkWirelessSsidSchedules200ResponseRangesInSecondsInner`

GetRangesInSeconds returns the RangesInSeconds field if non-nil, zero value otherwise.

### GetRangesInSecondsOk

`func (o *GetNetworkWirelessSsidSchedules200Response) GetRangesInSecondsOk() (*[]GetNetworkWirelessSsidSchedules200ResponseRangesInSecondsInner, bool)`

GetRangesInSecondsOk returns a tuple with the RangesInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangesInSeconds

`func (o *GetNetworkWirelessSsidSchedules200Response) SetRangesInSeconds(v []GetNetworkWirelessSsidSchedules200ResponseRangesInSecondsInner)`

SetRangesInSeconds sets RangesInSeconds field to given value.

### HasRangesInSeconds

`func (o *GetNetworkWirelessSsidSchedules200Response) HasRangesInSeconds() bool`

HasRangesInSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


