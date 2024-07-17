# \ThroughputTestAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceLiveToolsThroughputTest**](ThroughputTestAPI.md#CreateDeviceLiveToolsThroughputTest) | **Post** /devices/{serial}/liveTools/throughputTest | Enqueue a job to test a device throughput, the test will run for 10 secs to test throughput
[**GetDeviceLiveToolsThroughputTest**](ThroughputTestAPI.md#GetDeviceLiveToolsThroughputTest) | **Get** /devices/{serial}/liveTools/throughputTest/{throughputTestId} | Return a throughput test job



## CreateDeviceLiveToolsThroughputTest

> CreateDeviceLiveToolsThroughputTest201Response CreateDeviceLiveToolsThroughputTest(ctx, serial).CreateDeviceLiveToolsArpTableRequest(createDeviceLiveToolsArpTableRequest).Execute()

Enqueue a job to test a device throughput, the test will run for 10 secs to test throughput



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
	createDeviceLiveToolsArpTableRequest := *openapiclient.NewCreateDeviceLiveToolsArpTableRequest() // CreateDeviceLiveToolsArpTableRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThroughputTestAPI.CreateDeviceLiveToolsThroughputTest(context.Background(), serial).CreateDeviceLiveToolsArpTableRequest(createDeviceLiveToolsArpTableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThroughputTestAPI.CreateDeviceLiveToolsThroughputTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceLiveToolsThroughputTest`: CreateDeviceLiveToolsThroughputTest201Response
	fmt.Fprintf(os.Stdout, "Response from `ThroughputTestAPI.CreateDeviceLiveToolsThroughputTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsThroughputTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsArpTableRequest** | [**CreateDeviceLiveToolsArpTableRequest**](CreateDeviceLiveToolsArpTableRequest.md) |  | 

### Return type

[**CreateDeviceLiveToolsThroughputTest201Response**](CreateDeviceLiveToolsThroughputTest201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsThroughputTest

> DevicesSerialLiveToolsThroughputTestPostRequestMessage GetDeviceLiveToolsThroughputTest(ctx, serial, throughputTestId).Execute()

Return a throughput test job



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
	throughputTestId := "throughputTestId_example" // string | Throughput test ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThroughputTestAPI.GetDeviceLiveToolsThroughputTest(context.Background(), serial, throughputTestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThroughputTestAPI.GetDeviceLiveToolsThroughputTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceLiveToolsThroughputTest`: DevicesSerialLiveToolsThroughputTestPostRequestMessage
	fmt.Fprintf(os.Stdout, "Response from `ThroughputTestAPI.GetDeviceLiveToolsThroughputTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**throughputTestId** | **string** | Throughput test ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsThroughputTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DevicesSerialLiveToolsThroughputTestPostRequestMessage**](DevicesSerialLiveToolsThroughputTestPostRequestMessage.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

