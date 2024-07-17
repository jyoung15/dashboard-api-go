# \FirewalledServicesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkApplianceFirewallFirewalledService**](FirewalledServicesAPI.md#GetNetworkApplianceFirewallFirewalledService) | **Get** /networks/{networkId}/appliance/firewall/firewalledServices/{service} | Return the accessibility settings of the given service (&#39;ICMP&#39;, &#39;web&#39;, or &#39;SNMP&#39;)
[**GetNetworkApplianceFirewallFirewalledServices**](FirewalledServicesAPI.md#GetNetworkApplianceFirewallFirewalledServices) | **Get** /networks/{networkId}/appliance/firewall/firewalledServices | List the appliance services and their accessibility rules
[**UpdateNetworkApplianceFirewallFirewalledService**](FirewalledServicesAPI.md#UpdateNetworkApplianceFirewallFirewalledService) | **Put** /networks/{networkId}/appliance/firewall/firewalledServices/{service} | Updates the accessibility settings for the given service (&#39;ICMP&#39;, &#39;web&#39;, or &#39;SNMP&#39;)



## GetNetworkApplianceFirewallFirewalledService

> GetNetworkApplianceFirewallFirewalledServices200ResponseInner GetNetworkApplianceFirewallFirewalledService(ctx, networkId, service).Execute()

Return the accessibility settings of the given service ('ICMP', 'web', or 'SNMP')



### Example

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
	service := "service_example" // string | Service

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewalledServicesAPI.GetNetworkApplianceFirewallFirewalledService(context.Background(), networkId, service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewalledServicesAPI.GetNetworkApplianceFirewallFirewalledService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceFirewallFirewalledService`: GetNetworkApplianceFirewallFirewalledServices200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `FirewalledServicesAPI.GetNetworkApplianceFirewallFirewalledService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**service** | **string** | Service | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallFirewalledServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkApplianceFirewallFirewalledServices200ResponseInner**](GetNetworkApplianceFirewallFirewalledServices200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallFirewalledServices

> []GetNetworkApplianceFirewallFirewalledServices200ResponseInner GetNetworkApplianceFirewallFirewalledServices(ctx, networkId).Execute()

List the appliance services and their accessibility rules



### Example

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
	resp, r, err := apiClient.FirewalledServicesAPI.GetNetworkApplianceFirewallFirewalledServices(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewalledServicesAPI.GetNetworkApplianceFirewallFirewalledServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceFirewallFirewalledServices`: []GetNetworkApplianceFirewallFirewalledServices200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `FirewalledServicesAPI.GetNetworkApplianceFirewallFirewalledServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallFirewalledServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkApplianceFirewallFirewalledServices200ResponseInner**](GetNetworkApplianceFirewallFirewalledServices200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallFirewalledService

> GetNetworkApplianceFirewallFirewalledServices200ResponseInner UpdateNetworkApplianceFirewallFirewalledService(ctx, networkId, service).UpdateNetworkApplianceFirewallFirewalledServiceRequest(updateNetworkApplianceFirewallFirewalledServiceRequest).Execute()

Updates the accessibility settings for the given service ('ICMP', 'web', or 'SNMP')



### Example

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
	service := "service_example" // string | Service
	updateNetworkApplianceFirewallFirewalledServiceRequest := *openapiclient.NewUpdateNetworkApplianceFirewallFirewalledServiceRequest("Access_example") // UpdateNetworkApplianceFirewallFirewalledServiceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewalledServicesAPI.UpdateNetworkApplianceFirewallFirewalledService(context.Background(), networkId, service).UpdateNetworkApplianceFirewallFirewalledServiceRequest(updateNetworkApplianceFirewallFirewalledServiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewalledServicesAPI.UpdateNetworkApplianceFirewallFirewalledService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceFirewallFirewalledService`: GetNetworkApplianceFirewallFirewalledServices200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `FirewalledServicesAPI.UpdateNetworkApplianceFirewallFirewalledService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**service** | **string** | Service | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallFirewalledServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceFirewallFirewalledServiceRequest** | [**UpdateNetworkApplianceFirewallFirewalledServiceRequest**](UpdateNetworkApplianceFirewallFirewalledServiceRequest.md) |  | 

### Return type

[**GetNetworkApplianceFirewallFirewalledServices200ResponseInner**](GetNetworkApplianceFirewallFirewalledServices200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

