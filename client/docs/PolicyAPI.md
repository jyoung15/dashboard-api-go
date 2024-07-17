# \PolicyAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkClientPolicy**](PolicyAPI.md#GetNetworkClientPolicy) | **Get** /networks/{networkId}/clients/{clientId}/policy | Return the policy assigned to a client on the network
[**UpdateNetworkClientPolicy**](PolicyAPI.md#UpdateNetworkClientPolicy) | **Put** /networks/{networkId}/clients/{clientId}/policy | Update the policy assigned to a client on the network



## GetNetworkClientPolicy

> GetNetworkClientPolicy200Response GetNetworkClientPolicy(ctx, networkId, clientId).Execute()

Return the policy assigned to a client on the network



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
	clientId := "clientId_example" // string | Client ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.GetNetworkClientPolicy(context.Background(), networkId, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.GetNetworkClientPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkClientPolicy`: GetNetworkClientPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.GetNetworkClientPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkClientPolicy200Response**](GetNetworkClientPolicy200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkClientPolicy

> GetNetworkClientPolicy200Response UpdateNetworkClientPolicy(ctx, networkId, clientId).UpdateNetworkClientPolicyRequest(updateNetworkClientPolicyRequest).Execute()

Update the policy assigned to a client on the network



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
	clientId := "clientId_example" // string | Client ID
	updateNetworkClientPolicyRequest := *openapiclient.NewUpdateNetworkClientPolicyRequest("DevicePolicy_example") // UpdateNetworkClientPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyAPI.UpdateNetworkClientPolicy(context.Background(), networkId, clientId).UpdateNetworkClientPolicyRequest(updateNetworkClientPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyAPI.UpdateNetworkClientPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkClientPolicy`: GetNetworkClientPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `PolicyAPI.UpdateNetworkClientPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkClientPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkClientPolicyRequest** | [**UpdateNetworkClientPolicyRequest**](UpdateNetworkClientPolicyRequest.md) |  | 

### Return type

[**GetNetworkClientPolicy200Response**](GetNetworkClientPolicy200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

