# \LanAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceCellularGatewayLan**](LanAPI.md#GetDeviceCellularGatewayLan) | **Get** /devices/{serial}/cellularGateway/lan | Show the LAN Settings of a MG
[**UpdateDeviceCellularGatewayLan**](LanAPI.md#UpdateDeviceCellularGatewayLan) | **Put** /devices/{serial}/cellularGateway/lan | Update the LAN Settings for a single MG.



## GetDeviceCellularGatewayLan

> GetDeviceCellularGatewayLan200Response GetDeviceCellularGatewayLan(ctx, serial).Execute()

Show the LAN Settings of a MG



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
	resp, r, err := apiClient.LanAPI.GetDeviceCellularGatewayLan(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanAPI.GetDeviceCellularGatewayLan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCellularGatewayLan`: GetDeviceCellularGatewayLan200Response
	fmt.Fprintf(os.Stdout, "Response from `LanAPI.GetDeviceCellularGatewayLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCellularGatewayLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceCellularGatewayLan200Response**](GetDeviceCellularGatewayLan200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCellularGatewayLan

> GetDeviceCellularGatewayLan200Response UpdateDeviceCellularGatewayLan(ctx, serial).UpdateDeviceCellularGatewayLanRequest(updateDeviceCellularGatewayLanRequest).Execute()

Update the LAN Settings for a single MG.



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
	updateDeviceCellularGatewayLanRequest := *openapiclient.NewUpdateDeviceCellularGatewayLanRequest() // UpdateDeviceCellularGatewayLanRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LanAPI.UpdateDeviceCellularGatewayLan(context.Background(), serial).UpdateDeviceCellularGatewayLanRequest(updateDeviceCellularGatewayLanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LanAPI.UpdateDeviceCellularGatewayLan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceCellularGatewayLan`: GetDeviceCellularGatewayLan200Response
	fmt.Fprintf(os.Stdout, "Response from `LanAPI.UpdateDeviceCellularGatewayLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCellularGatewayLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCellularGatewayLanRequest** | [**UpdateDeviceCellularGatewayLanRequest**](UpdateDeviceCellularGatewayLanRequest.md) |  | 

### Return type

[**GetDeviceCellularGatewayLan200Response**](GetDeviceCellularGatewayLan200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

