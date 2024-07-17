# \PacketsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceSwitchPortsStatusesPackets**](PacketsAPI.md#GetDeviceSwitchPortsStatusesPackets) | **Get** /devices/{serial}/switch/ports/statuses/packets | Return the packet counters for all the ports of a switch



## GetDeviceSwitchPortsStatusesPackets

> []GetDeviceSwitchPortsStatusesPackets200ResponseInner GetDeviceSwitchPortsStatusesPackets(ctx, serial).T0(t0).Timespan(timespan).Execute()

Return the packet counters for all the ports of a switch



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
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 1 day from today. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 1 day. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PacketsAPI.GetDeviceSwitchPortsStatusesPackets(context.Background(), serial).T0(t0).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PacketsAPI.GetDeviceSwitchPortsStatusesPackets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSwitchPortsStatusesPackets`: []GetDeviceSwitchPortsStatusesPackets200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PacketsAPI.GetDeviceSwitchPortsStatusesPackets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchPortsStatusesPacketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 1 day from today. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 1 day. The default is 1 day. | 

### Return type

[**[]GetDeviceSwitchPortsStatusesPackets200ResponseInner**](GetDeviceSwitchPortsStatusesPackets200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

