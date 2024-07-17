# \AnalyticsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceCameraAnalyticsLive**](AnalyticsAPI.md#GetDeviceCameraAnalyticsLive) | **Get** /devices/{serial}/camera/analytics/live | Returns live state from camera analytics zones
[**GetDeviceCameraAnalyticsOverview**](AnalyticsAPI.md#GetDeviceCameraAnalyticsOverview) | **Get** /devices/{serial}/camera/analytics/overview | Returns an overview of aggregate analytics data for a timespan
[**GetDeviceCameraAnalyticsRecent**](AnalyticsAPI.md#GetDeviceCameraAnalyticsRecent) | **Get** /devices/{serial}/camera/analytics/recent | Returns most recent record for analytics zones
[**GetDeviceCameraAnalyticsZoneHistory**](AnalyticsAPI.md#GetDeviceCameraAnalyticsZoneHistory) | **Get** /devices/{serial}/camera/analytics/zones/{zoneId}/history | Return historical records for analytic zones
[**GetDeviceCameraAnalyticsZones**](AnalyticsAPI.md#GetDeviceCameraAnalyticsZones) | **Get** /devices/{serial}/camera/analytics/zones | Returns all configured analytic zones for this camera



## GetDeviceCameraAnalyticsLive

> GetDeviceCameraAnalyticsLive200Response GetDeviceCameraAnalyticsLive(ctx, serial).Execute()

Returns live state from camera analytics zones



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
	serial := "serial_example" // string | Serial

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetDeviceCameraAnalyticsLive(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetDeviceCameraAnalyticsLive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraAnalyticsLive`: GetDeviceCameraAnalyticsLive200Response
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetDeviceCameraAnalyticsLive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraAnalyticsLiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceCameraAnalyticsLive200Response**](GetDeviceCameraAnalyticsLive200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceCameraAnalyticsOverview

> []GetDeviceCameraAnalyticsOverview200ResponseInner GetDeviceCameraAnalyticsOverview(ctx, serial).T0(t0).T1(t1).Timespan(timespan).ObjectType(objectType).Execute()

Returns an overview of aggregate analytics data for a timespan



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
	serial := "serial_example" // string | Serial
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 1 hour. (optional)
	objectType := "objectType_example" // string | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetDeviceCameraAnalyticsOverview(context.Background(), serial).T0(t0).T1(t1).Timespan(timespan).ObjectType(objectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetDeviceCameraAnalyticsOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraAnalyticsOverview`: []GetDeviceCameraAnalyticsOverview200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetDeviceCameraAnalyticsOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraAnalyticsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 1 hour. | 
 **objectType** | **string** | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. | 

### Return type

[**[]GetDeviceCameraAnalyticsOverview200ResponseInner**](GetDeviceCameraAnalyticsOverview200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceCameraAnalyticsRecent

> []GetDeviceCameraAnalyticsRecent200ResponseInner GetDeviceCameraAnalyticsRecent(ctx, serial).ObjectType(objectType).Execute()

Returns most recent record for analytics zones



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
	serial := "serial_example" // string | Serial
	objectType := "objectType_example" // string | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetDeviceCameraAnalyticsRecent(context.Background(), serial).ObjectType(objectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetDeviceCameraAnalyticsRecent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraAnalyticsRecent`: []GetDeviceCameraAnalyticsRecent200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetDeviceCameraAnalyticsRecent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraAnalyticsRecentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectType** | **string** | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. | 

### Return type

[**[]GetDeviceCameraAnalyticsRecent200ResponseInner**](GetDeviceCameraAnalyticsRecent200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceCameraAnalyticsZoneHistory

> []GetDeviceCameraAnalyticsZoneHistory200ResponseInner GetDeviceCameraAnalyticsZoneHistory(ctx, serial, zoneId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).ObjectType(objectType).Execute()

Return historical records for analytic zones



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
	serial := "serial_example" // string | Serial
	zoneId := "zoneId_example" // string | Zone ID
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 14 hours after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 hours. The default is 1 hour. (optional)
	resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 60. The default is 60. (optional)
	objectType := "objectType_example" // string | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetDeviceCameraAnalyticsZoneHistory(context.Background(), serial, zoneId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).ObjectType(objectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetDeviceCameraAnalyticsZoneHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraAnalyticsZoneHistory`: []GetDeviceCameraAnalyticsZoneHistory200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetDeviceCameraAnalyticsZoneHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**zoneId** | **string** | Zone ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraAnalyticsZoneHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 14 hours after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 hours. The default is 1 hour. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 60. The default is 60. | 
 **objectType** | **string** | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. | 

### Return type

[**[]GetDeviceCameraAnalyticsZoneHistory200ResponseInner**](GetDeviceCameraAnalyticsZoneHistory200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceCameraAnalyticsZones

> []GetDeviceCameraAnalyticsZones200ResponseInner GetDeviceCameraAnalyticsZones(ctx, serial).Execute()

Returns all configured analytic zones for this camera



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
	serial := "serial_example" // string | Serial

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalyticsAPI.GetDeviceCameraAnalyticsZones(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.GetDeviceCameraAnalyticsZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraAnalyticsZones`: []GetDeviceCameraAnalyticsZones200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AnalyticsAPI.GetDeviceCameraAnalyticsZones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraAnalyticsZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetDeviceCameraAnalyticsZones200ResponseInner**](GetDeviceCameraAnalyticsZones200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

