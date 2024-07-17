# \WarmSpareAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceSwitchWarmSpare**](WarmSpareAPI.md#GetDeviceSwitchWarmSpare) | **Get** /devices/{serial}/switch/warmSpare | Return warm spare configuration for a switch
[**GetNetworkApplianceWarmSpare**](WarmSpareAPI.md#GetNetworkApplianceWarmSpare) | **Get** /networks/{networkId}/appliance/warmSpare | Return MX warm spare settings
[**SwapNetworkApplianceWarmSpare**](WarmSpareAPI.md#SwapNetworkApplianceWarmSpare) | **Post** /networks/{networkId}/appliance/warmSpare/swap | Swap MX primary and warm spare appliances
[**UpdateDeviceSwitchWarmSpare**](WarmSpareAPI.md#UpdateDeviceSwitchWarmSpare) | **Put** /devices/{serial}/switch/warmSpare | Update warm spare configuration for a switch
[**UpdateNetworkApplianceWarmSpare**](WarmSpareAPI.md#UpdateNetworkApplianceWarmSpare) | **Put** /networks/{networkId}/appliance/warmSpare | Update MX warm spare settings



## GetDeviceSwitchWarmSpare

> GetDeviceSwitchWarmSpare200Response GetDeviceSwitchWarmSpare(ctx, serial).Execute()

Return warm spare configuration for a switch



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
	serial := "serial_example" // string | Serial

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarmSpareAPI.GetDeviceSwitchWarmSpare(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarmSpareAPI.GetDeviceSwitchWarmSpare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchWarmSpare`: GetDeviceSwitchWarmSpare200Response
	fmt.Fprintf(os.Stdout, "Response from `WarmSpareAPI.GetDeviceSwitchWarmSpare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchWarmSpareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceSwitchWarmSpare200Response**](GetDeviceSwitchWarmSpare200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceWarmSpare

> GetNetworkApplianceWarmSpare200Response GetNetworkApplianceWarmSpare(ctx, networkId).Execute()

Return MX warm spare settings



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
	resp, r, err := apiClient.WarmSpareAPI.GetNetworkApplianceWarmSpare(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarmSpareAPI.GetNetworkApplianceWarmSpare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceWarmSpare`: GetNetworkApplianceWarmSpare200Response
	fmt.Fprintf(os.Stdout, "Response from `WarmSpareAPI.GetNetworkApplianceWarmSpare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceWarmSpareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceWarmSpare200Response**](GetNetworkApplianceWarmSpare200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwapNetworkApplianceWarmSpare

> GetNetworkApplianceWarmSpare200Response SwapNetworkApplianceWarmSpare(ctx, networkId).Execute()

Swap MX primary and warm spare appliances



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
	resp, r, err := apiClient.WarmSpareAPI.SwapNetworkApplianceWarmSpare(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarmSpareAPI.SwapNetworkApplianceWarmSpare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SwapNetworkApplianceWarmSpare`: GetNetworkApplianceWarmSpare200Response
	fmt.Fprintf(os.Stdout, "Response from `WarmSpareAPI.SwapNetworkApplianceWarmSpare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwapNetworkApplianceWarmSpareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceWarmSpare200Response**](GetNetworkApplianceWarmSpare200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceSwitchWarmSpare

> GetDeviceSwitchWarmSpare200Response UpdateDeviceSwitchWarmSpare(ctx, serial).UpdateDeviceSwitchWarmSpareRequest(updateDeviceSwitchWarmSpareRequest).Execute()

Update warm spare configuration for a switch



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
	serial := "serial_example" // string | Serial
	updateDeviceSwitchWarmSpareRequest := *openapiclient.NewUpdateDeviceSwitchWarmSpareRequest(false) // UpdateDeviceSwitchWarmSpareRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarmSpareAPI.UpdateDeviceSwitchWarmSpare(context.Background(), serial).UpdateDeviceSwitchWarmSpareRequest(updateDeviceSwitchWarmSpareRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarmSpareAPI.UpdateDeviceSwitchWarmSpare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceSwitchWarmSpare`: GetDeviceSwitchWarmSpare200Response
	fmt.Fprintf(os.Stdout, "Response from `WarmSpareAPI.UpdateDeviceSwitchWarmSpare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceSwitchWarmSpareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceSwitchWarmSpareRequest** | [**UpdateDeviceSwitchWarmSpareRequest**](UpdateDeviceSwitchWarmSpareRequest.md) |  | 

### Return type

[**GetDeviceSwitchWarmSpare200Response**](GetDeviceSwitchWarmSpare200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceWarmSpare

> GetNetworkApplianceWarmSpare200Response UpdateNetworkApplianceWarmSpare(ctx, networkId).UpdateNetworkApplianceWarmSpareRequest(updateNetworkApplianceWarmSpareRequest).Execute()

Update MX warm spare settings



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
	updateNetworkApplianceWarmSpareRequest := *openapiclient.NewUpdateNetworkApplianceWarmSpareRequest(false) // UpdateNetworkApplianceWarmSpareRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WarmSpareAPI.UpdateNetworkApplianceWarmSpare(context.Background(), networkId).UpdateNetworkApplianceWarmSpareRequest(updateNetworkApplianceWarmSpareRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WarmSpareAPI.UpdateNetworkApplianceWarmSpare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceWarmSpare`: GetNetworkApplianceWarmSpare200Response
	fmt.Fprintf(os.Stdout, "Response from `WarmSpareAPI.UpdateNetworkApplianceWarmSpare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceWarmSpareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceWarmSpareRequest** | [**UpdateNetworkApplianceWarmSpareRequest**](UpdateNetworkApplianceWarmSpareRequest.md) |  | 

### Return type

[**GetNetworkApplianceWarmSpare200Response**](GetNetworkApplianceWarmSpare200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

