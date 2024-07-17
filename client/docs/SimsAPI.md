# \SimsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceCellularSims**](SimsAPI.md#GetDeviceCellularSims) | **Get** /devices/{serial}/cellular/sims | Return the SIM and APN configurations for a cellular device.
[**UpdateDeviceCellularSims**](SimsAPI.md#UpdateDeviceCellularSims) | **Put** /devices/{serial}/cellular/sims | Updates the SIM and APN configurations for a cellular device.



## GetDeviceCellularSims

> GetDeviceCellularSims200Response GetDeviceCellularSims(ctx, serial).Execute()

Return the SIM and APN configurations for a cellular device.



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
	resp, r, err := apiClient.SimsAPI.GetDeviceCellularSims(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimsAPI.GetDeviceCellularSims``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCellularSims`: GetDeviceCellularSims200Response
	fmt.Fprintf(os.Stdout, "Response from `SimsAPI.GetDeviceCellularSims`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCellularSimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceCellularSims200Response**](GetDeviceCellularSims200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCellularSims

> GetDeviceCellularSims200Response UpdateDeviceCellularSims(ctx, serial).UpdateDeviceCellularSimsRequest(updateDeviceCellularSimsRequest).Execute()

Updates the SIM and APN configurations for a cellular device.



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
	updateDeviceCellularSimsRequest := *openapiclient.NewUpdateDeviceCellularSimsRequest() // UpdateDeviceCellularSimsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SimsAPI.UpdateDeviceCellularSims(context.Background(), serial).UpdateDeviceCellularSimsRequest(updateDeviceCellularSimsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SimsAPI.UpdateDeviceCellularSims``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceCellularSims`: GetDeviceCellularSims200Response
	fmt.Fprintf(os.Stdout, "Response from `SimsAPI.UpdateDeviceCellularSims`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCellularSimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCellularSimsRequest** | [**UpdateDeviceCellularSimsRequest**](UpdateDeviceCellularSimsRequest.md) |  | 

### Return type

[**GetDeviceCellularSims200Response**](GetDeviceCellularSims200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

