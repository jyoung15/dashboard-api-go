# GetNetworkGroupPolicies200ResponseInnerScheduling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether scheduling is enabled (true) or disabled (false). Defaults to false. If true, the schedule objects for each day of the week (monday - sunday) are parsed. | [optional] 
**Monday** | Pointer to [**GetNetworkGroupPolicies200ResponseInnerSchedulingMonday**](GetNetworkGroupPolicies200ResponseInnerSchedulingMonday.md) |  | [optional] 
**Tuesday** | Pointer to [**GetNetworkGroupPolicies200ResponseInnerSchedulingTuesday**](GetNetworkGroupPolicies200ResponseInnerSchedulingTuesday.md) |  | [optional] 
**Wednesday** | Pointer to [**GetNetworkGroupPolicies200ResponseInnerSchedulingWednesday**](GetNetworkGroupPolicies200ResponseInnerSchedulingWednesday.md) |  | [optional] 
**Thursday** | Pointer to [**GetNetworkGroupPolicies200ResponseInnerSchedulingThursday**](GetNetworkGroupPolicies200ResponseInnerSchedulingThursday.md) |  | [optional] 
**Friday** | Pointer to [**GetNetworkGroupPolicies200ResponseInnerSchedulingFriday**](GetNetworkGroupPolicies200ResponseInnerSchedulingFriday.md) |  | [optional] 
**Saturday** | Pointer to [**GetNetworkGroupPolicies200ResponseInnerSchedulingSaturday**](GetNetworkGroupPolicies200ResponseInnerSchedulingSaturday.md) |  | [optional] 
**Sunday** | Pointer to [**GetNetworkGroupPolicies200ResponseInnerSchedulingSunday**](GetNetworkGroupPolicies200ResponseInnerSchedulingSunday.md) |  | [optional] 

## Methods

### NewGetNetworkGroupPolicies200ResponseInnerScheduling

`func NewGetNetworkGroupPolicies200ResponseInnerScheduling() *GetNetworkGroupPolicies200ResponseInnerScheduling`

NewGetNetworkGroupPolicies200ResponseInnerScheduling instantiates a new GetNetworkGroupPolicies200ResponseInnerScheduling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkGroupPolicies200ResponseInnerSchedulingWithDefaults

`func NewGetNetworkGroupPolicies200ResponseInnerSchedulingWithDefaults() *GetNetworkGroupPolicies200ResponseInnerScheduling`

NewGetNetworkGroupPolicies200ResponseInnerSchedulingWithDefaults instantiates a new GetNetworkGroupPolicies200ResponseInnerScheduling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMonday

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) GetMonday() GetNetworkGroupPolicies200ResponseInnerSchedulingMonday`

GetMonday returns the Monday field if non-nil, zero value otherwise.

### GetMondayOk

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) GetMondayOk() (*GetNetworkGroupPolicies200ResponseInnerSchedulingMonday, bool)`

GetMondayOk returns a tuple with the Monday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonday

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) SetMonday(v GetNetworkGroupPolicies200ResponseInnerSchedulingMonday)`

SetMonday sets Monday field to given value.

### HasMonday

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) HasMonday() bool`

HasMonday returns a boolean if a field has been set.

### GetTuesday

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) GetTuesday() GetNetworkGroupPolicies200ResponseInnerSchedulingTuesday`

GetTuesday returns the Tuesday field if non-nil, zero value otherwise.

### GetTuesdayOk

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) GetTuesdayOk() (*GetNetworkGroupPolicies200ResponseInnerSchedulingTuesday, bool)`

GetTuesdayOk returns a tuple with the Tuesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesday

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) SetTuesday(v GetNetworkGroupPolicies200ResponseInnerSchedulingTuesday)`

SetTuesday sets Tuesday field to given value.

### HasTuesday

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) HasTuesday() bool`

HasTuesday returns a boolean if a field has been set.

### GetWednesday

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) GetWednesday() GetNetworkGroupPolicies200ResponseInnerSchedulingWednesday`

GetWednesday returns the Wednesday field if non-nil, zero value otherwise.

### GetWednesdayOk

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) GetWednesdayOk() (*GetNetworkGroupPolicies200ResponseInnerSchedulingWednesday, bool)`

GetWednesdayOk returns a tuple with the Wednesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesday

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) SetWednesday(v GetNetworkGroupPolicies200ResponseInnerSchedulingWednesday)`

SetWednesday sets Wednesday field to given value.

### HasWednesday

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) HasWednesday() bool`

HasWednesday returns a boolean if a field has been set.

### GetThursday

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) GetThursday() GetNetworkGroupPolicies200ResponseInnerSchedulingThursday`

GetThursday returns the Thursday field if non-nil, zero value otherwise.

### GetThursdayOk

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) GetThursdayOk() (*GetNetworkGroupPolicies200ResponseInnerSchedulingThursday, bool)`

GetThursdayOk returns a tuple with the Thursday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursday

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) SetThursday(v GetNetworkGroupPolicies200ResponseInnerSchedulingThursday)`

SetThursday sets Thursday field to given value.

### HasThursday

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) HasThursday() bool`

HasThursday returns a boolean if a field has been set.

### GetFriday

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) GetFriday() GetNetworkGroupPolicies200ResponseInnerSchedulingFriday`

GetFriday returns the Friday field if non-nil, zero value otherwise.

### GetFridayOk

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) GetFridayOk() (*GetNetworkGroupPolicies200ResponseInnerSchedulingFriday, bool)`

GetFridayOk returns a tuple with the Friday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriday

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) SetFriday(v GetNetworkGroupPolicies200ResponseInnerSchedulingFriday)`

SetFriday sets Friday field to given value.

### HasFriday

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) HasFriday() bool`

HasFriday returns a boolean if a field has been set.

### GetSaturday

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) GetSaturday() GetNetworkGroupPolicies200ResponseInnerSchedulingSaturday`

GetSaturday returns the Saturday field if non-nil, zero value otherwise.

### GetSaturdayOk

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) GetSaturdayOk() (*GetNetworkGroupPolicies200ResponseInnerSchedulingSaturday, bool)`

GetSaturdayOk returns a tuple with the Saturday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturday

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) SetSaturday(v GetNetworkGroupPolicies200ResponseInnerSchedulingSaturday)`

SetSaturday sets Saturday field to given value.

### HasSaturday

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) HasSaturday() bool`

HasSaturday returns a boolean if a field has been set.

### GetSunday

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) GetSunday() GetNetworkGroupPolicies200ResponseInnerSchedulingSunday`

GetSunday returns the Sunday field if non-nil, zero value otherwise.

### GetSundayOk

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) GetSundayOk() (*GetNetworkGroupPolicies200ResponseInnerSchedulingSunday, bool)`

GetSundayOk returns a tuple with the Sunday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunday

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) SetSunday(v GetNetworkGroupPolicies200ResponseInnerSchedulingSunday)`

SetSunday sets Sunday field to given value.

### HasSunday

`func (o *GetNetworkGroupPolicies200ResponseInnerScheduling) HasSunday() bool`

HasSunday returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


