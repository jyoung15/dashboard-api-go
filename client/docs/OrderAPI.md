# \OrderAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSwitchQosRulesOrder**](OrderAPI.md#GetNetworkSwitchQosRulesOrder) | **Get** /networks/{networkId}/switch/qosRules/order | Return the quality of service rule IDs by order in which they will be processed by the switch
[**UpdateNetworkSwitchQosRulesOrder**](OrderAPI.md#UpdateNetworkSwitchQosRulesOrder) | **Put** /networks/{networkId}/switch/qosRules/order | Update the order in which the rules should be processed by the switch



## GetNetworkSwitchQosRulesOrder

> GetNetworkSwitchQosRulesOrder200Response GetNetworkSwitchQosRulesOrder(ctx, networkId).Execute()

Return the quality of service rule IDs by order in which they will be processed by the switch



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
	resp, r, err := apiClient.OrderAPI.GetNetworkSwitchQosRulesOrder(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.GetNetworkSwitchQosRulesOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchQosRulesOrder`: GetNetworkSwitchQosRulesOrder200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.GetNetworkSwitchQosRulesOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchQosRulesOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchQosRulesOrder200Response**](GetNetworkSwitchQosRulesOrder200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchQosRulesOrder

> GetNetworkSwitchQosRulesOrder200Response UpdateNetworkSwitchQosRulesOrder(ctx, networkId).UpdateNetworkSwitchQosRulesOrderRequest(updateNetworkSwitchQosRulesOrderRequest).Execute()

Update the order in which the rules should be processed by the switch



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
	updateNetworkSwitchQosRulesOrderRequest := *openapiclient.NewUpdateNetworkSwitchQosRulesOrderRequest([]string{"RuleIds_example"}) // UpdateNetworkSwitchQosRulesOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.UpdateNetworkSwitchQosRulesOrder(context.Background(), networkId).UpdateNetworkSwitchQosRulesOrderRequest(updateNetworkSwitchQosRulesOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.UpdateNetworkSwitchQosRulesOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchQosRulesOrder`: GetNetworkSwitchQosRulesOrder200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.UpdateNetworkSwitchQosRulesOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchQosRulesOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchQosRulesOrderRequest** | [**UpdateNetworkSwitchQosRulesOrderRequest**](UpdateNetworkSwitchQosRulesOrderRequest.md) |  | 

### Return type

[**GetNetworkSwitchQosRulesOrder200Response**](GetNetworkSwitchQosRulesOrder200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

