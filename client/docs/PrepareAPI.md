# \PrepareAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationInventoryOnboardingCloudMonitoringPrepare**](PrepareAPI.md#CreateOrganizationInventoryOnboardingCloudMonitoringPrepare) | **Post** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/prepare | Initiates or updates an import session



## CreateOrganizationInventoryOnboardingCloudMonitoringPrepare

> []CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(ctx, organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest(createOrganizationInventoryOnboardingCloudMonitoringPrepareRequest).Execute()

Initiates or updates an import session



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
	createOrganizationInventoryOnboardingCloudMonitoringPrepareRequest := *openapiclient.NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest([]openapiclient.CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner{*openapiclient.NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner("Sudi_example")}) // CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrepareAPI.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(context.Background(), organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest(createOrganizationInventoryOnboardingCloudMonitoringPrepareRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrepareAPI.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationInventoryOnboardingCloudMonitoringPrepare`: []CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PrepareAPI.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInventoryOnboardingCloudMonitoringPrepareRequest** | [**CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest**](CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest.md) |  | 

### Return type

[**[]CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner**](CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

