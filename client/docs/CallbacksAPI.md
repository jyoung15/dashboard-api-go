# \CallbacksAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationWebhooksCallbacksStatus**](CallbacksAPI.md#GetOrganizationWebhooksCallbacksStatus) | **Get** /organizations/{organizationId}/webhooks/callbacks/statuses/{callbackId} | Return the status of an API callback



## GetOrganizationWebhooksCallbacksStatus

> GetOrganizationWebhooksCallbacksStatus200Response GetOrganizationWebhooksCallbacksStatus(ctx, organizationId, callbackId).Execute()

Return the status of an API callback



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
	organizationId := "organizationId_example" // string | Organization ID
	callbackId := "callbackId_example" // string | Callback ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CallbacksAPI.GetOrganizationWebhooksCallbacksStatus(context.Background(), organizationId, callbackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallbacksAPI.GetOrganizationWebhooksCallbacksStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationWebhooksCallbacksStatus`: GetOrganizationWebhooksCallbacksStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `CallbacksAPI.GetOrganizationWebhooksCallbacksStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**callbackId** | **string** | Callback ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWebhooksCallbacksStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationWebhooksCallbacksStatus200Response**](GetOrganizationWebhooksCallbacksStatus200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

