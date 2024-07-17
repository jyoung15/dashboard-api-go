# \InventoryAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClaimIntoOrganizationInventory**](InventoryAPI.md#ClaimIntoOrganizationInventory) | **Post** /organizations/{organizationId}/inventory/claim | Claim a list of devices, licenses, and/or orders into an organization inventory
[**CreateOrganizationInventoryDevicesSwapsBulk**](InventoryAPI.md#CreateOrganizationInventoryDevicesSwapsBulk) | **Post** /organizations/{organizationId}/inventory/devices/swaps/bulk | Swap the devices identified by devices.old with a devices.new, then perform the :afterAction on the devices.old.
[**CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent**](InventoryAPI.md#CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent) | **Post** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/exportEvents | Imports event logs related to the onboarding app into elastisearch
[**CreateOrganizationInventoryOnboardingCloudMonitoringImport**](InventoryAPI.md#CreateOrganizationInventoryOnboardingCloudMonitoringImport) | **Post** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports | Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.
[**CreateOrganizationInventoryOnboardingCloudMonitoringPrepare**](InventoryAPI.md#CreateOrganizationInventoryOnboardingCloudMonitoringPrepare) | **Post** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/prepare | Initiates or updates an import session
[**GetOrganizationInventoryDevice**](InventoryAPI.md#GetOrganizationInventoryDevice) | **Get** /organizations/{organizationId}/inventory/devices/{serial} | Return a single device from the inventory of an organization
[**GetOrganizationInventoryDevices**](InventoryAPI.md#GetOrganizationInventoryDevices) | **Get** /organizations/{organizationId}/inventory/devices | Return the device inventory for an organization
[**GetOrganizationInventoryDevicesSwapsBulk**](InventoryAPI.md#GetOrganizationInventoryDevicesSwapsBulk) | **Get** /organizations/{organizationId}/inventory/devices/swaps/bulk/{id} | List of device swaps for a given request ID ({id}).
[**GetOrganizationInventoryOnboardingCloudMonitoringImports**](InventoryAPI.md#GetOrganizationInventoryOnboardingCloudMonitoringImports) | **Get** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports | Check the status of a committed Import operation
[**GetOrganizationInventoryOnboardingCloudMonitoringNetworks**](InventoryAPI.md#GetOrganizationInventoryOnboardingCloudMonitoringNetworks) | **Get** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/networks | Returns list of networks eligible for adding cloud monitored device
[**ReleaseFromOrganizationInventory**](InventoryAPI.md#ReleaseFromOrganizationInventory) | **Post** /organizations/{organizationId}/inventory/release | Release a list of claimed devices from an organization.



## ClaimIntoOrganizationInventory

> ClaimIntoOrganization200Response ClaimIntoOrganizationInventory(ctx, organizationId).ClaimIntoOrganizationInventoryRequest(claimIntoOrganizationInventoryRequest).Execute()

Claim a list of devices, licenses, and/or orders into an organization inventory



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
	claimIntoOrganizationInventoryRequest := *openapiclient.NewClaimIntoOrganizationInventoryRequest() // ClaimIntoOrganizationInventoryRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.ClaimIntoOrganizationInventory(context.Background(), organizationId).ClaimIntoOrganizationInventoryRequest(claimIntoOrganizationInventoryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.ClaimIntoOrganizationInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClaimIntoOrganizationInventory`: ClaimIntoOrganization200Response
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.ClaimIntoOrganizationInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimIntoOrganizationInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **claimIntoOrganizationInventoryRequest** | [**ClaimIntoOrganizationInventoryRequest**](ClaimIntoOrganizationInventoryRequest.md) |  | 

### Return type

[**ClaimIntoOrganization200Response**](ClaimIntoOrganization200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.InventoryAPI.CreateOrganizationInventoryDevicesSwapsBulk(context.Background(), organizationId).CreateOrganizationInventoryDevicesSwapsBulkRequest(createOrganizationInventoryDevicesSwapsBulkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.CreateOrganizationInventoryDevicesSwapsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationInventoryDevicesSwapsBulk`: CreateOrganizationInventoryDevicesSwapsBulk207Response
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.CreateOrganizationInventoryDevicesSwapsBulk`: %v\n", resp)
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


## CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent

> map[string]interface{} CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent(ctx, organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest(createOrganizationInventoryOnboardingCloudMonitoringExportEventRequest).Execute()

Imports event logs related to the onboarding app into elastisearch



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
	createOrganizationInventoryOnboardingCloudMonitoringExportEventRequest := *openapiclient.NewCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest("LogEvent_example", int32(123)) // CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent(context.Background(), organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest(createOrganizationInventoryOnboardingCloudMonitoringExportEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInventoryOnboardingCloudMonitoringExportEventRequest** | [**CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest**](CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest.md) |  | 

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


## CreateOrganizationInventoryOnboardingCloudMonitoringImport

> []CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner CreateOrganizationInventoryOnboardingCloudMonitoringImport(ctx, organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest(createOrganizationInventoryOnboardingCloudMonitoringImportRequest).Execute()

Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.



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
	createOrganizationInventoryOnboardingCloudMonitoringImportRequest := *openapiclient.NewCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest([]openapiclient.CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner{*openapiclient.NewCreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner("DeviceId_example", "Udi_example", "NetworkId_example")}) // CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.CreateOrganizationInventoryOnboardingCloudMonitoringImport(context.Background(), organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest(createOrganizationInventoryOnboardingCloudMonitoringImportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.CreateOrganizationInventoryOnboardingCloudMonitoringImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationInventoryOnboardingCloudMonitoringImport`: []CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.CreateOrganizationInventoryOnboardingCloudMonitoringImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInventoryOnboardingCloudMonitoringImportRequest** | [**CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest**](CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest.md) |  | 

### Return type

[**[]CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner**](CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.InventoryAPI.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(context.Background(), organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest(createOrganizationInventoryOnboardingCloudMonitoringPrepareRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationInventoryOnboardingCloudMonitoringPrepare`: []CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare`: %v\n", resp)
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


## GetOrganizationInventoryDevice

> GetOrganizationInventoryDevices200ResponseInner GetOrganizationInventoryDevice(ctx, organizationId, serial).Execute()

Return a single device from the inventory of an organization



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
	serial := "serial_example" // string | Serial

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.GetOrganizationInventoryDevice(context.Background(), organizationId, serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.GetOrganizationInventoryDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationInventoryDevice`: GetOrganizationInventoryDevices200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.GetOrganizationInventoryDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationInventoryDevices200ResponseInner**](GetOrganizationInventoryDevices200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryDevices

> []GetOrganizationInventoryDevices200ResponseInner GetOrganizationInventoryDevices(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).UsedState(usedState).Search(search).Macs(macs).NetworkIds(networkIds).Serials(serials).Models(models).OrderNumbers(orderNumbers).Tags(tags).TagsFilterType(tagsFilterType).ProductTypes(productTypes).Execute()

Return the device inventory for an organization



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	usedState := "usedState_example" // string | Filter results by used or unused inventory. Accepted values are 'used' or 'unused'. (optional)
	search := "search_example" // string | Search for devices in inventory based on serial number, mac address, or model. (optional)
	macs := []string{"Inner_example"} // []string | Search for devices in inventory based on mac addresses. (optional)
	networkIds := []string{"Inner_example"} // []string | Search for devices in inventory based on network ids. Use explicit 'null' value to get available devices only. (optional)
	serials := []string{"Inner_example"} // []string | Search for devices in inventory based on serials. (optional)
	models := []string{"Inner_example"} // []string | Search for devices in inventory based on model. (optional)
	orderNumbers := []string{"Inner_example"} // []string | Search for devices in inventory based on order numbers. (optional)
	tags := []string{"Inner_example"} // []string | Filter devices by tags. The filtering is case-sensitive. If tags are included, 'tagsFilterType' should also be included (see below). (optional)
	tagsFilterType := "tagsFilterType_example" // string | To use with 'tags' parameter, to filter devices which contain ANY or ALL given tags. Accepted values are 'withAnyTags' or 'withAllTags', default is 'withAnyTags'. (optional)
	productTypes := []string{"ProductTypes_example"} // []string | Filter devices by product type. Accepted values are appliance, camera, cellularGateway, secureConnect, sensor, switch, systemsManager, wireless, and wirelessController. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.GetOrganizationInventoryDevices(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).UsedState(usedState).Search(search).Macs(macs).NetworkIds(networkIds).Serials(serials).Models(models).OrderNumbers(orderNumbers).Tags(tags).TagsFilterType(tagsFilterType).ProductTypes(productTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.GetOrganizationInventoryDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationInventoryDevices`: []GetOrganizationInventoryDevices200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.GetOrganizationInventoryDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **usedState** | **string** | Filter results by used or unused inventory. Accepted values are &#39;used&#39; or &#39;unused&#39;. | 
 **search** | **string** | Search for devices in inventory based on serial number, mac address, or model. | 
 **macs** | **[]string** | Search for devices in inventory based on mac addresses. | 
 **networkIds** | **[]string** | Search for devices in inventory based on network ids. Use explicit &#39;null&#39; value to get available devices only. | 
 **serials** | **[]string** | Search for devices in inventory based on serials. | 
 **models** | **[]string** | Search for devices in inventory based on model. | 
 **orderNumbers** | **[]string** | Search for devices in inventory based on order numbers. | 
 **tags** | **[]string** | Filter devices by tags. The filtering is case-sensitive. If tags are included, &#39;tagsFilterType&#39; should also be included (see below). | 
 **tagsFilterType** | **string** | To use with &#39;tags&#39; parameter, to filter devices which contain ANY or ALL given tags. Accepted values are &#39;withAnyTags&#39; or &#39;withAllTags&#39;, default is &#39;withAnyTags&#39;. | 
 **productTypes** | **[]string** | Filter devices by product type. Accepted values are appliance, camera, cellularGateway, secureConnect, sensor, switch, systemsManager, wireless, and wirelessController. | 

### Return type

[**[]GetOrganizationInventoryDevices200ResponseInner**](GetOrganizationInventoryDevices200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
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
	resp, r, err := apiClient.InventoryAPI.GetOrganizationInventoryDevicesSwapsBulk(context.Background(), organizationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.GetOrganizationInventoryDevicesSwapsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationInventoryDevicesSwapsBulk`: CreateOrganizationInventoryDevicesSwapsBulk207Response
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.GetOrganizationInventoryDevicesSwapsBulk`: %v\n", resp)
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


## GetOrganizationInventoryOnboardingCloudMonitoringImports

> []GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner GetOrganizationInventoryOnboardingCloudMonitoringImports(ctx, organizationId).ImportIds(importIds).Execute()

Check the status of a committed Import operation



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
	importIds := []string{"Inner_example"} // []string | import ids from an imports

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.GetOrganizationInventoryOnboardingCloudMonitoringImports(context.Background(), organizationId).ImportIds(importIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.GetOrganizationInventoryOnboardingCloudMonitoringImports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationInventoryOnboardingCloudMonitoringImports`: []GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.GetOrganizationInventoryOnboardingCloudMonitoringImports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importIds** | **[]string** | import ids from an imports | 

### Return type

[**[]GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner**](GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryOnboardingCloudMonitoringNetworks

> []GetNetwork200Response GetOrganizationInventoryOnboardingCloudMonitoringNetworks(ctx, organizationId).DeviceType(deviceType).Search(search).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Returns list of networks eligible for adding cloud monitored device



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
	deviceType := "deviceType_example" // string | Device Type switch or wireless controller
	search := "search_example" // string | Optional parameter to search on network name (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.GetOrganizationInventoryOnboardingCloudMonitoringNetworks(context.Background(), organizationId).DeviceType(deviceType).Search(search).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.GetOrganizationInventoryOnboardingCloudMonitoringNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationInventoryOnboardingCloudMonitoringNetworks`: []GetNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.GetOrganizationInventoryOnboardingCloudMonitoringNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceType** | **string** | Device Type switch or wireless controller | 
 **search** | **string** | Optional parameter to search on network name | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetwork200Response**](GetNetwork200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseFromOrganizationInventory

> ReleaseFromOrganizationInventory200Response ReleaseFromOrganizationInventory(ctx, organizationId).ReleaseFromOrganizationInventoryRequest(releaseFromOrganizationInventoryRequest).Execute()

Release a list of claimed devices from an organization.



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
	releaseFromOrganizationInventoryRequest := *openapiclient.NewReleaseFromOrganizationInventoryRequest() // ReleaseFromOrganizationInventoryRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryAPI.ReleaseFromOrganizationInventory(context.Background(), organizationId).ReleaseFromOrganizationInventoryRequest(releaseFromOrganizationInventoryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryAPI.ReleaseFromOrganizationInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReleaseFromOrganizationInventory`: ReleaseFromOrganizationInventory200Response
	fmt.Fprintf(os.Stdout, "Response from `InventoryAPI.ReleaseFromOrganizationInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseFromOrganizationInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **releaseFromOrganizationInventoryRequest** | [**ReleaseFromOrganizationInventoryRequest**](ReleaseFromOrganizationInventoryRequest.md) |  | 

### Return type

[**ReleaseFromOrganizationInventory200Response**](ReleaseFromOrganizationInventory200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

