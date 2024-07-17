# \AlertTypesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationWebhooksAlertTypes**](AlertTypesAPI.md#GetOrganizationWebhooksAlertTypes) | **Get** /organizations/{organizationId}/webhooks/alertTypes | Return a list of alert types to be used with managing webhook alerts



## GetOrganizationWebhooksAlertTypes

> GetOrganizationWebhooksAlertTypes200Response GetOrganizationWebhooksAlertTypes(ctx, organizationId).ProductType(productType).Execute()

Return a list of alert types to be used with managing webhook alerts



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
	productType := "productType_example" // string | Filter sample alerts to a specific product type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertTypesAPI.GetOrganizationWebhooksAlertTypes(context.Background(), organizationId).ProductType(productType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertTypesAPI.GetOrganizationWebhooksAlertTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationWebhooksAlertTypes`: GetOrganizationWebhooksAlertTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertTypesAPI.GetOrganizationWebhooksAlertTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWebhooksAlertTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productType** | **string** | Filter sample alerts to a specific product type | 

### Return type

[**GetOrganizationWebhooksAlertTypes200Response**](GetOrganizationWebhooksAlertTypes200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
