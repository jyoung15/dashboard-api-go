# \ClaimKeyAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey**](ClaimKeyAPI.md#ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey) | **Post** /administered/licensing/subscription/subscriptions/claimKey/validate | Find a subscription by claim key



## ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey

> GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey(ctx).ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest(validateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest).Execute()

Find a subscription by claim key



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
	validateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest := *openapiclient.NewValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest("ClaimKey_example") // ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClaimKeyAPI.ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey(context.Background()).ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest(validateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClaimKeyAPI.ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey`: GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ClaimKeyAPI.ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest** | [**ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest**](ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest.md) |  | 

### Return type

[**GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner**](GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

