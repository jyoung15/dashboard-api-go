# \RendezvousPointsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkSwitchRoutingMulticastRendezvousPoint**](RendezvousPointsAPI.md#CreateNetworkSwitchRoutingMulticastRendezvousPoint) | **Post** /networks/{networkId}/switch/routing/multicast/rendezvousPoints | Create a multicast rendezvous point
[**DeleteNetworkSwitchRoutingMulticastRendezvousPoint**](RendezvousPointsAPI.md#DeleteNetworkSwitchRoutingMulticastRendezvousPoint) | **Delete** /networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId} | Delete a multicast rendezvous point
[**GetNetworkSwitchRoutingMulticastRendezvousPoint**](RendezvousPointsAPI.md#GetNetworkSwitchRoutingMulticastRendezvousPoint) | **Get** /networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId} | Return a multicast rendezvous point
[**GetNetworkSwitchRoutingMulticastRendezvousPoints**](RendezvousPointsAPI.md#GetNetworkSwitchRoutingMulticastRendezvousPoints) | **Get** /networks/{networkId}/switch/routing/multicast/rendezvousPoints | List multicast rendezvous points
[**UpdateNetworkSwitchRoutingMulticastRendezvousPoint**](RendezvousPointsAPI.md#UpdateNetworkSwitchRoutingMulticastRendezvousPoint) | **Put** /networks/{networkId}/switch/routing/multicast/rendezvousPoints/{rendezvousPointId} | Update a multicast rendezvous point



## CreateNetworkSwitchRoutingMulticastRendezvousPoint

> GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner CreateNetworkSwitchRoutingMulticastRendezvousPoint(ctx, networkId).CreateNetworkSwitchRoutingMulticastRendezvousPointRequest(createNetworkSwitchRoutingMulticastRendezvousPointRequest).Execute()

Create a multicast rendezvous point



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
	createNetworkSwitchRoutingMulticastRendezvousPointRequest := *openapiclient.NewCreateNetworkSwitchRoutingMulticastRendezvousPointRequest("InterfaceIp_example", "MulticastGroup_example") // CreateNetworkSwitchRoutingMulticastRendezvousPointRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RendezvousPointsAPI.CreateNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId).CreateNetworkSwitchRoutingMulticastRendezvousPointRequest(createNetworkSwitchRoutingMulticastRendezvousPointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RendezvousPointsAPI.CreateNetworkSwitchRoutingMulticastRendezvousPoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSwitchRoutingMulticastRendezvousPoint`: GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RendezvousPointsAPI.CreateNetworkSwitchRoutingMulticastRendezvousPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchRoutingMulticastRendezvousPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchRoutingMulticastRendezvousPointRequest** | [**CreateNetworkSwitchRoutingMulticastRendezvousPointRequest**](CreateNetworkSwitchRoutingMulticastRendezvousPointRequest.md) |  | 

### Return type

[**GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner**](GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSwitchRoutingMulticastRendezvousPoint

> DeleteNetworkSwitchRoutingMulticastRendezvousPoint(ctx, networkId, rendezvousPointId).Execute()

Delete a multicast rendezvous point



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
	rendezvousPointId := "rendezvousPointId_example" // string | Rendezvous point ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RendezvousPointsAPI.DeleteNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RendezvousPointsAPI.DeleteNetworkSwitchRoutingMulticastRendezvousPoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rendezvousPointId** | **string** | Rendezvous point ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchRoutingMulticastRendezvousPointRequest struct via the builder pattern


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


## GetNetworkSwitchRoutingMulticastRendezvousPoint

> GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner GetNetworkSwitchRoutingMulticastRendezvousPoint(ctx, networkId, rendezvousPointId).Execute()

Return a multicast rendezvous point



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
	rendezvousPointId := "rendezvousPointId_example" // string | Rendezvous point ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RendezvousPointsAPI.GetNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RendezvousPointsAPI.GetNetworkSwitchRoutingMulticastRendezvousPoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchRoutingMulticastRendezvousPoint`: GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RendezvousPointsAPI.GetNetworkSwitchRoutingMulticastRendezvousPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rendezvousPointId** | **string** | Rendezvous point ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchRoutingMulticastRendezvousPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner**](GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchRoutingMulticastRendezvousPoints

> []GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner GetNetworkSwitchRoutingMulticastRendezvousPoints(ctx, networkId).Execute()

List multicast rendezvous points



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
	resp, r, err := apiClient.RendezvousPointsAPI.GetNetworkSwitchRoutingMulticastRendezvousPoints(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RendezvousPointsAPI.GetNetworkSwitchRoutingMulticastRendezvousPoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchRoutingMulticastRendezvousPoints`: []GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RendezvousPointsAPI.GetNetworkSwitchRoutingMulticastRendezvousPoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchRoutingMulticastRendezvousPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner**](GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchRoutingMulticastRendezvousPoint

> GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner UpdateNetworkSwitchRoutingMulticastRendezvousPoint(ctx, networkId, rendezvousPointId).CreateNetworkSwitchRoutingMulticastRendezvousPointRequest(createNetworkSwitchRoutingMulticastRendezvousPointRequest).Execute()

Update a multicast rendezvous point



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
	rendezvousPointId := "rendezvousPointId_example" // string | Rendezvous point ID
	createNetworkSwitchRoutingMulticastRendezvousPointRequest := *openapiclient.NewCreateNetworkSwitchRoutingMulticastRendezvousPointRequest("InterfaceIp_example", "MulticastGroup_example") // CreateNetworkSwitchRoutingMulticastRendezvousPointRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RendezvousPointsAPI.UpdateNetworkSwitchRoutingMulticastRendezvousPoint(context.Background(), networkId, rendezvousPointId).CreateNetworkSwitchRoutingMulticastRendezvousPointRequest(createNetworkSwitchRoutingMulticastRendezvousPointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RendezvousPointsAPI.UpdateNetworkSwitchRoutingMulticastRendezvousPoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchRoutingMulticastRendezvousPoint`: GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RendezvousPointsAPI.UpdateNetworkSwitchRoutingMulticastRendezvousPoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rendezvousPointId** | **string** | Rendezvous point ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchRoutingMulticastRendezvousPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createNetworkSwitchRoutingMulticastRendezvousPointRequest** | [**CreateNetworkSwitchRoutingMulticastRendezvousPointRequest**](CreateNetworkSwitchRoutingMulticastRendezvousPointRequest.md) |  | 

### Return type

[**GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner**](GetNetworkSwitchRoutingMulticastRendezvousPoints200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

