# \RadioAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceApplianceRadioSettings**](RadioAPI.md#GetDeviceApplianceRadioSettings) | **Get** /devices/{serial}/appliance/radio/settings | Return the radio settings of an appliance
[**GetDeviceWirelessRadioSettings**](RadioAPI.md#GetDeviceWirelessRadioSettings) | **Get** /devices/{serial}/wireless/radio/settings | Return the radio settings of a device
[**UpdateDeviceApplianceRadioSettings**](RadioAPI.md#UpdateDeviceApplianceRadioSettings) | **Put** /devices/{serial}/appliance/radio/settings | Update the radio settings of an appliance
[**UpdateDeviceWirelessRadioSettings**](RadioAPI.md#UpdateDeviceWirelessRadioSettings) | **Put** /devices/{serial}/wireless/radio/settings | Update the radio settings of a device



## GetDeviceApplianceRadioSettings

> GetDeviceApplianceRadioSettings200Response GetDeviceApplianceRadioSettings(ctx, serial).Execute()

Return the radio settings of an appliance



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
	resp, r, err := apiClient.RadioAPI.GetDeviceApplianceRadioSettings(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadioAPI.GetDeviceApplianceRadioSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceApplianceRadioSettings`: GetDeviceApplianceRadioSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `RadioAPI.GetDeviceApplianceRadioSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceApplianceRadioSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceApplianceRadioSettings200Response**](GetDeviceApplianceRadioSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceWirelessRadioSettings

> map[string]interface{} GetDeviceWirelessRadioSettings(ctx, serial).Execute()

Return the radio settings of a device



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
	resp, r, err := apiClient.RadioAPI.GetDeviceWirelessRadioSettings(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadioAPI.GetDeviceWirelessRadioSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceWirelessRadioSettings`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RadioAPI.GetDeviceWirelessRadioSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceWirelessRadioSettingsRequest struct via the builder pattern


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


## UpdateDeviceApplianceRadioSettings

> GetDeviceApplianceRadioSettings200Response UpdateDeviceApplianceRadioSettings(ctx, serial).UpdateDeviceApplianceRadioSettingsRequest(updateDeviceApplianceRadioSettingsRequest).Execute()

Update the radio settings of an appliance



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
	updateDeviceApplianceRadioSettingsRequest := *openapiclient.NewUpdateDeviceApplianceRadioSettingsRequest() // UpdateDeviceApplianceRadioSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RadioAPI.UpdateDeviceApplianceRadioSettings(context.Background(), serial).UpdateDeviceApplianceRadioSettingsRequest(updateDeviceApplianceRadioSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadioAPI.UpdateDeviceApplianceRadioSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceApplianceRadioSettings`: GetDeviceApplianceRadioSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `RadioAPI.UpdateDeviceApplianceRadioSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceApplianceRadioSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceApplianceRadioSettingsRequest** | [**UpdateDeviceApplianceRadioSettingsRequest**](UpdateDeviceApplianceRadioSettingsRequest.md) |  | 

### Return type

[**GetDeviceApplianceRadioSettings200Response**](GetDeviceApplianceRadioSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceWirelessRadioSettings

> map[string]interface{} UpdateDeviceWirelessRadioSettings(ctx, serial).UpdateDeviceWirelessRadioSettingsRequest(updateDeviceWirelessRadioSettingsRequest).Execute()

Update the radio settings of a device



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
	updateDeviceWirelessRadioSettingsRequest := *openapiclient.NewUpdateDeviceWirelessRadioSettingsRequest() // UpdateDeviceWirelessRadioSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RadioAPI.UpdateDeviceWirelessRadioSettings(context.Background(), serial).UpdateDeviceWirelessRadioSettingsRequest(updateDeviceWirelessRadioSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadioAPI.UpdateDeviceWirelessRadioSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceWirelessRadioSettings`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RadioAPI.UpdateDeviceWirelessRadioSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceWirelessRadioSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceWirelessRadioSettingsRequest** | [**UpdateDeviceWirelessRadioSettingsRequest**](UpdateDeviceWirelessRadioSettingsRequest.md) |  | 

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

