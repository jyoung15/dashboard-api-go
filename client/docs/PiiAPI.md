# \PiiAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkPiiRequest**](PiiAPI.md#CreateNetworkPiiRequest) | **Post** /networks/{networkId}/pii/requests | Submit a new delete or restrict processing PII request
[**DeleteNetworkPiiRequest**](PiiAPI.md#DeleteNetworkPiiRequest) | **Delete** /networks/{networkId}/pii/requests/{requestId} | Delete a restrict processing PII request
[**GetNetworkPiiPiiKeys**](PiiAPI.md#GetNetworkPiiPiiKeys) | **Get** /networks/{networkId}/pii/piiKeys | List the keys required to access Personally Identifiable Information (PII) for a given identifier
[**GetNetworkPiiRequest**](PiiAPI.md#GetNetworkPiiRequest) | **Get** /networks/{networkId}/pii/requests/{requestId} | Return a PII request
[**GetNetworkPiiRequests**](PiiAPI.md#GetNetworkPiiRequests) | **Get** /networks/{networkId}/pii/requests | List the PII requests for this network or organization
[**GetNetworkPiiSmDevicesForKey**](PiiAPI.md#GetNetworkPiiSmDevicesForKey) | **Get** /networks/{networkId}/pii/smDevicesForKey | Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier
[**GetNetworkPiiSmOwnersForKey**](PiiAPI.md#GetNetworkPiiSmOwnersForKey) | **Get** /networks/{networkId}/pii/smOwnersForKey | Given a piece of Personally Identifiable Information (PII), return the Systems Manager owner ID(s) associated with that identifier



## CreateNetworkPiiRequest

> GetNetworkPiiRequests200ResponseInner CreateNetworkPiiRequest(ctx, networkId).CreateNetworkPiiRequestRequest(createNetworkPiiRequestRequest).Execute()

Submit a new delete or restrict processing PII request



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
	createNetworkPiiRequestRequest := *openapiclient.NewCreateNetworkPiiRequestRequest() // CreateNetworkPiiRequestRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PiiAPI.CreateNetworkPiiRequest(context.Background(), networkId).CreateNetworkPiiRequestRequest(createNetworkPiiRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PiiAPI.CreateNetworkPiiRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkPiiRequest`: GetNetworkPiiRequests200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PiiAPI.CreateNetworkPiiRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkPiiRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkPiiRequestRequest** | [**CreateNetworkPiiRequestRequest**](CreateNetworkPiiRequestRequest.md) |  | 

### Return type

[**GetNetworkPiiRequests200ResponseInner**](GetNetworkPiiRequests200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkPiiRequest

> DeleteNetworkPiiRequest(ctx, networkId, requestId).Execute()

Delete a restrict processing PII request



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
	requestId := "requestId_example" // string | Request ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PiiAPI.DeleteNetworkPiiRequest(context.Background(), networkId, requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PiiAPI.DeleteNetworkPiiRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkPiiRequestRequest struct via the builder pattern


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


## GetNetworkPiiPiiKeys

> map[string]GetNetworkPiiPiiKeys200ResponseValue GetNetworkPiiPiiKeys(ctx, networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()

List the keys required to access Personally Identifiable Information (PII) for a given identifier



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
	username := "username_example" // string | The username of a Systems Manager user (optional)
	email := "email_example" // string | The email of a network user account or a Systems Manager device (optional)
	mac := "mac_example" // string | The MAC of a network client device or a Systems Manager device (optional)
	serial := "serial_example" // string | The serial of a Systems Manager device (optional)
	imei := "imei_example" // string | The IMEI of a Systems Manager device (optional)
	bluetoothMac := "bluetoothMac_example" // string | The MAC of a Bluetooth client (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PiiAPI.GetNetworkPiiPiiKeys(context.Background(), networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PiiAPI.GetNetworkPiiPiiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPiiPiiKeys`: map[string]GetNetworkPiiPiiKeys200ResponseValue
	fmt.Fprintf(os.Stdout, "Response from `PiiAPI.GetNetworkPiiPiiKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiPiiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** | The username of a Systems Manager user | 
 **email** | **string** | The email of a network user account or a Systems Manager device | 
 **mac** | **string** | The MAC of a network client device or a Systems Manager device | 
 **serial** | **string** | The serial of a Systems Manager device | 
 **imei** | **string** | The IMEI of a Systems Manager device | 
 **bluetoothMac** | **string** | The MAC of a Bluetooth client | 

### Return type

[**map[string]GetNetworkPiiPiiKeys200ResponseValue**](GetNetworkPiiPiiKeys200ResponseValue.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPiiRequest

> GetNetworkPiiRequests200ResponseInner GetNetworkPiiRequest(ctx, networkId, requestId).Execute()

Return a PII request



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
	requestId := "requestId_example" // string | Request ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PiiAPI.GetNetworkPiiRequest(context.Background(), networkId, requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PiiAPI.GetNetworkPiiRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPiiRequest`: GetNetworkPiiRequests200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PiiAPI.GetNetworkPiiRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkPiiRequests200ResponseInner**](GetNetworkPiiRequests200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPiiRequests

> []GetNetworkPiiRequests200ResponseInner GetNetworkPiiRequests(ctx, networkId).Execute()

List the PII requests for this network or organization



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
	resp, r, err := apiClient.PiiAPI.GetNetworkPiiRequests(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PiiAPI.GetNetworkPiiRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPiiRequests`: []GetNetworkPiiRequests200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PiiAPI.GetNetworkPiiRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkPiiRequests200ResponseInner**](GetNetworkPiiRequests200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPiiSmDevicesForKey

> map[string][]string GetNetworkPiiSmDevicesForKey(ctx, networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()

Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier



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
	username := "username_example" // string | The username of a Systems Manager user (optional)
	email := "email_example" // string | The email of a network user account or a Systems Manager device (optional)
	mac := "mac_example" // string | The MAC of a network client device or a Systems Manager device (optional)
	serial := "serial_example" // string | The serial of a Systems Manager device (optional)
	imei := "imei_example" // string | The IMEI of a Systems Manager device (optional)
	bluetoothMac := "bluetoothMac_example" // string | The MAC of a Bluetooth client (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PiiAPI.GetNetworkPiiSmDevicesForKey(context.Background(), networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PiiAPI.GetNetworkPiiSmDevicesForKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPiiSmDevicesForKey`: map[string][]string
	fmt.Fprintf(os.Stdout, "Response from `PiiAPI.GetNetworkPiiSmDevicesForKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiSmDevicesForKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** | The username of a Systems Manager user | 
 **email** | **string** | The email of a network user account or a Systems Manager device | 
 **mac** | **string** | The MAC of a network client device or a Systems Manager device | 
 **serial** | **string** | The serial of a Systems Manager device | 
 **imei** | **string** | The IMEI of a Systems Manager device | 
 **bluetoothMac** | **string** | The MAC of a Bluetooth client | 

### Return type

[**map[string][]string**](array.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPiiSmOwnersForKey

> map[string][]string GetNetworkPiiSmOwnersForKey(ctx, networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()

Given a piece of Personally Identifiable Information (PII), return the Systems Manager owner ID(s) associated with that identifier



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
	username := "username_example" // string | The username of a Systems Manager user (optional)
	email := "email_example" // string | The email of a network user account or a Systems Manager device (optional)
	mac := "mac_example" // string | The MAC of a network client device or a Systems Manager device (optional)
	serial := "serial_example" // string | The serial of a Systems Manager device (optional)
	imei := "imei_example" // string | The IMEI of a Systems Manager device (optional)
	bluetoothMac := "bluetoothMac_example" // string | The MAC of a Bluetooth client (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PiiAPI.GetNetworkPiiSmOwnersForKey(context.Background(), networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PiiAPI.GetNetworkPiiSmOwnersForKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPiiSmOwnersForKey`: map[string][]string
	fmt.Fprintf(os.Stdout, "Response from `PiiAPI.GetNetworkPiiSmOwnersForKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiSmOwnersForKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** | The username of a Systems Manager user | 
 **email** | **string** | The email of a network user account or a Systems Manager device | 
 **mac** | **string** | The MAC of a network client device or a Systems Manager device | 
 **serial** | **string** | The serial of a Systems Manager device | 
 **imei** | **string** | The IMEI of a Systems Manager device | 
 **bluetoothMac** | **string** | The MAC of a Bluetooth client | 

### Return type

[**map[string][]string**](array.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

