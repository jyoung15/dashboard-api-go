# \SplashLoginAttemptsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSplashLoginAttempts**](SplashLoginAttemptsAPI.md#GetNetworkSplashLoginAttempts) | **Get** /networks/{networkId}/splashLoginAttempts | List the splash login attempts for a network



## GetNetworkSplashLoginAttempts

> []GetNetworkSplashLoginAttempts200ResponseInner GetNetworkSplashLoginAttempts(ctx, networkId).SsidNumber(ssidNumber).LoginIdentifier(loginIdentifier).Timespan(timespan).Execute()

List the splash login attempts for a network



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
	ssidNumber := int32(56) // int32 | Only return the login attempts for the specified SSID (optional)
	loginIdentifier := "loginIdentifier_example" // string | The username, email, or phone number used during login (optional)
	timespan := int32(56) // int32 | The timespan, in seconds, for the login attempts. The period will be from [timespan] seconds ago until now. The maximum timespan is 3 months (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SplashLoginAttemptsAPI.GetNetworkSplashLoginAttempts(context.Background(), networkId).SsidNumber(ssidNumber).LoginIdentifier(loginIdentifier).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SplashLoginAttemptsAPI.GetNetworkSplashLoginAttempts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSplashLoginAttempts`: []GetNetworkSplashLoginAttempts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SplashLoginAttemptsAPI.GetNetworkSplashLoginAttempts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSplashLoginAttemptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ssidNumber** | **int32** | Only return the login attempts for the specified SSID | 
 **loginIdentifier** | **string** | The username, email, or phone number used during login | 
 **timespan** | **int32** | The timespan, in seconds, for the login attempts. The period will be from [timespan] seconds ago until now. The maximum timespan is 3 months | 

### Return type

[**[]GetNetworkSplashLoginAttempts200ResponseInner**](GetNetworkSplashLoginAttempts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

