# \ApplicationCategoriesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories**](ApplicationCategoriesAPI.md#GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories) | **Get** /networks/{networkId}/appliance/firewall/l7FirewallRules/applicationCategories | Return the L7 firewall application categories and their associated applications for an MX network
[**GetNetworkTrafficShapingApplicationCategories**](ApplicationCategoriesAPI.md#GetNetworkTrafficShapingApplicationCategories) | **Get** /networks/{networkId}/trafficShaping/applicationCategories | Returns the application categories for traffic shaping rules



## GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories

> GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories(ctx, networkId).Execute()

Return the L7 firewall application categories and their associated applications for an MX network



### Example

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
	resp, r, err := apiClient.ApplicationCategoriesAPI.GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories`: GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplicationCategoriesAPI.GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response**](GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response.md)

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
	resp, r, err := apiClient.ApplicationCategoriesAPI.GetNetworkTrafficShapingApplicationCategories(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.GetNetworkTrafficShapingApplicationCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkTrafficShapingApplicationCategories`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplicationCategoriesAPI.GetNetworkTrafficShapingApplicationCategories`: %v\n", resp)
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

