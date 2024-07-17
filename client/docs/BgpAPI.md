# \BgpAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkApplianceVpnBgp**](BgpAPI.md#GetNetworkApplianceVpnBgp) | **Get** /networks/{networkId}/appliance/vpn/bgp | Return a Hub BGP Configuration
[**UpdateNetworkApplianceVpnBgp**](BgpAPI.md#UpdateNetworkApplianceVpnBgp) | **Put** /networks/{networkId}/appliance/vpn/bgp | Update a Hub BGP Configuration



## GetNetworkApplianceVpnBgp

> GetNetworkApplianceVpnBgp200Response GetNetworkApplianceVpnBgp(ctx, networkId).Execute()

Return a Hub BGP Configuration



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
	resp, r, err := apiClient.BgpAPI.GetNetworkApplianceVpnBgp(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BgpAPI.GetNetworkApplianceVpnBgp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceVpnBgp`: GetNetworkApplianceVpnBgp200Response
	fmt.Fprintf(os.Stdout, "Response from `BgpAPI.GetNetworkApplianceVpnBgp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVpnBgpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceVpnBgp200Response**](GetNetworkApplianceVpnBgp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceVpnBgp

> GetNetworkApplianceVpnBgp200Response UpdateNetworkApplianceVpnBgp(ctx, networkId).UpdateNetworkApplianceVpnBgpRequest(updateNetworkApplianceVpnBgpRequest).Execute()

Update a Hub BGP Configuration



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
	updateNetworkApplianceVpnBgpRequest := *openapiclient.NewUpdateNetworkApplianceVpnBgpRequest(false) // UpdateNetworkApplianceVpnBgpRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BgpAPI.UpdateNetworkApplianceVpnBgp(context.Background(), networkId).UpdateNetworkApplianceVpnBgpRequest(updateNetworkApplianceVpnBgpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BgpAPI.UpdateNetworkApplianceVpnBgp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceVpnBgp`: GetNetworkApplianceVpnBgp200Response
	fmt.Fprintf(os.Stdout, "Response from `BgpAPI.UpdateNetworkApplianceVpnBgp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceVpnBgpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceVpnBgpRequest** | [**UpdateNetworkApplianceVpnBgpRequest**](UpdateNetworkApplianceVpnBgpRequest.md) |  | 

### Return type

[**GetNetworkApplianceVpnBgp200Response**](GetNetworkApplianceVpnBgp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

