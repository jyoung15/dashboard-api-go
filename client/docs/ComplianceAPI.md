# \ComplianceAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses**](ComplianceAPI.md#GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses) | **Get** /administered/licensing/subscription/subscriptions/compliance/statuses | Get compliance status for requested subscriptions



## GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses

> []GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInner GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses(ctx).OrganizationIds(organizationIds).SubscriptionIds(subscriptionIds).Execute()

Get compliance status for requested subscriptions



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
	organizationIds := []string{"Inner_example"} // []string | Organizations to get subscription compliance information for
	subscriptionIds := []string{"Inner_example"} // []string | Subscription ids (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceAPI.GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses(context.Background()).OrganizationIds(organizationIds).SubscriptionIds(subscriptionIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceAPI.GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses`: []GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ComplianceAPI.GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdministeredLicensingSubscriptionSubscriptionsComplianceStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationIds** | **[]string** | Organizations to get subscription compliance information for | 
 **subscriptionIds** | **[]string** | Subscription ids | 

### Return type

[**[]GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInner**](GetAdministeredLicensingSubscriptionSubscriptionsComplianceStatuses200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

