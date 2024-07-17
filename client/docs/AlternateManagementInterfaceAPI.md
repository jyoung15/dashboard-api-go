# \AlternateManagementInterfaceAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSwitchAlternateManagementInterface**](AlternateManagementInterfaceAPI.md#GetNetworkSwitchAlternateManagementInterface) | **Get** /networks/{networkId}/switch/alternateManagementInterface | Return the switch alternate management interface for the network
[**GetNetworkWirelessAlternateManagementInterface**](AlternateManagementInterfaceAPI.md#GetNetworkWirelessAlternateManagementInterface) | **Get** /networks/{networkId}/wireless/alternateManagementInterface | Return alternate management interface and devices with IP assigned
[**UpdateDeviceWirelessAlternateManagementInterfaceIpv6**](AlternateManagementInterfaceAPI.md#UpdateDeviceWirelessAlternateManagementInterfaceIpv6) | **Put** /devices/{serial}/wireless/alternateManagementInterface/ipv6 | Update alternate management interface IPv6 address
[**UpdateNetworkSwitchAlternateManagementInterface**](AlternateManagementInterfaceAPI.md#UpdateNetworkSwitchAlternateManagementInterface) | **Put** /networks/{networkId}/switch/alternateManagementInterface | Update the switch alternate management interface for the network
[**UpdateNetworkWirelessAlternateManagementInterface**](AlternateManagementInterfaceAPI.md#UpdateNetworkWirelessAlternateManagementInterface) | **Put** /networks/{networkId}/wireless/alternateManagementInterface | Update alternate management interface and device static IP



## GetNetworkSwitchAlternateManagementInterface

> GetNetworkSwitchAlternateManagementInterface200Response GetNetworkSwitchAlternateManagementInterface(ctx, networkId).Execute()

Return the switch alternate management interface for the network



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
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlternateManagementInterfaceAPI.GetNetworkSwitchAlternateManagementInterface(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternateManagementInterfaceAPI.GetNetworkSwitchAlternateManagementInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchAlternateManagementInterface`: GetNetworkSwitchAlternateManagementInterface200Response
	fmt.Fprintf(os.Stdout, "Response from `AlternateManagementInterfaceAPI.GetNetworkSwitchAlternateManagementInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchAlternateManagementInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchAlternateManagementInterface200Response**](GetNetworkSwitchAlternateManagementInterface200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessAlternateManagementInterface

> map[string]interface{} GetNetworkWirelessAlternateManagementInterface(ctx, networkId).Execute()

Return alternate management interface and devices with IP assigned



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
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlternateManagementInterfaceAPI.GetNetworkWirelessAlternateManagementInterface(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternateManagementInterfaceAPI.GetNetworkWirelessAlternateManagementInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessAlternateManagementInterface`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AlternateManagementInterfaceAPI.GetNetworkWirelessAlternateManagementInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessAlternateManagementInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceWirelessAlternateManagementInterfaceIpv6

> UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response UpdateDeviceWirelessAlternateManagementInterfaceIpv6(ctx, serial).UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request(updateDeviceWirelessAlternateManagementInterfaceIpv6Request).Execute()

Update alternate management interface IPv6 address



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
	updateDeviceWirelessAlternateManagementInterfaceIpv6Request := *openapiclient.NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request() // UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlternateManagementInterfaceAPI.UpdateDeviceWirelessAlternateManagementInterfaceIpv6(context.Background(), serial).UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request(updateDeviceWirelessAlternateManagementInterfaceIpv6Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternateManagementInterfaceAPI.UpdateDeviceWirelessAlternateManagementInterfaceIpv6``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceWirelessAlternateManagementInterfaceIpv6`: UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response
	fmt.Fprintf(os.Stdout, "Response from `AlternateManagementInterfaceAPI.UpdateDeviceWirelessAlternateManagementInterfaceIpv6`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceWirelessAlternateManagementInterfaceIpv6Request** | [**UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request**](UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request.md) |  | 

### Return type

[**UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response**](UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchAlternateManagementInterface

> GetNetworkSwitchAlternateManagementInterface200Response UpdateNetworkSwitchAlternateManagementInterface(ctx, networkId).UpdateNetworkSwitchAlternateManagementInterfaceRequest(updateNetworkSwitchAlternateManagementInterfaceRequest).Execute()

Update the switch alternate management interface for the network



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
	networkId := "networkId_example" // string | Network ID
	updateNetworkSwitchAlternateManagementInterfaceRequest := *openapiclient.NewUpdateNetworkSwitchAlternateManagementInterfaceRequest() // UpdateNetworkSwitchAlternateManagementInterfaceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlternateManagementInterfaceAPI.UpdateNetworkSwitchAlternateManagementInterface(context.Background(), networkId).UpdateNetworkSwitchAlternateManagementInterfaceRequest(updateNetworkSwitchAlternateManagementInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternateManagementInterfaceAPI.UpdateNetworkSwitchAlternateManagementInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchAlternateManagementInterface`: GetNetworkSwitchAlternateManagementInterface200Response
	fmt.Fprintf(os.Stdout, "Response from `AlternateManagementInterfaceAPI.UpdateNetworkSwitchAlternateManagementInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchAlternateManagementInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchAlternateManagementInterfaceRequest** | [**UpdateNetworkSwitchAlternateManagementInterfaceRequest**](UpdateNetworkSwitchAlternateManagementInterfaceRequest.md) |  | 

### Return type

[**GetNetworkSwitchAlternateManagementInterface200Response**](GetNetworkSwitchAlternateManagementInterface200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessAlternateManagementInterface

> map[string]interface{} UpdateNetworkWirelessAlternateManagementInterface(ctx, networkId).UpdateNetworkWirelessAlternateManagementInterfaceRequest(updateNetworkWirelessAlternateManagementInterfaceRequest).Execute()

Update alternate management interface and device static IP



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
	networkId := "networkId_example" // string | Network ID
	updateNetworkWirelessAlternateManagementInterfaceRequest := *openapiclient.NewUpdateNetworkWirelessAlternateManagementInterfaceRequest() // UpdateNetworkWirelessAlternateManagementInterfaceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlternateManagementInterfaceAPI.UpdateNetworkWirelessAlternateManagementInterface(context.Background(), networkId).UpdateNetworkWirelessAlternateManagementInterfaceRequest(updateNetworkWirelessAlternateManagementInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlternateManagementInterfaceAPI.UpdateNetworkWirelessAlternateManagementInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkWirelessAlternateManagementInterface`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AlternateManagementInterfaceAPI.UpdateNetworkWirelessAlternateManagementInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessAlternateManagementInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkWirelessAlternateManagementInterfaceRequest** | [**UpdateNetworkWirelessAlternateManagementInterfaceRequest**](UpdateNetworkWirelessAlternateManagementInterfaceRequest.md) |  | 

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

