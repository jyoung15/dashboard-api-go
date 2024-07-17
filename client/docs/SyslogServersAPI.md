# \SyslogServersAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSyslogServers**](SyslogServersAPI.md#GetNetworkSyslogServers) | **Get** /networks/{networkId}/syslogServers | List the syslog servers for a network
[**UpdateNetworkSyslogServers**](SyslogServersAPI.md#UpdateNetworkSyslogServers) | **Put** /networks/{networkId}/syslogServers | Update the syslog servers for a network



## GetNetworkSyslogServers

> GetNetworkSyslogServers200Response GetNetworkSyslogServers(ctx, networkId).Execute()

List the syslog servers for a network



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
	resp, r, err := apiClient.SyslogServersAPI.GetNetworkSyslogServers(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogServersAPI.GetNetworkSyslogServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSyslogServers`: GetNetworkSyslogServers200Response
	fmt.Fprintf(os.Stdout, "Response from `SyslogServersAPI.GetNetworkSyslogServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSyslogServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSyslogServers200Response**](GetNetworkSyslogServers200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSyslogServers

> GetNetworkSyslogServers200Response UpdateNetworkSyslogServers(ctx, networkId).UpdateNetworkSyslogServersRequest(updateNetworkSyslogServersRequest).Execute()

Update the syslog servers for a network



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
	updateNetworkSyslogServersRequest := *openapiclient.NewUpdateNetworkSyslogServersRequest([]openapiclient.UpdateNetworkSyslogServersRequestServersInner{*openapiclient.NewUpdateNetworkSyslogServersRequestServersInner("Host_example", int32(123), []string{"Roles_example"})}) // UpdateNetworkSyslogServersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SyslogServersAPI.UpdateNetworkSyslogServers(context.Background(), networkId).UpdateNetworkSyslogServersRequest(updateNetworkSyslogServersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogServersAPI.UpdateNetworkSyslogServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSyslogServers`: GetNetworkSyslogServers200Response
	fmt.Fprintf(os.Stdout, "Response from `SyslogServersAPI.UpdateNetworkSyslogServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSyslogServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSyslogServersRequest** | [**UpdateNetworkSyslogServersRequest**](UpdateNetworkSyslogServersRequest.md) |  | 

### Return type

[**GetNetworkSyslogServers200Response**](GetNetworkSyslogServers200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

