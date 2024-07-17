# GetNetworkSwitchRoutingOspf200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Boolean value to enable or disable OSPF routing. OSPF routing is disabled by default. | [optional] 
**HelloTimerInSeconds** | Pointer to **int32** | Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds. | [optional] 
**DeadTimerInSeconds** | Pointer to **int32** | Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535 | [optional] 
**Areas** | Pointer to [**[]GetNetworkSwitchRoutingOspf200ResponseAreasInner**](GetNetworkSwitchRoutingOspf200ResponseAreasInner.md) | OSPF areas | [optional] 
**V3** | Pointer to [**GetNetworkSwitchRoutingOspf200ResponseV3**](GetNetworkSwitchRoutingOspf200ResponseV3.md) |  | [optional] 
**Md5AuthenticationEnabled** | Pointer to **bool** | Boolean value to enable or disable MD5 authentication. MD5 authentication is disabled by default. | [optional] 
**Md5AuthenticationKey** | Pointer to [**GetNetworkSwitchRoutingOspf200ResponseMd5AuthenticationKey**](GetNetworkSwitchRoutingOspf200ResponseMd5AuthenticationKey.md) |  | [optional] 

## Methods

### NewGetNetworkSwitchRoutingOspf200Response

`func NewGetNetworkSwitchRoutingOspf200Response() *GetNetworkSwitchRoutingOspf200Response`

NewGetNetworkSwitchRoutingOspf200Response instantiates a new GetNetworkSwitchRoutingOspf200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSwitchRoutingOspf200ResponseWithDefaults

`func NewGetNetworkSwitchRoutingOspf200ResponseWithDefaults() *GetNetworkSwitchRoutingOspf200Response`

NewGetNetworkSwitchRoutingOspf200ResponseWithDefaults instantiates a new GetNetworkSwitchRoutingOspf200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetNetworkSwitchRoutingOspf200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkSwitchRoutingOspf200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkSwitchRoutingOspf200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkSwitchRoutingOspf200Response) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHelloTimerInSeconds

`func (o *GetNetworkSwitchRoutingOspf200Response) GetHelloTimerInSeconds() int32`

GetHelloTimerInSeconds returns the HelloTimerInSeconds field if non-nil, zero value otherwise.

### GetHelloTimerInSecondsOk

`func (o *GetNetworkSwitchRoutingOspf200Response) GetHelloTimerInSecondsOk() (*int32, bool)`

GetHelloTimerInSecondsOk returns a tuple with the HelloTimerInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloTimerInSeconds

`func (o *GetNetworkSwitchRoutingOspf200Response) SetHelloTimerInSeconds(v int32)`

SetHelloTimerInSeconds sets HelloTimerInSeconds field to given value.

### HasHelloTimerInSeconds

`func (o *GetNetworkSwitchRoutingOspf200Response) HasHelloTimerInSeconds() bool`

HasHelloTimerInSeconds returns a boolean if a field has been set.

### GetDeadTimerInSeconds

`func (o *GetNetworkSwitchRoutingOspf200Response) GetDeadTimerInSeconds() int32`

GetDeadTimerInSeconds returns the DeadTimerInSeconds field if non-nil, zero value otherwise.

### GetDeadTimerInSecondsOk

`func (o *GetNetworkSwitchRoutingOspf200Response) GetDeadTimerInSecondsOk() (*int32, bool)`

GetDeadTimerInSecondsOk returns a tuple with the DeadTimerInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadTimerInSeconds

`func (o *GetNetworkSwitchRoutingOspf200Response) SetDeadTimerInSeconds(v int32)`

SetDeadTimerInSeconds sets DeadTimerInSeconds field to given value.

### HasDeadTimerInSeconds

`func (o *GetNetworkSwitchRoutingOspf200Response) HasDeadTimerInSeconds() bool`

HasDeadTimerInSeconds returns a boolean if a field has been set.

### GetAreas

`func (o *GetNetworkSwitchRoutingOspf200Response) GetAreas() []GetNetworkSwitchRoutingOspf200ResponseAreasInner`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *GetNetworkSwitchRoutingOspf200Response) GetAreasOk() (*[]GetNetworkSwitchRoutingOspf200ResponseAreasInner, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *GetNetworkSwitchRoutingOspf200Response) SetAreas(v []GetNetworkSwitchRoutingOspf200ResponseAreasInner)`

SetAreas sets Areas field to given value.

### HasAreas

`func (o *GetNetworkSwitchRoutingOspf200Response) HasAreas() bool`

HasAreas returns a boolean if a field has been set.

### GetV3

`func (o *GetNetworkSwitchRoutingOspf200Response) GetV3() GetNetworkSwitchRoutingOspf200ResponseV3`

GetV3 returns the V3 field if non-nil, zero value otherwise.

### GetV3Ok

`func (o *GetNetworkSwitchRoutingOspf200Response) GetV3Ok() (*GetNetworkSwitchRoutingOspf200ResponseV3, bool)`

GetV3Ok returns a tuple with the V3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV3

`func (o *GetNetworkSwitchRoutingOspf200Response) SetV3(v GetNetworkSwitchRoutingOspf200ResponseV3)`

SetV3 sets V3 field to given value.

### HasV3

`func (o *GetNetworkSwitchRoutingOspf200Response) HasV3() bool`

HasV3 returns a boolean if a field has been set.

### GetMd5AuthenticationEnabled

`func (o *GetNetworkSwitchRoutingOspf200Response) GetMd5AuthenticationEnabled() bool`

GetMd5AuthenticationEnabled returns the Md5AuthenticationEnabled field if non-nil, zero value otherwise.

### GetMd5AuthenticationEnabledOk

`func (o *GetNetworkSwitchRoutingOspf200Response) GetMd5AuthenticationEnabledOk() (*bool, bool)`

GetMd5AuthenticationEnabledOk returns a tuple with the Md5AuthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5AuthenticationEnabled

`func (o *GetNetworkSwitchRoutingOspf200Response) SetMd5AuthenticationEnabled(v bool)`

SetMd5AuthenticationEnabled sets Md5AuthenticationEnabled field to given value.

### HasMd5AuthenticationEnabled

`func (o *GetNetworkSwitchRoutingOspf200Response) HasMd5AuthenticationEnabled() bool`

HasMd5AuthenticationEnabled returns a boolean if a field has been set.

### GetMd5AuthenticationKey

`func (o *GetNetworkSwitchRoutingOspf200Response) GetMd5AuthenticationKey() GetNetworkSwitchRoutingOspf200ResponseMd5AuthenticationKey`

GetMd5AuthenticationKey returns the Md5AuthenticationKey field if non-nil, zero value otherwise.

### GetMd5AuthenticationKeyOk

`func (o *GetNetworkSwitchRoutingOspf200Response) GetMd5AuthenticationKeyOk() (*GetNetworkSwitchRoutingOspf200ResponseMd5AuthenticationKey, bool)`

GetMd5AuthenticationKeyOk returns a tuple with the Md5AuthenticationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5AuthenticationKey

`func (o *GetNetworkSwitchRoutingOspf200Response) SetMd5AuthenticationKey(v GetNetworkSwitchRoutingOspf200ResponseMd5AuthenticationKey)`

SetMd5AuthenticationKey sets Md5AuthenticationKey field to given value.

### HasMd5AuthenticationKey

`func (o *GetNetworkSwitchRoutingOspf200Response) HasMd5AuthenticationKey() bool`

HasMd5AuthenticationKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


