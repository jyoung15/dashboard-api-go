# \OneToManyNatRulesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkApplianceFirewallOneToManyNatRules**](OneToManyNatRulesAPI.md#GetNetworkApplianceFirewallOneToManyNatRules) | **Get** /networks/{networkId}/appliance/firewall/oneToManyNatRules | Return the 1:Many NAT mapping rules for an MX network
[**UpdateNetworkApplianceFirewallOneToManyNatRules**](OneToManyNatRulesAPI.md#UpdateNetworkApplianceFirewallOneToManyNatRules) | **Put** /networks/{networkId}/appliance/firewall/oneToManyNatRules | Set the 1:Many NAT mapping rules for an MX network



## GetNetworkApplianceFirewallOneToManyNatRules

> map[string]interface{} GetNetworkApplianceFirewallOneToManyNatRules(ctx, networkId).Execute()

Return the 1:Many NAT mapping rules for an MX network



### Example

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
	resp, r, err := apiClient.OneToManyNatRulesAPI.GetNetworkApplianceFirewallOneToManyNatRules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OneToManyNatRulesAPI.GetNetworkApplianceFirewallOneToManyNatRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceFirewallOneToManyNatRules`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OneToManyNatRulesAPI.GetNetworkApplianceFirewallOneToManyNatRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallOneToManyNatRulesRequest struct via the builder pattern


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


## UpdateNetworkApplianceFirewallOneToManyNatRules

> map[string]interface{} UpdateNetworkApplianceFirewallOneToManyNatRules(ctx, networkId).UpdateNetworkApplianceFirewallOneToManyNatRulesRequest(updateNetworkApplianceFirewallOneToManyNatRulesRequest).Execute()

Set the 1:Many NAT mapping rules for an MX network



### Example

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
	updateNetworkApplianceFirewallOneToManyNatRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallOneToManyNatRulesRequest([]openapiclient.UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner{*openapiclient.NewUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner("PublicIp_example", "Uplink_example", []openapiclient.UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner{*openapiclient.NewUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner()})}) // UpdateNetworkApplianceFirewallOneToManyNatRulesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OneToManyNatRulesAPI.UpdateNetworkApplianceFirewallOneToManyNatRules(context.Background(), networkId).UpdateNetworkApplianceFirewallOneToManyNatRulesRequest(updateNetworkApplianceFirewallOneToManyNatRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OneToManyNatRulesAPI.UpdateNetworkApplianceFirewallOneToManyNatRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceFirewallOneToManyNatRules`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OneToManyNatRulesAPI.UpdateNetworkApplianceFirewallOneToManyNatRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallOneToManyNatRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallOneToManyNatRulesRequest** | [**UpdateNetworkApplianceFirewallOneToManyNatRulesRequest**](UpdateNetworkApplianceFirewallOneToManyNatRulesRequest.md) |  | 

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

