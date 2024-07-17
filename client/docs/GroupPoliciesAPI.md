# \GroupPoliciesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkGroupPolicy**](GroupPoliciesAPI.md#CreateNetworkGroupPolicy) | **Post** /networks/{networkId}/groupPolicies | Create a group policy
[**DeleteNetworkGroupPolicy**](GroupPoliciesAPI.md#DeleteNetworkGroupPolicy) | **Delete** /networks/{networkId}/groupPolicies/{groupPolicyId} | Delete a group policy
[**GetNetworkGroupPolicies**](GroupPoliciesAPI.md#GetNetworkGroupPolicies) | **Get** /networks/{networkId}/groupPolicies | List the group policies in a network
[**GetNetworkGroupPolicy**](GroupPoliciesAPI.md#GetNetworkGroupPolicy) | **Get** /networks/{networkId}/groupPolicies/{groupPolicyId} | Display a group policy
[**UpdateNetworkGroupPolicy**](GroupPoliciesAPI.md#UpdateNetworkGroupPolicy) | **Put** /networks/{networkId}/groupPolicies/{groupPolicyId} | Update a group policy



## CreateNetworkGroupPolicy

> GetNetworkGroupPolicies200ResponseInner CreateNetworkGroupPolicy(ctx, networkId).CreateNetworkGroupPolicyRequest(createNetworkGroupPolicyRequest).Execute()

Create a group policy



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
	createNetworkGroupPolicyRequest := *openapiclient.NewCreateNetworkGroupPolicyRequest("Name_example") // CreateNetworkGroupPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupPoliciesAPI.CreateNetworkGroupPolicy(context.Background(), networkId).CreateNetworkGroupPolicyRequest(createNetworkGroupPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupPoliciesAPI.CreateNetworkGroupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkGroupPolicy`: GetNetworkGroupPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `GroupPoliciesAPI.CreateNetworkGroupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkGroupPolicyRequest** | [**CreateNetworkGroupPolicyRequest**](CreateNetworkGroupPolicyRequest.md) |  | 

### Return type

[**GetNetworkGroupPolicies200ResponseInner**](GetNetworkGroupPolicies200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkGroupPolicy

> DeleteNetworkGroupPolicy(ctx, networkId, groupPolicyId).Force(force).Execute()

Delete a group policy



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
	groupPolicyId := "groupPolicyId_example" // string | Group policy ID
	force := true // bool | If true, the system deletes the GP even if there are active clients using the GP. After deletion, active clients that were assigned to that Group Policy will be left without any policy applied. Default is false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupPoliciesAPI.DeleteNetworkGroupPolicy(context.Background(), networkId, groupPolicyId).Force(force).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupPoliciesAPI.DeleteNetworkGroupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupPolicyId** | **string** | Group policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **bool** | If true, the system deletes the GP even if there are active clients using the GP. After deletion, active clients that were assigned to that Group Policy will be left without any policy applied. Default is false. | 

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


## GetNetworkGroupPolicies

> []GetNetworkGroupPolicies200ResponseInner GetNetworkGroupPolicies(ctx, networkId).Execute()

List the group policies in a network



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
	resp, r, err := apiClient.GroupPoliciesAPI.GetNetworkGroupPolicies(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupPoliciesAPI.GetNetworkGroupPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkGroupPolicies`: []GetNetworkGroupPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `GroupPoliciesAPI.GetNetworkGroupPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkGroupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkGroupPolicies200ResponseInner**](GetNetworkGroupPolicies200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkGroupPolicy

> GetNetworkGroupPolicies200ResponseInner GetNetworkGroupPolicy(ctx, networkId, groupPolicyId).Execute()

Display a group policy



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
	groupPolicyId := "groupPolicyId_example" // string | Group policy ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupPoliciesAPI.GetNetworkGroupPolicy(context.Background(), networkId, groupPolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupPoliciesAPI.GetNetworkGroupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkGroupPolicy`: GetNetworkGroupPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `GroupPoliciesAPI.GetNetworkGroupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupPolicyId** | **string** | Group policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkGroupPolicies200ResponseInner**](GetNetworkGroupPolicies200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkGroupPolicy

> GetNetworkGroupPolicies200ResponseInner UpdateNetworkGroupPolicy(ctx, networkId, groupPolicyId).UpdateNetworkGroupPolicyRequest(updateNetworkGroupPolicyRequest).Execute()

Update a group policy



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
	groupPolicyId := "groupPolicyId_example" // string | Group policy ID
	updateNetworkGroupPolicyRequest := *openapiclient.NewUpdateNetworkGroupPolicyRequest() // UpdateNetworkGroupPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupPoliciesAPI.UpdateNetworkGroupPolicy(context.Background(), networkId, groupPolicyId).UpdateNetworkGroupPolicyRequest(updateNetworkGroupPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupPoliciesAPI.UpdateNetworkGroupPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkGroupPolicy`: GetNetworkGroupPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `GroupPoliciesAPI.UpdateNetworkGroupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupPolicyId** | **string** | Group policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkGroupPolicyRequest** | [**UpdateNetworkGroupPolicyRequest**](UpdateNetworkGroupPolicyRequest.md) |  | 

### Return type

[**GetNetworkGroupPolicies200ResponseInner**](GetNetworkGroupPolicies200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

