# \LiveToolsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlinkDeviceLeds**](LiveToolsAPI.md#BlinkDeviceLeds) | **Post** /devices/{serial}/blinkLeds | Blink the LEDs on a device
[**CreateDeviceLiveToolsArpTable**](LiveToolsAPI.md#CreateDeviceLiveToolsArpTable) | **Post** /devices/{serial}/liveTools/arpTable | Enqueue a job to perform a ARP table request for the device
[**CreateDeviceLiveToolsCableTest**](LiveToolsAPI.md#CreateDeviceLiveToolsCableTest) | **Post** /devices/{serial}/liveTools/cableTest | Enqueue a job to perform a cable test for the device on the specified ports
[**CreateDeviceLiveToolsPing**](LiveToolsAPI.md#CreateDeviceLiveToolsPing) | **Post** /devices/{serial}/liveTools/ping | Enqueue a job to ping a target host from the device
[**CreateDeviceLiveToolsPingDevice**](LiveToolsAPI.md#CreateDeviceLiveToolsPingDevice) | **Post** /devices/{serial}/liveTools/pingDevice | Enqueue a job to check connectivity status to the device
[**CreateDeviceLiveToolsThroughputTest**](LiveToolsAPI.md#CreateDeviceLiveToolsThroughputTest) | **Post** /devices/{serial}/liveTools/throughputTest | Enqueue a job to test a device throughput, the test will run for 10 secs to test throughput
[**CreateDeviceLiveToolsWakeOnLan**](LiveToolsAPI.md#CreateDeviceLiveToolsWakeOnLan) | **Post** /devices/{serial}/liveTools/wakeOnLan | Enqueue a job to send a Wake-on-LAN packet from the device
[**CycleDeviceSwitchPorts**](LiveToolsAPI.md#CycleDeviceSwitchPorts) | **Post** /devices/{serial}/switch/ports/cycle | Cycle a set of switch ports
[**GetDeviceLiveToolsArpTable**](LiveToolsAPI.md#GetDeviceLiveToolsArpTable) | **Get** /devices/{serial}/liveTools/arpTable/{arpTableId} | Return an ARP table live tool job.
[**GetDeviceLiveToolsCableTest**](LiveToolsAPI.md#GetDeviceLiveToolsCableTest) | **Get** /devices/{serial}/liveTools/cableTest/{id} | Return a cable test live tool job.
[**GetDeviceLiveToolsPing**](LiveToolsAPI.md#GetDeviceLiveToolsPing) | **Get** /devices/{serial}/liveTools/ping/{id} | Return a ping job
[**GetDeviceLiveToolsPingDevice**](LiveToolsAPI.md#GetDeviceLiveToolsPingDevice) | **Get** /devices/{serial}/liveTools/pingDevice/{id} | Return a ping device job
[**GetDeviceLiveToolsThroughputTest**](LiveToolsAPI.md#GetDeviceLiveToolsThroughputTest) | **Get** /devices/{serial}/liveTools/throughputTest/{throughputTestId} | Return a throughput test job
[**GetDeviceLiveToolsWakeOnLan**](LiveToolsAPI.md#GetDeviceLiveToolsWakeOnLan) | **Get** /devices/{serial}/liveTools/wakeOnLan/{wakeOnLanId} | Return a Wake-on-LAN job
[**RebootDevice**](LiveToolsAPI.md#RebootDevice) | **Post** /devices/{serial}/reboot | Reboot a device



## BlinkDeviceLeds

> BlinkDeviceLeds202Response BlinkDeviceLeds(ctx, serial).BlinkDeviceLedsRequest(blinkDeviceLedsRequest).Execute()

Blink the LEDs on a device



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
	blinkDeviceLedsRequest := *openapiclient.NewBlinkDeviceLedsRequest() // BlinkDeviceLedsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveToolsAPI.BlinkDeviceLeds(context.Background(), serial).BlinkDeviceLedsRequest(blinkDeviceLedsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsAPI.BlinkDeviceLeds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlinkDeviceLeds`: BlinkDeviceLeds202Response
	fmt.Fprintf(os.Stdout, "Response from `LiveToolsAPI.BlinkDeviceLeds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlinkDeviceLedsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blinkDeviceLedsRequest** | [**BlinkDeviceLedsRequest**](BlinkDeviceLedsRequest.md) |  | 

### Return type

[**BlinkDeviceLeds202Response**](BlinkDeviceLeds202Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsArpTable

> CreateDeviceLiveToolsArpTable201Response CreateDeviceLiveToolsArpTable(ctx, serial).CreateDeviceLiveToolsArpTableRequest(createDeviceLiveToolsArpTableRequest).Execute()

Enqueue a job to perform a ARP table request for the device



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
	createDeviceLiveToolsArpTableRequest := *openapiclient.NewCreateDeviceLiveToolsArpTableRequest() // CreateDeviceLiveToolsArpTableRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveToolsAPI.CreateDeviceLiveToolsArpTable(context.Background(), serial).CreateDeviceLiveToolsArpTableRequest(createDeviceLiveToolsArpTableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsAPI.CreateDeviceLiveToolsArpTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceLiveToolsArpTable`: CreateDeviceLiveToolsArpTable201Response
	fmt.Fprintf(os.Stdout, "Response from `LiveToolsAPI.CreateDeviceLiveToolsArpTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsArpTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsArpTableRequest** | [**CreateDeviceLiveToolsArpTableRequest**](CreateDeviceLiveToolsArpTableRequest.md) |  | 

### Return type

[**CreateDeviceLiveToolsArpTable201Response**](CreateDeviceLiveToolsArpTable201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsCableTest

> CreateDeviceLiveToolsCableTest201Response CreateDeviceLiveToolsCableTest(ctx, serial).CreateDeviceLiveToolsCableTestRequest(createDeviceLiveToolsCableTestRequest).Execute()

Enqueue a job to perform a cable test for the device on the specified ports



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
	createDeviceLiveToolsCableTestRequest := *openapiclient.NewCreateDeviceLiveToolsCableTestRequest([]string{"Ports_example"}) // CreateDeviceLiveToolsCableTestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveToolsAPI.CreateDeviceLiveToolsCableTest(context.Background(), serial).CreateDeviceLiveToolsCableTestRequest(createDeviceLiveToolsCableTestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsAPI.CreateDeviceLiveToolsCableTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceLiveToolsCableTest`: CreateDeviceLiveToolsCableTest201Response
	fmt.Fprintf(os.Stdout, "Response from `LiveToolsAPI.CreateDeviceLiveToolsCableTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsCableTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsCableTestRequest** | [**CreateDeviceLiveToolsCableTestRequest**](CreateDeviceLiveToolsCableTestRequest.md) |  | 

### Return type

[**CreateDeviceLiveToolsCableTest201Response**](CreateDeviceLiveToolsCableTest201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsPing

> CreateDeviceLiveToolsPing201Response CreateDeviceLiveToolsPing(ctx, serial).CreateDeviceLiveToolsPingRequest(createDeviceLiveToolsPingRequest).Execute()

Enqueue a job to ping a target host from the device



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
	createDeviceLiveToolsPingRequest := *openapiclient.NewCreateDeviceLiveToolsPingRequest("Target_example") // CreateDeviceLiveToolsPingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveToolsAPI.CreateDeviceLiveToolsPing(context.Background(), serial).CreateDeviceLiveToolsPingRequest(createDeviceLiveToolsPingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsAPI.CreateDeviceLiveToolsPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceLiveToolsPing`: CreateDeviceLiveToolsPing201Response
	fmt.Fprintf(os.Stdout, "Response from `LiveToolsAPI.CreateDeviceLiveToolsPing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsPingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsPingRequest** | [**CreateDeviceLiveToolsPingRequest**](CreateDeviceLiveToolsPingRequest.md) |  | 

### Return type

[**CreateDeviceLiveToolsPing201Response**](CreateDeviceLiveToolsPing201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsPingDevice

> CreateDeviceLiveToolsPing201Response CreateDeviceLiveToolsPingDevice(ctx, serial).CreateDeviceLiveToolsPingDeviceRequest(createDeviceLiveToolsPingDeviceRequest).Execute()

Enqueue a job to check connectivity status to the device



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
	createDeviceLiveToolsPingDeviceRequest := *openapiclient.NewCreateDeviceLiveToolsPingDeviceRequest() // CreateDeviceLiveToolsPingDeviceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveToolsAPI.CreateDeviceLiveToolsPingDevice(context.Background(), serial).CreateDeviceLiveToolsPingDeviceRequest(createDeviceLiveToolsPingDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsAPI.CreateDeviceLiveToolsPingDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceLiveToolsPingDevice`: CreateDeviceLiveToolsPing201Response
	fmt.Fprintf(os.Stdout, "Response from `LiveToolsAPI.CreateDeviceLiveToolsPingDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsPingDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsPingDeviceRequest** | [**CreateDeviceLiveToolsPingDeviceRequest**](CreateDeviceLiveToolsPingDeviceRequest.md) |  | 

### Return type

[**CreateDeviceLiveToolsPing201Response**](CreateDeviceLiveToolsPing201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsThroughputTest

> CreateDeviceLiveToolsThroughputTest201Response CreateDeviceLiveToolsThroughputTest(ctx, serial).CreateDeviceLiveToolsArpTableRequest(createDeviceLiveToolsArpTableRequest).Execute()

Enqueue a job to test a device throughput, the test will run for 10 secs to test throughput



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
	createDeviceLiveToolsArpTableRequest := *openapiclient.NewCreateDeviceLiveToolsArpTableRequest() // CreateDeviceLiveToolsArpTableRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveToolsAPI.CreateDeviceLiveToolsThroughputTest(context.Background(), serial).CreateDeviceLiveToolsArpTableRequest(createDeviceLiveToolsArpTableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsAPI.CreateDeviceLiveToolsThroughputTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceLiveToolsThroughputTest`: CreateDeviceLiveToolsThroughputTest201Response
	fmt.Fprintf(os.Stdout, "Response from `LiveToolsAPI.CreateDeviceLiveToolsThroughputTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsThroughputTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsArpTableRequest** | [**CreateDeviceLiveToolsArpTableRequest**](CreateDeviceLiveToolsArpTableRequest.md) |  | 

### Return type

[**CreateDeviceLiveToolsThroughputTest201Response**](CreateDeviceLiveToolsThroughputTest201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDeviceLiveToolsWakeOnLan

> CreateDeviceLiveToolsWakeOnLan201Response CreateDeviceLiveToolsWakeOnLan(ctx, serial).CreateDeviceLiveToolsWakeOnLanRequest(createDeviceLiveToolsWakeOnLanRequest).Execute()

Enqueue a job to send a Wake-on-LAN packet from the device



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
	createDeviceLiveToolsWakeOnLanRequest := *openapiclient.NewCreateDeviceLiveToolsWakeOnLanRequest(int32(123), "Mac_example") // CreateDeviceLiveToolsWakeOnLanRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveToolsAPI.CreateDeviceLiveToolsWakeOnLan(context.Background(), serial).CreateDeviceLiveToolsWakeOnLanRequest(createDeviceLiveToolsWakeOnLanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsAPI.CreateDeviceLiveToolsWakeOnLan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceLiveToolsWakeOnLan`: CreateDeviceLiveToolsWakeOnLan201Response
	fmt.Fprintf(os.Stdout, "Response from `LiveToolsAPI.CreateDeviceLiveToolsWakeOnLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsWakeOnLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsWakeOnLanRequest** | [**CreateDeviceLiveToolsWakeOnLanRequest**](CreateDeviceLiveToolsWakeOnLanRequest.md) |  | 

### Return type

[**CreateDeviceLiveToolsWakeOnLan201Response**](CreateDeviceLiveToolsWakeOnLan201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CycleDeviceSwitchPorts

> CycleDeviceSwitchPorts200Response CycleDeviceSwitchPorts(ctx, serial).CycleDeviceSwitchPortsRequest(cycleDeviceSwitchPortsRequest).Execute()

Cycle a set of switch ports



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
	cycleDeviceSwitchPortsRequest := *openapiclient.NewCycleDeviceSwitchPortsRequest([]string{"Ports_example"}) // CycleDeviceSwitchPortsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveToolsAPI.CycleDeviceSwitchPorts(context.Background(), serial).CycleDeviceSwitchPortsRequest(cycleDeviceSwitchPortsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsAPI.CycleDeviceSwitchPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CycleDeviceSwitchPorts`: CycleDeviceSwitchPorts200Response
	fmt.Fprintf(os.Stdout, "Response from `LiveToolsAPI.CycleDeviceSwitchPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCycleDeviceSwitchPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cycleDeviceSwitchPortsRequest** | [**CycleDeviceSwitchPortsRequest**](CycleDeviceSwitchPortsRequest.md) |  | 

### Return type

[**CycleDeviceSwitchPorts200Response**](CycleDeviceSwitchPorts200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsArpTable

> DevicesSerialLiveToolsArpTablePostRequestMessage GetDeviceLiveToolsArpTable(ctx, serial, arpTableId).Execute()

Return an ARP table live tool job.



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
	arpTableId := "arpTableId_example" // string | Arp table ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveToolsAPI.GetDeviceLiveToolsArpTable(context.Background(), serial, arpTableId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsAPI.GetDeviceLiveToolsArpTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceLiveToolsArpTable`: DevicesSerialLiveToolsArpTablePostRequestMessage
	fmt.Fprintf(os.Stdout, "Response from `LiveToolsAPI.GetDeviceLiveToolsArpTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**arpTableId** | **string** | Arp table ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsArpTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DevicesSerialLiveToolsArpTablePostRequestMessage**](DevicesSerialLiveToolsArpTablePostRequestMessage.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsCableTest

> DevicesSerialLiveToolsCableTestPostRequestMessage GetDeviceLiveToolsCableTest(ctx, serial, id).Execute()

Return a cable test live tool job.



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
	id := "id_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveToolsAPI.GetDeviceLiveToolsCableTest(context.Background(), serial, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsAPI.GetDeviceLiveToolsCableTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceLiveToolsCableTest`: DevicesSerialLiveToolsCableTestPostRequestMessage
	fmt.Fprintf(os.Stdout, "Response from `LiveToolsAPI.GetDeviceLiveToolsCableTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsCableTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DevicesSerialLiveToolsCableTestPostRequestMessage**](DevicesSerialLiveToolsCableTestPostRequestMessage.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsPing

> DevicesSerialLiveToolsPingPostRequestMessage GetDeviceLiveToolsPing(ctx, serial, id).Execute()

Return a ping job



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
	id := "id_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveToolsAPI.GetDeviceLiveToolsPing(context.Background(), serial, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsAPI.GetDeviceLiveToolsPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceLiveToolsPing`: DevicesSerialLiveToolsPingPostRequestMessage
	fmt.Fprintf(os.Stdout, "Response from `LiveToolsAPI.GetDeviceLiveToolsPing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsPingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DevicesSerialLiveToolsPingPostRequestMessage**](DevicesSerialLiveToolsPingPostRequestMessage.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsPingDevice

> GetDeviceLiveToolsPingDevice200Response GetDeviceLiveToolsPingDevice(ctx, serial, id).Execute()

Return a ping device job



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
	id := "id_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveToolsAPI.GetDeviceLiveToolsPingDevice(context.Background(), serial, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsAPI.GetDeviceLiveToolsPingDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceLiveToolsPingDevice`: GetDeviceLiveToolsPingDevice200Response
	fmt.Fprintf(os.Stdout, "Response from `LiveToolsAPI.GetDeviceLiveToolsPingDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsPingDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDeviceLiveToolsPingDevice200Response**](GetDeviceLiveToolsPingDevice200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsThroughputTest

> DevicesSerialLiveToolsThroughputTestPostRequestMessage GetDeviceLiveToolsThroughputTest(ctx, serial, throughputTestId).Execute()

Return a throughput test job



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
	throughputTestId := "throughputTestId_example" // string | Throughput test ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveToolsAPI.GetDeviceLiveToolsThroughputTest(context.Background(), serial, throughputTestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsAPI.GetDeviceLiveToolsThroughputTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceLiveToolsThroughputTest`: DevicesSerialLiveToolsThroughputTestPostRequestMessage
	fmt.Fprintf(os.Stdout, "Response from `LiveToolsAPI.GetDeviceLiveToolsThroughputTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**throughputTestId** | **string** | Throughput test ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsThroughputTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DevicesSerialLiveToolsThroughputTestPostRequestMessage**](DevicesSerialLiveToolsThroughputTestPostRequestMessage.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsWakeOnLan

> DevicesSerialLiveToolsWakeOnLanPostRequestMessage GetDeviceLiveToolsWakeOnLan(ctx, serial, wakeOnLanId).Execute()

Return a Wake-on-LAN job



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
	wakeOnLanId := "wakeOnLanId_example" // string | Wake on lan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveToolsAPI.GetDeviceLiveToolsWakeOnLan(context.Background(), serial, wakeOnLanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsAPI.GetDeviceLiveToolsWakeOnLan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceLiveToolsWakeOnLan`: DevicesSerialLiveToolsWakeOnLanPostRequestMessage
	fmt.Fprintf(os.Stdout, "Response from `LiveToolsAPI.GetDeviceLiveToolsWakeOnLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**wakeOnLanId** | **string** | Wake on lan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsWakeOnLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DevicesSerialLiveToolsWakeOnLanPostRequestMessage**](DevicesSerialLiveToolsWakeOnLanPostRequestMessage.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootDevice

> RebootDevice202Response RebootDevice(ctx, serial).Execute()

Reboot a device



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
	resp, r, err := apiClient.LiveToolsAPI.RebootDevice(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveToolsAPI.RebootDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebootDevice`: RebootDevice202Response
	fmt.Fprintf(os.Stdout, "Response from `LiveToolsAPI.RebootDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RebootDevice202Response**](RebootDevice202Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

