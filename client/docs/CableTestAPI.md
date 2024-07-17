# \CableTestAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceLiveToolsCableTest**](CableTestAPI.md#CreateDeviceLiveToolsCableTest) | **Post** /devices/{serial}/liveTools/cableTest | Enqueue a job to perform a cable test for the device on the specified ports
[**GetDeviceLiveToolsCableTest**](CableTestAPI.md#GetDeviceLiveToolsCableTest) | **Get** /devices/{serial}/liveTools/cableTest/{id} | Return a cable test live tool job.



## CreateDeviceLiveToolsCableTest

> CreateDeviceLiveToolsCableTest201Response CreateDeviceLiveToolsCableTest(ctx, serial).CreateDeviceLiveToolsCableTestRequest(createDeviceLiveToolsCableTestRequest).Execute()

Enqueue a job to perform a cable test for the device on the specified ports



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
	createDeviceLiveToolsCableTestRequest := *openapiclient.NewCreateDeviceLiveToolsCableTestRequest([]string{"Ports_example"}) // CreateDeviceLiveToolsCableTestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CableTestAPI.CreateDeviceLiveToolsCableTest(context.Background(), serial).CreateDeviceLiveToolsCableTestRequest(createDeviceLiveToolsCableTestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CableTestAPI.CreateDeviceLiveToolsCableTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceLiveToolsCableTest`: CreateDeviceLiveToolsCableTest201Response
	fmt.Fprintf(os.Stdout, "Response from `CableTestAPI.CreateDeviceLiveToolsCableTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsCableTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsCableTestRequest** | [**CreateDeviceLiveToolsCableTestRequest**](CreateDeviceLiveToolsCableTestRequest.md) |  | 

### Return type

[**CreateDeviceLiveToolsCableTest201Response**](CreateDeviceLiveToolsCableTest201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsCableTest

> DevicesSerialLiveToolsCableTestPostRequestMessage GetDeviceLiveToolsCableTest(ctx, serial, id).Execute()

Return a cable test live tool job.



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
	resp, r, err := apiClient.CableTestAPI.GetDeviceLiveToolsCableTest(context.Background(), serial, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CableTestAPI.GetDeviceLiveToolsCableTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceLiveToolsCableTest`: DevicesSerialLiveToolsCableTestPostRequestMessage
	fmt.Fprintf(os.Stdout, "Response from `CableTestAPI.GetDeviceLiveToolsCableTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsCableTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DevicesSerialLiveToolsCableTestPostRequestMessage**](DevicesSerialLiveToolsCableTestPostRequestMessage.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

