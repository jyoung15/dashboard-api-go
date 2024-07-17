# \EntitlementsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdministeredLicensingSubscriptionEntitlements**](EntitlementsAPI.md#GetAdministeredLicensingSubscriptionEntitlements) | **Get** /administered/licensing/subscription/entitlements | Retrieve the list of purchasable entitlements



## GetAdministeredLicensingSubscriptionEntitlements

> GetAdministeredLicensingSubscriptionEntitlements200Response GetAdministeredLicensingSubscriptionEntitlements(ctx).Skus(skus).Execute()

Retrieve the list of purchasable entitlements



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
	skus := []string{"Inner_example"} // []string | Filter to entitlements with the specified SKUs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.GetAdministeredLicensingSubscriptionEntitlements(context.Background()).Skus(skus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetAdministeredLicensingSubscriptionEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdministeredLicensingSubscriptionEntitlements`: GetAdministeredLicensingSubscriptionEntitlements200Response
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetAdministeredLicensingSubscriptionEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdministeredLicensingSubscriptionEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skus** | **[]string** | Filter to entitlements with the specified SKUs | 

### Return type

[**GetAdministeredLicensingSubscriptionEntitlements200Response**](GetAdministeredLicensingSubscriptionEntitlements200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

