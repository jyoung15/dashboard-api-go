# GetNetworkWirelessSsidBonjourForwarding200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, Bonjour forwarding is enabled on the SSID. | [optional] 
**Exception** | Pointer to [**GetNetworkWirelessSsidBonjourForwarding200ResponseException**](GetNetworkWirelessSsidBonjourForwarding200ResponseException.md) |  | [optional] 
**Rules** | Pointer to [**[]GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner**](GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner.md) | Bonjour forwarding rules | [optional] 

## Methods

### NewGetNetworkWirelessSsidBonjourForwarding200Response

`func NewGetNetworkWirelessSsidBonjourForwarding200Response() *GetNetworkWirelessSsidBonjourForwarding200Response`

NewGetNetworkWirelessSsidBonjourForwarding200Response instantiates a new GetNetworkWirelessSsidBonjourForwarding200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWirelessSsidBonjourForwarding200ResponseWithDefaults

`func NewGetNetworkWirelessSsidBonjourForwarding200ResponseWithDefaults() *GetNetworkWirelessSsidBonjourForwarding200Response`

NewGetNetworkWirelessSsidBonjourForwarding200ResponseWithDefaults instantiates a new GetNetworkWirelessSsidBonjourForwarding200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetNetworkWirelessSsidBonjourForwarding200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkWirelessSsidBonjourForwarding200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkWirelessSsidBonjourForwarding200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkWirelessSsidBonjourForwarding200Response) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetException

`func (o *GetNetworkWirelessSsidBonjourForwarding200Response) GetException() GetNetworkWirelessSsidBonjourForwarding200ResponseException`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *GetNetworkWirelessSsidBonjourForwarding200Response) GetExceptionOk() (*GetNetworkWirelessSsidBonjourForwarding200ResponseException, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *GetNetworkWirelessSsidBonjourForwarding200Response) SetException(v GetNetworkWirelessSsidBonjourForwarding200ResponseException)`

SetException sets Exception field to given value.

### HasException

`func (o *GetNetworkWirelessSsidBonjourForwarding200Response) HasException() bool`

HasException returns a boolean if a field has been set.

### GetRules

`func (o *GetNetworkWirelessSsidBonjourForwarding200Response) GetRules() []GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GetNetworkWirelessSsidBonjourForwarding200Response) GetRulesOk() (*[]GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GetNetworkWirelessSsidBonjourForwarding200Response) SetRules(v []GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *GetNetworkWirelessSsidBonjourForwarding200Response) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


