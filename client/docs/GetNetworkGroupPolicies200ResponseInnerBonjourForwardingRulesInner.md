# GetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description for your Bonjour forwarding rule. Optional. | [optional] 
**VlanId** | **string** | The ID of the service VLAN. Required. | 
**Services** | **[]string** | A list of Bonjour services. At least one service must be specified. Available services are &#39;All Services&#39;, &#39;AirPlay&#39;, &#39;AFP&#39;, &#39;BitTorrent&#39;, &#39;FTP&#39;, &#39;iChat&#39;, &#39;iTunes&#39;, &#39;Printers&#39;, &#39;Samba&#39;, &#39;Scanners&#39; and &#39;SSH&#39; | 

## Methods

### NewGetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner

`func NewGetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner(vlanId string, services []string, ) *GetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner`

NewGetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner instantiates a new GetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInnerWithDefaults

`func NewGetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInnerWithDefaults() *GetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner`

NewGetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInnerWithDefaults instantiates a new GetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVlanId

`func (o *GetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner) GetVlanId() string`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *GetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner) GetVlanIdOk() (*string, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *GetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner) SetVlanId(v string)`

SetVlanId sets VlanId field to given value.


### GetServices

`func (o *GetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *GetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *GetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner) SetServices(v []string)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


