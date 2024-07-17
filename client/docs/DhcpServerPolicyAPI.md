# \DhcpServerPolicyAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer**](DhcpServerPolicyAPI.md#CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer) | **Post** /networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers | Add a server to be trusted by Dynamic ARP Inspection on this network
[**DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer**](DhcpServerPolicyAPI.md#DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer) | **Delete** /networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers/{trustedServerId} | Remove a server from being trusted by Dynamic ARP Inspection on this network
[**GetNetworkSwitchDhcpServerPolicy**](DhcpServerPolicyAPI.md#GetNetworkSwitchDhcpServerPolicy) | **Get** /networks/{networkId}/switch/dhcpServerPolicy | Return the DHCP server settings
[**GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers**](DhcpServerPolicyAPI.md#GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers) | **Get** /networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers | Return the list of servers trusted by Dynamic ARP Inspection on this network
[**GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice**](DhcpServerPolicyAPI.md#GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice) | **Get** /networks/{networkId}/switch/dhcpServerPolicy/arpInspection/warnings/byDevice | Return the devices that have a Dynamic ARP Inspection warning and their warnings
[**UpdateNetworkSwitchDhcpServerPolicy**](DhcpServerPolicyAPI.md#UpdateNetworkSwitchDhcpServerPolicy) | **Put** /networks/{networkId}/switch/dhcpServerPolicy | Update the DHCP server settings
[**UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer**](DhcpServerPolicyAPI.md#UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer) | **Put** /networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers/{trustedServerId} | Update a server that is trusted by Dynamic ARP Inspection on this network



## CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer

> GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx, networkId).CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest(createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest).Execute()

Add a server to be trusted by Dynamic ARP Inspection on this network



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
	createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest := *openapiclient.NewCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest("Mac_example", int32(123), *openapiclient.NewCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4()) // CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DhcpServerPolicyAPI.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(context.Background(), networkId).CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest(createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpServerPolicyAPI.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer`: GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DhcpServerPolicyAPI.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest** | [**CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest**](CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest.md) |  | 

### Return type

[**GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner**](GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer

> DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx, networkId, trustedServerId).Execute()

Remove a server from being trusted by Dynamic ARP Inspection on this network



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
	trustedServerId := "trustedServerId_example" // string | Trusted server ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DhcpServerPolicyAPI.DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(context.Background(), networkId, trustedServerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpServerPolicyAPI.DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**trustedServerId** | **string** | Trusted server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct via the builder pattern


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


## GetNetworkSwitchDhcpServerPolicy

> GetNetworkSwitchDhcpServerPolicy200Response GetNetworkSwitchDhcpServerPolicy(ctx, networkId).Execute()

Return the DHCP server settings



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
	resp, r, err := apiClient.DhcpServerPolicyAPI.GetNetworkSwitchDhcpServerPolicy(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpServerPolicyAPI.GetNetworkSwitchDhcpServerPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchDhcpServerPolicy`: GetNetworkSwitchDhcpServerPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `DhcpServerPolicyAPI.GetNetworkSwitchDhcpServerPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchDhcpServerPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchDhcpServerPolicy200Response**](GetNetworkSwitchDhcpServerPolicy200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers

> []GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers(ctx, networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return the list of servers trusted by Dynamic ARP Inspection on this network



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DhcpServerPolicyAPI.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers(context.Background(), networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpServerPolicyAPI.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers`: []GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DhcpServerPolicyAPI.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner**](GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice

> []GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice(ctx, networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return the devices that have a Dynamic ARP Inspection warning and their warnings



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
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DhcpServerPolicyAPI.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice(context.Background(), networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpServerPolicyAPI.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice`: []GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DhcpServerPolicyAPI.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner**](GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchDhcpServerPolicy

> GetNetworkSwitchDhcpServerPolicy200Response UpdateNetworkSwitchDhcpServerPolicy(ctx, networkId).UpdateNetworkSwitchDhcpServerPolicyRequest(updateNetworkSwitchDhcpServerPolicyRequest).Execute()

Update the DHCP server settings



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
	updateNetworkSwitchDhcpServerPolicyRequest := *openapiclient.NewUpdateNetworkSwitchDhcpServerPolicyRequest() // UpdateNetworkSwitchDhcpServerPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DhcpServerPolicyAPI.UpdateNetworkSwitchDhcpServerPolicy(context.Background(), networkId).UpdateNetworkSwitchDhcpServerPolicyRequest(updateNetworkSwitchDhcpServerPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpServerPolicyAPI.UpdateNetworkSwitchDhcpServerPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchDhcpServerPolicy`: GetNetworkSwitchDhcpServerPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `DhcpServerPolicyAPI.UpdateNetworkSwitchDhcpServerPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchDhcpServerPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchDhcpServerPolicyRequest** | [**UpdateNetworkSwitchDhcpServerPolicyRequest**](UpdateNetworkSwitchDhcpServerPolicyRequest.md) |  | 

### Return type

[**GetNetworkSwitchDhcpServerPolicy200Response**](GetNetworkSwitchDhcpServerPolicy200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer

> GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx, networkId, trustedServerId).UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest(updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest).Execute()

Update a server that is trusted by Dynamic ARP Inspection on this network



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
	trustedServerId := "trustedServerId_example" // string | Trusted server ID
	updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest := *openapiclient.NewUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest() // UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DhcpServerPolicyAPI.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(context.Background(), networkId, trustedServerId).UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest(updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpServerPolicyAPI.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer`: GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DhcpServerPolicyAPI.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**trustedServerId** | **string** | Trusted server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest** | [**UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest**](UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest.md) |  | 

### Return type

[**GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner**](GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

