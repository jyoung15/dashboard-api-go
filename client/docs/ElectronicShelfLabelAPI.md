# \ElectronicShelfLabelAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceWirelessElectronicShelfLabel**](ElectronicShelfLabelAPI.md#GetDeviceWirelessElectronicShelfLabel) | **Get** /devices/{serial}/wireless/electronicShelfLabel | Return the ESL settings of a device
[**GetNetworkWirelessElectronicShelfLabel**](ElectronicShelfLabelAPI.md#GetNetworkWirelessElectronicShelfLabel) | **Get** /networks/{networkId}/wireless/electronicShelfLabel | Return the ESL settings of a wireless network
[**GetNetworkWirelessElectronicShelfLabelConfiguredDevices**](ElectronicShelfLabelAPI.md#GetNetworkWirelessElectronicShelfLabelConfiguredDevices) | **Get** /networks/{networkId}/wireless/electronicShelfLabel/configuredDevices | Get a list of all ESL eligible devices of a network
[**UpdateDeviceWirelessElectronicShelfLabel**](ElectronicShelfLabelAPI.md#UpdateDeviceWirelessElectronicShelfLabel) | **Put** /devices/{serial}/wireless/electronicShelfLabel | Update the ESL settings of a device
[**UpdateNetworkWirelessElectronicShelfLabel**](ElectronicShelfLabelAPI.md#UpdateNetworkWirelessElectronicShelfLabel) | **Put** /networks/{networkId}/wireless/electronicShelfLabel | Update the ESL settings of a wireless network



## GetDeviceWirelessElectronicShelfLabel

> GetDeviceWirelessElectronicShelfLabel200Response GetDeviceWirelessElectronicShelfLabel(ctx, serial).Execute()

Return the ESL settings of a device



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
	resp, r, err := apiClient.ElectronicShelfLabelAPI.GetDeviceWirelessElectronicShelfLabel(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElectronicShelfLabelAPI.GetDeviceWirelessElectronicShelfLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceWirelessElectronicShelfLabel`: GetDeviceWirelessElectronicShelfLabel200Response
	fmt.Fprintf(os.Stdout, "Response from `ElectronicShelfLabelAPI.GetDeviceWirelessElectronicShelfLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceWirelessElectronicShelfLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceWirelessElectronicShelfLabel200Response**](GetDeviceWirelessElectronicShelfLabel200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessElectronicShelfLabel

> GetNetworkWirelessElectronicShelfLabel200Response GetNetworkWirelessElectronicShelfLabel(ctx, networkId).Execute()

Return the ESL settings of a wireless network



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
	resp, r, err := apiClient.ElectronicShelfLabelAPI.GetNetworkWirelessElectronicShelfLabel(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElectronicShelfLabelAPI.GetNetworkWirelessElectronicShelfLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessElectronicShelfLabel`: GetNetworkWirelessElectronicShelfLabel200Response
	fmt.Fprintf(os.Stdout, "Response from `ElectronicShelfLabelAPI.GetNetworkWirelessElectronicShelfLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessElectronicShelfLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkWirelessElectronicShelfLabel200Response**](GetNetworkWirelessElectronicShelfLabel200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessElectronicShelfLabelConfiguredDevices

> []GetNetworkWirelessElectronicShelfLabel200Response GetNetworkWirelessElectronicShelfLabelConfiguredDevices(ctx, networkId).Execute()

Get a list of all ESL eligible devices of a network



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
	resp, r, err := apiClient.ElectronicShelfLabelAPI.GetNetworkWirelessElectronicShelfLabelConfiguredDevices(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElectronicShelfLabelAPI.GetNetworkWirelessElectronicShelfLabelConfiguredDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessElectronicShelfLabelConfiguredDevices`: []GetNetworkWirelessElectronicShelfLabel200Response
	fmt.Fprintf(os.Stdout, "Response from `ElectronicShelfLabelAPI.GetNetworkWirelessElectronicShelfLabelConfiguredDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessElectronicShelfLabelConfiguredDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkWirelessElectronicShelfLabel200Response**](GetNetworkWirelessElectronicShelfLabel200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceWirelessElectronicShelfLabel

> GetDeviceWirelessElectronicShelfLabel200Response UpdateDeviceWirelessElectronicShelfLabel(ctx, serial).UpdateDeviceWirelessElectronicShelfLabelRequest(updateDeviceWirelessElectronicShelfLabelRequest).Execute()

Update the ESL settings of a device



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
	updateDeviceWirelessElectronicShelfLabelRequest := *openapiclient.NewUpdateDeviceWirelessElectronicShelfLabelRequest() // UpdateDeviceWirelessElectronicShelfLabelRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElectronicShelfLabelAPI.UpdateDeviceWirelessElectronicShelfLabel(context.Background(), serial).UpdateDeviceWirelessElectronicShelfLabelRequest(updateDeviceWirelessElectronicShelfLabelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElectronicShelfLabelAPI.UpdateDeviceWirelessElectronicShelfLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceWirelessElectronicShelfLabel`: GetDeviceWirelessElectronicShelfLabel200Response
	fmt.Fprintf(os.Stdout, "Response from `ElectronicShelfLabelAPI.UpdateDeviceWirelessElectronicShelfLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceWirelessElectronicShelfLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceWirelessElectronicShelfLabelRequest** | [**UpdateDeviceWirelessElectronicShelfLabelRequest**](UpdateDeviceWirelessElectronicShelfLabelRequest.md) |  | 

### Return type

[**GetDeviceWirelessElectronicShelfLabel200Response**](GetDeviceWirelessElectronicShelfLabel200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessElectronicShelfLabel

> GetNetworkWirelessElectronicShelfLabel200Response UpdateNetworkWirelessElectronicShelfLabel(ctx, networkId).UpdateNetworkWirelessElectronicShelfLabelRequest(updateNetworkWirelessElectronicShelfLabelRequest).Execute()

Update the ESL settings of a wireless network



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
	updateNetworkWirelessElectronicShelfLabelRequest := *openapiclient.NewUpdateNetworkWirelessElectronicShelfLabelRequest() // UpdateNetworkWirelessElectronicShelfLabelRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ElectronicShelfLabelAPI.UpdateNetworkWirelessElectronicShelfLabel(context.Background(), networkId).UpdateNetworkWirelessElectronicShelfLabelRequest(updateNetworkWirelessElectronicShelfLabelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ElectronicShelfLabelAPI.UpdateNetworkWirelessElectronicShelfLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkWirelessElectronicShelfLabel`: GetNetworkWirelessElectronicShelfLabel200Response
	fmt.Fprintf(os.Stdout, "Response from `ElectronicShelfLabelAPI.UpdateNetworkWirelessElectronicShelfLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessElectronicShelfLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkWirelessElectronicShelfLabelRequest** | [**UpdateNetworkWirelessElectronicShelfLabelRequest**](UpdateNetworkWirelessElectronicShelfLabelRequest.md) |  | 

### Return type

[**GetNetworkWirelessElectronicShelfLabel200Response**](GetNetworkWirelessElectronicShelfLabel200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

