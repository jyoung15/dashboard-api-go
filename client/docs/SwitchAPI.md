# \SwitchAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNetworkSwitchStack**](SwitchAPI.md#AddNetworkSwitchStack) | **Post** /networks/{networkId}/switch/stacks/{switchStackId}/add | Add a switch to a stack
[**CloneOrganizationSwitchDevices**](SwitchAPI.md#CloneOrganizationSwitchDevices) | **Post** /organizations/{organizationId}/switch/devices/clone | Clone port-level and some switch-level configuration settings from a source switch to one or more target switches
[**CreateDeviceSwitchRoutingInterface**](SwitchAPI.md#CreateDeviceSwitchRoutingInterface) | **Post** /devices/{serial}/switch/routing/interfaces | Create a layer 3 interface for a switch
[**CreateDeviceSwitchRoutingStaticRoute**](SwitchAPI.md#CreateDeviceSwitchRoutingStaticRoute) | **Post** /devices/{serial}/switch/routing/staticRoutes | Create a layer 3 static route for a switch
[**CreateNetworkSwitchAccessPolicy**](SwitchAPI.md#CreateNetworkSwitchAccessPolicy) | **Post** /networks/{networkId}/switch/accessPolicies | Create an access policy for a switch network
[**CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer**](SwitchAPI.md#CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer) | **Post** /networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers | Add a server to be trusted by Dynamic ARP Inspection on this network
[**CreateNetworkSwitchLinkAggregation**](SwitchAPI.md#CreateNetworkSwitchLinkAggregation) | **Post** /networks/{networkId}/switch/linkAggregations | Create a link aggregation group
[**CreateNetworkSwitchPortSchedule**](SwitchAPI.md#CreateNetworkSwitchPortSchedule) | **Post** /networks/{networkId}/switch/portSchedules | Add a switch port schedule
[**CreateNetworkSwitchQosRule**](SwitchAPI.md#CreateNetworkSwitchQosRule) | **Post** /networks/{networkId}/switch/qosRules | Add a quality of service rule
[**CreateNetworkSwitchRoutingMulticastRendezvousPoint**](SwitchAPI.md#CreateNetworkSwitchRoutingMulticastRendezvousPoint) | **Post** /networks/{networkId}/switch/routing/multicast/rendezvousPoints | Create a multicast rendezvous point
[**CreateNetworkSwitchStack**](SwitchAPI.md#CreateNetworkSwitchStack) | **Post** /networks/{networkId}/switch/stacks | Create a switch stack
[**CreateNetworkSwitchStackRoutingInterface**](SwitchAPI.md#CreateNetworkSwitchStackRoutingInterface) | **Post** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces | Create a layer 3 interface for a switch stack
[**CreateNetworkSwitchStackRoutingStaticRoute**](SwitchAPI.md#CreateNetworkSwitchStackRoutingStaticRoute) | **Post** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes | Create a layer 3 static route for a switch stack
[**CycleDeviceSwitchPorts**](SwitchAPI.md#CycleDeviceSwitchPorts) | **Post** /devices/{serial}/switch/ports/cycle | Cycle a set of switch ports
[**DeleteDeviceSwitchRoutingInterface**](SwitchAPI.md#DeleteDeviceSwitchRoutingInterface) | **Delete** /devices/{serial}/switch/routing/interfaces/{interfaceId} | Delete a layer 3 interface from the switch
[**DeleteDeviceSwitchRoutingStaticRoute**](SwitchAPI.md#DeleteDeviceSwitchRoutingStaticRoute) | **Delete** /devices/{serial}/switch/routing/staticRoutes/{staticRouteId} | Delete a layer 3 static route for a switch
[**DeleteNetworkSwitchAccessPolicy**](SwitchAPI.md#DeleteNetworkSwitchAccessPolicy) | **Delete** /networks/{networkId}/switch/accessPolicies/{accessPolicyNumber} | Delete an access policy for a switch network
[**DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer**](SwitchAPI.md#DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer) | **Delete** /networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers/{trustedServerId} | Remove a server from being trusted by Dynamic ARP Inspection on this network
[**DeleteNetworkSwitchLinkAggregation**](SwitchAPI.md#DeleteNetworkSwitchLinkAggregation) | **Delete** /networks/{networkId}/switch/linkAggregations/{linkAggregationId} | Split a link aggregation group into separate ports
[**DeleteNetworkSwitchPortSchedule**](SwitchAPI.md#DeleteNetworkSwitchPortSchedule) | **Delete** /networks/{networkId}/switch/portSchedules/{portScheduleId} | Delete a switch port schedule
[**DeleteNetworkSwitchQosRule**](SwitchAPI.md#DeleteNetworkSwitchQosRule) | **Delete** /networks/{networkId}/switch/qosRules/{qosRuleId} | Delete a quality of service rule
[**DeleteNetworkSwitchRoutingMulticastRendezvousPoint**](SwitchAPI.md#DeleteNetworkSwitchRoutingMulticastRendezvousPoint) | **Delete** /networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId} | Delete a multicast rendezvous point
[**DeleteNetworkSwitchStack**](SwitchAPI.md#DeleteNetworkSwitchStack) | **Delete** /networks/{networkId}/switch/stacks/{switchStackId} | Delete a stack
[**DeleteNetworkSwitchStackRoutingInterface**](SwitchAPI.md#DeleteNetworkSwitchStackRoutingInterface) | **Delete** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId} | Delete a layer 3 interface from a switch stack
[**DeleteNetworkSwitchStackRoutingStaticRoute**](SwitchAPI.md#DeleteNetworkSwitchStackRoutingStaticRoute) | **Delete** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes/{staticRouteId} | Delete a layer 3 static route for a switch stack
[**GetDeviceSwitchPort**](SwitchAPI.md#GetDeviceSwitchPort) | **Get** /devices/{serial}/switch/ports/{portId} | Return a switch port
[**GetDeviceSwitchPorts**](SwitchAPI.md#GetDeviceSwitchPorts) | **Get** /devices/{serial}/switch/ports | List the switch ports for a switch
[**GetDeviceSwitchPortsStatuses**](SwitchAPI.md#GetDeviceSwitchPortsStatuses) | **Get** /devices/{serial}/switch/ports/statuses | Return the status for all the ports of a switch
[**GetDeviceSwitchPortsStatusesPackets**](SwitchAPI.md#GetDeviceSwitchPortsStatusesPackets) | **Get** /devices/{serial}/switch/ports/statuses/packets | Return the packet counters for all the ports of a switch
[**GetDeviceSwitchRoutingInterface**](SwitchAPI.md#GetDeviceSwitchRoutingInterface) | **Get** /devices/{serial}/switch/routing/interfaces/{interfaceId} | Return a layer 3 interface for a switch
[**GetDeviceSwitchRoutingInterfaceDhcp**](SwitchAPI.md#GetDeviceSwitchRoutingInterfaceDhcp) | **Get** /devices/{serial}/switch/routing/interfaces/{interfaceId}/dhcp | Return a layer 3 interface DHCP configuration for a switch
[**GetDeviceSwitchRoutingInterfaces**](SwitchAPI.md#GetDeviceSwitchRoutingInterfaces) | **Get** /devices/{serial}/switch/routing/interfaces | List layer 3 interfaces for a switch
[**GetDeviceSwitchRoutingStaticRoute**](SwitchAPI.md#GetDeviceSwitchRoutingStaticRoute) | **Get** /devices/{serial}/switch/routing/staticRoutes/{staticRouteId} | Return a layer 3 static route for a switch
[**GetDeviceSwitchRoutingStaticRoutes**](SwitchAPI.md#GetDeviceSwitchRoutingStaticRoutes) | **Get** /devices/{serial}/switch/routing/staticRoutes | List layer 3 static routes for a switch
[**GetDeviceSwitchWarmSpare**](SwitchAPI.md#GetDeviceSwitchWarmSpare) | **Get** /devices/{serial}/switch/warmSpare | Return warm spare configuration for a switch
[**GetNetworkSwitchAccessControlLists**](SwitchAPI.md#GetNetworkSwitchAccessControlLists) | **Get** /networks/{networkId}/switch/accessControlLists | Return the access control lists for a MS network
[**GetNetworkSwitchAccessPolicies**](SwitchAPI.md#GetNetworkSwitchAccessPolicies) | **Get** /networks/{networkId}/switch/accessPolicies | List the access policies for a switch network
[**GetNetworkSwitchAccessPolicy**](SwitchAPI.md#GetNetworkSwitchAccessPolicy) | **Get** /networks/{networkId}/switch/accessPolicies/{accessPolicyNumber} | Return a specific access policy for a switch network
[**GetNetworkSwitchAlternateManagementInterface**](SwitchAPI.md#GetNetworkSwitchAlternateManagementInterface) | **Get** /networks/{networkId}/switch/alternateManagementInterface | Return the switch alternate management interface for the network
[**GetNetworkSwitchDhcpServerPolicy**](SwitchAPI.md#GetNetworkSwitchDhcpServerPolicy) | **Get** /networks/{networkId}/switch/dhcpServerPolicy | Return the DHCP server settings
[**GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers**](SwitchAPI.md#GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers) | **Get** /networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers | Return the list of servers trusted by Dynamic ARP Inspection on this network
[**GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice**](SwitchAPI.md#GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice) | **Get** /networks/{networkId}/switch/dhcpServerPolicy/arpInspection/warnings/byDevice | Return the devices that have a Dynamic ARP Inspection warning and their warnings
[**GetNetworkSwitchDhcpV4ServersSeen**](SwitchAPI.md#GetNetworkSwitchDhcpV4ServersSeen) | **Get** /networks/{networkId}/switch/dhcp/v4/servers/seen | Return the network&#39;s DHCPv4 servers seen within the selected timeframe (default 1 day)
[**GetNetworkSwitchDscpToCosMappings**](SwitchAPI.md#GetNetworkSwitchDscpToCosMappings) | **Get** /networks/{networkId}/switch/dscpToCosMappings | Return the DSCP to CoS mappings
[**GetNetworkSwitchLinkAggregations**](SwitchAPI.md#GetNetworkSwitchLinkAggregations) | **Get** /networks/{networkId}/switch/linkAggregations | List link aggregation groups
[**GetNetworkSwitchMtu**](SwitchAPI.md#GetNetworkSwitchMtu) | **Get** /networks/{networkId}/switch/mtu | Return the MTU configuration
[**GetNetworkSwitchPortSchedules**](SwitchAPI.md#GetNetworkSwitchPortSchedules) | **Get** /networks/{networkId}/switch/portSchedules | List switch port schedules
[**GetNetworkSwitchQosRule**](SwitchAPI.md#GetNetworkSwitchQosRule) | **Get** /networks/{networkId}/switch/qosRules/{qosRuleId} | Return a quality of service rule
[**GetNetworkSwitchQosRules**](SwitchAPI.md#GetNetworkSwitchQosRules) | **Get** /networks/{networkId}/switch/qosRules | List quality of service rules
[**GetNetworkSwitchQosRulesOrder**](SwitchAPI.md#GetNetworkSwitchQosRulesOrder) | **Get** /networks/{networkId}/switch/qosRules/order | Return the quality of service rule IDs by order in which they will be processed by the switch
[**GetNetworkSwitchRoutingMulticast**](SwitchAPI.md#GetNetworkSwitchRoutingMulticast) | **Get** /networks/{networkId}/switch/routing/multicast | Return multicast settings for a network
[**GetNetworkSwitchRoutingMulticastRendezvousPoint**](SwitchAPI.md#GetNetworkSwitchRoutingMulticastRendezvousPoint) | **Get** /networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId} | Return a multicast rendezvous point
[**GetNetworkSwitchRoutingMulticastRendezvousPoints**](SwitchAPI.md#GetNetworkSwitchRoutingMulticastRendezvousPoints) | **Get** /networks/{networkId}/switch/routing/multicast/rendezvousPoints | List multicast rendezvous points
[**GetNetworkSwitchRoutingOspf**](SwitchAPI.md#GetNetworkSwitchRoutingOspf) | **Get** /networks/{networkId}/switch/routing/ospf | Return layer 3 OSPF routing configuration
[**GetNetworkSwitchSettings**](SwitchAPI.md#GetNetworkSwitchSettings) | **Get** /networks/{networkId}/switch/settings | Returns the switch network settings
[**GetNetworkSwitchStack**](SwitchAPI.md#GetNetworkSwitchStack) | **Get** /networks/{networkId}/switch/stacks/{switchStackId} | Show a switch stack
[**GetNetworkSwitchStackRoutingInterface**](SwitchAPI.md#GetNetworkSwitchStackRoutingInterface) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId} | Return a layer 3 interface from a switch stack
[**GetNetworkSwitchStackRoutingInterfaceDhcp**](SwitchAPI.md#GetNetworkSwitchStackRoutingInterfaceDhcp) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp | Return a layer 3 interface DHCP configuration for a switch stack
[**GetNetworkSwitchStackRoutingInterfaces**](SwitchAPI.md#GetNetworkSwitchStackRoutingInterfaces) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces | List layer 3 interfaces for a switch stack
[**GetNetworkSwitchStackRoutingStaticRoute**](SwitchAPI.md#GetNetworkSwitchStackRoutingStaticRoute) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes/{staticRouteId} | Return a layer 3 static route for a switch stack
[**GetNetworkSwitchStackRoutingStaticRoutes**](SwitchAPI.md#GetNetworkSwitchStackRoutingStaticRoutes) | **Get** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes | List layer 3 static routes for a switch stack
[**GetNetworkSwitchStacks**](SwitchAPI.md#GetNetworkSwitchStacks) | **Get** /networks/{networkId}/switch/stacks | List the switch stacks in a network
[**GetNetworkSwitchStormControl**](SwitchAPI.md#GetNetworkSwitchStormControl) | **Get** /networks/{networkId}/switch/stormControl | Return the storm control configuration for a switch network
[**GetNetworkSwitchStp**](SwitchAPI.md#GetNetworkSwitchStp) | **Get** /networks/{networkId}/switch/stp | Returns STP settings
[**GetOrganizationConfigTemplateSwitchProfilePort**](SwitchAPI.md#GetOrganizationConfigTemplateSwitchProfilePort) | **Get** /organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles/{profileId}/ports/{portId} | Return a switch template port
[**GetOrganizationConfigTemplateSwitchProfilePorts**](SwitchAPI.md#GetOrganizationConfigTemplateSwitchProfilePorts) | **Get** /organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles/{profileId}/ports | Return all the ports of a switch template
[**GetOrganizationConfigTemplateSwitchProfiles**](SwitchAPI.md#GetOrganizationConfigTemplateSwitchProfiles) | **Get** /organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles | List the switch templates for your switch template configuration
[**GetOrganizationSummarySwitchPowerHistory**](SwitchAPI.md#GetOrganizationSummarySwitchPowerHistory) | **Get** /organizations/{organizationId}/summary/switch/power/history | Returns the total PoE power draw for all switch ports in the organization over the requested timespan (by default the last 24 hours)
[**GetOrganizationSwitchPortsBySwitch**](SwitchAPI.md#GetOrganizationSwitchPortsBySwitch) | **Get** /organizations/{organizationId}/switch/ports/bySwitch | List the switchports in an organization by switch
[**GetOrganizationSwitchPortsOverview**](SwitchAPI.md#GetOrganizationSwitchPortsOverview) | **Get** /organizations/{organizationId}/switch/ports/overview | Returns the counts of all active ports for the requested timespan, grouped by speed
[**RemoveNetworkSwitchStack**](SwitchAPI.md#RemoveNetworkSwitchStack) | **Post** /networks/{networkId}/switch/stacks/{switchStackId}/remove | Remove a switch from a stack
[**UpdateDeviceSwitchPort**](SwitchAPI.md#UpdateDeviceSwitchPort) | **Put** /devices/{serial}/switch/ports/{portId} | Update a switch port
[**UpdateDeviceSwitchRoutingInterface**](SwitchAPI.md#UpdateDeviceSwitchRoutingInterface) | **Put** /devices/{serial}/switch/routing/interfaces/{interfaceId} | Update a layer 3 interface for a switch
[**UpdateDeviceSwitchRoutingInterfaceDhcp**](SwitchAPI.md#UpdateDeviceSwitchRoutingInterfaceDhcp) | **Put** /devices/{serial}/switch/routing/interfaces/{interfaceId}/dhcp | Update a layer 3 interface DHCP configuration for a switch
[**UpdateDeviceSwitchRoutingStaticRoute**](SwitchAPI.md#UpdateDeviceSwitchRoutingStaticRoute) | **Put** /devices/{serial}/switch/routing/staticRoutes/{staticRouteId} | Update a layer 3 static route for a switch
[**UpdateDeviceSwitchWarmSpare**](SwitchAPI.md#UpdateDeviceSwitchWarmSpare) | **Put** /devices/{serial}/switch/warmSpare | Update warm spare configuration for a switch
[**UpdateNetworkSwitchAccessControlLists**](SwitchAPI.md#UpdateNetworkSwitchAccessControlLists) | **Put** /networks/{networkId}/switch/accessControlLists | Update the access control lists for a MS network
[**UpdateNetworkSwitchAccessPolicy**](SwitchAPI.md#UpdateNetworkSwitchAccessPolicy) | **Put** /networks/{networkId}/switch/accessPolicies/{accessPolicyNumber} | Update an access policy for a switch network
[**UpdateNetworkSwitchAlternateManagementInterface**](SwitchAPI.md#UpdateNetworkSwitchAlternateManagementInterface) | **Put** /networks/{networkId}/switch/alternateManagementInterface | Update the switch alternate management interface for the network
[**UpdateNetworkSwitchDhcpServerPolicy**](SwitchAPI.md#UpdateNetworkSwitchDhcpServerPolicy) | **Put** /networks/{networkId}/switch/dhcpServerPolicy | Update the DHCP server settings
[**UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer**](SwitchAPI.md#UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer) | **Put** /networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers/{trustedServerId} | Update a server that is trusted by Dynamic ARP Inspection on this network
[**UpdateNetworkSwitchDscpToCosMappings**](SwitchAPI.md#UpdateNetworkSwitchDscpToCosMappings) | **Put** /networks/{networkId}/switch/dscpToCosMappings | Update the DSCP to CoS mappings
[**UpdateNetworkSwitchLinkAggregation**](SwitchAPI.md#UpdateNetworkSwitchLinkAggregation) | **Put** /networks/{networkId}/switch/linkAggregations/{linkAggregationId} | Update a link aggregation group
[**UpdateNetworkSwitchMtu**](SwitchAPI.md#UpdateNetworkSwitchMtu) | **Put** /networks/{networkId}/switch/mtu | Update the MTU configuration
[**UpdateNetworkSwitchPortSchedule**](SwitchAPI.md#UpdateNetworkSwitchPortSchedule) | **Put** /networks/{networkId}/switch/portSchedules/{portScheduleId} | Update a switch port schedule
[**UpdateNetworkSwitchQosRule**](SwitchAPI.md#UpdateNetworkSwitchQosRule) | **Put** /networks/{networkId}/switch/qosRules/{qosRuleId} | Update a quality of service rule
[**UpdateNetworkSwitchQosRulesOrder**](SwitchAPI.md#UpdateNetworkSwitchQosRulesOrder) | **Put** /networks/{networkId}/switch/qosRules/order | Update the order in which the rules should be processed by the switch
[**UpdateNetworkSwitchRoutingMulticast**](SwitchAPI.md#UpdateNetworkSwitchRoutingMulticast) | **Put** /networks/{networkId}/switch/routing/multicast | Update multicast settings for a network
[**UpdateNetworkSwitchRoutingMulticastRendezvousPoint**](SwitchAPI.md#UpdateNetworkSwitchRoutingMulticastRendezvousPoint) | **Put** /networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId} | Update a multicast rendezvous point
[**UpdateNetworkSwitchRoutingOspf**](SwitchAPI.md#UpdateNetworkSwitchRoutingOspf) | **Put** /networks/{networkId}/switch/routing/ospf | Update layer 3 OSPF routing configuration
[**UpdateNetworkSwitchSettings**](SwitchAPI.md#UpdateNetworkSwitchSettings) | **Put** /networks/{networkId}/switch/settings | Update switch network settings
[**UpdateNetworkSwitchStackRoutingInterface**](SwitchAPI.md#UpdateNetworkSwitchStackRoutingInterface) | **Put** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId} | Update a layer 3 interface for a switch stack
[**UpdateNetworkSwitchStackRoutingInterfaceDhcp**](SwitchAPI.md#UpdateNetworkSwitchStackRoutingInterfaceDhcp) | **Put** /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp | Update a layer 3 interface DHCP configuration for a switch stack
[**UpdateNetworkSwitchStackRoutingStaticRoute**](SwitchAPI.md#UpdateNetworkSwitchStackRoutingStaticRoute) | **Put** /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes/{staticRouteId} | Update a layer 3 static route for a switch stack
[**UpdateNetworkSwitchStormControl**](SwitchAPI.md#UpdateNetworkSwitchStormControl) | **Put** /networks/{networkId}/switch/stormControl | Update the storm control configuration for a switch network
[**UpdateNetworkSwitchStp**](SwitchAPI.md#UpdateNetworkSwitchStp) | **Put** /networks/{networkId}/switch/stp | Updates STP settings
[**UpdateOrganizationConfigTemplateSwitchProfilePort**](SwitchAPI.md#UpdateOrganizationConfigTemplateSwitchProfilePort) | **Put** /organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles/{profileId}/ports/{portId} | Update a switch template port



## AddNetworkSwitchStack

> GetNetworkSwitchStacks200ResponseInner AddNetworkSwitchStack(ctx, networkId, switchStackId).AddNetworkSwitchStackRequest(addNetworkSwitchStackRequest).Execute()

Add a switch to a stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	switchStackId := "switchStackId_example" // string | Switch stack ID
	addNetworkSwitchStackRequest := *openapiclient.NewAddNetworkSwitchStackRequest("Serial_example") // AddNetworkSwitchStackRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.AddNetworkSwitchStack(context.Background(), networkId, switchStackId).AddNetworkSwitchStackRequest(addNetworkSwitchStackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.AddNetworkSwitchStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddNetworkSwitchStack`: GetNetworkSwitchStacks200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.AddNetworkSwitchStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddNetworkSwitchStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addNetworkSwitchStackRequest** | [**AddNetworkSwitchStackRequest**](AddNetworkSwitchStackRequest.md) |  | 

### Return type

[**GetNetworkSwitchStacks200ResponseInner**](GetNetworkSwitchStacks200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneOrganizationSwitchDevices

> CloneOrganizationSwitchDevices200Response CloneOrganizationSwitchDevices(ctx, organizationId).CloneOrganizationSwitchDevicesRequest(cloneOrganizationSwitchDevicesRequest).Execute()

Clone port-level and some switch-level configuration settings from a source switch to one or more target switches



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	cloneOrganizationSwitchDevicesRequest := *openapiclient.NewCloneOrganizationSwitchDevicesRequest("SourceSerial_example", []string{"TargetSerials_example"}) // CloneOrganizationSwitchDevicesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.CloneOrganizationSwitchDevices(context.Background(), organizationId).CloneOrganizationSwitchDevicesRequest(cloneOrganizationSwitchDevicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.CloneOrganizationSwitchDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloneOrganizationSwitchDevices`: CloneOrganizationSwitchDevices200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.CloneOrganizationSwitchDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneOrganizationSwitchDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloneOrganizationSwitchDevicesRequest** | [**CloneOrganizationSwitchDevicesRequest**](CloneOrganizationSwitchDevicesRequest.md) |  | 

### Return type

[**CloneOrganizationSwitchDevices200Response**](CloneOrganizationSwitchDevices200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceSwitchRoutingInterface

> GetDeviceSwitchRoutingInterfaces200ResponseInner CreateDeviceSwitchRoutingInterface(ctx, serial).CreateDeviceSwitchRoutingInterfaceRequest(createDeviceSwitchRoutingInterfaceRequest).Execute()

Create a layer 3 interface for a switch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial
	createDeviceSwitchRoutingInterfaceRequest := *openapiclient.NewCreateDeviceSwitchRoutingInterfaceRequest() // CreateDeviceSwitchRoutingInterfaceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.CreateDeviceSwitchRoutingInterface(context.Background(), serial).CreateDeviceSwitchRoutingInterfaceRequest(createDeviceSwitchRoutingInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.CreateDeviceSwitchRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceSwitchRoutingInterface`: GetDeviceSwitchRoutingInterfaces200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.CreateDeviceSwitchRoutingInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceSwitchRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceSwitchRoutingInterfaceRequest** | [**CreateDeviceSwitchRoutingInterfaceRequest**](CreateDeviceSwitchRoutingInterfaceRequest.md) |  | 

### Return type

[**GetDeviceSwitchRoutingInterfaces200ResponseInner**](GetDeviceSwitchRoutingInterfaces200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceSwitchRoutingStaticRoute

> GetDeviceSwitchRoutingStaticRoutes200ResponseInner CreateDeviceSwitchRoutingStaticRoute(ctx, serial).CreateDeviceSwitchRoutingStaticRouteRequest(createDeviceSwitchRoutingStaticRouteRequest).Execute()

Create a layer 3 static route for a switch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial
	createDeviceSwitchRoutingStaticRouteRequest := *openapiclient.NewCreateDeviceSwitchRoutingStaticRouteRequest("Subnet_example", "NextHopIp_example") // CreateDeviceSwitchRoutingStaticRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.CreateDeviceSwitchRoutingStaticRoute(context.Background(), serial).CreateDeviceSwitchRoutingStaticRouteRequest(createDeviceSwitchRoutingStaticRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.CreateDeviceSwitchRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceSwitchRoutingStaticRoute`: GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.CreateDeviceSwitchRoutingStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceSwitchRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceSwitchRoutingStaticRouteRequest** | [**CreateDeviceSwitchRoutingStaticRouteRequest**](CreateDeviceSwitchRoutingStaticRouteRequest.md) |  | 

### Return type

[**GetDeviceSwitchRoutingStaticRoutes200ResponseInner**](GetDeviceSwitchRoutingStaticRoutes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkSwitchAccessPolicy

> GetNetworkSwitchAccessPolicies200ResponseInner CreateNetworkSwitchAccessPolicy(ctx, networkId).CreateNetworkSwitchAccessPolicyRequest(createNetworkSwitchAccessPolicyRequest).Execute()

Create an access policy for a switch network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	createNetworkSwitchAccessPolicyRequest := *openapiclient.NewCreateNetworkSwitchAccessPolicyRequest("Name_example", []openapiclient.CreateNetworkSwitchAccessPolicyRequestRadiusServersInner{*openapiclient.NewCreateNetworkSwitchAccessPolicyRequestRadiusServersInner()}, false, false, false, "HostMode_example", false) // CreateNetworkSwitchAccessPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.CreateNetworkSwitchAccessPolicy(context.Background(), networkId).CreateNetworkSwitchAccessPolicyRequest(createNetworkSwitchAccessPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.CreateNetworkSwitchAccessPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSwitchAccessPolicy`: GetNetworkSwitchAccessPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.CreateNetworkSwitchAccessPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchAccessPolicyRequest** | [**CreateNetworkSwitchAccessPolicyRequest**](CreateNetworkSwitchAccessPolicyRequest.md) |  | 

### Return type

[**GetNetworkSwitchAccessPolicies200ResponseInner**](GetNetworkSwitchAccessPolicies200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer

> GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx, networkId).CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest(createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest).Execute()

Add a server to be trusted by Dynamic ARP Inspection on this network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest := *openapiclient.NewCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest("Mac_example", int32(123), *openapiclient.NewCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4()) // CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(context.Background(), networkId).CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest(createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer`: GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest** | [**CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest**](CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest.md) |  | 

### Return type

[**GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner**](GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkSwitchLinkAggregation

> GetNetworkSwitchLinkAggregations200ResponseInner CreateNetworkSwitchLinkAggregation(ctx, networkId).CreateNetworkSwitchLinkAggregationRequest(createNetworkSwitchLinkAggregationRequest).Execute()

Create a link aggregation group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	createNetworkSwitchLinkAggregationRequest := *openapiclient.NewCreateNetworkSwitchLinkAggregationRequest() // CreateNetworkSwitchLinkAggregationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.CreateNetworkSwitchLinkAggregation(context.Background(), networkId).CreateNetworkSwitchLinkAggregationRequest(createNetworkSwitchLinkAggregationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.CreateNetworkSwitchLinkAggregation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSwitchLinkAggregation`: GetNetworkSwitchLinkAggregations200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.CreateNetworkSwitchLinkAggregation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchLinkAggregationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchLinkAggregationRequest** | [**CreateNetworkSwitchLinkAggregationRequest**](CreateNetworkSwitchLinkAggregationRequest.md) |  | 

### Return type

[**GetNetworkSwitchLinkAggregations200ResponseInner**](GetNetworkSwitchLinkAggregations200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkSwitchPortSchedule

> GetNetworkSwitchPortSchedules200ResponseInner CreateNetworkSwitchPortSchedule(ctx, networkId).CreateNetworkSwitchPortScheduleRequest(createNetworkSwitchPortScheduleRequest).Execute()

Add a switch port schedule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	createNetworkSwitchPortScheduleRequest := *openapiclient.NewCreateNetworkSwitchPortScheduleRequest("Name_example") // CreateNetworkSwitchPortScheduleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.CreateNetworkSwitchPortSchedule(context.Background(), networkId).CreateNetworkSwitchPortScheduleRequest(createNetworkSwitchPortScheduleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.CreateNetworkSwitchPortSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSwitchPortSchedule`: GetNetworkSwitchPortSchedules200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.CreateNetworkSwitchPortSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchPortScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchPortScheduleRequest** | [**CreateNetworkSwitchPortScheduleRequest**](CreateNetworkSwitchPortScheduleRequest.md) |  | 

### Return type

[**GetNetworkSwitchPortSchedules200ResponseInner**](GetNetworkSwitchPortSchedules200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkSwitchQosRule

> GetNetworkSwitchQosRules200ResponseInner CreateNetworkSwitchQosRule(ctx, networkId).CreateNetworkSwitchQosRuleRequest(createNetworkSwitchQosRuleRequest).Execute()

Add a quality of service rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	createNetworkSwitchQosRuleRequest := *openapiclient.NewCreateNetworkSwitchQosRuleRequest(int32(123)) // CreateNetworkSwitchQosRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.CreateNetworkSwitchQosRule(context.Background(), networkId).CreateNetworkSwitchQosRuleRequest(createNetworkSwitchQosRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.CreateNetworkSwitchQosRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSwitchQosRule`: GetNetworkSwitchQosRules200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.CreateNetworkSwitchQosRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchQosRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchQosRuleRequest** | [**CreateNetworkSwitchQosRuleRequest**](CreateNetworkSwitchQosRuleRequest.md) |  | 

### Return type

[**GetNetworkSwitchQosRules200ResponseInner**](GetNetworkSwitchQosRules200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkSwitchRoutingMulticastRendezvousPoint

> GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner CreateNetworkSwitchRoutingMulticastRendezvousPoint(ctx, networkId).CreateNetworkSwitchRoutingMulticastRendezvousPointRequest(createNetworkSwitchRoutingMulticastRendezvousPointRequest).Execute()

Create a multicast rendezvous point



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	createNetworkSwitchRoutingMulticastRendezvousPointRequest := *openapiclient.NewCreateNetworkSwitchRoutingMulticastRendezvousPointRequest("InterfaceIp_example", "MulticastGroup_example") // CreateNetworkSwitchRoutingMulticastRendezvousPointRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.CreateNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId).CreateNetworkSwitchRoutingMulticastRendezvousPointRequest(createNetworkSwitchRoutingMulticastRendezvousPointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.CreateNetworkSwitchRoutingMulticastRendezvousPoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSwitchRoutingMulticastRendezvousPoint`: GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.CreateNetworkSwitchRoutingMulticastRendezvousPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchRoutingMulticastRendezvousPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchRoutingMulticastRendezvousPointRequest** | [**CreateNetworkSwitchRoutingMulticastRendezvousPointRequest**](CreateNetworkSwitchRoutingMulticastRendezvousPointRequest.md) |  | 

### Return type

[**GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner**](GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkSwitchStack

> GetNetworkSwitchStacks200ResponseInner CreateNetworkSwitchStack(ctx, networkId).CreateNetworkSwitchStackRequest(createNetworkSwitchStackRequest).Execute()

Create a switch stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	createNetworkSwitchStackRequest := *openapiclient.NewCreateNetworkSwitchStackRequest("Name_example", []string{"Serials_example"}) // CreateNetworkSwitchStackRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.CreateNetworkSwitchStack(context.Background(), networkId).CreateNetworkSwitchStackRequest(createNetworkSwitchStackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.CreateNetworkSwitchStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSwitchStack`: GetNetworkSwitchStacks200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.CreateNetworkSwitchStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchStackRequest** | [**CreateNetworkSwitchStackRequest**](CreateNetworkSwitchStackRequest.md) |  | 

### Return type

[**GetNetworkSwitchStacks200ResponseInner**](GetNetworkSwitchStacks200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkSwitchStackRoutingInterface

> GetDeviceSwitchRoutingInterfaces200ResponseInner CreateNetworkSwitchStackRoutingInterface(ctx, networkId, switchStackId).CreateNetworkSwitchStackRoutingInterfaceRequest(createNetworkSwitchStackRoutingInterfaceRequest).Execute()

Create a layer 3 interface for a switch stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	switchStackId := "switchStackId_example" // string | Switch stack ID
	createNetworkSwitchStackRoutingInterfaceRequest := *openapiclient.NewCreateNetworkSwitchStackRoutingInterfaceRequest("Name_example", int32(123)) // CreateNetworkSwitchStackRoutingInterfaceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.CreateNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId).CreateNetworkSwitchStackRoutingInterfaceRequest(createNetworkSwitchStackRoutingInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.CreateNetworkSwitchStackRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSwitchStackRoutingInterface`: GetDeviceSwitchRoutingInterfaces200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.CreateNetworkSwitchStackRoutingInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchStackRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createNetworkSwitchStackRoutingInterfaceRequest** | [**CreateNetworkSwitchStackRoutingInterfaceRequest**](CreateNetworkSwitchStackRoutingInterfaceRequest.md) |  | 

### Return type

[**GetDeviceSwitchRoutingInterfaces200ResponseInner**](GetDeviceSwitchRoutingInterfaces200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkSwitchStackRoutingStaticRoute

> GetDeviceSwitchRoutingStaticRoutes200ResponseInner CreateNetworkSwitchStackRoutingStaticRoute(ctx, networkId, switchStackId).CreateDeviceSwitchRoutingStaticRouteRequest(createDeviceSwitchRoutingStaticRouteRequest).Execute()

Create a layer 3 static route for a switch stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	switchStackId := "switchStackId_example" // string | Switch stack ID
	createDeviceSwitchRoutingStaticRouteRequest := *openapiclient.NewCreateDeviceSwitchRoutingStaticRouteRequest("Subnet_example", "NextHopIp_example") // CreateDeviceSwitchRoutingStaticRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.CreateNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId).CreateDeviceSwitchRoutingStaticRouteRequest(createDeviceSwitchRoutingStaticRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.CreateNetworkSwitchStackRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSwitchStackRoutingStaticRoute`: GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.CreateNetworkSwitchStackRoutingStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchStackRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createDeviceSwitchRoutingStaticRouteRequest** | [**CreateDeviceSwitchRoutingStaticRouteRequest**](CreateDeviceSwitchRoutingStaticRouteRequest.md) |  | 

### Return type

[**GetDeviceSwitchRoutingStaticRoutes200ResponseInner**](GetDeviceSwitchRoutingStaticRoutes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CycleDeviceSwitchPorts

> CycleDeviceSwitchPorts200Response CycleDeviceSwitchPorts(ctx, serial).CycleDeviceSwitchPortsRequest(cycleDeviceSwitchPortsRequest).Execute()

Cycle a set of switch ports



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial
	cycleDeviceSwitchPortsRequest := *openapiclient.NewCycleDeviceSwitchPortsRequest([]string{"Ports_example"}) // CycleDeviceSwitchPortsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.CycleDeviceSwitchPorts(context.Background(), serial).CycleDeviceSwitchPortsRequest(cycleDeviceSwitchPortsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.CycleDeviceSwitchPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CycleDeviceSwitchPorts`: CycleDeviceSwitchPorts200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.CycleDeviceSwitchPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCycleDeviceSwitchPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cycleDeviceSwitchPortsRequest** | [**CycleDeviceSwitchPortsRequest**](CycleDeviceSwitchPortsRequest.md) |  | 

### Return type

[**CycleDeviceSwitchPorts200Response**](CycleDeviceSwitchPorts200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeviceSwitchRoutingInterface

> DeleteDeviceSwitchRoutingInterface(ctx, serial, interfaceId).Execute()

Delete a layer 3 interface from the switch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial
	interfaceId := "interfaceId_example" // string | Interface ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SwitchAPI.DeleteDeviceSwitchRoutingInterface(context.Background(), serial, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.DeleteDeviceSwitchRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceSwitchRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeviceSwitchRoutingStaticRoute

> DeleteDeviceSwitchRoutingStaticRoute(ctx, serial, staticRouteId).Execute()

Delete a layer 3 static route for a switch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial
	staticRouteId := "staticRouteId_example" // string | Static route ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SwitchAPI.DeleteDeviceSwitchRoutingStaticRoute(context.Background(), serial, staticRouteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.DeleteDeviceSwitchRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceSwitchRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSwitchAccessPolicy

> DeleteNetworkSwitchAccessPolicy(ctx, networkId, accessPolicyNumber).Execute()

Delete an access policy for a switch network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	accessPolicyNumber := "accessPolicyNumber_example" // string | Access policy number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SwitchAPI.DeleteNetworkSwitchAccessPolicy(context.Background(), networkId, accessPolicyNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.DeleteNetworkSwitchAccessPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**accessPolicyNumber** | **string** | Access policy number | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer

> DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx, networkId, trustedServerId).Execute()

Remove a server from being trusted by Dynamic ARP Inspection on this network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	trustedServerId := "trustedServerId_example" // string | Trusted server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SwitchAPI.DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(context.Background(), networkId, trustedServerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**trustedServerId** | **string** | Trusted server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSwitchLinkAggregation

> DeleteNetworkSwitchLinkAggregation(ctx, networkId, linkAggregationId).Execute()

Split a link aggregation group into separate ports



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	linkAggregationId := "linkAggregationId_example" // string | Link aggregation ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SwitchAPI.DeleteNetworkSwitchLinkAggregation(context.Background(), networkId, linkAggregationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.DeleteNetworkSwitchLinkAggregation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**linkAggregationId** | **string** | Link aggregation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchLinkAggregationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSwitchPortSchedule

> DeleteNetworkSwitchPortSchedule(ctx, networkId, portScheduleId).Execute()

Delete a switch port schedule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	portScheduleId := "portScheduleId_example" // string | Port schedule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SwitchAPI.DeleteNetworkSwitchPortSchedule(context.Background(), networkId, portScheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.DeleteNetworkSwitchPortSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**portScheduleId** | **string** | Port schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchPortScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSwitchQosRule

> DeleteNetworkSwitchQosRule(ctx, networkId, qosRuleId).Execute()

Delete a quality of service rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	qosRuleId := "qosRuleId_example" // string | Qos rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SwitchAPI.DeleteNetworkSwitchQosRule(context.Background(), networkId, qosRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.DeleteNetworkSwitchQosRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**qosRuleId** | **string** | Qos rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchQosRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSwitchRoutingMulticastRendezvousPoint

> DeleteNetworkSwitchRoutingMulticastRendezvousPoint(ctx, networkId, rendezvousPointId).Execute()

Delete a multicast rendezvous point



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	rendezvousPointId := "rendezvousPointId_example" // string | Rendezvous point ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SwitchAPI.DeleteNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.DeleteNetworkSwitchRoutingMulticastRendezvousPoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rendezvousPointId** | **string** | Rendezvous point ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSwitchStack

> DeleteNetworkSwitchStack(ctx, networkId, switchStackId).Execute()

Delete a stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	switchStackId := "switchStackId_example" // string | Switch stack ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SwitchAPI.DeleteNetworkSwitchStack(context.Background(), networkId, switchStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.DeleteNetworkSwitchStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSwitchStackRoutingInterface

> DeleteNetworkSwitchStackRoutingInterface(ctx, networkId, switchStackId, interfaceId).Execute()

Delete a layer 3 interface from a switch stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	switchStackId := "switchStackId_example" // string | Switch stack ID
	interfaceId := "interfaceId_example" // string | Interface ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SwitchAPI.DeleteNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.DeleteNetworkSwitchStackRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchStackRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSwitchStackRoutingStaticRoute

> DeleteNetworkSwitchStackRoutingStaticRoute(ctx, networkId, switchStackId, staticRouteId).Execute()

Delete a layer 3 static route for a switch stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	switchStackId := "switchStackId_example" // string | Switch stack ID
	staticRouteId := "staticRouteId_example" // string | Static route ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SwitchAPI.DeleteNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.DeleteNetworkSwitchStackRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchStackRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSwitchPort

> GetDeviceSwitchPorts200ResponseInner GetDeviceSwitchPort(ctx, serial, portId).Execute()

Return a switch port



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial
	portId := "portId_example" // string | Port ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetDeviceSwitchPort(context.Background(), serial, portId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetDeviceSwitchPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchPort`: GetDeviceSwitchPorts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetDeviceSwitchPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDeviceSwitchPorts200ResponseInner**](GetDeviceSwitchPorts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSwitchPorts

> []GetDeviceSwitchPorts200ResponseInner GetDeviceSwitchPorts(ctx, serial).Execute()

List the switch ports for a switch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetDeviceSwitchPorts(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetDeviceSwitchPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchPorts`: []GetDeviceSwitchPorts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetDeviceSwitchPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetDeviceSwitchPorts200ResponseInner**](GetDeviceSwitchPorts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSwitchPortsStatuses

> []GetDeviceSwitchPortsStatuses200ResponseInner GetDeviceSwitchPortsStatuses(ctx, serial).T0(t0).Timespan(timespan).Execute()

Return the status for all the ports of a switch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetDeviceSwitchPortsStatuses(context.Background(), serial).T0(t0).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetDeviceSwitchPortsStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchPortsStatuses`: []GetDeviceSwitchPortsStatuses200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetDeviceSwitchPortsStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchPortsStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

[**[]GetDeviceSwitchPortsStatuses200ResponseInner**](GetDeviceSwitchPortsStatuses200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSwitchPortsStatusesPackets

> []GetDeviceSwitchPortsStatusesPackets200ResponseInner GetDeviceSwitchPortsStatusesPackets(ctx, serial).T0(t0).Timespan(timespan).Execute()

Return the packet counters for all the ports of a switch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 1 day from today. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 1 day. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetDeviceSwitchPortsStatusesPackets(context.Background(), serial).T0(t0).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetDeviceSwitchPortsStatusesPackets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchPortsStatusesPackets`: []GetDeviceSwitchPortsStatusesPackets200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetDeviceSwitchPortsStatusesPackets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchPortsStatusesPacketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 1 day from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 1 day. The default is 1 day. | 

### Return type

[**[]GetDeviceSwitchPortsStatusesPackets200ResponseInner**](GetDeviceSwitchPortsStatusesPackets200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSwitchRoutingInterface

> GetDeviceSwitchRoutingInterfaces200ResponseInner GetDeviceSwitchRoutingInterface(ctx, serial, interfaceId).Execute()

Return a layer 3 interface for a switch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial
	interfaceId := "interfaceId_example" // string | Interface ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetDeviceSwitchRoutingInterface(context.Background(), serial, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetDeviceSwitchRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchRoutingInterface`: GetDeviceSwitchRoutingInterfaces200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetDeviceSwitchRoutingInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDeviceSwitchRoutingInterfaces200ResponseInner**](GetDeviceSwitchRoutingInterfaces200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSwitchRoutingInterfaceDhcp

> GetDeviceSwitchRoutingInterfaceDhcp200Response GetDeviceSwitchRoutingInterfaceDhcp(ctx, serial, interfaceId).Execute()

Return a layer 3 interface DHCP configuration for a switch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial
	interfaceId := "interfaceId_example" // string | Interface ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetDeviceSwitchRoutingInterfaceDhcp(context.Background(), serial, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetDeviceSwitchRoutingInterfaceDhcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchRoutingInterfaceDhcp`: GetDeviceSwitchRoutingInterfaceDhcp200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetDeviceSwitchRoutingInterfaceDhcp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchRoutingInterfaceDhcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDeviceSwitchRoutingInterfaceDhcp200Response**](GetDeviceSwitchRoutingInterfaceDhcp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSwitchRoutingInterfaces

> []GetDeviceSwitchRoutingInterfaces200ResponseInner GetDeviceSwitchRoutingInterfaces(ctx, serial).Execute()

List layer 3 interfaces for a switch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetDeviceSwitchRoutingInterfaces(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetDeviceSwitchRoutingInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchRoutingInterfaces`: []GetDeviceSwitchRoutingInterfaces200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetDeviceSwitchRoutingInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchRoutingInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetDeviceSwitchRoutingInterfaces200ResponseInner**](GetDeviceSwitchRoutingInterfaces200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSwitchRoutingStaticRoute

> GetDeviceSwitchRoutingStaticRoutes200ResponseInner GetDeviceSwitchRoutingStaticRoute(ctx, serial, staticRouteId).Execute()

Return a layer 3 static route for a switch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial
	staticRouteId := "staticRouteId_example" // string | Static route ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetDeviceSwitchRoutingStaticRoute(context.Background(), serial, staticRouteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetDeviceSwitchRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchRoutingStaticRoute`: GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetDeviceSwitchRoutingStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDeviceSwitchRoutingStaticRoutes200ResponseInner**](GetDeviceSwitchRoutingStaticRoutes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSwitchRoutingStaticRoutes

> []GetDeviceSwitchRoutingStaticRoutes200ResponseInner GetDeviceSwitchRoutingStaticRoutes(ctx, serial).Execute()

List layer 3 static routes for a switch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetDeviceSwitchRoutingStaticRoutes(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetDeviceSwitchRoutingStaticRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchRoutingStaticRoutes`: []GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetDeviceSwitchRoutingStaticRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchRoutingStaticRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetDeviceSwitchRoutingStaticRoutes200ResponseInner**](GetDeviceSwitchRoutingStaticRoutes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSwitchWarmSpare

> GetDeviceSwitchWarmSpare200Response GetDeviceSwitchWarmSpare(ctx, serial).Execute()

Return warm spare configuration for a switch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetDeviceSwitchWarmSpare(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetDeviceSwitchWarmSpare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchWarmSpare`: GetDeviceSwitchWarmSpare200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetDeviceSwitchWarmSpare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchWarmSpareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceSwitchWarmSpare200Response**](GetDeviceSwitchWarmSpare200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchAccessControlLists

> GetNetworkSwitchAccessControlLists200Response GetNetworkSwitchAccessControlLists(ctx, networkId).Execute()

Return the access control lists for a MS network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchAccessControlLists(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchAccessControlLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchAccessControlLists`: GetNetworkSwitchAccessControlLists200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchAccessControlLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchAccessControlListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchAccessControlLists200Response**](GetNetworkSwitchAccessControlLists200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchAccessPolicies

> []GetNetworkSwitchAccessPolicies200ResponseInner GetNetworkSwitchAccessPolicies(ctx, networkId).Execute()

List the access policies for a switch network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchAccessPolicies(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchAccessPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchAccessPolicies`: []GetNetworkSwitchAccessPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchAccessPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchAccessPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkSwitchAccessPolicies200ResponseInner**](GetNetworkSwitchAccessPolicies200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchAccessPolicy

> GetNetworkSwitchAccessPolicies200ResponseInner GetNetworkSwitchAccessPolicy(ctx, networkId, accessPolicyNumber).Execute()

Return a specific access policy for a switch network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	accessPolicyNumber := "accessPolicyNumber_example" // string | Access policy number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchAccessPolicy(context.Background(), networkId, accessPolicyNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchAccessPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchAccessPolicy`: GetNetworkSwitchAccessPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchAccessPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**accessPolicyNumber** | **string** | Access policy number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkSwitchAccessPolicies200ResponseInner**](GetNetworkSwitchAccessPolicies200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchAlternateManagementInterface

> GetNetworkSwitchAlternateManagementInterface200Response GetNetworkSwitchAlternateManagementInterface(ctx, networkId).Execute()

Return the switch alternate management interface for the network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchAlternateManagementInterface(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchAlternateManagementInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchAlternateManagementInterface`: GetNetworkSwitchAlternateManagementInterface200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchAlternateManagementInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchAlternateManagementInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchAlternateManagementInterface200Response**](GetNetworkSwitchAlternateManagementInterface200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchDhcpServerPolicy

> GetNetworkSwitchDhcpServerPolicy200Response GetNetworkSwitchDhcpServerPolicy(ctx, networkId).Execute()

Return the DHCP server settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchDhcpServerPolicy(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchDhcpServerPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchDhcpServerPolicy`: GetNetworkSwitchDhcpServerPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchDhcpServerPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchDhcpServerPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchDhcpServerPolicy200Response**](GetNetworkSwitchDhcpServerPolicy200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers

> []GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers(ctx, networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return the list of servers trusted by Dynamic ARP Inspection on this network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers(context.Background(), networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers`: []GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner**](GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice

> []GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice(ctx, networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return the devices that have a Dynamic ARP Inspection warning and their warnings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice(context.Background(), networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice`: []GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner**](GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchDhcpV4ServersSeen

> []GetNetworkSwitchDhcpV4ServersSeen200ResponseInner GetNetworkSwitchDhcpV4ServersSeen(ctx, networkId).T0(t0).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return the network's DHCPv4 servers seen within the selected timeframe (default 1 day)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchDhcpV4ServersSeen(context.Background(), networkId).T0(t0).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchDhcpV4ServersSeen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchDhcpV4ServersSeen`: []GetNetworkSwitchDhcpV4ServersSeen200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchDhcpV4ServersSeen`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchDhcpV4ServersSeenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkSwitchDhcpV4ServersSeen200ResponseInner**](GetNetworkSwitchDhcpV4ServersSeen200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchDscpToCosMappings

> GetNetworkSwitchDscpToCosMappings200Response GetNetworkSwitchDscpToCosMappings(ctx, networkId).Execute()

Return the DSCP to CoS mappings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchDscpToCosMappings(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchDscpToCosMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchDscpToCosMappings`: GetNetworkSwitchDscpToCosMappings200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchDscpToCosMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchDscpToCosMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchDscpToCosMappings200Response**](GetNetworkSwitchDscpToCosMappings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchLinkAggregations

> []GetNetworkSwitchLinkAggregations200ResponseInner GetNetworkSwitchLinkAggregations(ctx, networkId).Execute()

List link aggregation groups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchLinkAggregations(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchLinkAggregations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchLinkAggregations`: []GetNetworkSwitchLinkAggregations200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchLinkAggregations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchLinkAggregationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkSwitchLinkAggregations200ResponseInner**](GetNetworkSwitchLinkAggregations200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchMtu

> GetNetworkSwitchMtu200Response GetNetworkSwitchMtu(ctx, networkId).Execute()

Return the MTU configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchMtu(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchMtu``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchMtu`: GetNetworkSwitchMtu200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchMtu`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchMtuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchMtu200Response**](GetNetworkSwitchMtu200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchPortSchedules

> []GetNetworkSwitchPortSchedules200ResponseInner GetNetworkSwitchPortSchedules(ctx, networkId).Execute()

List switch port schedules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchPortSchedules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchPortSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchPortSchedules`: []GetNetworkSwitchPortSchedules200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchPortSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchPortSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkSwitchPortSchedules200ResponseInner**](GetNetworkSwitchPortSchedules200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchQosRule

> GetNetworkSwitchQosRules200ResponseInner GetNetworkSwitchQosRule(ctx, networkId, qosRuleId).Execute()

Return a quality of service rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	qosRuleId := "qosRuleId_example" // string | Qos rule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchQosRule(context.Background(), networkId, qosRuleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchQosRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchQosRule`: GetNetworkSwitchQosRules200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchQosRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**qosRuleId** | **string** | Qos rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchQosRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkSwitchQosRules200ResponseInner**](GetNetworkSwitchQosRules200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchQosRules

> []GetNetworkSwitchQosRules200ResponseInner GetNetworkSwitchQosRules(ctx, networkId).Execute()

List quality of service rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchQosRules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchQosRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchQosRules`: []GetNetworkSwitchQosRules200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchQosRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchQosRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkSwitchQosRules200ResponseInner**](GetNetworkSwitchQosRules200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchQosRulesOrder

> GetNetworkSwitchQosRulesOrder200Response GetNetworkSwitchQosRulesOrder(ctx, networkId).Execute()

Return the quality of service rule IDs by order in which they will be processed by the switch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchQosRulesOrder(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchQosRulesOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchQosRulesOrder`: GetNetworkSwitchQosRulesOrder200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchQosRulesOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchQosRulesOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchQosRulesOrder200Response**](GetNetworkSwitchQosRulesOrder200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchRoutingMulticast

> GetNetworkSwitchRoutingMulticast200Response GetNetworkSwitchRoutingMulticast(ctx, networkId).Execute()

Return multicast settings for a network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchRoutingMulticast(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchRoutingMulticast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchRoutingMulticast`: GetNetworkSwitchRoutingMulticast200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchRoutingMulticast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchRoutingMulticastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchRoutingMulticast200Response**](GetNetworkSwitchRoutingMulticast200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchRoutingMulticastRendezvousPoint

> GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner GetNetworkSwitchRoutingMulticastRendezvousPoint(ctx, networkId, rendezvousPointId).Execute()

Return a multicast rendezvous point



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	rendezvousPointId := "rendezvousPointId_example" // string | Rendezvous point ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchRoutingMulticastRendezvousPoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchRoutingMulticastRendezvousPoint`: GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchRoutingMulticastRendezvousPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rendezvousPointId** | **string** | Rendezvous point ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchRoutingMulticastRendezvousPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner**](GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchRoutingMulticastRendezvousPoints

> []GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner GetNetworkSwitchRoutingMulticastRendezvousPoints(ctx, networkId).Execute()

List multicast rendezvous points



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchRoutingMulticastRendezvousPoints(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchRoutingMulticastRendezvousPoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchRoutingMulticastRendezvousPoints`: []GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchRoutingMulticastRendezvousPoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchRoutingMulticastRendezvousPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner**](GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchRoutingOspf

> GetNetworkSwitchRoutingOspf200Response GetNetworkSwitchRoutingOspf(ctx, networkId).Execute()

Return layer 3 OSPF routing configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchRoutingOspf(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchRoutingOspf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchRoutingOspf`: GetNetworkSwitchRoutingOspf200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchRoutingOspf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchRoutingOspfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchRoutingOspf200Response**](GetNetworkSwitchRoutingOspf200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchSettings

> GetNetworkSwitchSettings200Response GetNetworkSwitchSettings(ctx, networkId).Execute()

Returns the switch network settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchSettings(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchSettings`: GetNetworkSwitchSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchSettings200Response**](GetNetworkSwitchSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchStack

> GetNetworkSwitchStacks200ResponseInner GetNetworkSwitchStack(ctx, networkId, switchStackId).Execute()

Show a switch stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	switchStackId := "switchStackId_example" // string | Switch stack ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchStack(context.Background(), networkId, switchStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStack`: GetNetworkSwitchStacks200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkSwitchStacks200ResponseInner**](GetNetworkSwitchStacks200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchStackRoutingInterface

> GetDeviceSwitchRoutingInterfaces200ResponseInner GetNetworkSwitchStackRoutingInterface(ctx, networkId, switchStackId, interfaceId).Execute()

Return a layer 3 interface from a switch stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	switchStackId := "switchStackId_example" // string | Switch stack ID
	interfaceId := "interfaceId_example" // string | Interface ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchStackRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStackRoutingInterface`: GetDeviceSwitchRoutingInterfaces200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchStackRoutingInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetDeviceSwitchRoutingInterfaces200ResponseInner**](GetDeviceSwitchRoutingInterfaces200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchStackRoutingInterfaceDhcp

> GetDeviceSwitchRoutingInterfaceDhcp200Response GetNetworkSwitchStackRoutingInterfaceDhcp(ctx, networkId, switchStackId, interfaceId).Execute()

Return a layer 3 interface DHCP configuration for a switch stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	switchStackId := "switchStackId_example" // string | Switch stack ID
	interfaceId := "interfaceId_example" // string | Interface ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchStackRoutingInterfaceDhcp(context.Background(), networkId, switchStackId, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchStackRoutingInterfaceDhcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStackRoutingInterfaceDhcp`: GetDeviceSwitchRoutingInterfaceDhcp200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchStackRoutingInterfaceDhcp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRoutingInterfaceDhcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetDeviceSwitchRoutingInterfaceDhcp200Response**](GetDeviceSwitchRoutingInterfaceDhcp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchStackRoutingInterfaces

> []GetDeviceSwitchRoutingInterfaces200ResponseInner GetNetworkSwitchStackRoutingInterfaces(ctx, networkId, switchStackId).Execute()

List layer 3 interfaces for a switch stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	switchStackId := "switchStackId_example" // string | Switch stack ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchStackRoutingInterfaces(context.Background(), networkId, switchStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchStackRoutingInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStackRoutingInterfaces`: []GetDeviceSwitchRoutingInterfaces200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchStackRoutingInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRoutingInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetDeviceSwitchRoutingInterfaces200ResponseInner**](GetDeviceSwitchRoutingInterfaces200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchStackRoutingStaticRoute

> GetDeviceSwitchRoutingStaticRoutes200ResponseInner GetNetworkSwitchStackRoutingStaticRoute(ctx, networkId, switchStackId, staticRouteId).Execute()

Return a layer 3 static route for a switch stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	switchStackId := "switchStackId_example" // string | Switch stack ID
	staticRouteId := "staticRouteId_example" // string | Static route ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchStackRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStackRoutingStaticRoute`: GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchStackRoutingStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetDeviceSwitchRoutingStaticRoutes200ResponseInner**](GetDeviceSwitchRoutingStaticRoutes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchStackRoutingStaticRoutes

> []GetDeviceSwitchRoutingStaticRoutes200ResponseInner GetNetworkSwitchStackRoutingStaticRoutes(ctx, networkId, switchStackId).Execute()

List layer 3 static routes for a switch stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	switchStackId := "switchStackId_example" // string | Switch stack ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchStackRoutingStaticRoutes(context.Background(), networkId, switchStackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchStackRoutingStaticRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStackRoutingStaticRoutes`: []GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchStackRoutingStaticRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStackRoutingStaticRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetDeviceSwitchRoutingStaticRoutes200ResponseInner**](GetDeviceSwitchRoutingStaticRoutes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchStacks

> []GetNetworkSwitchStacks200ResponseInner GetNetworkSwitchStacks(ctx, networkId).Execute()

List the switch stacks in a network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchStacks(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchStacks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStacks`: []GetNetworkSwitchStacks200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchStacks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStacksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkSwitchStacks200ResponseInner**](GetNetworkSwitchStacks200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchStormControl

> GetNetworkSwitchStormControl200Response GetNetworkSwitchStormControl(ctx, networkId).Execute()

Return the storm control configuration for a switch network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchStormControl(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchStormControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStormControl`: GetNetworkSwitchStormControl200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchStormControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStormControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchStormControl200Response**](GetNetworkSwitchStormControl200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchStp

> GetNetworkSwitchStp200Response GetNetworkSwitchStp(ctx, networkId).Execute()

Returns STP settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetNetworkSwitchStp(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetNetworkSwitchStp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStp`: GetNetworkSwitchStp200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetNetworkSwitchStp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchStp200Response**](GetNetworkSwitchStp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConfigTemplateSwitchProfilePort

> GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner GetOrganizationConfigTemplateSwitchProfilePort(ctx, organizationId, configTemplateId, profileId, portId).Execute()

Return a switch template port



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	configTemplateId := "configTemplateId_example" // string | Config template ID
	profileId := "profileId_example" // string | Profile ID
	portId := "portId_example" // string | Port ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetOrganizationConfigTemplateSwitchProfilePort(context.Background(), organizationId, configTemplateId, profileId, portId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetOrganizationConfigTemplateSwitchProfilePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationConfigTemplateSwitchProfilePort`: GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetOrganizationConfigTemplateSwitchProfilePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 
**profileId** | **string** | Profile ID | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConfigTemplateSwitchProfilePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner**](GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConfigTemplateSwitchProfilePorts

> []GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner GetOrganizationConfigTemplateSwitchProfilePorts(ctx, organizationId, configTemplateId, profileId).Execute()

Return all the ports of a switch template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	configTemplateId := "configTemplateId_example" // string | Config template ID
	profileId := "profileId_example" // string | Profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetOrganizationConfigTemplateSwitchProfilePorts(context.Background(), organizationId, configTemplateId, profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetOrganizationConfigTemplateSwitchProfilePorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationConfigTemplateSwitchProfilePorts`: []GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetOrganizationConfigTemplateSwitchProfilePorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConfigTemplateSwitchProfilePortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner**](GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConfigTemplateSwitchProfiles

> []GetOrganizationConfigTemplateSwitchProfiles200ResponseInner GetOrganizationConfigTemplateSwitchProfiles(ctx, organizationId, configTemplateId).Execute()

List the switch templates for your switch template configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	configTemplateId := "configTemplateId_example" // string | Config template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetOrganizationConfigTemplateSwitchProfiles(context.Background(), organizationId, configTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetOrganizationConfigTemplateSwitchProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationConfigTemplateSwitchProfiles`: []GetOrganizationConfigTemplateSwitchProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetOrganizationConfigTemplateSwitchProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConfigTemplateSwitchProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetOrganizationConfigTemplateSwitchProfiles200ResponseInner**](GetOrganizationConfigTemplateSwitchProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummarySwitchPowerHistory

> []GetOrganizationSummarySwitchPowerHistory200ResponseInner GetOrganizationSummarySwitchPowerHistory(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Returns the total PoE power draw for all switch ports in the organization over the requested timespan (by default the last 24 hours)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 186 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetOrganizationSummarySwitchPowerHistory(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetOrganizationSummarySwitchPowerHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSummarySwitchPowerHistory`: []GetOrganizationSummarySwitchPowerHistory200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetOrganizationSummarySwitchPowerHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummarySwitchPowerHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 186 days. The default is 1 day. | 

### Return type

[**[]GetOrganizationSummarySwitchPowerHistory200ResponseInner**](GetOrganizationSummarySwitchPowerHistory200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSwitchPortsBySwitch

> []GetOrganizationSwitchPortsBySwitch200ResponseInner GetOrganizationSwitchPortsBySwitch(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).ConfigurationUpdatedAfter(configurationUpdatedAfter).Mac(mac).Macs(macs).Name(name).NetworkIds(networkIds).PortProfileIds(portProfileIds).Serial(serial).Serials(serials).Execute()

List the switchports in an organization by switch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 50. Default is 50. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	configurationUpdatedAfter := time.Now() // time.Time | Optional parameter to filter items to switches where the configuration has been updated after the given timestamp. (optional)
	mac := "mac_example" // string | Optional parameter to filter items to switches with MAC addresses that contain the search term or are an exact match. (optional)
	macs := []string{"Inner_example"} // []string | Optional parameter to filter items to switches that have one of the provided MAC addresses. (optional)
	name := "name_example" // string | Optional parameter to filter items to switches with names that contain the search term or are an exact match. (optional)
	networkIds := []string{"Inner_example"} // []string | Optional parameter to filter items to switches in one of the provided networks. (optional)
	portProfileIds := []string{"Inner_example"} // []string | Optional parameter to filter items to switches that contain switchports belonging to one of the specified port profiles. (optional)
	serial := "serial_example" // string | Optional parameter to filter items to switches with serial number that contains the search term or are an exact match. (optional)
	serials := []string{"Inner_example"} // []string | Optional parameter to filter items to switches that have one of the provided serials. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetOrganizationSwitchPortsBySwitch(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).ConfigurationUpdatedAfter(configurationUpdatedAfter).Mac(mac).Macs(macs).Name(name).NetworkIds(networkIds).PortProfileIds(portProfileIds).Serial(serial).Serials(serials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetOrganizationSwitchPortsBySwitch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSwitchPortsBySwitch`: []GetOrganizationSwitchPortsBySwitch200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetOrganizationSwitchPortsBySwitch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSwitchPortsBySwitchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 50. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **configurationUpdatedAfter** | **time.Time** | Optional parameter to filter items to switches where the configuration has been updated after the given timestamp. | 
 **mac** | **string** | Optional parameter to filter items to switches with MAC addresses that contain the search term or are an exact match. | 
 **macs** | **[]string** | Optional parameter to filter items to switches that have one of the provided MAC addresses. | 
 **name** | **string** | Optional parameter to filter items to switches with names that contain the search term or are an exact match. | 
 **networkIds** | **[]string** | Optional parameter to filter items to switches in one of the provided networks. | 
 **portProfileIds** | **[]string** | Optional parameter to filter items to switches that contain switchports belonging to one of the specified port profiles. | 
 **serial** | **string** | Optional parameter to filter items to switches with serial number that contains the search term or are an exact match. | 
 **serials** | **[]string** | Optional parameter to filter items to switches that have one of the provided serials. | 

### Return type

[**[]GetOrganizationSwitchPortsBySwitch200ResponseInner**](GetOrganizationSwitchPortsBySwitch200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSwitchPortsOverview

> GetOrganizationSwitchPortsOverview200Response GetOrganizationSwitchPortsOverview(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Returns the counts of all active ports for the requested timespan, grouped by speed



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 12 hours and be less than or equal to 186 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.GetOrganizationSwitchPortsOverview(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.GetOrganizationSwitchPortsOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSwitchPortsOverview`: GetOrganizationSwitchPortsOverview200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.GetOrganizationSwitchPortsOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSwitchPortsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 12 hours and be less than or equal to 186 days. The default is 1 day. | 

### Return type

[**GetOrganizationSwitchPortsOverview200Response**](GetOrganizationSwitchPortsOverview200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveNetworkSwitchStack

> GetNetworkSwitchStacks200ResponseInner RemoveNetworkSwitchStack(ctx, networkId, switchStackId).RemoveNetworkSwitchStackRequest(removeNetworkSwitchStackRequest).Execute()

Remove a switch from a stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	switchStackId := "switchStackId_example" // string | Switch stack ID
	removeNetworkSwitchStackRequest := *openapiclient.NewRemoveNetworkSwitchStackRequest("Serial_example") // RemoveNetworkSwitchStackRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.RemoveNetworkSwitchStack(context.Background(), networkId, switchStackId).RemoveNetworkSwitchStackRequest(removeNetworkSwitchStackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.RemoveNetworkSwitchStack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveNetworkSwitchStack`: GetNetworkSwitchStacks200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.RemoveNetworkSwitchStack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNetworkSwitchStackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **removeNetworkSwitchStackRequest** | [**RemoveNetworkSwitchStackRequest**](RemoveNetworkSwitchStackRequest.md) |  | 

### Return type

[**GetNetworkSwitchStacks200ResponseInner**](GetNetworkSwitchStacks200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceSwitchPort

> GetDeviceSwitchPorts200ResponseInner UpdateDeviceSwitchPort(ctx, serial, portId).UpdateDeviceSwitchPortRequest(updateDeviceSwitchPortRequest).Execute()

Update a switch port



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial
	portId := "portId_example" // string | Port ID
	updateDeviceSwitchPortRequest := *openapiclient.NewUpdateDeviceSwitchPortRequest() // UpdateDeviceSwitchPortRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateDeviceSwitchPort(context.Background(), serial, portId).UpdateDeviceSwitchPortRequest(updateDeviceSwitchPortRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateDeviceSwitchPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceSwitchPort`: GetDeviceSwitchPorts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateDeviceSwitchPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceSwitchPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDeviceSwitchPortRequest** | [**UpdateDeviceSwitchPortRequest**](UpdateDeviceSwitchPortRequest.md) |  | 

### Return type

[**GetDeviceSwitchPorts200ResponseInner**](GetDeviceSwitchPorts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceSwitchRoutingInterface

> GetDeviceSwitchRoutingInterfaces200ResponseInner UpdateDeviceSwitchRoutingInterface(ctx, serial, interfaceId).CreateDeviceSwitchRoutingInterfaceRequest(createDeviceSwitchRoutingInterfaceRequest).Execute()

Update a layer 3 interface for a switch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial
	interfaceId := "interfaceId_example" // string | Interface ID
	createDeviceSwitchRoutingInterfaceRequest := *openapiclient.NewCreateDeviceSwitchRoutingInterfaceRequest() // CreateDeviceSwitchRoutingInterfaceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateDeviceSwitchRoutingInterface(context.Background(), serial, interfaceId).CreateDeviceSwitchRoutingInterfaceRequest(createDeviceSwitchRoutingInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateDeviceSwitchRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceSwitchRoutingInterface`: GetDeviceSwitchRoutingInterfaces200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateDeviceSwitchRoutingInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceSwitchRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createDeviceSwitchRoutingInterfaceRequest** | [**CreateDeviceSwitchRoutingInterfaceRequest**](CreateDeviceSwitchRoutingInterfaceRequest.md) |  | 

### Return type

[**GetDeviceSwitchRoutingInterfaces200ResponseInner**](GetDeviceSwitchRoutingInterfaces200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceSwitchRoutingInterfaceDhcp

> GetDeviceSwitchRoutingInterfaceDhcp200Response UpdateDeviceSwitchRoutingInterfaceDhcp(ctx, serial, interfaceId).UpdateDeviceSwitchRoutingInterfaceDhcpRequest(updateDeviceSwitchRoutingInterfaceDhcpRequest).Execute()

Update a layer 3 interface DHCP configuration for a switch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial
	interfaceId := "interfaceId_example" // string | Interface ID
	updateDeviceSwitchRoutingInterfaceDhcpRequest := *openapiclient.NewUpdateDeviceSwitchRoutingInterfaceDhcpRequest() // UpdateDeviceSwitchRoutingInterfaceDhcpRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateDeviceSwitchRoutingInterfaceDhcp(context.Background(), serial, interfaceId).UpdateDeviceSwitchRoutingInterfaceDhcpRequest(updateDeviceSwitchRoutingInterfaceDhcpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateDeviceSwitchRoutingInterfaceDhcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceSwitchRoutingInterfaceDhcp`: GetDeviceSwitchRoutingInterfaceDhcp200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateDeviceSwitchRoutingInterfaceDhcp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceSwitchRoutingInterfaceDhcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDeviceSwitchRoutingInterfaceDhcpRequest** | [**UpdateDeviceSwitchRoutingInterfaceDhcpRequest**](UpdateDeviceSwitchRoutingInterfaceDhcpRequest.md) |  | 

### Return type

[**GetDeviceSwitchRoutingInterfaceDhcp200Response**](GetDeviceSwitchRoutingInterfaceDhcp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceSwitchRoutingStaticRoute

> GetDeviceSwitchRoutingStaticRoutes200ResponseInner UpdateDeviceSwitchRoutingStaticRoute(ctx, serial, staticRouteId).UpdateDeviceSwitchRoutingStaticRouteRequest(updateDeviceSwitchRoutingStaticRouteRequest).Execute()

Update a layer 3 static route for a switch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial
	staticRouteId := "staticRouteId_example" // string | Static route ID
	updateDeviceSwitchRoutingStaticRouteRequest := *openapiclient.NewUpdateDeviceSwitchRoutingStaticRouteRequest() // UpdateDeviceSwitchRoutingStaticRouteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateDeviceSwitchRoutingStaticRoute(context.Background(), serial, staticRouteId).UpdateDeviceSwitchRoutingStaticRouteRequest(updateDeviceSwitchRoutingStaticRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateDeviceSwitchRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceSwitchRoutingStaticRoute`: GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateDeviceSwitchRoutingStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceSwitchRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDeviceSwitchRoutingStaticRouteRequest** | [**UpdateDeviceSwitchRoutingStaticRouteRequest**](UpdateDeviceSwitchRoutingStaticRouteRequest.md) |  | 

### Return type

[**GetDeviceSwitchRoutingStaticRoutes200ResponseInner**](GetDeviceSwitchRoutingStaticRoutes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceSwitchWarmSpare

> GetDeviceSwitchWarmSpare200Response UpdateDeviceSwitchWarmSpare(ctx, serial).UpdateDeviceSwitchWarmSpareRequest(updateDeviceSwitchWarmSpareRequest).Execute()

Update warm spare configuration for a switch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial
	updateDeviceSwitchWarmSpareRequest := *openapiclient.NewUpdateDeviceSwitchWarmSpareRequest(false) // UpdateDeviceSwitchWarmSpareRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateDeviceSwitchWarmSpare(context.Background(), serial).UpdateDeviceSwitchWarmSpareRequest(updateDeviceSwitchWarmSpareRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateDeviceSwitchWarmSpare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceSwitchWarmSpare`: GetDeviceSwitchWarmSpare200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateDeviceSwitchWarmSpare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceSwitchWarmSpareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceSwitchWarmSpareRequest** | [**UpdateDeviceSwitchWarmSpareRequest**](UpdateDeviceSwitchWarmSpareRequest.md) |  | 

### Return type

[**GetDeviceSwitchWarmSpare200Response**](GetDeviceSwitchWarmSpare200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchAccessControlLists

> GetNetworkSwitchAccessControlLists200Response UpdateNetworkSwitchAccessControlLists(ctx, networkId).UpdateNetworkSwitchAccessControlListsRequest(updateNetworkSwitchAccessControlListsRequest).Execute()

Update the access control lists for a MS network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkSwitchAccessControlListsRequest := *openapiclient.NewUpdateNetworkSwitchAccessControlListsRequest([]openapiclient.UpdateNetworkSwitchAccessControlListsRequestRulesInner{*openapiclient.NewUpdateNetworkSwitchAccessControlListsRequestRulesInner("Policy_example", "Protocol_example", "SrcCidr_example", "DstCidr_example")}) // UpdateNetworkSwitchAccessControlListsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateNetworkSwitchAccessControlLists(context.Background(), networkId).UpdateNetworkSwitchAccessControlListsRequest(updateNetworkSwitchAccessControlListsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateNetworkSwitchAccessControlLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchAccessControlLists`: GetNetworkSwitchAccessControlLists200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateNetworkSwitchAccessControlLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchAccessControlListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchAccessControlListsRequest** | [**UpdateNetworkSwitchAccessControlListsRequest**](UpdateNetworkSwitchAccessControlListsRequest.md) |  | 

### Return type

[**GetNetworkSwitchAccessControlLists200Response**](GetNetworkSwitchAccessControlLists200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchAccessPolicy

> GetNetworkSwitchAccessPolicies200ResponseInner UpdateNetworkSwitchAccessPolicy(ctx, networkId, accessPolicyNumber).UpdateNetworkSwitchAccessPolicyRequest(updateNetworkSwitchAccessPolicyRequest).Execute()

Update an access policy for a switch network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	accessPolicyNumber := "accessPolicyNumber_example" // string | Access policy number
	updateNetworkSwitchAccessPolicyRequest := *openapiclient.NewUpdateNetworkSwitchAccessPolicyRequest() // UpdateNetworkSwitchAccessPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateNetworkSwitchAccessPolicy(context.Background(), networkId, accessPolicyNumber).UpdateNetworkSwitchAccessPolicyRequest(updateNetworkSwitchAccessPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateNetworkSwitchAccessPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchAccessPolicy`: GetNetworkSwitchAccessPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateNetworkSwitchAccessPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**accessPolicyNumber** | **string** | Access policy number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSwitchAccessPolicyRequest** | [**UpdateNetworkSwitchAccessPolicyRequest**](UpdateNetworkSwitchAccessPolicyRequest.md) |  | 

### Return type

[**GetNetworkSwitchAccessPolicies200ResponseInner**](GetNetworkSwitchAccessPolicies200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchAlternateManagementInterface

> GetNetworkSwitchAlternateManagementInterface200Response UpdateNetworkSwitchAlternateManagementInterface(ctx, networkId).UpdateNetworkSwitchAlternateManagementInterfaceRequest(updateNetworkSwitchAlternateManagementInterfaceRequest).Execute()

Update the switch alternate management interface for the network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkSwitchAlternateManagementInterfaceRequest := *openapiclient.NewUpdateNetworkSwitchAlternateManagementInterfaceRequest() // UpdateNetworkSwitchAlternateManagementInterfaceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateNetworkSwitchAlternateManagementInterface(context.Background(), networkId).UpdateNetworkSwitchAlternateManagementInterfaceRequest(updateNetworkSwitchAlternateManagementInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateNetworkSwitchAlternateManagementInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchAlternateManagementInterface`: GetNetworkSwitchAlternateManagementInterface200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateNetworkSwitchAlternateManagementInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchAlternateManagementInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchAlternateManagementInterfaceRequest** | [**UpdateNetworkSwitchAlternateManagementInterfaceRequest**](UpdateNetworkSwitchAlternateManagementInterfaceRequest.md) |  | 

### Return type

[**GetNetworkSwitchAlternateManagementInterface200Response**](GetNetworkSwitchAlternateManagementInterface200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchDhcpServerPolicy

> GetNetworkSwitchDhcpServerPolicy200Response UpdateNetworkSwitchDhcpServerPolicy(ctx, networkId).UpdateNetworkSwitchDhcpServerPolicyRequest(updateNetworkSwitchDhcpServerPolicyRequest).Execute()

Update the DHCP server settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkSwitchDhcpServerPolicyRequest := *openapiclient.NewUpdateNetworkSwitchDhcpServerPolicyRequest() // UpdateNetworkSwitchDhcpServerPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateNetworkSwitchDhcpServerPolicy(context.Background(), networkId).UpdateNetworkSwitchDhcpServerPolicyRequest(updateNetworkSwitchDhcpServerPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateNetworkSwitchDhcpServerPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchDhcpServerPolicy`: GetNetworkSwitchDhcpServerPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateNetworkSwitchDhcpServerPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchDhcpServerPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchDhcpServerPolicyRequest** | [**UpdateNetworkSwitchDhcpServerPolicyRequest**](UpdateNetworkSwitchDhcpServerPolicyRequest.md) |  | 

### Return type

[**GetNetworkSwitchDhcpServerPolicy200Response**](GetNetworkSwitchDhcpServerPolicy200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer

> GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx, networkId, trustedServerId).UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest(updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest).Execute()

Update a server that is trusted by Dynamic ARP Inspection on this network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	trustedServerId := "trustedServerId_example" // string | Trusted server ID
	updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest := *openapiclient.NewUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest() // UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(context.Background(), networkId, trustedServerId).UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest(updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer`: GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**trustedServerId** | **string** | Trusted server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest** | [**UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest**](UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest.md) |  | 

### Return type

[**GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner**](GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchDscpToCosMappings

> GetNetworkSwitchDscpToCosMappings200Response UpdateNetworkSwitchDscpToCosMappings(ctx, networkId).UpdateNetworkSwitchDscpToCosMappingsRequest(updateNetworkSwitchDscpToCosMappingsRequest).Execute()

Update the DSCP to CoS mappings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkSwitchDscpToCosMappingsRequest := *openapiclient.NewUpdateNetworkSwitchDscpToCosMappingsRequest([]openapiclient.UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner{*openapiclient.NewUpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner(int32(123), int32(123))}) // UpdateNetworkSwitchDscpToCosMappingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateNetworkSwitchDscpToCosMappings(context.Background(), networkId).UpdateNetworkSwitchDscpToCosMappingsRequest(updateNetworkSwitchDscpToCosMappingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateNetworkSwitchDscpToCosMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchDscpToCosMappings`: GetNetworkSwitchDscpToCosMappings200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateNetworkSwitchDscpToCosMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchDscpToCosMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchDscpToCosMappingsRequest** | [**UpdateNetworkSwitchDscpToCosMappingsRequest**](UpdateNetworkSwitchDscpToCosMappingsRequest.md) |  | 

### Return type

[**GetNetworkSwitchDscpToCosMappings200Response**](GetNetworkSwitchDscpToCosMappings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchLinkAggregation

> GetNetworkSwitchLinkAggregations200ResponseInner UpdateNetworkSwitchLinkAggregation(ctx, networkId, linkAggregationId).UpdateNetworkSwitchLinkAggregationRequest(updateNetworkSwitchLinkAggregationRequest).Execute()

Update a link aggregation group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	linkAggregationId := "linkAggregationId_example" // string | Link aggregation ID
	updateNetworkSwitchLinkAggregationRequest := *openapiclient.NewUpdateNetworkSwitchLinkAggregationRequest() // UpdateNetworkSwitchLinkAggregationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateNetworkSwitchLinkAggregation(context.Background(), networkId, linkAggregationId).UpdateNetworkSwitchLinkAggregationRequest(updateNetworkSwitchLinkAggregationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateNetworkSwitchLinkAggregation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchLinkAggregation`: GetNetworkSwitchLinkAggregations200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateNetworkSwitchLinkAggregation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**linkAggregationId** | **string** | Link aggregation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchLinkAggregationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSwitchLinkAggregationRequest** | [**UpdateNetworkSwitchLinkAggregationRequest**](UpdateNetworkSwitchLinkAggregationRequest.md) |  | 

### Return type

[**GetNetworkSwitchLinkAggregations200ResponseInner**](GetNetworkSwitchLinkAggregations200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchMtu

> GetNetworkSwitchMtu200Response UpdateNetworkSwitchMtu(ctx, networkId).UpdateNetworkSwitchMtuRequest(updateNetworkSwitchMtuRequest).Execute()

Update the MTU configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkSwitchMtuRequest := *openapiclient.NewUpdateNetworkSwitchMtuRequest() // UpdateNetworkSwitchMtuRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateNetworkSwitchMtu(context.Background(), networkId).UpdateNetworkSwitchMtuRequest(updateNetworkSwitchMtuRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateNetworkSwitchMtu``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchMtu`: GetNetworkSwitchMtu200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateNetworkSwitchMtu`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchMtuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchMtuRequest** | [**UpdateNetworkSwitchMtuRequest**](UpdateNetworkSwitchMtuRequest.md) |  | 

### Return type

[**GetNetworkSwitchMtu200Response**](GetNetworkSwitchMtu200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchPortSchedule

> GetNetworkSwitchPortSchedules200ResponseInner UpdateNetworkSwitchPortSchedule(ctx, networkId, portScheduleId).UpdateNetworkSwitchPortScheduleRequest(updateNetworkSwitchPortScheduleRequest).Execute()

Update a switch port schedule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	portScheduleId := "portScheduleId_example" // string | Port schedule ID
	updateNetworkSwitchPortScheduleRequest := *openapiclient.NewUpdateNetworkSwitchPortScheduleRequest() // UpdateNetworkSwitchPortScheduleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateNetworkSwitchPortSchedule(context.Background(), networkId, portScheduleId).UpdateNetworkSwitchPortScheduleRequest(updateNetworkSwitchPortScheduleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateNetworkSwitchPortSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchPortSchedule`: GetNetworkSwitchPortSchedules200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateNetworkSwitchPortSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**portScheduleId** | **string** | Port schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchPortScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSwitchPortScheduleRequest** | [**UpdateNetworkSwitchPortScheduleRequest**](UpdateNetworkSwitchPortScheduleRequest.md) |  | 

### Return type

[**GetNetworkSwitchPortSchedules200ResponseInner**](GetNetworkSwitchPortSchedules200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchQosRule

> GetNetworkSwitchQosRules200ResponseInner UpdateNetworkSwitchQosRule(ctx, networkId, qosRuleId).UpdateNetworkSwitchQosRuleRequest(updateNetworkSwitchQosRuleRequest).Execute()

Update a quality of service rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	qosRuleId := "qosRuleId_example" // string | Qos rule ID
	updateNetworkSwitchQosRuleRequest := *openapiclient.NewUpdateNetworkSwitchQosRuleRequest() // UpdateNetworkSwitchQosRuleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateNetworkSwitchQosRule(context.Background(), networkId, qosRuleId).UpdateNetworkSwitchQosRuleRequest(updateNetworkSwitchQosRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateNetworkSwitchQosRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchQosRule`: GetNetworkSwitchQosRules200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateNetworkSwitchQosRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**qosRuleId** | **string** | Qos rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchQosRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSwitchQosRuleRequest** | [**UpdateNetworkSwitchQosRuleRequest**](UpdateNetworkSwitchQosRuleRequest.md) |  | 

### Return type

[**GetNetworkSwitchQosRules200ResponseInner**](GetNetworkSwitchQosRules200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchQosRulesOrder

> GetNetworkSwitchQosRulesOrder200Response UpdateNetworkSwitchQosRulesOrder(ctx, networkId).UpdateNetworkSwitchQosRulesOrderRequest(updateNetworkSwitchQosRulesOrderRequest).Execute()

Update the order in which the rules should be processed by the switch



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkSwitchQosRulesOrderRequest := *openapiclient.NewUpdateNetworkSwitchQosRulesOrderRequest([]string{"RuleIds_example"}) // UpdateNetworkSwitchQosRulesOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateNetworkSwitchQosRulesOrder(context.Background(), networkId).UpdateNetworkSwitchQosRulesOrderRequest(updateNetworkSwitchQosRulesOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateNetworkSwitchQosRulesOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchQosRulesOrder`: GetNetworkSwitchQosRulesOrder200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateNetworkSwitchQosRulesOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchQosRulesOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchQosRulesOrderRequest** | [**UpdateNetworkSwitchQosRulesOrderRequest**](UpdateNetworkSwitchQosRulesOrderRequest.md) |  | 

### Return type

[**GetNetworkSwitchQosRulesOrder200Response**](GetNetworkSwitchQosRulesOrder200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchRoutingMulticast

> GetNetworkSwitchRoutingMulticast200Response UpdateNetworkSwitchRoutingMulticast(ctx, networkId).UpdateNetworkSwitchRoutingMulticastRequest(updateNetworkSwitchRoutingMulticastRequest).Execute()

Update multicast settings for a network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkSwitchRoutingMulticastRequest := *openapiclient.NewUpdateNetworkSwitchRoutingMulticastRequest() // UpdateNetworkSwitchRoutingMulticastRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateNetworkSwitchRoutingMulticast(context.Background(), networkId).UpdateNetworkSwitchRoutingMulticastRequest(updateNetworkSwitchRoutingMulticastRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateNetworkSwitchRoutingMulticast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchRoutingMulticast`: GetNetworkSwitchRoutingMulticast200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateNetworkSwitchRoutingMulticast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchRoutingMulticastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchRoutingMulticastRequest** | [**UpdateNetworkSwitchRoutingMulticastRequest**](UpdateNetworkSwitchRoutingMulticastRequest.md) |  | 

### Return type

[**GetNetworkSwitchRoutingMulticast200Response**](GetNetworkSwitchRoutingMulticast200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchRoutingMulticastRendezvousPoint

> GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner UpdateNetworkSwitchRoutingMulticastRendezvousPoint(ctx, networkId, rendezvousPointId).CreateNetworkSwitchRoutingMulticastRendezvousPointRequest(createNetworkSwitchRoutingMulticastRendezvousPointRequest).Execute()

Update a multicast rendezvous point



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	rendezvousPointId := "rendezvousPointId_example" // string | Rendezvous point ID
	createNetworkSwitchRoutingMulticastRendezvousPointRequest := *openapiclient.NewCreateNetworkSwitchRoutingMulticastRendezvousPointRequest("InterfaceIp_example", "MulticastGroup_example") // CreateNetworkSwitchRoutingMulticastRendezvousPointRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).CreateNetworkSwitchRoutingMulticastRendezvousPointRequest(createNetworkSwitchRoutingMulticastRendezvousPointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateNetworkSwitchRoutingMulticastRendezvousPoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchRoutingMulticastRendezvousPoint`: GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateNetworkSwitchRoutingMulticastRendezvousPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rendezvousPointId** | **string** | Rendezvous point ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createNetworkSwitchRoutingMulticastRendezvousPointRequest** | [**CreateNetworkSwitchRoutingMulticastRendezvousPointRequest**](CreateNetworkSwitchRoutingMulticastRendezvousPointRequest.md) |  | 

### Return type

[**GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner**](GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchRoutingOspf

> GetNetworkSwitchRoutingOspf200Response UpdateNetworkSwitchRoutingOspf(ctx, networkId).UpdateNetworkSwitchRoutingOspfRequest(updateNetworkSwitchRoutingOspfRequest).Execute()

Update layer 3 OSPF routing configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkSwitchRoutingOspfRequest := *openapiclient.NewUpdateNetworkSwitchRoutingOspfRequest() // UpdateNetworkSwitchRoutingOspfRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateNetworkSwitchRoutingOspf(context.Background(), networkId).UpdateNetworkSwitchRoutingOspfRequest(updateNetworkSwitchRoutingOspfRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateNetworkSwitchRoutingOspf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchRoutingOspf`: GetNetworkSwitchRoutingOspf200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateNetworkSwitchRoutingOspf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchRoutingOspfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchRoutingOspfRequest** | [**UpdateNetworkSwitchRoutingOspfRequest**](UpdateNetworkSwitchRoutingOspfRequest.md) |  | 

### Return type

[**GetNetworkSwitchRoutingOspf200Response**](GetNetworkSwitchRoutingOspf200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchSettings

> GetNetworkSwitchSettings200Response UpdateNetworkSwitchSettings(ctx, networkId).UpdateNetworkSwitchSettingsRequest(updateNetworkSwitchSettingsRequest).Execute()

Update switch network settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkSwitchSettingsRequest := *openapiclient.NewUpdateNetworkSwitchSettingsRequest() // UpdateNetworkSwitchSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateNetworkSwitchSettings(context.Background(), networkId).UpdateNetworkSwitchSettingsRequest(updateNetworkSwitchSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateNetworkSwitchSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchSettings`: GetNetworkSwitchSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateNetworkSwitchSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchSettingsRequest** | [**UpdateNetworkSwitchSettingsRequest**](UpdateNetworkSwitchSettingsRequest.md) |  | 

### Return type

[**GetNetworkSwitchSettings200Response**](GetNetworkSwitchSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchStackRoutingInterface

> UpdateNetworkSwitchStackRoutingInterface200Response UpdateNetworkSwitchStackRoutingInterface(ctx, networkId, switchStackId, interfaceId).UpdateNetworkSwitchStackRoutingInterfaceRequest(updateNetworkSwitchStackRoutingInterfaceRequest).Execute()

Update a layer 3 interface for a switch stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	switchStackId := "switchStackId_example" // string | Switch stack ID
	interfaceId := "interfaceId_example" // string | Interface ID
	updateNetworkSwitchStackRoutingInterfaceRequest := *openapiclient.NewUpdateNetworkSwitchStackRoutingInterfaceRequest() // UpdateNetworkSwitchStackRoutingInterfaceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateNetworkSwitchStackRoutingInterface(context.Background(), networkId, switchStackId, interfaceId).UpdateNetworkSwitchStackRoutingInterfaceRequest(updateNetworkSwitchStackRoutingInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateNetworkSwitchStackRoutingInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchStackRoutingInterface`: UpdateNetworkSwitchStackRoutingInterface200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateNetworkSwitchStackRoutingInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchStackRoutingInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateNetworkSwitchStackRoutingInterfaceRequest** | [**UpdateNetworkSwitchStackRoutingInterfaceRequest**](UpdateNetworkSwitchStackRoutingInterfaceRequest.md) |  | 

### Return type

[**UpdateNetworkSwitchStackRoutingInterface200Response**](UpdateNetworkSwitchStackRoutingInterface200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchStackRoutingInterfaceDhcp

> GetDeviceSwitchRoutingInterfaceDhcp200Response UpdateNetworkSwitchStackRoutingInterfaceDhcp(ctx, networkId, switchStackId, interfaceId).UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest(updateNetworkSwitchStackRoutingInterfaceDhcpRequest).Execute()

Update a layer 3 interface DHCP configuration for a switch stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	switchStackId := "switchStackId_example" // string | Switch stack ID
	interfaceId := "interfaceId_example" // string | Interface ID
	updateNetworkSwitchStackRoutingInterfaceDhcpRequest := *openapiclient.NewUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest() // UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateNetworkSwitchStackRoutingInterfaceDhcp(context.Background(), networkId, switchStackId, interfaceId).UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest(updateNetworkSwitchStackRoutingInterfaceDhcpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateNetworkSwitchStackRoutingInterfaceDhcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchStackRoutingInterfaceDhcp`: GetDeviceSwitchRoutingInterfaceDhcp200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateNetworkSwitchStackRoutingInterfaceDhcp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**interfaceId** | **string** | Interface ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchStackRoutingInterfaceDhcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateNetworkSwitchStackRoutingInterfaceDhcpRequest** | [**UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest**](UpdateNetworkSwitchStackRoutingInterfaceDhcpRequest.md) |  | 

### Return type

[**GetDeviceSwitchRoutingInterfaceDhcp200Response**](GetDeviceSwitchRoutingInterfaceDhcp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchStackRoutingStaticRoute

> GetDeviceSwitchRoutingStaticRoutes200ResponseInner UpdateNetworkSwitchStackRoutingStaticRoute(ctx, networkId, switchStackId, staticRouteId).UpdateDeviceSwitchRoutingStaticRouteRequest(updateDeviceSwitchRoutingStaticRouteRequest).Execute()

Update a layer 3 static route for a switch stack



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	switchStackId := "switchStackId_example" // string | Switch stack ID
	staticRouteId := "staticRouteId_example" // string | Static route ID
	updateDeviceSwitchRoutingStaticRouteRequest := *openapiclient.NewUpdateDeviceSwitchRoutingStaticRouteRequest() // UpdateDeviceSwitchRoutingStaticRouteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateNetworkSwitchStackRoutingStaticRoute(context.Background(), networkId, switchStackId, staticRouteId).UpdateDeviceSwitchRoutingStaticRouteRequest(updateDeviceSwitchRoutingStaticRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateNetworkSwitchStackRoutingStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchStackRoutingStaticRoute`: GetDeviceSwitchRoutingStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateNetworkSwitchStackRoutingStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**switchStackId** | **string** | Switch stack ID | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchStackRoutingStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateDeviceSwitchRoutingStaticRouteRequest** | [**UpdateDeviceSwitchRoutingStaticRouteRequest**](UpdateDeviceSwitchRoutingStaticRouteRequest.md) |  | 

### Return type

[**GetDeviceSwitchRoutingStaticRoutes200ResponseInner**](GetDeviceSwitchRoutingStaticRoutes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchStormControl

> GetNetworkSwitchStormControl200Response UpdateNetworkSwitchStormControl(ctx, networkId).UpdateNetworkSwitchStormControlRequest(updateNetworkSwitchStormControlRequest).Execute()

Update the storm control configuration for a switch network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkSwitchStormControlRequest := *openapiclient.NewUpdateNetworkSwitchStormControlRequest() // UpdateNetworkSwitchStormControlRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateNetworkSwitchStormControl(context.Background(), networkId).UpdateNetworkSwitchStormControlRequest(updateNetworkSwitchStormControlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateNetworkSwitchStormControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchStormControl`: GetNetworkSwitchStormControl200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateNetworkSwitchStormControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchStormControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchStormControlRequest** | [**UpdateNetworkSwitchStormControlRequest**](UpdateNetworkSwitchStormControlRequest.md) |  | 

### Return type

[**GetNetworkSwitchStormControl200Response**](GetNetworkSwitchStormControl200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchStp

> GetNetworkSwitchStp200Response UpdateNetworkSwitchStp(ctx, networkId).UpdateNetworkSwitchStpRequest(updateNetworkSwitchStpRequest).Execute()

Updates STP settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkSwitchStpRequest := *openapiclient.NewUpdateNetworkSwitchStpRequest() // UpdateNetworkSwitchStpRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateNetworkSwitchStp(context.Background(), networkId).UpdateNetworkSwitchStpRequest(updateNetworkSwitchStpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateNetworkSwitchStp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchStp`: GetNetworkSwitchStp200Response
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateNetworkSwitchStp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchStpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchStpRequest** | [**UpdateNetworkSwitchStpRequest**](UpdateNetworkSwitchStpRequest.md) |  | 

### Return type

[**GetNetworkSwitchStp200Response**](GetNetworkSwitchStp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationConfigTemplateSwitchProfilePort

> GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner UpdateOrganizationConfigTemplateSwitchProfilePort(ctx, organizationId, configTemplateId, profileId, portId).UpdateOrganizationConfigTemplateSwitchProfilePortRequest(updateOrganizationConfigTemplateSwitchProfilePortRequest).Execute()

Update a switch template port



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	configTemplateId := "configTemplateId_example" // string | Config template ID
	profileId := "profileId_example" // string | Profile ID
	portId := "portId_example" // string | Port ID
	updateOrganizationConfigTemplateSwitchProfilePortRequest := *openapiclient.NewUpdateOrganizationConfigTemplateSwitchProfilePortRequest() // UpdateOrganizationConfigTemplateSwitchProfilePortRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SwitchAPI.UpdateOrganizationConfigTemplateSwitchProfilePort(context.Background(), organizationId, configTemplateId, profileId, portId).UpdateOrganizationConfigTemplateSwitchProfilePortRequest(updateOrganizationConfigTemplateSwitchProfilePortRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SwitchAPI.UpdateOrganizationConfigTemplateSwitchProfilePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationConfigTemplateSwitchProfilePort`: GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SwitchAPI.UpdateOrganizationConfigTemplateSwitchProfilePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 
**profileId** | **string** | Profile ID | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationConfigTemplateSwitchProfilePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateOrganizationConfigTemplateSwitchProfilePortRequest** | [**UpdateOrganizationConfigTemplateSwitchProfilePortRequest**](UpdateOrganizationConfigTemplateSwitchProfilePortRequest.md) |  | 

### Return type

[**GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner**](GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

