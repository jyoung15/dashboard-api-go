# GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Description of the rule (optional) | [optional] 
**Policy** | Pointer to **string** | &#39;allow&#39; or &#39;deny&#39; traffic specified by this rule | [optional] 
**Protocol** | Pointer to **string** | The type of protocol (must be &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp&#39;, &#39;icmp6&#39; or &#39;any&#39;) | [optional] 
**SrcPort** | Pointer to **string** | Comma-separated list of source port(s) (integer in the range 1-65535), or &#39;any&#39; | [optional] 
**SrcCidr** | Pointer to **string** | Comma-separated list of source IP address(es) (in IP or CIDR notation), or &#39;any&#39; (note: FQDN not supported for source addresses) | [optional] 
**DestPort** | Pointer to **string** | Comma-separated list of destination port(s) (integer in the range 1-65535), or &#39;any&#39; | [optional] 
**DestCidr** | Pointer to **string** | Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or &#39;any&#39; | [optional] 
**SyslogEnabled** | Pointer to **bool** | Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional) | [optional] 

## Methods

### NewGetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner

`func NewGetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner() *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner`

NewGetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner instantiates a new GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInnerWithDefaults

`func NewGetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInnerWithDefaults() *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner`

NewGetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInnerWithDefaults instantiates a new GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPolicy

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetProtocol

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSrcPort

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) GetSrcPort() string`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) GetSrcPortOk() (*string, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) SetSrcPort(v string)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetSrcCidr

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) GetSrcCidr() string`

GetSrcCidr returns the SrcCidr field if non-nil, zero value otherwise.

### GetSrcCidrOk

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) GetSrcCidrOk() (*string, bool)`

GetSrcCidrOk returns a tuple with the SrcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCidr

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) SetSrcCidr(v string)`

SetSrcCidr sets SrcCidr field to given value.

### HasSrcCidr

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) HasSrcCidr() bool`

HasSrcCidr returns a boolean if a field has been set.

### GetDestPort

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) GetDestPort() string`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) GetDestPortOk() (*string, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) SetDestPort(v string)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetDestCidr

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) GetDestCidr() string`

GetDestCidr returns the DestCidr field if non-nil, zero value otherwise.

### GetDestCidrOk

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) GetDestCidrOk() (*string, bool)`

GetDestCidrOk returns a tuple with the DestCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestCidr

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) SetDestCidr(v string)`

SetDestCidr sets DestCidr field to given value.

### HasDestCidr

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) HasDestCidr() bool`

HasDestCidr returns a boolean if a field has been set.

### GetSyslogEnabled

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) GetSyslogEnabled() bool`

GetSyslogEnabled returns the SyslogEnabled field if non-nil, zero value otherwise.

### GetSyslogEnabledOk

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) GetSyslogEnabledOk() (*bool, bool)`

GetSyslogEnabledOk returns a tuple with the SyslogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogEnabled

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) SetSyslogEnabled(v bool)`

SetSyslogEnabled sets SyslogEnabled field to given value.

### HasSyslogEnabled

`func (o *GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) HasSyslogEnabled() bool`

HasSyslogEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


