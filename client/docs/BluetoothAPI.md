# \BluetoothAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceWirelessBluetoothSettings**](BluetoothAPI.md#GetDeviceWirelessBluetoothSettings) | **Get** /devices/{serial}/wireless/bluetooth/settings | Return the bluetooth settings for a wireless device
[**GetNetworkWirelessBluetoothSettings**](BluetoothAPI.md#GetNetworkWirelessBluetoothSettings) | **Get** /networks/{networkId}/wireless/bluetooth/settings | Return the Bluetooth settings for a network. &lt;a href&#x3D;\&quot;https://documentation.meraki.com/MR/Bluetooth/Bluetooth_Low_Energy_(BLE)\&quot;&gt;Bluetooth settings&lt;/a&gt; must be enabled on the network.
[**UpdateDeviceWirelessBluetoothSettings**](BluetoothAPI.md#UpdateDeviceWirelessBluetoothSettings) | **Put** /devices/{serial}/wireless/bluetooth/settings | Update the bluetooth settings for a wireless device
[**UpdateNetworkWirelessBluetoothSettings**](BluetoothAPI.md#UpdateNetworkWirelessBluetoothSettings) | **Put** /networks/{networkId}/wireless/bluetooth/settings | Update the Bluetooth settings for a network



## GetDeviceWirelessBluetoothSettings

> GetDeviceWirelessBluetoothSettings200Response GetDeviceWirelessBluetoothSettings(ctx, serial).Execute()

Return the bluetooth settings for a wireless device



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
	resp, r, err := apiClient.BluetoothAPI.GetDeviceWirelessBluetoothSettings(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.GetDeviceWirelessBluetoothSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceWirelessBluetoothSettings`: GetDeviceWirelessBluetoothSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.GetDeviceWirelessBluetoothSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceWirelessBluetoothSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceWirelessBluetoothSettings200Response**](GetDeviceWirelessBluetoothSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessBluetoothSettings

> GetNetworkWirelessBluetoothSettings200Response GetNetworkWirelessBluetoothSettings(ctx, networkId).Execute()

Return the Bluetooth settings for a network. <a href=\"https://documentation.meraki.com/MR/Bluetooth/Bluetooth_Low_Energy_(BLE)\">Bluetooth settings</a> must be enabled on the network.



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
	resp, r, err := apiClient.BluetoothAPI.GetNetworkWirelessBluetoothSettings(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.GetNetworkWirelessBluetoothSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessBluetoothSettings`: GetNetworkWirelessBluetoothSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.GetNetworkWirelessBluetoothSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessBluetoothSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkWirelessBluetoothSettings200Response**](GetNetworkWirelessBluetoothSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceWirelessBluetoothSettings

> GetDeviceWirelessBluetoothSettings200Response UpdateDeviceWirelessBluetoothSettings(ctx, serial).UpdateDeviceWirelessBluetoothSettingsRequest(updateDeviceWirelessBluetoothSettingsRequest).Execute()

Update the bluetooth settings for a wireless device



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
	updateDeviceWirelessBluetoothSettingsRequest := *openapiclient.NewUpdateDeviceWirelessBluetoothSettingsRequest() // UpdateDeviceWirelessBluetoothSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothAPI.UpdateDeviceWirelessBluetoothSettings(context.Background(), serial).UpdateDeviceWirelessBluetoothSettingsRequest(updateDeviceWirelessBluetoothSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.UpdateDeviceWirelessBluetoothSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceWirelessBluetoothSettings`: GetDeviceWirelessBluetoothSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.UpdateDeviceWirelessBluetoothSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceWirelessBluetoothSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceWirelessBluetoothSettingsRequest** | [**UpdateDeviceWirelessBluetoothSettingsRequest**](UpdateDeviceWirelessBluetoothSettingsRequest.md) |  | 

### Return type

[**GetDeviceWirelessBluetoothSettings200Response**](GetDeviceWirelessBluetoothSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessBluetoothSettings

> GetNetworkWirelessBluetoothSettings200Response UpdateNetworkWirelessBluetoothSettings(ctx, networkId).UpdateNetworkWirelessBluetoothSettingsRequest(updateNetworkWirelessBluetoothSettingsRequest).Execute()

Update the Bluetooth settings for a network



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
	updateNetworkWirelessBluetoothSettingsRequest := *openapiclient.NewUpdateNetworkWirelessBluetoothSettingsRequest() // UpdateNetworkWirelessBluetoothSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BluetoothAPI.UpdateNetworkWirelessBluetoothSettings(context.Background(), networkId).UpdateNetworkWirelessBluetoothSettingsRequest(updateNetworkWirelessBluetoothSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BluetoothAPI.UpdateNetworkWirelessBluetoothSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkWirelessBluetoothSettings`: GetNetworkWirelessBluetoothSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `BluetoothAPI.UpdateNetworkWirelessBluetoothSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessBluetoothSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkWirelessBluetoothSettingsRequest** | [**UpdateNetworkWirelessBluetoothSettingsRequest**](UpdateNetworkWirelessBluetoothSettingsRequest.md) |  | 

### Return type

[**GetNetworkWirelessBluetoothSettings200Response**](GetNetworkWirelessBluetoothSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

