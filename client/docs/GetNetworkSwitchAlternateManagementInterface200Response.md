# GetNetworkSwitchAlternateManagementInterface200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Boolean value to enable or disable AMI configuration. If enabled, VLAN and protocols must be set | [optional] 
**VlanId** | Pointer to **int32** | Alternate management VLAN, must be between 1 and 4094 | [optional] 
**Protocols** | Pointer to **[]string** | Can be one or more of the following values: &#39;radius&#39;, &#39;snmp&#39; or &#39;syslog&#39; | [optional] 
**Switches** | Pointer to [**[]GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner**](GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner.md) | Array of switch serial number and IP assignment. If parameter is present, it cannot have empty body. Note: switches parameter is not applicable for template networks, in other words, do not put &#39;switches&#39; in the body when updating template networks. Also, an empty &#39;switches&#39; array will remove all previous assignments | [optional] 

## Methods

### NewGetNetworkSwitchAlternateManagementInterface200Response

`func NewGetNetworkSwitchAlternateManagementInterface200Response() *GetNetworkSwitchAlternateManagementInterface200Response`

NewGetNetworkSwitchAlternateManagementInterface200Response instantiates a new GetNetworkSwitchAlternateManagementInterface200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSwitchAlternateManagementInterface200ResponseWithDefaults

`func NewGetNetworkSwitchAlternateManagementInterface200ResponseWithDefaults() *GetNetworkSwitchAlternateManagementInterface200Response`

NewGetNetworkSwitchAlternateManagementInterface200ResponseWithDefaults instantiates a new GetNetworkSwitchAlternateManagementInterface200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetNetworkSwitchAlternateManagementInterface200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkSwitchAlternateManagementInterface200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkSwitchAlternateManagementInterface200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkSwitchAlternateManagementInterface200Response) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVlanId

`func (o *GetNetworkSwitchAlternateManagementInterface200Response) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *GetNetworkSwitchAlternateManagementInterface200Response) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *GetNetworkSwitchAlternateManagementInterface200Response) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *GetNetworkSwitchAlternateManagementInterface200Response) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetProtocols

`func (o *GetNetworkSwitchAlternateManagementInterface200Response) GetProtocols() []string`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *GetNetworkSwitchAlternateManagementInterface200Response) GetProtocolsOk() (*[]string, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *GetNetworkSwitchAlternateManagementInterface200Response) SetProtocols(v []string)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *GetNetworkSwitchAlternateManagementInterface200Response) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### GetSwitches

`func (o *GetNetworkSwitchAlternateManagementInterface200Response) GetSwitches() []GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *GetNetworkSwitchAlternateManagementInterface200Response) GetSwitchesOk() (*[]GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *GetNetworkSwitchAlternateManagementInterface200Response) SetSwitches(v []GetNetworkSwitchAlternateManagementInterface200ResponseSwitchesInner)`

SetSwitches sets Switches field to given value.

### HasSwitches

`func (o *GetNetworkSwitchAlternateManagementInterface200Response) HasSwitches() bool`

HasSwitches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

