# \CustomPerformanceClassesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkApplianceTrafficShapingCustomPerformanceClass**](CustomPerformanceClassesAPI.md#CreateNetworkApplianceTrafficShapingCustomPerformanceClass) | **Post** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses | Add a custom performance class for an MX network
[**DeleteNetworkApplianceTrafficShapingCustomPerformanceClass**](CustomPerformanceClassesAPI.md#DeleteNetworkApplianceTrafficShapingCustomPerformanceClass) | **Delete** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId} | Delete a custom performance class from an MX network
[**GetNetworkApplianceTrafficShapingCustomPerformanceClass**](CustomPerformanceClassesAPI.md#GetNetworkApplianceTrafficShapingCustomPerformanceClass) | **Get** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId} | Return a custom performance class for an MX network
[**GetNetworkApplianceTrafficShapingCustomPerformanceClasses**](CustomPerformanceClassesAPI.md#GetNetworkApplianceTrafficShapingCustomPerformanceClasses) | **Get** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses | List all custom performance classes for an MX network
[**UpdateNetworkApplianceTrafficShapingCustomPerformanceClass**](CustomPerformanceClassesAPI.md#UpdateNetworkApplianceTrafficShapingCustomPerformanceClass) | **Put** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId} | Update a custom performance class for an MX network



## CreateNetworkApplianceTrafficShapingCustomPerformanceClass

> GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner CreateNetworkApplianceTrafficShapingCustomPerformanceClass(ctx, networkId).CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest(createNetworkApplianceTrafficShapingCustomPerformanceClassRequest).Execute()

Add a custom performance class for an MX network



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
	createNetworkApplianceTrafficShapingCustomPerformanceClassRequest := *openapiclient.NewCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest("Name_example") // CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomPerformanceClassesAPI.CreateNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId).CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest(createNetworkApplianceTrafficShapingCustomPerformanceClassRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPerformanceClassesAPI.CreateNetworkApplianceTrafficShapingCustomPerformanceClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkApplianceTrafficShapingCustomPerformanceClass`: GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CustomPerformanceClassesAPI.CreateNetworkApplianceTrafficShapingCustomPerformanceClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkApplianceTrafficShapingCustomPerformanceClassRequest** | [**CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest**](CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest.md) |  | 

### Return type

[**GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner**](GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkApplianceTrafficShapingCustomPerformanceClass

> DeleteNetworkApplianceTrafficShapingCustomPerformanceClass(ctx, networkId, customPerformanceClassId).Execute()

Delete a custom performance class from an MX network



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
	customPerformanceClassId := "customPerformanceClassId_example" // string | Custom performance class ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomPerformanceClassesAPI.DeleteNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId, customPerformanceClassId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPerformanceClassesAPI.DeleteNetworkApplianceTrafficShapingCustomPerformanceClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**customPerformanceClassId** | **string** | Custom performance class ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceTrafficShapingCustomPerformanceClass

> GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner GetNetworkApplianceTrafficShapingCustomPerformanceClass(ctx, networkId, customPerformanceClassId).Execute()

Return a custom performance class for an MX network



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
	customPerformanceClassId := "customPerformanceClassId_example" // string | Custom performance class ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomPerformanceClassesAPI.GetNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId, customPerformanceClassId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPerformanceClassesAPI.GetNetworkApplianceTrafficShapingCustomPerformanceClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceTrafficShapingCustomPerformanceClass`: GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CustomPerformanceClassesAPI.GetNetworkApplianceTrafficShapingCustomPerformanceClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**customPerformanceClassId** | **string** | Custom performance class ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner**](GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceTrafficShapingCustomPerformanceClasses

> []GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner GetNetworkApplianceTrafficShapingCustomPerformanceClasses(ctx, networkId).Execute()

List all custom performance classes for an MX network



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
	resp, r, err := apiClient.CustomPerformanceClassesAPI.GetNetworkApplianceTrafficShapingCustomPerformanceClasses(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPerformanceClassesAPI.GetNetworkApplianceTrafficShapingCustomPerformanceClasses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceTrafficShapingCustomPerformanceClasses`: []GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CustomPerformanceClassesAPI.GetNetworkApplianceTrafficShapingCustomPerformanceClasses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingCustomPerformanceClassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner**](GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceTrafficShapingCustomPerformanceClass

> GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner UpdateNetworkApplianceTrafficShapingCustomPerformanceClass(ctx, networkId, customPerformanceClassId).UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest(updateNetworkApplianceTrafficShapingCustomPerformanceClassRequest).Execute()

Update a custom performance class for an MX network



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
	customPerformanceClassId := "customPerformanceClassId_example" // string | Custom performance class ID
	updateNetworkApplianceTrafficShapingCustomPerformanceClassRequest := *openapiclient.NewUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest() // UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomPerformanceClassesAPI.UpdateNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId, customPerformanceClassId).UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest(updateNetworkApplianceTrafficShapingCustomPerformanceClassRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPerformanceClassesAPI.UpdateNetworkApplianceTrafficShapingCustomPerformanceClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceTrafficShapingCustomPerformanceClass`: GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CustomPerformanceClassesAPI.UpdateNetworkApplianceTrafficShapingCustomPerformanceClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**customPerformanceClassId** | **string** | Custom performance class ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceTrafficShapingCustomPerformanceClassRequest** | [**UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest**](UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest.md) |  | 

### Return type

[**GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner**](GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

