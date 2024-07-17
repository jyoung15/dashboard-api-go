# \CellularFirewallRulesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkApplianceFirewallCellularFirewallRules**](CellularFirewallRulesAPI.md#GetNetworkApplianceFirewallCellularFirewallRules) | **Get** /networks/{networkId}/appliance/firewall/cellularFirewallRules | Return the cellular firewall rules for an MX network
[**UpdateNetworkApplianceFirewallCellularFirewallRules**](CellularFirewallRulesAPI.md#UpdateNetworkApplianceFirewallCellularFirewallRules) | **Put** /networks/{networkId}/appliance/firewall/cellularFirewallRules | Update the cellular firewall rules of an MX network



## GetNetworkApplianceFirewallCellularFirewallRules

> map[string]interface{} GetNetworkApplianceFirewallCellularFirewallRules(ctx, networkId).Execute()

Return the cellular firewall rules for an MX network



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
	resp, r, err := apiClient.CellularFirewallRulesAPI.GetNetworkApplianceFirewallCellularFirewallRules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CellularFirewallRulesAPI.GetNetworkApplianceFirewallCellularFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceFirewallCellularFirewallRules`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CellularFirewallRulesAPI.GetNetworkApplianceFirewallCellularFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallCellularFirewallRulesRequest struct via the builder pattern


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


## UpdateNetworkApplianceFirewallCellularFirewallRules

> map[string]interface{} UpdateNetworkApplianceFirewallCellularFirewallRules(ctx, networkId).UpdateNetworkApplianceFirewallCellularFirewallRulesRequest(updateNetworkApplianceFirewallCellularFirewallRulesRequest).Execute()

Update the cellular firewall rules of an MX network



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
	updateNetworkApplianceFirewallCellularFirewallRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallCellularFirewallRulesRequest() // UpdateNetworkApplianceFirewallCellularFirewallRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CellularFirewallRulesAPI.UpdateNetworkApplianceFirewallCellularFirewallRules(context.Background(), networkId).UpdateNetworkApplianceFirewallCellularFirewallRulesRequest(updateNetworkApplianceFirewallCellularFirewallRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CellularFirewallRulesAPI.UpdateNetworkApplianceFirewallCellularFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceFirewallCellularFirewallRules`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CellularFirewallRulesAPI.UpdateNetworkApplianceFirewallCellularFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallCellularFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallCellularFirewallRulesRequest** | [**UpdateNetworkApplianceFirewallCellularFirewallRulesRequest**](UpdateNetworkApplianceFirewallCellularFirewallRulesRequest.md) |  | 

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

