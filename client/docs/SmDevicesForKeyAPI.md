# \SmDevicesForKeyAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkPiiSmDevicesForKey**](SmDevicesForKeyAPI.md#GetNetworkPiiSmDevicesForKey) | **Get** /networks/{networkId}/pii/smDevicesForKey | Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier



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
	resp, r, err := apiClient.SmDevicesForKeyAPI.GetNetworkPiiSmDevicesForKey(context.Background(), networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmDevicesForKeyAPI.GetNetworkPiiSmDevicesForKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPiiSmDevicesForKey`: map[string][]string
	fmt.Fprintf(os.Stdout, "Response from `SmDevicesForKeyAPI.GetNetworkPiiSmDevicesForKey`: %v\n", resp)
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

