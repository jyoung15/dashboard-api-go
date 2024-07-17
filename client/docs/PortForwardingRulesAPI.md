# \PortForwardingRulesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceCellularGatewayPortForwardingRules**](PortForwardingRulesAPI.md#GetDeviceCellularGatewayPortForwardingRules) | **Get** /devices/{serial}/cellularGateway/portForwardingRules | Returns the port forwarding rules for a single MG.
[**GetNetworkApplianceFirewallPortForwardingRules**](PortForwardingRulesAPI.md#GetNetworkApplianceFirewallPortForwardingRules) | **Get** /networks/{networkId}/appliance/firewall/portForwardingRules | Return the port forwarding rules for an MX network
[**UpdateDeviceCellularGatewayPortForwardingRules**](PortForwardingRulesAPI.md#UpdateDeviceCellularGatewayPortForwardingRules) | **Put** /devices/{serial}/cellularGateway/portForwardingRules | Updates the port forwarding rules for a single MG.
[**UpdateNetworkApplianceFirewallPortForwardingRules**](PortForwardingRulesAPI.md#UpdateNetworkApplianceFirewallPortForwardingRules) | **Put** /networks/{networkId}/appliance/firewall/portForwardingRules | Update the port forwarding rules for an MX network



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
	resp, r, err := apiClient.PortForwardingRulesAPI.GetDeviceCellularGatewayPortForwardingRules(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortForwardingRulesAPI.GetDeviceCellularGatewayPortForwardingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCellularGatewayPortForwardingRules`: GetDeviceCellularGatewayPortForwardingRules200Response
	fmt.Fprintf(os.Stdout, "Response from `PortForwardingRulesAPI.GetDeviceCellularGatewayPortForwardingRules`: %v\n", resp)
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


## GetNetworkApplianceFirewallPortForwardingRules

> GetNetworkApplianceFirewallPortForwardingRules200Response GetNetworkApplianceFirewallPortForwardingRules(ctx, networkId).Execute()

Return the port forwarding rules for an MX network



### Example

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
	resp, r, err := apiClient.PortForwardingRulesAPI.GetNetworkApplianceFirewallPortForwardingRules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortForwardingRulesAPI.GetNetworkApplianceFirewallPortForwardingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceFirewallPortForwardingRules`: GetNetworkApplianceFirewallPortForwardingRules200Response
	fmt.Fprintf(os.Stdout, "Response from `PortForwardingRulesAPI.GetNetworkApplianceFirewallPortForwardingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallPortForwardingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceFirewallPortForwardingRules200Response**](GetNetworkApplianceFirewallPortForwardingRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
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
	resp, r, err := apiClient.PortForwardingRulesAPI.UpdateDeviceCellularGatewayPortForwardingRules(context.Background(), serial).UpdateDeviceCellularGatewayPortForwardingRulesRequest(updateDeviceCellularGatewayPortForwardingRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortForwardingRulesAPI.UpdateDeviceCellularGatewayPortForwardingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceCellularGatewayPortForwardingRules`: GetDeviceCellularGatewayPortForwardingRules200Response
	fmt.Fprintf(os.Stdout, "Response from `PortForwardingRulesAPI.UpdateDeviceCellularGatewayPortForwardingRules`: %v\n", resp)
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


## UpdateNetworkApplianceFirewallPortForwardingRules

> GetNetworkApplianceFirewallPortForwardingRules200Response UpdateNetworkApplianceFirewallPortForwardingRules(ctx, networkId).UpdateNetworkApplianceFirewallPortForwardingRulesRequest(updateNetworkApplianceFirewallPortForwardingRulesRequest).Execute()

Update the port forwarding rules for an MX network



### Example

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
	updateNetworkApplianceFirewallPortForwardingRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallPortForwardingRulesRequest([]openapiclient.UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner{*openapiclient.NewUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner("LanIp_example", "PublicPort_example", "LocalPort_example", []string{"AllowedIps_example"}, "Protocol_example")}) // UpdateNetworkApplianceFirewallPortForwardingRulesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortForwardingRulesAPI.UpdateNetworkApplianceFirewallPortForwardingRules(context.Background(), networkId).UpdateNetworkApplianceFirewallPortForwardingRulesRequest(updateNetworkApplianceFirewallPortForwardingRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortForwardingRulesAPI.UpdateNetworkApplianceFirewallPortForwardingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceFirewallPortForwardingRules`: GetNetworkApplianceFirewallPortForwardingRules200Response
	fmt.Fprintf(os.Stdout, "Response from `PortForwardingRulesAPI.UpdateNetworkApplianceFirewallPortForwardingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallPortForwardingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallPortForwardingRulesRequest** | [**UpdateNetworkApplianceFirewallPortForwardingRulesRequest**](UpdateNetworkApplianceFirewallPortForwardingRulesRequest.md) |  | 

### Return type

[**GetNetworkApplianceFirewallPortForwardingRules200Response**](GetNetworkApplianceFirewallPortForwardingRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

