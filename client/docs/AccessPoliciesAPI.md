# \AccessPoliciesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkSwitchAccessPolicy**](AccessPoliciesAPI.md#CreateNetworkSwitchAccessPolicy) | **Post** /networks/{networkId}/switch/accessPolicies | Create an access policy for a switch network
[**DeleteNetworkSwitchAccessPolicy**](AccessPoliciesAPI.md#DeleteNetworkSwitchAccessPolicy) | **Delete** /networks/{networkId}/switch/accessPolicies/{accessPolicyNumber} | Delete an access policy for a switch network
[**GetNetworkSwitchAccessPolicies**](AccessPoliciesAPI.md#GetNetworkSwitchAccessPolicies) | **Get** /networks/{networkId}/switch/accessPolicies | List the access policies for a switch network
[**GetNetworkSwitchAccessPolicy**](AccessPoliciesAPI.md#GetNetworkSwitchAccessPolicy) | **Get** /networks/{networkId}/switch/accessPolicies/{accessPolicyNumber} | Return a specific access policy for a switch network
[**UpdateNetworkSwitchAccessPolicy**](AccessPoliciesAPI.md#UpdateNetworkSwitchAccessPolicy) | **Put** /networks/{networkId}/switch/accessPolicies/{accessPolicyNumber} | Update an access policy for a switch network



## CreateNetworkSwitchAccessPolicy

> GetNetworkSwitchAccessPolicies200ResponseInner CreateNetworkSwitchAccessPolicy(ctx, networkId).CreateNetworkSwitchAccessPolicyRequest(createNetworkSwitchAccessPolicyRequest).Execute()

Create an access policy for a switch network



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
	createNetworkSwitchAccessPolicyRequest := *openapiclient.NewCreateNetworkSwitchAccessPolicyRequest("Name_example", []openapiclient.CreateNetworkSwitchAccessPolicyRequestRadiusServersInner{*openapiclient.NewCreateNetworkSwitchAccessPolicyRequestRadiusServersInner()}, false, false, false, "HostMode_example", false) // CreateNetworkSwitchAccessPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessPoliciesAPI.CreateNetworkSwitchAccessPolicy(context.Background(), networkId).CreateNetworkSwitchAccessPolicyRequest(createNetworkSwitchAccessPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessPoliciesAPI.CreateNetworkSwitchAccessPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSwitchAccessPolicy`: GetNetworkSwitchAccessPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AccessPoliciesAPI.CreateNetworkSwitchAccessPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchAccessPolicyRequest** | [**CreateNetworkSwitchAccessPolicyRequest**](CreateNetworkSwitchAccessPolicyRequest.md) |  | 

### Return type

[**GetNetworkSwitchAccessPolicies200ResponseInner**](GetNetworkSwitchAccessPolicies200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSwitchAccessPolicy

> DeleteNetworkSwitchAccessPolicy(ctx, networkId, accessPolicyNumber).Execute()

Delete an access policy for a switch network



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
	accessPolicyNumber := "accessPolicyNumber_example" // string | Access policy number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccessPoliciesAPI.DeleteNetworkSwitchAccessPolicy(context.Background(), networkId, accessPolicyNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessPoliciesAPI.DeleteNetworkSwitchAccessPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**accessPolicyNumber** | **string** | Access policy number | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchAccessPolicyRequest struct via the builder pattern


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


## GetNetworkSwitchAccessPolicies

> []GetNetworkSwitchAccessPolicies200ResponseInner GetNetworkSwitchAccessPolicies(ctx, networkId).Execute()

List the access policies for a switch network



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
	resp, r, err := apiClient.AccessPoliciesAPI.GetNetworkSwitchAccessPolicies(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessPoliciesAPI.GetNetworkSwitchAccessPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchAccessPolicies`: []GetNetworkSwitchAccessPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AccessPoliciesAPI.GetNetworkSwitchAccessPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchAccessPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkSwitchAccessPolicies200ResponseInner**](GetNetworkSwitchAccessPolicies200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchAccessPolicy

> GetNetworkSwitchAccessPolicies200ResponseInner GetNetworkSwitchAccessPolicy(ctx, networkId, accessPolicyNumber).Execute()

Return a specific access policy for a switch network



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
	accessPolicyNumber := "accessPolicyNumber_example" // string | Access policy number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessPoliciesAPI.GetNetworkSwitchAccessPolicy(context.Background(), networkId, accessPolicyNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessPoliciesAPI.GetNetworkSwitchAccessPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchAccessPolicy`: GetNetworkSwitchAccessPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AccessPoliciesAPI.GetNetworkSwitchAccessPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**accessPolicyNumber** | **string** | Access policy number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkSwitchAccessPolicies200ResponseInner**](GetNetworkSwitchAccessPolicies200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchAccessPolicy

> GetNetworkSwitchAccessPolicies200ResponseInner UpdateNetworkSwitchAccessPolicy(ctx, networkId, accessPolicyNumber).UpdateNetworkSwitchAccessPolicyRequest(updateNetworkSwitchAccessPolicyRequest).Execute()

Update an access policy for a switch network



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
	accessPolicyNumber := "accessPolicyNumber_example" // string | Access policy number
	updateNetworkSwitchAccessPolicyRequest := *openapiclient.NewUpdateNetworkSwitchAccessPolicyRequest() // UpdateNetworkSwitchAccessPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessPoliciesAPI.UpdateNetworkSwitchAccessPolicy(context.Background(), networkId, accessPolicyNumber).UpdateNetworkSwitchAccessPolicyRequest(updateNetworkSwitchAccessPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessPoliciesAPI.UpdateNetworkSwitchAccessPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchAccessPolicy`: GetNetworkSwitchAccessPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AccessPoliciesAPI.UpdateNetworkSwitchAccessPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**accessPolicyNumber** | **string** | Access policy number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSwitchAccessPolicyRequest** | [**UpdateNetworkSwitchAccessPolicyRequest**](UpdateNetworkSwitchAccessPolicyRequest.md) |  | 

### Return type

[**GetNetworkSwitchAccessPolicies200ResponseInner**](GetNetworkSwitchAccessPolicies200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

