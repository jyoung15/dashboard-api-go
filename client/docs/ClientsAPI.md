# \ClientsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceClients**](ClientsAPI.md#GetDeviceClients) | **Get** /devices/{serial}/clients | List the clients of a device, up to a maximum of a month ago
[**GetNetworkApplianceClientSecurityEvents**](ClientsAPI.md#GetNetworkApplianceClientSecurityEvents) | **Get** /networks/{networkId}/appliance/clients/{clientId}/security/events | List the security events for a client
[**GetNetworkClient**](ClientsAPI.md#GetNetworkClient) | **Get** /networks/{networkId}/clients/{clientId} | Return the client associated with the given identifier
[**GetNetworkClientPolicy**](ClientsAPI.md#GetNetworkClientPolicy) | **Get** /networks/{networkId}/clients/{clientId}/policy | Return the policy assigned to a client on the network
[**GetNetworkClientSplashAuthorizationStatus**](ClientsAPI.md#GetNetworkClientSplashAuthorizationStatus) | **Get** /networks/{networkId}/clients/{clientId}/splashAuthorizationStatus | Return the splash authorization for a client, for each SSID they&#39;ve associated with through splash
[**GetNetworkClientTrafficHistory**](ClientsAPI.md#GetNetworkClientTrafficHistory) | **Get** /networks/{networkId}/clients/{clientId}/trafficHistory | Return the client&#39;s network traffic data over time
[**GetNetworkClientUsageHistory**](ClientsAPI.md#GetNetworkClientUsageHistory) | **Get** /networks/{networkId}/clients/{clientId}/usageHistory | Return the client&#39;s daily usage history
[**GetNetworkClients**](ClientsAPI.md#GetNetworkClients) | **Get** /networks/{networkId}/clients | List the clients that have used this network in the timespan
[**GetNetworkClientsApplicationUsage**](ClientsAPI.md#GetNetworkClientsApplicationUsage) | **Get** /networks/{networkId}/clients/applicationUsage | Return the application usage data for clients
[**GetNetworkClientsBandwidthUsageHistory**](ClientsAPI.md#GetNetworkClientsBandwidthUsageHistory) | **Get** /networks/{networkId}/clients/bandwidthUsageHistory | Returns a timeseries of total traffic consumption rates for all clients on a network within a given timespan, in megabits per second.
[**GetNetworkClientsOverview**](ClientsAPI.md#GetNetworkClientsOverview) | **Get** /networks/{networkId}/clients/overview | Return overview statistics for network clients
[**GetNetworkClientsUsageHistories**](ClientsAPI.md#GetNetworkClientsUsageHistories) | **Get** /networks/{networkId}/clients/usageHistories | Return the usage histories for clients
[**GetNetworkWirelessClientConnectionStats**](ClientsAPI.md#GetNetworkWirelessClientConnectionStats) | **Get** /networks/{networkId}/wireless/clients/{clientId}/connectionStats | Aggregated connectivity info for a given client on this network
[**GetNetworkWirelessClientConnectivityEvents**](ClientsAPI.md#GetNetworkWirelessClientConnectivityEvents) | **Get** /networks/{networkId}/wireless/clients/{clientId}/connectivityEvents | List the wireless connectivity events for a client within a network in the timespan.
[**GetNetworkWirelessClientLatencyHistory**](ClientsAPI.md#GetNetworkWirelessClientLatencyHistory) | **Get** /networks/{networkId}/wireless/clients/{clientId}/latencyHistory | Return the latency history for a client
[**GetNetworkWirelessClientLatencyStats**](ClientsAPI.md#GetNetworkWirelessClientLatencyStats) | **Get** /networks/{networkId}/wireless/clients/{clientId}/latencyStats | Aggregated latency info for a given client on this network
[**GetNetworkWirelessClientsConnectionStats**](ClientsAPI.md#GetNetworkWirelessClientsConnectionStats) | **Get** /networks/{networkId}/wireless/clients/connectionStats | Aggregated connectivity info for this network, grouped by clients
[**GetNetworkWirelessClientsLatencyStats**](ClientsAPI.md#GetNetworkWirelessClientsLatencyStats) | **Get** /networks/{networkId}/wireless/clients/latencyStats | Aggregated latency info for this network, grouped by clients
[**GetOrganizationClientsBandwidthUsageHistory**](ClientsAPI.md#GetOrganizationClientsBandwidthUsageHistory) | **Get** /organizations/{organizationId}/clients/bandwidthUsageHistory | Return data usage (in megabits per second) over time for all clients in the given organization within a given time range.
[**GetOrganizationClientsOverview**](ClientsAPI.md#GetOrganizationClientsOverview) | **Get** /organizations/{organizationId}/clients/overview | Return summary information around client data usage (in kb) across the given organization.
[**GetOrganizationClientsSearch**](ClientsAPI.md#GetOrganizationClientsSearch) | **Get** /organizations/{organizationId}/clients/search | Return the client details in an organization
[**GetOrganizationSummaryTopClientsByUsage**](ClientsAPI.md#GetOrganizationSummaryTopClientsByUsage) | **Get** /organizations/{organizationId}/summary/top/clients/byUsage | Return metrics for organization&#39;s top 10 clients by data usage (in mb) over given time range.
[**GetOrganizationSummaryTopClientsManufacturersByUsage**](ClientsAPI.md#GetOrganizationSummaryTopClientsManufacturersByUsage) | **Get** /organizations/{organizationId}/summary/top/clients/manufacturers/byUsage | Return metrics for organization&#39;s top clients by data usage (in mb) over given time range, grouped by manufacturer.
[**ProvisionNetworkClients**](ClientsAPI.md#ProvisionNetworkClients) | **Post** /networks/{networkId}/clients/provision | Provisions a client with a name and policy
[**UpdateNetworkClientPolicy**](ClientsAPI.md#UpdateNetworkClientPolicy) | **Put** /networks/{networkId}/clients/{clientId}/policy | Update the policy assigned to a client on the network
[**UpdateNetworkClientSplashAuthorizationStatus**](ClientsAPI.md#UpdateNetworkClientSplashAuthorizationStatus) | **Put** /networks/{networkId}/clients/{clientId}/splashAuthorizationStatus | Update a client&#39;s splash authorization



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
	resp, r, err := apiClient.ClientsAPI.GetDeviceClients(context.Background(), serial).T0(t0).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetDeviceClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceClients`: []GetDeviceClients200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetDeviceClients`: %v\n", resp)
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


## GetNetworkApplianceClientSecurityEvents

> []map[string]interface{} GetNetworkApplianceClientSecurityEvents(ctx, networkId, clientId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()

List the security events for a client



### Example

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
	t0 := "t0_example" // string | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 791 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 791 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 791 days. The default is 31 days. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	sortOrder := "sortOrder_example" // string | Sorted order of security events based on event detection time. Order options are 'ascending' or 'descending'. Default is ascending order. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.GetNetworkApplianceClientSecurityEvents(context.Background(), networkId, clientId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetNetworkApplianceClientSecurityEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceClientSecurityEvents`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetNetworkApplianceClientSecurityEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceClientSecurityEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **t0** | **string** | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 791 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 791 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 791 days. The default is 31 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of security events based on event detection time. Order options are &#39;ascending&#39; or &#39;descending&#39;. Default is ascending order. | 

### Return type

[**[]map[string]interface{}**](map.md)

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
	resp, r, err := apiClient.ClientsAPI.GetNetworkClient(context.Background(), networkId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetNetworkClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkClient`: GetNetworkClient200Response
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetNetworkClient`: %v\n", resp)
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
	resp, r, err := apiClient.ClientsAPI.GetNetworkClientPolicy(context.Background(), networkId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetNetworkClientPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkClientPolicy`: GetNetworkClientPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetNetworkClientPolicy`: %v\n", resp)
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
	resp, r, err := apiClient.ClientsAPI.GetNetworkClientSplashAuthorizationStatus(context.Background(), networkId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetNetworkClientSplashAuthorizationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkClientSplashAuthorizationStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetNetworkClientSplashAuthorizationStatus`: %v\n", resp)
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
	resp, r, err := apiClient.ClientsAPI.GetNetworkClientTrafficHistory(context.Background(), networkId, clientId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetNetworkClientTrafficHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkClientTrafficHistory`: []GetNetworkClientTrafficHistory200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetNetworkClientTrafficHistory`: %v\n", resp)
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
	resp, r, err := apiClient.ClientsAPI.GetNetworkClientUsageHistory(context.Background(), networkId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetNetworkClientUsageHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkClientUsageHistory`: []GetNetworkClientUsageHistory200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetNetworkClientUsageHistory`: %v\n", resp)
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
	resp, r, err := apiClient.ClientsAPI.GetNetworkClients(context.Background(), networkId).T0(t0).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Statuses(statuses).Ip(ip).Ip6(ip6).Ip6Local(ip6Local).Mac(mac).Os(os).PskGroup(pskGroup).Description(description).Vlan(vlan).NamedVlan(namedVlan).RecentDeviceConnections(recentDeviceConnections).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetNetworkClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkClients`: []GetNetworkClients200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetNetworkClients`: %v\n", resp)
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
	resp, r, err := apiClient.ClientsAPI.GetNetworkClientsApplicationUsage(context.Background(), networkId).Clients(clients).SsidNumber(ssidNumber).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetNetworkClientsApplicationUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkClientsApplicationUsage`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetNetworkClientsApplicationUsage`: %v\n", resp)
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
	resp, r, err := apiClient.ClientsAPI.GetNetworkClientsBandwidthUsageHistory(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetNetworkClientsBandwidthUsageHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkClientsBandwidthUsageHistory`: []GetNetworkClientsBandwidthUsageHistory200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetNetworkClientsBandwidthUsageHistory`: %v\n", resp)
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
	resp, r, err := apiClient.ClientsAPI.GetNetworkClientsOverview(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetNetworkClientsOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkClientsOverview`: GetNetworkClientsOverview200Response
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetNetworkClientsOverview`: %v\n", resp)
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
	resp, r, err := apiClient.ClientsAPI.GetNetworkClientsUsageHistories(context.Background(), networkId).Clients(clients).SsidNumber(ssidNumber).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetNetworkClientsUsageHistories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkClientsUsageHistories`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetNetworkClientsUsageHistories`: %v\n", resp)
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


## GetNetworkWirelessClientConnectionStats

> GetNetworkWirelessClientConnectionStats200Response GetNetworkWirelessClientConnectionStats(ctx, networkId, clientId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Execute()

Aggregated connectivity info for a given client on this network



### Example

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
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. (optional)
	band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information. (optional)
	ssid := int32(56) // int32 | Filter results by SSID (optional)
	vlan := int32(56) // int32 | Filter results by VLAN (optional)
	apTag := "apTag_example" // string | Filter results by AP Tag (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.GetNetworkWirelessClientConnectionStats(context.Background(), networkId, clientId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetNetworkWirelessClientConnectionStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessClientConnectionStats`: GetNetworkWirelessClientConnectionStats200Response
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetNetworkWirelessClientConnectionStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessClientConnectionStatsRequest struct via the builder pattern


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

[**GetNetworkWirelessClientConnectionStats200Response**](GetNetworkWirelessClientConnectionStats200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessClientConnectivityEvents

> []GetNetworkWirelessClientConnectivityEvents200ResponseInner GetNetworkWirelessClientConnectivityEvents(ctx, networkId, clientId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Types(types).Band(band).SsidNumber(ssidNumber).IncludedSeverities(includedSeverities).DeviceSerial(deviceSerial).Execute()

List the wireless connectivity events for a client within a network in the timespan.



### Example

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
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)
	types := []string{"Types_example"} // []string | A list of event types to include. If not specified, events of all types will be returned. Valid types are 'assoc', 'disassoc', 'auth', 'deauth', 'dns', 'dhcp', 'roam', 'connection' and/or 'sticky'. (optional)
	band := "band_example" // string | Filter results by band. Valid bands are '2.4', '5' or '6'. (optional)
	ssidNumber := int32(56) // int32 | Filter results by SSID. If not specified, events for all SSIDs will be returned. (optional)
	includedSeverities := []string{"IncludedSeverities_example"} // []string | A list of severities to include. If not specified, events of all severities will be returned. Valid severities are 'good', 'info', 'warn' and/or 'bad'. (optional)
	deviceSerial := "deviceSerial_example" // string | Filter results by an AP's serial number. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.GetNetworkWirelessClientConnectivityEvents(context.Background(), networkId, clientId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Types(types).Band(band).SsidNumber(ssidNumber).IncludedSeverities(includedSeverities).DeviceSerial(deviceSerial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetNetworkWirelessClientConnectivityEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessClientConnectivityEvents`: []GetNetworkWirelessClientConnectivityEvents200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetNetworkWirelessClientConnectivityEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessClientConnectivityEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 
 **types** | **[]string** | A list of event types to include. If not specified, events of all types will be returned. Valid types are &#39;assoc&#39;, &#39;disassoc&#39;, &#39;auth&#39;, &#39;deauth&#39;, &#39;dns&#39;, &#39;dhcp&#39;, &#39;roam&#39;, &#39;connection&#39; and/or &#39;sticky&#39;. | 
 **band** | **string** | Filter results by band. Valid bands are &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;. | 
 **ssidNumber** | **int32** | Filter results by SSID. If not specified, events for all SSIDs will be returned. | 
 **includedSeverities** | **[]string** | A list of severities to include. If not specified, events of all severities will be returned. Valid severities are &#39;good&#39;, &#39;info&#39;, &#39;warn&#39; and/or &#39;bad&#39;. | 
 **deviceSerial** | **string** | Filter results by an AP&#39;s serial number. | 

### Return type

[**[]GetNetworkWirelessClientConnectivityEvents200ResponseInner**](GetNetworkWirelessClientConnectivityEvents200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessClientLatencyHistory

> []GetNetworkWirelessClientLatencyHistory200ResponseInner GetNetworkWirelessClientLatencyHistory(ctx, networkId, clientId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Execute()

Return the latency history for a client



### Example

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
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 791 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 791 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 791 days. The default is 1 day. (optional)
	resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 86400. The default is 86400. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.GetNetworkWirelessClientLatencyHistory(context.Background(), networkId, clientId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetNetworkWirelessClientLatencyHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessClientLatencyHistory`: []GetNetworkWirelessClientLatencyHistory200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetNetworkWirelessClientLatencyHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessClientLatencyHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 791 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 791 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 791 days. The default is 1 day. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 86400. The default is 86400. | 

### Return type

[**[]GetNetworkWirelessClientLatencyHistory200ResponseInner**](GetNetworkWirelessClientLatencyHistory200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessClientLatencyStats

> map[string]interface{} GetNetworkWirelessClientLatencyStats(ctx, networkId, clientId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Fields(fields).Execute()

Aggregated latency info for a given client on this network



### Example

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
	resp, r, err := apiClient.ClientsAPI.GetNetworkWirelessClientLatencyStats(context.Background(), networkId, clientId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetNetworkWirelessClientLatencyStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessClientLatencyStats`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetNetworkWirelessClientLatencyStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessClientLatencyStatsRequest struct via the builder pattern


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

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessClientsConnectionStats

> []map[string]interface{} GetNetworkWirelessClientsConnectionStats(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Execute()

Aggregated connectivity info for this network, grouped by clients



### Example

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
	resp, r, err := apiClient.ClientsAPI.GetNetworkWirelessClientsConnectionStats(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetNetworkWirelessClientsConnectionStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessClientsConnectionStats`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetNetworkWirelessClientsConnectionStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessClientsConnectionStatsRequest struct via the builder pattern


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

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessClientsLatencyStats

> []map[string]interface{} GetNetworkWirelessClientsLatencyStats(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Fields(fields).Execute()

Aggregated latency info for this network, grouped by clients



### Example

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
	resp, r, err := apiClient.ClientsAPI.GetNetworkWirelessClientsLatencyStats(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetNetworkWirelessClientsLatencyStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessClientsLatencyStats`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetNetworkWirelessClientsLatencyStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessClientsLatencyStatsRequest struct via the builder pattern


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


## GetOrganizationClientsBandwidthUsageHistory

> []GetOrganizationClientsBandwidthUsageHistory200ResponseInner GetOrganizationClientsBandwidthUsageHistory(ctx, organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()

Return data usage (in megabits per second) over time for all clients in the given organization within a given time range.



### Example

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
	ssidName := "ssidName_example" // string | Filter results by ssid name (optional)
	usageUplink := "usageUplink_example" // string | Filter results by usage uplink (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 186 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.GetOrganizationClientsBandwidthUsageHistory(context.Background(), organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetOrganizationClientsBandwidthUsageHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationClientsBandwidthUsageHistory`: []GetOrganizationClientsBandwidthUsageHistory200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetOrganizationClientsBandwidthUsageHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationClientsBandwidthUsageHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkTag** | **string** | Match result to an exact network tag | 
 **deviceTag** | **string** | Match result to an exact device tag | 
 **networkId** | **string** | Match result to an exact network id | 
 **ssidName** | **string** | Filter results by ssid name | 
 **usageUplink** | **string** | Filter results by usage uplink | 
 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 186 days. The default is 1 day. | 

### Return type

[**[]GetOrganizationClientsBandwidthUsageHistory200ResponseInner**](GetOrganizationClientsBandwidthUsageHistory200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationClientsOverview

> GetOrganizationClientsOverview200Response GetOrganizationClientsOverview(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Return summary information around client data usage (in kb) across the given organization.



### Example

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
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.GetOrganizationClientsOverview(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetOrganizationClientsOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationClientsOverview`: GetOrganizationClientsOverview200Response
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetOrganizationClientsOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationClientsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

[**GetOrganizationClientsOverview200Response**](GetOrganizationClientsOverview200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationClientsSearch

> GetOrganizationClientsSearch200Response GetOrganizationClientsSearch(ctx, organizationId).Mac(mac).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return the client details in an organization



### Example

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
	mac := "mac_example" // string | The MAC address of the client. Required.
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 5. Default is 5. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.GetOrganizationClientsSearch(context.Background(), organizationId).Mac(mac).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetOrganizationClientsSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationClientsSearch`: GetOrganizationClientsSearch200Response
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetOrganizationClientsSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationClientsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **string** | The MAC address of the client. Required. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 5. Default is 5. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**GetOrganizationClientsSearch200Response**](GetOrganizationClientsSearch200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummaryTopClientsByUsage

> []GetOrganizationSummaryTopClientsByUsage200ResponseInner GetOrganizationSummaryTopClientsByUsage(ctx, organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()

Return metrics for organization's top 10 clients by data usage (in mb) over given time range.



### Example

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
	resp, r, err := apiClient.ClientsAPI.GetOrganizationSummaryTopClientsByUsage(context.Background(), organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetOrganizationSummaryTopClientsByUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSummaryTopClientsByUsage`: []GetOrganizationSummaryTopClientsByUsage200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetOrganizationSummaryTopClientsByUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopClientsByUsageRequest struct via the builder pattern


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

[**[]GetOrganizationSummaryTopClientsByUsage200ResponseInner**](GetOrganizationSummaryTopClientsByUsage200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummaryTopClientsManufacturersByUsage

> []GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner GetOrganizationSummaryTopClientsManufacturersByUsage(ctx, organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()

Return metrics for organization's top clients by data usage (in mb) over given time range, grouped by manufacturer.



### Example

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
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 186 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.GetOrganizationSummaryTopClientsManufacturersByUsage(context.Background(), organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetOrganizationSummaryTopClientsManufacturersByUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSummaryTopClientsManufacturersByUsage`: []GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetOrganizationSummaryTopClientsManufacturersByUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopClientsManufacturersByUsageRequest struct via the builder pattern


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
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 186 days. The default is 1 day. | 

### Return type

[**[]GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner**](GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner.md)

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
	resp, r, err := apiClient.ClientsAPI.ProvisionNetworkClients(context.Background(), networkId).ProvisionNetworkClientsRequest(provisionNetworkClientsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.ProvisionNetworkClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProvisionNetworkClients`: ProvisionNetworkClients201Response
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.ProvisionNetworkClients`: %v\n", resp)
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
	resp, r, err := apiClient.ClientsAPI.UpdateNetworkClientPolicy(context.Background(), networkId, clientId).UpdateNetworkClientPolicyRequest(updateNetworkClientPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.UpdateNetworkClientPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkClientPolicy`: GetNetworkClientPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.UpdateNetworkClientPolicy`: %v\n", resp)
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
	resp, r, err := apiClient.ClientsAPI.UpdateNetworkClientSplashAuthorizationStatus(context.Background(), networkId, clientId).UpdateNetworkClientSplashAuthorizationStatusRequest(updateNetworkClientSplashAuthorizationStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.UpdateNetworkClientSplashAuthorizationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkClientSplashAuthorizationStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.UpdateNetworkClientSplashAuthorizationStatus`: %v\n", resp)
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

