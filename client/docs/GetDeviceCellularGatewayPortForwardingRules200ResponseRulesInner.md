# GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A descriptive name for the rule | [optional] 
**LanIp** | Pointer to **string** | The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN | [optional] 
**PublicPort** | Pointer to **string** | A port or port ranges that will be forwarded to the host on the LAN | [optional] 
**LocalPort** | Pointer to **string** | A port or port ranges that will receive the forwarded traffic from the WAN | [optional] 
**AllowedIps** | Pointer to **[]string** | An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges. | [optional] 
**Protocol** | Pointer to **string** | TCP or UDP | [optional] 
**Access** | Pointer to **string** | &#x60;any&#x60; or &#x60;restricted&#x60;. Specify the right to make inbound connections on the specified ports or port ranges. If &#x60;restricted&#x60;, a list of allowed IPs is mandatory. | [optional] 

## Methods

### NewGetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner

`func NewGetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner() *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner`

NewGetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner instantiates a new GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceCellularGatewayPortForwardingRules200ResponseRulesInnerWithDefaults

`func NewGetDeviceCellularGatewayPortForwardingRules200ResponseRulesInnerWithDefaults() *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner`

NewGetDeviceCellularGatewayPortForwardingRules200ResponseRulesInnerWithDefaults instantiates a new GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLanIp

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) GetLanIp() string`

GetLanIp returns the LanIp field if non-nil, zero value otherwise.

### GetLanIpOk

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) GetLanIpOk() (*string, bool)`

GetLanIpOk returns a tuple with the LanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIp

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) SetLanIp(v string)`

SetLanIp sets LanIp field to given value.

### HasLanIp

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) HasLanIp() bool`

HasLanIp returns a boolean if a field has been set.

### GetPublicPort

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) GetPublicPort() string`

GetPublicPort returns the PublicPort field if non-nil, zero value otherwise.

### GetPublicPortOk

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) GetPublicPortOk() (*string, bool)`

GetPublicPortOk returns a tuple with the PublicPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPort

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) SetPublicPort(v string)`

SetPublicPort sets PublicPort field to given value.

### HasPublicPort

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) HasPublicPort() bool`

HasPublicPort returns a boolean if a field has been set.

### GetLocalPort

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) GetLocalPort() string`

GetLocalPort returns the LocalPort field if non-nil, zero value otherwise.

### GetLocalPortOk

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) GetLocalPortOk() (*string, bool)`

GetLocalPortOk returns a tuple with the LocalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPort

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) SetLocalPort(v string)`

SetLocalPort sets LocalPort field to given value.

### HasLocalPort

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) HasLocalPort() bool`

HasLocalPort returns a boolean if a field has been set.

### GetAllowedIps

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) GetAllowedIps() []string`

GetAllowedIps returns the AllowedIps field if non-nil, zero value otherwise.

### GetAllowedIpsOk

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) GetAllowedIpsOk() (*[]string, bool)`

GetAllowedIpsOk returns a tuple with the AllowedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIps

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) SetAllowedIps(v []string)`

SetAllowedIps sets AllowedIps field to given value.

### HasAllowedIps

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) HasAllowedIps() bool`

HasAllowedIps returns a boolean if a field has been set.

### GetProtocol

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetAccess

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *GetDeviceCellularGatewayPortForwardingRules200ResponseRulesInner) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


