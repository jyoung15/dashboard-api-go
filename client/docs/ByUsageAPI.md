# \ByUsageAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationSummaryTopApplicationsByUsage**](ByUsageAPI.md#GetOrganizationSummaryTopApplicationsByUsage) | **Get** /organizations/{organizationId}/summary/top/applications/byUsage | Return the top applications sorted by data usage over given time range
[**GetOrganizationSummaryTopApplicationsCategoriesByUsage**](ByUsageAPI.md#GetOrganizationSummaryTopApplicationsCategoriesByUsage) | **Get** /organizations/{organizationId}/summary/top/applications/categories/byUsage | Return the top application categories sorted by data usage over given time range
[**GetOrganizationSummaryTopClientsByUsage**](ByUsageAPI.md#GetOrganizationSummaryTopClientsByUsage) | **Get** /organizations/{organizationId}/summary/top/clients/byUsage | Return metrics for organization&#39;s top 10 clients by data usage (in mb) over given time range.
[**GetOrganizationSummaryTopClientsManufacturersByUsage**](ByUsageAPI.md#GetOrganizationSummaryTopClientsManufacturersByUsage) | **Get** /organizations/{organizationId}/summary/top/clients/manufacturers/byUsage | Return metrics for organization&#39;s top clients by data usage (in mb) over given time range, grouped by manufacturer.
[**GetOrganizationSummaryTopDevicesByUsage**](ByUsageAPI.md#GetOrganizationSummaryTopDevicesByUsage) | **Get** /organizations/{organizationId}/summary/top/devices/byUsage | Return metrics for organization&#39;s top 10 devices sorted by data usage over given time range
[**GetOrganizationSummaryTopDevicesModelsByUsage**](ByUsageAPI.md#GetOrganizationSummaryTopDevicesModelsByUsage) | **Get** /organizations/{organizationId}/summary/top/devices/models/byUsage | Return metrics for organization&#39;s top 10 device models sorted by data usage over given time range
[**GetOrganizationSummaryTopSsidsByUsage**](ByUsageAPI.md#GetOrganizationSummaryTopSsidsByUsage) | **Get** /organizations/{organizationId}/summary/top/ssids/byUsage | Return metrics for organization&#39;s top 10 ssids by data usage over given time range



## GetOrganizationSummaryTopApplicationsByUsage

> []GetOrganizationSummaryTopApplicationsByUsage200ResponseInner GetOrganizationSummaryTopApplicationsByUsage(ctx, organizationId).NetworkTag(networkTag).Device(device).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()

Return the top applications sorted by data usage over given time range



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
	networkTag := "networkTag_example" // string | Match result to an exact network tag (optional)
	device := "device_example" // string | Match result to an exact device tag (optional)
	networkId := "networkId_example" // string | Match result to an exact network id (optional)
	quantity := int32(56) // int32 | Set number of desired results to return. Default is 10. (optional)
	ssidName := "ssidName_example" // string | Filter results by ssid name (optional)
	usageUplink := "usageUplink_example" // string | Filter results by usage uplink (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 25 minutes and be less than or equal to 186 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ByUsageAPI.GetOrganizationSummaryTopApplicationsByUsage(context.Background(), organizationId).NetworkTag(networkTag).Device(device).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ByUsageAPI.GetOrganizationSummaryTopApplicationsByUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSummaryTopApplicationsByUsage`: []GetOrganizationSummaryTopApplicationsByUsage200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ByUsageAPI.GetOrganizationSummaryTopApplicationsByUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopApplicationsByUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkTag** | **string** | Match result to an exact network tag | 
 **device** | **string** | Match result to an exact device tag | 
 **networkId** | **string** | Match result to an exact network id | 
 **quantity** | **int32** | Set number of desired results to return. Default is 10. | 
 **ssidName** | **string** | Filter results by ssid name | 
 **usageUplink** | **string** | Filter results by usage uplink | 
 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 25 minutes and be less than or equal to 186 days. The default is 1 day. | 

### Return type

[**[]GetOrganizationSummaryTopApplicationsByUsage200ResponseInner**](GetOrganizationSummaryTopApplicationsByUsage200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummaryTopApplicationsCategoriesByUsage

> []GetOrganizationSummaryTopApplicationsCategoriesByUsage200ResponseInner GetOrganizationSummaryTopApplicationsCategoriesByUsage(ctx, organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()

Return the top application categories sorted by data usage over given time range



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
	networkTag := "networkTag_example" // string | Match result to an exact network tag (optional)
	deviceTag := "deviceTag_example" // string | Match result to an exact device tag (optional)
	networkId := "networkId_example" // string | Match result to an exact network id (optional)
	quantity := int32(56) // int32 | Set number of desired results to return. Default is 10. (optional)
	ssidName := "ssidName_example" // string | Filter results by ssid name (optional)
	usageUplink := "usageUplink_example" // string | Filter results by usage uplink (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 25 minutes and be less than or equal to 186 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ByUsageAPI.GetOrganizationSummaryTopApplicationsCategoriesByUsage(context.Background(), organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ByUsageAPI.GetOrganizationSummaryTopApplicationsCategoriesByUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSummaryTopApplicationsCategoriesByUsage`: []GetOrganizationSummaryTopApplicationsCategoriesByUsage200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ByUsageAPI.GetOrganizationSummaryTopApplicationsCategoriesByUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopApplicationsCategoriesByUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkTag** | **string** | Match result to an exact network tag | 
 **deviceTag** | **string** | Match result to an exact device tag | 
 **networkId** | **string** | Match result to an exact network id | 
 **quantity** | **int32** | Set number of desired results to return. Default is 10. | 
 **ssidName** | **string** | Filter results by ssid name | 
 **usageUplink** | **string** | Filter results by usage uplink | 
 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 25 minutes and be less than or equal to 186 days. The default is 1 day. | 

### Return type

[**[]GetOrganizationSummaryTopApplicationsCategoriesByUsage200ResponseInner**](GetOrganizationSummaryTopApplicationsCategoriesByUsage200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummaryTopClientsByUsage

> []GetOrganizationSummaryTopClientsByUsage200ResponseInner GetOrganizationSummaryTopClientsByUsage(ctx, organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()

Return metrics for organization's top 10 clients by data usage (in mb) over given time range.



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
	networkTag := "networkTag_example" // string | Match result to an exact network tag (optional)
	deviceTag := "deviceTag_example" // string | Match result to an exact device tag (optional)
	networkId := "networkId_example" // string | Match result to an exact network id (optional)
	quantity := int32(56) // int32 | Set number of desired results to return. Default is 10. (optional)
	ssidName := "ssidName_example" // string | Filter results by ssid name (optional)
	usageUplink := "usageUplink_example" // string | Filter results by usage uplink (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 8 hours and be less than or equal to 186 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ByUsageAPI.GetOrganizationSummaryTopClientsByUsage(context.Background(), organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ByUsageAPI.GetOrganizationSummaryTopClientsByUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSummaryTopClientsByUsage`: []GetOrganizationSummaryTopClientsByUsage200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ByUsageAPI.GetOrganizationSummaryTopClientsByUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopClientsByUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkTag** | **string** | Match result to an exact network tag | 
 **deviceTag** | **string** | Match result to an exact device tag | 
 **networkId** | **string** | Match result to an exact network id | 
 **quantity** | **int32** | Set number of desired results to return. Default is 10. | 
 **ssidName** | **string** | Filter results by ssid name | 
 **usageUplink** | **string** | Filter results by usage uplink | 
 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 8 hours and be less than or equal to 186 days. The default is 1 day. | 

### Return type

[**[]GetOrganizationSummaryTopClientsByUsage200ResponseInner**](GetOrganizationSummaryTopClientsByUsage200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummaryTopClientsManufacturersByUsage

> []GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner GetOrganizationSummaryTopClientsManufacturersByUsage(ctx, organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()

Return metrics for organization's top clients by data usage (in mb) over given time range, grouped by manufacturer.



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
	networkTag := "networkTag_example" // string | Match result to an exact network tag (optional)
	deviceTag := "deviceTag_example" // string | Match result to an exact device tag (optional)
	networkId := "networkId_example" // string | Match result to an exact network id (optional)
	quantity := int32(56) // int32 | Set number of desired results to return. Default is 10. (optional)
	ssidName := "ssidName_example" // string | Filter results by ssid name (optional)
	usageUplink := "usageUplink_example" // string | Filter results by usage uplink (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 186 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ByUsageAPI.GetOrganizationSummaryTopClientsManufacturersByUsage(context.Background(), organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ByUsageAPI.GetOrganizationSummaryTopClientsManufacturersByUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSummaryTopClientsManufacturersByUsage`: []GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ByUsageAPI.GetOrganizationSummaryTopClientsManufacturersByUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopClientsManufacturersByUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkTag** | **string** | Match result to an exact network tag | 
 **deviceTag** | **string** | Match result to an exact device tag | 
 **networkId** | **string** | Match result to an exact network id | 
 **quantity** | **int32** | Set number of desired results to return. Default is 10. | 
 **ssidName** | **string** | Filter results by ssid name | 
 **usageUplink** | **string** | Filter results by usage uplink | 
 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 186 days. The default is 1 day. | 

### Return type

[**[]GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner**](GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummaryTopDevicesByUsage

> []GetOrganizationSummaryTopDevicesByUsage200ResponseInner GetOrganizationSummaryTopDevicesByUsage(ctx, organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()

Return metrics for organization's top 10 devices sorted by data usage over given time range



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
	networkTag := "networkTag_example" // string | Match result to an exact network tag (optional)
	deviceTag := "deviceTag_example" // string | Match result to an exact device tag (optional)
	networkId := "networkId_example" // string | Match result to an exact network id (optional)
	quantity := int32(56) // int32 | Set number of desired results to return. Default is 10. (optional)
	ssidName := "ssidName_example" // string | Filter results by ssid name (optional)
	usageUplink := "usageUplink_example" // string | Filter results by usage uplink (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 8 hours and be less than or equal to 186 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ByUsageAPI.GetOrganizationSummaryTopDevicesByUsage(context.Background(), organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ByUsageAPI.GetOrganizationSummaryTopDevicesByUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSummaryTopDevicesByUsage`: []GetOrganizationSummaryTopDevicesByUsage200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ByUsageAPI.GetOrganizationSummaryTopDevicesByUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopDevicesByUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkTag** | **string** | Match result to an exact network tag | 
 **deviceTag** | **string** | Match result to an exact device tag | 
 **networkId** | **string** | Match result to an exact network id | 
 **quantity** | **int32** | Set number of desired results to return. Default is 10. | 
 **ssidName** | **string** | Filter results by ssid name | 
 **usageUplink** | **string** | Filter results by usage uplink | 
 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 8 hours and be less than or equal to 186 days. The default is 1 day. | 

### Return type

[**[]GetOrganizationSummaryTopDevicesByUsage200ResponseInner**](GetOrganizationSummaryTopDevicesByUsage200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummaryTopDevicesModelsByUsage

> []GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner GetOrganizationSummaryTopDevicesModelsByUsage(ctx, organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()

Return metrics for organization's top 10 device models sorted by data usage over given time range



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
	networkTag := "networkTag_example" // string | Match result to an exact network tag (optional)
	deviceTag := "deviceTag_example" // string | Match result to an exact device tag (optional)
	networkId := "networkId_example" // string | Match result to an exact network id (optional)
	quantity := int32(56) // int32 | Set number of desired results to return. Default is 10. (optional)
	ssidName := "ssidName_example" // string | Filter results by ssid name (optional)
	usageUplink := "usageUplink_example" // string | Filter results by usage uplink (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 8 hours and be less than or equal to 186 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ByUsageAPI.GetOrganizationSummaryTopDevicesModelsByUsage(context.Background(), organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ByUsageAPI.GetOrganizationSummaryTopDevicesModelsByUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSummaryTopDevicesModelsByUsage`: []GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ByUsageAPI.GetOrganizationSummaryTopDevicesModelsByUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopDevicesModelsByUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkTag** | **string** | Match result to an exact network tag | 
 **deviceTag** | **string** | Match result to an exact device tag | 
 **networkId** | **string** | Match result to an exact network id | 
 **quantity** | **int32** | Set number of desired results to return. Default is 10. | 
 **ssidName** | **string** | Filter results by ssid name | 
 **usageUplink** | **string** | Filter results by usage uplink | 
 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 8 hours and be less than or equal to 186 days. The default is 1 day. | 

### Return type

[**[]GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner**](GetOrganizationSummaryTopDevicesModelsByUsage200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSummaryTopSsidsByUsage

> []GetOrganizationSummaryTopSsidsByUsage200ResponseInner GetOrganizationSummaryTopSsidsByUsage(ctx, organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()

Return metrics for organization's top 10 ssids by data usage over given time range



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
	networkTag := "networkTag_example" // string | Match result to an exact network tag (optional)
	deviceTag := "deviceTag_example" // string | Match result to an exact device tag (optional)
	networkId := "networkId_example" // string | Match result to an exact network id (optional)
	quantity := int32(56) // int32 | Set number of desired results to return. Default is 10. (optional)
	ssidName := "ssidName_example" // string | Filter results by ssid name (optional)
	usageUplink := "usageUplink_example" // string | Filter results by usage uplink (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 8 hours and be less than or equal to 186 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ByUsageAPI.GetOrganizationSummaryTopSsidsByUsage(context.Background(), organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ByUsageAPI.GetOrganizationSummaryTopSsidsByUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSummaryTopSsidsByUsage`: []GetOrganizationSummaryTopSsidsByUsage200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ByUsageAPI.GetOrganizationSummaryTopSsidsByUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopSsidsByUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkTag** | **string** | Match result to an exact network tag | 
 **deviceTag** | **string** | Match result to an exact device tag | 
 **networkId** | **string** | Match result to an exact network id | 
 **quantity** | **int32** | Set number of desired results to return. Default is 10. | 
 **ssidName** | **string** | Filter results by ssid name | 
 **usageUplink** | **string** | Filter results by usage uplink | 
 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 186 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 8 hours and be less than or equal to 186 days. The default is 1 day. | 

### Return type

[**[]GetOrganizationSummaryTopSsidsByUsage200ResponseInner**](GetOrganizationSummaryTopSsidsByUsage200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

