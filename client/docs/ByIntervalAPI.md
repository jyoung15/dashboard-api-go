# \ByIntervalAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationApiRequestsOverviewResponseCodesByInterval**](ByIntervalAPI.md#GetOrganizationApiRequestsOverviewResponseCodesByInterval) | **Get** /organizations/{organizationId}/apiRequests/overview/responseCodes/byInterval | Tracks organizations&#39; API requests by response code across a given time period
[**GetOrganizationCameraDetectionsHistoryByBoundaryByInterval**](ByIntervalAPI.md#GetOrganizationCameraDetectionsHistoryByBoundaryByInterval) | **Get** /organizations/{organizationId}/camera/detections/history/byBoundary/byInterval | Returns analytics data for timespans
[**GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval**](ByIntervalAPI.md#GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval) | **Get** /organizations/{organizationId}/wireless/devices/channelUtilization/history/byDevice/byInterval | Get a time-series of average channel utilization for all bands, segmented by device.
[**GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval**](ByIntervalAPI.md#GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval) | **Get** /organizations/{organizationId}/wireless/devices/channelUtilization/history/byNetwork/byInterval | Get a time-series of average channel utilization for all bands



## GetOrganizationApiRequestsOverviewResponseCodesByInterval

> []GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner GetOrganizationApiRequestsOverviewResponseCodesByInterval(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Version(version).OperationIds(operationIds).SourceIps(sourceIps).AdminIds(adminIds).UserAgent(userAgent).Execute()

Tracks organizations' API requests by response code across a given time period



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
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days. If interval is provided, the timespan will be autocalculated. (optional)
	interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 120, 3600, 14400, 21600. The default is 21600. Interval is calculated if time params are provided. (optional)
	version := int32(56) // int32 | Filter by API version of the endpoint. Allowable values are: [0, 1] (optional)
	operationIds := []string{"Inner_example"} // []string | Filter by operation ID of the endpoint (optional)
	sourceIps := []string{"Inner_example"} // []string | Filter by source IP that made the API request (optional)
	adminIds := []string{"Inner_example"} // []string | Filter by admin ID of user that made the API request (optional)
	userAgent := "userAgent_example" // string | Filter by user agent string for API request. This will filter by a complete or partial match. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ByIntervalAPI.GetOrganizationApiRequestsOverviewResponseCodesByInterval(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Version(version).OperationIds(operationIds).SourceIps(sourceIps).AdminIds(adminIds).UserAgent(userAgent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ByIntervalAPI.GetOrganizationApiRequestsOverviewResponseCodesByInterval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationApiRequestsOverviewResponseCodesByInterval`: []GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ByIntervalAPI.GetOrganizationApiRequestsOverviewResponseCodesByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days. If interval is provided, the timespan will be autocalculated. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 120, 3600, 14400, 21600. The default is 21600. Interval is calculated if time params are provided. | 
 **version** | **int32** | Filter by API version of the endpoint. Allowable values are: [0, 1] | 
 **operationIds** | **[]string** | Filter by operation ID of the endpoint | 
 **sourceIps** | **[]string** | Filter by source IP that made the API request | 
 **adminIds** | **[]string** | Filter by admin ID of user that made the API request | 
 **userAgent** | **string** | Filter by user agent string for API request. This will filter by a complete or partial match. | 

### Return type

[**[]GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner**](GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCameraDetectionsHistoryByBoundaryByInterval

> []GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInner GetOrganizationCameraDetectionsHistoryByBoundaryByInterval(ctx, organizationId).BoundaryIds(boundaryIds).Ranges(ranges).Duration(duration).PerPage(perPage).BoundaryTypes(boundaryTypes).Execute()

Returns analytics data for timespans



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
	boundaryIds := []string{"Inner_example"} // []string | A list of boundary ids. The returned cameras will be filtered to only include these ids.
	ranges := []openapiclient.GetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRangesParameterInner{*openapiclient.NewGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRangesParameterInner("StartTime_example", "EndTime_example", int32(123))} // []GetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRangesParameterInner | A list of time ranges with intervals
	duration := int32(56) // int32 | The minimum time, in seconds, that the person or car remains in the area to be counted. Defaults to boundary configuration or 60. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 1 - 1000. Defaults to 1000. (optional)
	boundaryTypes := []string{"BoundaryTypes_example"} // []string | The detection types. Defaults to 'person'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ByIntervalAPI.GetOrganizationCameraDetectionsHistoryByBoundaryByInterval(context.Background(), organizationId).BoundaryIds(boundaryIds).Ranges(ranges).Duration(duration).PerPage(perPage).BoundaryTypes(boundaryTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ByIntervalAPI.GetOrganizationCameraDetectionsHistoryByBoundaryByInterval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationCameraDetectionsHistoryByBoundaryByInterval`: []GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ByIntervalAPI.GetOrganizationCameraDetectionsHistoryByBoundaryByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **boundaryIds** | **[]string** | A list of boundary ids. The returned cameras will be filtered to only include these ids. | 
 **ranges** | [**[]GetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRangesParameterInner**](GetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRangesParameterInner.md) | A list of time ranges with intervals | 
 **duration** | **int32** | The minimum time, in seconds, that the person or car remains in the area to be counted. Defaults to boundary configuration or 60. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 1 - 1000. Defaults to 1000. | 
 **boundaryTypes** | **[]string** | The detection types. Defaults to &#39;person&#39;. | 

### Return type

[**[]GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInner**](GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval

> []GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval(ctx, organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()

Get a time-series of average channel utilization for all bands, segmented by device.



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
	networkIds := []string{"Inner_example"} // []string | Filter results by network. (optional)
	serials := []string{"Inner_example"} // []string | Filter results by device. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
	interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ByIntervalAPI.GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ByIntervalAPI.GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval`: []GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ByIntervalAPI.GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter results by network. | 
 **serials** | **[]string** | Filter results by device. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. | 

### Return type

[**[]GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner**](GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval

> []GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval(ctx, organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()

Get a time-series of average channel utilization for all bands



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
	networkIds := []string{"Inner_example"} // []string | Filter results by network. (optional)
	serials := []string{"Inner_example"} // []string | Filter results by device. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
	interval := int32(56) // int32 | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ByIntervalAPI.GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval(context.Background(), organizationId).NetworkIds(networkIds).Serials(serials).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ByIntervalAPI.GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval`: []GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ByIntervalAPI.GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkIds** | **[]string** | Filter results by network. | 
 **serials** | **[]string** | Filter results by device. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **interval** | **int32** | The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600. | 

### Return type

[**[]GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner**](GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

