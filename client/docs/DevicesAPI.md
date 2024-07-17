# \DevicesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlinkDeviceLeds**](DevicesAPI.md#BlinkDeviceLeds) | **Post** /devices/{serial}/blinkLeds | Blink the LEDs on a device
[**BulkUpdateOrganizationDevicesDetails**](DevicesAPI.md#BulkUpdateOrganizationDevicesDetails) | **Post** /organizations/{organizationId}/devices/details/bulkUpdate | Updating device details (currently only used for Catalyst devices)
[**CheckinNetworkSmDevices**](DevicesAPI.md#CheckinNetworkSmDevices) | **Post** /networks/{networkId}/sm/devices/checkin | Force check-in a set of devices
[**ClaimNetworkDevices**](DevicesAPI.md#ClaimNetworkDevices) | **Post** /networks/{networkId}/devices/claim | Claim devices into a network. (Note: for recently claimed devices, it may take a few minutes for API requests against that device to succeed)
[**CloneOrganizationSwitchDevices**](DevicesAPI.md#CloneOrganizationSwitchDevices) | **Post** /organizations/{organizationId}/switch/devices/clone | Clone port-level and some switch-level configuration settings from a source switch to one or more target switches
[**CreateDeviceLiveToolsArpTable**](DevicesAPI.md#CreateDeviceLiveToolsArpTable) | **Post** /devices/{serial}/liveTools/arpTable | Enqueue a job to perform a ARP table request for the device
[**CreateDeviceLiveToolsCableTest**](DevicesAPI.md#CreateDeviceLiveToolsCableTest) | **Post** /devices/{serial}/liveTools/cableTest | Enqueue a job to perform a cable test for the device on the specified ports
[**CreateDeviceLiveToolsPing**](DevicesAPI.md#CreateDeviceLiveToolsPing) | **Post** /devices/{serial}/liveTools/ping | Enqueue a job to ping a target host from the device
[**CreateDeviceLiveToolsPingDevice**](DevicesAPI.md#CreateDeviceLiveToolsPingDevice) | **Post** /devices/{serial}/liveTools/pingDevice | Enqueue a job to check connectivity status to the device
[**CreateDeviceLiveToolsThroughputTest**](DevicesAPI.md#CreateDeviceLiveToolsThroughputTest) | **Post** /devices/{serial}/liveTools/throughputTest | Enqueue a job to test a device throughput, the test will run for 10 secs to test throughput
[**CreateDeviceLiveToolsWakeOnLan**](DevicesAPI.md#CreateDeviceLiveToolsWakeOnLan) | **Post** /devices/{serial}/liveTools/wakeOnLan | Enqueue a job to send a Wake-on-LAN packet from the device
[**CreateOrganizationInventoryDevicesSwapsBulk**](DevicesAPI.md#CreateOrganizationInventoryDevicesSwapsBulk) | **Post** /organizations/{organizationId}/inventory/devices/swaps/bulk | Swap the devices identified by devices.old with a devices.new, then perform the :afterAction on the devices.old.
[**GetDevice**](DevicesAPI.md#GetDevice) | **Get** /devices/{serial} | Return a single device
[**GetDeviceCellularSims**](DevicesAPI.md#GetDeviceCellularSims) | **Get** /devices/{serial}/cellular/sims | Return the SIM and APN configurations for a cellular device.
[**GetDeviceClients**](DevicesAPI.md#GetDeviceClients) | **Get** /devices/{serial}/clients | List the clients of a device, up to a maximum of a month ago
[**GetDeviceLiveToolsArpTable**](DevicesAPI.md#GetDeviceLiveToolsArpTable) | **Get** /devices/{serial}/liveTools/arpTable/{arpTableId} | Return an ARP table live tool job.
[**GetDeviceLiveToolsCableTest**](DevicesAPI.md#GetDeviceLiveToolsCableTest) | **Get** /devices/{serial}/liveTools/cableTest/{id} | Return a cable test live tool job.
[**GetDeviceLiveToolsPing**](DevicesAPI.md#GetDeviceLiveToolsPing) | **Get** /devices/{serial}/liveTools/ping/{id} | Return a ping job
[**GetDeviceLiveToolsPingDevice**](DevicesAPI.md#GetDeviceLiveToolsPingDevice) | **Get** /devices/{serial}/liveTools/pingDevice/{id} | Return a ping device job
[**GetDeviceLiveToolsThroughputTest**](DevicesAPI.md#GetDeviceLiveToolsThroughputTest) | **Get** /devices/{serial}/liveTools/throughputTest/{throughputTestId} | Return a throughput test job
[**GetDeviceLiveToolsWakeOnLan**](DevicesAPI.md#GetDeviceLiveToolsWakeOnLan) | **Get** /devices/{serial}/liveTools/wakeOnLan/{wakeOnLanId} | Return a Wake-on-LAN job
[**GetDeviceLldpCdp**](DevicesAPI.md#GetDeviceLldpCdp) | **Get** /devices/{serial}/lldpCdp | List LLDP and CDP information for a device
[**GetDeviceLossAndLatencyHistory**](DevicesAPI.md#GetDeviceLossAndLatencyHistory) | **Get** /devices/{serial}/lossAndLatencyHistory | Get the uplink loss percentage and latency in milliseconds, and goodput in kilobits per second for MX, MG and Z devices.
[**GetDeviceManagementInterface**](DevicesAPI.md#GetDeviceManagementInterface) | **Get** /devices/{serial}/managementInterface | Return the management interface settings for a device
[**GetNetworkDevices**](DevicesAPI.md#GetNetworkDevices) | **Get** /networks/{networkId}/devices | List the devices in a network
[**GetNetworkSmDeviceCellularUsageHistory**](DevicesAPI.md#GetNetworkSmDeviceCellularUsageHistory) | **Get** /networks/{networkId}/sm/devices/{deviceId}/cellularUsageHistory | Return the client&#39;s daily cellular data usage history
[**GetNetworkSmDeviceCerts**](DevicesAPI.md#GetNetworkSmDeviceCerts) | **Get** /networks/{networkId}/sm/devices/{deviceId}/certs | List the certs on a device
[**GetNetworkSmDeviceConnectivity**](DevicesAPI.md#GetNetworkSmDeviceConnectivity) | **Get** /networks/{networkId}/sm/devices/{deviceId}/connectivity | Returns historical connectivity data (whether a device is regularly checking in to Dashboard).
[**GetNetworkSmDeviceDesktopLogs**](DevicesAPI.md#GetNetworkSmDeviceDesktopLogs) | **Get** /networks/{networkId}/sm/devices/{deviceId}/desktopLogs | Return historical records of various Systems Manager network connection details for desktop devices.
[**GetNetworkSmDeviceDeviceCommandLogs**](DevicesAPI.md#GetNetworkSmDeviceDeviceCommandLogs) | **Get** /networks/{networkId}/sm/devices/{deviceId}/deviceCommandLogs | Return historical records of commands sent to Systems Manager devices
[**GetNetworkSmDeviceDeviceProfiles**](DevicesAPI.md#GetNetworkSmDeviceDeviceProfiles) | **Get** /networks/{networkId}/sm/devices/{deviceId}/deviceProfiles | Get the installed profiles associated with a device
[**GetNetworkSmDeviceNetworkAdapters**](DevicesAPI.md#GetNetworkSmDeviceNetworkAdapters) | **Get** /networks/{networkId}/sm/devices/{deviceId}/networkAdapters | List the network adapters of a device
[**GetNetworkSmDevicePerformanceHistory**](DevicesAPI.md#GetNetworkSmDevicePerformanceHistory) | **Get** /networks/{networkId}/sm/devices/{deviceId}/performanceHistory | Return historical records of various Systems Manager client metrics for desktop devices.
[**GetNetworkSmDeviceRestrictions**](DevicesAPI.md#GetNetworkSmDeviceRestrictions) | **Get** /networks/{networkId}/sm/devices/{deviceId}/restrictions | List the restrictions on a device
[**GetNetworkSmDeviceSecurityCenters**](DevicesAPI.md#GetNetworkSmDeviceSecurityCenters) | **Get** /networks/{networkId}/sm/devices/{deviceId}/securityCenters | List the security centers on a device
[**GetNetworkSmDeviceSoftwares**](DevicesAPI.md#GetNetworkSmDeviceSoftwares) | **Get** /networks/{networkId}/sm/devices/{deviceId}/softwares | Get a list of softwares associated with a device
[**GetNetworkSmDeviceWlanLists**](DevicesAPI.md#GetNetworkSmDeviceWlanLists) | **Get** /networks/{networkId}/sm/devices/{deviceId}/wlanLists | List the saved SSID names on a device
[**GetNetworkSmDevices**](DevicesAPI.md#GetNetworkSmDevices) | **Get** /networks/{networkId}/sm/devices | List the devices enrolled in an SM network with various specified fields and filters
[**GetNetworkWirelessDevicesConnectionStats**](DevicesAPI.md#GetNetworkWirelessDevicesConnectionStats) | **Get** /networks/{networkId}/wireless/devices/connectionStats | Aggregated connectivity info for this network, grouped by node
[**GetNetworkWirelessDevicesLatencyStats**](DevicesAPI.md#GetNetworkWirelessDevicesLatencyStats) | **Get** /networks/{networkId}/wireless/devices/latencyStats | Aggregated latency info for this network, grouped by node
[**GetOrganizationDevices**](DevicesAPI.md#GetOrganizationDevices) | **Get** /organizations/{organizationId}/devices | List the devices in an organization
[**GetOrganizationDevicesAvailabilities**](DevicesAPI.md#GetOrganizationDevicesAvailabilities) | **Get** /organizations/{organizationId}/devices/availabilities | List the availability information for devices in an organization
[**GetOrganizationDevicesAvailabilitiesChangeHistory**](DevicesAPI.md#GetOrganizationDevicesAvailabilitiesChangeHistory) | **Get** /organizations/{organizationId}/devices/availabilities/changeHistory | List the availability history information for devices in an organization.
[**GetOrganizationDevicesOverviewByModel**](DevicesAPI.md#GetOrganizationDevicesOverviewByModel) | **Get** /organizations/{organizationId}/devices/overview/byModel | Lists the count for each device model
[**GetOrganizationDevicesPowerModulesStatusesByDevice**](DevicesAPI.md#GetOrganizationDevicesPowerModulesStatusesByDevice) | **Get** /organizations/{organizationId}/devices/powerModules/statuses/byDevice | List the most recent status information for power modules in rackmount MX and MS devices that support them
[**GetOrganizationDevicesProvisioningStatuses**](DevicesAPI.md#GetOrganizationDevicesProvisioningStatuses) | **Get** /organizations/{organizationId}/devices/provisioning/statuses | List the provisioning statuses information for devices in an organization.
[**GetOrganizationDevicesStatuses**](DevicesAPI.md#GetOrganizationDevicesStatuses) | **Get** /organizations/{organizationId}/devices/statuses | List the status of every Meraki device in the organization
[**GetOrganizationDevicesStatusesOverview**](DevicesAPI.md#GetOrganizationDevicesStatusesOverview) | **Get** /organizations/{organizationId}/devices/statuses/overview | Return an overview of current device statuses
[**GetOrganizationDevicesUplinksAddressesByDevice**](DevicesAPI.md#GetOrganizationDevicesUplinksAddressesByDevice) | **Get** /organizations/{organizationId}/devices/uplinks/addresses/byDevice | List the current uplink addresses for devices in an organization.
[**GetOrganizationDevicesUplinksLossAndLatency**](DevicesAPI.md#GetOrganizationDevicesUplinksLossAndLatency) | **Get** /organizations/{organizationId}/devices/uplinksLossAndLatency | Return the uplink loss and latency for every MX in the organization from at latest 2 minutes ago
[**GetOrganizationInventoryDevice**](DevicesAPI.md#GetOrganizationInventoryDevice) | **Get** /organizations/{organizationId}/inventory/devices/{serial} | Return a single device from the inventory of an organization
[**GetOrganizationInventoryDevices**](DevicesAPI.md#GetOrganizationInventoryDevices) | **Get** /organizations/{organizationId}/inventory/devices | Return the device inventory for an organization
[**GetOrganizationInventoryDevicesSwapsBulk**](DevicesAPI.md#GetOrganizationInventoryDevicesSwapsBulk) | **Get** /organizations/{organizationId}/inventory/devices/swaps/bulk/{id} | List of device swaps for a given request ID ({id}).
[**GetOrganizationSummaryTopDevicesByUsage**](DevicesAPI.md#GetOrganizationSummaryTopDevicesByUsage) | **Get** /organizations/{organizationId}/summary/top/devices/byUsage | Return metrics for organization&#39;s top 10 devices sorted by data usage over given time range
[**GetOrganizationSummaryTopDevicesModelsByUsage**](DevicesAPI.md#GetOrganizationSummaryTopDevicesModelsByUsage) | **Get** /organizations/{organizationId}/summary/top/devices/models/byUsage | Return metrics for organization&#39;s top 10 device models sorted by data usage over given time range
[**GetOrganizationWirelessDevicesChannelUtilizationByDevice**](DevicesAPI.md#GetOrganizationWirelessDevicesChannelUtilizationByDevice) | **Get** /organizations/{organizationId}/wireless/devices/channelUtilization/byDevice | Get average channel utilization for all bands in a network, split by AP
[**GetOrganizationWirelessDevicesChannelUtilizationByNetwork**](DevicesAPI.md#GetOrganizationWirelessDevicesChannelUtilizationByNetwork) | **Get** /organizations/{organizationId}/wireless/devices/channelUtilization/byNetwork | Get average channel utilization across all bands for all networks in the organization
[**GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval**](DevicesAPI.md#GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval) | **Get** /organizations/{organizationId}/wireless/devices/channelUtilization/history/byDevice/byInterval | Get a time-series of average channel utilization for all bands, segmented by device.
[**GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval**](DevicesAPI.md#GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval) | **Get** /organizations/{organizationId}/wireless/devices/channelUtilization/history/byNetwork/byInterval | Get a time-series of average channel utilization for all bands
[**GetOrganizationWirelessDevicesEthernetStatuses**](DevicesAPI.md#GetOrganizationWirelessDevicesEthernetStatuses) | **Get** /organizations/{organizationId}/wireless/devices/ethernet/statuses | List the most recent Ethernet link speed, duplex, aggregation and power mode and status information for wireless devices.
[**GetOrganizationWirelessDevicesPacketLossByClient**](DevicesAPI.md#GetOrganizationWirelessDevicesPacketLossByClient) | **Get** /organizations/{organizationId}/wireless/devices/packetLoss/byClient | Get average packet loss for the given timespan for all clients in the organization.
[**GetOrganizationWirelessDevicesPacketLossByDevice**](DevicesAPI.md#GetOrganizationWirelessDevicesPacketLossByDevice) | **Get** /organizations/{organizationId}/wireless/devices/packetLoss/byDevice | Get average packet loss for the given timespan for all devices in the organization
[**GetOrganizationWirelessDevicesPacketLossByNetwork**](DevicesAPI.md#GetOrganizationWirelessDevicesPacketLossByNetwork) | **Get** /organizations/{organizationId}/wireless/devices/packetLoss/byNetwork | Get average packet loss for the given timespan for all networks in the organization.
[**InstallNetworkSmDeviceApps**](DevicesAPI.md#InstallNetworkSmDeviceApps) | **Post** /networks/{networkId}/sm/devices/{deviceId}/installApps | Install applications on a device
[**LockNetworkSmDevices**](DevicesAPI.md#LockNetworkSmDevices) | **Post** /networks/{networkId}/sm/devices/lock | Lock a set of devices
[**ModifyNetworkSmDevicesTags**](DevicesAPI.md#ModifyNetworkSmDevicesTags) | **Post** /networks/{networkId}/sm/devices/modifyTags | Add, delete, or update the tags of a set of devices
[**MoveNetworkSmDevices**](DevicesAPI.md#MoveNetworkSmDevices) | **Post** /networks/{networkId}/sm/devices/move | Move a set of devices to a new network
[**RebootDevice**](DevicesAPI.md#RebootDevice) | **Post** /devices/{serial}/reboot | Reboot a device
[**RebootNetworkSmDevices**](DevicesAPI.md#RebootNetworkSmDevices) | **Post** /networks/{networkId}/sm/devices/reboot | Reboot a set of endpoints
[**RefreshNetworkSmDeviceDetails**](DevicesAPI.md#RefreshNetworkSmDeviceDetails) | **Post** /networks/{networkId}/sm/devices/{deviceId}/refreshDetails | Refresh the details of a device
[**RemoveNetworkDevices**](DevicesAPI.md#RemoveNetworkDevices) | **Post** /networks/{networkId}/devices/remove | Remove a single device
[**ShutdownNetworkSmDevices**](DevicesAPI.md#ShutdownNetworkSmDevices) | **Post** /networks/{networkId}/sm/devices/shutdown | Shutdown a set of endpoints
[**UnenrollNetworkSmDevice**](DevicesAPI.md#UnenrollNetworkSmDevice) | **Post** /networks/{networkId}/sm/devices/{deviceId}/unenroll | Unenroll a device
[**UninstallNetworkSmDeviceApps**](DevicesAPI.md#UninstallNetworkSmDeviceApps) | **Post** /networks/{networkId}/sm/devices/{deviceId}/uninstallApps | Uninstall applications on a device
[**UpdateDevice**](DevicesAPI.md#UpdateDevice) | **Put** /devices/{serial} | Update the attributes of a device
[**UpdateDeviceCellularSims**](DevicesAPI.md#UpdateDeviceCellularSims) | **Put** /devices/{serial}/cellular/sims | Updates the SIM and APN configurations for a cellular device.
[**UpdateDeviceManagementInterface**](DevicesAPI.md#UpdateDeviceManagementInterface) | **Put** /devices/{serial}/managementInterface | Update the management interface settings for a device
[**UpdateNetworkSmDevicesFields**](DevicesAPI.md#UpdateNetworkSmDevicesFields) | **Put** /networks/{networkId}/sm/devices/fields | Modify the fields of a device
[**VmxNetworkDevicesClaim**](DevicesAPI.md#VmxNetworkDevicesClaim) | **Post** /networks/{networkId}/devices/claim/vmx | Claim a vMX into a network
[**WipeNetworkSmDevices**](DevicesAPI.md#WipeNetworkSmDevices) | **Post** /networks/{networkId}/sm/devices/wipe | Wipe a device



## BlinkDeviceLeds

> BlinkDeviceLeds202Response BlinkDeviceLeds(ctx, serial).BlinkDeviceLedsRequest(blinkDeviceLedsRequest).Execute()

Blink the LEDs on a device



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
	blinkDeviceLedsRequest := *openapiclient.NewBlinkDeviceLedsRequest() // BlinkDeviceLedsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.BlinkDeviceLeds(context.Background(), serial).BlinkDeviceLedsRequest(blinkDeviceLedsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.BlinkDeviceLeds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlinkDeviceLeds`: BlinkDeviceLeds202Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.BlinkDeviceLeds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlinkDeviceLedsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blinkDeviceLedsRequest** | [**BlinkDeviceLedsRequest**](BlinkDeviceLedsRequest.md) |  | 

### Return type

[**BlinkDeviceLeds202Response**](BlinkDeviceLeds202Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdateOrganizationDevicesDetails

> BulkUpdateOrganizationDevicesDetails200Response BulkUpdateOrganizationDevicesDetails(ctx, organizationId).BulkUpdateOrganizationDevicesDetailsRequest(bulkUpdateOrganizationDevicesDetailsRequest).Execute()

Updating device details (currently only used for Catalyst devices)



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
	bulkUpdateOrganizationDevicesDetailsRequest := *openapiclient.NewBulkUpdateOrganizationDevicesDetailsRequest([]string{"Serials_example"}, []openapiclient.BulkUpdateOrganizationDevicesDetailsRequestDetailsInner{*openapiclient.NewBulkUpdateOrganizationDevicesDetailsRequestDetailsInner("Name_example")}) // BulkUpdateOrganizationDevicesDetailsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.BulkUpdateOrganizationDevicesDetails(context.Background(), organizationId).BulkUpdateOrganizationDevicesDetailsRequest(bulkUpdateOrganizationDevicesDetailsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.BulkUpdateOrganizationDevicesDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkUpdateOrganizationDevicesDetails`: BulkUpdateOrganizationDevicesDetails200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.BulkUpdateOrganizationDevicesDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateOrganizationDevicesDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkUpdateOrganizationDevicesDetailsRequest** | [**BulkUpdateOrganizationDevicesDetailsRequest**](BulkUpdateOrganizationDevicesDetailsRequest.md) |  | 

### Return type

[**BulkUpdateOrganizationDevicesDetails200Response**](BulkUpdateOrganizationDevicesDetails200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckinNetworkSmDevices

> CheckinNetworkSmDevices200Response CheckinNetworkSmDevices(ctx, networkId).CheckinNetworkSmDevicesRequest(checkinNetworkSmDevicesRequest).Execute()

Force check-in a set of devices



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
	checkinNetworkSmDevicesRequest := *openapiclient.NewCheckinNetworkSmDevicesRequest() // CheckinNetworkSmDevicesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.CheckinNetworkSmDevices(context.Background(), networkId).CheckinNetworkSmDevicesRequest(checkinNetworkSmDevicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.CheckinNetworkSmDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckinNetworkSmDevices`: CheckinNetworkSmDevices200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.CheckinNetworkSmDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckinNetworkSmDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkinNetworkSmDevicesRequest** | [**CheckinNetworkSmDevicesRequest**](CheckinNetworkSmDevicesRequest.md) |  | 

### Return type

[**CheckinNetworkSmDevices200Response**](CheckinNetworkSmDevices200Response.md)

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
	resp, r, err := apiClient.DevicesAPI.ClaimNetworkDevices(context.Background(), networkId).ClaimNetworkDevicesRequest(claimNetworkDevicesRequest).AddAtomically(addAtomically).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.ClaimNetworkDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClaimNetworkDevices`: ClaimNetworkDevices200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.ClaimNetworkDevices`: %v\n", resp)
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
	resp, r, err := apiClient.DevicesAPI.CloneOrganizationSwitchDevices(context.Background(), organizationId).CloneOrganizationSwitchDevicesRequest(cloneOrganizationSwitchDevicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.CloneOrganizationSwitchDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloneOrganizationSwitchDevices`: CloneOrganizationSwitchDevices200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.CloneOrganizationSwitchDevices`: %v\n", resp)
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


## CreateDeviceLiveToolsArpTable

> CreateDeviceLiveToolsArpTable201Response CreateDeviceLiveToolsArpTable(ctx, serial).CreateDeviceLiveToolsArpTableRequest(createDeviceLiveToolsArpTableRequest).Execute()

Enqueue a job to perform a ARP table request for the device



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
	createDeviceLiveToolsArpTableRequest := *openapiclient.NewCreateDeviceLiveToolsArpTableRequest() // CreateDeviceLiveToolsArpTableRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.CreateDeviceLiveToolsArpTable(context.Background(), serial).CreateDeviceLiveToolsArpTableRequest(createDeviceLiveToolsArpTableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.CreateDeviceLiveToolsArpTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceLiveToolsArpTable`: CreateDeviceLiveToolsArpTable201Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.CreateDeviceLiveToolsArpTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsArpTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsArpTableRequest** | [**CreateDeviceLiveToolsArpTableRequest**](CreateDeviceLiveToolsArpTableRequest.md) |  | 

### Return type

[**CreateDeviceLiveToolsArpTable201Response**](CreateDeviceLiveToolsArpTable201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsCableTest

> CreateDeviceLiveToolsCableTest201Response CreateDeviceLiveToolsCableTest(ctx, serial).CreateDeviceLiveToolsCableTestRequest(createDeviceLiveToolsCableTestRequest).Execute()

Enqueue a job to perform a cable test for the device on the specified ports



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
	createDeviceLiveToolsCableTestRequest := *openapiclient.NewCreateDeviceLiveToolsCableTestRequest([]string{"Ports_example"}) // CreateDeviceLiveToolsCableTestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.CreateDeviceLiveToolsCableTest(context.Background(), serial).CreateDeviceLiveToolsCableTestRequest(createDeviceLiveToolsCableTestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.CreateDeviceLiveToolsCableTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceLiveToolsCableTest`: CreateDeviceLiveToolsCableTest201Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.CreateDeviceLiveToolsCableTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsCableTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsCableTestRequest** | [**CreateDeviceLiveToolsCableTestRequest**](CreateDeviceLiveToolsCableTestRequest.md) |  | 

### Return type

[**CreateDeviceLiveToolsCableTest201Response**](CreateDeviceLiveToolsCableTest201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsPing

> CreateDeviceLiveToolsPing201Response CreateDeviceLiveToolsPing(ctx, serial).CreateDeviceLiveToolsPingRequest(createDeviceLiveToolsPingRequest).Execute()

Enqueue a job to ping a target host from the device



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
	createDeviceLiveToolsPingRequest := *openapiclient.NewCreateDeviceLiveToolsPingRequest("Target_example") // CreateDeviceLiveToolsPingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.CreateDeviceLiveToolsPing(context.Background(), serial).CreateDeviceLiveToolsPingRequest(createDeviceLiveToolsPingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.CreateDeviceLiveToolsPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceLiveToolsPing`: CreateDeviceLiveToolsPing201Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.CreateDeviceLiveToolsPing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsPingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsPingRequest** | [**CreateDeviceLiveToolsPingRequest**](CreateDeviceLiveToolsPingRequest.md) |  | 

### Return type

[**CreateDeviceLiveToolsPing201Response**](CreateDeviceLiveToolsPing201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsPingDevice

> CreateDeviceLiveToolsPing201Response CreateDeviceLiveToolsPingDevice(ctx, serial).CreateDeviceLiveToolsPingDeviceRequest(createDeviceLiveToolsPingDeviceRequest).Execute()

Enqueue a job to check connectivity status to the device



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
	createDeviceLiveToolsPingDeviceRequest := *openapiclient.NewCreateDeviceLiveToolsPingDeviceRequest() // CreateDeviceLiveToolsPingDeviceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.CreateDeviceLiveToolsPingDevice(context.Background(), serial).CreateDeviceLiveToolsPingDeviceRequest(createDeviceLiveToolsPingDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.CreateDeviceLiveToolsPingDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceLiveToolsPingDevice`: CreateDeviceLiveToolsPing201Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.CreateDeviceLiveToolsPingDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsPingDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsPingDeviceRequest** | [**CreateDeviceLiveToolsPingDeviceRequest**](CreateDeviceLiveToolsPingDeviceRequest.md) |  | 

### Return type

[**CreateDeviceLiveToolsPing201Response**](CreateDeviceLiveToolsPing201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsThroughputTest

> CreateDeviceLiveToolsThroughputTest201Response CreateDeviceLiveToolsThroughputTest(ctx, serial).CreateDeviceLiveToolsArpTableRequest(createDeviceLiveToolsArpTableRequest).Execute()

Enqueue a job to test a device throughput, the test will run for 10 secs to test throughput



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
	createDeviceLiveToolsArpTableRequest := *openapiclient.NewCreateDeviceLiveToolsArpTableRequest() // CreateDeviceLiveToolsArpTableRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.CreateDeviceLiveToolsThroughputTest(context.Background(), serial).CreateDeviceLiveToolsArpTableRequest(createDeviceLiveToolsArpTableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.CreateDeviceLiveToolsThroughputTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceLiveToolsThroughputTest`: CreateDeviceLiveToolsThroughputTest201Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.CreateDeviceLiveToolsThroughputTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsThroughputTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsArpTableRequest** | [**CreateDeviceLiveToolsArpTableRequest**](CreateDeviceLiveToolsArpTableRequest.md) |  | 

### Return type

[**CreateDeviceLiveToolsThroughputTest201Response**](CreateDeviceLiveToolsThroughputTest201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsWakeOnLan

> CreateDeviceLiveToolsWakeOnLan201Response CreateDeviceLiveToolsWakeOnLan(ctx, serial).CreateDeviceLiveToolsWakeOnLanRequest(createDeviceLiveToolsWakeOnLanRequest).Execute()

Enqueue a job to send a Wake-on-LAN packet from the device



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
	createDeviceLiveToolsWakeOnLanRequest := *openapiclient.NewCreateDeviceLiveToolsWakeOnLanRequest(int32(123), "Mac_example") // CreateDeviceLiveToolsWakeOnLanRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.CreateDeviceLiveToolsWakeOnLan(context.Background(), serial).CreateDeviceLiveToolsWakeOnLanRequest(createDeviceLiveToolsWakeOnLanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.CreateDeviceLiveToolsWakeOnLan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceLiveToolsWakeOnLan`: CreateDeviceLiveToolsWakeOnLan201Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.CreateDeviceLiveToolsWakeOnLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsWakeOnLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsWakeOnLanRequest** | [**CreateDeviceLiveToolsWakeOnLanRequest**](CreateDeviceLiveToolsWakeOnLanRequest.md) |  | 

### Return type

[**CreateDeviceLiveToolsWakeOnLan201Response**](CreateDeviceLiveToolsWakeOnLan201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationInventoryDevicesSwapsBulk

> CreateOrganizationInventoryDevicesSwapsBulk207Response CreateOrganizationInventoryDevicesSwapsBulk(ctx, organizationId).CreateOrganizationInventoryDevicesSwapsBulkRequest(createOrganizationInventoryDevicesSwapsBulkRequest).Execute()

Swap the devices identified by devices.old with a devices.new, then perform the :afterAction on the devices.old.



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
	createOrganizationInventoryDevicesSwapsBulkRequest := *openapiclient.NewCreateOrganizationInventoryDevicesSwapsBulkRequest([]openapiclient.CreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInner{*openapiclient.NewCreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInner(*openapiclient.NewCreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInnerDevices("Old_example", "New_example"), "AfterAction_example")}) // CreateOrganizationInventoryDevicesSwapsBulkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.CreateOrganizationInventoryDevicesSwapsBulk(context.Background(), organizationId).CreateOrganizationInventoryDevicesSwapsBulkRequest(createOrganizationInventoryDevicesSwapsBulkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.CreateOrganizationInventoryDevicesSwapsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationInventoryDevicesSwapsBulk`: CreateOrganizationInventoryDevicesSwapsBulk207Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.CreateOrganizationInventoryDevicesSwapsBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInventoryDevicesSwapsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInventoryDevicesSwapsBulkRequest** | [**CreateOrganizationInventoryDevicesSwapsBulkRequest**](CreateOrganizationInventoryDevicesSwapsBulkRequest.md) |  | 

### Return type

[**CreateOrganizationInventoryDevicesSwapsBulk207Response**](CreateOrganizationInventoryDevicesSwapsBulk207Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevice

> GetDevice200Response GetDevice(ctx, serial).Execute()

Return a single device



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
	resp, r, err := apiClient.DevicesAPI.GetDevice(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevice`: GetDevice200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDevice200Response**](GetDevice200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceCellularSims

> GetDeviceCellularSims200Response GetDeviceCellularSims(ctx, serial).Execute()

Return the SIM and APN configurations for a cellular device.



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
	resp, r, err := apiClient.DevicesAPI.GetDeviceCellularSims(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDeviceCellularSims``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCellularSims`: GetDeviceCellularSims200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDeviceCellularSims`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCellularSimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceCellularSims200Response**](GetDeviceCellularSims200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceClients

> []GetDeviceClients200ResponseInner GetDeviceClients(ctx, serial).T0(t0).Timespan(timespan).Execute()

List the clients of a device, up to a maximum of a month ago



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
	resp, r, err := apiClient.DevicesAPI.GetDeviceClients(context.Background(), serial).T0(t0).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDeviceClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceClients`: []GetDeviceClients200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDeviceClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

[**[]GetDeviceClients200ResponseInner**](GetDeviceClients200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsArpTable

> DevicesSerialLiveToolsArpTablePostRequestMessage GetDeviceLiveToolsArpTable(ctx, serial, arpTableId).Execute()

Return an ARP table live tool job.



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
	arpTableId := "arpTableId_example" // string | Arp table ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetDeviceLiveToolsArpTable(context.Background(), serial, arpTableId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDeviceLiveToolsArpTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceLiveToolsArpTable`: DevicesSerialLiveToolsArpTablePostRequestMessage
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDeviceLiveToolsArpTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**arpTableId** | **string** | Arp table ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsArpTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DevicesSerialLiveToolsArpTablePostRequestMessage**](DevicesSerialLiveToolsArpTablePostRequestMessage.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsCableTest

> DevicesSerialLiveToolsCableTestPostRequestMessage GetDeviceLiveToolsCableTest(ctx, serial, id).Execute()

Return a cable test live tool job.



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
	id := "id_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetDeviceLiveToolsCableTest(context.Background(), serial, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDeviceLiveToolsCableTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceLiveToolsCableTest`: DevicesSerialLiveToolsCableTestPostRequestMessage
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDeviceLiveToolsCableTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsCableTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DevicesSerialLiveToolsCableTestPostRequestMessage**](DevicesSerialLiveToolsCableTestPostRequestMessage.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsPing

> DevicesSerialLiveToolsPingPostRequestMessage GetDeviceLiveToolsPing(ctx, serial, id).Execute()

Return a ping job



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
	id := "id_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetDeviceLiveToolsPing(context.Background(), serial, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDeviceLiveToolsPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceLiveToolsPing`: DevicesSerialLiveToolsPingPostRequestMessage
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDeviceLiveToolsPing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsPingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DevicesSerialLiveToolsPingPostRequestMessage**](DevicesSerialLiveToolsPingPostRequestMessage.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsPingDevice

> GetDeviceLiveToolsPingDevice200Response GetDeviceLiveToolsPingDevice(ctx, serial, id).Execute()

Return a ping device job



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
	id := "id_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetDeviceLiveToolsPingDevice(context.Background(), serial, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDeviceLiveToolsPingDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceLiveToolsPingDevice`: GetDeviceLiveToolsPingDevice200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDeviceLiveToolsPingDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsPingDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDeviceLiveToolsPingDevice200Response**](GetDeviceLiveToolsPingDevice200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsThroughputTest

> DevicesSerialLiveToolsThroughputTestPostRequestMessage GetDeviceLiveToolsThroughputTest(ctx, serial, throughputTestId).Execute()

Return a throughput test job



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
	throughputTestId := "throughputTestId_example" // string | Throughput test ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetDeviceLiveToolsThroughputTest(context.Background(), serial, throughputTestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDeviceLiveToolsThroughputTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceLiveToolsThroughputTest`: DevicesSerialLiveToolsThroughputTestPostRequestMessage
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDeviceLiveToolsThroughputTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**throughputTestId** | **string** | Throughput test ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsThroughputTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DevicesSerialLiveToolsThroughputTestPostRequestMessage**](DevicesSerialLiveToolsThroughputTestPostRequestMessage.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsWakeOnLan

> DevicesSerialLiveToolsWakeOnLanPostRequestMessage GetDeviceLiveToolsWakeOnLan(ctx, serial, wakeOnLanId).Execute()

Return a Wake-on-LAN job



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
	wakeOnLanId := "wakeOnLanId_example" // string | Wake on lan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetDeviceLiveToolsWakeOnLan(context.Background(), serial, wakeOnLanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDeviceLiveToolsWakeOnLan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceLiveToolsWakeOnLan`: DevicesSerialLiveToolsWakeOnLanPostRequestMessage
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDeviceLiveToolsWakeOnLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**wakeOnLanId** | **string** | Wake on lan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsWakeOnLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DevicesSerialLiveToolsWakeOnLanPostRequestMessage**](DevicesSerialLiveToolsWakeOnLanPostRequestMessage.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLldpCdp

> GetDeviceLldpCdp200Response GetDeviceLldpCdp(ctx, serial).Execute()

List LLDP and CDP information for a device



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
	resp, r, err := apiClient.DevicesAPI.GetDeviceLldpCdp(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDeviceLldpCdp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceLldpCdp`: GetDeviceLldpCdp200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDeviceLldpCdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLldpCdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceLldpCdp200Response**](GetDeviceLldpCdp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLossAndLatencyHistory

> []GetDeviceLossAndLatencyHistory200ResponseInner GetDeviceLossAndLatencyHistory(ctx, serial).Ip(ip).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Uplink(uplink).Execute()

Get the uplink loss percentage and latency in milliseconds, and goodput in kilobits per second for MX, MG and Z devices.



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
	ip := "ip_example" // string | The destination IP used to obtain the requested stats. This is required.
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 60 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)
	resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 60, 600, 3600, 86400. The default is 60. (optional)
	uplink := "uplink_example" // string | The WAN uplink used to obtain the requested stats. Valid uplinks are wan1, wan2, wan3, cellular. The default is wan1. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetDeviceLossAndLatencyHistory(context.Background(), serial).Ip(ip).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Uplink(uplink).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDeviceLossAndLatencyHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceLossAndLatencyHistory`: []GetDeviceLossAndLatencyHistory200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDeviceLossAndLatencyHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLossAndLatencyHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ip** | **string** | The destination IP used to obtain the requested stats. This is required. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 60 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 60, 600, 3600, 86400. The default is 60. | 
 **uplink** | **string** | The WAN uplink used to obtain the requested stats. Valid uplinks are wan1, wan2, wan3, cellular. The default is wan1. | 

### Return type

[**[]GetDeviceLossAndLatencyHistory200ResponseInner**](GetDeviceLossAndLatencyHistory200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceManagementInterface

> GetDeviceManagementInterface200Response GetDeviceManagementInterface(ctx, serial).Execute()

Return the management interface settings for a device



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
	resp, r, err := apiClient.DevicesAPI.GetDeviceManagementInterface(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDeviceManagementInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceManagementInterface`: GetDeviceManagementInterface200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDeviceManagementInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceManagementInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceManagementInterface200Response**](GetDeviceManagementInterface200Response.md)

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
	resp, r, err := apiClient.DevicesAPI.GetNetworkDevices(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetNetworkDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkDevices`: []GetDevice200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetNetworkDevices`: %v\n", resp)
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


## GetNetworkSmDeviceCellularUsageHistory

> []GetNetworkSmDeviceCellularUsageHistory200ResponseInner GetNetworkSmDeviceCellularUsageHistory(ctx, networkId, deviceId).Execute()

Return the client's daily cellular data usage history



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
	deviceId := "deviceId_example" // string | Device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetNetworkSmDeviceCellularUsageHistory(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetNetworkSmDeviceCellularUsageHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceCellularUsageHistory`: []GetNetworkSmDeviceCellularUsageHistory200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetNetworkSmDeviceCellularUsageHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceCellularUsageHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceCellularUsageHistory200ResponseInner**](GetNetworkSmDeviceCellularUsageHistory200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceCerts

> []GetNetworkSmDeviceCerts200ResponseInner GetNetworkSmDeviceCerts(ctx, networkId, deviceId).Execute()

List the certs on a device



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
	deviceId := "deviceId_example" // string | Device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetNetworkSmDeviceCerts(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetNetworkSmDeviceCerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceCerts`: []GetNetworkSmDeviceCerts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetNetworkSmDeviceCerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceCertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceCerts200ResponseInner**](GetNetworkSmDeviceCerts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceConnectivity

> []GetNetworkSmDeviceConnectivity200ResponseInner GetNetworkSmDeviceConnectivity(ctx, networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Returns historical connectivity data (whether a device is regularly checking in to Dashboard).



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
	deviceId := "deviceId_example" // string | Device ID
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetNetworkSmDeviceConnectivity(context.Background(), networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetNetworkSmDeviceConnectivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceConnectivity`: []GetNetworkSmDeviceConnectivity200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetNetworkSmDeviceConnectivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceConnectivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkSmDeviceConnectivity200ResponseInner**](GetNetworkSmDeviceConnectivity200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceDesktopLogs

> []GetNetworkSmDeviceDesktopLogs200ResponseInner GetNetworkSmDeviceDesktopLogs(ctx, networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return historical records of various Systems Manager network connection details for desktop devices.



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
	deviceId := "deviceId_example" // string | Device ID
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetNetworkSmDeviceDesktopLogs(context.Background(), networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetNetworkSmDeviceDesktopLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceDesktopLogs`: []GetNetworkSmDeviceDesktopLogs200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetNetworkSmDeviceDesktopLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceDesktopLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkSmDeviceDesktopLogs200ResponseInner**](GetNetworkSmDeviceDesktopLogs200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceDeviceCommandLogs

> []GetNetworkSmDeviceDeviceCommandLogs200ResponseInner GetNetworkSmDeviceDeviceCommandLogs(ctx, networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return historical records of commands sent to Systems Manager devices



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
	deviceId := "deviceId_example" // string | Device ID
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetNetworkSmDeviceDeviceCommandLogs(context.Background(), networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetNetworkSmDeviceDeviceCommandLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceDeviceCommandLogs`: []GetNetworkSmDeviceDeviceCommandLogs200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetNetworkSmDeviceDeviceCommandLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceDeviceCommandLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkSmDeviceDeviceCommandLogs200ResponseInner**](GetNetworkSmDeviceDeviceCommandLogs200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceDeviceProfiles

> []GetNetworkSmDeviceDeviceProfiles200ResponseInner GetNetworkSmDeviceDeviceProfiles(ctx, networkId, deviceId).Execute()

Get the installed profiles associated with a device



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
	deviceId := "deviceId_example" // string | Device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetNetworkSmDeviceDeviceProfiles(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetNetworkSmDeviceDeviceProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceDeviceProfiles`: []GetNetworkSmDeviceDeviceProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetNetworkSmDeviceDeviceProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceDeviceProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceDeviceProfiles200ResponseInner**](GetNetworkSmDeviceDeviceProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceNetworkAdapters

> []GetNetworkSmDeviceNetworkAdapters200ResponseInner GetNetworkSmDeviceNetworkAdapters(ctx, networkId, deviceId).Execute()

List the network adapters of a device



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
	deviceId := "deviceId_example" // string | Device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetNetworkSmDeviceNetworkAdapters(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetNetworkSmDeviceNetworkAdapters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceNetworkAdapters`: []GetNetworkSmDeviceNetworkAdapters200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetNetworkSmDeviceNetworkAdapters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceNetworkAdaptersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceNetworkAdapters200ResponseInner**](GetNetworkSmDeviceNetworkAdapters200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDevicePerformanceHistory

> []GetNetworkSmDevicePerformanceHistory200ResponseInner GetNetworkSmDevicePerformanceHistory(ctx, networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return historical records of various Systems Manager client metrics for desktop devices.



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
	deviceId := "deviceId_example" // string | Device ID
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetNetworkSmDevicePerformanceHistory(context.Background(), networkId, deviceId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetNetworkSmDevicePerformanceHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDevicePerformanceHistory`: []GetNetworkSmDevicePerformanceHistory200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetNetworkSmDevicePerformanceHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDevicePerformanceHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkSmDevicePerformanceHistory200ResponseInner**](GetNetworkSmDevicePerformanceHistory200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceRestrictions

> GetNetworkSmDeviceRestrictions200Response GetNetworkSmDeviceRestrictions(ctx, networkId, deviceId).Execute()

List the restrictions on a device



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
	deviceId := "deviceId_example" // string | Device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetNetworkSmDeviceRestrictions(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetNetworkSmDeviceRestrictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceRestrictions`: GetNetworkSmDeviceRestrictions200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetNetworkSmDeviceRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkSmDeviceRestrictions200Response**](GetNetworkSmDeviceRestrictions200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceSecurityCenters

> []GetNetworkSmDeviceSecurityCenters200ResponseInner GetNetworkSmDeviceSecurityCenters(ctx, networkId, deviceId).Execute()

List the security centers on a device



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
	deviceId := "deviceId_example" // string | Device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetNetworkSmDeviceSecurityCenters(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetNetworkSmDeviceSecurityCenters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceSecurityCenters`: []GetNetworkSmDeviceSecurityCenters200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetNetworkSmDeviceSecurityCenters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceSecurityCentersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceSecurityCenters200ResponseInner**](GetNetworkSmDeviceSecurityCenters200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceSoftwares

> []GetNetworkSmDeviceSoftwares200ResponseInner GetNetworkSmDeviceSoftwares(ctx, networkId, deviceId).Execute()

Get a list of softwares associated with a device



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
	deviceId := "deviceId_example" // string | Device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetNetworkSmDeviceSoftwares(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetNetworkSmDeviceSoftwares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceSoftwares`: []GetNetworkSmDeviceSoftwares200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetNetworkSmDeviceSoftwares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceSoftwaresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceSoftwares200ResponseInner**](GetNetworkSmDeviceSoftwares200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDeviceWlanLists

> []GetNetworkSmDeviceWlanLists200ResponseInner GetNetworkSmDeviceWlanLists(ctx, networkId, deviceId).Execute()

List the saved SSID names on a device



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
	deviceId := "deviceId_example" // string | Device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetNetworkSmDeviceWlanLists(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetNetworkSmDeviceWlanLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceWlanLists`: []GetNetworkSmDeviceWlanLists200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetNetworkSmDeviceWlanLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceWlanListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceWlanLists200ResponseInner**](GetNetworkSmDeviceWlanLists200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmDevices

> []GetNetworkSmDevices200ResponseInner GetNetworkSmDevices(ctx, networkId).Fields(fields).WifiMacs(wifiMacs).Serials(serials).Ids(ids).Uuids(uuids).SystemTypes(systemTypes).Scope(scope).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the devices enrolled in an SM network with various specified fields and filters



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
	fields := []string{"Inner_example"} // []string | Additional fields that will be displayed for each device.     The default fields are: id, name, tags, ssid, wifiMac, osName, systemModel, uuid, and serialNumber. The additional fields are: ip,     systemType, availableDeviceCapacity, kioskAppName, biosVersion, lastConnected, missingAppsCount, userSuppliedAddress, location, lastUser,     ownerEmail, ownerUsername, osBuild, publicIp, phoneNumber, diskInfoJson, deviceCapacity, isManaged, hadMdm, isSupervised, meid, imei, iccid,     simCarrierNetwork, cellularDataUsed, isHotspotEnabled, createdAt, batteryEstCharge, quarantined, avName, avRunning, asName, fwName,     isRooted, loginRequired, screenLockEnabled, screenLockDelay, autoLoginDisabled, autoTags, hasMdm, hasDesktopAgent, diskEncryptionEnabled,     hardwareEncryptionCaps, passCodeLock, usesHardwareKeystore, androidSecurityPatchVersion, cellular, and url. (optional)
	wifiMacs := []string{"Inner_example"} // []string | Filter devices by wifi mac(s). (optional)
	serials := []string{"Inner_example"} // []string | Filter devices by serial(s). (optional)
	ids := []string{"Inner_example"} // []string | Filter devices by id(s). (optional)
	uuids := []string{"Inner_example"} // []string | Filter devices by uuid(s). (optional)
	systemTypes := []string{"Inner_example"} // []string | Filter devices by system type(s). (optional)
	scope := []string{"Inner_example"} // []string | Specify a scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetNetworkSmDevices(context.Background(), networkId).Fields(fields).WifiMacs(wifiMacs).Serials(serials).Ids(ids).Uuids(uuids).SystemTypes(systemTypes).Scope(scope).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetNetworkSmDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDevices`: []GetNetworkSmDevices200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetNetworkSmDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | Additional fields that will be displayed for each device.     The default fields are: id, name, tags, ssid, wifiMac, osName, systemModel, uuid, and serialNumber. The additional fields are: ip,     systemType, availableDeviceCapacity, kioskAppName, biosVersion, lastConnected, missingAppsCount, userSuppliedAddress, location, lastUser,     ownerEmail, ownerUsername, osBuild, publicIp, phoneNumber, diskInfoJson, deviceCapacity, isManaged, hadMdm, isSupervised, meid, imei, iccid,     simCarrierNetwork, cellularDataUsed, isHotspotEnabled, createdAt, batteryEstCharge, quarantined, avName, avRunning, asName, fwName,     isRooted, loginRequired, screenLockEnabled, screenLockDelay, autoLoginDisabled, autoTags, hasMdm, hasDesktopAgent, diskEncryptionEnabled,     hardwareEncryptionCaps, passCodeLock, usesHardwareKeystore, androidSecurityPatchVersion, cellular, and url. | 
 **wifiMacs** | **[]string** | Filter devices by wifi mac(s). | 
 **serials** | **[]string** | Filter devices by serial(s). | 
 **ids** | **[]string** | Filter devices by id(s). | 
 **uuids** | **[]string** | Filter devices by uuid(s). | 
 **systemTypes** | **[]string** | Filter devices by system type(s). | 
 **scope** | **[]string** | Specify a scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkSmDevices200ResponseInner**](GetNetworkSmDevices200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessDevicesConnectionStats

> []GetDeviceWirelessConnectionStats200Response GetNetworkWirelessDevicesConnectionStats(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Execute()

Aggregated connectivity info for this network, grouped by node



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
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. (optional)
	band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information. (optional)
	ssid := int32(56) // int32 | Filter results by SSID (optional)
	vlan := int32(56) // int32 | Filter results by VLAN (optional)
	apTag := "apTag_example" // string | Filter results by AP Tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetNetworkWirelessDevicesConnectionStats(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetNetworkWirelessDevicesConnectionStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessDevicesConnectionStats`: []GetDeviceWirelessConnectionStats200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetNetworkWirelessDevicesConnectionStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessDevicesConnectionStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information. | 
 **ssid** | **int32** | Filter results by SSID | 
 **vlan** | **int32** | Filter results by VLAN | 
 **apTag** | **string** | Filter results by AP Tag | 

### Return type

[**[]GetDeviceWirelessConnectionStats200Response**](GetDeviceWirelessConnectionStats200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessDevicesLatencyStats

> []map[string]interface{} GetNetworkWirelessDevicesLatencyStats(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Fields(fields).Execute()

Aggregated latency info for this network, grouped by node



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
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. (optional)
	band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information. (optional)
	ssid := int32(56) // int32 | Filter results by SSID (optional)
	vlan := int32(56) // int32 | Filter results by VLAN (optional)
	apTag := "apTag_example" // string | Filter results by AP Tag (optional)
	fields := "fields_example" // string | Partial selection: If present, this call will return only the selected fields of [\"rawDistribution\", \"avg\"]. All fields will be returned by default. Selected fields must be entered as a comma separated string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetNetworkWirelessDevicesLatencyStats(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetNetworkWirelessDevicesLatencyStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessDevicesLatencyStats`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetNetworkWirelessDevicesLatencyStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessDevicesLatencyStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information. | 
 **ssid** | **int32** | Filter results by SSID | 
 **vlan** | **int32** | Filter results by VLAN | 
 **apTag** | **string** | Filter results by AP Tag | 
 **fields** | **string** | Partial selection: If present, this call will return only the selected fields of [\&quot;rawDistribution\&quot;, \&quot;avg\&quot;]. All fields will be returned by default. Selected fields must be entered as a comma separated string. | 

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


## GetOrganizationDevices

> []VmxNetworkDevicesClaim200Response GetOrganizationDevices(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).ConfigurationUpdatedAfter(configurationUpdatedAfter).NetworkIds(networkIds).ProductTypes(productTypes).Tags(tags).TagsFilterType(tagsFilterType).Name(name).Mac(mac).Serial(serial).Model(model).Macs(macs).Serials(serials).SensorMetrics(sensorMetrics).SensorAlertProfileIds(sensorAlertProfileIds).Models(models).Execute()

List the devices in an organization



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	configurationUpdatedAfter := "configurationUpdatedAfter_example" // string | Filter results by whether or not the device's configuration has been updated after the given timestamp (optional)
	networkIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by network. (optional)
	productTypes := []string{"ProductTypes_example"} // []string | Optional parameter to filter devices by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, and secureConnect. (optional)
	tags := []string{"Inner_example"} // []string | Optional parameter to filter devices by tags. (optional)
	tagsFilterType := "tagsFilterType_example" // string | Optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return networks which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)
	name := "name_example" // string | Optional parameter to filter devices by name. All returned devices will have a name that contains the search term or is an exact match. (optional)
	mac := "mac_example" // string | Optional parameter to filter devices by MAC address. All returned devices will have a MAC address that contains the search term or is an exact match. (optional)
	serial := "serial_example" // string | Optional parameter to filter devices by serial number. All returned devices will have a serial number that contains the search term or is an exact match. (optional)
	model := "model_example" // string | Optional parameter to filter devices by model. All returned devices will have a model that contains the search term or is an exact match. (optional)
	macs := []string{"Inner_example"} // []string | Optional parameter to filter devices by one or more MAC addresses. All returned devices will have a MAC address that is an exact match. (optional)
	serials := []string{"Inner_example"} // []string | Optional parameter to filter devices by one or more serial numbers. All returned devices will have a serial number that is an exact match. (optional)
	sensorMetrics := []string{"Inner_example"} // []string | Optional parameter to filter devices by the metrics that they provide. Only applies to sensor devices. (optional)
	sensorAlertProfileIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by the alert profiles that are bound to them. Only applies to sensor devices. (optional)
	models := []string{"Inner_example"} // []string | Optional parameter to filter devices by one or more models. All returned devices will have a model that is an exact match. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationDevices(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).ConfigurationUpdatedAfter(configurationUpdatedAfter).NetworkIds(networkIds).ProductTypes(productTypes).Tags(tags).TagsFilterType(tagsFilterType).Name(name).Mac(mac).Serial(serial).Model(model).Macs(macs).Serials(serials).SensorMetrics(sensorMetrics).SensorAlertProfileIds(sensorAlertProfileIds).Models(models).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationDevices`: []VmxNetworkDevicesClaim200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **configurationUpdatedAfter** | **string** | Filter results by whether or not the device&#39;s configuration has been updated after the given timestamp | 
 **networkIds** | **[]string** | Optional parameter to filter devices by network. | 
 **productTypes** | **[]string** | Optional parameter to filter devices by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, and secureConnect. | 
 **tags** | **[]string** | Optional parameter to filter devices by tags. | 
 **tagsFilterType** | **string** | Optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return networks which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 
 **name** | **string** | Optional parameter to filter devices by name. All returned devices will have a name that contains the search term or is an exact match. | 
 **mac** | **string** | Optional parameter to filter devices by MAC address. All returned devices will have a MAC address that contains the search term or is an exact match. | 
 **serial** | **string** | Optional parameter to filter devices by serial number. All returned devices will have a serial number that contains the search term or is an exact match. | 
 **model** | **string** | Optional parameter to filter devices by model. All returned devices will have a model that contains the search term or is an exact match. | 
 **macs** | **[]string** | Optional parameter to filter devices by one or more MAC addresses. All returned devices will have a MAC address that is an exact match. | 
 **serials** | **[]string** | Optional parameter to filter devices by one or more serial numbers. All returned devices will have a serial number that is an exact match. | 
 **sensorMetrics** | **[]string** | Optional parameter to filter devices by the metrics that they provide. Only applies to sensor devices. | 
 **sensorAlertProfileIds** | **[]string** | Optional parameter to filter devices by the alert profiles that are bound to them. Only applies to sensor devices. | 
 **models** | **[]string** | Optional parameter to filter devices by one or more models. All returned devices will have a model that is an exact match. | 

### Return type

[**[]VmxNetworkDevicesClaim200Response**](VmxNetworkDevicesClaim200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesAvailabilities

> []GetOrganizationDevicesAvailabilities200ResponseInner GetOrganizationDevicesAvailabilities(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Tags(tags).TagsFilterType(tagsFilterType).Execute()

List the availability information for devices in an organization



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	networkIds := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities by network ID. This filter uses multiple exact matches. (optional)
	productTypes := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities by device product types. This filter uses multiple exact matches. (optional)
	serials := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches. (optional)
	tags := []string{"Inner_example"} // []string | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). This filter uses multiple exact matches. (optional)
	tagsFilterType := "tagsFilterType_example" // string | An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationDevicesAvailabilities(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Tags(tags).TagsFilterType(tagsFilterType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationDevicesAvailabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationDevicesAvailabilities`: []GetOrganizationDevicesAvailabilities200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationDevicesAvailabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesAvailabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter device availabilities by network ID. This filter uses multiple exact matches. | 
 **productTypes** | **[]string** | Optional parameter to filter device availabilities by device product types. This filter uses multiple exact matches. | 
 **serials** | **[]string** | Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches. | 
 **tags** | **[]string** | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). This filter uses multiple exact matches. | 
 **tagsFilterType** | **string** | An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 

### Return type

[**[]GetOrganizationDevicesAvailabilities200ResponseInner**](GetOrganizationDevicesAvailabilities200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesAvailabilitiesChangeHistory

> []GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner GetOrganizationDevicesAvailabilitiesChangeHistory(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Serials(serials).ProductTypes(productTypes).NetworkIds(networkIds).Statuses(statuses).Execute()

List the availability history information for devices in an organization.



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)
	serials := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities history by device serial numbers (optional)
	productTypes := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities history by device product types (optional)
	networkIds := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities history by network IDs (optional)
	statuses := []string{"Statuses_example"} // []string | Optional parameter to filter device availabilities history by device statuses (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationDevicesAvailabilitiesChangeHistory(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Serials(serials).ProductTypes(productTypes).NetworkIds(networkIds).Statuses(statuses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationDevicesAvailabilitiesChangeHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationDevicesAvailabilitiesChangeHistory`: []GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationDevicesAvailabilitiesChangeHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesAvailabilitiesChangeHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 
 **serials** | **[]string** | Optional parameter to filter device availabilities history by device serial numbers | 
 **productTypes** | **[]string** | Optional parameter to filter device availabilities history by device product types | 
 **networkIds** | **[]string** | Optional parameter to filter device availabilities history by network IDs | 
 **statuses** | **[]string** | Optional parameter to filter device availabilities history by device statuses | 

### Return type

[**[]GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner**](GetOrganizationDevicesAvailabilitiesChangeHistory200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesOverviewByModel

> GetOrganizationDevicesOverviewByModel200Response GetOrganizationDevicesOverviewByModel(ctx, organizationId).Models(models).NetworkIds(networkIds).ProductTypes(productTypes).Execute()

Lists the count for each device model



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
	models := []string{"Inner_example"} // []string | Optional parameter to filter devices by one or more models. All returned devices will have a model that is an exact match. (optional)
	networkIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by networkId. (optional)
	productTypes := []string{"ProductTypes_example"} // []string | Optional parameter to filter device by device product types. This filter uses multiple exact matches. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationDevicesOverviewByModel(context.Background(), organizationId).Models(models).NetworkIds(networkIds).ProductTypes(productTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationDevicesOverviewByModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationDevicesOverviewByModel`: GetOrganizationDevicesOverviewByModel200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationDevicesOverviewByModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesOverviewByModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **models** | **[]string** | Optional parameter to filter devices by one or more models. All returned devices will have a model that is an exact match. | 
 **networkIds** | **[]string** | Optional parameter to filter devices by networkId. | 
 **productTypes** | **[]string** | Optional parameter to filter device by device product types. This filter uses multiple exact matches. | 

### Return type

[**GetOrganizationDevicesOverviewByModel200Response**](GetOrganizationDevicesOverviewByModel200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesPowerModulesStatusesByDevice

> []GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner GetOrganizationDevicesPowerModulesStatusesByDevice(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Tags(tags).TagsFilterType(tagsFilterType).Execute()

List the most recent status information for power modules in rackmount MX and MS devices that support them



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	networkIds := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities by network ID. This filter uses multiple exact matches. (optional)
	productTypes := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities by device product types. This filter uses multiple exact matches. (optional)
	serials := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches. (optional)
	tags := []string{"Inner_example"} // []string | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). This filter uses multiple exact matches. (optional)
	tagsFilterType := "tagsFilterType_example" // string | An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationDevicesPowerModulesStatusesByDevice(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Tags(tags).TagsFilterType(tagsFilterType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationDevicesPowerModulesStatusesByDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationDevicesPowerModulesStatusesByDevice`: []GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationDevicesPowerModulesStatusesByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesPowerModulesStatusesByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter device availabilities by network ID. This filter uses multiple exact matches. | 
 **productTypes** | **[]string** | Optional parameter to filter device availabilities by device product types. This filter uses multiple exact matches. | 
 **serials** | **[]string** | Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches. | 
 **tags** | **[]string** | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). This filter uses multiple exact matches. | 
 **tagsFilterType** | **string** | An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 

### Return type

[**[]GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner**](GetOrganizationDevicesPowerModulesStatusesByDevice200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesProvisioningStatuses

> []GetOrganizationDevicesProvisioningStatuses200ResponseInner GetOrganizationDevicesProvisioningStatuses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Status(status).Tags(tags).TagsFilterType(tagsFilterType).Execute()

List the provisioning statuses information for devices in an organization.



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	networkIds := []string{"Inner_example"} // []string | Optional parameter to filter device by network ID. This filter uses multiple exact matches. (optional)
	productTypes := []string{"Inner_example"} // []string | Optional parameter to filter device by device product types. This filter uses multiple exact matches. (optional)
	serials := []string{"Inner_example"} // []string | Optional parameter to filter device by device serial numbers. This filter uses multiple exact matches. (optional)
	status := "status_example" // string | An optional parameter to filter devices by the provisioning status. Accepted statuses: unprovisioned, incomplete, complete. (optional)
	tags := []string{"Inner_example"} // []string | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). This filter uses multiple exact matches. (optional)
	tagsFilterType := "tagsFilterType_example" // string | An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationDevicesProvisioningStatuses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Status(status).Tags(tags).TagsFilterType(tagsFilterType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationDevicesProvisioningStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationDevicesProvisioningStatuses`: []GetOrganizationDevicesProvisioningStatuses200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationDevicesProvisioningStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesProvisioningStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter device by network ID. This filter uses multiple exact matches. | 
 **productTypes** | **[]string** | Optional parameter to filter device by device product types. This filter uses multiple exact matches. | 
 **serials** | **[]string** | Optional parameter to filter device by device serial numbers. This filter uses multiple exact matches. | 
 **status** | **string** | An optional parameter to filter devices by the provisioning status. Accepted statuses: unprovisioned, incomplete, complete. | 
 **tags** | **[]string** | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). This filter uses multiple exact matches. | 
 **tagsFilterType** | **string** | An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 

### Return type

[**[]GetOrganizationDevicesProvisioningStatuses200ResponseInner**](GetOrganizationDevicesProvisioningStatuses200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesStatuses

> []GetOrganizationDevicesStatuses200ResponseInner GetOrganizationDevicesStatuses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Statuses(statuses).ProductTypes(productTypes).Models(models).Tags(tags).TagsFilterType(tagsFilterType).Execute()

List the status of every Meraki device in the organization



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	networkIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by network ids. (optional)
	serials := []string{"Inner_example"} // []string | Optional parameter to filter devices by serials. (optional)
	statuses := []string{"Statuses_example"} // []string | Optional parameter to filter devices by statuses. Valid statuses are [\"online\", \"alerting\", \"offline\", \"dormant\"]. (optional)
	productTypes := []string{"ProductTypes_example"} // []string | An optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, and secureConnect. (optional)
	models := []string{"Inner_example"} // []string | Optional parameter to filter devices by models. (optional)
	tags := []string{"Inner_example"} // []string | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). (optional)
	tagsFilterType := "tagsFilterType_example" // string | An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationDevicesStatuses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Statuses(statuses).ProductTypes(productTypes).Models(models).Tags(tags).TagsFilterType(tagsFilterType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationDevicesStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationDevicesStatuses`: []GetOrganizationDevicesStatuses200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationDevicesStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter devices by network ids. | 
 **serials** | **[]string** | Optional parameter to filter devices by serials. | 
 **statuses** | **[]string** | Optional parameter to filter devices by statuses. Valid statuses are [\&quot;online\&quot;, \&quot;alerting\&quot;, \&quot;offline\&quot;, \&quot;dormant\&quot;]. | 
 **productTypes** | **[]string** | An optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, and secureConnect. | 
 **models** | **[]string** | Optional parameter to filter devices by models. | 
 **tags** | **[]string** | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). | 
 **tagsFilterType** | **string** | An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 

### Return type

[**[]GetOrganizationDevicesStatuses200ResponseInner**](GetOrganizationDevicesStatuses200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesStatusesOverview

> GetOrganizationDevicesStatusesOverview200Response GetOrganizationDevicesStatusesOverview(ctx, organizationId).ProductTypes(productTypes).NetworkIds(networkIds).Execute()

Return an overview of current device statuses



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
	productTypes := []string{"ProductTypes_example"} // []string | An optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, and secureConnect. (optional)
	networkIds := []string{"Inner_example"} // []string | An optional parameter to filter device statuses by network. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationDevicesStatusesOverview(context.Background(), organizationId).ProductTypes(productTypes).NetworkIds(networkIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationDevicesStatusesOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationDevicesStatusesOverview`: GetOrganizationDevicesStatusesOverview200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationDevicesStatusesOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesStatusesOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productTypes** | **[]string** | An optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, and secureConnect. | 
 **networkIds** | **[]string** | An optional parameter to filter device statuses by network. | 

### Return type

[**GetOrganizationDevicesStatusesOverview200Response**](GetOrganizationDevicesStatusesOverview200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesUplinksAddressesByDevice

> []GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner GetOrganizationDevicesUplinksAddressesByDevice(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Tags(tags).TagsFilterType(tagsFilterType).Execute()

List the current uplink addresses for devices in an organization.



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	networkIds := []string{"Inner_example"} // []string | Optional parameter to filter device uplinks by network ID. This filter uses multiple exact matches. (optional)
	productTypes := []string{"Inner_example"} // []string | Optional parameter to filter device uplinks by device product types. This filter uses multiple exact matches. (optional)
	serials := []string{"Inner_example"} // []string | Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches. (optional)
	tags := []string{"Inner_example"} // []string | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). This filter uses multiple exact matches. (optional)
	tagsFilterType := "tagsFilterType_example" // string | An optional parameter of value 'withAnyTags' or 'withAllTags' to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, 'withAnyTags' will be selected. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationDevicesUplinksAddressesByDevice(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).ProductTypes(productTypes).Serials(serials).Tags(tags).TagsFilterType(tagsFilterType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationDevicesUplinksAddressesByDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationDevicesUplinksAddressesByDevice`: []GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationDevicesUplinksAddressesByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesUplinksAddressesByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter device uplinks by network ID. This filter uses multiple exact matches. | 
 **productTypes** | **[]string** | Optional parameter to filter device uplinks by device product types. This filter uses multiple exact matches. | 
 **serials** | **[]string** | Optional parameter to filter device availabilities by device serial numbers. This filter uses multiple exact matches. | 
 **tags** | **[]string** | An optional parameter to filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). This filter uses multiple exact matches. | 
 **tagsFilterType** | **string** | An optional parameter of value &#39;withAnyTags&#39; or &#39;withAllTags&#39; to indicate whether to return devices which contain ANY or ALL of the included tags. If no type is included, &#39;withAnyTags&#39; will be selected. | 

### Return type

[**[]GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner**](GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationDevicesUplinksLossAndLatency

> []GetOrganizationDevicesUplinksLossAndLatency200ResponseInner GetOrganizationDevicesUplinksLossAndLatency(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Uplink(uplink).Ip(ip).Execute()

Return the uplink loss and latency for every MX in the organization from at latest 2 minutes ago



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
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 60 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 5 minutes after t0. The latest possible time that t1 can be is 2 minutes into the past. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 5 minutes. The default is 5 minutes. (optional)
	uplink := "uplink_example" // string | Optional filter for a specific WAN uplink. Valid uplinks are wan1, wan2, wan3, cellular. Default will return all uplinks. (optional)
	ip := "ip_example" // string | Optional filter for a specific destination IP. Default will return all destination IPs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationDevicesUplinksLossAndLatency(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Uplink(uplink).Ip(ip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationDevicesUplinksLossAndLatency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationDevicesUplinksLossAndLatency`: []GetOrganizationDevicesUplinksLossAndLatency200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationDevicesUplinksLossAndLatency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationDevicesUplinksLossAndLatencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 60 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 5 minutes after t0. The latest possible time that t1 can be is 2 minutes into the past. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 5 minutes. The default is 5 minutes. | 
 **uplink** | **string** | Optional filter for a specific WAN uplink. Valid uplinks are wan1, wan2, wan3, cellular. Default will return all uplinks. | 
 **ip** | **string** | Optional filter for a specific destination IP. Default will return all destination IPs. | 

### Return type

[**[]GetOrganizationDevicesUplinksLossAndLatency200ResponseInner**](GetOrganizationDevicesUplinksLossAndLatency200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryDevice

> GetOrganizationInventoryDevices200ResponseInner GetOrganizationInventoryDevice(ctx, organizationId, serial).Execute()

Return a single device from the inventory of an organization



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
	serial := "serial_example" // string | Serial

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationInventoryDevice(context.Background(), organizationId, serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationInventoryDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationInventoryDevice`: GetOrganizationInventoryDevices200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationInventoryDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationInventoryDevices200ResponseInner**](GetOrganizationInventoryDevices200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryDevices

> []GetOrganizationInventoryDevices200ResponseInner GetOrganizationInventoryDevices(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).UsedState(usedState).Search(search).Macs(macs).NetworkIds(networkIds).Serials(serials).Models(models).OrderNumbers(orderNumbers).Tags(tags).TagsFilterType(tagsFilterType).ProductTypes(productTypes).Execute()

Return the device inventory for an organization



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	usedState := "usedState_example" // string | Filter results by used or unused inventory. Accepted values are 'used' or 'unused'. (optional)
	search := "search_example" // string | Search for devices in inventory based on serial number, mac address, or model. (optional)
	macs := []string{"Inner_example"} // []string | Search for devices in inventory based on mac addresses. (optional)
	networkIds := []string{"Inner_example"} // []string | Search for devices in inventory based on network ids. Use explicit 'null' value to get available devices only. (optional)
	serials := []string{"Inner_example"} // []string | Search for devices in inventory based on serials. (optional)
	models := []string{"Inner_example"} // []string | Search for devices in inventory based on model. (optional)
	orderNumbers := []string{"Inner_example"} // []string | Search for devices in inventory based on order numbers. (optional)
	tags := []string{"Inner_example"} // []string | Filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). (optional)
	tagsFilterType := "tagsFilterType_example" // string | To use with 'tags' parameter, to filter devices which contain ANY or ALL given tags. Accepted values are 'withAnyTags' or 'withAllTags', default is 'withAnyTags'. (optional)
	productTypes := []string{"ProductTypes_example"} // []string | Filter devices by product type. Accepted values are appliance, camera, cellularGateway, secureConnect, sensor, switch, systemsManager, wireless, and wirelessController. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationInventoryDevices(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).UsedState(usedState).Search(search).Macs(macs).NetworkIds(networkIds).Serials(serials).Models(models).OrderNumbers(orderNumbers).Tags(tags).TagsFilterType(tagsFilterType).ProductTypes(productTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationInventoryDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationInventoryDevices`: []GetOrganizationInventoryDevices200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationInventoryDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **usedState** | **string** | Filter results by used or unused inventory. Accepted values are &#39;used&#39; or &#39;unused&#39;. | 
 **search** | **string** | Search for devices in inventory based on serial number, mac address, or model. | 
 **macs** | **[]string** | Search for devices in inventory based on mac addresses. | 
 **networkIds** | **[]string** | Search for devices in inventory based on network ids. Use explicit &#39;null&#39; value to get available devices only. | 
 **serials** | **[]string** | Search for devices in inventory based on serials. | 
 **models** | **[]string** | Search for devices in inventory based on model. | 
 **orderNumbers** | **[]string** | Search for devices in inventory based on order numbers. | 
 **tags** | **[]string** | Filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). | 
 **tagsFilterType** | **string** | To use with &#39;tags&#39; parameter, to filter devices which contain ANY or ALL given tags. Accepted values are &#39;withAnyTags&#39; or &#39;withAllTags&#39;, default is &#39;withAnyTags&#39;. | 
 **productTypes** | **[]string** | Filter devices by product type. Accepted values are appliance, camera, cellularGateway, secureConnect, sensor, switch, systemsManager, wireless, and wirelessController. | 

### Return type

[**[]GetOrganizationInventoryDevices200ResponseInner**](GetOrganizationInventoryDevices200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryDevicesSwapsBulk

> CreateOrganizationInventoryDevicesSwapsBulk207Response GetOrganizationInventoryDevicesSwapsBulk(ctx, organizationId, id).Execute()

List of device swaps for a given request ID ({id}).



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
	id := "id_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationInventoryDevicesSwapsBulk(context.Background(), organizationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationInventoryDevicesSwapsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationInventoryDevicesSwapsBulk`: CreateOrganizationInventoryDevicesSwapsBulk207Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationInventoryDevicesSwapsBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryDevicesSwapsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateOrganizationInventoryDevicesSwapsBulk207Response**](CreateOrganizationInventoryDevicesSwapsBulk207Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummaryTopDevicesByUsage

> []GetOrganizationSummaryTopDevicesByUsage200ResponseInner GetOrganizationSummaryTopDevicesByUsage(ctx, organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()

Return metrics for organization's top 10 devices sorted by data usage over given time range



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
	t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 8 hours and be less than or equal to 186 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationSummaryTopDevicesByUsage(context.Background(), organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationSummaryTopDevicesByUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSummaryTopDevicesByUsage`: []GetOrganizationSummaryTopDevicesByUsage200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationSummaryTopDevicesByUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopDevicesByUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkTag** | **string** | Match result to an exact network tag | 
 **deviceTag** | **string** | Match result to an exact device tag | 
 **networkId** | **string** | Match result to an exact network id | 
 **quantity** | **int32** | Set number of desired results to return. Default is 10. | 
 **ssidName** | **string** | Filter results by ssid name | 
 **usageUplink** | **string** | Filter results by usage uplink | 
 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 8 hours and be less than or equal to 186 days. The default is 1 day. | 

### Return type

[**[]GetOrganizationSummaryTopDevicesByUsage200ResponseInner**](GetOrganizationSummaryTopDevicesByUsage200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummaryTopDevicesModelsByUsage

> []GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner GetOrganizationSummaryTopDevicesModelsByUsage(ctx, organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()

Return metrics for organization's top 10 device models sorted by data usage over given time range



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
	t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 8 hours and be less than or equal to 186 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationSummaryTopDevicesModelsByUsage(context.Background(), organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationSummaryTopDevicesModelsByUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSummaryTopDevicesModelsByUsage`: []GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationSummaryTopDevicesModelsByUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopDevicesModelsByUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkTag** | **string** | Match result to an exact network tag | 
 **deviceTag** | **string** | Match result to an exact device tag | 
 **networkId** | **string** | Match result to an exact network id | 
 **quantity** | **int32** | Set number of desired results to return. Default is 10. | 
 **ssidName** | **string** | Filter results by ssid name | 
 **usageUplink** | **string** | Filter results by usage uplink | 
 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 8 hours and be less than or equal to 186 days. The default is 1 day. | 

### Return type

[**[]GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner**](GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesChannelUtilizationByDevice

> []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner GetOrganizationWirelessDevicesChannelUtilizationByDevice(ctx, organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()

Get average channel utilization for all bands in a network, split by AP



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
	networkIds := []string{"Inner_example"} // []string | Filter results by network. (optional)
	serials := []string{"Inner_example"} // []string | Filter results by device. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 90 days. The default is 7 days. (optional)
	interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationWirelessDevicesChannelUtilizationByDevice(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationWirelessDevicesChannelUtilizationByDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationWirelessDevicesChannelUtilizationByDevice`: []GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationWirelessDevicesChannelUtilizationByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesChannelUtilizationByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter results by network. | 
 **serials** | **[]string** | Filter results by device. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 90 days. The default is 7 days. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. | 

### Return type

[**[]GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner**](GetOrganizationWirelessDevicesChannelUtilizationByDevice200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesChannelUtilizationByNetwork

> []GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner GetOrganizationWirelessDevicesChannelUtilizationByNetwork(ctx, organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()

Get average channel utilization across all bands for all networks in the organization



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
	networkIds := []string{"Inner_example"} // []string | Filter results by network. (optional)
	serials := []string{"Inner_example"} // []string | Filter results by device. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 90 days. The default is 7 days. (optional)
	interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationWirelessDevicesChannelUtilizationByNetwork(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationWirelessDevicesChannelUtilizationByNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationWirelessDevicesChannelUtilizationByNetwork`: []GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationWirelessDevicesChannelUtilizationByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesChannelUtilizationByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter results by network. | 
 **serials** | **[]string** | Filter results by device. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 90 days. The default is 7 days. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. | 

### Return type

[**[]GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner**](GetOrganizationWirelessDevicesChannelUtilizationByNetwork200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval

> []GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval(ctx, organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()

Get a time-series of average channel utilization for all bands, segmented by device.



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
	networkIds := []string{"Inner_example"} // []string | Filter results by network. (optional)
	serials := []string{"Inner_example"} // []string | Filter results by device. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
	interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval`: []GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter results by network. | 
 **serials** | **[]string** | Filter results by device. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. | 

### Return type

[**[]GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner**](GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval

> []GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval(ctx, organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()

Get a time-series of average channel utilization for all bands



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
	networkIds := []string{"Inner_example"} // []string | Filter results by network. (optional)
	serials := []string{"Inner_example"} // []string | Filter results by device. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
	interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval`: []GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter results by network. | 
 **serials** | **[]string** | Filter results by device. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. | 

### Return type

[**[]GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner**](GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesEthernetStatuses

> []GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner GetOrganizationWirelessDevicesEthernetStatuses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

List the most recent Ethernet link speed, duplex, aggregation and power mode and status information for wireless devices.



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	networkIds := []string{"Inner_example"} // []string | A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]=N_12345678&networkIds[]=L_3456 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationWirelessDevicesEthernetStatuses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationWirelessDevicesEthernetStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationWirelessDevicesEthernetStatuses`: []GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationWirelessDevicesEthernetStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesEthernetStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]&#x3D;N_12345678&amp;networkIds[]&#x3D;L_3456 | 

### Return type

[**[]GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner**](GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesPacketLossByClient

> []GetOrganizationWirelessDevicesPacketLossByClient200ResponseInner GetOrganizationWirelessDevicesPacketLossByClient(ctx, organizationId).NetworkIds(networkIds).Ssids(ssids).Bands(bands).Macs(macs).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()

Get average packet loss for the given timespan for all clients in the organization.



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
	networkIds := []string{"Inner_example"} // []string | Filter results by network. (optional)
	ssids := []int32{int32(123)} // []int32 | Filter results by SSID number. (optional)
	bands := []string{"Inner_example"} // []string | Filter results by band. Valid bands are: 2.4, 5, and 6. (optional)
	macs := []string{"Inner_example"} // []string | Filter results by client mac address(es). (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 90 days. The default is 7 days. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationWirelessDevicesPacketLossByClient(context.Background(), organizationId).NetworkIds(networkIds).Ssids(ssids).Bands(bands).Macs(macs).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationWirelessDevicesPacketLossByClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationWirelessDevicesPacketLossByClient`: []GetOrganizationWirelessDevicesPacketLossByClient200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationWirelessDevicesPacketLossByClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesPacketLossByClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter results by network. | 
 **ssids** | **[]int32** | Filter results by SSID number. | 
 **bands** | **[]string** | Filter results by band. Valid bands are: 2.4, 5, and 6. | 
 **macs** | **[]string** | Filter results by client mac address(es). | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 90 days. The default is 7 days. | 

### Return type

[**[]GetOrganizationWirelessDevicesPacketLossByClient200ResponseInner**](GetOrganizationWirelessDevicesPacketLossByClient200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesPacketLossByDevice

> []GetOrganizationWirelessDevicesPacketLossByDevice200ResponseInner GetOrganizationWirelessDevicesPacketLossByDevice(ctx, organizationId).NetworkIds(networkIds).Serials(serials).Ssids(ssids).Bands(bands).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()

Get average packet loss for the given timespan for all devices in the organization



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
	networkIds := []string{"Inner_example"} // []string | Filter results by network. (optional)
	serials := []string{"Inner_example"} // []string | Filter results by device. (optional)
	ssids := []int32{int32(123)} // []int32 | Filter results by SSID number. (optional)
	bands := []string{"Inner_example"} // []string | Filter results by band. Valid bands are: 2.4, 5, and 6. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 90 days. The default is 7 days. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationWirelessDevicesPacketLossByDevice(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).Ssids(ssids).Bands(bands).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationWirelessDevicesPacketLossByDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationWirelessDevicesPacketLossByDevice`: []GetOrganizationWirelessDevicesPacketLossByDevice200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationWirelessDevicesPacketLossByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesPacketLossByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter results by network. | 
 **serials** | **[]string** | Filter results by device. | 
 **ssids** | **[]int32** | Filter results by SSID number. | 
 **bands** | **[]string** | Filter results by band. Valid bands are: 2.4, 5, and 6. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 90 days. The default is 7 days. | 

### Return type

[**[]GetOrganizationWirelessDevicesPacketLossByDevice200ResponseInner**](GetOrganizationWirelessDevicesPacketLossByDevice200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesPacketLossByNetwork

> []GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner GetOrganizationWirelessDevicesPacketLossByNetwork(ctx, organizationId).NetworkIds(networkIds).Serials(serials).Ssids(ssids).Bands(bands).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()

Get average packet loss for the given timespan for all networks in the organization.



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
	networkIds := []string{"Inner_example"} // []string | Filter results by network. (optional)
	serials := []string{"Inner_example"} // []string | Filter results by device. (optional)
	ssids := []int32{int32(123)} // []int32 | Filter results by SSID number. (optional)
	bands := []string{"Inner_example"} // []string | Filter results by band. Valid bands are: 2.4, 5, and 6. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 90 days. The default is 7 days. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetOrganizationWirelessDevicesPacketLossByNetwork(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).Ssids(ssids).Bands(bands).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetOrganizationWirelessDevicesPacketLossByNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationWirelessDevicesPacketLossByNetwork`: []GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetOrganizationWirelessDevicesPacketLossByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesPacketLossByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter results by network. | 
 **serials** | **[]string** | Filter results by device. | 
 **ssids** | **[]int32** | Filter results by SSID number. | 
 **bands** | **[]string** | Filter results by band. Valid bands are: 2.4, 5, and 6. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 90 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 5 minutes and be less than or equal to 90 days. The default is 7 days. | 

### Return type

[**[]GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner**](GetOrganizationWirelessDevicesPacketLossByNetwork200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallNetworkSmDeviceApps

> map[string]interface{} InstallNetworkSmDeviceApps(ctx, networkId, deviceId).InstallNetworkSmDeviceAppsRequest(installNetworkSmDeviceAppsRequest).Execute()

Install applications on a device



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
	deviceId := "deviceId_example" // string | Device ID
	installNetworkSmDeviceAppsRequest := *openapiclient.NewInstallNetworkSmDeviceAppsRequest([]string{"AppIds_example"}) // InstallNetworkSmDeviceAppsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.InstallNetworkSmDeviceApps(context.Background(), networkId, deviceId).InstallNetworkSmDeviceAppsRequest(installNetworkSmDeviceAppsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.InstallNetworkSmDeviceApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InstallNetworkSmDeviceApps`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.InstallNetworkSmDeviceApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallNetworkSmDeviceAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **installNetworkSmDeviceAppsRequest** | [**InstallNetworkSmDeviceAppsRequest**](InstallNetworkSmDeviceAppsRequest.md) |  | 

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


## LockNetworkSmDevices

> CheckinNetworkSmDevices200Response LockNetworkSmDevices(ctx, networkId).LockNetworkSmDevicesRequest(lockNetworkSmDevicesRequest).Execute()

Lock a set of devices



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
	lockNetworkSmDevicesRequest := *openapiclient.NewLockNetworkSmDevicesRequest() // LockNetworkSmDevicesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.LockNetworkSmDevices(context.Background(), networkId).LockNetworkSmDevicesRequest(lockNetworkSmDevicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.LockNetworkSmDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LockNetworkSmDevices`: CheckinNetworkSmDevices200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.LockNetworkSmDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiLockNetworkSmDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lockNetworkSmDevicesRequest** | [**LockNetworkSmDevicesRequest**](LockNetworkSmDevicesRequest.md) |  | 

### Return type

[**CheckinNetworkSmDevices200Response**](CheckinNetworkSmDevices200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyNetworkSmDevicesTags

> []ModifyNetworkSmDevicesTags200ResponseInner ModifyNetworkSmDevicesTags(ctx, networkId).ModifyNetworkSmDevicesTagsRequest(modifyNetworkSmDevicesTagsRequest).Execute()

Add, delete, or update the tags of a set of devices



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
	modifyNetworkSmDevicesTagsRequest := *openapiclient.NewModifyNetworkSmDevicesTagsRequest([]string{"Tags_example"}, "UpdateAction_example") // ModifyNetworkSmDevicesTagsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.ModifyNetworkSmDevicesTags(context.Background(), networkId).ModifyNetworkSmDevicesTagsRequest(modifyNetworkSmDevicesTagsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.ModifyNetworkSmDevicesTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyNetworkSmDevicesTags`: []ModifyNetworkSmDevicesTags200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.ModifyNetworkSmDevicesTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyNetworkSmDevicesTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyNetworkSmDevicesTagsRequest** | [**ModifyNetworkSmDevicesTagsRequest**](ModifyNetworkSmDevicesTagsRequest.md) |  | 

### Return type

[**[]ModifyNetworkSmDevicesTags200ResponseInner**](ModifyNetworkSmDevicesTags200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveNetworkSmDevices

> MoveNetworkSmDevices200Response MoveNetworkSmDevices(ctx, networkId).MoveNetworkSmDevicesRequest(moveNetworkSmDevicesRequest).Execute()

Move a set of devices to a new network



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
	moveNetworkSmDevicesRequest := *openapiclient.NewMoveNetworkSmDevicesRequest("NewNetwork_example") // MoveNetworkSmDevicesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.MoveNetworkSmDevices(context.Background(), networkId).MoveNetworkSmDevicesRequest(moveNetworkSmDevicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.MoveNetworkSmDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveNetworkSmDevices`: MoveNetworkSmDevices200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.MoveNetworkSmDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveNetworkSmDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveNetworkSmDevicesRequest** | [**MoveNetworkSmDevicesRequest**](MoveNetworkSmDevicesRequest.md) |  | 

### Return type

[**MoveNetworkSmDevices200Response**](MoveNetworkSmDevices200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootDevice

> RebootDevice202Response RebootDevice(ctx, serial).Execute()

Reboot a device



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
	resp, r, err := apiClient.DevicesAPI.RebootDevice(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.RebootDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebootDevice`: RebootDevice202Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.RebootDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RebootDevice202Response**](RebootDevice202Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootNetworkSmDevices

> RebootNetworkSmDevices200Response RebootNetworkSmDevices(ctx, networkId).RebootNetworkSmDevicesRequest(rebootNetworkSmDevicesRequest).Execute()

Reboot a set of endpoints



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
	rebootNetworkSmDevicesRequest := *openapiclient.NewRebootNetworkSmDevicesRequest() // RebootNetworkSmDevicesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.RebootNetworkSmDevices(context.Background(), networkId).RebootNetworkSmDevicesRequest(rebootNetworkSmDevicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.RebootNetworkSmDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebootNetworkSmDevices`: RebootNetworkSmDevices200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.RebootNetworkSmDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootNetworkSmDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rebootNetworkSmDevicesRequest** | [**RebootNetworkSmDevicesRequest**](RebootNetworkSmDevicesRequest.md) |  | 

### Return type

[**RebootNetworkSmDevices200Response**](RebootNetworkSmDevices200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshNetworkSmDeviceDetails

> map[string]interface{} RefreshNetworkSmDeviceDetails(ctx, networkId, deviceId).Execute()

Refresh the details of a device



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
	deviceId := "deviceId_example" // string | Device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.RefreshNetworkSmDeviceDetails(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.RefreshNetworkSmDeviceDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshNetworkSmDeviceDetails`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.RefreshNetworkSmDeviceDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshNetworkSmDeviceDetailsRequest struct via the builder pattern


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
	r, err := apiClient.DevicesAPI.RemoveNetworkDevices(context.Background(), networkId).RemoveNetworkDevicesRequest(removeNetworkDevicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.RemoveNetworkDevices``: %v\n", err)
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


## ShutdownNetworkSmDevices

> RebootNetworkSmDevices200Response ShutdownNetworkSmDevices(ctx, networkId).ShutdownNetworkSmDevicesRequest(shutdownNetworkSmDevicesRequest).Execute()

Shutdown a set of endpoints



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
	shutdownNetworkSmDevicesRequest := *openapiclient.NewShutdownNetworkSmDevicesRequest() // ShutdownNetworkSmDevicesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.ShutdownNetworkSmDevices(context.Background(), networkId).ShutdownNetworkSmDevicesRequest(shutdownNetworkSmDevicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.ShutdownNetworkSmDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShutdownNetworkSmDevices`: RebootNetworkSmDevices200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.ShutdownNetworkSmDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShutdownNetworkSmDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shutdownNetworkSmDevicesRequest** | [**ShutdownNetworkSmDevicesRequest**](ShutdownNetworkSmDevicesRequest.md) |  | 

### Return type

[**RebootNetworkSmDevices200Response**](RebootNetworkSmDevices200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnenrollNetworkSmDevice

> UnenrollNetworkSmDevice200Response UnenrollNetworkSmDevice(ctx, networkId, deviceId).Execute()

Unenroll a device



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
	deviceId := "deviceId_example" // string | Device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.UnenrollNetworkSmDevice(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.UnenrollNetworkSmDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnenrollNetworkSmDevice`: UnenrollNetworkSmDevice200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.UnenrollNetworkSmDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnenrollNetworkSmDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UnenrollNetworkSmDevice200Response**](UnenrollNetworkSmDevice200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UninstallNetworkSmDeviceApps

> map[string]interface{} UninstallNetworkSmDeviceApps(ctx, networkId, deviceId).UninstallNetworkSmDeviceAppsRequest(uninstallNetworkSmDeviceAppsRequest).Execute()

Uninstall applications on a device



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
	deviceId := "deviceId_example" // string | Device ID
	uninstallNetworkSmDeviceAppsRequest := *openapiclient.NewUninstallNetworkSmDeviceAppsRequest([]string{"AppIds_example"}) // UninstallNetworkSmDeviceAppsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.UninstallNetworkSmDeviceApps(context.Background(), networkId, deviceId).UninstallNetworkSmDeviceAppsRequest(uninstallNetworkSmDeviceAppsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.UninstallNetworkSmDeviceApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UninstallNetworkSmDeviceApps`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.UninstallNetworkSmDeviceApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUninstallNetworkSmDeviceAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uninstallNetworkSmDeviceAppsRequest** | [**UninstallNetworkSmDeviceAppsRequest**](UninstallNetworkSmDeviceAppsRequest.md) |  | 

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


## UpdateDevice

> GetDevice200Response UpdateDevice(ctx, serial).UpdateDeviceRequest(updateDeviceRequest).Execute()

Update the attributes of a device



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
	updateDeviceRequest := *openapiclient.NewUpdateDeviceRequest() // UpdateDeviceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.UpdateDevice(context.Background(), serial).UpdateDeviceRequest(updateDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.UpdateDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDevice`: GetDevice200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.UpdateDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceRequest** | [**UpdateDeviceRequest**](UpdateDeviceRequest.md) |  | 

### Return type

[**GetDevice200Response**](GetDevice200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCellularSims

> GetDeviceCellularSims200Response UpdateDeviceCellularSims(ctx, serial).UpdateDeviceCellularSimsRequest(updateDeviceCellularSimsRequest).Execute()

Updates the SIM and APN configurations for a cellular device.



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
	updateDeviceCellularSimsRequest := *openapiclient.NewUpdateDeviceCellularSimsRequest() // UpdateDeviceCellularSimsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.UpdateDeviceCellularSims(context.Background(), serial).UpdateDeviceCellularSimsRequest(updateDeviceCellularSimsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.UpdateDeviceCellularSims``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceCellularSims`: GetDeviceCellularSims200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.UpdateDeviceCellularSims`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCellularSimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCellularSimsRequest** | [**UpdateDeviceCellularSimsRequest**](UpdateDeviceCellularSimsRequest.md) |  | 

### Return type

[**GetDeviceCellularSims200Response**](GetDeviceCellularSims200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceManagementInterface

> GetDeviceManagementInterface200Response UpdateDeviceManagementInterface(ctx, serial).UpdateDeviceManagementInterfaceRequest(updateDeviceManagementInterfaceRequest).Execute()

Update the management interface settings for a device



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
	updateDeviceManagementInterfaceRequest := *openapiclient.NewUpdateDeviceManagementInterfaceRequest() // UpdateDeviceManagementInterfaceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.UpdateDeviceManagementInterface(context.Background(), serial).UpdateDeviceManagementInterfaceRequest(updateDeviceManagementInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.UpdateDeviceManagementInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceManagementInterface`: GetDeviceManagementInterface200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.UpdateDeviceManagementInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceManagementInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceManagementInterfaceRequest** | [**UpdateDeviceManagementInterfaceRequest**](UpdateDeviceManagementInterfaceRequest.md) |  | 

### Return type

[**GetDeviceManagementInterface200Response**](GetDeviceManagementInterface200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSmDevicesFields

> []UpdateNetworkSmDevicesFields200ResponseInner UpdateNetworkSmDevicesFields(ctx, networkId).UpdateNetworkSmDevicesFieldsRequest(updateNetworkSmDevicesFieldsRequest).Execute()

Modify the fields of a device



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
	updateNetworkSmDevicesFieldsRequest := *openapiclient.NewUpdateNetworkSmDevicesFieldsRequest(*openapiclient.NewUpdateNetworkSmDevicesFieldsRequestDeviceFields()) // UpdateNetworkSmDevicesFieldsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.UpdateNetworkSmDevicesFields(context.Background(), networkId).UpdateNetworkSmDevicesFieldsRequest(updateNetworkSmDevicesFieldsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.UpdateNetworkSmDevicesFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSmDevicesFields`: []UpdateNetworkSmDevicesFields200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.UpdateNetworkSmDevicesFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSmDevicesFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSmDevicesFieldsRequest** | [**UpdateNetworkSmDevicesFieldsRequest**](UpdateNetworkSmDevicesFieldsRequest.md) |  | 

### Return type

[**[]UpdateNetworkSmDevicesFields200ResponseInner**](UpdateNetworkSmDevicesFields200ResponseInner.md)

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
	resp, r, err := apiClient.DevicesAPI.VmxNetworkDevicesClaim(context.Background(), networkId).VmxNetworkDevicesClaimRequest(vmxNetworkDevicesClaimRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.VmxNetworkDevicesClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VmxNetworkDevicesClaim`: VmxNetworkDevicesClaim200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.VmxNetworkDevicesClaim`: %v\n", resp)
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


## WipeNetworkSmDevices

> WipeNetworkSmDevices200Response WipeNetworkSmDevices(ctx, networkId).WipeNetworkSmDevicesRequest(wipeNetworkSmDevicesRequest).Execute()

Wipe a device



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
	wipeNetworkSmDevicesRequest := *openapiclient.NewWipeNetworkSmDevicesRequest() // WipeNetworkSmDevicesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.WipeNetworkSmDevices(context.Background(), networkId).WipeNetworkSmDevicesRequest(wipeNetworkSmDevicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.WipeNetworkSmDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WipeNetworkSmDevices`: WipeNetworkSmDevices200Response
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.WipeNetworkSmDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiWipeNetworkSmDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wipeNetworkSmDevicesRequest** | [**WipeNetworkSmDevicesRequest**](WipeNetworkSmDevicesRequest.md) |  | 

### Return type

[**WipeNetworkSmDevices200Response**](WipeNetworkSmDevices200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

