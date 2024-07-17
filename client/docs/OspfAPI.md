# \OspfAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSwitchRoutingOspf**](OspfAPI.md#GetNetworkSwitchRoutingOspf) | **Get** /networks/{networkId}/switch/routing/ospf | Return layer 3 OSPF routing configuration
[**UpdateNetworkSwitchRoutingOspf**](OspfAPI.md#UpdateNetworkSwitchRoutingOspf) | **Put** /networks/{networkId}/switch/routing/ospf | Update layer 3 OSPF routing configuration



## GetNetworkSwitchRoutingOspf

> GetNetworkSwitchRoutingOspf200Response GetNetworkSwitchRoutingOspf(ctx, networkId).Execute()

Return layer 3 OSPF routing configuration



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
	resp, r, err := apiClient.OspfAPI.GetNetworkSwitchRoutingOspf(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspfAPI.GetNetworkSwitchRoutingOspf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchRoutingOspf`: GetNetworkSwitchRoutingOspf200Response
	fmt.Fprintf(os.Stdout, "Response from `OspfAPI.GetNetworkSwitchRoutingOspf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchRoutingOspfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchRoutingOspf200Response**](GetNetworkSwitchRoutingOspf200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchRoutingOspf

> GetNetworkSwitchRoutingOspf200Response UpdateNetworkSwitchRoutingOspf(ctx, networkId).UpdateNetworkSwitchRoutingOspfRequest(updateNetworkSwitchRoutingOspfRequest).Execute()

Update layer 3 OSPF routing configuration



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
	updateNetworkSwitchRoutingOspfRequest := *openapiclient.NewUpdateNetworkSwitchRoutingOspfRequest() // UpdateNetworkSwitchRoutingOspfRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OspfAPI.UpdateNetworkSwitchRoutingOspf(context.Background(), networkId).UpdateNetworkSwitchRoutingOspfRequest(updateNetworkSwitchRoutingOspfRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OspfAPI.UpdateNetworkSwitchRoutingOspf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchRoutingOspf`: GetNetworkSwitchRoutingOspf200Response
	fmt.Fprintf(os.Stdout, "Response from `OspfAPI.UpdateNetworkSwitchRoutingOspf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchRoutingOspfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchRoutingOspfRequest** | [**UpdateNetworkSwitchRoutingOspfRequest**](UpdateNetworkSwitchRoutingOspfRequest.md) |  | 

### Return type

[**GetNetworkSwitchRoutingOspf200Response**](GetNetworkSwitchRoutingOspf200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

