# \InboundCellularFirewallRulesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkApplianceFirewallInboundCellularFirewallRules**](InboundCellularFirewallRulesAPI.md#GetNetworkApplianceFirewallInboundCellularFirewallRules) | **Get** /networks/{networkId}/appliance/firewall/inboundCellularFirewallRules | Return the inbound cellular firewall rules for an MX network
[**UpdateNetworkApplianceFirewallInboundCellularFirewallRules**](InboundCellularFirewallRulesAPI.md#UpdateNetworkApplianceFirewallInboundCellularFirewallRules) | **Put** /networks/{networkId}/appliance/firewall/inboundCellularFirewallRules | Update the inbound cellular firewall rules of an MX network



## GetNetworkApplianceFirewallInboundCellularFirewallRules

> GetNetworkApplianceFirewallInboundCellularFirewallRules200Response GetNetworkApplianceFirewallInboundCellularFirewallRules(ctx, networkId).Execute()

Return the inbound cellular firewall rules for an MX network



### Example

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
	resp, r, err := apiClient.InboundCellularFirewallRulesAPI.GetNetworkApplianceFirewallInboundCellularFirewallRules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboundCellularFirewallRulesAPI.GetNetworkApplianceFirewallInboundCellularFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceFirewallInboundCellularFirewallRules`: GetNetworkApplianceFirewallInboundCellularFirewallRules200Response
	fmt.Fprintf(os.Stdout, "Response from `InboundCellularFirewallRulesAPI.GetNetworkApplianceFirewallInboundCellularFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallInboundCellularFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceFirewallInboundCellularFirewallRules200Response**](GetNetworkApplianceFirewallInboundCellularFirewallRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallInboundCellularFirewallRules

> GetNetworkApplianceFirewallInboundCellularFirewallRules200Response UpdateNetworkApplianceFirewallInboundCellularFirewallRules(ctx, networkId).UpdateNetworkApplianceFirewallInboundCellularFirewallRulesRequest(updateNetworkApplianceFirewallInboundCellularFirewallRulesRequest).Execute()

Update the inbound cellular firewall rules of an MX network



### Example

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
	updateNetworkApplianceFirewallInboundCellularFirewallRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallInboundCellularFirewallRulesRequest() // UpdateNetworkApplianceFirewallInboundCellularFirewallRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InboundCellularFirewallRulesAPI.UpdateNetworkApplianceFirewallInboundCellularFirewallRules(context.Background(), networkId).UpdateNetworkApplianceFirewallInboundCellularFirewallRulesRequest(updateNetworkApplianceFirewallInboundCellularFirewallRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InboundCellularFirewallRulesAPI.UpdateNetworkApplianceFirewallInboundCellularFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceFirewallInboundCellularFirewallRules`: GetNetworkApplianceFirewallInboundCellularFirewallRules200Response
	fmt.Fprintf(os.Stdout, "Response from `InboundCellularFirewallRulesAPI.UpdateNetworkApplianceFirewallInboundCellularFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallInboundCellularFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallInboundCellularFirewallRulesRequest** | [**UpdateNetworkApplianceFirewallInboundCellularFirewallRulesRequest**](UpdateNetworkApplianceFirewallInboundCellularFirewallRulesRequest.md) |  | 

### Return type

[**GetNetworkApplianceFirewallInboundCellularFirewallRules200Response**](GetNetworkApplianceFirewallInboundCellularFirewallRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

