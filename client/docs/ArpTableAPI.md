# \ArpTableAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceLiveToolsArpTable**](ArpTableAPI.md#CreateDeviceLiveToolsArpTable) | **Post** /devices/{serial}/liveTools/arpTable | Enqueue a job to perform a ARP table request for the device
[**GetDeviceLiveToolsArpTable**](ArpTableAPI.md#GetDeviceLiveToolsArpTable) | **Get** /devices/{serial}/liveTools/arpTable/{arpTableId} | Return an ARP table live tool job.



## CreateDeviceLiveToolsArpTable

> CreateDeviceLiveToolsArpTable201Response CreateDeviceLiveToolsArpTable(ctx, serial).CreateDeviceLiveToolsArpTableRequest(createDeviceLiveToolsArpTableRequest).Execute()

Enqueue a job to perform a ARP table request for the device



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
	resp, r, err := apiClient.ArpTableAPI.CreateDeviceLiveToolsArpTable(context.Background(), serial).CreateDeviceLiveToolsArpTableRequest(createDeviceLiveToolsArpTableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArpTableAPI.CreateDeviceLiveToolsArpTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceLiveToolsArpTable`: CreateDeviceLiveToolsArpTable201Response
	fmt.Fprintf(os.Stdout, "Response from `ArpTableAPI.CreateDeviceLiveToolsArpTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsArpTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsArpTableRequest** | [**CreateDeviceLiveToolsArpTableRequest**](CreateDeviceLiveToolsArpTableRequest.md) |  | 

### Return type

[**CreateDeviceLiveToolsArpTable201Response**](CreateDeviceLiveToolsArpTable201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsArpTable

> DevicesSerialLiveToolsArpTablePostRequestMessage GetDeviceLiveToolsArpTable(ctx, serial, arpTableId).Execute()

Return an ARP table live tool job.



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
	arpTableId := "arpTableId_example" // string | Arp table ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArpTableAPI.GetDeviceLiveToolsArpTable(context.Background(), serial, arpTableId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArpTableAPI.GetDeviceLiveToolsArpTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceLiveToolsArpTable`: DevicesSerialLiveToolsArpTablePostRequestMessage
	fmt.Fprintf(os.Stdout, "Response from `ArpTableAPI.GetDeviceLiveToolsArpTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**arpTableId** | **string** | Arp table ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsArpTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DevicesSerialLiveToolsArpTablePostRequestMessage**](DevicesSerialLiveToolsArpTablePostRequestMessage.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

