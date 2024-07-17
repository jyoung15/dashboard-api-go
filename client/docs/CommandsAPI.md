# \CommandsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceSensorCommand**](CommandsAPI.md#CreateDeviceSensorCommand) | **Post** /devices/{serial}/sensor/commands | Sends a command to a sensor
[**GetDeviceSensorCommand**](CommandsAPI.md#GetDeviceSensorCommand) | **Get** /devices/{serial}/sensor/commands/{commandId} | Returns information about the command&#39;s execution, including the status
[**GetDeviceSensorCommands**](CommandsAPI.md#GetDeviceSensorCommands) | **Get** /devices/{serial}/sensor/commands | Returns a historical log of all commands



## CreateDeviceSensorCommand

> GetDeviceSensorCommands200ResponseInner CreateDeviceSensorCommand(ctx, serial).CreateDeviceSensorCommandRequest(createDeviceSensorCommandRequest).Execute()

Sends a command to a sensor



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
	createDeviceSensorCommandRequest := *openapiclient.NewCreateDeviceSensorCommandRequest("Operation_example") // CreateDeviceSensorCommandRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommandsAPI.CreateDeviceSensorCommand(context.Background(), serial).CreateDeviceSensorCommandRequest(createDeviceSensorCommandRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommandsAPI.CreateDeviceSensorCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceSensorCommand`: GetDeviceSensorCommands200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CommandsAPI.CreateDeviceSensorCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceSensorCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceSensorCommandRequest** | [**CreateDeviceSensorCommandRequest**](CreateDeviceSensorCommandRequest.md) |  | 

### Return type

[**GetDeviceSensorCommands200ResponseInner**](GetDeviceSensorCommands200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSensorCommand

> GetDeviceSensorCommands200ResponseInner GetDeviceSensorCommand(ctx, serial, commandId).Execute()

Returns information about the command's execution, including the status



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
	commandId := "commandId_example" // string | Command ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommandsAPI.GetDeviceSensorCommand(context.Background(), serial, commandId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommandsAPI.GetDeviceSensorCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSensorCommand`: GetDeviceSensorCommands200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CommandsAPI.GetDeviceSensorCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**commandId** | **string** | Command ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSensorCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDeviceSensorCommands200ResponseInner**](GetDeviceSensorCommands200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceSensorCommands

> []GetDeviceSensorCommands200ResponseInner GetDeviceSensorCommands(ctx, serial).Operations(operations).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).T0(t0).T1(t1).Timespan(timespan).Execute()

Returns a historical log of all commands



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
	operations := []string{"Inner_example"} // []string | Optional parameter to filter commands by operation. Allowed values are disableDownstreamPower, enableDownstreamPower, cycleDownstreamPower, and refreshData. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	sortOrder := "sortOrder_example" // string | Sorted order of entries. Order options are 'ascending' and 'descending'. Default is 'descending'. (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 30 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 30 days. The default is 30 days. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommandsAPI.GetDeviceSensorCommands(context.Background(), serial).Operations(operations).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommandsAPI.GetDeviceSensorCommands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSensorCommands`: []GetDeviceSensorCommands200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CommandsAPI.GetDeviceSensorCommands`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSensorCommandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **operations** | **[]string** | Optional parameter to filter commands by operation. Allowed values are disableDownstreamPower, enableDownstreamPower, cycleDownstreamPower, and refreshData. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 10. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of entries. Order options are &#39;ascending&#39; and &#39;descending&#39;. Default is &#39;descending&#39;. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 30 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 30 days. The default is 30 days. | 

### Return type

[**[]GetDeviceSensorCommands200ResponseInner**](GetDeviceSensorCommands200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

