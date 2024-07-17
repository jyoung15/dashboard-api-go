# GetNetworkAlertsSettings200ResponseAlertsInnerFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]GetNetworkAlertsSettings200ResponseAlertsInnerFiltersConditionsInner**](GetNetworkAlertsSettings200ResponseAlertsInnerFiltersConditionsInner.md) | Conditions | [optional] 
**FailureType** | Pointer to **string** | Failure Type | [optional] 
**LookbackWindow** | Pointer to **int32** | Loopback Window (in sec) | [optional] 
**MinDuration** | Pointer to **int32** | Min Duration | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Period** | Pointer to **int32** | Period | [optional] 
**Priority** | Pointer to **string** | Priority | [optional] 
**Regex** | Pointer to **string** | Regex | [optional] 
**Selector** | Pointer to **string** | Selector | [optional] 
**Serials** | Pointer to **[]string** | Serials | [optional] 
**SsidNum** | Pointer to **int32** | SSID Number | [optional] 
**Tag** | Pointer to **string** | Tag | [optional] 
**Threshold** | Pointer to **int32** | Threshold | [optional] 
**Timeout** | Pointer to **int32** | Timeout | [optional] 

## Methods

### NewGetNetworkAlertsSettings200ResponseAlertsInnerFilters

`func NewGetNetworkAlertsSettings200ResponseAlertsInnerFilters() *GetNetworkAlertsSettings200ResponseAlertsInnerFilters`

NewGetNetworkAlertsSettings200ResponseAlertsInnerFilters instantiates a new GetNetworkAlertsSettings200ResponseAlertsInnerFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkAlertsSettings200ResponseAlertsInnerFiltersWithDefaults

`func NewGetNetworkAlertsSettings200ResponseAlertsInnerFiltersWithDefaults() *GetNetworkAlertsSettings200ResponseAlertsInnerFilters`

NewGetNetworkAlertsSettings200ResponseAlertsInnerFiltersWithDefaults instantiates a new GetNetworkAlertsSettings200ResponseAlertsInnerFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetConditions() []GetNetworkAlertsSettings200ResponseAlertsInnerFiltersConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetConditionsOk() (*[]GetNetworkAlertsSettings200ResponseAlertsInnerFiltersConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) SetConditions(v []GetNetworkAlertsSettings200ResponseAlertsInnerFiltersConditionsInner)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetFailureType

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetFailureType() string`

GetFailureType returns the FailureType field if non-nil, zero value otherwise.

### GetFailureTypeOk

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetFailureTypeOk() (*string, bool)`

GetFailureTypeOk returns a tuple with the FailureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureType

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) SetFailureType(v string)`

SetFailureType sets FailureType field to given value.

### HasFailureType

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) HasFailureType() bool`

HasFailureType returns a boolean if a field has been set.

### GetLookbackWindow

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetLookbackWindow() int32`

GetLookbackWindow returns the LookbackWindow field if non-nil, zero value otherwise.

### GetLookbackWindowOk

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetLookbackWindowOk() (*int32, bool)`

GetLookbackWindowOk returns a tuple with the LookbackWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookbackWindow

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) SetLookbackWindow(v int32)`

SetLookbackWindow sets LookbackWindow field to given value.

### HasLookbackWindow

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) HasLookbackWindow() bool`

HasLookbackWindow returns a boolean if a field has been set.

### GetMinDuration

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetMinDuration() int32`

GetMinDuration returns the MinDuration field if non-nil, zero value otherwise.

### GetMinDurationOk

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetMinDurationOk() (*int32, bool)`

GetMinDurationOk returns a tuple with the MinDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDuration

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) SetMinDuration(v int32)`

SetMinDuration sets MinDuration field to given value.

### HasMinDuration

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) HasMinDuration() bool`

HasMinDuration returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPeriod

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPriority

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRegex

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetSelector

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetSerials

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetSsidNum

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetSsidNum() int32`

GetSsidNum returns the SsidNum field if non-nil, zero value otherwise.

### GetSsidNumOk

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetSsidNumOk() (*int32, bool)`

GetSsidNumOk returns a tuple with the SsidNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNum

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) SetSsidNum(v int32)`

SetSsidNum sets SsidNum field to given value.

### HasSsidNum

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) HasSsidNum() bool`

HasSsidNum returns a boolean if a field has been set.

### GetTag

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetThreshold

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetTimeout

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GetNetworkAlertsSettings200ResponseAlertsInnerFilters) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


