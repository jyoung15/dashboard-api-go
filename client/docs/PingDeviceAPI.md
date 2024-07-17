# \PingDeviceAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceLiveToolsPingDevice**](PingDeviceAPI.md#CreateDeviceLiveToolsPingDevice) | **Post** /devices/{serial}/liveTools/pingDevice | Enqueue a job to check connectivity status to the device
[**GetDeviceLiveToolsPingDevice**](PingDeviceAPI.md#GetDeviceLiveToolsPingDevice) | **Get** /devices/{serial}/liveTools/pingDevice/{id} | Return a ping device job



## CreateDeviceLiveToolsPingDevice

> CreateDeviceLiveToolsPing201Response CreateDeviceLiveToolsPingDevice(ctx, serial).CreateDeviceLiveToolsPingDeviceRequest(createDeviceLiveToolsPingDeviceRequest).Execute()

Enqueue a job to check connectivity status to the device



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
	createDeviceLiveToolsPingDeviceRequest := *openapiclient.NewCreateDeviceLiveToolsPingDeviceRequest() // CreateDeviceLiveToolsPingDeviceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PingDeviceAPI.CreateDeviceLiveToolsPingDevice(context.Background(), serial).CreateDeviceLiveToolsPingDeviceRequest(createDeviceLiveToolsPingDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PingDeviceAPI.CreateDeviceLiveToolsPingDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceLiveToolsPingDevice`: CreateDeviceLiveToolsPing201Response
	fmt.Fprintf(os.Stdout, "Response from `PingDeviceAPI.CreateDeviceLiveToolsPingDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsPingDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsPingDeviceRequest** | [**CreateDeviceLiveToolsPingDeviceRequest**](CreateDeviceLiveToolsPingDeviceRequest.md) |  | 

### Return type

[**CreateDeviceLiveToolsPing201Response**](CreateDeviceLiveToolsPing201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsPingDevice

> GetDeviceLiveToolsPingDevice200Response GetDeviceLiveToolsPingDevice(ctx, serial, id).Execute()

Return a ping device job



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
	id := "id_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PingDeviceAPI.GetDeviceLiveToolsPingDevice(context.Background(), serial, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PingDeviceAPI.GetDeviceLiveToolsPingDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceLiveToolsPingDevice`: GetDeviceLiveToolsPingDevice200Response
	fmt.Fprintf(os.Stdout, "Response from `PingDeviceAPI.GetDeviceLiveToolsPingDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsPingDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDeviceLiveToolsPingDevice200Response**](GetDeviceLiveToolsPingDevice200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

