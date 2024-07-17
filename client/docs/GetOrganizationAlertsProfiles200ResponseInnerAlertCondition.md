# GetOrganizationAlertsProfiles200ResponseInnerAlertCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | The total duration in seconds that the threshold should be crossed before alerting | [optional] 
**Window** | Pointer to **int32** | The look back period in seconds for sensing the alert | [optional] 
**BitRateBps** | Pointer to **int32** | The threshold the metric must cross to be valid for alerting. Used only for WAN Utilization alerts. | [optional] 
**Interface** | Pointer to **string** | The uplink observed for the alert | [optional] 

## Methods

### NewGetOrganizationAlertsProfiles200ResponseInnerAlertCondition

`func NewGetOrganizationAlertsProfiles200ResponseInnerAlertCondition() *GetOrganizationAlertsProfiles200ResponseInnerAlertCondition`

NewGetOrganizationAlertsProfiles200ResponseInnerAlertCondition instantiates a new GetOrganizationAlertsProfiles200ResponseInnerAlertCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationAlertsProfiles200ResponseInnerAlertConditionWithDefaults

`func NewGetOrganizationAlertsProfiles200ResponseInnerAlertConditionWithDefaults() *GetOrganizationAlertsProfiles200ResponseInnerAlertCondition`

NewGetOrganizationAlertsProfiles200ResponseInnerAlertConditionWithDefaults instantiates a new GetOrganizationAlertsProfiles200ResponseInnerAlertCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *GetOrganizationAlertsProfiles200ResponseInnerAlertCondition) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetOrganizationAlertsProfiles200ResponseInnerAlertCondition) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetOrganizationAlertsProfiles200ResponseInnerAlertCondition) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetOrganizationAlertsProfiles200ResponseInnerAlertCondition) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetWindow

`func (o *GetOrganizationAlertsProfiles200ResponseInnerAlertCondition) GetWindow() int32`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *GetOrganizationAlertsProfiles200ResponseInnerAlertCondition) GetWindowOk() (*int32, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *GetOrganizationAlertsProfiles200ResponseInnerAlertCondition) SetWindow(v int32)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *GetOrganizationAlertsProfiles200ResponseInnerAlertCondition) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetBitRateBps

`func (o *GetOrganizationAlertsProfiles200ResponseInnerAlertCondition) GetBitRateBps() int32`

GetBitRateBps returns the BitRateBps field if non-nil, zero value otherwise.

### GetBitRateBpsOk

`func (o *GetOrganizationAlertsProfiles200ResponseInnerAlertCondition) GetBitRateBpsOk() (*int32, bool)`

GetBitRateBpsOk returns a tuple with the BitRateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitRateBps

`func (o *GetOrganizationAlertsProfiles200ResponseInnerAlertCondition) SetBitRateBps(v int32)`

SetBitRateBps sets BitRateBps field to given value.

### HasBitRateBps

`func (o *GetOrganizationAlertsProfiles200ResponseInnerAlertCondition) HasBitRateBps() bool`

HasBitRateBps returns a boolean if a field has been set.

### GetInterface

`func (o *GetOrganizationAlertsProfiles200ResponseInnerAlertCondition) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *GetOrganizationAlertsProfiles200ResponseInnerAlertCondition) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *GetOrganizationAlertsProfiles200ResponseInnerAlertCondition) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *GetOrganizationAlertsProfiles200ResponseInnerAlertCondition) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


