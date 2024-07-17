# \BonjourForwardingAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkWirelessSsidBonjourForwarding**](BonjourForwardingAPI.md#GetNetworkWirelessSsidBonjourForwarding) | **Get** /networks/{networkId}/wireless/ssids/{number}/bonjourForwarding | List the Bonjour forwarding setting and rules for the SSID
[**UpdateNetworkWirelessSsidBonjourForwarding**](BonjourForwardingAPI.md#UpdateNetworkWirelessSsidBonjourForwarding) | **Put** /networks/{networkId}/wireless/ssids/{number}/bonjourForwarding | Update the bonjour forwarding setting and rules for the SSID



## GetNetworkWirelessSsidBonjourForwarding

> GetNetworkWirelessSsidBonjourForwarding200Response GetNetworkWirelessSsidBonjourForwarding(ctx, networkId, number).Execute()

List the Bonjour forwarding setting and rules for the SSID



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
	resp, r, err := apiClient.BonjourForwardingAPI.GetNetworkWirelessSsidBonjourForwarding(context.Background(), networkId, number).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BonjourForwardingAPI.GetNetworkWirelessSsidBonjourForwarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessSsidBonjourForwarding`: GetNetworkWirelessSsidBonjourForwarding200Response
	fmt.Fprintf(os.Stdout, "Response from `BonjourForwardingAPI.GetNetworkWirelessSsidBonjourForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidBonjourForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkWirelessSsidBonjourForwarding200Response**](GetNetworkWirelessSsidBonjourForwarding200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidBonjourForwarding

> GetNetworkWirelessSsidBonjourForwarding200Response UpdateNetworkWirelessSsidBonjourForwarding(ctx, networkId, number).UpdateNetworkWirelessSsidBonjourForwardingRequest(updateNetworkWirelessSsidBonjourForwardingRequest).Execute()

Update the bonjour forwarding setting and rules for the SSID



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
	updateNetworkWirelessSsidBonjourForwardingRequest := *openapiclient.NewUpdateNetworkWirelessSsidBonjourForwardingRequest() // UpdateNetworkWirelessSsidBonjourForwardingRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BonjourForwardingAPI.UpdateNetworkWirelessSsidBonjourForwarding(context.Background(), networkId, number).UpdateNetworkWirelessSsidBonjourForwardingRequest(updateNetworkWirelessSsidBonjourForwardingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BonjourForwardingAPI.UpdateNetworkWirelessSsidBonjourForwarding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkWirelessSsidBonjourForwarding`: GetNetworkWirelessSsidBonjourForwarding200Response
	fmt.Fprintf(os.Stdout, "Response from `BonjourForwardingAPI.UpdateNetworkWirelessSsidBonjourForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidBonjourForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidBonjourForwardingRequest** | [**UpdateNetworkWirelessSsidBonjourForwardingRequest**](UpdateNetworkWirelessSsidBonjourForwardingRequest.md) |  | 

### Return type

[**GetNetworkWirelessSsidBonjourForwarding200Response**](GetNetworkWirelessSsidBonjourForwarding200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

