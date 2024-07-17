# \NetworksAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BindNetwork**](NetworksAPI.md#BindNetwork) | **Post** /networks/{networkId}/bind | Bind a network to a template.
[**ClaimNetworkDevices**](NetworksAPI.md#ClaimNetworkDevices) | **Post** /networks/{networkId}/devices/claim | Claim devices into a network. (Note: for recently claimed devices, it may take a few minutes for API requests against that device to succeed)
[**CombineOrganizationNetworks**](NetworksAPI.md#CombineOrganizationNetworks) | **Post** /organizations/{organizationId}/networks/combine | Combine multiple networks into a single network
[**CreateNetworkFirmwareUpgradesRollback**](NetworksAPI.md#CreateNetworkFirmwareUpgradesRollback) | **Post** /networks/{networkId}/firmwareUpgrades/rollbacks | Rollback a Firmware Upgrade For A Network
[**CreateNetworkFirmwareUpgradesStagedEvent**](NetworksAPI.md#CreateNetworkFirmwareUpgradesStagedEvent) | **Post** /networks/{networkId}/firmwareUpgrades/staged/events | Create a Staged Upgrade Event for a network
[**CreateNetworkFirmwareUpgradesStagedGroup**](NetworksAPI.md#CreateNetworkFirmwareUpgradesStagedGroup) | **Post** /networks/{networkId}/firmwareUpgrades/staged/groups | Create a Staged Upgrade Group for a network
[**CreateNetworkFloorPlan**](NetworksAPI.md#CreateNetworkFloorPlan) | **Post** /networks/{networkId}/floorPlans | Upload a floor plan
[**CreateNetworkGroupPolicy**](NetworksAPI.md#CreateNetworkGroupPolicy) | **Post** /networks/{networkId}/groupPolicies | Create a group policy
[**CreateNetworkMerakiAuthUser**](NetworksAPI.md#CreateNetworkMerakiAuthUser) | **Post** /networks/{networkId}/merakiAuthUsers | Authorize a user configured with Meraki Authentication for a network (currently supports 802.1X, splash guest, and client VPN users, and currently, organizations have a 50,000 user cap)
[**CreateNetworkMqttBroker**](NetworksAPI.md#CreateNetworkMqttBroker) | **Post** /networks/{networkId}/mqttBrokers | Add an MQTT broker
[**CreateNetworkPiiRequest**](NetworksAPI.md#CreateNetworkPiiRequest) | **Post** /networks/{networkId}/pii/requests | Submit a new delete or restrict processing PII request
[**CreateNetworkVlanProfile**](NetworksAPI.md#CreateNetworkVlanProfile) | **Post** /networks/{networkId}/vlanProfiles | Create a VLAN profile for a network
[**CreateNetworkWebhooksHttpServer**](NetworksAPI.md#CreateNetworkWebhooksHttpServer) | **Post** /networks/{networkId}/webhooks/httpServers | Add an HTTP server to a network
[**CreateNetworkWebhooksPayloadTemplate**](NetworksAPI.md#CreateNetworkWebhooksPayloadTemplate) | **Post** /networks/{networkId}/webhooks/payloadTemplates | Create a webhook payload template for a network
[**CreateNetworkWebhooksWebhookTest**](NetworksAPI.md#CreateNetworkWebhooksWebhookTest) | **Post** /networks/{networkId}/webhooks/webhookTests | Send a test webhook for a network
[**CreateOrganizationNetwork**](NetworksAPI.md#CreateOrganizationNetwork) | **Post** /organizations/{organizationId}/networks | Create a network
[**DeferNetworkFirmwareUpgradesStagedEvents**](NetworksAPI.md#DeferNetworkFirmwareUpgradesStagedEvents) | **Post** /networks/{networkId}/firmwareUpgrades/staged/events/defer | Postpone by 1 week all pending staged upgrade stages for a network
[**DeleteNetwork**](NetworksAPI.md#DeleteNetwork) | **Delete** /networks/{networkId} | Delete a network
[**DeleteNetworkFirmwareUpgradesStagedGroup**](NetworksAPI.md#DeleteNetworkFirmwareUpgradesStagedGroup) | **Delete** /networks/{networkId}/firmwareUpgrades/staged/groups/{groupId} | Delete a Staged Upgrade Group
[**DeleteNetworkFloorPlan**](NetworksAPI.md#DeleteNetworkFloorPlan) | **Delete** /networks/{networkId}/floorPlans/{floorPlanId} | Destroy a floor plan
[**DeleteNetworkGroupPolicy**](NetworksAPI.md#DeleteNetworkGroupPolicy) | **Delete** /networks/{networkId}/groupPolicies/{groupPolicyId} | Delete a group policy
[**DeleteNetworkMerakiAuthUser**](NetworksAPI.md#DeleteNetworkMerakiAuthUser) | **Delete** /networks/{networkId}/merakiAuthUsers/{merakiAuthUserId} | Delete an 802.1X RADIUS user, or deauthorize and optionally delete a splash guest or client VPN user.
[**DeleteNetworkMqttBroker**](NetworksAPI.md#DeleteNetworkMqttBroker) | **Delete** /networks/{networkId}/mqttBrokers/{mqttBrokerId} | Delete an MQTT broker
[**DeleteNetworkPiiRequest**](NetworksAPI.md#DeleteNetworkPiiRequest) | **Delete** /networks/{networkId}/pii/requests/{requestId} | Delete a restrict processing PII request
[**DeleteNetworkVlanProfile**](NetworksAPI.md#DeleteNetworkVlanProfile) | **Delete** /networks/{networkId}/vlanProfiles/{iname} | Delete a VLAN profile of a network
[**DeleteNetworkWebhooksHttpServer**](NetworksAPI.md#DeleteNetworkWebhooksHttpServer) | **Delete** /networks/{networkId}/webhooks/httpServers/{httpServerId} | Delete an HTTP server from a network
[**DeleteNetworkWebhooksPayloadTemplate**](NetworksAPI.md#DeleteNetworkWebhooksPayloadTemplate) | **Delete** /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId} | Destroy a webhook payload template for a network
[**GetNetwork**](NetworksAPI.md#GetNetwork) | **Get** /networks/{networkId} | Return a network
[**GetNetworkAlertsHistory**](NetworksAPI.md#GetNetworkAlertsHistory) | **Get** /networks/{networkId}/alerts/history | Return the alert history for this network
[**GetNetworkAlertsSettings**](NetworksAPI.md#GetNetworkAlertsSettings) | **Get** /networks/{networkId}/alerts/settings | Return the alert configuration for this network
[**GetNetworkBluetoothClient**](NetworksAPI.md#GetNetworkBluetoothClient) | **Get** /networks/{networkId}/bluetoothClients/{bluetoothClientId} | Return a Bluetooth client
[**GetNetworkBluetoothClients**](NetworksAPI.md#GetNetworkBluetoothClients) | **Get** /networks/{networkId}/bluetoothClients | List the Bluetooth clients seen by APs in this network
[**GetNetworkClient**](NetworksAPI.md#GetNetworkClient) | **Get** /networks/{networkId}/clients/{clientId} | Return the client associated with the given identifier
[**GetNetworkClientPolicy**](NetworksAPI.md#GetNetworkClientPolicy) | **Get** /networks/{networkId}/clients/{clientId}/policy | Return the policy assigned to a client on the network
[**GetNetworkClientSplashAuthorizationStatus**](NetworksAPI.md#GetNetworkClientSplashAuthorizationStatus) | **Get** /networks/{networkId}/clients/{clientId}/splashAuthorizationStatus | Return the splash authorization for a client, for each SSID they&#39;ve associated with through splash
[**GetNetworkClientTrafficHistory**](NetworksAPI.md#GetNetworkClientTrafficHistory) | **Get** /networks/{networkId}/clients/{clientId}/trafficHistory | Return the client&#39;s network traffic data over time
[**GetNetworkClientUsageHistory**](NetworksAPI.md#GetNetworkClientUsageHistory) | **Get** /networks/{networkId}/clients/{clientId}/usageHistory | Return the client&#39;s daily usage history
[**GetNetworkClients**](NetworksAPI.md#GetNetworkClients) | **Get** /networks/{networkId}/clients | List the clients that have used this network in the timespan
[**GetNetworkClientsApplicationUsage**](NetworksAPI.md#GetNetworkClientsApplicationUsage) | **Get** /networks/{networkId}/clients/applicationUsage | Return the application usage data for clients
[**GetNetworkClientsBandwidthUsageHistory**](NetworksAPI.md#GetNetworkClientsBandwidthUsageHistory) | **Get** /networks/{networkId}/clients/bandwidthUsageHistory | Returns a timeseries of total traffic consumption rates for all clients on a network within a given timespan, in megabits per second.
[**GetNetworkClientsOverview**](NetworksAPI.md#GetNetworkClientsOverview) | **Get** /networks/{networkId}/clients/overview | Return overview statistics for network clients
[**GetNetworkClientsUsageHistories**](NetworksAPI.md#GetNetworkClientsUsageHistories) | **Get** /networks/{networkId}/clients/usageHistories | Return the usage histories for clients
[**GetNetworkDevices**](NetworksAPI.md#GetNetworkDevices) | **Get** /networks/{networkId}/devices | List the devices in a network
[**GetNetworkEvents**](NetworksAPI.md#GetNetworkEvents) | **Get** /networks/{networkId}/events | List the events for the network
[**GetNetworkEventsEventTypes**](NetworksAPI.md#GetNetworkEventsEventTypes) | **Get** /networks/{networkId}/events/eventTypes | List the event type to human-readable description
[**GetNetworkFirmwareUpgrades**](NetworksAPI.md#GetNetworkFirmwareUpgrades) | **Get** /networks/{networkId}/firmwareUpgrades | Get firmware upgrade information for a network
[**GetNetworkFirmwareUpgradesStagedEvents**](NetworksAPI.md#GetNetworkFirmwareUpgradesStagedEvents) | **Get** /networks/{networkId}/firmwareUpgrades/staged/events | Get the Staged Upgrade Event from a network
[**GetNetworkFirmwareUpgradesStagedGroup**](NetworksAPI.md#GetNetworkFirmwareUpgradesStagedGroup) | **Get** /networks/{networkId}/firmwareUpgrades/staged/groups/{groupId} | Get a Staged Upgrade Group from a network
[**GetNetworkFirmwareUpgradesStagedGroups**](NetworksAPI.md#GetNetworkFirmwareUpgradesStagedGroups) | **Get** /networks/{networkId}/firmwareUpgrades/staged/groups | List of Staged Upgrade Groups in a network
[**GetNetworkFirmwareUpgradesStagedStages**](NetworksAPI.md#GetNetworkFirmwareUpgradesStagedStages) | **Get** /networks/{networkId}/firmwareUpgrades/staged/stages | Order of Staged Upgrade Groups in a network
[**GetNetworkFloorPlan**](NetworksAPI.md#GetNetworkFloorPlan) | **Get** /networks/{networkId}/floorPlans/{floorPlanId} | Find a floor plan by ID
[**GetNetworkFloorPlans**](NetworksAPI.md#GetNetworkFloorPlans) | **Get** /networks/{networkId}/floorPlans | List the floor plans that belong to your network
[**GetNetworkGroupPolicies**](NetworksAPI.md#GetNetworkGroupPolicies) | **Get** /networks/{networkId}/groupPolicies | List the group policies in a network
[**GetNetworkGroupPolicy**](NetworksAPI.md#GetNetworkGroupPolicy) | **Get** /networks/{networkId}/groupPolicies/{groupPolicyId} | Display a group policy
[**GetNetworkHealthAlerts**](NetworksAPI.md#GetNetworkHealthAlerts) | **Get** /networks/{networkId}/health/alerts | Return all global alerts on this network
[**GetNetworkMerakiAuthUser**](NetworksAPI.md#GetNetworkMerakiAuthUser) | **Get** /networks/{networkId}/merakiAuthUsers/{merakiAuthUserId} | Return the Meraki Auth splash guest, RADIUS, or client VPN user
[**GetNetworkMerakiAuthUsers**](NetworksAPI.md#GetNetworkMerakiAuthUsers) | **Get** /networks/{networkId}/merakiAuthUsers | List the users configured under Meraki Authentication for a network (splash guest or RADIUS users for a wireless network, or client VPN users for a MX network)
[**GetNetworkMqttBroker**](NetworksAPI.md#GetNetworkMqttBroker) | **Get** /networks/{networkId}/mqttBrokers/{mqttBrokerId} | Return an MQTT broker
[**GetNetworkMqttBrokers**](NetworksAPI.md#GetNetworkMqttBrokers) | **Get** /networks/{networkId}/mqttBrokers | List the MQTT brokers for this network
[**GetNetworkNetflow**](NetworksAPI.md#GetNetworkNetflow) | **Get** /networks/{networkId}/netflow | Return the NetFlow traffic reporting settings for a network
[**GetNetworkNetworkHealthChannelUtilization**](NetworksAPI.md#GetNetworkNetworkHealthChannelUtilization) | **Get** /networks/{networkId}/networkHealth/channelUtilization | Get the channel utilization over each radio for all APs in a network.
[**GetNetworkPiiPiiKeys**](NetworksAPI.md#GetNetworkPiiPiiKeys) | **Get** /networks/{networkId}/pii/piiKeys | List the keys required to access Personally Identifiable Information (PII) for a given identifier
[**GetNetworkPiiRequest**](NetworksAPI.md#GetNetworkPiiRequest) | **Get** /networks/{networkId}/pii/requests/{requestId} | Return a PII request
[**GetNetworkPiiRequests**](NetworksAPI.md#GetNetworkPiiRequests) | **Get** /networks/{networkId}/pii/requests | List the PII requests for this network or organization
[**GetNetworkPiiSmDevicesForKey**](NetworksAPI.md#GetNetworkPiiSmDevicesForKey) | **Get** /networks/{networkId}/pii/smDevicesForKey | Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier
[**GetNetworkPiiSmOwnersForKey**](NetworksAPI.md#GetNetworkPiiSmOwnersForKey) | **Get** /networks/{networkId}/pii/smOwnersForKey | Given a piece of Personally Identifiable Information (PII), return the Systems Manager owner ID(s) associated with that identifier
[**GetNetworkPoliciesByClient**](NetworksAPI.md#GetNetworkPoliciesByClient) | **Get** /networks/{networkId}/policies/byClient | Get policies for all clients with policies
[**GetNetworkSettings**](NetworksAPI.md#GetNetworkSettings) | **Get** /networks/{networkId}/settings | Return the settings for a network
[**GetNetworkSnmp**](NetworksAPI.md#GetNetworkSnmp) | **Get** /networks/{networkId}/snmp | Return the SNMP settings for a network
[**GetNetworkSplashLoginAttempts**](NetworksAPI.md#GetNetworkSplashLoginAttempts) | **Get** /networks/{networkId}/splashLoginAttempts | List the splash login attempts for a network
[**GetNetworkSyslogServers**](NetworksAPI.md#GetNetworkSyslogServers) | **Get** /networks/{networkId}/syslogServers | List the syslog servers for a network
[**GetNetworkTopologyLinkLayer**](NetworksAPI.md#GetNetworkTopologyLinkLayer) | **Get** /networks/{networkId}/topology/linkLayer | List the LLDP and CDP information for all discovered devices and connections in a network
[**GetNetworkTraffic**](NetworksAPI.md#GetNetworkTraffic) | **Get** /networks/{networkId}/traffic | Return the traffic analysis data for this network
[**GetNetworkTrafficAnalysis**](NetworksAPI.md#GetNetworkTrafficAnalysis) | **Get** /networks/{networkId}/trafficAnalysis | Return the traffic analysis settings for a network
[**GetNetworkTrafficShapingApplicationCategories**](NetworksAPI.md#GetNetworkTrafficShapingApplicationCategories) | **Get** /networks/{networkId}/trafficShaping/applicationCategories | Returns the application categories for traffic shaping rules
[**GetNetworkTrafficShapingDscpTaggingOptions**](NetworksAPI.md#GetNetworkTrafficShapingDscpTaggingOptions) | **Get** /networks/{networkId}/trafficShaping/dscpTaggingOptions | Returns the available DSCP tagging options for your traffic shaping rules.
[**GetNetworkVlanProfile**](NetworksAPI.md#GetNetworkVlanProfile) | **Get** /networks/{networkId}/vlanProfiles/{iname} | Get an existing VLAN profile of a network
[**GetNetworkVlanProfiles**](NetworksAPI.md#GetNetworkVlanProfiles) | **Get** /networks/{networkId}/vlanProfiles | List VLAN profiles for a network
[**GetNetworkVlanProfilesAssignmentsByDevice**](NetworksAPI.md#GetNetworkVlanProfilesAssignmentsByDevice) | **Get** /networks/{networkId}/vlanProfiles/assignments/byDevice | Get the assigned VLAN Profiles for devices in a network
[**GetNetworkWebhooksHttpServer**](NetworksAPI.md#GetNetworkWebhooksHttpServer) | **Get** /networks/{networkId}/webhooks/httpServers/{httpServerId} | Return an HTTP server for a network
[**GetNetworkWebhooksHttpServers**](NetworksAPI.md#GetNetworkWebhooksHttpServers) | **Get** /networks/{networkId}/webhooks/httpServers | List the HTTP servers for a network
[**GetNetworkWebhooksPayloadTemplate**](NetworksAPI.md#GetNetworkWebhooksPayloadTemplate) | **Get** /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId} | Get the webhook payload template for a network
[**GetNetworkWebhooksPayloadTemplates**](NetworksAPI.md#GetNetworkWebhooksPayloadTemplates) | **Get** /networks/{networkId}/webhooks/payloadTemplates | List the webhook payload templates for a network
[**GetNetworkWebhooksWebhookTest**](NetworksAPI.md#GetNetworkWebhooksWebhookTest) | **Get** /networks/{networkId}/webhooks/webhookTests/{webhookTestId} | Return the status of a webhook test for a network
[**GetOrganizationInventoryOnboardingCloudMonitoringNetworks**](NetworksAPI.md#GetOrganizationInventoryOnboardingCloudMonitoringNetworks) | **Get** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/networks | Returns list of networks eligible for adding cloud monitored device
[**GetOrganizationNetworks**](NetworksAPI.md#GetOrganizationNetworks) | **Get** /organizations/{organizationId}/networks | List the networks that the user has privileges on in an organization
[**GetOrganizationSummaryTopNetworksByStatus**](NetworksAPI.md#GetOrganizationSummaryTopNetworksByStatus) | **Get** /organizations/{organizationId}/summary/top/networks/byStatus | List the client and status overview information for the networks in an organization
[**ProvisionNetworkClients**](NetworksAPI.md#ProvisionNetworkClients) | **Post** /networks/{networkId}/clients/provision | Provisions a client with a name and policy
[**ReassignNetworkVlanProfilesAssignments**](NetworksAPI.md#ReassignNetworkVlanProfilesAssignments) | **Post** /networks/{networkId}/vlanProfiles/assignments/reassign | Update the assigned VLAN Profile for devices in a network
[**RemoveNetworkDevices**](NetworksAPI.md#RemoveNetworkDevices) | **Post** /networks/{networkId}/devices/remove | Remove a single device
[**RollbacksNetworkFirmwareUpgradesStagedEvents**](NetworksAPI.md#RollbacksNetworkFirmwareUpgradesStagedEvents) | **Post** /networks/{networkId}/firmwareUpgrades/staged/events/rollbacks | Rollback a Staged Upgrade Event for a network
[**SplitNetwork**](NetworksAPI.md#SplitNetwork) | **Post** /networks/{networkId}/split | Split a combined network into individual networks for each type of device
[**UnbindNetwork**](NetworksAPI.md#UnbindNetwork) | **Post** /networks/{networkId}/unbind | Unbind a network from a template.
[**UpdateNetwork**](NetworksAPI.md#UpdateNetwork) | **Put** /networks/{networkId} | Update a network
[**UpdateNetworkAlertsSettings**](NetworksAPI.md#UpdateNetworkAlertsSettings) | **Put** /networks/{networkId}/alerts/settings | Update the alert configuration for this network
[**UpdateNetworkClientPolicy**](NetworksAPI.md#UpdateNetworkClientPolicy) | **Put** /networks/{networkId}/clients/{clientId}/policy | Update the policy assigned to a client on the network
[**UpdateNetworkClientSplashAuthorizationStatus**](NetworksAPI.md#UpdateNetworkClientSplashAuthorizationStatus) | **Put** /networks/{networkId}/clients/{clientId}/splashAuthorizationStatus | Update a client&#39;s splash authorization
[**UpdateNetworkFirmwareUpgrades**](NetworksAPI.md#UpdateNetworkFirmwareUpgrades) | **Put** /networks/{networkId}/firmwareUpgrades | Update firmware upgrade information for a network
[**UpdateNetworkFirmwareUpgradesStagedEvents**](NetworksAPI.md#UpdateNetworkFirmwareUpgradesStagedEvents) | **Put** /networks/{networkId}/firmwareUpgrades/staged/events | Update the Staged Upgrade Event for a network
[**UpdateNetworkFirmwareUpgradesStagedGroup**](NetworksAPI.md#UpdateNetworkFirmwareUpgradesStagedGroup) | **Put** /networks/{networkId}/firmwareUpgrades/staged/groups/{groupId} | Update a Staged Upgrade Group for a network
[**UpdateNetworkFirmwareUpgradesStagedStages**](NetworksAPI.md#UpdateNetworkFirmwareUpgradesStagedStages) | **Put** /networks/{networkId}/firmwareUpgrades/staged/stages | Assign Staged Upgrade Group order in the sequence.
[**UpdateNetworkFloorPlan**](NetworksAPI.md#UpdateNetworkFloorPlan) | **Put** /networks/{networkId}/floorPlans/{floorPlanId} | Update a floor plan&#39;s geolocation and other meta data
[**UpdateNetworkGroupPolicy**](NetworksAPI.md#UpdateNetworkGroupPolicy) | **Put** /networks/{networkId}/groupPolicies/{groupPolicyId} | Update a group policy
[**UpdateNetworkMerakiAuthUser**](NetworksAPI.md#UpdateNetworkMerakiAuthUser) | **Put** /networks/{networkId}/merakiAuthUsers/{merakiAuthUserId} | Update a user configured with Meraki Authentication (currently, 802.1X RADIUS, splash guest, and client VPN users can be updated)
[**UpdateNetworkMqttBroker**](NetworksAPI.md#UpdateNetworkMqttBroker) | **Put** /networks/{networkId}/mqttBrokers/{mqttBrokerId} | Update an MQTT broker
[**UpdateNetworkNetflow**](NetworksAPI.md#UpdateNetworkNetflow) | **Put** /networks/{networkId}/netflow | Update the NetFlow traffic reporting settings for a network
[**UpdateNetworkSettings**](NetworksAPI.md#UpdateNetworkSettings) | **Put** /networks/{networkId}/settings | Update the settings for a network
[**UpdateNetworkSnmp**](NetworksAPI.md#UpdateNetworkSnmp) | **Put** /networks/{networkId}/snmp | Update the SNMP settings for a network
[**UpdateNetworkSyslogServers**](NetworksAPI.md#UpdateNetworkSyslogServers) | **Put** /networks/{networkId}/syslogServers | Update the syslog servers for a network
[**UpdateNetworkTrafficAnalysis**](NetworksAPI.md#UpdateNetworkTrafficAnalysis) | **Put** /networks/{networkId}/trafficAnalysis | Update the traffic analysis settings for a network
[**UpdateNetworkVlanProfile**](NetworksAPI.md#UpdateNetworkVlanProfile) | **Put** /networks/{networkId}/vlanProfiles/{iname} | Update an existing VLAN profile of a network
[**UpdateNetworkWebhooksHttpServer**](NetworksAPI.md#UpdateNetworkWebhooksHttpServer) | **Put** /networks/{networkId}/webhooks/httpServers/{httpServerId} | Update an HTTP server
[**UpdateNetworkWebhooksPayloadTemplate**](NetworksAPI.md#UpdateNetworkWebhooksPayloadTemplate) | **Put** /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId} | Update a webhook payload template for a network
[**VmxNetworkDevicesClaim**](NetworksAPI.md#VmxNetworkDevicesClaim) | **Post** /networks/{networkId}/devices/claim/vmx | Claim a vMX into a network



## BindNetwork

> BindNetwork200Response BindNetwork(ctx, networkId).BindNetworkRequest(bindNetworkRequest).Execute()

Bind a network to a template.



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
	bindNetworkRequest := *openapiclient.NewBindNetworkRequest("ConfigTemplateId_example") // BindNetworkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.BindNetwork(context.Background(), networkId).BindNetworkRequest(bindNetworkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.BindNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BindNetwork`: BindNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.BindNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBindNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bindNetworkRequest** | [**BindNetworkRequest**](BindNetworkRequest.md) |  | 

### Return type

[**BindNetwork200Response**](BindNetwork200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClaimNetworkDevices

> ClaimNetworkDevices200Response ClaimNetworkDevices(ctx, networkId).ClaimNetworkDevicesRequest(claimNetworkDevicesRequest).AddAtomically(addAtomically).Execute()

Claim devices into a network. (Note: for recently claimed devices, it may take a few minutes for API requests against that device to succeed)



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
	claimNetworkDevicesRequest := *openapiclient.NewClaimNetworkDevicesRequest([]string{"Serials_example"}) // ClaimNetworkDevicesRequest | 
	addAtomically := true // bool | Whether to claim devices atomically. If true, all devices will be claimed or none will be claimed. Default is true. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.ClaimNetworkDevices(context.Background(), networkId).ClaimNetworkDevicesRequest(claimNetworkDevicesRequest).AddAtomically(addAtomically).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.ClaimNetworkDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClaimNetworkDevices`: ClaimNetworkDevices200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.ClaimNetworkDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimNetworkDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **claimNetworkDevicesRequest** | [**ClaimNetworkDevicesRequest**](ClaimNetworkDevicesRequest.md) |  | 
 **addAtomically** | **bool** | Whether to claim devices atomically. If true, all devices will be claimed or none will be claimed. Default is true. | 

### Return type

[**ClaimNetworkDevices200Response**](ClaimNetworkDevices200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CombineOrganizationNetworks

> CombineOrganizationNetworks200Response CombineOrganizationNetworks(ctx, organizationId).CombineOrganizationNetworksRequest(combineOrganizationNetworksRequest).Execute()

Combine multiple networks into a single network



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
	combineOrganizationNetworksRequest := *openapiclient.NewCombineOrganizationNetworksRequest("Name_example", []string{"NetworkIds_example"}) // CombineOrganizationNetworksRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CombineOrganizationNetworks(context.Background(), organizationId).CombineOrganizationNetworksRequest(combineOrganizationNetworksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CombineOrganizationNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CombineOrganizationNetworks`: CombineOrganizationNetworks200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CombineOrganizationNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCombineOrganizationNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **combineOrganizationNetworksRequest** | [**CombineOrganizationNetworksRequest**](CombineOrganizationNetworksRequest.md) |  | 

### Return type

[**CombineOrganizationNetworks200Response**](CombineOrganizationNetworks200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkFirmwareUpgradesRollback

> CreateNetworkFirmwareUpgradesRollback200Response CreateNetworkFirmwareUpgradesRollback(ctx, networkId).CreateNetworkFirmwareUpgradesRollbackRequest(createNetworkFirmwareUpgradesRollbackRequest).Execute()

Rollback a Firmware Upgrade For A Network



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
	createNetworkFirmwareUpgradesRollbackRequest := *openapiclient.NewCreateNetworkFirmwareUpgradesRollbackRequest([]openapiclient.CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner{*openapiclient.NewCreateNetworkFirmwareUpgradesRollbackRequestReasonsInner("Category_example", "Comment_example")}) // CreateNetworkFirmwareUpgradesRollbackRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkFirmwareUpgradesRollback(context.Background(), networkId).CreateNetworkFirmwareUpgradesRollbackRequest(createNetworkFirmwareUpgradesRollbackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkFirmwareUpgradesRollback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkFirmwareUpgradesRollback`: CreateNetworkFirmwareUpgradesRollback200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkFirmwareUpgradesRollback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFirmwareUpgradesRollbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkFirmwareUpgradesRollbackRequest** | [**CreateNetworkFirmwareUpgradesRollbackRequest**](CreateNetworkFirmwareUpgradesRollbackRequest.md) |  | 

### Return type

[**CreateNetworkFirmwareUpgradesRollback200Response**](CreateNetworkFirmwareUpgradesRollback200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkFirmwareUpgradesStagedEvent

> GetNetworkFirmwareUpgradesStagedEvents200Response CreateNetworkFirmwareUpgradesStagedEvent(ctx, networkId).CreateNetworkFirmwareUpgradesStagedEventRequest(createNetworkFirmwareUpgradesStagedEventRequest).Execute()

Create a Staged Upgrade Event for a network



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
	createNetworkFirmwareUpgradesStagedEventRequest := *openapiclient.NewCreateNetworkFirmwareUpgradesStagedEventRequest([]openapiclient.UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner{*openapiclient.NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner()}) // CreateNetworkFirmwareUpgradesStagedEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkFirmwareUpgradesStagedEvent(context.Background(), networkId).CreateNetworkFirmwareUpgradesStagedEventRequest(createNetworkFirmwareUpgradesStagedEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkFirmwareUpgradesStagedEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkFirmwareUpgradesStagedEvent`: GetNetworkFirmwareUpgradesStagedEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkFirmwareUpgradesStagedEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFirmwareUpgradesStagedEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkFirmwareUpgradesStagedEventRequest** | [**CreateNetworkFirmwareUpgradesStagedEventRequest**](CreateNetworkFirmwareUpgradesStagedEventRequest.md) |  | 

### Return type

[**GetNetworkFirmwareUpgradesStagedEvents200Response**](GetNetworkFirmwareUpgradesStagedEvents200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkFirmwareUpgradesStagedGroup

> GetNetworkFirmwareUpgradesStagedGroups200ResponseInner CreateNetworkFirmwareUpgradesStagedGroup(ctx, networkId).CreateNetworkFirmwareUpgradesStagedGroupRequest(createNetworkFirmwareUpgradesStagedGroupRequest).Execute()

Create a Staged Upgrade Group for a network



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
	createNetworkFirmwareUpgradesStagedGroupRequest := *openapiclient.NewCreateNetworkFirmwareUpgradesStagedGroupRequest("Name_example", false) // CreateNetworkFirmwareUpgradesStagedGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId).CreateNetworkFirmwareUpgradesStagedGroupRequest(createNetworkFirmwareUpgradesStagedGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkFirmwareUpgradesStagedGroup`: GetNetworkFirmwareUpgradesStagedGroups200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkFirmwareUpgradesStagedGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFirmwareUpgradesStagedGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkFirmwareUpgradesStagedGroupRequest** | [**CreateNetworkFirmwareUpgradesStagedGroupRequest**](CreateNetworkFirmwareUpgradesStagedGroupRequest.md) |  | 

### Return type

[**GetNetworkFirmwareUpgradesStagedGroups200ResponseInner**](GetNetworkFirmwareUpgradesStagedGroups200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkFloorPlan

> GetNetworkFloorPlans200ResponseInner CreateNetworkFloorPlan(ctx, networkId).CreateNetworkFloorPlanRequest(createNetworkFloorPlanRequest).Execute()

Upload a floor plan



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
	createNetworkFloorPlanRequest := *openapiclient.NewCreateNetworkFloorPlanRequest("Name_example", string(123)) // CreateNetworkFloorPlanRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkFloorPlan(context.Background(), networkId).CreateNetworkFloorPlanRequest(createNetworkFloorPlanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkFloorPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkFloorPlan`: GetNetworkFloorPlans200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkFloorPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFloorPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkFloorPlanRequest** | [**CreateNetworkFloorPlanRequest**](CreateNetworkFloorPlanRequest.md) |  | 

### Return type

[**GetNetworkFloorPlans200ResponseInner**](GetNetworkFloorPlans200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkGroupPolicy

> GetNetworkGroupPolicies200ResponseInner CreateNetworkGroupPolicy(ctx, networkId).CreateNetworkGroupPolicyRequest(createNetworkGroupPolicyRequest).Execute()

Create a group policy



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
	createNetworkGroupPolicyRequest := *openapiclient.NewCreateNetworkGroupPolicyRequest("Name_example") // CreateNetworkGroupPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkGroupPolicy(context.Background(), networkId).CreateNetworkGroupPolicyRequest(createNetworkGroupPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkGroupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkGroupPolicy`: GetNetworkGroupPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkGroupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkGroupPolicyRequest** | [**CreateNetworkGroupPolicyRequest**](CreateNetworkGroupPolicyRequest.md) |  | 

### Return type

[**GetNetworkGroupPolicies200ResponseInner**](GetNetworkGroupPolicies200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkMerakiAuthUser

> GetNetworkMerakiAuthUsers200ResponseInner CreateNetworkMerakiAuthUser(ctx, networkId).CreateNetworkMerakiAuthUserRequest(createNetworkMerakiAuthUserRequest).Execute()

Authorize a user configured with Meraki Authentication for a network (currently supports 802.1X, splash guest, and client VPN users, and currently, organizations have a 50,000 user cap)



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
	createNetworkMerakiAuthUserRequest := *openapiclient.NewCreateNetworkMerakiAuthUserRequest("Email_example", []openapiclient.CreateNetworkMerakiAuthUserRequestAuthorizationsInner{*openapiclient.NewCreateNetworkMerakiAuthUserRequestAuthorizationsInner()}) // CreateNetworkMerakiAuthUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkMerakiAuthUser(context.Background(), networkId).CreateNetworkMerakiAuthUserRequest(createNetworkMerakiAuthUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkMerakiAuthUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkMerakiAuthUser`: GetNetworkMerakiAuthUsers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkMerakiAuthUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkMerakiAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkMerakiAuthUserRequest** | [**CreateNetworkMerakiAuthUserRequest**](CreateNetworkMerakiAuthUserRequest.md) |  | 

### Return type

[**GetNetworkMerakiAuthUsers200ResponseInner**](GetNetworkMerakiAuthUsers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkMqttBroker

> GetNetworkMqttBrokers200ResponseInner CreateNetworkMqttBroker(ctx, networkId).CreateNetworkMqttBrokerRequest(createNetworkMqttBrokerRequest).Execute()

Add an MQTT broker



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
	createNetworkMqttBrokerRequest := *openapiclient.NewCreateNetworkMqttBrokerRequest("Name_example", "Host_example", int32(123)) // CreateNetworkMqttBrokerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkMqttBroker(context.Background(), networkId).CreateNetworkMqttBrokerRequest(createNetworkMqttBrokerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkMqttBroker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkMqttBroker`: GetNetworkMqttBrokers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkMqttBroker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkMqttBrokerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkMqttBrokerRequest** | [**CreateNetworkMqttBrokerRequest**](CreateNetworkMqttBrokerRequest.md) |  | 

### Return type

[**GetNetworkMqttBrokers200ResponseInner**](GetNetworkMqttBrokers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkPiiRequest

> GetNetworkPiiRequests200ResponseInner CreateNetworkPiiRequest(ctx, networkId).CreateNetworkPiiRequestRequest(createNetworkPiiRequestRequest).Execute()

Submit a new delete or restrict processing PII request



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
	createNetworkPiiRequestRequest := *openapiclient.NewCreateNetworkPiiRequestRequest() // CreateNetworkPiiRequestRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkPiiRequest(context.Background(), networkId).CreateNetworkPiiRequestRequest(createNetworkPiiRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkPiiRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkPiiRequest`: GetNetworkPiiRequests200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkPiiRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkPiiRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkPiiRequestRequest** | [**CreateNetworkPiiRequestRequest**](CreateNetworkPiiRequestRequest.md) |  | 

### Return type

[**GetNetworkPiiRequests200ResponseInner**](GetNetworkPiiRequests200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkVlanProfile

> GetNetworkVlanProfiles200ResponseInner CreateNetworkVlanProfile(ctx, networkId).CreateNetworkVlanProfileRequest(createNetworkVlanProfileRequest).Execute()

Create a VLAN profile for a network



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
	createNetworkVlanProfileRequest := *openapiclient.NewCreateNetworkVlanProfileRequest("Name_example", []openapiclient.CreateNetworkVlanProfileRequestVlanNamesInner{*openapiclient.NewCreateNetworkVlanProfileRequestVlanNamesInner("Name_example", "VlanId_example")}, []openapiclient.CreateNetworkVlanProfileRequestVlanGroupsInner{*openapiclient.NewCreateNetworkVlanProfileRequestVlanGroupsInner("Name_example", "VlanIds_example")}, "Iname_example") // CreateNetworkVlanProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkVlanProfile(context.Background(), networkId).CreateNetworkVlanProfileRequest(createNetworkVlanProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkVlanProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkVlanProfile`: GetNetworkVlanProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkVlanProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkVlanProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkVlanProfileRequest** | [**CreateNetworkVlanProfileRequest**](CreateNetworkVlanProfileRequest.md) |  | 

### Return type

[**GetNetworkVlanProfiles200ResponseInner**](GetNetworkVlanProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkWebhooksHttpServer

> GetNetworkWebhooksHttpServers200ResponseInner CreateNetworkWebhooksHttpServer(ctx, networkId).CreateNetworkWebhooksHttpServerRequest(createNetworkWebhooksHttpServerRequest).Execute()

Add an HTTP server to a network



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
	createNetworkWebhooksHttpServerRequest := *openapiclient.NewCreateNetworkWebhooksHttpServerRequest("Name_example", "Url_example") // CreateNetworkWebhooksHttpServerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkWebhooksHttpServer(context.Background(), networkId).CreateNetworkWebhooksHttpServerRequest(createNetworkWebhooksHttpServerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkWebhooksHttpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkWebhooksHttpServer`: GetNetworkWebhooksHttpServers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkWebhooksHttpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWebhooksHttpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkWebhooksHttpServerRequest** | [**CreateNetworkWebhooksHttpServerRequest**](CreateNetworkWebhooksHttpServerRequest.md) |  | 

### Return type

[**GetNetworkWebhooksHttpServers200ResponseInner**](GetNetworkWebhooksHttpServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkWebhooksPayloadTemplate

> GetNetworkWebhooksPayloadTemplates200ResponseInner CreateNetworkWebhooksPayloadTemplate(ctx, networkId).CreateNetworkWebhooksPayloadTemplateRequest(createNetworkWebhooksPayloadTemplateRequest).Execute()

Create a webhook payload template for a network



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
	createNetworkWebhooksPayloadTemplateRequest := *openapiclient.NewCreateNetworkWebhooksPayloadTemplateRequest("Name_example") // CreateNetworkWebhooksPayloadTemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkWebhooksPayloadTemplate(context.Background(), networkId).CreateNetworkWebhooksPayloadTemplateRequest(createNetworkWebhooksPayloadTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkWebhooksPayloadTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkWebhooksPayloadTemplate`: GetNetworkWebhooksPayloadTemplates200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkWebhooksPayloadTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWebhooksPayloadTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkWebhooksPayloadTemplateRequest** | [**CreateNetworkWebhooksPayloadTemplateRequest**](CreateNetworkWebhooksPayloadTemplateRequest.md) |  | 

### Return type

[**GetNetworkWebhooksPayloadTemplates200ResponseInner**](GetNetworkWebhooksPayloadTemplates200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkWebhooksWebhookTest

> CreateNetworkWebhooksWebhookTest201Response CreateNetworkWebhooksWebhookTest(ctx, networkId).CreateNetworkWebhooksWebhookTestRequest(createNetworkWebhooksWebhookTestRequest).Execute()

Send a test webhook for a network



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
	createNetworkWebhooksWebhookTestRequest := *openapiclient.NewCreateNetworkWebhooksWebhookTestRequest("Url_example") // CreateNetworkWebhooksWebhookTestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateNetworkWebhooksWebhookTest(context.Background(), networkId).CreateNetworkWebhooksWebhookTestRequest(createNetworkWebhooksWebhookTestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateNetworkWebhooksWebhookTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkWebhooksWebhookTest`: CreateNetworkWebhooksWebhookTest201Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateNetworkWebhooksWebhookTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWebhooksWebhookTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkWebhooksWebhookTestRequest** | [**CreateNetworkWebhooksWebhookTestRequest**](CreateNetworkWebhooksWebhookTestRequest.md) |  | 

### Return type

[**CreateNetworkWebhooksWebhookTest201Response**](CreateNetworkWebhooksWebhookTest201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationNetwork

> GetNetwork200Response CreateOrganizationNetwork(ctx, organizationId).CreateOrganizationNetworkRequest(createOrganizationNetworkRequest).Execute()

Create a network



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
	createOrganizationNetworkRequest := *openapiclient.NewCreateOrganizationNetworkRequest("Name_example", []string{"ProductTypes_example"}) // CreateOrganizationNetworkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateOrganizationNetwork(context.Background(), organizationId).CreateOrganizationNetworkRequest(createOrganizationNetworkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateOrganizationNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationNetwork`: GetNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateOrganizationNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationNetworkRequest** | [**CreateOrganizationNetworkRequest**](CreateOrganizationNetworkRequest.md) |  | 

### Return type

[**GetNetwork200Response**](GetNetwork200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeferNetworkFirmwareUpgradesStagedEvents

> GetNetworkFirmwareUpgradesStagedEvents200Response DeferNetworkFirmwareUpgradesStagedEvents(ctx, networkId).Execute()

Postpone by 1 week all pending staged upgrade stages for a network



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
	resp, r, err := apiClient.NetworksAPI.DeferNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeferNetworkFirmwareUpgradesStagedEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeferNetworkFirmwareUpgradesStagedEvents`: GetNetworkFirmwareUpgradesStagedEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeferNetworkFirmwareUpgradesStagedEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeferNetworkFirmwareUpgradesStagedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkFirmwareUpgradesStagedEvents200Response**](GetNetworkFirmwareUpgradesStagedEvents200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetwork

> DeleteNetwork(ctx, networkId).Execute()

Delete a network



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
	r, err := apiClient.NetworksAPI.DeleteNetwork(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRequest struct via the builder pattern


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


## DeleteNetworkFirmwareUpgradesStagedGroup

> DeleteNetworkFirmwareUpgradesStagedGroup(ctx, networkId, groupId).Execute()

Delete a Staged Upgrade Group



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
	groupId := "groupId_example" // string | Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworksAPI.DeleteNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkFirmwareUpgradesStagedGroupRequest struct via the builder pattern


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


## DeleteNetworkFloorPlan

> GetNetworkFloorPlans200ResponseInner DeleteNetworkFloorPlan(ctx, networkId, floorPlanId).Execute()

Destroy a floor plan



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
	floorPlanId := "floorPlanId_example" // string | Floor plan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteNetworkFloorPlan(context.Background(), networkId, floorPlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkFloorPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNetworkFloorPlan`: GetNetworkFloorPlans200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteNetworkFloorPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**floorPlanId** | **string** | Floor plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkFloorPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkFloorPlans200ResponseInner**](GetNetworkFloorPlans200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkGroupPolicy

> DeleteNetworkGroupPolicy(ctx, networkId, groupPolicyId).Force(force).Execute()

Delete a group policy



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
	groupPolicyId := "groupPolicyId_example" // string | Group policy ID
	force := true // bool | If true, the system deletes the GP even if there are active clients using the GP. After deletion, active clients that were assigned to that Group Policy will be left without any policy applied. Default is false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworksAPI.DeleteNetworkGroupPolicy(context.Background(), networkId, groupPolicyId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkGroupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupPolicyId** | **string** | Group policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** | If true, the system deletes the GP even if there are active clients using the GP. After deletion, active clients that were assigned to that Group Policy will be left without any policy applied. Default is false. | 

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


## DeleteNetworkMerakiAuthUser

> DeleteNetworkMerakiAuthUser(ctx, networkId, merakiAuthUserId).Delete(delete).Execute()

Delete an 802.1X RADIUS user, or deauthorize and optionally delete a splash guest or client VPN user.



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
	merakiAuthUserId := "merakiAuthUserId_example" // string | Meraki auth user ID
	delete := true // bool | If the ID supplied is for a splash guest or client VPN user, and that user is not authorized for any other networks in the organization, then also delete the user. 802.1X RADIUS users are always deleted regardless of this optional attribute. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworksAPI.DeleteNetworkMerakiAuthUser(context.Background(), networkId, merakiAuthUserId).Delete(delete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkMerakiAuthUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**merakiAuthUserId** | **string** | Meraki auth user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkMerakiAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **delete** | **bool** | If the ID supplied is for a splash guest or client VPN user, and that user is not authorized for any other networks in the organization, then also delete the user. 802.1X RADIUS users are always deleted regardless of this optional attribute. | 

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


## DeleteNetworkMqttBroker

> DeleteNetworkMqttBroker(ctx, networkId, mqttBrokerId).Execute()

Delete an MQTT broker



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
	mqttBrokerId := "mqttBrokerId_example" // string | Mqtt broker ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworksAPI.DeleteNetworkMqttBroker(context.Background(), networkId, mqttBrokerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkMqttBroker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**mqttBrokerId** | **string** | Mqtt broker ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkMqttBrokerRequest struct via the builder pattern


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


## DeleteNetworkPiiRequest

> DeleteNetworkPiiRequest(ctx, networkId, requestId).Execute()

Delete a restrict processing PII request



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
	requestId := "requestId_example" // string | Request ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworksAPI.DeleteNetworkPiiRequest(context.Background(), networkId, requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkPiiRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkPiiRequestRequest struct via the builder pattern


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


## DeleteNetworkVlanProfile

> DeleteNetworkVlanProfile(ctx, networkId, iname).Execute()

Delete a VLAN profile of a network



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
	iname := "iname_example" // string | Iname

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworksAPI.DeleteNetworkVlanProfile(context.Background(), networkId, iname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkVlanProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**iname** | **string** | Iname | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkVlanProfileRequest struct via the builder pattern


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


## DeleteNetworkWebhooksHttpServer

> DeleteNetworkWebhooksHttpServer(ctx, networkId, httpServerId).Execute()

Delete an HTTP server from a network



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
	httpServerId := "httpServerId_example" // string | Http server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworksAPI.DeleteNetworkWebhooksHttpServer(context.Background(), networkId, httpServerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkWebhooksHttpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**httpServerId** | **string** | Http server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkWebhooksHttpServerRequest struct via the builder pattern


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


## DeleteNetworkWebhooksPayloadTemplate

> DeleteNetworkWebhooksPayloadTemplate(ctx, networkId, payloadTemplateId).Execute()

Destroy a webhook payload template for a network



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
	payloadTemplateId := "payloadTemplateId_example" // string | Payload template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworksAPI.DeleteNetworkWebhooksPayloadTemplate(context.Background(), networkId, payloadTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteNetworkWebhooksPayloadTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**payloadTemplateId** | **string** | Payload template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkWebhooksPayloadTemplateRequest struct via the builder pattern


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


## GetNetwork

> GetNetwork200Response GetNetwork(ctx, networkId).Execute()

Return a network



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
	resp, r, err := apiClient.NetworksAPI.GetNetwork(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetwork`: GetNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetwork200Response**](GetNetwork200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAlertsHistory

> []GetNetworkAlertsHistory200ResponseInner GetNetworkAlertsHistory(ctx, networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return the alert history for this network



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkAlertsHistory(context.Background(), networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkAlertsHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkAlertsHistory`: []GetNetworkAlertsHistory200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkAlertsHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAlertsHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkAlertsHistory200ResponseInner**](GetNetworkAlertsHistory200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAlertsSettings

> GetNetworkAlertsSettings200Response GetNetworkAlertsSettings(ctx, networkId).Execute()

Return the alert configuration for this network



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkAlertsSettings(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkAlertsSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkAlertsSettings`: GetNetworkAlertsSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkAlertsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAlertsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkAlertsSettings200Response**](GetNetworkAlertsSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkBluetoothClient

> GetNetworkBluetoothClients200ResponseInner GetNetworkBluetoothClient(ctx, networkId, bluetoothClientId).IncludeConnectivityHistory(includeConnectivityHistory).ConnectivityHistoryTimespan(connectivityHistoryTimespan).Execute()

Return a Bluetooth client



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
	bluetoothClientId := "bluetoothClientId_example" // string | Bluetooth client ID
	includeConnectivityHistory := true // bool | Include the connectivity history for this client (optional)
	connectivityHistoryTimespan := int32(56) // int32 | The timespan, in seconds, for the connectivityHistory data. By default 1 day, 86400, will be used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkBluetoothClient(context.Background(), networkId, bluetoothClientId).IncludeConnectivityHistory(includeConnectivityHistory).ConnectivityHistoryTimespan(connectivityHistoryTimespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkBluetoothClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkBluetoothClient`: GetNetworkBluetoothClients200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkBluetoothClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**bluetoothClientId** | **string** | Bluetooth client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkBluetoothClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeConnectivityHistory** | **bool** | Include the connectivity history for this client | 
 **connectivityHistoryTimespan** | **int32** | The timespan, in seconds, for the connectivityHistory data. By default 1 day, 86400, will be used. | 

### Return type

[**GetNetworkBluetoothClients200ResponseInner**](GetNetworkBluetoothClients200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkBluetoothClients

> []GetNetworkBluetoothClients200ResponseInner GetNetworkBluetoothClients(ctx, networkId).T0(t0).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).IncludeConnectivityHistory(includeConnectivityHistory).Execute()

List the Bluetooth clients seen by APs in this network



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
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 7 days from today. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 7 days. The default is 1 day. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 5 - 1000. Default is 10. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	includeConnectivityHistory := true // bool | Include the connectivity history for this client (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkBluetoothClients(context.Background(), networkId).T0(t0).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).IncludeConnectivityHistory(includeConnectivityHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkBluetoothClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkBluetoothClients`: []GetNetworkBluetoothClients200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkBluetoothClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkBluetoothClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 7 days from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 7 days. The default is 1 day. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 5 - 1000. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **includeConnectivityHistory** | **bool** | Include the connectivity history for this client | 

### Return type

[**[]GetNetworkBluetoothClients200ResponseInner**](GetNetworkBluetoothClients200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkClient

> GetNetworkClient200Response GetNetworkClient(ctx, networkId, clientId).Execute()

Return the client associated with the given identifier



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
	clientId := "clientId_example" // string | Client ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkClient(context.Background(), networkId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkClient`: GetNetworkClient200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkClient200Response**](GetNetworkClient200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkClientPolicy

> GetNetworkClientPolicy200Response GetNetworkClientPolicy(ctx, networkId, clientId).Execute()

Return the policy assigned to a client on the network



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
	clientId := "clientId_example" // string | Client ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkClientPolicy(context.Background(), networkId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkClientPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkClientPolicy`: GetNetworkClientPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkClientPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkClientPolicy200Response**](GetNetworkClientPolicy200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkClientSplashAuthorizationStatus

> map[string]interface{} GetNetworkClientSplashAuthorizationStatus(ctx, networkId, clientId).Execute()

Return the splash authorization for a client, for each SSID they've associated with through splash



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
	clientId := "clientId_example" // string | Client ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkClientSplashAuthorizationStatus(context.Background(), networkId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkClientSplashAuthorizationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkClientSplashAuthorizationStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkClientSplashAuthorizationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientSplashAuthorizationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkClientTrafficHistory

> []GetNetworkClientTrafficHistory200ResponseInner GetNetworkClientTrafficHistory(ctx, networkId, clientId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return the client's network traffic data over time



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
	clientId := "clientId_example" // string | Client ID
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkClientTrafficHistory(context.Background(), networkId, clientId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkClientTrafficHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkClientTrafficHistory`: []GetNetworkClientTrafficHistory200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkClientTrafficHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientTrafficHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkClientTrafficHistory200ResponseInner**](GetNetworkClientTrafficHistory200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkClientUsageHistory

> []GetNetworkClientUsageHistory200ResponseInner GetNetworkClientUsageHistory(ctx, networkId, clientId).Execute()

Return the client's daily usage history



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
	clientId := "clientId_example" // string | Client ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkClientUsageHistory(context.Background(), networkId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkClientUsageHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkClientUsageHistory`: []GetNetworkClientUsageHistory200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkClientUsageHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientUsageHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkClientUsageHistory200ResponseInner**](GetNetworkClientUsageHistory200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkClients

> []GetNetworkClients200ResponseInner GetNetworkClients(ctx, networkId).T0(t0).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Statuses(statuses).Ip(ip).Ip6(ip6).Ip6Local(ip6Local).Mac(mac).Os(os).PskGroup(pskGroup).Description(description).Vlan(vlan).NamedVlan(namedVlan).RecentDeviceConnections(recentDeviceConnections).Execute()

List the clients that have used this network in the timespan



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 5000. Default is 10. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	statuses := []string{"Statuses_example"} // []string | Filters clients based on status. Can be one of 'Online' or 'Offline'. (optional)
	ip := "ip_example" // string | Filters clients based on a partial or full match for the ip address field. (optional)
	ip6 := "ip6_example" // string | Filters clients based on a partial or full match for the ip6 address field. (optional)
	ip6Local := "ip6Local_example" // string | Filters clients based on a partial or full match for the ip6Local address field. (optional)
	mac := "mac_example" // string | Filters clients based on a partial or full match for the mac address field. (optional)
	os := "os_example" // string | Filters clients based on a partial or full match for the os (operating system) field. (optional)
	pskGroup := "pskGroup_example" // string | Filters clients based on partial or full match for the iPSK name field. (optional)
	description := "description_example" // string | Filters clients based on a partial or full match for the description field. (optional)
	vlan := "vlan_example" // string | Filters clients based on the full match for the VLAN field. (optional)
	namedVlan := "namedVlan_example" // string | Filters clients based on the partial or full match for the named VLAN field. (optional)
	recentDeviceConnections := []string{"RecentDeviceConnections_example"} // []string | Filters clients based on recent connection type. Can be one of 'Wired' or 'Wireless'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkClients(context.Background(), networkId).T0(t0).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Statuses(statuses).Ip(ip).Ip6(ip6).Ip6Local(ip6Local).Mac(mac).Os(os).PskGroup(pskGroup).Description(description).Vlan(vlan).NamedVlan(namedVlan).RecentDeviceConnections(recentDeviceConnections).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkClients`: []GetNetworkClients200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 5000. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **statuses** | **[]string** | Filters clients based on status. Can be one of &#39;Online&#39; or &#39;Offline&#39;. | 
 **ip** | **string** | Filters clients based on a partial or full match for the ip address field. | 
 **ip6** | **string** | Filters clients based on a partial or full match for the ip6 address field. | 
 **ip6Local** | **string** | Filters clients based on a partial or full match for the ip6Local address field. | 
 **mac** | **string** | Filters clients based on a partial or full match for the mac address field. | 
 **os** | **string** | Filters clients based on a partial or full match for the os (operating system) field. | 
 **pskGroup** | **string** | Filters clients based on partial or full match for the iPSK name field. | 
 **description** | **string** | Filters clients based on a partial or full match for the description field. | 
 **vlan** | **string** | Filters clients based on the full match for the VLAN field. | 
 **namedVlan** | **string** | Filters clients based on the partial or full match for the named VLAN field. | 
 **recentDeviceConnections** | **[]string** | Filters clients based on recent connection type. Can be one of &#39;Wired&#39; or &#39;Wireless&#39;. | 

### Return type

[**[]GetNetworkClients200ResponseInner**](GetNetworkClients200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkClientsApplicationUsage

> []map[string]interface{} GetNetworkClientsApplicationUsage(ctx, networkId).Clients(clients).SsidNumber(ssidNumber).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()

Return the application usage data for clients



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
	clients := "clients_example" // string | A list of client keys, MACs or IPs separated by comma.
	ssidNumber := int32(56) // int32 | An SSID number to include. If not specified, eveusage histories application usagents for all SSIDs will be returned. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkClientsApplicationUsage(context.Background(), networkId).Clients(clients).SsidNumber(ssidNumber).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkClientsApplicationUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkClientsApplicationUsage`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkClientsApplicationUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientsApplicationUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clients** | **string** | A list of client keys, MACs or IPs separated by comma. | 
 **ssidNumber** | **int32** | An SSID number to include. If not specified, eveusage histories application usagents for all SSIDs will be returned. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkClientsBandwidthUsageHistory

> []GetNetworkClientsBandwidthUsageHistory200ResponseInner GetNetworkClientsBandwidthUsageHistory(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Returns a timeseries of total traffic consumption rates for all clients on a network within a given timespan, in megabits per second.



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
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkClientsBandwidthUsageHistory(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkClientsBandwidthUsageHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkClientsBandwidthUsageHistory`: []GetNetworkClientsBandwidthUsageHistory200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkClientsBandwidthUsageHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientsBandwidthUsageHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkClientsBandwidthUsageHistory200ResponseInner**](GetNetworkClientsBandwidthUsageHistory200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkClientsOverview

> GetNetworkClientsOverview200Response GetNetworkClientsOverview(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Execute()

Return overview statistics for network clients



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
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)
	resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 7200, 86400, 604800, 2592000. The default is 604800. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkClientsOverview(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkClientsOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkClientsOverview`: GetNetworkClientsOverview200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkClientsOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 7200, 86400, 604800, 2592000. The default is 604800. | 

### Return type

[**GetNetworkClientsOverview200Response**](GetNetworkClientsOverview200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkClientsUsageHistories

> []map[string]interface{} GetNetworkClientsUsageHistories(ctx, networkId).Clients(clients).SsidNumber(ssidNumber).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()

Return the usage histories for clients



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
	clients := "clients_example" // string | A list of client keys, MACs or IPs separated by comma.
	ssidNumber := int32(56) // int32 | An SSID number to include. If not specified, events for all SSIDs will be returned. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkClientsUsageHistories(context.Background(), networkId).Clients(clients).SsidNumber(ssidNumber).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkClientsUsageHistories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkClientsUsageHistories`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkClientsUsageHistories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientsUsageHistoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clients** | **string** | A list of client keys, MACs or IPs separated by comma. | 
 **ssidNumber** | **int32** | An SSID number to include. If not specified, events for all SSIDs will be returned. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkDevices

> []GetDevice200Response GetNetworkDevices(ctx, networkId).Execute()

List the devices in a network



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkDevices(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDevices`: []GetDevice200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetDevice200Response**](GetDevice200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkEvents

> GetNetworkEvents200Response GetNetworkEvents(ctx, networkId).ProductType(productType).IncludedEventTypes(includedEventTypes).ExcludedEventTypes(excludedEventTypes).DeviceMac(deviceMac).DeviceSerial(deviceSerial).DeviceName(deviceName).ClientIp(clientIp).ClientMac(clientMac).ClientName(clientName).SmDeviceMac(smDeviceMac).SmDeviceName(smDeviceName).EventDetails(eventDetails).EventSeverity(eventSeverity).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the events for the network



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
	productType := "productType_example" // string | The product type to fetch events for. This parameter is required for networks with multiple device types. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, wirelessController, and secureConnect (optional)
	includedEventTypes := []string{"Inner_example"} // []string | A list of event types. The returned events will be filtered to only include events with these types. (optional)
	excludedEventTypes := []string{"Inner_example"} // []string | A list of event types. The returned events will be filtered to exclude events with these types. (optional)
	deviceMac := "deviceMac_example" // string | The MAC address of the Meraki device which the list of events will be filtered with (optional)
	deviceSerial := "deviceSerial_example" // string | The serial of the Meraki device which the list of events will be filtered with (optional)
	deviceName := "deviceName_example" // string | The name of the Meraki device which the list of events will be filtered with (optional)
	clientIp := "clientIp_example" // string | The IP of the client which the list of events will be filtered with. Only supported for track-by-IP networks. (optional)
	clientMac := "clientMac_example" // string | The MAC address of the client which the list of events will be filtered with. Only supported for track-by-MAC networks. (optional)
	clientName := "clientName_example" // string | The name, or partial name, of the client which the list of events will be filtered with (optional)
	smDeviceMac := "smDeviceMac_example" // string | The MAC address of the Systems Manager device which the list of events will be filtered with (optional)
	smDeviceName := "smDeviceName_example" // string | The name of the Systems Manager device which the list of events will be filtered with (optional)
	eventDetails := "eventDetails_example" // string | The details of the event(Catalyst device only) which the list of events will be filtered with (optional)
	eventSeverity := "eventSeverity_example" // string | The severity of the event(Catalyst device only) which the list of events will be filtered with (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkEvents(context.Background(), networkId).ProductType(productType).IncludedEventTypes(includedEventTypes).ExcludedEventTypes(excludedEventTypes).DeviceMac(deviceMac).DeviceSerial(deviceSerial).DeviceName(deviceName).ClientIp(clientIp).ClientMac(clientMac).ClientName(clientName).SmDeviceMac(smDeviceMac).SmDeviceName(smDeviceName).EventDetails(eventDetails).EventSeverity(eventSeverity).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEvents`: GetNetworkEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productType** | **string** | The product type to fetch events for. This parameter is required for networks with multiple device types. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, wirelessController, and secureConnect | 
 **includedEventTypes** | **[]string** | A list of event types. The returned events will be filtered to only include events with these types. | 
 **excludedEventTypes** | **[]string** | A list of event types. The returned events will be filtered to exclude events with these types. | 
 **deviceMac** | **string** | The MAC address of the Meraki device which the list of events will be filtered with | 
 **deviceSerial** | **string** | The serial of the Meraki device which the list of events will be filtered with | 
 **deviceName** | **string** | The name of the Meraki device which the list of events will be filtered with | 
 **clientIp** | **string** | The IP of the client which the list of events will be filtered with. Only supported for track-by-IP networks. | 
 **clientMac** | **string** | The MAC address of the client which the list of events will be filtered with. Only supported for track-by-MAC networks. | 
 **clientName** | **string** | The name, or partial name, of the client which the list of events will be filtered with | 
 **smDeviceMac** | **string** | The MAC address of the Systems Manager device which the list of events will be filtered with | 
 **smDeviceName** | **string** | The name of the Systems Manager device which the list of events will be filtered with | 
 **eventDetails** | **string** | The details of the event(Catalyst device only) which the list of events will be filtered with | 
 **eventSeverity** | **string** | The severity of the event(Catalyst device only) which the list of events will be filtered with | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**GetNetworkEvents200Response**](GetNetworkEvents200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkEventsEventTypes

> []GetNetworkEventsEventTypes200ResponseInner GetNetworkEventsEventTypes(ctx, networkId).Execute()

List the event type to human-readable description



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkEventsEventTypes(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkEventsEventTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEventsEventTypes`: []GetNetworkEventsEventTypes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkEventsEventTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEventsEventTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkEventsEventTypes200ResponseInner**](GetNetworkEventsEventTypes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirmwareUpgrades

> GetNetworkFirmwareUpgrades200Response GetNetworkFirmwareUpgrades(ctx, networkId).Execute()

Get firmware upgrade information for a network



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkFirmwareUpgrades(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkFirmwareUpgrades``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkFirmwareUpgrades`: GetNetworkFirmwareUpgrades200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkFirmwareUpgrades`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirmwareUpgradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkFirmwareUpgrades200Response**](GetNetworkFirmwareUpgrades200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirmwareUpgradesStagedEvents

> GetNetworkFirmwareUpgradesStagedEvents200Response GetNetworkFirmwareUpgradesStagedEvents(ctx, networkId).Execute()

Get the Staged Upgrade Event from a network



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkFirmwareUpgradesStagedEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkFirmwareUpgradesStagedEvents`: GetNetworkFirmwareUpgradesStagedEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkFirmwareUpgradesStagedEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirmwareUpgradesStagedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkFirmwareUpgradesStagedEvents200Response**](GetNetworkFirmwareUpgradesStagedEvents200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirmwareUpgradesStagedGroup

> GetNetworkFirmwareUpgradesStagedGroups200ResponseInner GetNetworkFirmwareUpgradesStagedGroup(ctx, networkId, groupId).Execute()

Get a Staged Upgrade Group from a network



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
	groupId := "groupId_example" // string | Group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkFirmwareUpgradesStagedGroup`: GetNetworkFirmwareUpgradesStagedGroups200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkFirmwareUpgradesStagedGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirmwareUpgradesStagedGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkFirmwareUpgradesStagedGroups200ResponseInner**](GetNetworkFirmwareUpgradesStagedGroups200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirmwareUpgradesStagedGroups

> []GetNetworkFirmwareUpgradesStagedGroups200ResponseInner GetNetworkFirmwareUpgradesStagedGroups(ctx, networkId).Execute()

List of Staged Upgrade Groups in a network



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkFirmwareUpgradesStagedGroups(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkFirmwareUpgradesStagedGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkFirmwareUpgradesStagedGroups`: []GetNetworkFirmwareUpgradesStagedGroups200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkFirmwareUpgradesStagedGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirmwareUpgradesStagedGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkFirmwareUpgradesStagedGroups200ResponseInner**](GetNetworkFirmwareUpgradesStagedGroups200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirmwareUpgradesStagedStages

> []GetNetworkFirmwareUpgradesStagedStages200ResponseInner GetNetworkFirmwareUpgradesStagedStages(ctx, networkId).Execute()

Order of Staged Upgrade Groups in a network



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkFirmwareUpgradesStagedStages(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkFirmwareUpgradesStagedStages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkFirmwareUpgradesStagedStages`: []GetNetworkFirmwareUpgradesStagedStages200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkFirmwareUpgradesStagedStages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirmwareUpgradesStagedStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkFirmwareUpgradesStagedStages200ResponseInner**](GetNetworkFirmwareUpgradesStagedStages200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFloorPlan

> GetNetworkFloorPlans200ResponseInner GetNetworkFloorPlan(ctx, networkId, floorPlanId).Execute()

Find a floor plan by ID



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
	floorPlanId := "floorPlanId_example" // string | Floor plan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkFloorPlan(context.Background(), networkId, floorPlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkFloorPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkFloorPlan`: GetNetworkFloorPlans200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkFloorPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**floorPlanId** | **string** | Floor plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFloorPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkFloorPlans200ResponseInner**](GetNetworkFloorPlans200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFloorPlans

> []GetNetworkFloorPlans200ResponseInner GetNetworkFloorPlans(ctx, networkId).Execute()

List the floor plans that belong to your network



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkFloorPlans(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkFloorPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkFloorPlans`: []GetNetworkFloorPlans200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkFloorPlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFloorPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkFloorPlans200ResponseInner**](GetNetworkFloorPlans200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkGroupPolicies

> []GetNetworkGroupPolicies200ResponseInner GetNetworkGroupPolicies(ctx, networkId).Execute()

List the group policies in a network



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkGroupPolicies(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkGroupPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkGroupPolicies`: []GetNetworkGroupPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkGroupPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkGroupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkGroupPolicies200ResponseInner**](GetNetworkGroupPolicies200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkGroupPolicy

> GetNetworkGroupPolicies200ResponseInner GetNetworkGroupPolicy(ctx, networkId, groupPolicyId).Execute()

Display a group policy



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
	groupPolicyId := "groupPolicyId_example" // string | Group policy ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkGroupPolicy(context.Background(), networkId, groupPolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkGroupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkGroupPolicy`: GetNetworkGroupPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkGroupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupPolicyId** | **string** | Group policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkGroupPolicies200ResponseInner**](GetNetworkGroupPolicies200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkHealthAlerts

> []GetNetworkHealthAlerts200ResponseInner GetNetworkHealthAlerts(ctx, networkId).Execute()

Return all global alerts on this network



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkHealthAlerts(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkHealthAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkHealthAlerts`: []GetNetworkHealthAlerts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkHealthAlerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkHealthAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkHealthAlerts200ResponseInner**](GetNetworkHealthAlerts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkMerakiAuthUser

> GetNetworkMerakiAuthUsers200ResponseInner GetNetworkMerakiAuthUser(ctx, networkId, merakiAuthUserId).Execute()

Return the Meraki Auth splash guest, RADIUS, or client VPN user



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
	merakiAuthUserId := "merakiAuthUserId_example" // string | Meraki auth user ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkMerakiAuthUser(context.Background(), networkId, merakiAuthUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkMerakiAuthUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkMerakiAuthUser`: GetNetworkMerakiAuthUsers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkMerakiAuthUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**merakiAuthUserId** | **string** | Meraki auth user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkMerakiAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkMerakiAuthUsers200ResponseInner**](GetNetworkMerakiAuthUsers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkMerakiAuthUsers

> []GetNetworkMerakiAuthUsers200ResponseInner GetNetworkMerakiAuthUsers(ctx, networkId).Execute()

List the users configured under Meraki Authentication for a network (splash guest or RADIUS users for a wireless network, or client VPN users for a MX network)



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkMerakiAuthUsers(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkMerakiAuthUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkMerakiAuthUsers`: []GetNetworkMerakiAuthUsers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkMerakiAuthUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkMerakiAuthUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkMerakiAuthUsers200ResponseInner**](GetNetworkMerakiAuthUsers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkMqttBroker

> GetNetworkMqttBrokers200ResponseInner GetNetworkMqttBroker(ctx, networkId, mqttBrokerId).Execute()

Return an MQTT broker



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
	mqttBrokerId := "mqttBrokerId_example" // string | Mqtt broker ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkMqttBroker(context.Background(), networkId, mqttBrokerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkMqttBroker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkMqttBroker`: GetNetworkMqttBrokers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkMqttBroker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**mqttBrokerId** | **string** | Mqtt broker ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkMqttBrokerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkMqttBrokers200ResponseInner**](GetNetworkMqttBrokers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkMqttBrokers

> []GetNetworkMqttBrokers200ResponseInner GetNetworkMqttBrokers(ctx, networkId).Execute()

List the MQTT brokers for this network



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkMqttBrokers(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkMqttBrokers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkMqttBrokers`: []GetNetworkMqttBrokers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkMqttBrokers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkMqttBrokersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkMqttBrokers200ResponseInner**](GetNetworkMqttBrokers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkNetflow

> GetNetworkNetflow200Response GetNetworkNetflow(ctx, networkId).Execute()

Return the NetFlow traffic reporting settings for a network



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkNetflow(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkNetflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkNetflow`: GetNetworkNetflow200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkNetflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkNetflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkNetflow200Response**](GetNetworkNetflow200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkNetworkHealthChannelUtilization

> []GetNetworkNetworkHealthChannelUtilization200ResponseInner GetNetworkNetworkHealthChannelUtilization(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Get the channel utilization over each radio for all APs in a network.



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
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)
	resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 600. The default is 600. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 10. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkNetworkHealthChannelUtilization(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkNetworkHealthChannelUtilization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkNetworkHealthChannelUtilization`: []GetNetworkNetworkHealthChannelUtilization200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkNetworkHealthChannelUtilization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkNetworkHealthChannelUtilizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 600. The default is 600. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkNetworkHealthChannelUtilization200ResponseInner**](GetNetworkNetworkHealthChannelUtilization200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPiiPiiKeys

> map[string]GetNetworkPiiPiiKeys200ResponseValue GetNetworkPiiPiiKeys(ctx, networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()

List the keys required to access Personally Identifiable Information (PII) for a given identifier



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
	username := "username_example" // string | The username of a Systems Manager user (optional)
	email := "email_example" // string | The email of a network user account or a Systems Manager device (optional)
	mac := "mac_example" // string | The MAC of a network client device or a Systems Manager device (optional)
	serial := "serial_example" // string | The serial of a Systems Manager device (optional)
	imei := "imei_example" // string | The IMEI of a Systems Manager device (optional)
	bluetoothMac := "bluetoothMac_example" // string | The MAC of a Bluetooth client (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkPiiPiiKeys(context.Background(), networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkPiiPiiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPiiPiiKeys`: map[string]GetNetworkPiiPiiKeys200ResponseValue
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkPiiPiiKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiPiiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** | The username of a Systems Manager user | 
 **email** | **string** | The email of a network user account or a Systems Manager device | 
 **mac** | **string** | The MAC of a network client device or a Systems Manager device | 
 **serial** | **string** | The serial of a Systems Manager device | 
 **imei** | **string** | The IMEI of a Systems Manager device | 
 **bluetoothMac** | **string** | The MAC of a Bluetooth client | 

### Return type

[**map[string]GetNetworkPiiPiiKeys200ResponseValue**](GetNetworkPiiPiiKeys200ResponseValue.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPiiRequest

> GetNetworkPiiRequests200ResponseInner GetNetworkPiiRequest(ctx, networkId, requestId).Execute()

Return a PII request



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
	requestId := "requestId_example" // string | Request ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkPiiRequest(context.Background(), networkId, requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkPiiRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPiiRequest`: GetNetworkPiiRequests200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkPiiRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkPiiRequests200ResponseInner**](GetNetworkPiiRequests200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPiiRequests

> []GetNetworkPiiRequests200ResponseInner GetNetworkPiiRequests(ctx, networkId).Execute()

List the PII requests for this network or organization



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkPiiRequests(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkPiiRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPiiRequests`: []GetNetworkPiiRequests200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkPiiRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkPiiRequests200ResponseInner**](GetNetworkPiiRequests200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPiiSmDevicesForKey

> map[string][]string GetNetworkPiiSmDevicesForKey(ctx, networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()

Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier



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
	username := "username_example" // string | The username of a Systems Manager user (optional)
	email := "email_example" // string | The email of a network user account or a Systems Manager device (optional)
	mac := "mac_example" // string | The MAC of a network client device or a Systems Manager device (optional)
	serial := "serial_example" // string | The serial of a Systems Manager device (optional)
	imei := "imei_example" // string | The IMEI of a Systems Manager device (optional)
	bluetoothMac := "bluetoothMac_example" // string | The MAC of a Bluetooth client (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkPiiSmDevicesForKey(context.Background(), networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkPiiSmDevicesForKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPiiSmDevicesForKey`: map[string][]string
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkPiiSmDevicesForKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiSmDevicesForKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** | The username of a Systems Manager user | 
 **email** | **string** | The email of a network user account or a Systems Manager device | 
 **mac** | **string** | The MAC of a network client device or a Systems Manager device | 
 **serial** | **string** | The serial of a Systems Manager device | 
 **imei** | **string** | The IMEI of a Systems Manager device | 
 **bluetoothMac** | **string** | The MAC of a Bluetooth client | 

### Return type

[**map[string][]string**](array.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPiiSmOwnersForKey

> map[string][]string GetNetworkPiiSmOwnersForKey(ctx, networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()

Given a piece of Personally Identifiable Information (PII), return the Systems Manager owner ID(s) associated with that identifier



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
	username := "username_example" // string | The username of a Systems Manager user (optional)
	email := "email_example" // string | The email of a network user account or a Systems Manager device (optional)
	mac := "mac_example" // string | The MAC of a network client device or a Systems Manager device (optional)
	serial := "serial_example" // string | The serial of a Systems Manager device (optional)
	imei := "imei_example" // string | The IMEI of a Systems Manager device (optional)
	bluetoothMac := "bluetoothMac_example" // string | The MAC of a Bluetooth client (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkPiiSmOwnersForKey(context.Background(), networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkPiiSmOwnersForKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPiiSmOwnersForKey`: map[string][]string
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkPiiSmOwnersForKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiSmOwnersForKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** | The username of a Systems Manager user | 
 **email** | **string** | The email of a network user account or a Systems Manager device | 
 **mac** | **string** | The MAC of a network client device or a Systems Manager device | 
 **serial** | **string** | The serial of a Systems Manager device | 
 **imei** | **string** | The IMEI of a Systems Manager device | 
 **bluetoothMac** | **string** | The MAC of a Bluetooth client | 

### Return type

[**map[string][]string**](array.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPoliciesByClient

> []GetNetworkPoliciesByClient200ResponseInner GetNetworkPoliciesByClient(ctx, networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).Timespan(timespan).Execute()

Get policies for all clients with policies



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkPoliciesByClient(context.Background(), networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkPoliciesByClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPoliciesByClient`: []GetNetworkPoliciesByClient200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkPoliciesByClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPoliciesByClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

[**[]GetNetworkPoliciesByClient200ResponseInner**](GetNetworkPoliciesByClient200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSettings

> GetNetworkSettings200Response GetNetworkSettings(ctx, networkId).Execute()

Return the settings for a network



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkSettings(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSettings`: GetNetworkSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSettings200Response**](GetNetworkSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSnmp

> GetNetworkSnmp200Response GetNetworkSnmp(ctx, networkId).Execute()

Return the SNMP settings for a network



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkSnmp(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkSnmp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSnmp`: GetNetworkSnmp200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkSnmp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSnmpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSnmp200Response**](GetNetworkSnmp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSplashLoginAttempts

> []GetNetworkSplashLoginAttempts200ResponseInner GetNetworkSplashLoginAttempts(ctx, networkId).SsidNumber(ssidNumber).LoginIdentifier(loginIdentifier).Timespan(timespan).Execute()

List the splash login attempts for a network



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
	ssidNumber := int32(56) // int32 | Only return the login attempts for the specified SSID (optional)
	loginIdentifier := "loginIdentifier_example" // string | The username, email, or phone number used during login (optional)
	timespan := int32(56) // int32 | The timespan, in seconds, for the login attempts. The period will be from [timespan] seconds ago until now. The maximum timespan is 3 months (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkSplashLoginAttempts(context.Background(), networkId).SsidNumber(ssidNumber).LoginIdentifier(loginIdentifier).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkSplashLoginAttempts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSplashLoginAttempts`: []GetNetworkSplashLoginAttempts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkSplashLoginAttempts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSplashLoginAttemptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ssidNumber** | **int32** | Only return the login attempts for the specified SSID | 
 **loginIdentifier** | **string** | The username, email, or phone number used during login | 
 **timespan** | **int32** | The timespan, in seconds, for the login attempts. The period will be from [timespan] seconds ago until now. The maximum timespan is 3 months | 

### Return type

[**[]GetNetworkSplashLoginAttempts200ResponseInner**](GetNetworkSplashLoginAttempts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSyslogServers

> GetNetworkSyslogServers200Response GetNetworkSyslogServers(ctx, networkId).Execute()

List the syslog servers for a network



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkSyslogServers(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkSyslogServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSyslogServers`: GetNetworkSyslogServers200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkSyslogServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSyslogServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSyslogServers200Response**](GetNetworkSyslogServers200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkTopologyLinkLayer

> map[string]interface{} GetNetworkTopologyLinkLayer(ctx, networkId).Execute()

List the LLDP and CDP information for all discovered devices and connections in a network



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkTopologyLinkLayer(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkTopologyLinkLayer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkTopologyLinkLayer`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkTopologyLinkLayer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTopologyLinkLayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkTraffic

> []map[string]interface{} GetNetworkTraffic(ctx, networkId).T0(t0).Timespan(timespan).DeviceType(deviceType).Execute()

Return the traffic analysis data for this network



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
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 30 days. (optional)
	deviceType := "deviceType_example" // string | Filter the data by device type: 'combined', 'wireless', 'switch' or 'appliance'. Defaults to 'combined'. When using 'combined', for each rule the data will come from the device type with the most usage. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkTraffic(context.Background(), networkId).T0(t0).Timespan(timespan).DeviceType(deviceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkTraffic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkTraffic`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkTraffic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTrafficRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 30 days. | 
 **deviceType** | **string** | Filter the data by device type: &#39;combined&#39;, &#39;wireless&#39;, &#39;switch&#39; or &#39;appliance&#39;. Defaults to &#39;combined&#39;. When using &#39;combined&#39;, for each rule the data will come from the device type with the most usage. | 

### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkTrafficAnalysis

> GetNetworkTrafficAnalysis200Response GetNetworkTrafficAnalysis(ctx, networkId).Execute()

Return the traffic analysis settings for a network



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkTrafficAnalysis(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkTrafficAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkTrafficAnalysis`: GetNetworkTrafficAnalysis200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkTrafficAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTrafficAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkTrafficAnalysis200Response**](GetNetworkTrafficAnalysis200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkTrafficShapingApplicationCategories

> map[string]interface{} GetNetworkTrafficShapingApplicationCategories(ctx, networkId).Execute()

Returns the application categories for traffic shaping rules



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkTrafficShapingApplicationCategories(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkTrafficShapingApplicationCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkTrafficShapingApplicationCategories`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkTrafficShapingApplicationCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTrafficShapingApplicationCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkTrafficShapingDscpTaggingOptions

> []map[string]interface{} GetNetworkTrafficShapingDscpTaggingOptions(ctx, networkId).Execute()

Returns the available DSCP tagging options for your traffic shaping rules.



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkTrafficShapingDscpTaggingOptions(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkTrafficShapingDscpTaggingOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkTrafficShapingDscpTaggingOptions`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkTrafficShapingDscpTaggingOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTrafficShapingDscpTaggingOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkVlanProfile

> GetNetworkVlanProfiles200ResponseInner GetNetworkVlanProfile(ctx, networkId, iname).Execute()

Get an existing VLAN profile of a network



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
	iname := "iname_example" // string | Iname

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkVlanProfile(context.Background(), networkId, iname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkVlanProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkVlanProfile`: GetNetworkVlanProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkVlanProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**iname** | **string** | Iname | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkVlanProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkVlanProfiles200ResponseInner**](GetNetworkVlanProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkVlanProfiles

> []GetNetworkVlanProfiles200ResponseInner GetNetworkVlanProfiles(ctx, networkId).Execute()

List VLAN profiles for a network



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkVlanProfiles(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkVlanProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkVlanProfiles`: []GetNetworkVlanProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkVlanProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkVlanProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkVlanProfiles200ResponseInner**](GetNetworkVlanProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkVlanProfilesAssignmentsByDevice

> []GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner GetNetworkVlanProfilesAssignmentsByDevice(ctx, networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Serials(serials).ProductTypes(productTypes).StackIds(stackIds).Execute()

Get the assigned VLAN Profiles for devices in a network



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
	serials := []string{"Inner_example"} // []string | Optional parameter to filter devices by serials. All devices returned belong to serial numbers that are an exact match. (optional)
	productTypes := []string{"ProductTypes_example"} // []string | Optional parameter to filter devices by product types. (optional)
	stackIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by Switch Stack ids. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkVlanProfilesAssignmentsByDevice(context.Background(), networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Serials(serials).ProductTypes(productTypes).StackIds(stackIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkVlanProfilesAssignmentsByDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkVlanProfilesAssignmentsByDevice`: []GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkVlanProfilesAssignmentsByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkVlanProfilesAssignmentsByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **serials** | **[]string** | Optional parameter to filter devices by serials. All devices returned belong to serial numbers that are an exact match. | 
 **productTypes** | **[]string** | Optional parameter to filter devices by product types. | 
 **stackIds** | **[]string** | Optional parameter to filter devices by Switch Stack ids. | 

### Return type

[**[]GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner**](GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWebhooksHttpServer

> GetNetworkWebhooksHttpServers200ResponseInner GetNetworkWebhooksHttpServer(ctx, networkId, httpServerId).Execute()

Return an HTTP server for a network



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
	httpServerId := "httpServerId_example" // string | Http server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkWebhooksHttpServer(context.Background(), networkId, httpServerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkWebhooksHttpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWebhooksHttpServer`: GetNetworkWebhooksHttpServers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkWebhooksHttpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**httpServerId** | **string** | Http server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksHttpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkWebhooksHttpServers200ResponseInner**](GetNetworkWebhooksHttpServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWebhooksHttpServers

> []GetNetworkWebhooksHttpServers200ResponseInner GetNetworkWebhooksHttpServers(ctx, networkId).Execute()

List the HTTP servers for a network



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkWebhooksHttpServers(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkWebhooksHttpServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWebhooksHttpServers`: []GetNetworkWebhooksHttpServers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkWebhooksHttpServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksHttpServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkWebhooksHttpServers200ResponseInner**](GetNetworkWebhooksHttpServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWebhooksPayloadTemplate

> GetNetworkWebhooksPayloadTemplates200ResponseInner GetNetworkWebhooksPayloadTemplate(ctx, networkId, payloadTemplateId).Execute()

Get the webhook payload template for a network



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
	payloadTemplateId := "payloadTemplateId_example" // string | Payload template ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkWebhooksPayloadTemplate(context.Background(), networkId, payloadTemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkWebhooksPayloadTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWebhooksPayloadTemplate`: GetNetworkWebhooksPayloadTemplates200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkWebhooksPayloadTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**payloadTemplateId** | **string** | Payload template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksPayloadTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkWebhooksPayloadTemplates200ResponseInner**](GetNetworkWebhooksPayloadTemplates200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWebhooksPayloadTemplates

> []GetNetworkWebhooksPayloadTemplates200ResponseInner GetNetworkWebhooksPayloadTemplates(ctx, networkId).Execute()

List the webhook payload templates for a network



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
	resp, r, err := apiClient.NetworksAPI.GetNetworkWebhooksPayloadTemplates(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkWebhooksPayloadTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWebhooksPayloadTemplates`: []GetNetworkWebhooksPayloadTemplates200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkWebhooksPayloadTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksPayloadTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkWebhooksPayloadTemplates200ResponseInner**](GetNetworkWebhooksPayloadTemplates200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWebhooksWebhookTest

> CreateNetworkWebhooksWebhookTest201Response GetNetworkWebhooksWebhookTest(ctx, networkId, webhookTestId).Execute()

Return the status of a webhook test for a network



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
	webhookTestId := "webhookTestId_example" // string | Webhook test ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetNetworkWebhooksWebhookTest(context.Background(), networkId, webhookTestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworkWebhooksWebhookTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWebhooksWebhookTest`: CreateNetworkWebhooksWebhookTest201Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworkWebhooksWebhookTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**webhookTestId** | **string** | Webhook test ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksWebhookTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateNetworkWebhooksWebhookTest201Response**](CreateNetworkWebhooksWebhookTest201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryOnboardingCloudMonitoringNetworks

> []GetNetwork200Response GetOrganizationInventoryOnboardingCloudMonitoringNetworks(ctx, organizationId).DeviceType(deviceType).Search(search).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Returns list of networks eligible for adding cloud monitored device



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
	deviceType := "deviceType_example" // string | Device Type switch or wireless controller
	search := "search_example" // string | Optional parameter to search on network name (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetOrganizationInventoryOnboardingCloudMonitoringNetworks(context.Background(), organizationId).DeviceType(deviceType).Search(search).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetOrganizationInventoryOnboardingCloudMonitoringNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationInventoryOnboardingCloudMonitoringNetworks`: []GetNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetOrganizationInventoryOnboardingCloudMonitoringNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceType** | **string** | Device Type switch or wireless controller | 
 **search** | **string** | Optional parameter to search on network name | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetwork200Response**](GetNetwork200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationNetworks

> []GetNetwork200Response GetOrganizationNetworks(ctx, organizationId).ConfigTemplateId(configTemplateId).IsBoundToConfigTemplate(isBoundToConfigTemplate).Tags(tags).TagsFilterType(tagsFilterType).ProductTypes(productTypes).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the networks that the user has privileges on in an organization



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
	configTemplateId := "configTemplateId_example" // string | An optional parameter that is the ID of a config template. Will return all networks bound to that template. (optional)
	isBoundToConfigTemplate := true // bool | An optional parameter to filter config template bound networks. If configTemplateId is set, this cannot be false. (optional)
	tags := []string{"Inner_example"} // []string | An optional parameter to filter networks by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). (optional)
	tagsFilterType := "tagsFilterType_example" // string | An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return networks which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)
	productTypes := []string{"ProductTypes_example"} // []string | An optional parameter to filter networks by product type. Results will have at least one of the included product types. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetOrganizationNetworks(context.Background(), organizationId).ConfigTemplateId(configTemplateId).IsBoundToConfigTemplate(isBoundToConfigTemplate).Tags(tags).TagsFilterType(tagsFilterType).ProductTypes(productTypes).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetOrganizationNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationNetworks`: []GetNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetOrganizationNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configTemplateId** | **string** | An optional parameter that is the ID of a config template. Will return all networks bound to that template. | 
 **isBoundToConfigTemplate** | **bool** | An optional parameter to filter config template bound networks. If configTemplateId is set, this cannot be false. | 
 **tags** | **[]string** | An optional parameter to filter networks by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). | 
 **tagsFilterType** | **string** | An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return networks which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 
 **productTypes** | **[]string** | An optional parameter to filter networks by product type. Results will have at least one of the included product types. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetwork200Response**](GetNetwork200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummaryTopNetworksByStatus

> []GetOrganizationSummaryTopNetworksByStatus200ResponseInner GetOrganizationSummaryTopNetworksByStatus(ctx, organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the client and status overview information for the networks in an organization



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
	networkTag := "networkTag_example" // string | Match result to an exact network tag (optional)
	deviceTag := "deviceTag_example" // string | Match result to an exact device tag (optional)
	networkId := "networkId_example" // string | Match result to an exact network id (optional)
	quantity := int32(56) // int32 | Set number of desired results to return. Default is 10. (optional)
	ssidName := "ssidName_example" // string | Filter results by ssid name (optional)
	usageUplink := "usageUplink_example" // string | Filter results by usage uplink (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 5000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetOrganizationSummaryTopNetworksByStatus(context.Background(), organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetOrganizationSummaryTopNetworksByStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSummaryTopNetworksByStatus`: []GetOrganizationSummaryTopNetworksByStatus200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetOrganizationSummaryTopNetworksByStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopNetworksByStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkTag** | **string** | Match result to an exact network tag | 
 **deviceTag** | **string** | Match result to an exact device tag | 
 **networkId** | **string** | Match result to an exact network id | 
 **quantity** | **int32** | Set number of desired results to return. Default is 10. | 
 **ssidName** | **string** | Filter results by ssid name | 
 **usageUplink** | **string** | Filter results by usage uplink | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 5000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetOrganizationSummaryTopNetworksByStatus200ResponseInner**](GetOrganizationSummaryTopNetworksByStatus200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisionNetworkClients

> ProvisionNetworkClients201Response ProvisionNetworkClients(ctx, networkId).ProvisionNetworkClientsRequest(provisionNetworkClientsRequest).Execute()

Provisions a client with a name and policy



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
	provisionNetworkClientsRequest := *openapiclient.NewProvisionNetworkClientsRequest([]openapiclient.ProvisionNetworkClientsRequestClientsInner{*openapiclient.NewProvisionNetworkClientsRequestClientsInner("Mac_example")}, "DevicePolicy_example") // ProvisionNetworkClientsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.ProvisionNetworkClients(context.Background(), networkId).ProvisionNetworkClientsRequest(provisionNetworkClientsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.ProvisionNetworkClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvisionNetworkClients`: ProvisionNetworkClients201Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.ProvisionNetworkClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisionNetworkClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **provisionNetworkClientsRequest** | [**ProvisionNetworkClientsRequest**](ProvisionNetworkClientsRequest.md) |  | 

### Return type

[**ProvisionNetworkClients201Response**](ProvisionNetworkClients201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReassignNetworkVlanProfilesAssignments

> ReassignNetworkVlanProfilesAssignments200Response ReassignNetworkVlanProfilesAssignments(ctx, networkId).ReassignNetworkVlanProfilesAssignmentsRequest(reassignNetworkVlanProfilesAssignmentsRequest).Execute()

Update the assigned VLAN Profile for devices in a network



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
	reassignNetworkVlanProfilesAssignmentsRequest := *openapiclient.NewReassignNetworkVlanProfilesAssignmentsRequest([]string{"Serials_example"}, []string{"StackIds_example"}) // ReassignNetworkVlanProfilesAssignmentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.ReassignNetworkVlanProfilesAssignments(context.Background(), networkId).ReassignNetworkVlanProfilesAssignmentsRequest(reassignNetworkVlanProfilesAssignmentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.ReassignNetworkVlanProfilesAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReassignNetworkVlanProfilesAssignments`: ReassignNetworkVlanProfilesAssignments200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.ReassignNetworkVlanProfilesAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReassignNetworkVlanProfilesAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reassignNetworkVlanProfilesAssignmentsRequest** | [**ReassignNetworkVlanProfilesAssignmentsRequest**](ReassignNetworkVlanProfilesAssignmentsRequest.md) |  | 

### Return type

[**ReassignNetworkVlanProfilesAssignments200Response**](ReassignNetworkVlanProfilesAssignments200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveNetworkDevices

> RemoveNetworkDevices(ctx, networkId).RemoveNetworkDevicesRequest(removeNetworkDevicesRequest).Execute()

Remove a single device



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
	removeNetworkDevicesRequest := *openapiclient.NewRemoveNetworkDevicesRequest("Serial_example") // RemoveNetworkDevicesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworksAPI.RemoveNetworkDevices(context.Background(), networkId).RemoveNetworkDevicesRequest(removeNetworkDevicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.RemoveNetworkDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNetworkDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeNetworkDevicesRequest** | [**RemoveNetworkDevicesRequest**](RemoveNetworkDevicesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbacksNetworkFirmwareUpgradesStagedEvents

> GetNetworkFirmwareUpgradesStagedEvents200Response RollbacksNetworkFirmwareUpgradesStagedEvents(ctx, networkId).RollbacksNetworkFirmwareUpgradesStagedEventsRequest(rollbacksNetworkFirmwareUpgradesStagedEventsRequest).Execute()

Rollback a Staged Upgrade Event for a network



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
	rollbacksNetworkFirmwareUpgradesStagedEventsRequest := *openapiclient.NewRollbacksNetworkFirmwareUpgradesStagedEventsRequest([]openapiclient.UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner{*openapiclient.NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner()}) // RollbacksNetworkFirmwareUpgradesStagedEventsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.RollbacksNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).RollbacksNetworkFirmwareUpgradesStagedEventsRequest(rollbacksNetworkFirmwareUpgradesStagedEventsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.RollbacksNetworkFirmwareUpgradesStagedEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RollbacksNetworkFirmwareUpgradesStagedEvents`: GetNetworkFirmwareUpgradesStagedEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.RollbacksNetworkFirmwareUpgradesStagedEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbacksNetworkFirmwareUpgradesStagedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rollbacksNetworkFirmwareUpgradesStagedEventsRequest** | [**RollbacksNetworkFirmwareUpgradesStagedEventsRequest**](RollbacksNetworkFirmwareUpgradesStagedEventsRequest.md) |  | 

### Return type

[**GetNetworkFirmwareUpgradesStagedEvents200Response**](GetNetworkFirmwareUpgradesStagedEvents200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SplitNetwork

> SplitNetwork200Response SplitNetwork(ctx, networkId).Execute()

Split a combined network into individual networks for each type of device



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
	resp, r, err := apiClient.NetworksAPI.SplitNetwork(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.SplitNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SplitNetwork`: SplitNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.SplitNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSplitNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SplitNetwork200Response**](SplitNetwork200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnbindNetwork

> GetNetwork200Response UnbindNetwork(ctx, networkId).UnbindNetworkRequest(unbindNetworkRequest).Execute()

Unbind a network from a template.



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
	unbindNetworkRequest := *openapiclient.NewUnbindNetworkRequest() // UnbindNetworkRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UnbindNetwork(context.Background(), networkId).UnbindNetworkRequest(unbindNetworkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UnbindNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnbindNetwork`: GetNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UnbindNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnbindNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unbindNetworkRequest** | [**UnbindNetworkRequest**](UnbindNetworkRequest.md) |  | 

### Return type

[**GetNetwork200Response**](GetNetwork200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetwork

> GetNetwork200Response UpdateNetwork(ctx, networkId).UpdateNetworkRequest(updateNetworkRequest).Execute()

Update a network



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
	updateNetworkRequest := *openapiclient.NewUpdateNetworkRequest() // UpdateNetworkRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetwork(context.Background(), networkId).UpdateNetworkRequest(updateNetworkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetwork`: GetNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkRequest** | [**UpdateNetworkRequest**](UpdateNetworkRequest.md) |  | 

### Return type

[**GetNetwork200Response**](GetNetwork200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkAlertsSettings

> GetNetworkAlertsSettings200Response UpdateNetworkAlertsSettings(ctx, networkId).UpdateNetworkAlertsSettingsRequest(updateNetworkAlertsSettingsRequest).Execute()

Update the alert configuration for this network



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
	updateNetworkAlertsSettingsRequest := *openapiclient.NewUpdateNetworkAlertsSettingsRequest() // UpdateNetworkAlertsSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkAlertsSettings(context.Background(), networkId).UpdateNetworkAlertsSettingsRequest(updateNetworkAlertsSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkAlertsSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkAlertsSettings`: GetNetworkAlertsSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkAlertsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkAlertsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkAlertsSettingsRequest** | [**UpdateNetworkAlertsSettingsRequest**](UpdateNetworkAlertsSettingsRequest.md) |  | 

### Return type

[**GetNetworkAlertsSettings200Response**](GetNetworkAlertsSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkClientPolicy

> GetNetworkClientPolicy200Response UpdateNetworkClientPolicy(ctx, networkId, clientId).UpdateNetworkClientPolicyRequest(updateNetworkClientPolicyRequest).Execute()

Update the policy assigned to a client on the network



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
	clientId := "clientId_example" // string | Client ID
	updateNetworkClientPolicyRequest := *openapiclient.NewUpdateNetworkClientPolicyRequest("DevicePolicy_example") // UpdateNetworkClientPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkClientPolicy(context.Background(), networkId, clientId).UpdateNetworkClientPolicyRequest(updateNetworkClientPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkClientPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkClientPolicy`: GetNetworkClientPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkClientPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkClientPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkClientPolicyRequest** | [**UpdateNetworkClientPolicyRequest**](UpdateNetworkClientPolicyRequest.md) |  | 

### Return type

[**GetNetworkClientPolicy200Response**](GetNetworkClientPolicy200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkClientSplashAuthorizationStatus

> map[string]interface{} UpdateNetworkClientSplashAuthorizationStatus(ctx, networkId, clientId).UpdateNetworkClientSplashAuthorizationStatusRequest(updateNetworkClientSplashAuthorizationStatusRequest).Execute()

Update a client's splash authorization



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
	clientId := "clientId_example" // string | Client ID
	updateNetworkClientSplashAuthorizationStatusRequest := *openapiclient.NewUpdateNetworkClientSplashAuthorizationStatusRequest(*openapiclient.NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids()) // UpdateNetworkClientSplashAuthorizationStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkClientSplashAuthorizationStatus(context.Background(), networkId, clientId).UpdateNetworkClientSplashAuthorizationStatusRequest(updateNetworkClientSplashAuthorizationStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkClientSplashAuthorizationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkClientSplashAuthorizationStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkClientSplashAuthorizationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkClientSplashAuthorizationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkClientSplashAuthorizationStatusRequest** | [**UpdateNetworkClientSplashAuthorizationStatusRequest**](UpdateNetworkClientSplashAuthorizationStatusRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkFirmwareUpgrades

> GetNetworkFirmwareUpgrades200Response UpdateNetworkFirmwareUpgrades(ctx, networkId).UpdateNetworkFirmwareUpgradesRequest(updateNetworkFirmwareUpgradesRequest).Execute()

Update firmware upgrade information for a network



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
	updateNetworkFirmwareUpgradesRequest := *openapiclient.NewUpdateNetworkFirmwareUpgradesRequest() // UpdateNetworkFirmwareUpgradesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkFirmwareUpgrades(context.Background(), networkId).UpdateNetworkFirmwareUpgradesRequest(updateNetworkFirmwareUpgradesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkFirmwareUpgrades``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkFirmwareUpgrades`: GetNetworkFirmwareUpgrades200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkFirmwareUpgrades`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFirmwareUpgradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkFirmwareUpgradesRequest** | [**UpdateNetworkFirmwareUpgradesRequest**](UpdateNetworkFirmwareUpgradesRequest.md) |  | 

### Return type

[**GetNetworkFirmwareUpgrades200Response**](GetNetworkFirmwareUpgrades200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkFirmwareUpgradesStagedEvents

> GetNetworkFirmwareUpgradesStagedEvents200Response UpdateNetworkFirmwareUpgradesStagedEvents(ctx, networkId).UpdateNetworkFirmwareUpgradesStagedEventsRequest(updateNetworkFirmwareUpgradesStagedEventsRequest).Execute()

Update the Staged Upgrade Event for a network



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
	updateNetworkFirmwareUpgradesStagedEventsRequest := *openapiclient.NewUpdateNetworkFirmwareUpgradesStagedEventsRequest([]openapiclient.UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner{*openapiclient.NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner()}) // UpdateNetworkFirmwareUpgradesStagedEventsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).UpdateNetworkFirmwareUpgradesStagedEventsRequest(updateNetworkFirmwareUpgradesStagedEventsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkFirmwareUpgradesStagedEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkFirmwareUpgradesStagedEvents`: GetNetworkFirmwareUpgradesStagedEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkFirmwareUpgradesStagedEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFirmwareUpgradesStagedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkFirmwareUpgradesStagedEventsRequest** | [**UpdateNetworkFirmwareUpgradesStagedEventsRequest**](UpdateNetworkFirmwareUpgradesStagedEventsRequest.md) |  | 

### Return type

[**GetNetworkFirmwareUpgradesStagedEvents200Response**](GetNetworkFirmwareUpgradesStagedEvents200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkFirmwareUpgradesStagedGroup

> GetNetworkFirmwareUpgradesStagedGroups200ResponseInner UpdateNetworkFirmwareUpgradesStagedGroup(ctx, networkId, groupId).CreateNetworkFirmwareUpgradesStagedGroupRequest(createNetworkFirmwareUpgradesStagedGroupRequest).Execute()

Update a Staged Upgrade Group for a network



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
	groupId := "groupId_example" // string | Group ID
	createNetworkFirmwareUpgradesStagedGroupRequest := *openapiclient.NewCreateNetworkFirmwareUpgradesStagedGroupRequest("Name_example", false) // CreateNetworkFirmwareUpgradesStagedGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).CreateNetworkFirmwareUpgradesStagedGroupRequest(createNetworkFirmwareUpgradesStagedGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkFirmwareUpgradesStagedGroup`: GetNetworkFirmwareUpgradesStagedGroups200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkFirmwareUpgradesStagedGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFirmwareUpgradesStagedGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createNetworkFirmwareUpgradesStagedGroupRequest** | [**CreateNetworkFirmwareUpgradesStagedGroupRequest**](CreateNetworkFirmwareUpgradesStagedGroupRequest.md) |  | 

### Return type

[**GetNetworkFirmwareUpgradesStagedGroups200ResponseInner**](GetNetworkFirmwareUpgradesStagedGroups200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkFirmwareUpgradesStagedStages

> []GetNetworkFirmwareUpgradesStagedStages200ResponseInner UpdateNetworkFirmwareUpgradesStagedStages(ctx, networkId).UpdateNetworkFirmwareUpgradesStagedStagesRequest(updateNetworkFirmwareUpgradesStagedStagesRequest).Execute()

Assign Staged Upgrade Group order in the sequence.



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
	updateNetworkFirmwareUpgradesStagedStagesRequest := *openapiclient.NewUpdateNetworkFirmwareUpgradesStagedStagesRequest() // UpdateNetworkFirmwareUpgradesStagedStagesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkFirmwareUpgradesStagedStages(context.Background(), networkId).UpdateNetworkFirmwareUpgradesStagedStagesRequest(updateNetworkFirmwareUpgradesStagedStagesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkFirmwareUpgradesStagedStages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkFirmwareUpgradesStagedStages`: []GetNetworkFirmwareUpgradesStagedStages200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkFirmwareUpgradesStagedStages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFirmwareUpgradesStagedStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkFirmwareUpgradesStagedStagesRequest** | [**UpdateNetworkFirmwareUpgradesStagedStagesRequest**](UpdateNetworkFirmwareUpgradesStagedStagesRequest.md) |  | 

### Return type

[**[]GetNetworkFirmwareUpgradesStagedStages200ResponseInner**](GetNetworkFirmwareUpgradesStagedStages200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkFloorPlan

> GetNetworkFloorPlans200ResponseInner UpdateNetworkFloorPlan(ctx, networkId, floorPlanId).UpdateNetworkFloorPlanRequest(updateNetworkFloorPlanRequest).Execute()

Update a floor plan's geolocation and other meta data



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
	floorPlanId := "floorPlanId_example" // string | Floor plan ID
	updateNetworkFloorPlanRequest := *openapiclient.NewUpdateNetworkFloorPlanRequest() // UpdateNetworkFloorPlanRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkFloorPlan(context.Background(), networkId, floorPlanId).UpdateNetworkFloorPlanRequest(updateNetworkFloorPlanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkFloorPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkFloorPlan`: GetNetworkFloorPlans200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkFloorPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**floorPlanId** | **string** | Floor plan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFloorPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkFloorPlanRequest** | [**UpdateNetworkFloorPlanRequest**](UpdateNetworkFloorPlanRequest.md) |  | 

### Return type

[**GetNetworkFloorPlans200ResponseInner**](GetNetworkFloorPlans200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkGroupPolicy

> GetNetworkGroupPolicies200ResponseInner UpdateNetworkGroupPolicy(ctx, networkId, groupPolicyId).UpdateNetworkGroupPolicyRequest(updateNetworkGroupPolicyRequest).Execute()

Update a group policy



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
	groupPolicyId := "groupPolicyId_example" // string | Group policy ID
	updateNetworkGroupPolicyRequest := *openapiclient.NewUpdateNetworkGroupPolicyRequest() // UpdateNetworkGroupPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkGroupPolicy(context.Background(), networkId, groupPolicyId).UpdateNetworkGroupPolicyRequest(updateNetworkGroupPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkGroupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkGroupPolicy`: GetNetworkGroupPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkGroupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupPolicyId** | **string** | Group policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkGroupPolicyRequest** | [**UpdateNetworkGroupPolicyRequest**](UpdateNetworkGroupPolicyRequest.md) |  | 

### Return type

[**GetNetworkGroupPolicies200ResponseInner**](GetNetworkGroupPolicies200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkMerakiAuthUser

> GetNetworkMerakiAuthUsers200ResponseInner UpdateNetworkMerakiAuthUser(ctx, networkId, merakiAuthUserId).UpdateNetworkMerakiAuthUserRequest(updateNetworkMerakiAuthUserRequest).Execute()

Update a user configured with Meraki Authentication (currently, 802.1X RADIUS, splash guest, and client VPN users can be updated)



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
	merakiAuthUserId := "merakiAuthUserId_example" // string | Meraki auth user ID
	updateNetworkMerakiAuthUserRequest := *openapiclient.NewUpdateNetworkMerakiAuthUserRequest() // UpdateNetworkMerakiAuthUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkMerakiAuthUser(context.Background(), networkId, merakiAuthUserId).UpdateNetworkMerakiAuthUserRequest(updateNetworkMerakiAuthUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkMerakiAuthUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkMerakiAuthUser`: GetNetworkMerakiAuthUsers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkMerakiAuthUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**merakiAuthUserId** | **string** | Meraki auth user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkMerakiAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkMerakiAuthUserRequest** | [**UpdateNetworkMerakiAuthUserRequest**](UpdateNetworkMerakiAuthUserRequest.md) |  | 

### Return type

[**GetNetworkMerakiAuthUsers200ResponseInner**](GetNetworkMerakiAuthUsers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkMqttBroker

> GetNetworkMqttBrokers200ResponseInner UpdateNetworkMqttBroker(ctx, networkId, mqttBrokerId).UpdateNetworkMqttBrokerRequest(updateNetworkMqttBrokerRequest).Execute()

Update an MQTT broker



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
	mqttBrokerId := "mqttBrokerId_example" // string | Mqtt broker ID
	updateNetworkMqttBrokerRequest := *openapiclient.NewUpdateNetworkMqttBrokerRequest() // UpdateNetworkMqttBrokerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkMqttBroker(context.Background(), networkId, mqttBrokerId).UpdateNetworkMqttBrokerRequest(updateNetworkMqttBrokerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkMqttBroker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkMqttBroker`: GetNetworkMqttBrokers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkMqttBroker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**mqttBrokerId** | **string** | Mqtt broker ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkMqttBrokerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkMqttBrokerRequest** | [**UpdateNetworkMqttBrokerRequest**](UpdateNetworkMqttBrokerRequest.md) |  | 

### Return type

[**GetNetworkMqttBrokers200ResponseInner**](GetNetworkMqttBrokers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkNetflow

> GetNetworkNetflow200Response UpdateNetworkNetflow(ctx, networkId).UpdateNetworkNetflowRequest(updateNetworkNetflowRequest).Execute()

Update the NetFlow traffic reporting settings for a network



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
	updateNetworkNetflowRequest := *openapiclient.NewUpdateNetworkNetflowRequest() // UpdateNetworkNetflowRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkNetflow(context.Background(), networkId).UpdateNetworkNetflowRequest(updateNetworkNetflowRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkNetflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkNetflow`: GetNetworkNetflow200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkNetflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkNetflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkNetflowRequest** | [**UpdateNetworkNetflowRequest**](UpdateNetworkNetflowRequest.md) |  | 

### Return type

[**GetNetworkNetflow200Response**](GetNetworkNetflow200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSettings

> GetNetworkSettings200Response UpdateNetworkSettings(ctx, networkId).UpdateNetworkSettingsRequest(updateNetworkSettingsRequest).Execute()

Update the settings for a network



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
	updateNetworkSettingsRequest := *openapiclient.NewUpdateNetworkSettingsRequest() // UpdateNetworkSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkSettings(context.Background(), networkId).UpdateNetworkSettingsRequest(updateNetworkSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSettings`: GetNetworkSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSettingsRequest** | [**UpdateNetworkSettingsRequest**](UpdateNetworkSettingsRequest.md) |  | 

### Return type

[**GetNetworkSettings200Response**](GetNetworkSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSnmp

> GetNetworkSnmp200Response UpdateNetworkSnmp(ctx, networkId).UpdateNetworkSnmpRequest(updateNetworkSnmpRequest).Execute()

Update the SNMP settings for a network



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
	updateNetworkSnmpRequest := *openapiclient.NewUpdateNetworkSnmpRequest() // UpdateNetworkSnmpRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkSnmp(context.Background(), networkId).UpdateNetworkSnmpRequest(updateNetworkSnmpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkSnmp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSnmp`: GetNetworkSnmp200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkSnmp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSnmpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSnmpRequest** | [**UpdateNetworkSnmpRequest**](UpdateNetworkSnmpRequest.md) |  | 

### Return type

[**GetNetworkSnmp200Response**](GetNetworkSnmp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSyslogServers

> GetNetworkSyslogServers200Response UpdateNetworkSyslogServers(ctx, networkId).UpdateNetworkSyslogServersRequest(updateNetworkSyslogServersRequest).Execute()

Update the syslog servers for a network



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
	updateNetworkSyslogServersRequest := *openapiclient.NewUpdateNetworkSyslogServersRequest([]openapiclient.UpdateNetworkSyslogServersRequestServersInner{*openapiclient.NewUpdateNetworkSyslogServersRequestServersInner("Host_example", int32(123), []string{"Roles_example"})}) // UpdateNetworkSyslogServersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkSyslogServers(context.Background(), networkId).UpdateNetworkSyslogServersRequest(updateNetworkSyslogServersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkSyslogServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSyslogServers`: GetNetworkSyslogServers200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkSyslogServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSyslogServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSyslogServersRequest** | [**UpdateNetworkSyslogServersRequest**](UpdateNetworkSyslogServersRequest.md) |  | 

### Return type

[**GetNetworkSyslogServers200Response**](GetNetworkSyslogServers200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkTrafficAnalysis

> GetNetworkTrafficAnalysis200Response UpdateNetworkTrafficAnalysis(ctx, networkId).UpdateNetworkTrafficAnalysisRequest(updateNetworkTrafficAnalysisRequest).Execute()

Update the traffic analysis settings for a network



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
	updateNetworkTrafficAnalysisRequest := *openapiclient.NewUpdateNetworkTrafficAnalysisRequest() // UpdateNetworkTrafficAnalysisRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkTrafficAnalysis(context.Background(), networkId).UpdateNetworkTrafficAnalysisRequest(updateNetworkTrafficAnalysisRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkTrafficAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkTrafficAnalysis`: GetNetworkTrafficAnalysis200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkTrafficAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkTrafficAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkTrafficAnalysisRequest** | [**UpdateNetworkTrafficAnalysisRequest**](UpdateNetworkTrafficAnalysisRequest.md) |  | 

### Return type

[**GetNetworkTrafficAnalysis200Response**](GetNetworkTrafficAnalysis200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkVlanProfile

> GetNetworkVlanProfiles200ResponseInner UpdateNetworkVlanProfile(ctx, networkId, iname).UpdateNetworkVlanProfileRequest(updateNetworkVlanProfileRequest).Execute()

Update an existing VLAN profile of a network



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
	iname := "iname_example" // string | Iname
	updateNetworkVlanProfileRequest := *openapiclient.NewUpdateNetworkVlanProfileRequest("Name_example", []openapiclient.CreateNetworkVlanProfileRequestVlanNamesInner{*openapiclient.NewCreateNetworkVlanProfileRequestVlanNamesInner("Name_example", "VlanId_example")}, []openapiclient.CreateNetworkVlanProfileRequestVlanGroupsInner{*openapiclient.NewCreateNetworkVlanProfileRequestVlanGroupsInner("Name_example", "VlanIds_example")}) // UpdateNetworkVlanProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkVlanProfile(context.Background(), networkId, iname).UpdateNetworkVlanProfileRequest(updateNetworkVlanProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkVlanProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkVlanProfile`: GetNetworkVlanProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkVlanProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**iname** | **string** | Iname | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkVlanProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkVlanProfileRequest** | [**UpdateNetworkVlanProfileRequest**](UpdateNetworkVlanProfileRequest.md) |  | 

### Return type

[**GetNetworkVlanProfiles200ResponseInner**](GetNetworkVlanProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWebhooksHttpServer

> GetNetworkWebhooksHttpServers200ResponseInner UpdateNetworkWebhooksHttpServer(ctx, networkId, httpServerId).UpdateNetworkWebhooksHttpServerRequest(updateNetworkWebhooksHttpServerRequest).Execute()

Update an HTTP server



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
	httpServerId := "httpServerId_example" // string | Http server ID
	updateNetworkWebhooksHttpServerRequest := *openapiclient.NewUpdateNetworkWebhooksHttpServerRequest() // UpdateNetworkWebhooksHttpServerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkWebhooksHttpServer(context.Background(), networkId, httpServerId).UpdateNetworkWebhooksHttpServerRequest(updateNetworkWebhooksHttpServerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkWebhooksHttpServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkWebhooksHttpServer`: GetNetworkWebhooksHttpServers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkWebhooksHttpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**httpServerId** | **string** | Http server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWebhooksHttpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWebhooksHttpServerRequest** | [**UpdateNetworkWebhooksHttpServerRequest**](UpdateNetworkWebhooksHttpServerRequest.md) |  | 

### Return type

[**GetNetworkWebhooksHttpServers200ResponseInner**](GetNetworkWebhooksHttpServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWebhooksPayloadTemplate

> GetNetworkWebhooksPayloadTemplates200ResponseInner UpdateNetworkWebhooksPayloadTemplate(ctx, networkId, payloadTemplateId).UpdateNetworkWebhooksPayloadTemplateRequest(updateNetworkWebhooksPayloadTemplateRequest).Execute()

Update a webhook payload template for a network



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
	payloadTemplateId := "payloadTemplateId_example" // string | Payload template ID
	updateNetworkWebhooksPayloadTemplateRequest := *openapiclient.NewUpdateNetworkWebhooksPayloadTemplateRequest() // UpdateNetworkWebhooksPayloadTemplateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateNetworkWebhooksPayloadTemplate(context.Background(), networkId, payloadTemplateId).UpdateNetworkWebhooksPayloadTemplateRequest(updateNetworkWebhooksPayloadTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateNetworkWebhooksPayloadTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkWebhooksPayloadTemplate`: GetNetworkWebhooksPayloadTemplates200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateNetworkWebhooksPayloadTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**payloadTemplateId** | **string** | Payload template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWebhooksPayloadTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWebhooksPayloadTemplateRequest** | [**UpdateNetworkWebhooksPayloadTemplateRequest**](UpdateNetworkWebhooksPayloadTemplateRequest.md) |  | 

### Return type

[**GetNetworkWebhooksPayloadTemplates200ResponseInner**](GetNetworkWebhooksPayloadTemplates200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmxNetworkDevicesClaim

> VmxNetworkDevicesClaim200Response VmxNetworkDevicesClaim(ctx, networkId).VmxNetworkDevicesClaimRequest(vmxNetworkDevicesClaimRequest).Execute()

Claim a vMX into a network



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
	vmxNetworkDevicesClaimRequest := *openapiclient.NewVmxNetworkDevicesClaimRequest("Size_example") // VmxNetworkDevicesClaimRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.VmxNetworkDevicesClaim(context.Background(), networkId).VmxNetworkDevicesClaimRequest(vmxNetworkDevicesClaimRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.VmxNetworkDevicesClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VmxNetworkDevicesClaim`: VmxNetworkDevicesClaim200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.VmxNetworkDevicesClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmxNetworkDevicesClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vmxNetworkDevicesClaimRequest** | [**VmxNetworkDevicesClaimRequest**](VmxNetworkDevicesClaimRequest.md) |  | 

### Return type

[**VmxNetworkDevicesClaim200Response**](VmxNetworkDevicesClaim200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

