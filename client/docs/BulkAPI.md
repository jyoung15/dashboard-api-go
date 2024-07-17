# \BulkAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationInventoryDevicesSwapsBulk**](BulkAPI.md#CreateOrganizationInventoryDevicesSwapsBulk) | **Post** /organizations/{organizationId}/inventory/devices/swaps/bulk | Swap the devices identified by devices.old with a devices.new, then perform the :afterAction on the devices.old.
[**GetOrganizationInventoryDevicesSwapsBulk**](BulkAPI.md#GetOrganizationInventoryDevicesSwapsBulk) | **Get** /organizations/{organizationId}/inventory/devices/swaps/bulk/{id} | List of device swaps for a given request ID ({id}).



## CreateOrganizationInventoryDevicesSwapsBulk

> CreateOrganizationInventoryDevicesSwapsBulk207Response CreateOrganizationInventoryDevicesSwapsBulk(ctx, organizationId).CreateOrganizationInventoryDevicesSwapsBulkRequest(createOrganizationInventoryDevicesSwapsBulkRequest).Execute()

Swap the devices identified by devices.old with a devices.new, then perform the :afterAction on the devices.old.



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
	createOrganizationInventoryDevicesSwapsBulkRequest := *openapiclient.NewCreateOrganizationInventoryDevicesSwapsBulkRequest([]openapiclient.CreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInner{*openapiclient.NewCreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInner(*openapiclient.NewCreateOrganizationInventoryDevicesSwapsBulkRequestSwapsInnerDevices("Old_example", "New_example"), "AfterAction_example")}) // CreateOrganizationInventoryDevicesSwapsBulkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkAPI.CreateOrganizationInventoryDevicesSwapsBulk(context.Background(), organizationId).CreateOrganizationInventoryDevicesSwapsBulkRequest(createOrganizationInventoryDevicesSwapsBulkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkAPI.CreateOrganizationInventoryDevicesSwapsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationInventoryDevicesSwapsBulk`: CreateOrganizationInventoryDevicesSwapsBulk207Response
	fmt.Fprintf(os.Stdout, "Response from `BulkAPI.CreateOrganizationInventoryDevicesSwapsBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInventoryDevicesSwapsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInventoryDevicesSwapsBulkRequest** | [**CreateOrganizationInventoryDevicesSwapsBulkRequest**](CreateOrganizationInventoryDevicesSwapsBulkRequest.md) |  | 

### Return type

[**CreateOrganizationInventoryDevicesSwapsBulk207Response**](CreateOrganizationInventoryDevicesSwapsBulk207Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryDevicesSwapsBulk

> CreateOrganizationInventoryDevicesSwapsBulk207Response GetOrganizationInventoryDevicesSwapsBulk(ctx, organizationId, id).Execute()

List of device swaps for a given request ID ({id}).



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
	id := "id_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkAPI.GetOrganizationInventoryDevicesSwapsBulk(context.Background(), organizationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkAPI.GetOrganizationInventoryDevicesSwapsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationInventoryDevicesSwapsBulk`: CreateOrganizationInventoryDevicesSwapsBulk207Response
	fmt.Fprintf(os.Stdout, "Response from `BulkAPI.GetOrganizationInventoryDevicesSwapsBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryDevicesSwapsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateOrganizationInventoryDevicesSwapsBulk207Response**](CreateOrganizationInventoryDevicesSwapsBulk207Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

