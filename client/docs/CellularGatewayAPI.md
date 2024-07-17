# \CellularGatewayAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceCellularGatewayLan**](CellularGatewayAPI.md#GetDeviceCellularGatewayLan) | **Get** /devices/{serial}/cellularGateway/lan | Show the LAN Settings of a MG
[**GetDeviceCellularGatewayPortForwardingRules**](CellularGatewayAPI.md#GetDeviceCellularGatewayPortForwardingRules) | **Get** /devices/{serial}/cellularGateway/portForwardingRules | Returns the port forwarding rules for a single MG.
[**GetNetworkCellularGatewayConnectivityMonitoringDestinations**](CellularGatewayAPI.md#GetNetworkCellularGatewayConnectivityMonitoringDestinations) | **Get** /networks/{networkId}/cellularGateway/connectivityMonitoringDestinations | Return the connectivity testing destinations for an MG network
[**GetNetworkCellularGatewayDhcp**](CellularGatewayAPI.md#GetNetworkCellularGatewayDhcp) | **Get** /networks/{networkId}/cellularGateway/dhcp | List common DHCP settings of MGs
[**GetNetworkCellularGatewaySubnetPool**](CellularGatewayAPI.md#GetNetworkCellularGatewaySubnetPool) | **Get** /networks/{networkId}/cellularGateway/subnetPool | Return the subnet pool and mask configured for MGs in the network.
[**GetNetworkCellularGatewayUplink**](CellularGatewayAPI.md#GetNetworkCellularGatewayUplink) | **Get** /networks/{networkId}/cellularGateway/uplink | Returns the uplink settings for your MG network.
[**GetOrganizationCellularGatewayUplinkStatuses**](CellularGatewayAPI.md#GetOrganizationCellularGatewayUplinkStatuses) | **Get** /organizations/{organizationId}/cellularGateway/uplink/statuses | List the uplink status of every Meraki MG cellular gateway in the organization
[**UpdateDeviceCellularGatewayLan**](CellularGatewayAPI.md#UpdateDeviceCellularGatewayLan) | **Put** /devices/{serial}/cellularGateway/lan | Update the LAN Settings for a single MG.
[**UpdateDeviceCellularGatewayPortForwardingRules**](CellularGatewayAPI.md#UpdateDeviceCellularGatewayPortForwardingRules) | **Put** /devices/{serial}/cellularGateway/portForwardingRules | Updates the port forwarding rules for a single MG.
[**UpdateNetworkCellularGatewayConnectivityMonitoringDestinations**](CellularGatewayAPI.md#UpdateNetworkCellularGatewayConnectivityMonitoringDestinations) | **Put** /networks/{networkId}/cellularGateway/connectivityMonitoringDestinations | Update the connectivity testing destinations for an MG network
[**UpdateNetworkCellularGatewayDhcp**](CellularGatewayAPI.md#UpdateNetworkCellularGatewayDhcp) | **Put** /networks/{networkId}/cellularGateway/dhcp | Update common DHCP settings of MGs
[**UpdateNetworkCellularGatewaySubnetPool**](CellularGatewayAPI.md#UpdateNetworkCellularGatewaySubnetPool) | **Put** /networks/{networkId}/cellularGateway/subnetPool | Update the subnet pool and mask configuration for MGs in the network.
[**UpdateNetworkCellularGatewayUplink**](CellularGatewayAPI.md#UpdateNetworkCellularGatewayUplink) | **Put** /networks/{networkId}/cellularGateway/uplink | Updates the uplink settings for your MG network.



## GetDeviceCellularGatewayLan

> GetDeviceCellularGatewayLan200Response GetDeviceCellularGatewayLan(ctx, serial).Execute()

Show the LAN Settings of a MG



### Example

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
	resp, r, err := apiClient.CellularGatewayAPI.GetDeviceCellularGatewayLan(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayAPI.GetDeviceCellularGatewayLan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCellularGatewayLan`: GetDeviceCellularGatewayLan200Response
	fmt.Fprintf(os.Stdout, "Response from `CellularGatewayAPI.GetDeviceCellularGatewayLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCellularGatewayLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceCellularGatewayLan200Response**](GetDeviceCellularGatewayLan200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceCellularGatewayPortForwardingRules

> GetDeviceCellularGatewayPortForwardingRules200Response GetDeviceCellularGatewayPortForwardingRules(ctx, serial).Execute()

Returns the port forwarding rules for a single MG.



### Example

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
	resp, r, err := apiClient.CellularGatewayAPI.GetDeviceCellularGatewayPortForwardingRules(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayAPI.GetDeviceCellularGatewayPortForwardingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCellularGatewayPortForwardingRules`: GetDeviceCellularGatewayPortForwardingRules200Response
	fmt.Fprintf(os.Stdout, "Response from `CellularGatewayAPI.GetDeviceCellularGatewayPortForwardingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCellularGatewayPortForwardingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceCellularGatewayPortForwardingRules200Response**](GetDeviceCellularGatewayPortForwardingRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkCellularGatewayConnectivityMonitoringDestinations

> GetNetworkApplianceConnectivityMonitoringDestinations200Response GetNetworkCellularGatewayConnectivityMonitoringDestinations(ctx, networkId).Execute()

Return the connectivity testing destinations for an MG network



### Example

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
	resp, r, err := apiClient.CellularGatewayAPI.GetNetworkCellularGatewayConnectivityMonitoringDestinations(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayAPI.GetNetworkCellularGatewayConnectivityMonitoringDestinations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkCellularGatewayConnectivityMonitoringDestinations`: GetNetworkApplianceConnectivityMonitoringDestinations200Response
	fmt.Fprintf(os.Stdout, "Response from `CellularGatewayAPI.GetNetworkCellularGatewayConnectivityMonitoringDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCellularGatewayConnectivityMonitoringDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceConnectivityMonitoringDestinations200Response**](GetNetworkApplianceConnectivityMonitoringDestinations200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkCellularGatewayDhcp

> GetNetworkCellularGatewayDhcp200Response GetNetworkCellularGatewayDhcp(ctx, networkId).Execute()

List common DHCP settings of MGs



### Example

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
	resp, r, err := apiClient.CellularGatewayAPI.GetNetworkCellularGatewayDhcp(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayAPI.GetNetworkCellularGatewayDhcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkCellularGatewayDhcp`: GetNetworkCellularGatewayDhcp200Response
	fmt.Fprintf(os.Stdout, "Response from `CellularGatewayAPI.GetNetworkCellularGatewayDhcp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCellularGatewayDhcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkCellularGatewayDhcp200Response**](GetNetworkCellularGatewayDhcp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkCellularGatewaySubnetPool

> GetNetworkCellularGatewaySubnetPool200Response GetNetworkCellularGatewaySubnetPool(ctx, networkId).Execute()

Return the subnet pool and mask configured for MGs in the network.



### Example

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
	resp, r, err := apiClient.CellularGatewayAPI.GetNetworkCellularGatewaySubnetPool(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayAPI.GetNetworkCellularGatewaySubnetPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkCellularGatewaySubnetPool`: GetNetworkCellularGatewaySubnetPool200Response
	fmt.Fprintf(os.Stdout, "Response from `CellularGatewayAPI.GetNetworkCellularGatewaySubnetPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCellularGatewaySubnetPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkCellularGatewaySubnetPool200Response**](GetNetworkCellularGatewaySubnetPool200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkCellularGatewayUplink

> GetNetworkCellularGatewayUplink200Response GetNetworkCellularGatewayUplink(ctx, networkId).Execute()

Returns the uplink settings for your MG network.



### Example

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
	resp, r, err := apiClient.CellularGatewayAPI.GetNetworkCellularGatewayUplink(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayAPI.GetNetworkCellularGatewayUplink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkCellularGatewayUplink`: GetNetworkCellularGatewayUplink200Response
	fmt.Fprintf(os.Stdout, "Response from `CellularGatewayAPI.GetNetworkCellularGatewayUplink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCellularGatewayUplinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkCellularGatewayUplink200Response**](GetNetworkCellularGatewayUplink200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCellularGatewayUplinkStatuses

> []GetOrganizationCellularGatewayUplinkStatuses200ResponseInner GetOrganizationCellularGatewayUplinkStatuses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Iccids(iccids).Execute()

List the uplink status of every Meraki MG cellular gateway in the organization



### Example

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
	networkIds := []string{"Inner_example"} // []string | A list of network IDs. The returned devices will be filtered to only include these networks. (optional)
	serials := []string{"Inner_example"} // []string | A list of serial numbers. The returned devices will be filtered to only include these serials. (optional)
	iccids := []string{"Inner_example"} // []string | A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CellularGatewayAPI.GetOrganizationCellularGatewayUplinkStatuses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Iccids(iccids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayAPI.GetOrganizationCellularGatewayUplinkStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationCellularGatewayUplinkStatuses`: []GetOrganizationCellularGatewayUplinkStatuses200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CellularGatewayAPI.GetOrganizationCellularGatewayUplinkStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCellularGatewayUplinkStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | A list of network IDs. The returned devices will be filtered to only include these networks. | 
 **serials** | **[]string** | A list of serial numbers. The returned devices will be filtered to only include these serials. | 
 **iccids** | **[]string** | A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs. | 

### Return type

[**[]GetOrganizationCellularGatewayUplinkStatuses200ResponseInner**](GetOrganizationCellularGatewayUplinkStatuses200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCellularGatewayLan

> GetDeviceCellularGatewayLan200Response UpdateDeviceCellularGatewayLan(ctx, serial).UpdateDeviceCellularGatewayLanRequest(updateDeviceCellularGatewayLanRequest).Execute()

Update the LAN Settings for a single MG.



### Example

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
	updateDeviceCellularGatewayLanRequest := *openapiclient.NewUpdateDeviceCellularGatewayLanRequest() // UpdateDeviceCellularGatewayLanRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CellularGatewayAPI.UpdateDeviceCellularGatewayLan(context.Background(), serial).UpdateDeviceCellularGatewayLanRequest(updateDeviceCellularGatewayLanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayAPI.UpdateDeviceCellularGatewayLan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceCellularGatewayLan`: GetDeviceCellularGatewayLan200Response
	fmt.Fprintf(os.Stdout, "Response from `CellularGatewayAPI.UpdateDeviceCellularGatewayLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCellularGatewayLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCellularGatewayLanRequest** | [**UpdateDeviceCellularGatewayLanRequest**](UpdateDeviceCellularGatewayLanRequest.md) |  | 

### Return type

[**GetDeviceCellularGatewayLan200Response**](GetDeviceCellularGatewayLan200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCellularGatewayPortForwardingRules

> GetDeviceCellularGatewayPortForwardingRules200Response UpdateDeviceCellularGatewayPortForwardingRules(ctx, serial).UpdateDeviceCellularGatewayPortForwardingRulesRequest(updateDeviceCellularGatewayPortForwardingRulesRequest).Execute()

Updates the port forwarding rules for a single MG.



### Example

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
	updateDeviceCellularGatewayPortForwardingRulesRequest := *openapiclient.NewUpdateDeviceCellularGatewayPortForwardingRulesRequest() // UpdateDeviceCellularGatewayPortForwardingRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CellularGatewayAPI.UpdateDeviceCellularGatewayPortForwardingRules(context.Background(), serial).UpdateDeviceCellularGatewayPortForwardingRulesRequest(updateDeviceCellularGatewayPortForwardingRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayAPI.UpdateDeviceCellularGatewayPortForwardingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceCellularGatewayPortForwardingRules`: GetDeviceCellularGatewayPortForwardingRules200Response
	fmt.Fprintf(os.Stdout, "Response from `CellularGatewayAPI.UpdateDeviceCellularGatewayPortForwardingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCellularGatewayPortForwardingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCellularGatewayPortForwardingRulesRequest** | [**UpdateDeviceCellularGatewayPortForwardingRulesRequest**](UpdateDeviceCellularGatewayPortForwardingRulesRequest.md) |  | 

### Return type

[**GetDeviceCellularGatewayPortForwardingRules200Response**](GetDeviceCellularGatewayPortForwardingRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkCellularGatewayConnectivityMonitoringDestinations

> GetNetworkApplianceConnectivityMonitoringDestinations200Response UpdateNetworkCellularGatewayConnectivityMonitoringDestinations(ctx, networkId).UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest(updateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest).Execute()

Update the connectivity testing destinations for an MG network



### Example

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
	updateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest := *openapiclient.NewUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest() // UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CellularGatewayAPI.UpdateNetworkCellularGatewayConnectivityMonitoringDestinations(context.Background(), networkId).UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest(updateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayAPI.UpdateNetworkCellularGatewayConnectivityMonitoringDestinations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkCellularGatewayConnectivityMonitoringDestinations`: GetNetworkApplianceConnectivityMonitoringDestinations200Response
	fmt.Fprintf(os.Stdout, "Response from `CellularGatewayAPI.UpdateNetworkCellularGatewayConnectivityMonitoringDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest** | [**UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest**](UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest.md) |  | 

### Return type

[**GetNetworkApplianceConnectivityMonitoringDestinations200Response**](GetNetworkApplianceConnectivityMonitoringDestinations200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkCellularGatewayDhcp

> GetNetworkCellularGatewayDhcp200Response UpdateNetworkCellularGatewayDhcp(ctx, networkId).UpdateNetworkCellularGatewayDhcpRequest(updateNetworkCellularGatewayDhcpRequest).Execute()

Update common DHCP settings of MGs



### Example

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
	updateNetworkCellularGatewayDhcpRequest := *openapiclient.NewUpdateNetworkCellularGatewayDhcpRequest() // UpdateNetworkCellularGatewayDhcpRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CellularGatewayAPI.UpdateNetworkCellularGatewayDhcp(context.Background(), networkId).UpdateNetworkCellularGatewayDhcpRequest(updateNetworkCellularGatewayDhcpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayAPI.UpdateNetworkCellularGatewayDhcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkCellularGatewayDhcp`: GetNetworkCellularGatewayDhcp200Response
	fmt.Fprintf(os.Stdout, "Response from `CellularGatewayAPI.UpdateNetworkCellularGatewayDhcp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCellularGatewayDhcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkCellularGatewayDhcpRequest** | [**UpdateNetworkCellularGatewayDhcpRequest**](UpdateNetworkCellularGatewayDhcpRequest.md) |  | 

### Return type

[**GetNetworkCellularGatewayDhcp200Response**](GetNetworkCellularGatewayDhcp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkCellularGatewaySubnetPool

> GetNetworkCellularGatewaySubnetPool200Response UpdateNetworkCellularGatewaySubnetPool(ctx, networkId).UpdateNetworkCellularGatewaySubnetPoolRequest(updateNetworkCellularGatewaySubnetPoolRequest).Execute()

Update the subnet pool and mask configuration for MGs in the network.



### Example

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
	updateNetworkCellularGatewaySubnetPoolRequest := *openapiclient.NewUpdateNetworkCellularGatewaySubnetPoolRequest() // UpdateNetworkCellularGatewaySubnetPoolRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CellularGatewayAPI.UpdateNetworkCellularGatewaySubnetPool(context.Background(), networkId).UpdateNetworkCellularGatewaySubnetPoolRequest(updateNetworkCellularGatewaySubnetPoolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayAPI.UpdateNetworkCellularGatewaySubnetPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkCellularGatewaySubnetPool`: GetNetworkCellularGatewaySubnetPool200Response
	fmt.Fprintf(os.Stdout, "Response from `CellularGatewayAPI.UpdateNetworkCellularGatewaySubnetPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCellularGatewaySubnetPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkCellularGatewaySubnetPoolRequest** | [**UpdateNetworkCellularGatewaySubnetPoolRequest**](UpdateNetworkCellularGatewaySubnetPoolRequest.md) |  | 

### Return type

[**GetNetworkCellularGatewaySubnetPool200Response**](GetNetworkCellularGatewaySubnetPool200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkCellularGatewayUplink

> GetNetworkCellularGatewayUplink200Response UpdateNetworkCellularGatewayUplink(ctx, networkId).UpdateNetworkCellularGatewayUplinkRequest(updateNetworkCellularGatewayUplinkRequest).Execute()

Updates the uplink settings for your MG network.



### Example

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
	updateNetworkCellularGatewayUplinkRequest := *openapiclient.NewUpdateNetworkCellularGatewayUplinkRequest() // UpdateNetworkCellularGatewayUplinkRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CellularGatewayAPI.UpdateNetworkCellularGatewayUplink(context.Background(), networkId).UpdateNetworkCellularGatewayUplinkRequest(updateNetworkCellularGatewayUplinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CellularGatewayAPI.UpdateNetworkCellularGatewayUplink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkCellularGatewayUplink`: GetNetworkCellularGatewayUplink200Response
	fmt.Fprintf(os.Stdout, "Response from `CellularGatewayAPI.UpdateNetworkCellularGatewayUplink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCellularGatewayUplinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkCellularGatewayUplinkRequest** | [**UpdateNetworkCellularGatewayUplinkRequest**](UpdateNetworkCellularGatewayUplinkRequest.md) |  | 

### Return type

[**GetNetworkCellularGatewayUplink200Response**](GetNetworkCellularGatewayUplink200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

