# \EapOverrideAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkWirelessSsidEapOverride**](EapOverrideAPI.md#GetNetworkWirelessSsidEapOverride) | **Get** /networks/{networkId}/wireless/ssids/{number}/eapOverride | Return the EAP overridden parameters for an SSID
[**UpdateNetworkWirelessSsidEapOverride**](EapOverrideAPI.md#UpdateNetworkWirelessSsidEapOverride) | **Put** /networks/{networkId}/wireless/ssids/{number}/eapOverride | Update the EAP overridden parameters for an SSID.



## GetNetworkWirelessSsidEapOverride

> GetNetworkWirelessSsidEapOverride200Response GetNetworkWirelessSsidEapOverride(ctx, networkId, number).Execute()

Return the EAP overridden parameters for an SSID



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
	resp, r, err := apiClient.EapOverrideAPI.GetNetworkWirelessSsidEapOverride(context.Background(), networkId, number).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EapOverrideAPI.GetNetworkWirelessSsidEapOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessSsidEapOverride`: GetNetworkWirelessSsidEapOverride200Response
	fmt.Fprintf(os.Stdout, "Response from `EapOverrideAPI.GetNetworkWirelessSsidEapOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidEapOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkWirelessSsidEapOverride200Response**](GetNetworkWirelessSsidEapOverride200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidEapOverride

> GetNetworkWirelessSsidEapOverride200Response UpdateNetworkWirelessSsidEapOverride(ctx, networkId, number).UpdateNetworkWirelessSsidEapOverrideRequest(updateNetworkWirelessSsidEapOverrideRequest).Execute()

Update the EAP overridden parameters for an SSID.



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
	updateNetworkWirelessSsidEapOverrideRequest := *openapiclient.NewUpdateNetworkWirelessSsidEapOverrideRequest() // UpdateNetworkWirelessSsidEapOverrideRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EapOverrideAPI.UpdateNetworkWirelessSsidEapOverride(context.Background(), networkId, number).UpdateNetworkWirelessSsidEapOverrideRequest(updateNetworkWirelessSsidEapOverrideRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EapOverrideAPI.UpdateNetworkWirelessSsidEapOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkWirelessSsidEapOverride`: GetNetworkWirelessSsidEapOverride200Response
	fmt.Fprintf(os.Stdout, "Response from `EapOverrideAPI.UpdateNetworkWirelessSsidEapOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidEapOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidEapOverrideRequest** | [**UpdateNetworkWirelessSsidEapOverrideRequest**](UpdateNetworkWirelessSsidEapOverrideRequest.md) |  | 

### Return type

[**GetNetworkWirelessSsidEapOverride200Response**](GetNetworkWirelessSsidEapOverride200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

