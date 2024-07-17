# \PrefixesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkAppliancePrefixesDelegatedStatic**](PrefixesAPI.md#CreateNetworkAppliancePrefixesDelegatedStatic) | **Post** /networks/{networkId}/appliance/prefixes/delegated/statics | Add a static delegated prefix from a network
[**DeleteNetworkAppliancePrefixesDelegatedStatic**](PrefixesAPI.md#DeleteNetworkAppliancePrefixesDelegatedStatic) | **Delete** /networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId} | Delete a static delegated prefix from a network
[**GetDeviceAppliancePrefixesDelegated**](PrefixesAPI.md#GetDeviceAppliancePrefixesDelegated) | **Get** /devices/{serial}/appliance/prefixes/delegated | Return current delegated IPv6 prefixes on an appliance.
[**GetDeviceAppliancePrefixesDelegatedVlanAssignments**](PrefixesAPI.md#GetDeviceAppliancePrefixesDelegatedVlanAssignments) | **Get** /devices/{serial}/appliance/prefixes/delegated/vlanAssignments | Return prefixes assigned to all IPv6 enabled VLANs on an appliance.
[**GetNetworkAppliancePrefixesDelegatedStatic**](PrefixesAPI.md#GetNetworkAppliancePrefixesDelegatedStatic) | **Get** /networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId} | Return a static delegated prefix from a network
[**GetNetworkAppliancePrefixesDelegatedStatics**](PrefixesAPI.md#GetNetworkAppliancePrefixesDelegatedStatics) | **Get** /networks/{networkId}/appliance/prefixes/delegated/statics | List static delegated prefixes for a network
[**UpdateNetworkAppliancePrefixesDelegatedStatic**](PrefixesAPI.md#UpdateNetworkAppliancePrefixesDelegatedStatic) | **Put** /networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId} | Update a static delegated prefix from a network



## CreateNetworkAppliancePrefixesDelegatedStatic

> map[string]interface{} CreateNetworkAppliancePrefixesDelegatedStatic(ctx, networkId).CreateNetworkAppliancePrefixesDelegatedStaticRequest(createNetworkAppliancePrefixesDelegatedStaticRequest).Execute()

Add a static delegated prefix from a network



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
	createNetworkAppliancePrefixesDelegatedStaticRequest := *openapiclient.NewCreateNetworkAppliancePrefixesDelegatedStaticRequest("Prefix_example", *openapiclient.NewCreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin()) // CreateNetworkAppliancePrefixesDelegatedStaticRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrefixesAPI.CreateNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId).CreateNetworkAppliancePrefixesDelegatedStaticRequest(createNetworkAppliancePrefixesDelegatedStaticRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrefixesAPI.CreateNetworkAppliancePrefixesDelegatedStatic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkAppliancePrefixesDelegatedStatic`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrefixesAPI.CreateNetworkAppliancePrefixesDelegatedStatic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkAppliancePrefixesDelegatedStaticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkAppliancePrefixesDelegatedStaticRequest** | [**CreateNetworkAppliancePrefixesDelegatedStaticRequest**](CreateNetworkAppliancePrefixesDelegatedStaticRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkAppliancePrefixesDelegatedStatic

> DeleteNetworkAppliancePrefixesDelegatedStatic(ctx, networkId, staticDelegatedPrefixId).Execute()

Delete a static delegated prefix from a network



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
	staticDelegatedPrefixId := "staticDelegatedPrefixId_example" // string | Static delegated prefix ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PrefixesAPI.DeleteNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId, staticDelegatedPrefixId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrefixesAPI.DeleteNetworkAppliancePrefixesDelegatedStatic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticDelegatedPrefixId** | **string** | Static delegated prefix ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkAppliancePrefixesDelegatedStaticRequest struct via the builder pattern


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


## GetDeviceAppliancePrefixesDelegated

> []map[string]interface{} GetDeviceAppliancePrefixesDelegated(ctx, serial).Execute()

Return current delegated IPv6 prefixes on an appliance.



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
	resp, r, err := apiClient.PrefixesAPI.GetDeviceAppliancePrefixesDelegated(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrefixesAPI.GetDeviceAppliancePrefixesDelegated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceAppliancePrefixesDelegated`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrefixesAPI.GetDeviceAppliancePrefixesDelegated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceAppliancePrefixesDelegatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceAppliancePrefixesDelegatedVlanAssignments

> []map[string]interface{} GetDeviceAppliancePrefixesDelegatedVlanAssignments(ctx, serial).Execute()

Return prefixes assigned to all IPv6 enabled VLANs on an appliance.



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
	resp, r, err := apiClient.PrefixesAPI.GetDeviceAppliancePrefixesDelegatedVlanAssignments(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrefixesAPI.GetDeviceAppliancePrefixesDelegatedVlanAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceAppliancePrefixesDelegatedVlanAssignments`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrefixesAPI.GetDeviceAppliancePrefixesDelegatedVlanAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAppliancePrefixesDelegatedStatic

> GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner GetNetworkAppliancePrefixesDelegatedStatic(ctx, networkId, staticDelegatedPrefixId).Execute()

Return a static delegated prefix from a network



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
	staticDelegatedPrefixId := "staticDelegatedPrefixId_example" // string | Static delegated prefix ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrefixesAPI.GetNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId, staticDelegatedPrefixId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrefixesAPI.GetNetworkAppliancePrefixesDelegatedStatic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkAppliancePrefixesDelegatedStatic`: GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PrefixesAPI.GetNetworkAppliancePrefixesDelegatedStatic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticDelegatedPrefixId** | **string** | Static delegated prefix ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAppliancePrefixesDelegatedStaticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner**](GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAppliancePrefixesDelegatedStatics

> []GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner GetNetworkAppliancePrefixesDelegatedStatics(ctx, networkId).Execute()

List static delegated prefixes for a network



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
	resp, r, err := apiClient.PrefixesAPI.GetNetworkAppliancePrefixesDelegatedStatics(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrefixesAPI.GetNetworkAppliancePrefixesDelegatedStatics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkAppliancePrefixesDelegatedStatics`: []GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PrefixesAPI.GetNetworkAppliancePrefixesDelegatedStatics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAppliancePrefixesDelegatedStaticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner**](GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkAppliancePrefixesDelegatedStatic

> map[string]interface{} UpdateNetworkAppliancePrefixesDelegatedStatic(ctx, networkId, staticDelegatedPrefixId).UpdateNetworkAppliancePrefixesDelegatedStaticRequest(updateNetworkAppliancePrefixesDelegatedStaticRequest).Execute()

Update a static delegated prefix from a network



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
	staticDelegatedPrefixId := "staticDelegatedPrefixId_example" // string | Static delegated prefix ID
	updateNetworkAppliancePrefixesDelegatedStaticRequest := *openapiclient.NewUpdateNetworkAppliancePrefixesDelegatedStaticRequest() // UpdateNetworkAppliancePrefixesDelegatedStaticRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrefixesAPI.UpdateNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId, staticDelegatedPrefixId).UpdateNetworkAppliancePrefixesDelegatedStaticRequest(updateNetworkAppliancePrefixesDelegatedStaticRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrefixesAPI.UpdateNetworkAppliancePrefixesDelegatedStatic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkAppliancePrefixesDelegatedStatic`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrefixesAPI.UpdateNetworkAppliancePrefixesDelegatedStatic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticDelegatedPrefixId** | **string** | Static delegated prefix ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkAppliancePrefixesDelegatedStaticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkAppliancePrefixesDelegatedStaticRequest** | [**UpdateNetworkAppliancePrefixesDelegatedStaticRequest**](UpdateNetworkAppliancePrefixesDelegatedStaticRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

