# \L3FirewallRulesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkApplianceFirewallL3FirewallRules**](L3FirewallRulesAPI.md#GetNetworkApplianceFirewallL3FirewallRules) | **Get** /networks/{networkId}/appliance/firewall/l3FirewallRules | Return the L3 firewall rules for an MX network
[**GetNetworkWirelessSsidFirewallL3FirewallRules**](L3FirewallRulesAPI.md#GetNetworkWirelessSsidFirewallL3FirewallRules) | **Get** /networks/{networkId}/wireless/ssids/{number}/firewall/l3FirewallRules | Return the L3 firewall rules for an SSID on an MR network
[**UpdateNetworkApplianceFirewallL3FirewallRules**](L3FirewallRulesAPI.md#UpdateNetworkApplianceFirewallL3FirewallRules) | **Put** /networks/{networkId}/appliance/firewall/l3FirewallRules | Update the L3 firewall rules of an MX network
[**UpdateNetworkWirelessSsidFirewallL3FirewallRules**](L3FirewallRulesAPI.md#UpdateNetworkWirelessSsidFirewallL3FirewallRules) | **Put** /networks/{networkId}/wireless/ssids/{number}/firewall/l3FirewallRules | Update the L3 firewall rules of an SSID on an MR network



## GetNetworkApplianceFirewallL3FirewallRules

> map[string]interface{} GetNetworkApplianceFirewallL3FirewallRules(ctx, networkId).Execute()

Return the L3 firewall rules for an MX network



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
	resp, r, err := apiClient.L3FirewallRulesAPI.GetNetworkApplianceFirewallL3FirewallRules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `L3FirewallRulesAPI.GetNetworkApplianceFirewallL3FirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceFirewallL3FirewallRules`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `L3FirewallRulesAPI.GetNetworkApplianceFirewallL3FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallL3FirewallRulesRequest struct via the builder pattern


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


## GetNetworkWirelessSsidFirewallL3FirewallRules

> GetNetworkWirelessSsidFirewallL3FirewallRules200Response GetNetworkWirelessSsidFirewallL3FirewallRules(ctx, networkId, number).Execute()

Return the L3 firewall rules for an SSID on an MR network



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
	number := "number_example" // string | Number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.L3FirewallRulesAPI.GetNetworkWirelessSsidFirewallL3FirewallRules(context.Background(), networkId, number).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `L3FirewallRulesAPI.GetNetworkWirelessSsidFirewallL3FirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessSsidFirewallL3FirewallRules`: GetNetworkWirelessSsidFirewallL3FirewallRules200Response
	fmt.Fprintf(os.Stdout, "Response from `L3FirewallRulesAPI.GetNetworkWirelessSsidFirewallL3FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidFirewallL3FirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkWirelessSsidFirewallL3FirewallRules200Response**](GetNetworkWirelessSsidFirewallL3FirewallRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallL3FirewallRules

> map[string]interface{} UpdateNetworkApplianceFirewallL3FirewallRules(ctx, networkId).UpdateNetworkApplianceFirewallL3FirewallRulesRequest(updateNetworkApplianceFirewallL3FirewallRulesRequest).Execute()

Update the L3 firewall rules of an MX network



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
	updateNetworkApplianceFirewallL3FirewallRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallL3FirewallRulesRequest() // UpdateNetworkApplianceFirewallL3FirewallRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.L3FirewallRulesAPI.UpdateNetworkApplianceFirewallL3FirewallRules(context.Background(), networkId).UpdateNetworkApplianceFirewallL3FirewallRulesRequest(updateNetworkApplianceFirewallL3FirewallRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `L3FirewallRulesAPI.UpdateNetworkApplianceFirewallL3FirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceFirewallL3FirewallRules`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `L3FirewallRulesAPI.UpdateNetworkApplianceFirewallL3FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallL3FirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallL3FirewallRulesRequest** | [**UpdateNetworkApplianceFirewallL3FirewallRulesRequest**](UpdateNetworkApplianceFirewallL3FirewallRulesRequest.md) |  | 

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


## UpdateNetworkWirelessSsidFirewallL3FirewallRules

> GetNetworkWirelessSsidFirewallL3FirewallRules200Response UpdateNetworkWirelessSsidFirewallL3FirewallRules(ctx, networkId, number).UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest(updateNetworkWirelessSsidFirewallL3FirewallRulesRequest).Execute()

Update the L3 firewall rules of an SSID on an MR network



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
	number := "number_example" // string | Number
	updateNetworkWirelessSsidFirewallL3FirewallRulesRequest := *openapiclient.NewUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest() // UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.L3FirewallRulesAPI.UpdateNetworkWirelessSsidFirewallL3FirewallRules(context.Background(), networkId, number).UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest(updateNetworkWirelessSsidFirewallL3FirewallRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `L3FirewallRulesAPI.UpdateNetworkWirelessSsidFirewallL3FirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkWirelessSsidFirewallL3FirewallRules`: GetNetworkWirelessSsidFirewallL3FirewallRules200Response
	fmt.Fprintf(os.Stdout, "Response from `L3FirewallRulesAPI.UpdateNetworkWirelessSsidFirewallL3FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidFirewallL3FirewallRulesRequest** | [**UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest**](UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest.md) |  | 

### Return type

[**GetNetworkWirelessSsidFirewallL3FirewallRules200Response**](GetNetworkWirelessSsidFirewallL3FirewallRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

