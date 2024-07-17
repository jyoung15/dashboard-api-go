# \InboundFirewallRulesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkApplianceFirewallInboundFirewallRules**](InboundFirewallRulesAPI.md#GetNetworkApplianceFirewallInboundFirewallRules) | **Get** /networks/{networkId}/appliance/firewall/inboundFirewallRules | Return the inbound firewall rules for an MX network
[**UpdateNetworkApplianceFirewallInboundFirewallRules**](InboundFirewallRulesAPI.md#UpdateNetworkApplianceFirewallInboundFirewallRules) | **Put** /networks/{networkId}/appliance/firewall/inboundFirewallRules | Update the inbound firewall rules of an MX network



## GetNetworkApplianceFirewallInboundFirewallRules

> GetNetworkApplianceFirewallInboundFirewallRules200Response GetNetworkApplianceFirewallInboundFirewallRules(ctx, networkId).Execute()

Return the inbound firewall rules for an MX network



### Example

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
	resp, r, err := apiClient.InboundFirewallRulesAPI.GetNetworkApplianceFirewallInboundFirewallRules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboundFirewallRulesAPI.GetNetworkApplianceFirewallInboundFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceFirewallInboundFirewallRules`: GetNetworkApplianceFirewallInboundFirewallRules200Response
	fmt.Fprintf(os.Stdout, "Response from `InboundFirewallRulesAPI.GetNetworkApplianceFirewallInboundFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallInboundFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceFirewallInboundFirewallRules200Response**](GetNetworkApplianceFirewallInboundFirewallRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallInboundFirewallRules

> GetNetworkApplianceFirewallInboundFirewallRules200Response UpdateNetworkApplianceFirewallInboundFirewallRules(ctx, networkId).UpdateNetworkApplianceFirewallInboundFirewallRulesRequest(updateNetworkApplianceFirewallInboundFirewallRulesRequest).Execute()

Update the inbound firewall rules of an MX network



### Example

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
	updateNetworkApplianceFirewallInboundFirewallRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallInboundFirewallRulesRequest() // UpdateNetworkApplianceFirewallInboundFirewallRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboundFirewallRulesAPI.UpdateNetworkApplianceFirewallInboundFirewallRules(context.Background(), networkId).UpdateNetworkApplianceFirewallInboundFirewallRulesRequest(updateNetworkApplianceFirewallInboundFirewallRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboundFirewallRulesAPI.UpdateNetworkApplianceFirewallInboundFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceFirewallInboundFirewallRules`: GetNetworkApplianceFirewallInboundFirewallRules200Response
	fmt.Fprintf(os.Stdout, "Response from `InboundFirewallRulesAPI.UpdateNetworkApplianceFirewallInboundFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallInboundFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallInboundFirewallRulesRequest** | [**UpdateNetworkApplianceFirewallInboundFirewallRulesRequest**](UpdateNetworkApplianceFirewallInboundFirewallRulesRequest.md) |  | 

### Return type

[**GetNetworkApplianceFirewallInboundFirewallRules200Response**](GetNetworkApplianceFirewallInboundFirewallRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

