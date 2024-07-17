# \WirelessProfilesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkCameraWirelessProfile**](WirelessProfilesAPI.md#CreateNetworkCameraWirelessProfile) | **Post** /networks/{networkId}/camera/wirelessProfiles | Creates a new camera wireless profile for this network.
[**DeleteNetworkCameraWirelessProfile**](WirelessProfilesAPI.md#DeleteNetworkCameraWirelessProfile) | **Delete** /networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId} | Delete an existing camera wireless profile for this network.
[**GetDeviceCameraWirelessProfiles**](WirelessProfilesAPI.md#GetDeviceCameraWirelessProfiles) | **Get** /devices/{serial}/camera/wirelessProfiles | Returns wireless profile assigned to the given camera
[**GetNetworkCameraWirelessProfile**](WirelessProfilesAPI.md#GetNetworkCameraWirelessProfile) | **Get** /networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId} | Retrieve a single camera wireless profile.
[**GetNetworkCameraWirelessProfiles**](WirelessProfilesAPI.md#GetNetworkCameraWirelessProfiles) | **Get** /networks/{networkId}/camera/wirelessProfiles | List the camera wireless profiles for this network.
[**UpdateDeviceCameraWirelessProfiles**](WirelessProfilesAPI.md#UpdateDeviceCameraWirelessProfiles) | **Put** /devices/{serial}/camera/wirelessProfiles | Assign wireless profiles to the given camera
[**UpdateNetworkCameraWirelessProfile**](WirelessProfilesAPI.md#UpdateNetworkCameraWirelessProfile) | **Put** /networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId} | Update an existing camera wireless profile in this network.



## CreateNetworkCameraWirelessProfile

> GetNetworkCameraWirelessProfiles200ResponseInner CreateNetworkCameraWirelessProfile(ctx, networkId).CreateNetworkCameraWirelessProfileRequest(createNetworkCameraWirelessProfileRequest).Execute()

Creates a new camera wireless profile for this network.



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
	createNetworkCameraWirelessProfileRequest := *openapiclient.NewCreateNetworkCameraWirelessProfileRequest("Name_example", *openapiclient.NewCreateNetworkCameraWirelessProfileRequestSsid()) // CreateNetworkCameraWirelessProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessProfilesAPI.CreateNetworkCameraWirelessProfile(context.Background(), networkId).CreateNetworkCameraWirelessProfileRequest(createNetworkCameraWirelessProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessProfilesAPI.CreateNetworkCameraWirelessProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkCameraWirelessProfile`: GetNetworkCameraWirelessProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `WirelessProfilesAPI.CreateNetworkCameraWirelessProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkCameraWirelessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkCameraWirelessProfileRequest** | [**CreateNetworkCameraWirelessProfileRequest**](CreateNetworkCameraWirelessProfileRequest.md) |  | 

### Return type

[**GetNetworkCameraWirelessProfiles200ResponseInner**](GetNetworkCameraWirelessProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkCameraWirelessProfile

> DeleteNetworkCameraWirelessProfile(ctx, networkId, wirelessProfileId).Execute()

Delete an existing camera wireless profile for this network.



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
	wirelessProfileId := "wirelessProfileId_example" // string | Wireless profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WirelessProfilesAPI.DeleteNetworkCameraWirelessProfile(context.Background(), networkId, wirelessProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessProfilesAPI.DeleteNetworkCameraWirelessProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**wirelessProfileId** | **string** | Wireless profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkCameraWirelessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceCameraWirelessProfiles

> map[string]interface{} GetDeviceCameraWirelessProfiles(ctx, serial).Execute()

Returns wireless profile assigned to the given camera



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
	resp, r, err := apiClient.WirelessProfilesAPI.GetDeviceCameraWirelessProfiles(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessProfilesAPI.GetDeviceCameraWirelessProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraWirelessProfiles`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WirelessProfilesAPI.GetDeviceCameraWirelessProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraWirelessProfilesRequest struct via the builder pattern


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


## GetNetworkCameraWirelessProfile

> GetNetworkCameraWirelessProfiles200ResponseInner GetNetworkCameraWirelessProfile(ctx, networkId, wirelessProfileId).Execute()

Retrieve a single camera wireless profile.



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
	wirelessProfileId := "wirelessProfileId_example" // string | Wireless profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessProfilesAPI.GetNetworkCameraWirelessProfile(context.Background(), networkId, wirelessProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessProfilesAPI.GetNetworkCameraWirelessProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkCameraWirelessProfile`: GetNetworkCameraWirelessProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `WirelessProfilesAPI.GetNetworkCameraWirelessProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**wirelessProfileId** | **string** | Wireless profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCameraWirelessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkCameraWirelessProfiles200ResponseInner**](GetNetworkCameraWirelessProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkCameraWirelessProfiles

> []GetNetworkCameraWirelessProfiles200ResponseInner GetNetworkCameraWirelessProfiles(ctx, networkId).Execute()

List the camera wireless profiles for this network.



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
	resp, r, err := apiClient.WirelessProfilesAPI.GetNetworkCameraWirelessProfiles(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessProfilesAPI.GetNetworkCameraWirelessProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkCameraWirelessProfiles`: []GetNetworkCameraWirelessProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `WirelessProfilesAPI.GetNetworkCameraWirelessProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCameraWirelessProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkCameraWirelessProfiles200ResponseInner**](GetNetworkCameraWirelessProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCameraWirelessProfiles

> map[string]interface{} UpdateDeviceCameraWirelessProfiles(ctx, serial).UpdateDeviceCameraWirelessProfilesRequest(updateDeviceCameraWirelessProfilesRequest).Execute()

Assign wireless profiles to the given camera



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
	updateDeviceCameraWirelessProfilesRequest := *openapiclient.NewUpdateDeviceCameraWirelessProfilesRequest(*openapiclient.NewUpdateDeviceCameraWirelessProfilesRequestIds()) // UpdateDeviceCameraWirelessProfilesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessProfilesAPI.UpdateDeviceCameraWirelessProfiles(context.Background(), serial).UpdateDeviceCameraWirelessProfilesRequest(updateDeviceCameraWirelessProfilesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessProfilesAPI.UpdateDeviceCameraWirelessProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceCameraWirelessProfiles`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WirelessProfilesAPI.UpdateDeviceCameraWirelessProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCameraWirelessProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCameraWirelessProfilesRequest** | [**UpdateDeviceCameraWirelessProfilesRequest**](UpdateDeviceCameraWirelessProfilesRequest.md) |  | 

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


## UpdateNetworkCameraWirelessProfile

> GetNetworkCameraWirelessProfiles200ResponseInner UpdateNetworkCameraWirelessProfile(ctx, networkId, wirelessProfileId).UpdateNetworkCameraWirelessProfileRequest(updateNetworkCameraWirelessProfileRequest).Execute()

Update an existing camera wireless profile in this network.



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
	wirelessProfileId := "wirelessProfileId_example" // string | Wireless profile ID
	updateNetworkCameraWirelessProfileRequest := *openapiclient.NewUpdateNetworkCameraWirelessProfileRequest() // UpdateNetworkCameraWirelessProfileRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessProfilesAPI.UpdateNetworkCameraWirelessProfile(context.Background(), networkId, wirelessProfileId).UpdateNetworkCameraWirelessProfileRequest(updateNetworkCameraWirelessProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessProfilesAPI.UpdateNetworkCameraWirelessProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkCameraWirelessProfile`: GetNetworkCameraWirelessProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `WirelessProfilesAPI.UpdateNetworkCameraWirelessProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**wirelessProfileId** | **string** | Wireless profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCameraWirelessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkCameraWirelessProfileRequest** | [**UpdateNetworkCameraWirelessProfileRequest**](UpdateNetworkCameraWirelessProfileRequest.md) |  | 

### Return type

[**GetNetworkCameraWirelessProfiles200ResponseInner**](GetNetworkCameraWirelessProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

