# \BypassActivationLockAttemptsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkSmBypassActivationLockAttempt**](BypassActivationLockAttemptsAPI.md#CreateNetworkSmBypassActivationLockAttempt) | **Post** /networks/{networkId}/sm/bypassActivationLockAttempts | Bypass activation lock attempt
[**GetNetworkSmBypassActivationLockAttempt**](BypassActivationLockAttemptsAPI.md#GetNetworkSmBypassActivationLockAttempt) | **Get** /networks/{networkId}/sm/bypassActivationLockAttempts/{attemptId} | Bypass activation lock attempt status



## CreateNetworkSmBypassActivationLockAttempt

> map[string]interface{} CreateNetworkSmBypassActivationLockAttempt(ctx, networkId).CreateNetworkSmBypassActivationLockAttemptRequest(createNetworkSmBypassActivationLockAttemptRequest).Execute()

Bypass activation lock attempt



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
	createNetworkSmBypassActivationLockAttemptRequest := *openapiclient.NewCreateNetworkSmBypassActivationLockAttemptRequest([]string{"Ids_example"}) // CreateNetworkSmBypassActivationLockAttemptRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BypassActivationLockAttemptsAPI.CreateNetworkSmBypassActivationLockAttempt(context.Background(), networkId).CreateNetworkSmBypassActivationLockAttemptRequest(createNetworkSmBypassActivationLockAttemptRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BypassActivationLockAttemptsAPI.CreateNetworkSmBypassActivationLockAttempt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSmBypassActivationLockAttempt`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BypassActivationLockAttemptsAPI.CreateNetworkSmBypassActivationLockAttempt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSmBypassActivationLockAttemptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSmBypassActivationLockAttemptRequest** | [**CreateNetworkSmBypassActivationLockAttemptRequest**](CreateNetworkSmBypassActivationLockAttemptRequest.md) |  | 

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


## GetNetworkSmBypassActivationLockAttempt

> map[string]interface{} GetNetworkSmBypassActivationLockAttempt(ctx, networkId, attemptId).Execute()

Bypass activation lock attempt status



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
	attemptId := "attemptId_example" // string | Attempt ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BypassActivationLockAttemptsAPI.GetNetworkSmBypassActivationLockAttempt(context.Background(), networkId, attemptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BypassActivationLockAttemptsAPI.GetNetworkSmBypassActivationLockAttempt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmBypassActivationLockAttempt`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BypassActivationLockAttemptsAPI.GetNetworkSmBypassActivationLockAttempt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**attemptId** | **string** | Attempt ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmBypassActivationLockAttemptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

